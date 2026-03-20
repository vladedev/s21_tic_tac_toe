package datasource

type GameDS struct {
	ID     string
	Matrix [3][3]int
	Turn   int
	Winner int
}
