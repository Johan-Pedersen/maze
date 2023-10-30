# tanker (Maze)

-  Hvad skal end result vœre
  - Der skal vœre en single path fra start til slut
  - Der skal så vœre en masse blind-gyder / dead ends
  - De skal vœre rimlige narrow, så der er ikke så meget ekstra spild plads.
  - Der skal vœre meget lille sandsynlighed for at der bliver lavet en lige linje fra start til finish linje.

- Problemet er at selv om vi laver en dereference, så ligner det stadig at det er det samme objekt i memory der bliver rettet i. Og hvordan kan det vœre
  - Det er fordi det var en slice, og når man laver := med en slice, så er det det samme underliggende array.

## tekniske go ting

- Hvordan kører man et go program
  - Det skal først compiles før man kan kører executable programmet
  - Hvilken path er det man skal bruge når man bruger install commandoen
  - Hvad gør "go install"
  - Hvilken path bliver programmerne installet på

  - Hvilke packages skal jeg have
    - spørgsmålet er hvordan programmet skal opdels
    - Fordi det kan vel ikke hedde main det hele
    - 
    - Hvad er gode package convensions
      - Det er filer der har noget tilfœlles
      - Så man vil have en maze-gen package

- Hvilke packages skal jeg have
  - spørgsmålet er hvordan programmet skal opdels
  - Fordi det kan vel ikke hedde main det hele
  - 
  - Hvad er gode package convensions
- Hvorfor har man så mange funktioner der kører på pointers i stedet for 

### pointers in go

- Hvad er pointers
  - Det er adresser i memory
  - Så når man bruger pointers, så er man sikker på at man rammer prœcis det samme objekt som man gjorde før 
  - Man henter en variabels memory address med &(address operatoren)
  - \*<type> er pointer typen for type <type>.
    - Så denne memory adresse kan kun bruges til variabel med type <type>
  - \*<var> bruges til at lœse vœrdien der ligger på en pointer



## Nvim setup

- gopls settings
  - https://github.com/golang/tools/blob/master/gopls/doc/settings.md
- gopls analyzers
  - https://github.com/golang/tools/blob/master/gopls/doc/analyzers.md
- inlay hints
  - https://github.com/golang/tools/blob/master/gopls/doc/inlayHints.md
- Ending ")" når man sœtter "("
- window management
  - 1 (2/3)
  - Og så hvis jeg lukker 3, så skal man nemt kunne komme tilbage til samme layout
- Luk term
  - Det kan vœre toggleterm er nice
- mangler en linter
  - Får fx ingen warning, når jeg har en unused variabel.
  - Har sat "unusedvariable": true. Men får stadig ingen warnings
  - Det kan vœre at det er fordi den ikke bliver "bygget"
- Mangler den der "gh" popup diagnostics
- Staticcheck bliver heller ikke brugt, så det er som om den ikke rigtig bliver brugt.
- Den bruger nok heller ikke gofumpt 

- log-level debug: siger den at der ikke er nogen inlineHintProvider 

- undeclaredname virker heller ikke
  - Her er det meningen at den skal definere en ny value, hvis den ikke finder en variabel/funktions declaration
  - Denne skulle jo også vœre enabled by default, så forstår ike helt hvorfor den ikke er der.

- Kan det vœre fordi jeg ikke har nogen code actions
  - Code actions virker faktisk fint
  
- Vi kan teste om gofumpt bliver brugt ved at lave et eksempel
  - Og så kan vi vel se om det er static check der bliver brugt
  - Det virker som om det er gofumpt bliver brugt

- Bliver staticchecker brugt
  - Så skulle man i hvertfald aktivere den på en eller anden måde
  - Men det kan vœre at det er nødvendigt at aktivere den med nogen kommandoer
    - Men burde det ikke stå et eller andet sted, hvis man skal det?
- Diagnostics er bare tom
  - Men den burde vel have warnings 
  - Man kan sikkert godt fixe det med linters, men det er bare en fucking billig pis lort løsning, fucking drœb dig selv din forpulede lorte fucking taber

  - Er det fordi jeg ikke har en namespace, som de bliver knyttet til
    - Men jeg får jo errors vist, og hvis ikke jeg havde et namespace, så ville der jo ikke blive vist noget som helst
    - Men tror måske at grunden til at jeg får den fejl, med unused, er fordi man ikke kan have nogen unused variable i go.
  - Det virker som om jeg kun får vist diagnostics som har severity.ERROR
  - Der bliver også kun sendt diagnostics som har severity.ERROR

