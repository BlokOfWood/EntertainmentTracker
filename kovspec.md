# Entertainment Tracker
## Áttekíntés
Mindannyian sok filmet, könyvet vagy sorozatot elkezdünk aztán félbehagyunk, elfelejtve hogy meddig jutottunk. Ez az applikáció ezt a problémát segítené azzal, hogy lehetővé teszi a megkezdett szórakoztató tartalmak feljegyzését és a haladás megjelölését.

## Jelenlegi helyzet
Jelenleg a szórakoztató médiák számontartása különböző appokon és webhelyeken keresztül tehető meg, amik nem kifejezetten abból a célból jöttek létre, hogy a személyes média fogyasztásunkat tartsuk számon. Ezért olyan módszerekhez kell folyamodni, amik nem kézenfekvőek például egy Excel tábla.

## Vágyálom rendszer
A projekt célja egy olyan rendszer kialakítása, ami a különböző médiák fogyasztását egy felületen kezeli. A művek egy táblázatban kerülnek megjelenítésre, mely különböző tulajdonságok alapján szűréssel kereshető. Ezen a felületen tárolásra kerülhet az adott műbeli haladás, ilyen állapotokkal például, hogy: még nem kezdte el a felhasználó, de szeretné megnézni vagy már nem szeretné folytatni. A program saját eszközön tárolt adatokkal dolgozik, így a személyes adatok nincsenek veszélynek kitéve.

## Funkcionális követelmények
- Művek megjelenítése
- Előrehaladás megjelenítése
- Új művek hozzáadása
- Művek módosítása
- Művek szűrése
- Művek törlése

## Rendszerre vonatkozó törvények, szabványok, ajánlások
A rendszernek jogszabályi kérdésekkel nem kell foglalkoznia, mert a minimális személyes adatok, amik a program működéséhez szükségesek is lokálisan kerülnek tárolásra, nem távoli szerveren.
A rendszerre fejenként 100-200 kódsoros és 4 hetes határidő megszorítás vonatkozik.

## Követelménylista
|Modul|ID|Név|v. |Kifejtés|
|:----|:-|:--|:--|:-------|
|Felület|K1|Mű táblázat|1.0|A felület, amelyen az összes követett mű kijelzésre kerül.|
|Felület|K2|Mű hozzáadás/módosítás ablak|1.0|A felület, amelyen új követett művet lehet felvenni vagy módosítani egy meglévőt.|
|Felület|K3|Művek szűrése|1.0|Lehessen szűrni a műveket a név és állapot alapján.|
|Felület|K4|Mű előrehaladás kezelése|1.0|A különböző típusú művekkel az előrehaladást típusnak megfelelően lehet követni. Például: Könyvnél lap szám/max lapszám, filmnél idő/max idő|
|Felület|K5|Mű törlése megerősítése|1.0|A törlés megerősítésére szolgáló felugró ablak.|
|Modifikáció|K6|Mű felvétele|1.0|A felhasználó fel tudjon venni egy új művet.|
|Modifikáció|K7|Mű módosítása|1.0|A felhasználó egy meglévő művön tudja módosítani a nevet, állapotát és előrehaladást a művel.|
|Modifikáció|K8|Mű törlése|1.0|A felhasználó egy meglévő művet töröl a rendszerből.|
|Perziesztencia|K9|Művek perzisztenciája|1.0|Művek tárolása a rendszer indítások között.|
