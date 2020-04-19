package path

import . "ch6/test-distance/point"

type Path []Point

func (path Path) Distance() int {
	sum := 0

	for i := range path {
		if i > 0 {
			sum += path[i].Distance()
		}
	}

	return sum
}