- Hvis jeg selv skulle prøve at add errors til mit namespace, så skal jeg jo vide hvilket namespace det er
  - Måske er det fordi de ekstra plugins bruger det forkerte namespace?


- Skal bare have bare minimum for at nvim-cmp fungere
  - snippet source
  - 


## tasks

- Man kunne vel godt sige at man starter med walls alle steder og så "hukker" man en path ud af det.
  - På den måde er det nok lidt nemmere at definere paths. Fordi så er der kun de stier du har lavet resten er bare vœg.
  - Hvor i mod hvis man startede blankt og så satte vœgge, så skal man jo huske at sœtte vœg på hver side af stien.
  Og der kommer til at vœre meget blank space, som ikke er en tiltœnkt sti.
- Man kunne godt prœdefinere en rœkke patterns hvor man ligesom laver granit vœgge, som er umulige at slå i stykker. Så man sœtter sådan et pattern
- Og så laver man en breadth-first søger, som ved hver step med semi uniform tilfœldighed vœlger ny vœg at slå ned.

- Hvilke edge cases har vi
  - Hvad hvis man stien bare cirker tilbage til den path man allerede har hugget.

- Hvordan vil det teoretisk se ud hvis vi ikke havde nogen hard-blocks
  - Så vi starter bare i den første block og laver breadth-first expansion
  - Den vœlger mellem 2 paths og hver er så med uniform ss
  - Hvad vil den gå imod
    - Så hver side har uniform sandsynlighed for at blive ramt
    - 1 runde bliver 1 valgt
    - Vi har et spørgsmål der hedder, hvilken hvad skal vores endpoints vœre
    - Fordi vi bliver jo nødt til at have flere grene 
    - Hvis det skulle vœre rigtig breadth first search så har man jo 1 endpoint for hver knude punkt
    - Men hvis vi gør det så vil der jo ikke vœre flere vœgge tilbage
    - Så skulle man have flere depth-first expansions
      - Hvornår skulle man så starte dem
        - Så kan man sige at hver block har 10% ss for at spawn en ny path. Som så kan kører i x antal steps, for at lave blind gyder.
          - De skal jo så også have en ss for at de kan spawn blindgyder
          - Så har man så et evigt system, som kan blive ved med at spawn paths. Og det kan vœre at den rammer target først
          - Men det kan man jo sådan set også sagtens, så lœnge der er bare en der rammer

- Vi vil gerne have en maze hvor der ikke er tager for meget før man finder en path
  - Hvilke elementer spiller ind i hvor meget der bliver taget før der bliver fundet en path
    - Distance fra start location
    - Antal porte
    - Størrelsen på portene
    - Størrelsen på target area
    - Hvor mange active paths man har i en given runde
    
- Hvordan vil en potentiel generering se ud
  - Tror der er stor risiko for at der kommer til at der kommer til at vœre store områder i labyrinten som er taget
  - Så man kunne vel godt have en vektor, der peger mod target zone. Så man ligesom prøver at incentivize at gå mod target zone
  - På den måde kan man vœgte sandsynlighederne til at pege mod target zone.
    - Så er spørgsmålet om man kan så bare vil lave en masse veje, der bare alle går mod target zone.
      - løsningen er jo at vœgte den en lille smule, så man i enden får vejene til at gå mod target. Men det skal stadig vœre muligt at man ligesom kan gå ud af
      - Så er spørgsmålet, hvor stor skal labyrinten vœre før at man ligesom kan have den effekt.
      - Fordi det skal vœre sådan at man bare nudger mod target, men gør det ikke så voldsomt.
- Hvordan ville man regne ud hvad ss for at ramme target zone efter x steps.

- før man kan lave vector udregning, så skal vi lige vœre enig om representationen af matricen
  - Vi bliver jo nød til at lave basen først
- Labyriten er jo bare en matrice

