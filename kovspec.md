# Entertainment Tracker

## Szabadriport
Egy olyan rendszer, amiben követni lehet a könyvekkel, Netflix és egyéb sorozatokkal, filmekkel, YouTube videókkal való haladást. Külön megjelenítés jó lenne médiatípusonként, a rendszer az adott művekhez automatikusan begyűjti az adatokat külső forrásból. A tárolt művek közül lehessen törölni, lehessen őket rendezni a dashboard-on.
Különböző eszközökön való bejelentkezés után ugyanazt lássa a felhasználó.
Az adatok táblázatosan kerüljenek megjelenítésre.

## Irányítottriport

- **Külső forrsából begyűjtött adatokként pontosan mire gondol a megrendelő?**
- Kategóriák, összefoglaló, könyveknél író.
<br>

- **Homescreen-t szeretne-e a felhasználó?**
- Nem, a dashboard-ra vigyen bejelentkezés után, onnan lehessen tovább navigálni.
<br>

- **Külön megjelenítés médiatípusonként pontosan mit jelent?**
- A haladáskövetés médiatípusonként specifikusan legyen, könyveknél oldalszám, filmeknél perc, sorozatoknál epizódszám, YouTube videonál óra : perc : másodperc formátum.
<br>

- **Legyen-e beágyazva a YouTube videó?**
- Igen.
<br>

- **Szerkeszthet-e több felhasználó egy műhöz tárolt haladást?**
- Nem, csak én tudjam szerkeszteni, amit én vettem fel, más csak láthassa, ha megosztom.
- **Legyen külön nézet a "megosztott bejegyzések", "saját bejegyzések"-nek?**
- Ne, csak lehessen rendezni.
<br>

- **Pontosan mi alapján lehessen rendezni?**
- Név, típus, haladás és saját vagy megosztott bejegyzés.

## Áttekíntés
Mindannyian sok filmet, könyvet vagy sorozatot elkezdünk aztán félbehagyunk, elfelejtve hogy meddig jutottunk. Ez az applikáció ezt a problémát segítené azzal, hogy lehetővé teszi a megkezdett szórakoztató tartalmak feljegyzését és a haladás megjelölését.

## Jelenlegi helyzet
Jelenleg a szórakoztató médiák számontartása különböző appokon és webhelyeken keresztül tehető meg, amik nem kifejezetten abból a célból jöttek létre, hogy a személyes média fogyasztásunkat tartsuk számon. Ezért olyan módszerekhez kell folyamodni, amik nem kézenfekvőek például egy Excel tábla.

## Vágyálom rendszer
A projekt célja egy olyan rendszer kialakítása, ami a különböző médiák fogyasztását egy felületen kezeli. A művek egy táblázatban kerülnek megjelenítésre, melyet különböző tulajdonságok alapján rendezni lehet. Ezen a felületen tárolásra kerülhet az adott műbeli haladás. A rendszer a felhasználó eszközei között szinkronizálva tárolja a hozzáadott bejegyzéseket.

## Funkcionális követelmények
- Művek megjelenítése
- Előrehaladás megjelenítése
- Új művek hozzáadása
- Művek módosítása
- Művek szűrése
- Művek törlése
- Metaadat betöltése a művekhez
- Felhasználók kezelése
- Bejegyzések megosztása más felhasználókkal

## Rendszerre vonatkozó törvények, szabványok, ajánlások
### Adatvédelem és adatkezelés
A rendszer, amely a felhasználók email-címét, felhasználónevét, jelszavának hashelt változatát és előrehaladási adatait tárolja, köteles betartani a [GDPR (General Data Protection Regulation)](https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX:02016R0679-20160504) szabályait. Ez magában foglalja a felhasználók adatainak védelmét, az adatok biztonságos tárolását és a felhasználók jogainak tiszteletben tartását, mint például az adatkezelési tájékoztatás, hozzáférés, adatmódosítás és törlés lehetősége.

### Szerzői jogok és API-k használata
Az alkalmazásban tárolt vagy felhasznált harmadik féltől származó adatokra vonatkozóan be kell tartani a szerzői jogi és licencelési szabályokat:
- TheMovieDB: https://www.themoviedb.org/api-terms-of-use
- Google API: https://developers.google.com/books/terms

## Követelménylista
|Modul|ID|Név|v. |Kifejtés|
|:----|:-|:--|:--|:-------|
|Felület|K1|Mű táblázat|1.0|A felület, amelyen az összes követett mű kijelzésre kerül.|
|Felület|K2|Mű hozzáadás/módosítás ablak|1.0|A felület, amelyen új követett művet lehet felvenni vagy módosítani egy meglévőt.|
|Felület|K3|Művek rendezése|1.0|Lehessen rendezni a műveket a név, típus, haladás és megosztás alapján.|
|Felület|K4|Mű előrehaladás kezelése|1.0|A különböző típusú művekkel az előrehaladást típusnak megfelelően lehet követni. Például: Könyvnél lap szám/max lapszám, filmnél idő/max idő|
|Felület|K5|Mű törlése megerősítése|1.0|A törlés megerősítésére szolgáló felugró ablak.|
|Modifikáció|K6|Mű felvétele|1.0|A felhasználó fel tudjon venni egy új művet.|
|Modifikáció|K7|Mű módosítása|1.0|A felhasználó egy meglévő művön tudja módosítani a nevet, állapotot és előrehaladást a művel.|
|Modifikáció|K8|Mű törlése|1.0|A felhasználó egy meglévő művet töröl a rendszerből.|
|Perziesztencia|K9|Művek perzisztenciája|1.0|Művek tárolása a rendszer indítások között.|
|Felhasználó kezelés|K10|Felhasználó regisztráció|1.0|Egy regisztrálatlan felhasználó fel tudja magát venni a rendszerbe.|
|Felhasználó kezelés|K11|Felhasználó bejelentkezés|1.0|Egy regisztrált felhasználó be tud jelentkezni a rendszerbe.|
|Felület|K12|Metaadat betöltése|1.0|A felhasználó láthat külső forrásból információkat anélkül, hogy meg kell adnia.|
|Felhasználó kezelés|K13|Bejegyzés megosztása másik felhasználóval|1.0|Egy felhasználó meg tud osztani egy saját bejegyzést egy másik felhasználóval.|
|Felhasználó kezelés|K14|A felhasználó kijelentkeztetése a rendszerből|1.0|Egy bejelentkezett felhasználó kijelentkezik a rendszerből.|
