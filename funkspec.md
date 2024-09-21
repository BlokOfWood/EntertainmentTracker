# Entertainment Tracker
## A rendszer céljai és nem céljai
Cél, hogy a projekt segítsen a szórakoztató médiák fogyasztásának a felhasználó általi követésében, hogy ne kelljen több különböző oldalt használni, amelyek nem is erre a célra lettek elkészítve. 
Nem cél, hogy az előrehaladás követése automatikusan történjen, ezt a felhasználónak saját magának kell, hogy kövesse. Nem cél az sem, hogy ezekhez a művekhez automatikusan szolgáltassunk metaadatot.

## [Jelenlegi helyzet](https://github.com/BlokOfWood/EntertainmentTracker/blob/main/kovspec.md#jelenlegi-helyzet)

## [Vágyálom rendszer](https://github.com/BlokOfWood/EntertainmentTracker/blob/main/kovspec.md#v%C3%A1gy%C3%A1lom-rendszer)

## [Rendszerre vonatkozó törvények, szabványok, ajánlások](https://github.com/BlokOfWood/EntertainmentTracker/blob/main/kovspec.md#rendszerre-vonatkoz%C3%B3-t%C3%B6rv%C3%A9nyek-szabv%C3%A1nyok-aj%C3%A1nl%C3%A1sok)

## [Követelménylista](https://github.com/BlokOfWood/EntertainmentTracker/blob/main/kovspec.md#k%C3%B6vetelm%C3%A9nylista)

## Használati esetek
### Első indítás
- leírás: a felhasználó első alkalommal elindítja a rendszert, a rendszer inicializálja a működés feltételeit
- előfeltételek: az aktuális indítás a rendszer első indítása
- utófeltételek: a rendszer készen áll a további működésre
- egyéb: létrehozza a helyi adatbázist
### Új mű hozzáadása
- leírás: a felhasználó a rendszerben még nem szereplő művet vesz fel, a rendszer hozzáadja az új művet
- előfeltételek: a mű még nem szerepel a rendszerben, az első indítás megtörtént
- utófeltételek: a rendszerben tárolásra került az új mű
- egyéb: a mű hozzáadás után megtekinthető
### Meglévő mű szerkesztése
- leírás: a felhasználó a rendszerben már szereplő műről tárolt adatokat frissíti, a rendszer frissíti a mű adatai
- előfeltételek: a mű már létezik a rendszerben
- utófeltételek: a rendszer a művet a frissített adatokkal tárolja
- egyéb: a mű a frissített adatokkal tekinthető meg
### Meglévő mű törlése
- leírás: a felhasználó a rendszerben már szereplő művet töröl, a rendszer eltávolítja a művet és a róla tárolt adatokat
- előfeltételek: a mű már létezik a rendszerben
- utófeltételek: a rendszer nem tárolja a művet
- egyéb: a mű többé nem tekintehtő meg
### Művek megtekintése
- leírás: a felhasználó megtekinti a rendszerben tárolt műveket
- előfeltételek: az első indítás megtörtént
- egyéb: az összes tárolt mű megtekinthető
### Művek szűrése
- leírás: a felhasználó a tárolt művek között szűr, a rendszer csak a szűrésnek megfelelő műveket jeleníti meg
- előfeltételek: az első indítás megtörtént

## Megfeleltetés
|Használati Esetek|Követelmény ID|
|:----------------|:--------------|
|Első indítás|K9|
|Új mű hozzáadása|K2, K6|
|Meglévő mű szerkesztése|K2, K4, K7|
|Meglévő mű törlése|K5, K8|
|Művek megtekíntése|K1|
|Művek szűrése|K3|

## Képernyő tervek
### Művek megtekintése képernyő
![Művek megtekintése képernyő](./assets/main_screen.jpg)

### Mű hozzáadás vagy módisítás képernyő
![Mű hozzáadás vagy módisítás képernyő](./assets/add_or_modify_screen.jpg)

### Törlés megerősítés képernyő
![Törlés megerősítés képernyő](./assets/alert.jpg)

## Forgatókönyvek
- Haladását módosítja egy adott műben a felhasználó: mű melletti szerkesztés gomb -> módosítás ablak -> átírja a haladást -> mentés
- Új mű hozzáadása: új mű gomb -> új mű ablak -> kitölti az űrlapot -> mentés gomb
- Mű szerkesztése: adott művön szerkesztés gomb -> mű szerkesztés ablak -> átírja amit szeretne -> mentés gomb
- Mű törlése: mű törlése gomb -> törlés megerősítés felugró ablak -> megerősítés gomb
- Művek szűrése: adott kategória szűrő mező -> szűrési minta beírása
