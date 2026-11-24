# Kontext projektu

Poslední věcné ověření: 2026-11-20. Dokument popisuje stav souborů, nikoli
záruku, že právě běží server. Git a prostředí ověř při každé nové relaci.

## Účel a způsob práce

Vlastní rezervační web pro Martinu, jednu učitelku online angličtiny.
Martina spravuje dostupnost a lekce; studenti si ve veřejné části vybírají
uvolněné termíny. Oba pohledy budou používat stejná data. Veřejná část nesmí
odhalovat údaje ostatních studentů a rezervace se nesmějí překrývat.

Vývoj současně slouží jako návrat uživatele ke Go. Postupujeme po malých
funkčních krocích s vysvětlením a zastavením na dotazy. Aktuální úkol ani
obnovení chatu není pokyn k vytvoření celého rezervačního systému.

Konkrétní koncepční a architektonické volby řešíme primárně zde nad skutečným
repozitářem. Uživatel rozhoduje o požadavcích a významných volbách; Codex
připravuje návrhy, implementaci, testy a dokumentaci. Obecné vzdělávání může
probíhat jinde, ale projektová paměť uchovává aktuální projektová rozhodnutí.
Komunikace a Markdown dokumentace jsou česky, identifikátory, komentáře v kódu
a commit messages anglicky. Jazyk UI se řeší samostatně; současná stránka je česky.

Významné změny vyžadují nezávislé review, ideálně v samostatném chatu, a po
opravách opakování relevantních testů. Drobné změny ověřujeme testy a diffem.
Commit vyžaduje výslovné schválení; push a deployment samostatný souhlas.

## Ověřený technologický stav

- Modul `github.com/skybedy/martina-anglicky`, `go.mod` deklaruje Go 1.24.0.
- Pouze standardní knihovna: HTTP server a router `net/http`, `html/template`,
  `bytes`, `log`. Žádný Gin, Echo ani Gorilla, žádné externí Go závislosti.
- Jedna česká HTML stránka bez CSS či JavaScriptu. Bez databáze a konfigurace `.env`.
- Testy v `internal/web/web_test.go` používají standardní `testing` a `httptest`.
  CI/CD ani kontejnerová konfigurace nejsou zavedené.

Plánovaný, zatím nezavedený stack: MariaDB/MySQL přes `database/sql` a ovladač,
SQL migrace `golang-migrate`. Tailwind CSS 4 a vanilla JavaScript jsou dosavadní
preferovaný frontendový směr, FullCalendar je návrh pro administraci. Jejich
finální použití se potvrdí při návrhu UX/UI; žádná z těchto frontendových
závislostí není zavedena. Důvody jsou v [DECISIONS.md](DECISIONS.md).

## Stav UX/UI návrhu

Aktuální změna směru (2026-11-24): práce na rezervačním systému je pozastavena.
Na větvi `feat/martina-website` vzniká návrh hlavního prezentačního webu pro
Martinu jako online lektorku angličtiny. Soubor `prototype/index.html` je první
statický vizuální nástřel s ukázkovými texty a pracovně upraveným portrétem
odvozeným z dodané fotografie. Před zveřejněním je nutné použít originální soubor;
návrh ještě není schválen a není součástí běžící Go aplikace. Původní základ
zůstává ve větvi `archive/booking-foundation`.

V běžící Go aplikaci existuje pouze základní HTML stránka. Vizuální styl
prezentačního webu zatím není odsouhlasený. `DESIGN.md` neexistuje.

Nejprve společně navrhneme procesy a obrazovky, potom odsouhlasíme koncepci,
případně vytvoříme prototyp a po schválení řešení přistoupíme k implementaci.
Mobilní ovládání, přehlednost a konzistence jsou požadavky, ne již ověřené vlastnosti.

## Architektura a tok požadavku

- `cmd/server/main.go`: načte šablonu, vytvoří `web.App` a spustí server.
- `internal/web/app.go`: `App` drží `*template.Template`; `New` ji přijímá jako závislost.
- `internal/web/routes.go`: `Routes() http.Handler` vytvoří `http.ServeMux`
  a registruje `GET /{$}` na pojmenovanou metodu `app.home`.
- `internal/web/handlers.go`: sestaví `homePageData`, vykreslí šablonu do bufferu
  a odešle HTML. Při chybě vykreslení vrátí 500, chyby zápisu loguje.
- `templates/home.html`: titulek a popis z dat, text o připravovaných rezervacích.

Požadavek jde přes HTTP server do routeru, metody `app.home` a HTML šablony.
Šablona se parsuje jednou při startu; po její úpravě je nutný restart.
Rezervační pravidla mají později patřit do `internal/booking` a SQL do
`internal/store`; tyto balíčky dosud neexistují.

## Spuštění a ověření

