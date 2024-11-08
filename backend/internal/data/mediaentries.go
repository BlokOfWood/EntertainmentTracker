package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/validator"
)

type MediaEntry struct {
	ID              int64  `json:"id"`
	UserID          int64  `json:"user_id"`
	ThirdPartyID    string `json:"third_party_id"`
	Title           string `json:"title"`
	Type            Type   `json:"type"`
	Status          Status `json:"status"`
	CurrentProgress int    `json:"current_progress"`
	TargetProgress  int    `json:"target_progress"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Version         int    `json:"version"`
}

type Type string

const (
	TypeMovie   Type = "movie"
	TypeBook    Type = "book"
	TypeShow    Type = "show"
	TypeYoutube Type = "youtube"
)

type Status string

const (
	StatusNotStarted Status = "not_started"
	StatusWatching   Status = "watching"
	StatusCompleted  Status = "completed"
)

type MediaEntryModel struct {
	DB *sql.DB
}

func ValidateType(v *validator.Validator, mediaType string) {
	v.Check(mediaType != "", "type", "must be provided")
	v.Check(mediaType == string(TypeMovie) || mediaType == string(TypeBook) || mediaType == string(TypeShow) || mediaType == string(TypeYoutube), "type", "must be a valid type (movie, book, show, youtube)")
}

func ValidateStatus(v *validator.Validator, status string) {
	v.Check(status != "", "status", "must be provided")
	v.Check(status == string(StatusNotStarted) || status == string(StatusWatching) || status == string(StatusCompleted), "status", "must be a valid status (not_started, watching, completed)")
}

func ValidateProgress(v *validator.Validator, progress int) {
	v.Check(progress >= 0, "progress", "must be 0 or greater")
}

func ValidateMediaEntry(v *validator.Validator, mediaEntry *MediaEntry) {
	v.Check(mediaEntry.Title != "", "title", "must be provided")
	v.Check(mediaEntry.Type != "", "type", "must be provided")
	ValidateStatus(v, string(mediaEntry.Status))
	ValidateProgress(v, mediaEntry.CurrentProgress)
	ValidateProgress(v, mediaEntry.TargetProgress)
	v.Check(mediaEntry.CurrentProgress <= mediaEntry.TargetProgress, "current_progress", "must be less than or equal to target progress")
	v.Check(mediaEntry.Status != StatusCompleted || mediaEntry.CurrentProgress == mediaEntry.TargetProgress, "current_progress", "must be equal to target progress if status is completed")
	v.Check(mediaEntry.Status != StatusNotStarted || mediaEntry.CurrentProgress == 0, "current_progress", "must be 0 if status is not started")
}

func (m MediaEntryModel) Insert(mediaEntry *MediaEntry) error {
	query := `
		INSERT INTO media_entries (user_id, third_party_id, title, type, status, current_progress, target_progress)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at, version`

	args := []interface{}{
		mediaEntry.UserID,
		mediaEntry.ThirdPartyID,
		mediaEntry.Title,
		mediaEntry.Type,
		mediaEntry.Status,
		mediaEntry.CurrentProgress,
		mediaEntry.TargetProgress,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&mediaEntry.ID, &mediaEntry.CreatedAt, &mediaEntry.UpdatedAt, &mediaEntry.Version)
	if err != nil {
		return err
	}

	return nil
}

func (m MediaEntryModel) Get(id int64, userId int64) (*MediaEntry, error) {
	query := `
		SELECT id, user_id, third_party_id, title, type, status, current_progress, target_progress, created_at, updated_at, version
		FROM media_entries
		WHERE id = $1 AND user_id = $2`

	var mediaEntry MediaEntry

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id, userId).Scan(
		&mediaEntry.ID,
		&mediaEntry.UserID,
		&mediaEntry.ThirdPartyID,
		&mediaEntry.Title,
		&mediaEntry.Type,
		&mediaEntry.Status,
		&mediaEntry.CurrentProgress,
		&mediaEntry.TargetProgress,
		&mediaEntry.CreatedAt,
		&mediaEntry.UpdatedAt,
		&mediaEntry.Version,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &mediaEntry, nil
}

func (m MediaEntryModel) GetAll(userID int64) ([]*MediaEntry, error) {
	query := `
		SELECT id, user_id, third_party_id, title, type, status, current_progress, target_progress, created_at, updated_at, version
		FROM media_entries
		WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mediaEntries := []*MediaEntry{}

	for rows.Next() {
		var mediaEntry MediaEntry

		err := rows.Scan(
			&mediaEntry.ID,
			&mediaEntry.UserID,
			&mediaEntry.ThirdPartyID,
			&mediaEntry.Title,
			&mediaEntry.Type,
			&mediaEntry.Status,
			&mediaEntry.CurrentProgress,
			&mediaEntry.TargetProgress,
			&mediaEntry.CreatedAt,
			&mediaEntry.UpdatedAt,
			&mediaEntry.Version,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrRecordNotFound
			}
			return nil, err
		}

		mediaEntries = append(mediaEntries, &mediaEntry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return mediaEntries, nil
}

func (m MediaEntryModel) Update(mediaEntry *MediaEntry) error {
	query := `
		UPDATE media_entries
		SET title = $1, type = $2, status = $3, current_progress = $4, target_progress = $5, version = version + 1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $6 AND version = $7
		RETURNING version, updated_at`

	args := []interface{}{
		mediaEntry.Title,
		mediaEntry.Type,
		mediaEntry.Status,
		mediaEntry.CurrentProgress,
		mediaEntry.TargetProgress,
		mediaEntry.ID,
		mediaEntry.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&mediaEntry.Version, &mediaEntry.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m MediaEntryModel) Delete(id int64, userId int64) error {
	query := `
		DELETE FROM media_entries
		WHERE id = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id, userId)
	if err != nil {
		return err
	}

	return nil
}
