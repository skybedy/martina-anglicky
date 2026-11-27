# Důležitá rozhodnutí

Data níže vycházejí z dosavadní konverzace projektu. Implementovaný stav se
ověřuje v kódu; technologický směr není totéž jako nainstalovaná závislost.

## 2026-11-19 — Vývoj jako praktické osvěžení Go

- Postupovat po malých funkčních celcích, vysvětlovat důležité koncepty
  a po kroku dát prostor na otázky. Nevytvořit celý systém na jedno zadání.
- Důvod: uživatel je programátor a vrací se ke Go po několika letech.

## 2026-11-19 — HTTP vrstva bez frameworku

- Používat standardní `net/http` a `http.ServeMux`.
- Důvod: dostatečné pro tento projekt, přímo procvičuje Go a snižuje závislosti.
- Zvažováno: Gin, Echo a dřívější zkušenost s Gorilla mux. Echo bylo krátce
  zavedeno podle Tene Life, pak na výslovné přání nahrazeno standardní knihovnou.
- Stav: implementováno; v `go.mod` nejsou žádné externí závislosti.

## 2026-11-19 — Jedna učitelka, dva pohledy na stejné rezervace

- Administrace pro Martinu a veřejná nabídka volných termínů pro studenty.
  Rezervace nebo ruční lekce musí blokovat odpovídající čas pro všechny.
- Důvod: vlastní rezervační systém odpovídající způsobu práce Martiny.
- Studenti nesmějí vidět osobní údaje jiných studentů. MVP bez studentských
  účtů, plateb, více učitelů a složitých integrací.
- Stav: produktový směr, rezervační logika zatím není implementována.

## 2026-11-19 — Plánovaný zbytek stacku

- Go, serverové HTML šablony, MariaDB/MySQL přes `database/sql`, SQL migrace
  přes `golang-migrate`, Tailwind CSS 4 a vanilla JavaScript.
- Důvod: jednoduchost a návaznost na známé postupy Tene Life. Pro HTTP platí
  následné rozhodnutí o `net/http`, nikoli Echo používané v Tene Life.
- SQLite bylo zvažováno jako jednodušší začátek; směr se změnil na MariaDB
  při sjednocení s Tene Life. React nebyl pro dosavadní rozsah potřeba.
- Stav: Go a `html/template` existují; databáze, migrace a frontendový build ne.
  FullCalendar zůstává návrhem pro administraci; konkrétní verze nejsou zvolené.

## 2026-11-20 — Serverové šablony načítané při startu

- HTML patří do `templates/`; data se předávají strukturou. Šablona se načte
  při startu a při požadavku se vykreslí do bufferu před odesláním odpovědi.
- Důvod: oddělení HTML od obsluhy a možnost při chybě vykreslení vrátit 500
  namísto již částečně odeslané stránky se stavem 200.
- Stav: implementováno; relativní cesta vyžaduje spuštění z kořene projektu
  a změna šablony vyžaduje restart. `embed` ani automatické načítání změn není zavedeno.

## 2026-11-20 — Oddělení vstupu, rout a handlerů

- `cmd/server` sestaví závislosti a spustí server. `internal/web` obsahuje
  `app.go`, `routes.go` a `handlers.go`; handler je pojmenovaná metoda `*App`.
- Důvod: uživatel výslovně preferuje přehled rout na jednom místě a oddělené
  handlery; struktura aplikace jim poskytuje šablonu bez globálních proměnných.
- Budoucí `internal/booking` oddělí pravidla rezervací a `internal/store` SQL.
  Balíčky se vytvářejí až při reálné potřebě.
- Stav: implementováno pro web; nový příkaz je `go run ./cmd/server`.

## 2026-11-20 — Trvalá paměť v repozitáři

- `AGENTS.md` určuje práci v relacích, `PROJECT_CONTEXT.md` skutečný stav,
  `TODO.md` úkoly, tento soubor důvody a `START_NEW_CODEX_CHAT.md` obnovení kontextu.
- Důvod: pokračování napříč chaty bez spoléhání na historii konverzace.
- Paměť aktualizuje Codex po významné práci. Žádný commit, push ani deployment
  automaticky; existující necommitnuté změny se zachovávají.

## 2026-11-20 — Automatické testy jako součást vývoje

- Používat standardní `testing` a `net/http/httptest`, bez dalšího frameworku.
  Nové funkce doprovázet testy; při opravě pokud možno nejprve regresní test.
