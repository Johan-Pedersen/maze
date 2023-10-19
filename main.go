package main

import (
	"fmt"

	"maze/model/mazeGen"
)

// Burde vœre inline hints om at den er unused -> Det kommer op hvis jeg kører staticcheck ./... selv. Men vil gerne have det inline

func main() {
	fmt.Println("Suck big cock")

	//  fmt.Println(yes3)
	// Den her lorte luder ting, bliver vist uanset, ligemeget om unused variable analyses er på eller ej
	//	fmt.Println(no)
	//	noyesno := noyesno
	// foo(buf)

	//    mazeGen.CreatePath(mazeGen.GenerateMaze(2,3))

	mazeGen.PrintMaze(mazeGen.GenerateMaze(10, 13))
}
