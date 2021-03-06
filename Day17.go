package main

func findHighestPointAndVelocityCount(min point, max point) (int, point, int) {
	h := 0
	v := point{0, min.y}
	c := 0
	for x := 0; x < max.x; x++ {
		for y := min.y; y < min.y*-1; y++ {
			nh := nextStep(point{0, 0}, point{x, y}, min, max, 0)
			if nh > h {
				h = nh
				v = point{x, y}
			}
			if nh != -1 {
				c += 1
			}
		}
	}
	return h, v, c
}

func nextStep(p point, v point, min point, max point, h int) int {
	np := p.move(v)
	nv := v.decreaseVelocity()

	if np.isHigherThan(h) {
		h = np.y
	}
	if np.isOnTarget(min, max) {
		return h
	}
	if np.isPastTarget(min) {
		return -1
	}

	return nextStep(np, nv, min, max, h)
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

func (p point) isHigherThan(h int) bool {
	return p.y > h
}

func (p point) isOnTarget(min point, max point) bool {
	return p.x >= min.x && p.y >= min.y && p.x <= max.x && p.y <= max.y
}

func (p point) isPastTarget(min point) bool {
	return p.y < min.y
}