- Důvod: ověřovat chování průběžně a mít sadu spustitelnou přes `go test ./...`
  bez závislosti na běžícím serveru nebo produkčních datech.
- Stav: testy existující webové vrstvy a skutečné HTML šablony; proces startu
  není automaticky testovaný. Databáze a rezervace vzniknou až s vlastní implementací.
- Budoucí databázové testy musí být izolované. CI/CD se nyní nezavádí.

## 2026-11-20 — Jazyk a odpovědnosti při spolupráci

- Komunikace a projektové Markdown dokumenty jsou česky; zdrojové identifikátory,
  komentáře v kódu a commit messages anglicky. Jazyk UI se řeší samostatně.
- Konkrétní projektové volby řešíme primárně v Codexu nad repozitářem. Uživatel
  definuje požadavky a schvaluje významné volby, Codex zajišťuje technickou přípravu.
- Důvod: srozumitelnost dokumentace, návaznost na kód a omezení rozporů mezi chaty.

## 2026-11-20 — Nezávislé review u významných změn

- Uživatel schválil review významných změn, zejména architektury, rezervací,
  oprávnění, databáze, souběžnosti a podstatných uživatelských procesů.
- Cyklus: implementace → testy → nezávislé review → opravy → opakované testy
  → schválení commitu. Review chat kontroluje a hlásí nálezy, sám neupravuje soubory.
- Drobné změny mohou projít jen relevantními testy a kontrolou diffu.
- Důvod: zaměřit nezávislou kontrolu na podstatná rizika bez zbytečné administrativy.
  Povinné nezávislé review každé drobné úpravy nebylo zvoleno.
- Commit vyžaduje schválení; jeho schválení nezahrnuje push ani deployment.
- Review probíhá nad dočasně zmrazeným sdíleným checkoutem. Reviewer nemění
  ani projektovou paměť, index či historii; opravy provádí implementační chat.
- Důvod upřesnění: zabránit souběžným změnám kontrolovaného stavu a konfliktu
  mezi čtecím review režimem a obecnou povinností aktualizovat dokumentaci.

## 2026-11-20 — UX/UI před implementací rozhraní

- Společný návrh procesů a obrazovek obou částí aplikace, schválení koncepce,
  případný funkční vizuální prototyp, schválení řešení a teprve potom implementace.
- Důvod: ověřit práci Martiny i rezervaci studenta včas, včetně mobilního ovládání.
- `DESIGN.md` vznikne až s prvním odsouhlaseným návrhem, nebude prázdnou šablonou.
- Upřesňuje dosavadní frontendový směr: Tailwind CSS, vanilla JavaScript
  a FullCalendar nejsou zavedené; jejich finální použití se potvrdí při návrhu.
  Toto rozhodnutí nemění Go, `net/http` ani současnou funkční HTML kostru.
- Stav: pravidla schválena a zapsána; vznikl neschválený prezentační prototyp,
  `DESIGN.md` zatím neexistuje.

## 2026-11-24 — Prezentační web je aktuální priorita

- Práci na rezervačním systému pozastavit a výchozí stav uchovat ve větvi
  `archive/booking-foundation`.
- Na větvi `feat/martina-website` nejprve připravit základní vizuální a obsahový
  nástřel webu Martiny pro české studenty, klidně s jasně označenými ukázkovými údaji.
- Důvod: Martina a uživatel potřebují konkrétní podklad pro diskusi o směru webu;
  rezervační funkce lze doplnit později.

## 2026-11-27 — Prototyp spouštět jako Go aplikaci

- Vizuální návrh obsluhuje existující `net/http` aplikace na kořenové routě,
  nikoli samostatný statický server.
- Šablona i pracovní portrét se načtou při startu a předají do `web.App`;
  portrét vrací samostatný pojmenovaný handler.
- Důvod: zachovat jeden běžný způsob lokálního spuštění `go run ./cmd/server`
  a současnou architekturu oddělených rout, handlerů a závislostí.
- Stav: implementováno jako pracovní prototyp; obsah ani vzhled nejsou schválené.

## Dosud neuzavřeno

Konkrétní datové schéma, přihlašovací knihovna, rezervační předstihy, rušení,
potvrzování, přesný model časových pásem, UX/UI návrh a finální frontendové nástroje.
Dříve navržené hodnoty nejsou
schválenou implementační specifikací; navazující úkoly jsou v [TODO.md](TODO.md).
