# Martina – online výuka angličtiny

Prezentační web Martiny jako online lektorky angličtiny pro české studenty.
Práce na rezervačním systému je nyní pozastavená a její základ je bezpečně
uložený ve větvi `archive/booking-foundation`.

Projekt vzniká po malých krocích jako praktické opakování Go. Každý krok
má vlastní vysvětlení a ověření; další krok následuje po společném projití.

## Projektová paměť

Pravidla práce jsou v [AGENTS.md](AGENTS.md), aktuální stav v
[PROJECT_CONTEXT.md](PROJECT_CONTEXT.md), úkoly v [TODO.md](TODO.md)
a důvody rozhodnutí v [DECISIONS.md](DECISIONS.md). Při otevření dalšího chatu
použij [START_NEW_CODEX_CHAT.md](START_NEW_CODEX_CHAT.md).

## Technologický směr

Současný prezentační prototyp používá Go, standardní HTTP server a router
`net/http`, `html/template` a vlastní responzivní CSS přímo v šabloně. Nemá
frontendový build ani externí závislosti.

MariaDB/MySQL, SQL migrace, Tailwind CSS, vanilla JavaScript a FullCalendar
patří pouze k dřívějšímu směru pozastaveného rezervačního systému. Jejich
budoucí použití se potvrdí až podle skutečné potřeby.

HTTP vrstvu stavíme bez frameworku, abychom při vývoji procvičovali přímo
standardní knihovnu Go. Ostatní závislosti budeme přidávat postupně,
až je budeme potřebovat.

## Struktura a spuštění

```text
cmd/
  server/
    main.go          spuštění serveru a sestavení závislostí
internal/
  web/
    app.go           struktura webové aplikace a její vytvoření
    routes.go        registrace rout
    handlers.go      obsluha HTTP požadavků
    web_test.go      testy rout, vykreslování a HTML šablony
templates/
  home.html          responzivní prezentační stránka
prototype/
  index.html         původní statický nástřel z cloudové větve
  martina-koncept.png pracovní portrét načítaný Go aplikací
go.mod               identita modulu a verze Go
```

Z kořenového adresáře projektu:

```sh
go run ./cmd/server
```

Stránka je na <http://127.0.0.1:8080/>. Server zastavíme pomocí `Ctrl+C`.
Pracovní adresář zůstává kořenem projektu, aby se našla šablona i portrét.

Dodržujeme tyto hranice:

- `cmd/server` připravuje závislosti a spouští aplikaci.
- `internal/web` zpracovává HTTP požadavky, registruje routy a vykresluje HTML.
- `templates` obsahuje HTML; obchodní pravidla do šablon nepatří.
- Budoucí rezervační pravidla budou v `internal/booking` a databázové
  dotazy v `internal/store`. Tyto balíčky přidáme, až začneme dané části psát.
- Routy registrujeme v `routes.go`; jejich obsluhu píšeme jako pojmenované
  metody v `handlers.go`. Až přibudou další oblasti, handlery lze rozdělit
  například na `handlers_booking.go` a `handlers_admin.go`.

## Automatické testování

Používáme standardní `testing` a `net/http/httptest`, bez externího testovacího
frameworku. Z kořene repozitáře spusť:

```sh
go test ./...
```

Testy v `internal/web/web_test.go` procházejí skutečným routerem, handlerem
a souborem `templates/home.html`. Požadavky i odpovědi jsou v paměti přes
`httptest.NewRequest` a `httptest.NewRecorder`; nespouštějí HTTP server,
neobsazují port a nepotřebují databázi, síť, `.env` ani další služby.

Pokryté jsou prezentační HTML a jeho data, portrét včetně MIME typu a cache
hlavičky, query parametr, přijetí `HEAD`, přesné směrování cest, HTTP 404 a 405
včetně `Allow`, chyba vykreslení bez úniku částečného obsahu či interní chyby,
escapování HTML a prázdná pole šablony.
Test `HEAD` ověřuje stav a hlavičky handleru; potlačení těla na síťové vrstvě
`net/http.Server` není součástí tohoto testu v paměti.

Podrobnější výpis a konkrétní test:

```sh
go test -count=1 -v ./...
go test -run '^TestHomeTemplateError$' -v ./internal/web
```

Další relevantní kontroly:

```sh
go test -race ./...
go test -coverprofile=/tmp/martina-anglicky-coverage.out ./...
go tool cover -func=/tmp/martina-anglicky-coverage.out
go vet ./...
go build ./...
```

Detektor datových závodů `-race` potřebuje podporovanou platformu, zapnuté CGo
a C překladač; běžné testy další nástroje nepotřebují. V omezeném prostředí
lze před příkazy nastavit `GOCACHE=/tmp/martina-anglicky-go-build`.

Spuštění procesu v `cmd/server`, chyba odeslání odpovědi při odpojení klienta
a skutečná síťová komunikace zatím nemají automatické testy. Obchodní logika
rezervací, přihlášení, práce s časem, databáze ani API rezervací ještě nejsou
implementované; jejich testy budou vznikat spolu s nimi. Budoucí databázové
integrační testy musí používat samostatné testovací prostředí, nikdy produkci.
CI/CD zatím nezavádíme; testy jsou připravené ke spuštění standardním Go příkazem.