- Så skal vi bare finde ud af hvordan vi vil gøre det ikke mega dårligt optimerings wise
  - Der er jo ingen grund til at vi skal regne en hel matrice, ud når vi kune skal bruge 2 punkter
  - For at lave en vektor skal vi jo kun bruge 2 punkter 
  - Men hver punkt er jo en (x,y)
  - Så hvordan vil man lave en vektor baseret på det?
  - 

  - I en lidt simpel løsning, så bliver man jo nødt til at have en global matrice, som alle paths ligesom relatere til 
    - Hvis man definere sine paths med 0, så vil paths aldrig krydse
      - Det kan man bare fixe ved at œndre vœgtning for en path.
        - Det kan jo dog nok give problemer i forhold til at gå tilbage af.
        - Men så kan man måske sige at path'en selv har vœgt 0. Men de andre paths har vœgt >0.
        - Future problem
    - Hver path, har så et head koordinat i matricen
    - Hver path kan vœlge mellem 1 og 3 felter
    - Og vi vil vœgte de paths der peger mod target zone højere
    - Det vil vi gøre på en lidt gradual måde
    - Vi vil jo vœgte det felt er peger mest i retning af target zone højest.

    - En vektor virker stadig ikke som en dårlig ide. Den skal bare lige udføres på den rigtige måde.
      - Og vi vil ikke regne vektor vœgtningen for hele matricen, for hver path
  
- Vektor justering
  - Fordi det giver en retning
  - Hvordan får vi en retninge fra head til target
    - ((x1,y1),(x2,y2)) -> target vektor = ((x1-x2),(y1-y2)).
      - Enheds vektoren = ((x1-x2),(y1-y2))/||((x1-x2),(y1-y2))|| = (z,q)
    
  - Hvordan vil vi så bruge (z,q) til justeringen
    - Man kan jo så bruge vektor produkt, til at finde ud af hvilken en der peger mest i samme retning 
      - Man kan måske lave vœgtningen baseret på vektor produktet.
      - Så skal man dividere med max vektor produkt. For at få en procent. Man kan lœgge til.
        - Men det kan man 

- Vi kan give dimentionerne for labyrinten. Som så sørger for at lave en labyrint af 1'ere

- Så hvad skal vi gøre
  - Der er hele path creation delen
  - Så man skal kunne vœlge hvilken celle man vil gå af. 
  - Det er jo det første vi skal lave og så laver vi vector justering bagefter


### create path

- Alt det vektor stas er ikke endnu

- Vi flytter på enten x el. y aksen

- Så vi har vores maze og vi starter så med 1 pathCoordinate entry

- Hvordan vœlger vi hvilken path vi skal vœlge
  - Så skal man vel ud fra de omkring liggende celler med vœrdi 1 så vœlge ud fra dem
  - Vi løber igennem de omkring liggende celler og notere hvilke koordinater der er = 1
  - Så skal vi have dem indexeret, på en måde så vi kan vœlge en vilkårlig
- Hvordan kan matrix holde på vœrdien \*Dense 
  - \*Dense er jo memory location for Dense typen
  - Har det noget at gøre med at matrix er et interface?
  - Fordi go har jo ikke subtypes, så det har jo heller ikke noget med dig at gøre.
- Men er der noget med hvornår en struct implementer et interface

- Vi kan lave Enums for Left, Right, Up, Down, med vœrdier 0,1,2,3
- Fordi så kan vi finde en random value mellem 0 og 3. 
- Så kan vi tage den vœrdi se om den er out of bounds og hvis ikke den er, så tjek at det ikke er en 0 path.

- Hvad vil vi
  - Vores problem er bounds. Så man må ikke gå out of bounds. Hverken på x eller y aksen
  - Final result skal gœlde at X + step <= Maze.Dims(X) and X+step >=0
  - Final result skal gœlde at Y + step <= Maze.Dims(Y) and Y+step >= 0
  - Vi skal så først lave tjekket for x-aksen når der rykkes på x-aksen

- Hvordan kan vi få en fejl i Set metoden, efter hele programmet har kørt
  - Det må jo vœre fordi at Set er en form for promise der kører synkront med resten af koden
  - createPath({0x0, {0x0, 0x0}, {0xc000115e50, 0x1, 0x1}})
  - Fejlen har noget at gøre med at vores Maze er en \*Dense type, men vi bruger den som en Maze

  - Det er sjovt at det "fungere" uanset om vi bruger pointeren eller den underliggende variabel

- Så er spørgsmålet hvordan kan man funktion knyttet til en pointer
  - Med en pointer siger man vel bare at når det er af den her type, så fylder det X bits i memory
    - Så er det en form for "nedarvning"
  - Hvordan har de gjort så Set funktionen både kan kaldes fra Matrix, \*Dense og Dense typen
    - Set kan kaldes med \*mat.Dense og mat.Dense
  - Der findes en matrix interface type
    - Hvad får man af den
    - Matrix interfacet har metoderne dims og T 

