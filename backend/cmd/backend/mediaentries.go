package main

import (
	"net/http"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/validator"
)

func (app *application) listMediaEntriesHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	entries, err := app.models.MediaEntries.GetAll(user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"mediaEntries": entries}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showMediaEntryHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	entry, err := app.models.MediaEntries.Get(id, user.ID)
	if err != nil {
		switch {
		case err.Error() == "record not found":
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"mediaEntry": entry}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createMediaEntryHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	var input struct {
		ThirdPartyID    string `json:"third_party_id"`
		Title           string `json:"title"`
		Type            string `json:"type"`
		Status          string `json:"status"`
		CurrentProgress int    `json:"current_progress"`
		TargetProgress  int    `json:"target_progress"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateStatus(v, input.Status)
	data.ValidateType(v, input.Type)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
	}

	status := data.Status(input.Status)
	mediaType := data.Type(input.Type)

	mediaEntry := &data.MediaEntry{
		UserID:          user.ID,
		ThirdPartyID:    input.ThirdPartyID,
		Title:           input.Title,
		Type:            mediaType,
		Status:          status,
		CurrentProgress: input.CurrentProgress,
		TargetProgress:  input.TargetProgress,
	}

	if data.ValidateMediaEntry(v, mediaEntry); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.MediaEntries.Insert(mediaEntry)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"mediaEntry": mediaEntry}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateMediaEntryHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	entry, err := app.models.MediaEntries.Get(id, user.ID)
	if err != nil {
		switch {
		case err.Error() == "record not found":
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title           *string `json:"title"`
		Type            *string `json:"type"`
		Status          *string `json:"status"`
		CurrentProgress *int    `json:"current_progress"`
		TargetProgress  *int    `json:"target_progress"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if input.Title != nil {
		entry.Title = *input.Title
	}

	if input.Type != nil {
		data.ValidateType(v, *input.Type)
		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}
		entry.Type = data.Type(*input.Type)
	}

	if input.Status != nil {
		data.ValidateStatus(v, *input.Status)
		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}
		entry.Status = data.Status(*input.Status)
	}

	if input.CurrentProgress != nil {
		data.ValidateProgress(v, *input.CurrentProgress)
		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}
		entry.CurrentProgress = *input.CurrentProgress
	}

	if input.TargetProgress != nil {
		data.ValidateProgress(v, *input.TargetProgress)
		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}
		entry.TargetProgress = *input.TargetProgress
	}

	if data.ValidateMediaEntry(v, entry); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.MediaEntries.Update(entry)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"mediaEntry": entry}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteMediaEntryHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.MediaEntries.Delete(id, user.ID)
	if err != nil {
		switch {
		case err.Error() == "record not found":
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "media entry deleted successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
