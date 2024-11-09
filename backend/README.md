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
    DB_PATH=<DB_PATH> # Az SQLite adatbázis elérési útja
    GOOGLE_BOOKS_API_KEY=<GOOGLE_BOOKS_API_KEY> # A Google Books API kulcsa
    TMDB_API_KEY=<TMDB_API_KEY> # A The Movie Database API kulcsa
    API_PORT=<API_PORT> # Az API portja
    CORS_TRUSTED_ORIGINS=<CORS_ORIGINS> # A CORS engedélyezett eredetek szóközzel elválasztva
    JWT_SECRET=<JWT_SECRET> # A JWT titkosításhoz használt kulcs
    ```
5. Indítsd el a szervert:
    ```bash
    go run main.go
    ```
6. A szerver elérhető a `http://localhost:<API_PORT>` címen.
