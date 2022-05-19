package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	side  uint
}

func (r *Square) Area() uint {
	return r.side * r.side
}

func (r Square) Perimeter() uint {
	return 4 * r.side
}

func (r Square) End() Point {
	out := Point{r.start.x + int(r.side/2), r.start.y + int(r.side/2)}
	return out
}

// package main

// func main() {
// 	//	m := "+1-2 -3 -5"
// 	m := "-a-2"
// 	var a string
// 	fmt.Println(m)
// 	a, e := StringSum(m)
// 	fmt.Println("string", a)
// 	fmt.Println("error: ", e)
// }
