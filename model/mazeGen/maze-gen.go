package mazeGen

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

//Hvorfor skal alle disse vœre public
type TargetCoordinate struct {
	X, Y int
}

type PathCoordinate struct {
	X, Y int
}

type Maze struct {
	MazeTrack *mat.Dense

	Target TargetCoordinate

	Paths []PathCoordinate
}

type TargetVektor struct{
  //BaseX og BaseY skal nok vœre optional. Fordi de afhœnger jo af hvilket head man har
  
  //Vi skal også lige overveje hvad der skal ske når vi opdatere en pathcoordinate, fordi den skal jo følge med. Så det skal vœre et objekt for hver
  BaseX,BaseY, HeadX,HeadY int 
} 

type StepDirection int 

const (
  Left  StepDirection = 0
  Right StepDirection = 1
  Up    StepDirection = 2 
  Down  StepDirection = 3 
)

var yes = 0

func (path PathCoordinate) String() string{
  return fmt.Sprintf("X: %d, Y: %d", path.X, path.Y)
}

func GenerateMaze(r,c int) Maze {
  data := make([]float64, r*c)
  for i := 0; i<r*c; i++ {
    data[i] = 1
  }
  mz := Maze{ 
		MazeTrack:   mat.NewDense(r, c, data),
		Target: TargetCoordinate{},
		Paths:  []PathCoordinate{},
	}



  return createPath(mz)
}

func createPath(mz Maze) Maze{

  //Init first position
  mz.Paths = append(mz.Paths, PathCoordinate{0,0})

  mzTrack := *mz.MazeTrack
  mzTrack.Set(0,0, float64(0))

  //direction := [4] StepDirection {Left, Right, Up, Down}

  println("left:", Left)
  println("Right", Right)
  println("Up:", Up)
  println("Down:", Down)
  yAxisBound, xAxisBound := mz.MazeTrack.Dims()
  println("yaxis bound:", yAxisBound, "xaxixbound", xAxisBound)
  r,c := mz.MazeTrack.Dims()
  fmt.Println("dims", r,c)
  // Tager 10 skridt 
  for i:=0; i<10; i++{

    println("******************",i)

    dir := StepDirection(rand.Intn(4))
    
    //Denne kunne man godt lave til en pointer. Så man kunne skrive direkte til memory location
    head := mz.Paths[0]

    println(dir)
    fmt.Println("initial head", head)


    if dir == Left && head.X == 0{
        continue
    }
    if dir == Right && head.X == xAxisBound-1{
        continue
    }
    if dir == Up && head.Y == yAxisBound-1{
        continue
    }
    if dir == Down && head.Y == 0{
        continue
    }

    switch dir{
    case Left:
        head.X = head.X -1
    case Right:
        head.X = head.X +1
    case Up:
        head.Y = head.Y +1
    case Down:
        head.Y = head.Y -1
    }

    mz.MazeTrack.Set(head.Y, head.X, 0)

    fmt.Println("head", head)

    mz.Paths[0] = head

  }

  fmt.Println(mz.MazeTrack)

  return mz 
}