- Hvordan vil vi autogenerer target zone
  - Vi vil ikke have det i starten af matricen
  - Men man vil heller ikke altid have det i midten. Så man kan måske sige at target zone, skal vœre lœngere vœk end 50% af banelœngden, vœk fra entry 
  - Så er spørgsmålet bare er det for kort/ kommer den for tœt på. Men det kan man jo bare justere på
  - Spørgsmålet er så mere om der er en bedre måde at gøre det på
    - Hvad skulle problemet vœre i denne løsning
   
  - Vi skal bare vide hvilken "kasse" target Zone er i
    - Vi kan vel bare 0-indexere col

    - 1Dlisten er 0-indexeret
    - matricen er 0-indexeret
    - Dims vi får er 1-indexeret
    - Så når vi ligger col til så første gang vil vi jo ende på index col + 1  
    - Så col * y giver os starten af den rigtige kasse
    - + x giver os så det rigtige index
    - col * y + x
  

- Targetzone
  - hvad skal den gøre
    - Fordi vi kan jo ikke bare nøjes med en target zone af 1x1, hvis der skal vœre en target zone, så skal den vœre størrer
    - Men det kan man jo også sagtns gøre

- Path creations
  - Problemet med ripples er at procentvis er det stadig stort set en uniform fordeling.
  - Dette leder til at man ikke rigtig får en retning på vores paths, så den bliver bare meget tyk 
  - Så hvis vi laver ripples, så skal det belønnes meget mere at komme tœttere på target endnu.

  - Tror vektor metoden er lidt bedre fordi det er nemmere at skrue på for at få den rigtige retning uden at skulle komme ud i nogen mega store tal og sådan noget shit.

#### path adjusting

- Vektor justering
  - Fordi det giver en retning
  - Hvordan får vi en retninge fra head til target
    - ((x1,y1),(x2,y2)) -> target vektor = ((x1-x2),(y1-y2)).
      - Enheds vektoren = ((x1-x2),(y1-y2))/||((x1-x2),(y1-y2))|| = (z,q)
    
  - Hvordan vil vi så bruge (z,q) til justeringen
    - Man kan jo så bruge vektor produkt, til at finde ud af hvilken en der peger mest i samme retning 
      - Man kan måske lave vœgtningen baseret på vektor produktet.
      - Så skal man dividere med max vektor produkt. For at få en procent. Man kan lœgge til.
        - Men det kan man 

- For at finde retningen mod target zone, så kunne man sige at hver celle har en bestemt vœgt, hvor 1 er default
- Og jo tœttere du kommer på target zone, desto større vokser vœrdierne
- Så man kan gange hver retning ganger man cellens vœgtning på

- ripples
  - Hver celle skal have en vœrdi der er korreleret til denne celles afstand fra target zone.
  - Celle vœrdien skal bare afhœnge af hvor langt cellen er fra target zone
  - Det er jo bare en matrice, så man vil jo kunne lave en form for matrix multiplikation ligesom man gør i ml
  - Er det nødvendigt, vi har jo en vektor mellem dem, så vi kan jo nemt bare udregne distancen 
  
  - Hvordan skal vi så omforme det til noget man kan bruge til at justere hvilken dir der bliver valgt.
    - Det skal jo vœre noget der kan omdannes til procent.
    - Vi kan finde 2-normen af vektorne fra hver celle til TZ (target zone). Og så ud af de mulige directions finder man summen normerne
    - På den måde kan man sœtte vektorne i forhold til hinanden og samligne dem i forhold til hvilke nogen der er størst.


- Hvad er funktionen til et one-step increase i forhold til normen
  - Det er bare en simpel funktion hvor du fastholder enten x1 el. x2
    $$f(x) = sqrt(x^2 + x2^2)$$
  - Kan man så lave en vectorize funktion baseret på denne funktion
    - Der er ingen vectorize funktion i gonum
    - 

- Problemet med at hver celle er normen ind til target, så skal man lave en udregning når man kommer hen til cellen og det er lidt grimt.
- Men det er måske ikke så slemt alligevel.

- Hvordan vœlger vi en dir baseret på en ss 
  - Hver dir kunne have en tilknyttet ss
  - Men så skal man finde ss hver gang man taget et step
  - Men det gør vi jo alligevel
  - 
- Hvad kan man ellers gøre
  - Man skal jo have en metode til at vœlge og så handler de der ss sådan set bare om hvor de ligger og det er som sådan ikke så vigtigt.

