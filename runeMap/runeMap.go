package runeMap

type Pos struct {
	I, J int
}

func Legal(m, n, i, j int) bool {
	if i < 0 || i >= m {
		return false
	}
	if j < 0 || j >= n {
		return false
	}
	return true
}

func NeighborsWithMNF(now Pos, maxI, maxJ int, f func(neighborI, neighborJ int) bool) []Pos {
	i := now.I
	j := now.J
	mapHeight := maxI
	mapLength := maxJ
	res := make([]Pos, 0)
	if Legal(mapHeight, mapLength, i+1, j) && f(i+1, j) {
		res = append(res, Pos{
			I: i + 1,
			J: j,
		})
	}
	if Legal(mapHeight, mapLength, i-1, j) && f(i-1, j) {
		res = append(res, Pos{
			I: i - 1,
			J: j,
		})
	}
	if Legal(mapHeight, mapLength, i, j+1) && f(i, j+1) {
		res = append(res, Pos{
			I: i,
			J: j + 1,
		})
	}
	if Legal(mapHeight, mapLength, i, j-1) && f(i, j-1) {
		res = append(res, Pos{
			I: i,
			J: j - 1,
		})
	}
	return res
}
func NeighborsF(now Pos, mP *[][]rune, f func(neighborI, neighborJ int) bool) []Pos {
	i := now.I
	j := now.J
	m := *mP
	mapHeight := len(m)
	mapLength := len(m[0])
	res := make([]Pos, 0)
	if Legal(mapHeight, mapLength, i+1, j) && f(i+1, j) {
		res = append(res, Pos{
			I: i + 1,
			J: j,
		})
	}
	if Legal(mapHeight, mapLength, i-1, j) && f(i-1, j) {
		res = append(res, Pos{
			I: i - 1,
			J: j,
		})
	}
	if Legal(mapHeight, mapLength, i, j+1) && f(i, j+1) {
		res = append(res, Pos{
			I: i,
			J: j + 1,
		})
	}
	if Legal(mapHeight, mapLength, i, j-1) && f(i, j-1) {
		res = append(res, Pos{
			I: i,
			J: j - 1,
		})
	}
	return res
}

func IllegalNeighborsF(now Pos, mP *[][]rune, f func(neighborI, neighborJ int) bool) []Pos {
	i := now.I
	j := now.J
	m := *mP
	mapHeight := len(m)
	mapLength := len(m[0])
	res := make([]Pos, 0)
	if !Legal(mapHeight, mapLength, i+1, j) || f(i+1, j) {
		res = append(res, Pos{
			I: i + 1,
			J: j,
		})
	}
	if !Legal(mapHeight, mapLength, i-1, j) || f(i-1, j) {
		res = append(res, Pos{
			I: i - 1,
			J: j,
		})
	}
	if !Legal(mapHeight, mapLength, i, j+1) || f(i, j+1) {
		res = append(res, Pos{
			I: i,
			J: j + 1,
		})
	}
	if !Legal(mapHeight, mapLength, i, j-1) || f(i, j-1) {
		res = append(res, Pos{
			I: i,
			J: j - 1,
		})
	}
	return res
}
