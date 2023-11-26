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
