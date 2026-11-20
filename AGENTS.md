# Pravidla pro práci s Codexem

## Zahájení relace

1. Ověř pracovní adresář (`pwd`) a přečti platné nadřazené a lokální projektové
   instrukce. Před prací v podadresářích zkontroluj i jejich scoped instrukce.
2. Načti [PROJECT_CONTEXT.md](PROJECT_CONTEXT.md), [TODO.md](TODO.md),
   [DECISIONS.md](DECISIONS.md) a relevantní části [README.md](README.md).
3. Zkontroluj `git status --short --branch`, staged i unstaged diff a nové soubory.
   Prohlédni relevantní historii, pokud existuje; absence commitů není chyba k opravě.
   Samotný `git diff` neukazuje obsah untracked souborů, ty přečti zvlášť.
4. Ověř dokumentaci proti aktuálnímu kódu a konfiguraci. Při rozporu rozliš skutečný
   stav a dohodnutý záměr; chybějící informace označ jako neověřené či nerozhodnuté.
5. Zjisti rozpracované úkoly. Nový chat sám o sobě není pokyn implementovat celý TODO.

## Spolupráce a implementace

- Projekt slouží i jako osvěžení Go. Pracuj po malých funkčních krocích.
  Před významnější implementací stručně popiš postup; po kroku vysvětli důležité
  koncepty a dej prostor na otázky. Nepřecházej automaticky na další etapu.
- Odpovídej česky. Přizpůsob vysvětlení programátorovi, který se vrací ke Go.
- Uživatel definuje požadavky a obchodní pravidla, rozhoduje o významných
  architektonických a technologických volbách a posuzuje výsledky. Codex analyzuje,
  navrhuje, implementuje, testuje a připravuje dokumentaci, review a návrhy commitů.
- Konkrétní koncepční rozhodnutí řeš primárně v tomto repozitáři a zapisuj je
  do projektové paměti. Podklady z jiných chatů porovnej s aktuálním stavem.
  Významné návrhy nejprve předlož uživateli a implementuj po jejich odsouhlasení.
- Respektuj aktuální architekturu a rozhodnutí. Technologický plán a existující
  závislosti ověřuj v repozitáři; pravidla práce platí i při budoucí změně stacku.
- Routy registruj odděleně od pojmenovaných handlerů. Dodržuj odpovědnosti balíčků
  popsané v kontextu a README. Nové vrstvy přidávej až podle skutečné potřeby.
- Nepřidávej frameworky, závislosti nebo infrastrukturu bez důvodu v zadání.
- Dotaz na vysvětlení, diagnostiku nebo review sám neopravňuje k opravě kódu.
  Měň jen soubory potřebné pro zadaný úkol a zachovej nesouvisející změny.

## Jazyková pravidla

- Komunikace a všechny projektové Markdown dokumenty včetně tohoto jsou česky.
- Zdrojový kód, názvy funkcí, proměnných a technické identifikátory jsou anglicky.
- Komentáře v kódu a Git commit messages jsou anglicky.
- Jazyk uživatelského rozhraní se rozhoduje samostatně; současná česká stránka
  není automatickým rozhodnutím o všech budoucích jazycích aplikace.

## Návrh uživatelského rozhraní

- UX/UI řeš od začátku pro administraci Martiny i veřejné rezervace studentů.
  Návrh má být přehledný, konzistentní a dobře ovladatelný na mobilu, s minimem kroků.
- Postupuj: společný návrh uživatelských procesů a obrazovek → schválení koncepce
  → případný funkční vizuální prototyp → schválení řešení → implementace.
- Neimplementuj celý frontend ani nepřidávej dosud neschválené frontendové
  technologie jen na základě obecného směru. Preferovaný stack odlišuj od finální volby.
- `DESIGN.md` vytvoř až s prvním odsouhlaseným návrhem a udržuj v něm pravidla UI.
  Do té doby zaznamenej otevřené otázky do TODO, nevytvářej prázdné dokumenty.

