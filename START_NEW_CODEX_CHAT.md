# Zahájení nového chatu

Otevři tento repozitář jako pracovní adresář Codexu. Projektové soubory jsou
trvalým kontextem; nový chat nesmí spoléhat na to, že zná předchozí konverzaci.
Paměť musí být dostupná v pracovním stromu; po přenosu na jiný počítač přenes
i necommitnuté soubory. Existenci commitů a remote vždy ověř podle postupu níže.

## Postup obnovení

1. Přečti platné instrukce `AGENTS.md` a případné scoped instrukce.
2. Načti `PROJECT_CONTEXT.md`, `TODO.md`, `DECISIONS.md` a relevantní README.
   Pokud již existuje schválený `DESIGN.md` a úkol se týká UI, načti také ten.
3. Ověř aktuální soubory a stav Gitu:

   ```sh
   pwd
   git status --short --branch
   git diff
   git diff --cached
   git ls-files --others --exclude-standard
   ```

4. Pokud `git rev-parse --verify HEAD` uspěje, prohlédni `git log -5 --oneline`
   a relevantní commity. Jinak jde o repozitář bez prvního commitu, ne o důvod
   k jeho automatickému vytvoření. Nové soubory nejsou ve výstupu běžného diffu;
   přečti je samostatně, případně použij `git diff --no-index /dev/null <soubor>`.
5. Ověř dokumentovaný stav proti kódu. Krátce shrň, co je hotové, rozpracované
   a jaký jeden krok bude následovat. Rozsah se řídí aktuálním zadáním.

Aktuální spuštění: `go run ./cmd/server` z kořene projektu. Další příkazy
a omezení jsou v PROJECT_CONTEXT.md. Dostupnost serveru vždy znovu ověř.

## Úvodní prompt

```text
Pokračujeme v tomto repozitáři. Nejprve načti platné projektové instrukce
a soubory AGENTS.md, PROJECT_CONTEXT.md, TODO.md, DECISIONS.md a README.md.
Zkontroluj pracovní adresář, aktuální Git stav, staged i unstaged změny,
nové soubory a dostupnou relevantní historii. Dokumentaci ověř proti kódu.
Zachovej existující necommitnuté změny.

Vývoj vedeme po malých funkčních krocích s vysvětlováním Go. Krátce shrň
aktuální stav a navazující krok; pokud níže zadávám implementaci, proveď jen
tento krok a ověř ho. Po jeho dokončení se zastav na společné projití.
Neimplementuj automaticky celý TODO.

Dodrž jazyková pravidla, rozdělení odpovědností, nezávislé review významných
změn a postup návrhu a schválení UX/UI podle AGENTS.md. Preferované, ale dosud
nepotvrzené frontendové nástroje nepovažuj za definitivní volbu.

Po smysluplné práci aktualizuj relevantní projektovou paměť. Bez výslovného
souhlasu nedělej commit; push a deployment potřebují samostatný souhlas.
Na konci shrň změny, ověření a další krok.

Můj dnešní úkol: [doplň konkrétní požadavek]
```

Aktualizaci paměti provádí Codex při práci podle AGENTS.md; na pozadí neběží
žádný synchronizační proces. Tyto dokumenty se samy nemění mimo relaci.

## Samostatný review chat

Před předáním dokonči relevantní testy a kontrolu diffu. Urči rozsah review
a pozastav změny ve sdíleném checkoutu, včetně práce ostatních chatů.
Reviewer musí pracovat nad stejným adresářem, nikoli jinou kopií projektu.

Před review i po něm porovnej Git stav, index a obsah kontrolovaných souborů.
Samotný `git status` nezachytí změnu obsahu již existujícího untracked souboru.
Kontrolní výstupy předej v chatu; nevytvářej kvůli nim soubory v repozitáři.
Pro kontrolu stavu a obsahu neignorovaných projektových souborů z kořene použij:

```sh
export GIT_OPTIONAL_LOCKS=0
git status --short --branch
git ls-files --stage
git ls-files --cached --others --exclude-standard -z |
  xargs -0 -r sha256sum --
```

Úvodní prompt pro review (nepoužívej implementační prompt výše):

```text
Jsi nezávislý reviewer tohoto repozitáře. Pracuj pouze čtecím způsobem
podle review pravidel AGENTS.md. Ověř pracovní adresář a načti platné
instrukce i relevantní projektovou dokumentaci.

Neměň žádné soubory, dokumentaci, nastavení, Git index ani historii.
Neprováděj opravy a neaktualizuj projektovou paměť. Pokud by kontrola
vyžadovala zápis nebo rozšíření oprávnění, pouze uveď omezení.

Rozsah review: [konkrétní soubory nebo změny]
Kontrolovaný stav: [Git stav a identifikace obsahu]
Výsledky kontrol implementačního chatu: [příkazy a výsledky]

Zkontroluj skutečný obsah, včetně určených untracked souborů.
Nálezy seřaď podle závažnosti; uveď soubor a řádek, konkrétní scénář
selhání a stručné odůvodnění. Nevytvářej spekulativní ani stylistické nálezy.
Pokud skutečný problém nenajdeš, výslovně to napiš. Uveď také rozsah
provedeného ověření a zbývající mezery. Potom se zastav.
```

Čtecí oprávnění nastav a ověř v konkrétní review relaci; samotný prompt
zápis technicky neblokuje. Opravy i zápis výsledků do projektové paměti
provádí původní implementační chat až po ukončení review.
