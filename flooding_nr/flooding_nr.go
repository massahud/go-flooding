package flooding_nr

type position struct {
	x int
	y int
}

func Flood(terrain *[][]rune, x, y int) {
	size := 0
	for _, line := range *terrain {
		size += len(line)
	}

	// TODO: use real FIFO
	// using a channel instead of a FIFO forces me to set the buffer
	// to the maximum possible number of items for it not to hang
	c := make(chan position, size)

	c <- position{x, y}

	for len(c) > 0 {
		pos := <-c

		if (*terrain)[pos.y][pos.x] == '.' {
			(*terrain)[pos.y][pos.x] = 'X'
			if validatePos(terrain, pos.x+1, pos.y) {
				c <- position{pos.x + 1, pos.y}
			}
			if validatePos(terrain, pos.x-1, pos.y) {
				c <- position{pos.x - 1, pos.y}
			}
			if validatePos(terrain, pos.x, pos.y-1) {
				c <- position{pos.x, pos.y - 1}
			}
			if validatePos(terrain, pos.x, pos.y+1) {
				c <- position{pos.x, pos.y + 1}
			}

		}
	}
}

func validatePos(terrain *[][]rune, x, y int) bool {
	if y < 0 || y >= len(*terrain) || x < 0 || x >= len((*terrain)[y]) {
		return false
	}
	return (*terrain)[y][x] == '.'
}
