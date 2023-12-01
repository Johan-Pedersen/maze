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
