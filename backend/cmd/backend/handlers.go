package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/validator"
	tmdb "github.com/cyruzin/golang-tmdb"
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
			app.alreadyExistsResponse(w, r, v.Errors)
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

	w.WriteHeader(http.StatusOK)
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

type movieSearchResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Popularity  float32 `json:"popularity"`
	VoteAverage float32 `json:"vote_average"`
	Thumbnail   string  `json:"thumbnail"`
}

func (app *application) searchMoviesByTitleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		app.badRequestResponse(w, r, errors.New("missing q query parameter"))
		return
	}

	resp, err := app.tmdb.GetSearchMovies(query, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var movies []movieSearchResponse
	for _, movie := range resp.Results {
		movies = append(movies, movieSearchResponse{
			ID:          int(movie.ID),
			Title:       movie.Title,
			ReleaseDate: movie.ReleaseDate,
			Popularity:  movie.Popularity,
			VoteAverage: movie.VoteAverage,
			Thumbnail:   tmdb.GetImageURL(movie.PosterPath, "w92"),
		})
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movies": movies}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type movieResponse struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	ReleaseDate string   `json:"release_date"`
	Overview    string   `json:"overview"`
	Popularity  float32  `json:"popularity"`
	Thumbnail   string   `json:"thumbnail"`
	Genres      []string `json:"genres"`
	Runtime     int      `json:"runtime"`
}

func (app *application) getMoviesByIMDbHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.badRequestResponse(w, r, errors.New("missing id query parameter"))
		return
	}

	results, err := app.tmdb.GetFindByID(id, map[string]string{"external_source": "imdb_id"})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if len(results.MovieResults) == 0 {
		app.notFoundResponse(w, r)
		return
	}

	movieDetails, err := app.tmdb.GetMovieDetails(int(results.MovieResults[0].ID), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var genres []string
	for _, genre := range movieDetails.Genres {
		genres = append(genres, genre.Name)
	}

	movie := movieResponse{
		ID:          int(movieDetails.ID),
		Title:       movieDetails.Title,
		ReleaseDate: movieDetails.ReleaseDate,
		Overview:    movieDetails.Overview,
		Popularity:  movieDetails.Popularity,
		Thumbnail:   tmdb.GetImageURL(movieDetails.PosterPath, "w92"),
		Genres:      genres,
		Runtime:     movieDetails.Runtime,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movieDetails, err := app.tmdb.GetMovieDetails(int(id), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var genres []string
	for _, genre := range movieDetails.Genres {
		genres = append(genres, genre.Name)
	}

	movie := movieResponse{
		ID:          int(movieDetails.ID),
		Title:       movieDetails.Title,
		ReleaseDate: movieDetails.ReleaseDate,
		Overview:    movieDetails.Overview,
		Popularity:  movieDetails.Popularity,
		Thumbnail:   tmdb.GetImageURL(movieDetails.PosterPath, "w92"),
		Genres:      genres,
		Runtime:     movieDetails.Runtime,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type tvShowSearchResponse struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	FirstAirDate string  `json:"first_air_date"`
	Popularity   float32 `json:"popularity"`
	VoteAverage  float32 `json:"vote_average"`
	Thumbnail    string  `json:"thumbnail"`
}

func (app *application) searchTVShowsByTitleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		app.badRequestResponse(w, r, errors.New("missing q query parameter"))
		return
	}

	resp, err := app.tmdb.GetSearchTVShow(query, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var tvShows []tvShowSearchResponse
	for _, tvShow := range resp.Results {
		tvShows = append(tvShows, tvShowSearchResponse{
			ID:           int(tvShow.ID),
			Title:        tvShow.Name,
			FirstAirDate: tvShow.FirstAirDate,
			Popularity:   tvShow.Popularity,
			VoteAverage:  tvShow.VoteAverage,
			Thumbnail:    tmdb.GetImageURL(tvShow.PosterPath, "w92"),
		})
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"tvshows": tvShows}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type tvShowResponse struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	FirstAirDate     string   `json:"first_air_date"`
	Overview         string   `json:"overview"`
	Popularity       float32  `json:"popularity"`
	Thumbnail        string   `json:"thumbnail"`
	Genres           []string `json:"genres"`
	NumberOfSeasons  int      `json:"number_of_seasons"`
	NumberOfEpisodes int      `json:"number_of_episodes"`
}

func (app *application) getTVShowsByIMDbHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.badRequestResponse(w, r, errors.New("missing id query parameter"))
		return
	}

	results, err := app.tmdb.GetFindByID(id, map[string]string{"external_source": "imdb_id"})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if len(results.TvResults) == 0 {
		app.notFoundResponse(w, r)
		return
	}

	tvShowDetails, err := app.tmdb.GetTVDetails(int(results.TvResults[0].ID), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var genres []string
	for _, genre := range tvShowDetails.Genres {
		genres = append(genres, genre.Name)
	}

	tvShow := tvShowResponse{
		ID:               int(tvShowDetails.ID),
		Title:            tvShowDetails.Name,
		FirstAirDate:     tvShowDetails.FirstAirDate,
		Overview:         tvShowDetails.Overview,
		Popularity:       tvShowDetails.Popularity,
		Thumbnail:        tmdb.GetImageURL(tvShowDetails.PosterPath, "w92"),
		Genres:           genres,
		NumberOfSeasons:  tvShowDetails.NumberOfSeasons,
		NumberOfEpisodes: tvShowDetails.NumberOfEpisodes,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"tvshow": tvShow}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getTVShowHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	tvShowDetails, err := app.tmdb.GetTVDetails(int(id), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var genres []string
	for _, genre := range tvShowDetails.Genres {
		genres = append(genres, genre.Name)
	}

	tvShow := tvShowResponse{
		ID:               int(tvShowDetails.ID),
		Title:            tvShowDetails.Name,
		FirstAirDate:     tvShowDetails.FirstAirDate,
		Overview:         tvShowDetails.Overview,
		Popularity:       tvShowDetails.Popularity,
		Thumbnail:        tmdb.GetImageURL(tvShowDetails.PosterPath, "w92"),
		Genres:           genres,
		NumberOfSeasons:  tvShowDetails.NumberOfSeasons,
		NumberOfEpisodes: tvShowDetails.NumberOfEpisodes,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"tvshow": tvShow}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type bookSearchResponse struct {
	ID        string `json:"id"`
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	PageCount int    `json:"page_count"`
	Thumbnail string `json:"thumbnail"`
}

func (app *application) searchBooksByTitleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		app.badRequestResponse(w, r, errors.New("missing q query parameter"))
		return
	}

	resp, err := app.books.Volumes.List(query).PrintType("BOOKS").Do()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var books []bookSearchResponse
	for _, book := range resp.Items {
		var isbn string
		if len(book.VolumeInfo.IndustryIdentifiers) > 0 {
			isbn = book.VolumeInfo.IndustryIdentifiers[0].Identifier
		}

		var thumbnail string
		if book.VolumeInfo.ImageLinks != nil {
			thumbnail = book.VolumeInfo.ImageLinks.Thumbnail
		}

		books = append(books, bookSearchResponse{
			ID:        book.Id,
			ISBN:      isbn,
			Title:     book.VolumeInfo.Title,
			Author:    strings.Join(book.VolumeInfo.Authors, ", "),
			PageCount: int(book.VolumeInfo.PageCount),
			Thumbnail: thumbnail,
		})
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type bookResponse struct {
	ID            string   `json:"id"`
	ISBN          string   `json:"isbn"`
	Title         string   `json:"title"`
	Author        string   `json:"author"`
	Description   string   `json:"description"`
	PageCount     int      `json:"page_count"`
	Thumbnail     string   `json:"thumbnail"`
	Categories    []string `json:"categories"`
	PublishedDate string   `json:"published_date"`
	Publisher     string   `json:"publisher"`
	Language      string   `json:"language"`
}

func (app *application) searchBooksByISBNHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	if query == "" {
		app.badRequestResponse(w, r, errors.New("missing id query parameter"))
		return
	}

	queryParam := "isbn:" + query

	resp, err := app.books.Volumes.List(queryParam).Do()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if len(resp.Items) == 0 {
		app.notFoundResponse(w, r)
		return
	}

	result := resp.Items[0]
	thumbnail := ""
	if result.VolumeInfo.ImageLinks != nil {
		thumbnail = result.VolumeInfo.ImageLinks.Thumbnail
	}

	var book bookResponse
	book = bookResponse{
		ID:            result.Id,
		ISBN:          query,
		Title:         result.VolumeInfo.Title,
		Author:        strings.Join(result.VolumeInfo.Authors, ", "),
		Description:   result.VolumeInfo.Description,
		PageCount:     int(result.VolumeInfo.PageCount),
		Thumbnail:     thumbnail,
		Categories:    result.VolumeInfo.Categories,
		PublishedDate: result.VolumeInfo.PublishedDate,
		Publisher:     result.VolumeInfo.Publisher,
		Language:      result.VolumeInfo.Language,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readGoogleIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	resp, err := app.books.Volumes.Get(id).Do()
	if err != nil {
		switch {
		case err.Error() == "googleapi: Error 404: The volume ID could not be found., notFound":
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var book bookResponse
	isbn := ""
	if len(resp.VolumeInfo.IndustryIdentifiers) > 0 {
		isbn = resp.VolumeInfo.IndustryIdentifiers[0].Identifier
	}
	thumbnail := ""
	if resp.VolumeInfo.ImageLinks != nil {
		thumbnail = resp.VolumeInfo.ImageLinks.Thumbnail
	}

	book = bookResponse{
		ID:            resp.Id,
		ISBN:          isbn,
		Title:         resp.VolumeInfo.Title,
		Author:        strings.Join(resp.VolumeInfo.Authors, ", "),
		Description:   resp.VolumeInfo.Description,
		PageCount:     int(resp.VolumeInfo.PageCount),
		Thumbnail:     thumbnail,
		Categories:    resp.VolumeInfo.Categories,
		PublishedDate: resp.VolumeInfo.PublishedDate,
		Publisher:     resp.VolumeInfo.Publisher,
		Language:      resp.VolumeInfo.Language,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type youtubeVideoResponse struct {
	VideoID   string `json:"video_id"`
	VideoURL  string `json:"video_url"`
	Title     string `json:"title"`
	Channel   string `json:"channel"`
	Thumbnail string `json:"thumbnail"`
	Duration  int    `json:"duration"`
	Published string `json:"published"`
}

func (app *application) getYoutubeVideoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.badRequestResponse(w, r, errors.New("missing id query parameter"))
		return
	}

	resp, err := app.youtube.Videos.List([]string{"snippet", "contentDetails"}).Id(id).Do()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if len(resp.Items) == 0 {
		app.notFoundResponse(w, r)
		return
	}

	video := resp.Items[0]
	thumbnail := video.Snippet.Thumbnails.Default.Url
	duration, err := app.parseDuration(video.ContentDetails.Duration)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	ytVideo := youtubeVideoResponse{
		VideoID:   video.Id,
		VideoURL:  "https://www.youtube.com/watch?v=" + video.Id,
		Title:     video.Snippet.Title,
		Channel:   video.Snippet.ChannelTitle,
		Thumbnail: thumbnail,
		Duration:  duration,
		Published: video.Snippet.PublishedAt,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"video": ytVideo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) shareMediaEntryHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	var input struct {
		MediaEntry int    `json:"media_entry"`
		ShareWith  string `json:"share_with"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.logger.Error(err.Error())
		app.badRequestResponse(w, r, err)
		return
	}

	mediaEntry, err := app.models.MediaEntries.Get(int64(input.MediaEntry), user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	sharedWith, err := app.models.Users.GetByEmail(input.ShareWith)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	validator := validator.New()
	data.ValidateIds(validator, user.ID, int64(sharedWith.ID))
	if !validator.Valid() {
		app.failedValidationResponse(w, r, validator.Errors)
		return
	}

	err = app.models.SharedEntries.Insert(mediaEntry.ID, user.ID, sharedWith.ID)
	if err != nil {
		switch {
		case err.Error() == "constraint failed: UNIQUE constraint failed: shared_entries.entry_id, shared_entries.shared_by, shared_entries.shared_with (2067)":
			app.alreadyExistsResponse(w, r, map[string]string{"shared_with": "media entry already shared with this user"})
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"message": "media entry shared successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type sharedEntryResponse struct {
	ID         int64           `json:"id"`
	SharedBy   string          `json:"shared_by"`
	SharedWith string          `json:"shared_with"`
	MediaEntry data.MediaEntry `json:"media_entry"`
	CreatedAt  string          `json:"created_at"`
}

func (app *application) listSharedEntriesHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	sharedEntries, err := app.models.SharedEntries.GetBySharedWithID(user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var sharedEntriesResponse []sharedEntryResponse
	for _, sharedEntry := range sharedEntries {
		sharedBy, err := app.models.Users.Get(sharedEntry.SharedBy)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		mediaEntry, err := app.models.MediaEntries.Get(sharedEntry.EntryID, sharedEntry.SharedBy)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		sharedEntriesResponse = append(sharedEntriesResponse, sharedEntryResponse{
			ID:         sharedEntry.ID,
			SharedBy:   sharedBy.Email,
			SharedWith: user.Email,
			MediaEntry: *mediaEntry,
			CreatedAt:  sharedEntry.CreatedAt,
		})
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sharedEntries": sharedEntriesResponse}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteSharedEntryHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.SharedEntries.Delete(id, user.ID)
	if err != nil {
		switch {
		case err.Error() == "record not found":
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "shared entry deleted successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
