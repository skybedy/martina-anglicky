# Kontext projektu

Poslední věcné ověření: 2026-11-27. Dokument popisuje stav souborů, nikoli
záruku, že právě běží server. Git a prostředí ověř při každé nové relaci.

## Účel a způsob práce

Aktuálním cílem je základní prezentační web Martiny jako online lektorky
angličtiny pro české dospělé studenty. Práce na rezervačním systému je
pozastavena a její výchozí stav uchovává větev `archive/booking-foundation`.
K rezervacím se lze později vrátit bez toho, aby nyní určovaly podobu webu.

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
  `bytes`, `log`, `os`. Žádný Gin, Echo ani Gorilla, žádné externí Go závislosti.
- Jednostránkový český prezentační návrh s responzivním CSS přímo v šabloně,
  bez JavaScriptu. Bez databáze a konfigurace `.env`.
- Testy v `internal/web/web_test.go` používají standardní `testing` a `httptest`.
  CI/CD ani kontejnerová konfigurace nejsou zavedené.

Pro pozastavené rezervace je plánovaná MariaDB/MySQL přes `database/sql`,
SQL migrace `golang-migrate` a případně FullCalendar. Tailwind CSS 4 a vanilla
JavaScript zůstávají pouze dřívějším směrem; prezentační prototyp používá vlastní
CSS bez frontendového buildu. Finální nástroje se zvolí až podle schváleného návrhu.

## Stav UX/UI návrhu

Na větvi `feat/martina-website` je první jednostránkový vizuální nástřel.
Původní cloudový soubor `prototype/index.html` byl převeden do
`templates/home.html`, takže jej vykresluje běžící Go aplikace na `/`.
Pracovně upravený portrét z `prototype/martina-koncept.png` server načte při
startu a zpřístupní přes přesnou asset routu. Před zveřejněním je nutné použít
originální fotografii.

Návrh, texty, kontakt, barvy ani konkrétní fotografie zatím nejsou schválené
Martinou. Po porovnání původní složité kompozice, skutečného kruhového ořezu
a jednoduchého obdélníku uživatel zvolil svislý portrét s mírně zaoblenými
rohy a jemným stínem, bez cedulky a posunutého barevného bloku. Pravidlo je
v `DESIGN.md`; zbývající vizuální směr je ještě potřeba projít s Martinou.

## Architektura a tok požadavku

- `cmd/server/main.go`: načte šablonu a portrét, vytvoří `web.App` a spustí server.
- `internal/web/app.go`: `App` drží `*template.Template` a data portrétu;
  `New` obě závislosti přijímá.
- `internal/web/routes.go`: `Routes() http.Handler` vytvoří `http.ServeMux`
  a registruje pojmenované handlery pro `GET /{$}` a portrét.
- `internal/web/handlers.go`: vykreslí HTML do bufferu a odešle portrét se
  zjištěným MIME typem. Chyby vykreslení či zápisu ošetří a loguje.
- `templates/home.html`: responzivní prezentační stránka s pracovními texty.
- `prototype/`: původní statický návrh a pracovní obrazový soubor.

Požadavek jde přes HTTP server do routeru a pojmenovaného handleru. Šablona
i portrét se načítají jednou při startu; po jejich úpravě je nutný restart.
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

Server používá výchozí adresu `127.0.0.1:8080`; proměnná `ADDR` ji může změnit,
například pro systemd službu na `127.0.0.1:18081`. Před spuštěním ověř případné
obsazení portu, neukončuj cizí proces. `go run .` po přesunu do `cmd/server` nefunguje.
V omezeném prostředí lze použít `GOCACHE=/tmp/martina-anglicky-go-build`.
Oprávnění k síti, cache či socketům řeš standardním schvalovacím mechanismem.

## Produkční nasazení

- Veřejná adresa je <https://martina-anglicky.online/>; HTTP se přesměruje na HTTPS.
- DNS domény míří na server dostupný přes místní SSH alias `vps`.
- Systemd jednotka `martina-anglicky.service` je aktivní a zapnutá po startu.
  Aplikace naslouchá jen na `127.0.0.1:18081` a veřejně ji zpřístupňuje Nginx.
- Release `55e601152c62` je v
  `/home/skybedy/apps/martina-anglicky/releases/55e601152c62`. Symlink `current`
  umožňuje atomické přepnutí verze a služba z něj načítá binárku, šablonu i portrét.
- Reprodukovatelné konfigurace jsou v `deploy/`. Certifikát Let's Encrypt je
  spravovaný Certbotem, automaticky se obnovuje a aktuálně platí do 2027-02-25.
