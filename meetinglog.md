# Meeting log

## 2024.10.07. - 17 perc

- nagy projekt témájának véglegesítése
- ötletelés a rendszer lehetséges funkcióiról

## 2024.10.10. - 1 óra 8 perc

- rendszer feladatainak meghatározása, működésének és lehetőségeinek feltárása
- ezen feladatok és működésük pontosítása
- követelmény lista elemeinek felvétele
- első feladatok megbeszélése
- projektmenedzsment eszközként a Jira-ra esett a választás, scrum és agilis fejlesztéshez nyújtott támogatása miatt

## 2024.10.19. - 1 óra 22 perc

- vágyálom rendszer, rendszer céljai, nem céljai frissítése
- követelménylista kiegészítése
- használati esetek bővítése, pontosítása
- használati esetek-követelmény lista megfeleltetés frissítése
<br>

- API választás, és ez alapján a rendszerre vontakozó követelmények frissítése feladat kiadása
- use case UML ábrák átbeszélése, bővítési feladat kiadása
<br>

rendszertervbeli feladatok meghatározása:
- rendszer célja - belinkelni funkspecből
- projekt terv
- üzleti folyamatok modellje ábra formában - Excalidraw
- funkcionális terv UML ábrákként - PlantUML
- fizikai környezet
- architekturális terv - backend, frontend
- adatbázis terv
- tesztterv - backend, frontend
- telepítési terv - backend, frontend

## 2024.11.03. - 1 óra 40 perc
projektterv megírása:
 - szerepkörök meghatározása
 - sprintekben elvégzendő feladatok meghatározása
 - mérföldkövek meghatározása

rendszerterv átnézése:
- menü-hierarchiák javítása
- fizikai környezet kiegészítése
- adatbázis terv frissítése

## 2024.11.09. - 1 óra 20 perc

### Code review:

#### backend észrevételek/teendők:
- project felépítés refaktorálása
- .env fájl betöltése indításkor
- hibakezelés igazítása
- register objektum pontosítása - meeting során megoldva
- README.md frissítése a változtatások alapján

#### frontend észrevételek/teendők:

felhasználókezelés:
- átirányítás igazítása
- input binding igazítása - meeting során megoldva
- access token kezelés implementálása

dashboard:
- header igazítás kód egyszerűsítés céljából

## 2024.11.17. - 38 perc

### Code review:

#### frontend észrevételek/teendők:
- dashboard scroll overflow - meeting alatt megoldva
- scroll-olásnál csak a művek kerüljenek görgetésre, a fejléc maradjon (gyermekelembe rakni a műveket)
- mű szerkesztése után visszanavigálni a dashboard-ra és újra fetch-elni
- delete media pop-up - képernyőterv alapján alakítani

- YouTube video hozzáadásánál URL megadása csak
- YouTube videónál display name helyett, progress-t lehessen módosítani

- megosztott művek dashboard-on jelölve jelenjenek meg (képernyőterv + implementálás)

## 2024.11.24. - 1 óra 5 perc

### Code review:

#### frontend észrevételek/teendők:
- jelenlegi progress megjelenítése progress szerkesztésénél
- YouTube videó beágyazása javítása - meeting alatt megoldva
- YouTube videó hozzáadása - meeting alatt megoldva
- törlés és szerkesztés utáni táblázatbetöltés javítása - meeting alatt megoldva
- rendezés után a nyilak egyértelműek legyenek
  - lefele mutat -> csökkenő (v. nem shared)
  - felfele mutat -> növekvő (v. shared)
  - mind a két nyíl -> rendezetlen 
- felhasználóval megosztott és felhasználó által felvett művek beolvasása után dátum szerinti sorrendbe rendezni (legutóbb hozzáadott legelőre)
- százalék megjelenítése progress-nél

## 2024.12.02. - 22 perc

### további teendők:
- tesztelés elvégzése
  - end to end
  - unit test
  - integration
- shared media törlésének implementálása frontend-en (az által, akivel meg lett osztva)
- felhasználónév kiírása, profilképgenerálás
- dokumentációk átnézése
