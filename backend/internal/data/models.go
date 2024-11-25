package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Users         UserModel
	Tokens        TokenModel
	MediaEntries  MediaEntryModel
	SharedEntries SharedEntryModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:         &UserModelDB{DB: db},
		Tokens:        &TokenModelDB{DB: db},
		MediaEntries:  &MediaEntryModelDB{DB: db},
		SharedEntries: &SharedEntryModelDB{DB: db},
	}
}
