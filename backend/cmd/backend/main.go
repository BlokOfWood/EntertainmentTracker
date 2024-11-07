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
	_ "modernc.org/sqlite"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		db_path string
	}
	cors struct {
		trustedOrigins []string
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
	var cfg config
	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		port = 4000
	}
	flag.IntVar(&cfg.port, "port", port, "API server port")
	flag.StringVar(&cfg.env, "env", os.Getenv("API_ENV"), "Environment (development|production)")
	flag.StringVar(&cfg.db.db_path, "db-path", os.Getenv("API_DB_PATH"), "Path to SQLite database file")
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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
	db, err := sql.Open("sqlite", cfg.db.db_path)
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
