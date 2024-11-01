# Entertainment Tracker

## [Rendszer céljai és nemcéljai](https://github.com/BlokOfWood/EntertainmentTracker/blob/main/funkspec.md#a-rendszer-céljai-és-nem-céljai)

## Üzleti folyamatok modellje

Kiváltandó üzleti folyamatok modellje:

![Kiváltandó üzleti folyamatok modellje](./assets/uzleti_folyamatok_abrak/kivaltando_uzleti_folyamatok_abra.jpg)

MediaMind üzleti folyamatainak modellje:

![MediaMind üzleti folyamatainak modellje](./assets/uzleti_folyamatok_abrak/mediamind_uzleti_folyamata_abra.jpg)

## Funkcionális terv

Rendszerszereplők:

- felhasználó

Rendszerhasználati esetek és lefutásaik:

- új felhasználó regisztrál a rendszerbe

![Regisztráció](./assets/usecase_umls/register.jpg)

- már regisztrált felhasználó bejelentkezik a rendszerbe

![Bejelentkezés](./assets/usecase_umls/login.jpg)

- bejelentkezett felhasználó kijelentkezik a rendszerből

![Kijelentkezés](./assets/usecase_umls/logout.jpg)

- bejelentkezett felhasználó új művet ad hozzá a tárolt műveihez

![Új mű hozzáadása](./assets/usecase_umls/new_media.jpg)

- bejelentekezett felhasználó már meglévő művet szerkeszt

![Már hozzáadott mű szerkesztése](./assets/usecase_umls/edit_media.jpg)

- bejelentkezett felhasználó meglévő művet töröl

![Már hozzadott mű törlése](./assets/usecase_umls/delete_media.jpg)

- bejelentkezett felhasználó megtekinti a mentett műveket

![Mentett művek megtekintése](./assets/usecase_umls/view_saved_media.jpg)

- bejelentkezett felhasználó szűr a mentett művek között

![Szűrés mentett művek között](./assets/usecase_umls/filter_saved_media.jpg)

- bejelentkezett felhasználó műben való haladást oszt meg

![Műben való haladás megosztása](./assets/usecase_umls/share_media.jpg)

Menü-hierarchiák:

- Navbar:
  - művek megtekintése
  - új mű hozzáadása
  - kijelentkezés
- Dashboard:
  - mentett művek megtekintése
  - szűrés mentett művek között
  - műben való haladás megosztása
  - mű szerkesztése
  - mű törlése
- Add Media:
  - könyv hozzáadása
  - sorozat hozzáadása
  - film hozzáadása
  - YouTube video hozzáadása
- Books:
  - könyv keresése cím, ISBN azonosító alapján
  - találatok közüli könyv hozzáadása
- TV Shows:
  - sorozat keresése cím, IMDb ID alapján
  - találatok közüli könyv hozzáadása
- Movies:
  - film keresése cím, IMDb ID alapján
  - találatok közüli könyv hozzáadása
- YouTube:
  - megjelenítési név megadása
  - YouTube URL megadása

## Adatbázis terv
```mermaid
erDiagram
    USERS {
        INTEGER user_id PK "Unique identifier for users"
        TEXT username "Username"
        TEXT password "Password (hashed)"
        TEXT email "Email for login"
        DATETIME created_at "User creation date"
    }
    
    MEDIA_ENTRIES {
        INTEGER entry_id PK "Unique identifier for media entries"
        INTEGER user_id FK "User who added the media"
        TEXT title "Title of the media"
        TEXT type "Type of the media (book, tv show, movie, youtube)"
        TEXT status "Status of the media (not started, in progress, completed)"
        DATETIME created_at "Media entry creation date"
        DATETIME updated_at "Last update date"
    }
    
    PROGRESS {
        INTEGER progress_id PK "Unique identifier for progress"
        INTEGER entry_id FK "Media entry being tracked"
        INTEGER current "Current progress"
        INTEGER target "Target progress"
        DATETIME updated_at "Last update date"
    }
    
    SHARED_ENTRIES {
        INTEGER share_id PK "Unique identifier for shared entries"
        INTEGER entry_id FK "Media entry being shared"
        INTEGER shared_by FK "User who shared the media"
        INTEGER shared_with FK "User who the media is shared with"
        DATETIME created_at "Share creation date"
    }

    USERS ||--o{ MEDIA_ENTRIES : "1:n owns"
    USERS ||--o{ SHARED_ENTRIES : "1:n shares"
    MEDIA_ENTRIES ||--|| PROGRESS : "1:1 tracks"
    MEDIA_ENTRIES }o--o{ SHARED_ENTRIES : "n:m is shared in"
```


