package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/validator"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]interface{}{
			"environment":          app.config.env,
			"cors_trusted_origins": app.config.cors.trustedOrigins,
			"version":              version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.logger.Error(err.Error())
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:  input.Name,
		Email: input.Email,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "This email address already in use")
			app.emailAlreadyExistsResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	err := app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateEmail(v, input.Email)
	data.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidCredentialsResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	hours := app.config.auth.expireTime * 24
	token, err := app.models.Tokens.New(user.ID, time.Duration(hours)*time.Hour)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) invalidateAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	err := app.models.Tokens.DeleteAllForUser(user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

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
