package main

type command struct {
	direction string
	value     int
}

type position struct {
	x int
	y int
}

type aimPosition struct {
	x   int
	y   int
	aim int
}

func moveToTarget(coms []command, pos position) position {
	if len(coms) == 0 {
		return pos
	}
	return moveToTarget(coms[1:], move(coms[0], pos))
}

func moveToTargetWithAim(coms []command, pos aimPosition) aimPosition {
	if len(coms) == 0 {
		return pos
	}
	return moveToTargetWithAim(coms[1:], moveWithAim(coms[0], pos))
}

func move(com command, pos position) position {
	switch com.direction {
	case "forward":
		return position{pos.x + com.value, pos.y}
	case "down":
		return position{pos.x, pos.y + com.value}
	case "up":
		return position{pos.x, pos.y - com.value}
	default:
		return pos
	}
}

func moveWithAim(com command, ap aimPosition) aimPosition {
	switch com.direction {
	case "forward":
		return aimPosition{ap.x + com.value, ap.y + (com.value * ap.aim), ap.aim}
	case "down":
		return aimPosition{ap.x, ap.y, ap.aim + com.value}
	case "up":
		return aimPosition{ap.x, ap.y, ap.aim - com.value}
	default:
		return ap
	}
}