## Frontend tesztterv

A tesztelés célja a frontend megfelelő működésének vizsgálata.

A frontendnek a felület megnyitásakor a felhasználónak elérhetőve kell tenni a regisztráció vagy bejelentkezés opciót.
Bejelentkezés után a Dashboard-on a mentett műveknek megtekinhetőnek kell lenniük, köztük szűrni lehessen, új mű hozzáadásának vagy már létező mű szerkesztésének, törlésének lehetségesnek kell lenni.
A felhasználónak lehetősége kell legyen műben való haladást megosztani, megosztott műben való haladást megtekinteni.

Unit teszt:
A tesztelés fejlesztési idő alatt történik.
A megjelenített adatok pontosságának ellenőrzése, új mű hozzádása vagy már hozzáadott műveken végzett műveletek a megfelelő következményeket vonja maga után.
Megbizonyosodni a frontend komponenseinek jelenlétéről a megfelelő funkciókkal, megjelenéssel.
Komponensek és oldalak közötti útvonalak helyes működésének ellenőrzése.
Frontend metódusok megfelelő működésének ellenőrzése.

Alfa teszt:
Fejlesztők által végzett teszt. Elsődleges célja az rendszer működésének felhasználó szemszögéből való ellenőrzése.
A rendszer a fejlesztési folyamat alatt, valamint kész állapotában is tesztelésre kerül, esetlegesen felmerülő problémák észlelése, javítása miatt.

## Backend tesztterv

A tesztelés célja a backend megfelelő működésének és az adatkezelés pontosságának vizsgálata.

A backendnek biztosítania kell az adatbázis műveletek megfelelő végrehajtását, például művek hozzáadását, szerkesztését, törlését, a felhasználói hitelesítést, valamint a megosztott művek kezelését. A cél az API végpontok helyes működésének ellenőrzése, hogy a felhasználó be tudjon jelentkezni, regisztrálni tudjon, és a felhasználói műveletek a megfelelő adatbázis-frissítéseket eredményezzék. A backendnek a műveletek során fellépő hibák esetén is kezelnie kell a helyzetet, megfelelő hibaüzeneteket visszaadva.

### Unit teszt:
- **Cél:** Fejlesztési idő alatt történő részletes ellenőrzés az egyes API végpontok és adatkezelési logika helyességéről.
- **Tesztelendő területek:**
  - **Felhasználói hitelesítés:** Regisztráció és bejelentkezés végpontjainak tesztelése, helyes és helytelen adatokra adott válaszok ellenőrzése.
  - **CRUD műveletek művekhez:** Mű hozzáadásának, módosításának, törlésének tesztelése, ellenőrizve az adatbázis frissülését és a helyes válaszok visszaadását.
  - **Szűrés és lekérdezés:** Szűrési funkciók tesztelése különböző paraméterekkel (például név, státusz) a helyes találati eredmények biztosítására.
  - **Megosztási funkció:** Megosztott művek kezelésének tesztelése, művek megosztása és megtekintése, jogosultságok ellenőrzése.
  - **Adat integritás:** A backend biztosítja, hogy minden adatbázis-módosítás konzisztense legyen, például a törölt művek nem maradhatnak megosztva.
- **Módszer:** Minden végponthoz külön tesztesetek létrehozása, pozitív és negatív esetek lefuttatása, illetve adatbázis állapotok ellenőrzése.

### Alfa teszt:
- **Cél:** A backend működésének felhasználói szemszögből történő ellenőrzése, a rendszer teljes funkcionalitásának tesztelése.
- **Tesztelendő területek:**
  - **Felhasználói élmény:** A felhasználók számára elérhető funkciók helyes működésének ellenőrzése, például regisztráció, bejelentkezés, művek kezelése.
  - **Adatkezelés:** A felhasználói műveletek helyes hatásainak ellenőrzése az adatbázisban, például művek hozzáadása, szerkesztése, törlése.
  - **Hibakezelés:** A rendszer helyes hibaüzeneteket ad-e vissza, ha a felhasználó hibás adatokat ad meg vagy hibás műveleteket próbál végrehajtani.
- **Módszer:** Manuális tesztelés a rendszer teljes funkcionalitásának ellenőrzésére, különböző felhasználói szituációk szimulálása, hibás adatokkal való próbálkozás.
