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
