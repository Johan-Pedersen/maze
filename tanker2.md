# Tanker2 (Maze)


## productSum == 0

- Hvornår er productSum == 0
- Hvordan skal vi fixe det
  - Vi kan vel bare sige at hvis productSum = 0, så kan det vœre fordi at der ikke er nogen sti man kan gå

- 2. path1 = -path2
  - Hvornår kan dette ske
    - Det kan godt vœre at det er smart hvis man kunne rgne ud hvor mange grader man peger i den rigtige retning ud fra indre produktet.
- 3.
  - Hvad skal vi gøre her
    - Der er jo stadig en valid path
    - Men der sker nok ikke noget ved at sige vi bare lukker den
    - Fordi på det tidspunkt det sker, så er der nok andre heads der er blevet generetet. Så der sker ikke noget ved at lukke denne.


### Omdan indre produkt til grader

- Der må jo vœre en sammenhœng, fordi ved indreprod = 0, så er de vinkel ret på hinanden.

- Hvordan vil vi bruge dette
  - Hvis der er en mere direkte måde at omdanne det til grader, så kunne man jo bruge dette til at vœlge dir ud fra.
  - Det kan også give lidt bedre logning
  - bedre forståelse af havd der forgår
  - 
- Hvordan er setup
  - Vi har en target vector, der går fra head til target
  - Så har vi en cell vektor fra, head til alle valid dirs
  - Ud fra det finder vi så indre produktet mellem disse, og finder hvor stor en procent stas de udgør af det samlede indre produkt.

- Hvad mener jeg overhovedet
  - Vinklen mellem cellevektoren og targetvektoren
  - 

## Håndter flere heads

- Vi har jo et array af heads
- Men det er måske bedre bare at spawn en ny thread som også bare laver en path

### thread per head

- Man kan sige, hver gang vi laver et nyt target, så kan vi spawn 1 ny thread
- Så skal vi have fjernet logikken med arrayet af heads
- Ellers skal vi lave heads arrayet til shared memory. Men det kan jeg ikke rigtig se vœrdien i

- selve maze'en skal vœre shared.
  - Hvordan laver vi den shared når heads ikke er.
  - Så skal man skile heads og maze af.

- Hvordan skal Maze struct se ud
  - Spørgsmålet er om target og paths skal vœre i den
  - Bruger vi Target field fra maze. Eller bliver den i forvejen bare givet med fra siden af?
  - paths kan vi godt skilde fra hinanden
  - Men hvis ikke vi har Target og paths, giver det så overhovedet mening at skulle have en bestemt bestemt Maze struct, da den egenligt bare er en mat.Dense og yBound, xBound, allerede ligger i det objekt.
  - Og når det allerede er en pointer, så er den vel sådan set allerede shared et eller andet sted.

- Der er måske problemer i step, fordi der brugte jeg en pointer til head, men det er det ikke nu.

- Head bliver nødt til at vœre en pointer
- Hvordan sikre vi access til shared memory
  - Skal nok have noget semaphore.
  - Maze vil give raceConditions, semaphore til at bestemme access.
    - Sharedmemory bliver fixed ved Channels 

- For at vi kan kalde det concurrently, så skal vi have samlet det i en metode
- initial head variabel {0,0}
- Indhold i concurrent metode
  - head variabel
  - genTargetZone
  - StepVectorProduct()
  - printMaze
  - Genere ny thread efter x steps
- 1 enkelt inital loop er nok

- Vil fortsœtte til alle celler er fyldt.
  - Så vi skal have en eller anden form for break.
  - 


- shared memory via channels
  - Hvordan fungerer det
    - Hvordan bliver shared memory locked når man bruger channles.
    - Send og read fra channels er blocking by default
    - Så send blokere så lœnge der ikke bliver lœst fra den channel og read blokere hvis den prøver at lœse fra en channel, men der ikke er noget data i.

    - Hvad menes med blokere
      - Det betyder vel bare at programmet ikke fortsœtter efter en read, hvis ikke der bliver lœst noget.
    - Har ikke noget at gøre med read/write access til shared memory.

  - Hvordan gør man
    - Maze kan leve på channels
    - Så før du lœser fra og skrive til maze, skal du have det fra en channel
    - Dette tror jeg stadig kan lede til fejl, hvor en goroutine lœser fra en gammel
      - Det tror jeg ikke. Fordi det kan godt vœre denne go rutine er gammel, men den får stadig den seneste aktuelle maze.
    - Så read fra channel inden du laver read fra maze og write til channel efter du har skrevet til den.
    - Det fungere jo ligesom en read write lock.


- hvordan stopper man en gorutine
  - https://yourbasic.org/golang/stop-goroutine/
- Måden vi har gjort det på, så kører programmet bare rent sekventielt  fordi vi låser alt logikken så lœnge du ikke er den der har maze.
  - Problemet er bare også at man kan ikke lœse chan, give access videre og så œndre i maze. Fordi så passer dit billede ikke med den faktiske verden.
  - Så det giver ikke rigtig mening at bruge channels.


- Sekeventielt løsning
  - Er det bedst at have et heads array 
    - Det bliver vi jo nødt til når vi har flere heads
- Skal det ligge i Maze strukturen eller parsere vi det med til funktionerne som vi har gjort her.
  - Fordele
    - Skal ikke til at lave det om
    - Hvis vi skiller os af med Target også, så kan vi helt droppe Maze strukt. Og har så ikke Maze.Maze
  - Ulemper
    
