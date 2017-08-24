package asdfg

func Flood(terrain *[][]rune, x, y int) {
	if y < 0 || y >= len((*terrain)) || x < 0 || x >= len((*terrain)[y]) {
		return
	}
	if (*terrain)[y][x] == '.' {
		(*terrain)[y][x] = 'X'
		Flood(terrain, x+1, y)
		Flood(terrain, x-1, y)
		Flood(terrain, x, y-1)
		Flood(terrain, x, y+1)
	}
}