- Po nasazení prošly `nginx -t`, kontrola aktivní a povolené služby, loopback HTTP,
  veřejné HTTPS HTML i PNG a vizuální kontrola produkční stránky v Chromiu.

Nasazení další verze vyžaduje sestavit linux/amd64 binárku, vytvořit nový
verzovaný release, přepnout `current` a restartovat službu. Žádný z těchto kroků
se neprovádí automaticky pouze na základě commitu nebo pushnutí větve.

## Stav automatických testů

Před zavedením sady nebyly žádné testy ani testovací framework. Obchodní logika
dosud neexistuje, takže testujeme dokončenou HTTP vrstvu a šablonu:

- `TestServerAddress`: výchozí lokální adresa a změna přes `ADDR` pro nasazení.
- `TestRoutes`: osm scénářů pro skutečný router, handler a šablonu; prezentační HTML
  a data, query parametr, `HEAD`, neznámé cesty a nepovolené metody včetně `Allow`.
- `TestMartinaPortraitImage`: ověří stav, MIME typ, cache hlavičku a přesná data
  portrétu vrácená asset handlerem.
- `TestHomeTemplateError`: selhání vykreslení po částečném zápisu musí vrátit
  500 a veřejnou zprávu bez rozpracovaného HTML nebo interních detailů.
- `TestHomeTemplateData`: skutečná šablona escapuje HTML v textových polích
  a zvládá prázdné hodnoty.

Jde celkem o 14 scénářů v pěti testovacích funkcích. Testy požadavků používají
`httptest.NewRecorder` bez otevřeného socketu; procházejí propojením komponent
v paměti. Nevyžadují služby, databázi ani konfiguraci. Routy a datové scénáře
se spouštějí i souběžně pomocí `t.Parallel()`.

Ověření 2026-11-27: po propojení prototypu a přípravě nasazení prošly
`go test -count=1 ./...`, `go test -race -count=1 ./...`, `go build ./...`
a `go vet ./...`. Skutečný server s `ADDR=127.0.0.1:18081` vrátil HTML 200.
Lokální i produkční podoba byla vykreslena v headless Chromiu; před nasazením
byly ověřeny také HTML a PNG odpovědi. Návod je v
[README.md](README.md#automatické-testování).

Automaticky netestujeme start procesu a chyby načtení šablony v `cmd/server`,
chybu zápisu odpovědi při odpojení klienta ani síťové chování HTTP serveru
(včetně potlačení těla `HEAD` na síti). Integrace s databází, rezervace,
kolize, oprávnění a časová pásma čekají na implementaci; budoucí databázové
testy musí mít izolované testovací prostředí, nikdy produkční databázi.

## Git, rozpracovanost a omezení

- Remote `origin` je `https://github.com/skybedy/martina-anglicky.git`.
  Aktivní pracovní větev je `feat/martina-website`; archiv rezervačního základu
  je `archive/booking-foundation`. Lokální `main` obsahuje navíc dokumentační
  commit `e242467`; před budoucím začleněním webu se musí zachovat nebo vědomě
  nahradit, nikoli bez kontroly přepsat.
- Uživatel 2026-11-20 předal výsledek nezávislého CLI review všech staged,
  unstaged a untracked souborů: bez konkrétních funkčních či bezpečnostních
  nálezů, formátování a kontroly diffu prošly. Reviewer nemohl spustit testy,
  build ani vet kvůli zákazu zápisu do cache a dočasného adresáře.
- Implementační chat po review ověřil shodu kontrolních součtů všech 14 souborů
  s předaným stavem a úspěšně zopakoval `go test -count=1 ./...`,
  `go test -race -count=1 ./...`, `go build ./...` a `go vet ./...`
  s `GOCACHE=/tmp/martina-anglicky-go-build`. Po review se změnila pouze
  projektová dokumentace; aplikační kód ani testy se nezměnily.
- Nasazený je pracovní prezentační prototyp na Go kostře. Neexistuje přihlášení,
  evidence studentů, dostupnost, kalendář ani rezervace.
- Testy dosavadní webové vrstvy neodhalily funkční chybu. Adresu lze změnit přes
  `ADDR`; závislost na pracovním adresáři zůstává současným omezením.

MVP zatím nezahrnuje platby, studentské účty, více učitelů ani složité integrace.
Délky lekcí, potvrzování, rušení, předstihy a chování při změnách času nejsou
uzavřené implementační specifikace. Zohlednit bude nutné Tenerife a časová pásma
studentů; návrh `Atlantic/Canary` a ukládání okamžiků v UTC se má před implementací upřesnit.

Další práce: [TODO.md](TODO.md). Pravidla relací: [AGENTS.md](AGENTS.md).