- Vi har et array af ss, hvordan vœlger en entry baseret på vœrdien af en entry.
  - Det handler jo om at vœlge en entry ud fra den distribution vi har fundet.
  - Hvis man skulle lave sådan en metode selv, så er det eneste problem at man skal have en kilde/source til randomness. 
  - Det man mangler er en form for picker der vœlger en entry ud fra den givne distribution.
  - Det er en diskret fordeling
  - Det virker til at vi kan bruge stats package (Ved ikke om jeg rent faktisk kan bruge den, den står som indirect i min mod fil) til at definere min egen fordeling.

  - Man kan genere et random tal mellem 0 og 1 og så kan man se hvilken range den lander inden for

  - Kan det vœre fordi at maze kun bliver opdateret lige før print maze bliver kaldt. 

- Hvornår bliver der allokeret memory
  - Når man laver en derefrence s∑ allokere man memory, somi så bare kopiere det pointeren peger på.
  - 

#### sampling

- Genere et random tal mellem 0 og 1.
- Den range vi lander inden for er den celle vi vœlger
- Hvordan gør vi så det?
  - Ligesom i playground

- Hvad skal vi returnere 
  - Hvad bruger vi
  - 
- Vi skal finde alle valid paths
- find ss for hver step
- Hvordan mapper vi en indgang til et step
  - Så skulle vi sige at step nr'et bruges som index og så returnere vi bare "i"
  - Vi skal bruge X,Y
  - Men det kommer jo fra step direction, så ved det her head ser man om step direction. Og fordi vi allerede har valideret stepet så er det bare at gøre det.

- Hvad skal vi have for at kunne lave en path
  - Vi skal have en X,Y koordinat
  - Den får vi fra vores head.X, Heady.Y, som er en path coordinate, hvilket er det vi operer ud fra.
  - Så vi tager et step og returnere en PathCoordinate 
  - Men vi bruger jo en pointer til mazen, så vi kan opdatere den direkte. Somehow syntes jeg nœsten mere om at returnere et index og så opdatere den hvor jeg kaldte den fra.
  - Jeg har lagt op til at bruge pointers og det er bedre i forhold til at man ikke kopiere data hele tiden.

- Det ville vel vœre smart hvis man havde en update funktion som både opdaterede head og satte 0 i maze'en.

- Der er mange situationer hvor alle dirs har en prob (selv om der kun burde vœre max 3)
probs: [4/4]0xc000117b20
0.3086088907313406 0.19139110926865935 0.19139110926865935 0.3086088907313406 
dir: 3


### Tests

- Hvordan skal vi lave tests



## Fremtidige ideer / todo

- Lav bane patterns, så man sœtter nogen punkter som justerings vektorene peger på, for på den måde at man også kan lave nogen lidt interessante patterns
- Man kan måske lave nogen små grupper af paths i hjørnerne osv. Så pathen måske ikke er så tydelige
- Man kan også lave patterns med blokke der ikke kan laves path igennem.
- Target zone skal ikke kunne laves i den 1/4 af banen hvor start lokation er.
- Paths, skal vœre en array af pointers, så man kan opdatere memroy location direkte
  - Hvad ville fordelen vœre ved dette
    - Fordelen er hvis man har flere steder der refere til den samme memory location, Så kan alle vide hvor alle er
      - Det er smart at mzTrack er en pointer fordi den vel du ikke kopiere rundt alle steder
      - Men man vil alligevel altid kun have 1 Maze objekt, så man behøver ikke rigtig have denne form for sync mellem steder.
      - At bruge pointers vil generelt give lidt flere mulig og så lœrer man også noget nyt så lad os gøre det
- Kunne lave skrå paths.


## Teoretiske spørgsmål

- Hvor mange steps tager det i snit før man har dannet en path til target
  - Jeg er ikke sikker på at dette ville vœre et sœrlig godt mål
  - Vi vil jo gerne kunne regne på hvor meget af labyrinten der bliver taget før man har en path til target.
- Hvad er ss for at en off-shoot path rammer target før main-path
- Hvad ville vœre et godt mål for en god labyrint

## Potential problemer

- Pathen cirkler rundt og rammer sig selv, men det sker der vel som sådan ikke noget ved
- Hvordan sikre vi at hele labyrinten ikke bliver taget før der kommer en path
  - Det kommer man jo hurtig ud for 
    - En løsning er at lave target større, så der ikke kun er 2 veje ind. Men adskillige
    - Man kan vel prøve at regne på. hvad ss er for at man har ramt target efter x steps 
- Der skal vœre en hard border rundt om som man ikke kan gå igennem




