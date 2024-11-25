package data

import (
	"database/sql"

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
	Delete(id int64) error
}

type SharedEntryModelDB struct {
	DB *sql.DB
}

func ValidateIds(v *validator.Validator, sharedBy, sharedWith int64) {
	v.Check(sharedBy != sharedWith, "shared_with", "cannot share with yourself")
}

func (m *SharedEntryModelDB) Insert(entryID, sharedBy, sharedWith int64) error {
	stmt := `INSERT INTO shared_entries (entry_id, shared_by, shared_with) VALUES (?, ?, ?)`
	_, err := m.DB.Exec(stmt, entryID, sharedBy, sharedWith)
	return err
}

func (m *SharedEntryModelDB) Get(id int64) (*SharedEntry, error) {
	stmt := `SELECT id, entry_id, shared_by, shared_with, created_at
		FROM shared_entries WHERE id = ?`
	sharedEntry := &SharedEntry{}
	err := m.DB.QueryRow(stmt, id).Scan(&sharedEntry.ID, &sharedEntry.EntryID, &sharedEntry.SharedBy, &sharedEntry.SharedWith, &sharedEntry.CreatedAt)
	if err != nil {
		return nil, err
	}
	return sharedEntry, nil
}

func (m *SharedEntryModelDB) GetBySharedWithID(sharedWith int64) ([]*SharedEntry, error) {
	stmt := `SELECT id, entry_id, shared_by, shared_with, created_at
		FROM shared_entries WHERE shared_with = ?`
	rows, err := m.DB.Query(stmt, sharedWith)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sharedEntries := []*SharedEntry{}
	for rows.Next() {
		sharedEntry := &SharedEntry{}
		err := rows.Scan(&sharedEntry.ID, &sharedEntry.EntryID, &sharedEntry.SharedBy, &sharedEntry.SharedWith, &sharedEntry.CreatedAt)
		if err != nil {
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
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sharedEntries := []*SharedEntry{}
	for rows.Next() {
		sharedEntry := &SharedEntry{}
		err := rows.Scan(&sharedEntry.ID, &sharedEntry.EntryID, &sharedEntry.SharedBy, &sharedEntry.SharedWith, &sharedEntry.CreatedAt)
		if err != nil {
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
	stmt := `UPDATE shared_entries SET entry_id = ?, shared_by = ?, shared_with = ?
		WHERE id = ?`
	_, err := m.DB.Exec(stmt, sharedEntry.EntryID, sharedEntry.SharedBy, sharedEntry.SharedWith, sharedEntry.ID)
	return err
}

func (m *SharedEntryModelDB) Delete(id int64) error {
	stmt := `DELETE FROM shared_entries WHERE id = ?`
	_, err := m.DB.Exec(stmt, id)
	return err
}
