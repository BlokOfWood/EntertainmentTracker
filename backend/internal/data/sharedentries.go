package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/validator"
)

type SharedEntry struct {
	ID         int64  `json:"id"`
	EntryID    int64  `json:"entry_id"`
	SharedBy   int64  `json:"shared_by"`
	SharedWith int64  `json:"shared_with"`
	CreatedAt  string `json:"created_at"`
}

type SharedEntryModel interface {
	Insert(entryID, sharedBy, sharedWith int64) error
	Get(id int64) (*SharedEntry, error)
	GetBySharedWithID(sharedWith int64) ([]*SharedEntry, error)
	GetAll() ([]*SharedEntry, error)
	Update(sharedEntry *SharedEntry) error
	Delete(id, userID int64) error
}

type SharedEntryModelDB struct {
	DB *sql.DB
}

func ValidateIds(v *validator.Validator, sharedBy, sharedWith int64) {
	v.Check(sharedBy != sharedWith, "shared_with", "cannot share with yourself")
}

func (m *SharedEntryModelDB) Insert(entryID, sharedBy, sharedWith int64) error {
	stmt := `INSERT INTO shared_entries (entry_id, shared_by, shared_with) VALUES ($1, $2, $3) RETURNING id, created_at`
	args := []interface{}{entryID, sharedBy, sharedWith}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.QueryContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return nil
}

func (m *SharedEntryModelDB) Get(id int64) (*SharedEntry, error) {
	stmt := `SELECT id, entry_id, shared_by, shared_with, created_at
		FROM shared_entries WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sharedEntry := &SharedEntry{}
	err := m.DB.QueryRowContext(ctx, stmt, id).Scan(&sharedEntry.ID, &sharedEntry.EntryID, &sharedEntry.SharedBy, &sharedEntry.SharedWith, &sharedEntry.CreatedAt)
	if err != nil {
		return nil, err
	}
	return sharedEntry, nil
}

func (m *SharedEntryModelDB) GetBySharedWithID(sharedWith int64) ([]*SharedEntry, error) {
	stmt := `SELECT id, entry_id, shared_by, shared_with, created_at
		FROM shared_entries WHERE shared_with = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, stmt, sharedWith)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sharedEntries := []*SharedEntry{}

	for rows.Next() {
		sharedEntry := &SharedEntry{}
		err := rows.Scan(
			&sharedEntry.ID,
			&sharedEntry.EntryID,
			&sharedEntry.SharedBy,
			&sharedEntry.SharedWith,
			&sharedEntry.CreatedAt,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrRecordNotFound
			}
			return nil, err
		}
		sharedEntries = append(sharedEntries, sharedEntry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sharedEntries, nil
}

func (m *SharedEntryModelDB) GetAll() ([]*SharedEntry, error) {
	stmt := `SELECT id, entry_id, shared_by, shared_with, created_at
		FROM shared_entries`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sharedEntries := []*SharedEntry{}

	for rows.Next() {
		sharedEntry := &SharedEntry{}
		err := rows.Scan(
			&sharedEntry.ID,
			&sharedEntry.EntryID,
			&sharedEntry.SharedBy,
			&sharedEntry.SharedWith,
			&sharedEntry.CreatedAt,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrRecordNotFound
			}
			return nil, err
		}
		sharedEntries = append(sharedEntries, sharedEntry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sharedEntries, nil
}

func (m *SharedEntryModelDB) Update(sharedEntry *SharedEntry) error {
	stmt := `UPDATE shared_entries SET entry_id = $1, shared_by = $2, shared_with = $3
		WHERE id = $4`
	args := []interface{}{sharedEntry.EntryID, sharedEntry.SharedBy, sharedEntry.SharedWith, sharedEntry.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.QueryContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return nil
}

func (m *SharedEntryModelDB) Delete(id, userID int64) error {
	stmt := `DELETE FROM shared_entries WHERE entry_id = $1 and shared_with = $2`
	args := []interface{}{id, userID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.QueryContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return nil
}