- plan for sekventiel
  - Fjern maze strukt -> Der da ikke lœngere er nogen justification for at have den
  - Lav et array af heads i createPath


- Efter hver stepsPerRound antal steps, s∑ skal vi generere et nyt head og ligge til heads (med en eller anden ss)
- Hver runde vœlger man et head og et target, tager $stepsPerRound antal steps mod det. genere et nt head. vœlge head og target og state forfra

- Vi skal ikke crash når der kommer en fejl.
  - kan bare return i stedet

- Current problem
  - Det er ikke hele maze'en der bliver fyldt
  - Det er oftes altid inde i midten og så bliver det nogen bare store klumper af "path"
  - Det er bedre hvis det er mere de samme stier der bliver bygget ud 
  - Fordi sådan får man bare en masse små stier uden nogen rigtig dybte
  - Det behøver vel heller ikke rigtig vœre tilfœldigt når vi vœlger hvilken path der skal increase.
  - Man kan godt gå på runde
  - Hvor mnage paths skal vi så have
    - 1 per runde bliver formange
    - Man kunne godt isge at man har 5 paths, som så for lov til at sprede sig, og så tager man de nΩste 5. Og så gøre det x gange.
    

- Hver sub-path, skal kunne lave sin egen path
  - Sådan som det er nu har vi bare en rœkke forgreninger der efter hver runde.
  - Det er vel også ok

  - Når vi laver et nyt head, så skal det vœre ud fra de heads vi har i forvejen. Det skal ikke altid vœre den seneste.
    - Det virker vel sådan set fint nok, fordi vi vœlger et vilkårligt head hver runde.


## Rethink forgreninger

- Vi vil gerne have flere lœngere forgreninger der spreder sig ud til kanterne af maze.

- Man kører på runder, hvor hver head i tager x steps og med en hvis ss adder et nyt head til heads som vil tages med nœste runde

- Target selection
  - Der skal også vœre paths langs kanten
    - Dem derer svœre at nå er for alle y | x = max && alle x | y = max
    - Det kan styres med target selection.
  - Regler for target selection
    - De skal prioritere spredning vœk fra midten

  - Man kan give dem områder, så et target holder sig inden for et given område
  - Gør så de ikke kryder på tvœrs hele tiden og dermed samler sig i midten.
  - Ved at dele op på midten skaber det bare 2 nye centrumer, hvor paths vil samle sig.
  - Det kan måske betale sig at lav et framework til at sœtte zoner op. Så man kan prøve nogen forskellige patterns
  - Og så skal man måske have nogen targets som gør som normalt til at forbinde zonerne.


- skal have en måde at steppe i gennem koden og se hvad der sker
  - Hvis man kunne se terminal output fra debuggeren, så vil det kunne bruges.


- Man vil sætte targets væk fra de andre paths og fordi vi har arbejder i en matrice så kan man lave nogen operationer der giver et billede af hvordan matricen ser ud og så kan man lave det baseret op det
- Meget ligesom vi har set i ml
- Når man vælger et target, så løber man matricen i gennem og laver en gradient, sådan at hvis du er ved siden af en path, trækkes x fra ( afhængig af hvor tæt du er på en path) (kan man sikkert også lave et filter på)
Når man så har den, så kan man scanne den med et filter for at finde de steder længst væk fra en path, som man så kan sætte som target

- Er det projectioner vi skal bruge til at "omdanne" matricen
- spørgsmålet er om det kan betale sig at lave alle de ekstra operationer
  - Vi bliver jo nødt til at lave gradieringen
  - Men behøver vi at køre filteret over
    - Nok ikke. Fordi den gradiering er jo allerede en form for filter
    - Når man så skal finde et target område, så går man jo bare ud af hvor cellerne bliver større og større.

## Kernels

- Hvilke kernels skal vi bruge
- Vi har avg kernes og max-kernels

- Det er måske fint at bruge avg kernels og når kernelen ikke er så stor så kan det give en fin ide
om der er en path eller ej.

- Ellers skulle man bruge en kernel der ser linjer
- Det er ikke effektivt fordi det er ikke nødvendigvis linjer

- avg-kernel er den mest indlysende at bruge
  - Så passer det også nemt med der er 0 i en path.
  - Og så vœgter man de entries tœttes på 1 højest.

- Hvordan bruger man den til at vœlge et target
  - Man kan bruge det til at lave en target vector
    - Man skal stadig have nogen koordinator
    - Og man kan ikke rigtig omdanne tilbage igen
    - Det ville i hvertfald ikke give mening at give mening nu.
  - Man kan lave en estimering baseret på hvor meget matricen bliver scalet ned. og så gang med den faktor 



- Hvordan vil man lave en kernel
  - Vi vil lave en avg kernel
  - Hvis det bare er en avg kernel, så summere man bare alle entries og dividere med antallet
  - Så man behøver ikke definere en decideret kernel matrice som man kørere over
  - men bare tage den lidt i etaper
  - Det er vel bare et forloop, hvor man summere man dem alle sammen, avg, og så smid det ind i en result matrice

  - Hvor meget bliver matricen skåret ned.
    -
  - Du x og y bliver reduceret med hvor stor kernel matricen er

