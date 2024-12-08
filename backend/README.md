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
    API_PORT=6969
    API_ENV="development"
    API_DB_PATH="db.db"
    CORS_TRUSTED_ORIGINS="http://localhost:3000 http://localhost:3001"
    AUTH_EXPIRE_TIME=14
    TMDB_API_KEY="<api-key>"
    GOOGLE_API_KEY="<api-key>"
    ```
5. Indítsd el a szervert:
    ```bash
    go run main.go
    ```
6. A szerver elérhető a `http://localhost:<API_PORT>` címen.

# API dokumentáció
Az API dokumentációja a [Swagger Editor](https://editor-next.swagger.io/) segítségével készült. Az API dokumentációja a `./swagger/index.html` fájlban érhető el.