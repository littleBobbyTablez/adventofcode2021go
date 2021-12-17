package main

func findHighestPoint(min point, max point) (int, point) {
	h := 0
	v := point{0, min.y}
	for x := 0; x < max.x; x++ {
		for y := min.y; y < 500; y++ {
			nh := moveStep(point{0, 0}, point{x, y}, min, max, 0)
			if nh > h {
				h = nh
				v = point{x, y}
			}
		}
	}
	return h, v
}

func findNumberOfPoints(min point, max point) int {
	c := 0
	var points []point
	for x := 0; x <= max.x; x++ {
		for y := min.y; y <= min.y*-1; y++ {
			nh := moveStep(point{0, 0}, point{x, y}, min, max, 0)
			if nh != -1 {
				c += 1
				points = append(points, point{x, y})
			}
		}
	}
	return c
}

func moveStep(p point, v point, min point, max point, h int) int {
	np := p.move(v)
	nv := v.decreaseVelocity()

	if np.y > h {
		h = np.y
	}

	if np.isOnTarget(min, max) {
		return h
	}
	if np.isPastTarget(min) {
		return -1
	}

	return moveStep(np, nv, min, max, h)
}

func (p point) move(v point) point {
	return point{p.x + v.x, p.y + v.y}
}

func (p point) decreaseVelocity() point {
	switch {
	case p.x < 0:
		return point{p.x + 1, p.y - 1}
	case p.x > 0:
		return point{p.x - 1, p.y - 1}
	default:
		return point{p.x, p.y - 1}
	}
}

func (p point) isOnTarget(min point, max point) bool {
	return p.x >= min.x && p.y >= min.y && p.x <= max.x && p.y <= max.y
}

func (p point) isPastTarget(min point) bool {
	return p.y < min.y
}
