package main

import (
	"database/sql"
	"errors"
	"flag"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		path string
	}
	cors struct {
		trustedOrigins []string
	}
	auth struct {
		expireTime int
	}
}

type application struct {
	config config
	logger *slog.Logger
	db     *sql.DB
	models data.Models
	wg     sync.WaitGroup
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	var cfg config

	err := godotenv.Load()
	if err != nil {
		logger.Warn("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		port = 4000
	}
	cfg.port = port
	cfg.env = os.Getenv("API_ENV")
	cfg.db.path = os.Getenv("API_DB_PATH")
	cfg.cors.trustedOrigins = strings.Fields(os.Getenv("CORS_TRUSTED_ORIGINS"))
	cfg.auth.expireTime, err = strconv.Atoi(os.Getenv("AUTH_EXPIRE_TIME"))
	if err != nil {
		cfg.auth.expireTime = 14
	}

	flag.IntVar(&cfg.port, "port", cfg.port, "API server port")
	flag.StringVar(&cfg.env, "env", cfg.env, "Environment (development|production)")
	flag.StringVar(&cfg.db.path, "db-path", cfg.db.path, "Path to SQLite database file")
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		if val != "" {
			cfg.cors.trustedOrigins = strings.Fields(val)
		}
		return nil
	})
	flag.IntVar(&cfg.auth.expireTime, "auth-expire-time", cfg.auth.expireTime, "Auth token expire time in days")
	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		db:     db,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.db.path)
	if err != nil {
		return nil, err
	}
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "sqlite", driver)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
