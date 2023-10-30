package maze

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

// Hvorfor skal alle disse vœre public
type TargetCoordinate struct {
	X, Y int `json:"x,omitempty"`
}

type PathCoordinate struct {
	X, Y int
}

type Maze struct {
	Maze *mat.Dense `json:"maze_track,omitempty"`

	Target TargetCoordinate `json:"target,omitempty"`

	Paths []PathCoordinate `json:"paths,omitempty"`

	YBound, XBound int
}

type TargetVektor struct {
	// BaseX og BaseY skal nok vœre optional. Fordi de afhœnger jo af hvilket head man har

	// Vi skal også lige overveje hvad der skal ske når vi opdatere en pathcoordinate, fordi den skal jo følge med. Så det skal vœre et objekt for hver
	BaseX, BaseY, HeadX, HeadY int
}

type StepDirection int

const (
	Left  StepDirection = 0
	Right StepDirection = 1
	Up    StepDirection = 2
	Down  StepDirection = 3
)

func (path PathCoordinate) String() string {
	return fmt.Sprintf("X: %d, Y: %d", path.X, path.Y)
}

func NewMaze(r, c int) Maze {
	data := make([]float64, r*c)
	for i := 0; i < r*c; i++ {
		data[i] = 1
	}

	// Generate 2x2 target zone
	targetX, targetY := TargetZone(r, c)

	mz := &Maze{
		Maze:   mat.NewDense(r, c, data),
		Target: TargetCoordinate{targetX, targetY},
		Paths:  []PathCoordinate{},
		YBound: r,
		XBound: c,
	}

	data2 := make([]float64, r*c)

	copy(data2, data)

	ripple(mz.Maze, mz.Target)

	maze2 := mat.NewDense(r, c, data2)

	ripple(maze2, mz.Target)

	mz2 := Maze{
		Maze:   maze2,
		Target: mz.Target,
		Paths:  append(mz.Paths, PathCoordinate{0, 0}),
		YBound: mz.YBound,
		XBound: mz.XBound,
	}

	createPath(mz)
	println("Target: (", targetX, targetY, ")")

	// dum ide, bare lav et nyt objekt
	PrintMaze(mz2)
	println(&maze2)
	println(&mz.Maze)
	return *mz
}

func ripple(maze *mat.Dense, target TargetCoordinate) {
	yBound, xBound := maze.Dims()
	for x := 0; x < xBound; x++ {
		for y := 0; y < yBound; y++ {
			maze.Set(y, x, Norm(target.X-x, target.Y-y))
		}
	}
}

/*
Checks if the step direction is valid from given path coordinate

Kan fejle ved dir == Down og head.Y = 0
*/

func isDirValid(dir StepDirection, head PathCoordinate,
	xAxisBound, yAxisBound int,
) bool {
	return (dir == Left && head.X != 0) ||
		(dir == Right && head.X != xAxisBound-1) ||
		(dir == Up && head.Y != yAxisBound-1) ||
		(dir == Down && head.Y != 0)
}

func createPath(mz *Maze) {
	// Init first position
	mz.Paths = append(mz.Paths, PathCoordinate{0, 0})

	// mzTrack := mz.Maze
	mz.Maze.Set(0, 0, float64(0))

	// direction := [4] StepDirection {Left, Right, Up, Down}

	println("left:", Left)
	println("Right", Right)
	println("Up:", Up)
	println("Down:", Down)
	fmt.Println("dims", mz.XBound, mz.YBound)
	// Tager 10 skridt
	for i := 0; i < 20; i++ {

		println("******************", i)

		step(mz)

		PrintMaze(*mz)

	}
}

/*
Tag et step for en PathCoordinate

Hvad skal den gøre

  - Vi skal tage et step

  - Det skal vœre et valid step

  - Ud af de valid steps man kan tage vœlger man dem med en ss der går mod target zone.

  - Det betyder vi skal have have alle valid steps, deres totale sum, for hver dir skal vi så tillœkke en sandsynlig

  - head skal vœre en pointer, fordi så kan vi opdatere head direkte
*/
func step(maze *Maze) {
	// Valid dirs
	weights := make([]float64, 4)
	probs := make([]float64, 4)

	// Det er noget pointer shit.
	// Er det en pointer når det er et array ->
	head := &maze.Paths[0]

	var distanceSum float64

	if isDirValid(Left, *head, maze.XBound, maze.YBound) {
		weights[Left] = maze.Maze.At(head.Y, head.X-1)
		distanceSum += maze.Maze.At(head.Y, head.X-1)
	}
	if isDirValid(Right, *head, maze.XBound, maze.YBound) {
		weights[Right] = maze.Maze.At(head.Y, head.X+1)
		distanceSum += maze.Maze.At(head.Y, head.X+1)
	}
	if isDirValid(Up, *head, maze.XBound, maze.YBound) {
		weights[Up] = maze.Maze.At(head.Y+1, head.X)
		distanceSum += maze.Maze.At(head.Y+1, head.X)
	}
	if isDirValid(Down, *head, maze.XBound, maze.YBound) {
		weights[Down] = maze.Maze.At(head.Y-1, head.X)
		distanceSum += maze.Maze.At(head.Y-1, head.X)
	}

	for i := 0; i < 4; i++ {
		// Det her vil fejle fordi vi ikke har en dir ikke har en PathCoordinate.

		//Dette er forkert, her er det den der ligger lœngest vœk som har størst ss for at blive valgt.
			probs[i] = weights[i] / distanceSum
	}

	dir := Sample(probs)

	println("probs:")
	println("Left:", Left, "Right:", Right, "Up:", Up, "Down:", Down)
	for _, v := range probs {
		fmt.Print(v, " ")
	}
	println("\ndir:", dir)

	switch dir {
	case Left:
		head.X = head.X - 1
	case Right:
		head.X = head.X + 1
	case Up:
		head.Y = head.Y + 1
	case Down:
		head.Y = head.Y - 1
	}

	maze.Maze.Set(head.Y, head.X, 0)
}

func Sample(probs []float64) StepDirection {
	summedProbs := make([]float64, len(probs))

	// Det her er vel noget pointer shit

	// Når man arbejder med arrays, så er det vel pointers til arrayet i stedet for

	// Undersøg
	summedProbs[0] = probs[0]
	step := rand.Float64()

	// Det her ville man kunne gøre meget smartere i scala
	for i := 1; i < len(probs); i++ {
		summedProbs[i] = summedProbs[i-1] + probs[i]
	}

	for i := 0; i < len(summedProbs); i++ {
		if step < summedProbs[i] {
			return StepDirection(i)
		}
	}
	// Should be an error
	return StepDirection(0)
}
