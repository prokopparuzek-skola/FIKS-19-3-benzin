package main

import "fmt"
import "math"

type Blot struct {
	x int
	y int
	r int
}
type Map struct {
	X      int
	Y      int
	B      int
	blots  []Blot
	lblots map[int]bool
	ublots map[int]bool
	rblots map[int]bool
	dblots map[int]bool
}

func findRoute(in Map) bool {
	var neighs [][]int
	neighs = make([][]int, in.B)
	for i := range neighs {
		neighs[i] = make([]int, 0)
	}
	for i, b := range in.blots {
		for j := i + 1; j < len(in.blots); j++ {
			nb := in.blots[j]
			od1 := math.Abs(float64(b.x) - float64(nb.x))
			od2 := math.Abs(float64(b.y) - float64(nb.y))
			prep := math.Sqrt(math.Pow(od1, 2) + math.Pow(od2, 2))
			if prep <= float64(b.r+nb.r) {
				neighs[i] = append(neighs[i], j)
				neighs[j] = append(neighs[j], i)
			}
		}
	}
	return true
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		var X, Y, B int
		var can bool = true
		var in Map
		fmt.Scanf("%d%d%d", &X, &Y, &B)
		in.blots = make([]Blot, B)
		in.lblots = make(map[int]bool)
		in.ublots = make(map[int]bool)
		in.rblots = make(map[int]bool)
		in.dblots = make(map[int]bool)
		in.X = X
		in.Y = Y
		in.B = B
		for j := 0; j < B; j++ {
			fmt.Scanf("%d%d%d", &in.blots[j].x, &in.blots[j].y, &in.blots[j].r)
			left := (in.blots[j].x - in.blots[j].r) <= 0
			right := (in.blots[j].x + in.blots[j].r) >= in.X
			up := (in.blots[j].y + in.blots[j].r) >= in.Y
			down := (in.blots[j].y - in.blots[j].r) <= 0
			if (left && down) || (right && up) {
				fmt.Println("Nelze projet")
				can = false
			} else if left {
				in.lblots[j] = true
			} else if right {
				in.rblots[j] = true
			} else if up {
				in.ublots[j] = true
			} else if down {
				in.dblots[j] = true
			}

		}
		if can {
			findRoute(in)
		} else {
			continue
		}
	}
}
