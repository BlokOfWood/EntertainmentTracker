# Backend telepítési útmutató

## Előfeltételek
- [Go 1.23](https://go.dev/dl)
- [SQLite 3](https://www.sqlite.org/download.html)

## Telepítés
1. Klónozd le a repót:
    ```bash
    git clone https://github.com/BlokOfWood/EntertainmentTracker.git
    ```
2. Navigálj a backend mappába:
    ```bash
    cd backend
    ```
3. Telepítsd a függőségeket:
    ```bash
    go mod download
    ```
4. Hozd létre a .env fájlt a következő tartalommal:
    ```env
    API_PORT=6969 # Port, amin a szerver futni fog
    API_ENV="development" # Környezet, amin a szerver futni fog
    API_DB_PATH="db.db" # Adatbázis fájl elérési útja
    CORS_TRUSTED_ORIGINS="http://localhost:3000 http://localhost:3001" # Megbízhat a CORS originok
    AUTH_EXPIRE_TIME=14 # Auth token lejárati ideje napokban
    ```
5. Indítsd el a szervert:
    ```bash
    go run main.go
    ```
6. A szerver elérhető a `http://localhost:<API_PORT>` címen.

# API dokumentáció

## Áttekintés
Base URL: `http://localhost:<API_PORT>/v1`
### Authentication
Az API token alapú autentikációt használ. Az összes kérésnek tartalmaznia kell, kivéve a `/healthcheck`, `/user/login` és a `/user/register` endpointokat.

```http
Authorization: Bearer <token>
```

## Endpoints
### 1. Healthcheck
### Endpoint
```http
GET /healthcheck
```

### Description
Ellenőrzi, hogy a szerver elérhető-e és rendszer információkat ad vissza.

### Parameters
Nincsenek

### Request Example
```http
GET /healthcheck
```
### Response

Status: `200 OK`

### Example Response:
```json
{
    "status": "available",
    "system_info": {
        "cors_trusted_origins": [
            "http://localhost:3000",
            "http://localhost:3001"
        ],
        "environment": "development",
        "version": "1.0.0"
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
500|Internal Server Error|An error occurred on the server

### 2. User Registration
### Endpoint
```http
POST /user/register
```

### Description
Regisztrál egy új felhasználót.

### Parameters
Name|Type|Description
--- | --- | ---
username|string|Felhasználónév
email|string|Email cím
password|string|Jelszó

### Request Example
```http
POST /user/register
Content-Type: application/json

{
    "username": "test",
    "email": "test@test.com",
    "password": "test"
}
```

### Response

Status: `201 Created`

### Example Response:
```json
{
    "user": {
        "id": 2,
        "created_at": "2024-11-14T16:36:15Z",
        "name": "John Doe",
        "email": "john.doe@example.com"
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
409|Conflict|Az email cím már foglalt
422|Unprocessable Entity|Validációs hiba
500|Internal Server Error|An error occurred on the server

### 3. User Login
### Endpoint
```http
POST /users/login
```

### Description
Bejelentkezik egy felhasználóval.

### Parameters
Name|Type|Description
--- | --- | ---
email|string|Email cím
password|string|Jelszó

### Request Example
```http
POST /users/login
Content-Type: application/json

{
    "email": "test@test.com",
    "password": "test"
}
```

### Response

Status: `201 Created`

### Example Response:
```json
{
    "authentication_token": {
        "token": "GF433ERQINPBI24A2OBH5MZ5LI",       
        "expiry": "2024-11-28T17:41:42.6475501+01:00"
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
401|Unauthorized|Hibás email cím vagy jelszó
422|Unprocessable Entity|Validációs hiba
500|Internal Server Error|An error occurred on the server

### 4. User Logout
### Endpoint
```http
GET /users/logout
```

### Description
Kijelentkezik a felhasználó.

### Parameters
Nincsenek

### Request Example
```http
GET /users/logout
```

### Response

Status: `200 OK`

### Example Response:
No content

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
500|Internal Server Error|An error occurred on the server

### 5. User Profile
### Endpoint
```http
GET /users/me
```

### Description
Visszaadja a bejelentkezett felhasználó adatait.

### Parameters
Nincsenek

### Request Example
```http
GET /users/me
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "user": {       
        "id": 1,
        "created_at": "2024-11-14T14:06:37Z",
        "name": "John Doe",
        "email": "john.doe@example.com"
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
401|Unauthorized|Nincs bejelentkezve
500|Internal Server Error|An error occurred on the server

### 6. Media Entry List

### Endpoint
```http
GET /mediaentries
```

### Description
Visszaadja az összes média bejegyzést.

### Parameters
Nincsenek

### Request Example
```http
GET /mediaentries
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "mediaEntries": [
        {        
            "id": 1,
            "user_id": 1,
            "third_party_id": "12345",
            "title": "Updated Media Title",
            "type": "movie",
            "status": "completed",
            "current_progress": 100,
            "target_progress": 100,
            "created_at": "2024-11-14T14:08:33Z",
            "updated_at": "2024-11-14T14:09:01Z",
            "version": 2
        }
    ]
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
500|Internal Server Error|An error occurred on the server

### 7. Media Entry Create

### Endpoint
```http
POST /mediaentries
```

### Description
Létrehoz egy új média bejegyzést.

### Parameters
Name|Type|Description
--- | --- | ---
third_party_id|string|Külső szolgáltató azonosító
title|string|Cím
type|string|Típus (movie, book, show, youtube)
status|string|Státusz (not_started, watching, completed)
current_progress|number|Jelenlegi állapot
target_progress|number|Célállapot

### Request Example
```http
POST /mediaentries
Content-Type: application/json

{
    "third_party_id": "12345",
    "title": "Media Title",
    "type": "movie",
    "status": "not_started",
    "current_progress": 0,
    "target_progress": 100
}
```

### Response

Status: `201 Created`

### Example Response:
```json
{
    "mediaEntry": {
        "id": 1,
        "user_id": 1,
        "third_party_id": "12345",
        "title": "Media Title",
        "type": "movie",
        "status": "not_started",
        "current_progress": 0,
        "target_progress": 100,
        "created_at": "2024-11-14T14:08:33Z",
        "updated_at": "2024-11-14T14:08:33Z",
        "version": 1
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
422|Unprocessable Entity|Validációs hiba
500|Internal Server Error|An error occurred on the server

### 8. Media Entry Detail

### Endpoint
```http
GET /mediaentries/:id
```

### Description
Visszaadja a média bejegyzés részleteit.

### Parameters
Nincsenek

### Request Example
```http
GET /mediaentries/1
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "mediaEntry": {      
        "id": 1,     
        "user_id": 1,
        "third_party_id": "12345",
        "title": "Media Title",
        "type": "movie",
        "status": "not_started",
        "current_progress": 0,
        "target_progress": 100,
        "created_at": "2024-11-14T14:08:33Z",
        "updated_at": "2024-11-14T14:08:33Z",
        "version": 1
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
404|Not Found|Nem található a bejegyzés
500|Internal Server Error|An error occurred on the server


### 9. Media Entry Update

### Endpoint
```http
PATCH /mediaentries/:id
```

### Description
Frissíti a média bejegyzést.

### Parameters
Name|Type|Description
--- | --- | ---
title|string|Cím
type|string|Típus (movie, book, show, youtube)
status|string|Státusz (not_started, watching, completed)
current_progress|number|Jelenlegi állapot
target_progress|number|Célállapot

### Request Example
```http
PATCH /mediaentries/1
Content-Type: application/json

{
    "title": "Updated Media Title",
    "type": "movie",
    "status": "completed",
    "current_progress": 100,
    "target_progress": 100
}
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "mediaEntry": {      
        "id": 1,     
        "user_id": 1,
        "third_party_id": "12345",
        "title": "Updated Media Title",
        "type": "movie",
        "status": "completed",
        "current_progress": 100,
        "target_progress": 100,
        "created_at": "2024-11-14T14:08:33Z",
        "updated_at": "2024-11-14T16:59:47Z",
        "version": 3
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
404|Not Found|Nem található a bejegyzés
422|Unprocessable Entity|Validációs hiba
500|Internal Server Error|An error occurred on the server

### 10. Media Entry Delete

### Endpoint
```http
DELETE /mediaentries/:id
```

### Description
Törli a média bejegyzést.

### Parameters
Nincsenek

### Request Example
```http
DELETE /mediaentries/1
```

### Response

Status: `200 OK`

### Example Response:
No content

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
404|Not Found|Nem található a bejegyzés
500|Internal Server Error|An error occurred on the server

### 11. Movie Search

### Endpoint
```http
GET /search/movies
```

### Description
Film keresés a TMDB API-ban.

### Parameters
Name|Type|Description
--- | --- | ---
q|string|Keresési kifejezés

### Request Example
```http
GET /search/movies?q=inception
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "movies": [
        {
            "id": 27205,
            "title": "Inception",
            "release_date": "2010-07-15",
            "popularity": 112.09,
            "thumbnail": "https://image.tmdb.org/t/p/w92/ljsZTbVsrQSqZgWeep2B1QiDKuh.jpg"
        },
        .
        .
        .
        {
            "id": 973484,
            "title": "Inception: Music from the Motion Picture",
            "release_date": "2010-12-07",
            "popularity": 1.463,
            "thumbnail": "https://image.tmdb.org/t/p/w92/7uM4DyRVAcgagvhZoWrkrqMPbqV.jpg"
        },
    ]
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
500|Internal Server Error|An error occurred on the server

### 12. TV Show Search

### Endpoint
```http
GET /search/tvshows
```

### Description
Sorozat keresés a TMDB API-ban.

### Parameters
Name|Type|Description
--- | --- | ---
q|string|Keresési kifejezés

### Request Example
```http
GET /search/tvshows?q=breaking+bad
```

### Response

Status: `200 OK`

### Example Response:
```json
{
        "tvshows": [
                {
                        "id": 1396,
                        "title": "Breaking Bad",
                        "first_air_date": "2008-01-20",
                        "popularity": 841.54,
                        "thumbnail": "https://image.tmdb.org/t/p/w92/ztkUQFLlC19CCMYHW9o1zWhJRNq.jpg"
                },
                {
                        "id": 232533,
                        "title": "Breaking Bad Fortune Teller",
                        "first_air_date": "2016-04-11",
                        "popularity": 7.101,
                        "thumbnail": "https://image.tmdb.org/t/p/w92/n8dKkbPkdEPFCFmjNvrUl1gUAYr.jpg"
                }
        ]
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
500|Internal Server Error|An error occurred on the server

### 13. Book Search

### Endpoint
```http
GET /search/books
```

### Description
Könyv keresés a Google Books API-ban.

### Parameters
Name|Type|Description
--- | --- | ---
q|string|Keresési kifejezés

### Request Example
```http
GET /search/books?q=1984
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "books": [
        {
            "id": "PpcZEAAAQBAJ",
            "isbn": "",
            "title": "1984",
            "author": "George Orwell",
            "page_count": 275,
            "thumbnail": "http://books.google.com/books/content?id=PpcZEAAAQBAJ\u0026printsec=frontcover\u0026img=1\u0026zoom=1\u0026edge=curl\u0026source=gbs_api"
        },
        .
        .
        .
        {
            "id": "sG8WAAAAIAAJ",
            "isbn": "STANFORD:36105126766430",
            "title": "Veterans' Administration Health Care Amendments of 1984",
            "author": "United States. Congress. Senate. Committee on Veterans' Affairs",
            "page_count": 718,
            "thumbnail": "http://books.google.com/books/content?id=sG8WAAAAIAAJ\u0026printsec=frontcover\u0026img=1\u0026zoom=1\u0026edge=curl\u0026source=gbs_api"
        }
    ]
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
500|Internal Server Error|An error occurred on the server

### 14. Find Movie by Imdb ID

### Endpoint
```http
GET /find/movie
```

### Description
Film keresése az IMDB API-ban.

### Parameters
Name|Type|Description
--- | --- | ---
id|string|IMDB azonosító

### Request Example
```http
GET /find/movie?id=tt1375666
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "movie": {
        "id": 27205,
        "title": "Inception",
        "release_date": "2010-07-15",
        "overview": "Cobb, a skilled thief who commits corporate espionage by infiltrating the subconscious of his targets is offered a chance to regain his old life as payment for a task considered to be impossible: \"inception\", the implantation of another person's idea into a target's subconscious.",
        "popularity": 112.09,
        "thumbnail": "https://image.tmdb.org/t/p/w92/ljsZTbVsrQSqZgWeep2B1QiDKuh.jpg",
        "genres": [
            "Action",
            "Science Fiction",
            "Adventure"
        ],
        "runtime": 148
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
404|Not Found|Nem található a film
500|Internal Server Error|An error occurred on the server

### 15. Find TV Show by Imdb ID

### Endpoint
```http
GET /find/tvshow
```

### Description
Sorozat keresése az IMDB API-ban.

### Parameters
Name|Type|Description
--- | --- | ---
id|string|IMDB azonosító

### Request Example
```http
GET /find/tvshow?id=tt0903747
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "tvshow": {
        "id": 1396,
        "title": "Breaking Bad",
        "first_air_date": "2008-01-20",
        "overview": "Walter White, a New Mexico chemistry teacher, is diagnosed with Stage III cancer and given a prognosis of only two years left to live. He becomes filled with a sense of fearlessness and an unrelenting desire to secure his family's financial future at any cost as he enters the dangerous world of drugs and crime.",
        "popularity": 841.54,
        "thumbnail": "https://image.tmdb.org/t/p/w92/ztkUQFLlC19CCMYHW9o1zWhJRNq.jpg",
        "genres": [
            "Drama",
            "Crime"
        ],
        "number_of_seasons": 5,
        "number_of_episodes": 62
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
404|Not Found|Nem található a sorozat
500|Internal Server Error|An error occurred on the server

### 16. Find Book by ISBN ID

### Endpoint
```http
GET /find/book
```

### Description
Könyv keresése a Google Books API-ban ISBN id alapján.

### Parameters
Name|Type|Description
--- | --- | ---
id|string|ISBN azonosító

### Request Example
```http
GET /find/book?id=9780451524935
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "book": {
        "id": "wGQ5mwEACAAJ",
        "isbn": "isbn:9630792257",
        "title": "1984",
        "author": "George Orwell",
        "description": "",
        "page_count": 358,
        "thumbnail": "",
        "categories": null,
        "published_date": "2011",
        "publisher": "",
        "language": "hu"
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
404|Not Found|Nem található a könyv
500|Internal Server Error|An error occurred on the server

### 17. Get Movie By TMDB ID

### Endpoint
```http
GET /movies/:id
```

### Description
Film részleteinek lekérése TMDB ID alapján.

### Parameters
Nincsenek

### Request Example
```http
GET /movies/27205
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "movie": {
        "id": 27205,
        "title": "Inception",
        "release_date": "2010-07-15",
        "overview": "Cobb, a skilled thief who commits corporate espionage by infiltrating the subconscious of his targets is offered a chance to regain his old life as payment for a task considered to be impossible: \"inception\", the implantation of another person's idea into a target's subconscious.",
        "popularity": 112.09,
        "thumbnail": "https://image.tmdb.org/t/p/w92/ljsZTbVsrQSqZgWeep2B1QiDKuh.jpg",
        "genres": [
            "Action",
            "Science Fiction",
            "Adventure"
        ],
        "runtime": 148
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
404|Not Found|Nem található a film
500|Internal Server Error|An error occurred on the server

### 18. Get TV Show By TMDB ID

### Endpoint
```http
GET /tvshows/:id
```

### Description
Sorozat részleteinek lekérése TMDB ID alapján.

### Parameters
Nincsenek

### Request Example
```http
GET /tvshows/1396
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "tvshow": {
        "id": 1396,
        "title": "Breaking Bad",
        "first_air_date": "2008-01-20",
        "overview": "Walter White, a New Mexico chemistry teacher, is diagnosed with Stage III cancer and given a prognosis of only 
two years left to live. He becomes filled with a sense of fearlessness and an unrelenting desire to secure his family's financial future at 
any cost as he enters the dangerous world of drugs and crime.",
        "popularity": 841.54,
        "thumbnail": "https://image.tmdb.org/t/p/w92/ztkUQFLlC19CCMYHW9o1zWhJRNq.jpg",
        "genres": [
            "Drama",
            "Crime"
        ],
        "number_of_seasons": 5,
        "number_of_episodes": 62
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
404|Not Found|Nem található a sorozat
500|Internal Server Error|An error occurred on the server

### 19. Get Book By Google Books ID

### Endpoint
```http
GET /books/:id
```

### Description
Könyv részleteinek lekérése Google Books ID alapján.

### Parameters
Nincsenek

### Request Example
```http
GET /books/wGQ5mwEACAAJ
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "book": {
        "id": "wGQ5mwEACAAJ",
        "isbn": "9630792257",
        "title": "1984",
        "author": "George Orwell",
        "description": "",
        "page_count": 358,
        "thumbnail": "",
        "categories": null,
        "published_date": "2011",
        "publisher": "Európa",
        "language": "hu"
    }
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
404|Not Found|Nem található a könyv
500|Internal Server Error|An error occurred on the server

### 20. Share Media Entry

### Endpoint
```http
POST /share
```

### Description
Megosztja a média bejegyzést.

### Parameters
Name|Type|Description
--- | --- | ---
media_entry|number|Média bejegyzés azonosító
share_with|string|Email cím

### Request Example
```http
POST /share
Content-Type: application/json

{
    "media_entry": 1,
    "share_with": "john.doe@example.com"
}
```

### Response

Status: `200 OK`

### Example Response:
No content

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
400|Bad Request|Hibás kérés
404|Not Found|Nem található a bejegyzés
500|Internal Server Error|An error occurred on the server

### 21. Get Shared Media Entries

### Endpoint
```http
GET /shared
```

### Description
Visszaadja a megosztott média bejegyzéseket.

### Parameters
Nincsenek

### Request Example
```http
GET /shared
```

### Response

Status: `200 OK`

### Example Response:
```json
{
    "sharedEntries": [
        {
            "id": 1,
            "entry_id": 1,
            "shared_by": 2,
            "shared_with": 1,
            "created_at": "2024-11-14T18:18:08Z"
        }
    ]
}
```

### Error Responses
Status Code|Error Message|Description
--- | --- | ---
500|Internal Server Error|An error occurred on the server