## Bezpečnost a Git

- Nikdy automaticky neprováděj commit, push ani deployment bez výslovného souhlasu.
  Souhlas s commitem neopravňuje k pushi ani deploymentu; vyžadují samostatný souhlas.
  Nestaguj soubory jen kvůli kontrole diffu.
- Nevkládej do dokumentace, logů ani Gitu hesla, tokeny, osobní údaje studentů
  nebo obsah lokálních tajných konfigurací. Příklady musí být anonymní.
- Nespouštěj destruktivní Git operace a nemaž či nepřepisuj uživatelskou práci
  bez jasného oprávnění. Před restartem procesu ověř, že patří tomuto projektu.
- Nevymýšlej výsledky kontrol, stav nasazení ani dostupnost běžících procesů.

## Příprava a návrh Git commitů

- Pracuj v malých, logicky oddělených celcích. Jeden commit má představovat
  jeden smysluplný dokončený úkol.
- Po dokončení každého smysluplného úkolu aktualizuj relevantní projektovou
  dokumentaci, spusť dostupné relevantní testy a zkontroluj Git diff včetně
  staged i untracked souborů podle pravidel ověření níže.
- U významných změn dokonči nezávislé code review podle pravidel níže před
  předložením commitu ke schválení. U drobných změn postačí relevantní testy a diff.
- Pokud jsou změny dokončené a připravené k uložení, automaticky navrhni Git
  commit: uveď konkrétní soubory a připrav stručnou commit message v angličtině
  ve formátu Conventional Commits s vhodným prefixem `feat:`, `fix:`,
  `refactor:`, `docs:` nebo `test:`.
- Do navrženého a případně schváleného commitu zahrnuj pouze soubory související
  s dokončeným úkolem. Nepřibírej ostatní necommitnuté změny; používáš-li soubor
  se změnami více úkolů, zahrň jen příslušné části.
- Pokud testy selžou nebo změny nejsou dokončené, upozorni na to před návrhem
  commitu a neoznačuj je za připravené. Pokud testy chybí nebo je nelze spustit,
  výslovně uveď omezení ověření.
- Návrh commitu není jeho provedení. Nikdy automaticky neprováděj `git commit`,
  push ani deployment bez výslovného souhlasu uživatele.

## Ověření a code review

- Při implementaci nové funkcionality vytvářej odpovídající automatické testy.
  Při opravě chyby pokud možno nejprve vytvoř regresní test, který ji prokáže.
- Používej primárně standardní `testing` a `net/http/httptest`. Testuj existující
  dokončené chování, úspěšné, chybové a hraniční scénáře; rezervace, kolize,
  oprávnění a časová pásma pokrývej spolu s jejich implementací.
- Integrační testy musí používat izolované prostředí. Nikdy je nepřipojuj
  k produkční databázi ani nepřebírej produkční konfiguraci jako výchozí hodnotu.
  Potřebné služby, testovací konfiguraci a úklid dat zdokumentuj.
- E2E testy přidávej tam, kde mají smysl pro skutečně implementované uživatelské
  procesy. Nový testovací framework ani CI/CD nezaváděj bez skutečné potřeby a zadání.
- Průběžně spouštěj relevantní testy, build a statické kontroly podle povahy změny.
  Aktuální příkazy jsou v PROJECT_CONTEXT.md; při změně stacku je aktualizuj.
- Po každé významné změně spusť relevantní testy; před návrhem commitu ověř
  stav odpovídající testovací sady. Pokud test selže, zjisti příčinu a transparentně
  uveď výsledek i případné nevyřešené selhání.
- Nikdy neupravuj testy pouze proto, aby prošly. Změnu nesprávného očekávání
  zdůvodni specifikací nebo opravenou chybou testu. Existující testy nemaž
  bez zdůvodnění a nepřepisuj obchodní logiku jen kvůli pohodlnějším testům.
- Testuj hlavně chování a rizikové větve. Nepřidávej testy jen pro formu.
  Pokud testy neexistují nebo kontrola nejde spustit, výslovně to uveď.
