# Úkoly

Aktualizováno 2026-11-27. Priority určují pořadí diskuse a práce, ne oprávnění
k automatické implementaci celého seznamu.

## Nyní / rozpracováno

- [ ] P0: Spustit Go aplikaci a projít s Martinou vizuální nástřel na `/`, nahradit
  ukázkové texty skutečnou nabídkou a rozhodnout o fotografiích, tónu a barvách.
- [ ] P1: Před zveřejněním nahradit pracovně upravený portrét originální
  fotografií a zopakovat vizuální kontrolu.
- [ ] P1: Další odsouhlasená pravidla průběžně doplňovat do `DESIGN.md` a po
  schválení konceptu rozhodnout o začlenění větve do `main`; rezervační systém
  zůstává pozastaven.

## Pozastavená práce na rezervacích

- [ ] Společně projít novou strukturu, balíčky, `web.New`, receiver `*App`
  a předávání šablony handleru. Kód přesunu je hotový; výukové projití navazuje.

## Následující malé kroky rezervací (pozastaveno)

- [ ] P1: Společně navrhnout UX administrace Martiny a veřejných rezervací:
  hlavní uživatelské procesy, potřebné údaje a chybové či prázdné stavy.
- [ ] P1: Navrhnout hlavní obrazovky a jejich mobilní ovládání. Po odsouhlasení
  koncepce založit `DESIGN.md` s pravidly rozhraní; zatím žádný návrh schválený není.
- [ ] P1: Podle potřeby připravit funkční vizuální prototyp. Finální frontendové
  nástroje potvrdit podle návrhu; stylování a implementace navazují na schválení
  řešení, nikoli na samotný dosavadní návrh používat Tailwind či FullCalendar.
- [ ] P1: S dalšími změnami doplnit zbývající testy startu a chyb načtení šablony
  v `cmd/server`; případně ověřit chybu zápisu odpovědi a skutečné síťové chování.
  Současná sada pokrývá routy, vykreslení, chybovou odpověď a escapování.
- [ ] P1: Konfigurace a připojení k MariaDB, pak první migrace. Závisí na dohodě
  o lokální databázi a parametrech; tajné údaje nepatří do dokumentace ani Gitu.

## Další etapy MVP

- [ ] P2: Uzavřít pravidla délek lekcí, potvrzování, rušení, předstihu rezervací,
  pracovních intervalů, výjimek a časových pásem před implementací rezervací.
- [ ] P2: Datový model a migrace pro administraci, studenty, typy lekcí,
  dostupnost, blokace a rezervace. Závisí na upřesnění pravidel.
- [ ] P2: Zabezpečené přihlášení Martiny a základ administrace.
- [ ] P2: Evidence studentů a ruční správa lekcí; poté administrátorský kalendář.
- [ ] P2: Výpočet volných termínů z dostupnosti, blokací a existujících lekcí.
- [ ] P2: Veřejný výběr termínu a rezervační formulář; ochrana osobních údajů.
- [ ] P2: Atomické vytváření rezervací a test souběžného pokusu o stejný čas.
- [ ] P2: S implementací přidávat testy oprávnění, časových pásem a změn letního
  času. Databázové a rezervační API integrační testy spouštět jen v izolovaném
  testovacím prostředí; zatím pro ně neexistuje implementace ani testovací služba.
- [ ] P2: Mobilní ověření, provozní konfigurace, zálohy a příprava nasazení.
  Samotný deployment vyžaduje výslovné zadání.
- [ ] P2: Podle implementovaných uživatelských procesů doplnit vhodné E2E testy.
  Nástroje vybírat až podle potřeby; nyní E2E sada ani CI/CD neexistují.

## Později, mimo MVP

- [ ] P3: Potvrzovací e-maily a připomínky, případně opakované lekce.
- [ ] P3: Integrace kalendářů a platby pouze podle budoucího zadání.

## Stručná historie dokončené práce

- 2026-11-19: Základ Go modulu a lokálního Git repozitáře; první HTTP server;
  rozhodnutí používat standardní `net/http`, odstranění Echo a jeho závislostí.
- 2026-11-20: HTML šablona s daty a bufferem; přesun vstupního programu do
  `cmd/server`, oddělení rout, handlerů a závislostí do `internal/web`.
- 2026-11-20: Zavedení projektové paměti a návodu pro nový chat. Build a vet
  prošly; v té době ještě neexistovaly testovací soubory.
- 2026-11-20: Automatické testování přes `testing` a `httptest`: 11 scénářů
  routování a šablon včetně chyb a hraničních dat. Sada, race detektor, build
  a vet prošly. Pokrytí `internal/web` je 92,3 %, celého projektu 60 %.
  Doplněna trvalá testovací pravidla; bez změny aplikačního kódu a bez CI/CD.
- 2026-11-20: Sladěna koncepce spolupráce: jazyková pravidla, role uživatele
  a Codexu, nezávislé review u významných změn a UX/UI návrh před implementací.
  DESIGN.md se vytvoří až s odsouhlaseným návrhem; commit nyní nebyl proveden.
- 2026-11-20: Uživatel předal výsledek nezávislého CLI review všech dosavadních
  souborů bez konkrétních funkčních či bezpečnostních nálezů. Reviewer nemohl
  spustit Go kontroly kvůli čtecímu prostředí; implementační chat ověřil nezměněný
  obsah souborů a úspěšně zopakoval testy, race detektor, build a vet.
- 2026-11-20: Uživatel schválil první commit ověřeného základu i následný push.
  Remote `origin` nyní odkazuje na GitHub repozitář projektu.
- 2026-11-24: Cloudová větev přidala statický vizuální prototyp prezentačního
  webu a pracovní portrét; rezervační základ byl zakonzervován ve vedlejší větvi.
- 2026-11-27: Vizuální prototyp byl převeden do Go šablony, portrét dostal
  vlastní handler a návrh je připravený ke spuštění přes `go run ./cmd/server`.
- 2026-11-27: Porovnána původní kompozice, jednoduchý obdélník a skutečný
  kruhový ořez portrétu. Zvolen svislý portrét s mírně zaoblenými rohy bez
  cedulky a posunutého barevného bloku; pravidlo je v `DESIGN.md`.
- 2026-11-27: Doplněna konfigurovatelná adresa serveru přes `ADDR` se zachováním
  lokálního výchozího portu 8080 a testem pro nasazení za reverzní proxy.
