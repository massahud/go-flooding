// Non recursive flooding, using a queue
package flooding_nr

type position struct {
	x int
	y int
}

func Flood(terrain *[][]rune, x, y int) {
	queue := make([]position, 0)
	queue = append(queue, position{x,y})

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if (*terrain)[pos.y][pos.x] == '.' {
			(*terrain)[pos.y][pos.x] = 'X'
			if validatePos(terrain, pos.x+1, pos.y) {
				queue = append(queue,position{pos.x + 1, pos.y})
			}
			if validatePos(terrain, pos.x-1, pos.y) {
				queue = append(queue, position{pos.x - 1, pos.y})
			}
			if validatePos(terrain, pos.x, pos.y-1) {
				queue = append(queue, position{pos.x, pos.y - 1})
			}
			if validatePos(terrain, pos.x, pos.y+1) {
				queue = append(queue, position{pos.x, pos.y + 1})
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