Spouštěj z kořene repozitáře, protože cesta k šabloně je relativní:

```sh
go run ./cmd/server
go build ./...
go test ./...
go test -race ./...
go vet ./...
```

Server má pevnou adresu `127.0.0.1:8080`. Před spuštěním ověř případné obsazení
portu, neukončuj cizí proces. `go run .` již po přesunu do `cmd/server` nefunguje.
V omezeném prostředí lze použít `GOCACHE=/tmp/martina-anglicky-go-build`.
Oprávnění k síti, cache či socketům řeš standardním schvalovacím mechanismem.

## Stav automatických testů

Před zavedením sady nebyly žádné testy ani testovací framework. Obchodní logika
dosud neexistuje, takže testujeme dokončenou HTTP vrstvu a šablonu:

- `TestRoutes`: osm scénářů pro skutečný router, handler a šablonu; úvodní HTML
  a data, query parametr, `HEAD`, neznámé cesty a nepovolené metody včetně `Allow`.
- `TestHomeTemplateError`: selhání vykreslení po částečném zápisu musí vrátit
  500 a veřejnou zprávu bez rozpracovaného HTML nebo interních detailů.
- `TestHomeTemplateData`: skutečná šablona escapuje HTML v textových polích
  a zvládá prázdné hodnoty.

Jde celkem o 11 scénářů ve třech testovacích funkcích. Testy požadavků používají
`httptest.NewRecorder` bez otevřeného socketu; procházejí propojením komponent
v paměti. Nevyžadují služby, databázi ani konfiguraci. Routy a datové scénáře
se spouštějí i souběžně pomocí `t.Parallel()`.

Ověření 2026-11-20: `go test ./...`, podrobná sada `go test -count=1 -v ./...`,
`go test -race -count=1 ./...`, měření pokrytí, `go build ./...` a `go vet ./...`
prošly. Pokrytí příkazů je 92,3 % pro `internal/web`, 0 % pro `cmd/server`
a 60 % celkem. Pokrytí není důkazem správnosti. Návod a příkazy jsou v
[README.md](README.md#automatické-testování).

Automaticky netestujeme start procesu a chyby načtení šablony v `cmd/server`,
chybu zápisu odpovědi při odpojení klienta ani síťové chování HTTP serveru
(včetně potlačení těla `HEAD` na síti). Integrace s databází, rezervace,
kolize, oprávnění a časová pásma čekají na implementaci; budoucí databázové
testy musí mít izolované testovací prostředí, nikdy produkční databázi.

## Git, rozpracovanost a omezení

- Při inicializaci paměti: větev `main`, žádné commity, žádné remotes;
  všechny projektové soubory jsou untracked. Není commitová historie k analýze.
  Toto je časový snímek, nikoli pokyn k automatickému commitu.
- Dosavadní změny se zachovávají. Zavedení testů a úpravy pravidel spolupráce
  nemění aplikační kód ani závislosti. Uživatel 2026-11-20 schválil první commit
  ověřené kostry, testů a dokumentace i následný push. Skutečný stav commitů
  vždy ověř v Gitu; vzdálený repozitář zatím není nastavený a jeho URL chybí.
- Uživatel 2026-11-20 předal výsledek nezávislého CLI review všech staged,
  unstaged a untracked souborů: bez konkrétních funkčních či bezpečnostních
  nálezů, formátování a kontroly diffu prošly. Reviewer nemohl spustit testy,
  build ani vet kvůli zákazu zápisu do cache a dočasného adresáře.
- Implementační chat po review ověřil shodu kontrolních součtů všech 14 souborů
  s předaným stavem a úspěšně zopakoval `go test -count=1 ./...`,
  `go test -race -count=1 ./...`, `go build ./...` a `go vet ./...`
  s `GOCACHE=/tmp/martina-anglicky-go-build`. Po review se změnila pouze
  projektová dokumentace; aplikační kód ani testy se nezměnily.
- Hotová je pouze výuková kostra. Neexistuje přihlášení, evidence studentů,
  dostupnost, kalendář, rezervace, e-maily ani nasazení.
- Testy dosavadní webové vrstvy neodhalily funkční chybu. Pevný port a závislost
  na pracovním adresáři jsou současná omezení kostry; zbylé mezery v testech jsou výše.

MVP zatím nezahrnuje platby, studentské účty, více učitelů ani složité integrace.
Délky lekcí, potvrzování, rušení, předstihy a chování při změnách času nejsou
uzavřené implementační specifikace. Zohlednit bude nutné Tenerife a časová pásma
studentů; návrh `Atlantic/Canary` a ukládání okamžiků v UTC se má před implementací upřesnit.

Další práce: [TODO.md](TODO.md). Pravidla relací: [AGENTS.md](AGENTS.md).