- Před dokončením vždy uveď spuštěné testovací příkazy, výsledky a oblasti,
  které zůstávají neotestované. Procento pokrytí samo nedokazuje správnost.
- Před dokončením projdi `git diff`, `git diff --cached`, `git diff --check`
  a obsah nových souborů. Bez commitů kontroluj nové soubory proti `/dev/null`
  pomocí `git diff --no-index`; návratový kód 1 znamená nalezené rozdíly.
- U významných změn používej cyklus: implementace → automatické testy → nezávislé
  code review → případné opravy → opakované relevantní testy → schválení commitu.
  Významné jsou zejména změny architektury, pravidel rezervací, oprávnění,
  práce s databází, souběžnosti a podstatných uživatelských procesů.
- Nezávislé review významné změny prováděj v novém samostatném nebo detached
  chatu nad stejným lokálním checkoutem. Detached chat neznamená změnu Git HEAD.
  U drobných úprav dokumentace či obdobně malých změn není samostatné review povinné.
- Před review předej přesný rozsah, stav Gitu a výsledky relevantních kontrol.
  Zahrň staged, unstaged i související untracked soubory; bez prvního commitu
  kontroluj výslovně určené nové soubory, nikoli prázdný běžný diff.
- Od předání do ukončení review neměň kontrolovaný checkout ani Git index
  či historii. Platí to pro implementační chat, review chat i další souběžnou
  práci. Pokud se stav změní, review přeruš a domluv nový kontrolovaný stav.
- Review hledá konkrétní logické a bezpečnostní chyby, problémy souběžnosti,
  ošetření chyb, architektury a chybějící důležité testy. Respektuj rozsah a formát
  zadání; nevytvářej nálezy jen kvůli stylu.
- Review chat pouze čte a hlásí nálezy. Nemění soubory včetně dokumentace
  a projektové paměti, Git index, historii ani nastavení. Neprovádí opravy,
  formátování, generování souborů, staging, commit, push ani deployment.
  Pro čtecí příkazy Git používej GIT_OPTIONAL_LOCKS=0.
- Reviewer uvádí konkrétní opravitelné problémy podle závažnosti. U každého
  uvede soubor a řádek, scénář selhání a stručné odůvodnění. Pokud skutečný
  problém nenajde, výslovně to uvede. Vždy popíše rozsah, provedené kontroly
  a zbývající mezery v ověření; převzaté výsledky testů odliší od vlastních.
- Reviewer spouští jen kontroly, které nemění checkout ani externí data.
  Pokud bezpečné spuštění není možné, uvede omezení místo změny oprávnění.
- Opravy a aktualizaci projektové paměti provádí původní implementační chat
  až po ukončení review. Následují relevantní testy a podle významu oprav
  následné review. Commit navrhni až po dokončení tohoto cyklu.
  Netvrď, že nezávislé review proběhlo, pokud nebylo provedeno.

## Projektová paměť a ukončení relace

- Po dokončení každého smysluplného úkolu aktualizuj relevantní kontext, úkoly
  nebo rozhodnutí podle povahy změny.
  Aktualizace je součást práce Codexu, nikoli samostatný běžící automat.
- Kontext popisuje aktuální implementaci; TODO priority, závislosti a další práci;
  DECISIONS důvody důležitých rozhodnutí. README zachovej jako návod a výukový text.
- Piš stručně, bez přepisování chatů nebo evidence každé drobné změny. Dokončené
  úkoly přesuň do krátké historie. Datum rozhodnutí uváděj jen při známém původu.
- Při změně architektury sjednoť související dokumenty. Plán neoznačuj jako hotový.
- Na závěr shrň změny, výsledky ověření, omezení a doporučený další malý krok.
  Uveď nevyřešenou práci potřebnou pro pokračování.

Návod a prompt pro další chat: [START_NEW_CODEX_CHAT.md](START_NEW_CODEX_CHAT.md).