## Krok 1: základ Go projektu

V prvním kroku jsme založili Git repozitář a Go modul. Program vypsal
zprávu do terminálu.

- `go.mod` určuje identitu modulu a požadovanou verzi Go.
- `main.go` obsahoval vstupní funkci programu (od kroku 4 je v `cmd/server`).
- `.gitignore` vynechává lokální konfiguraci a generované soubory z Gitu.

Cesta modulu `github.com/skybedy/martina-anglicky` je zvolený identifikátor;
remote `origin` odkazuje na stejnojmenný GitHub repozitář.

## Krok 2: první HTTP server

Používáme pouze standardní knihovnu Go, bez externích závislostí.
Server spustíme z kořenového adresáře projektu:

```sh
go run ./cmd/server
```

V prohlížeči otevřeme <http://127.0.0.1:8080/>. V tomto kroku stránka vracela
prostý text; od kroku 3 na této adrese vracíme HTML:

```text
Martina – online výuka angličtiny
```

Server běží, dokud jej v terminálu nezastavíme pomocí `Ctrl+C`.
Naslouchá pouze na lokálním rozhraní `127.0.0.1`.

- `http.NewServeMux()` vytvoří router.
- `mux.HandleFunc("GET /{$}", handler)` zaregistruje obsluhu kořenové adresy.
  Zápis `{$}` označuje konec cesty, takže například `/rezervace` tuto routu
  nepoužije. Metodě `GET` odpovídá také `HEAD`.
- `http.ResponseWriter` slouží k nastavení hlaviček a zápisu odpovědi.
- `*http.Request` obsahuje příchozí požadavek.
- `fmt.Fprintln(w, ...)` zapíše text do HTTP odpovědi. Při prvním zápisu
  se bez výslovného nastavení jiného stavu automaticky použije HTTP 200.
- `http.ListenAndServe(...)` spustí server, který předává požadavky routeru.
- Router automaticky vrátí HTTP 404 pro neznámou cestu a HTTP 405 pro
  nepovolenou metodu na známé cestě.

Soubor `go.sum` zatím nepotřebujeme, protože nemáme žádné externí moduly.

HTTP odpověď lze ověřit z dalšího terminálu:

```sh
curl -i http://127.0.0.1:8080/
```

## Krok 3: první HTML šablona

Úvodní stránka nyní používá `templates/home.html`. Server spouštíme ze
kořenového adresáře projektu příkazem `go run ./cmd/server`.

- `homePageData` je struktura s titulkem a popisem stránky.
- `template.ParseFiles(...)` načte a zpracuje šablonu jednou při startu.
  Pokud je soubor nedostupný nebo má chybnou syntaxi, server se nespustí.
- `homeTemplate.Execute(&output, data)` při každém požadavku vloží data
  do šablony. `{{.Title}}` a `{{.Description}}` odkazují na pole struktury.
- `html/template` automaticky escapuje běžné řetězce podle jejich umístění
  v HTML; například znak `<` v textovém obsahu se nezmění v HTML značku.
- `bytes.Buffer` dočasně drží vykreslený obsah v paměti. Teprve po úspěšném
  vykreslení nastavíme HTML hlavičku a odešleme stránku. Při chybě vykreslení
  tak můžeme vrátit HTTP 500 místo částečné stránky se stavem 200.

Po úpravě šablony, portrétu nebo Go kódu je potřeba server restartovat.
Současné responzivní stylování je pro jednoduchost přímo v šabloně.

## Krok 4: balíčky, routy a pojmenované handlery

Vstupní program je v `cmd/server/main.go` a má `package main`. Webová část
je samostatný balíček `web` v `internal/web`; všechny jeho Go soubory patří
do stejného balíčku a sdílejí jeho deklarace.

`main` načte šablonu i pracovní portrét a předá je do
`web.New(homeTemplate, martinaPortrait)`. Konstruktor vrátí `*web.App`.
Struktura drží závislosti, které potřebují její handlery. Šablonu ani obrázek
tak nemusíme ukládat do globálních proměnných.

Metoda `app.Routes()` vytvoří router a vrátí jej jako `http.Handler`.
Registrace `mux.HandleFunc("GET /{$}", app.home)` předává routeru
pojmenovanou metodu bez jejího okamžitého zavolání.

Handler `func (app *App) home(...)` má přes receiver `app` přístup
k načtené šabloně. Malé `home` a `homeTemplate` znamená dostupnost uvnitř
balíčku `web`; velká `App`, `New` a `Routes` jsou dostupná i z `main`.

`cmd` je běžná adresářová konvence. Omezení adresáře `internal` vynucuje Go:
jeho balíčky mohou importovat jen balíčky ve stromu nadřazeného adresáře.
Zde tedy slouží jako vnitřní implementace tohoto modulu.

Další krok: společné projití nové struktury a předávání závislostí.
