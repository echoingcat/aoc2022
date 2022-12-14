package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n")
	sensors, beacons, dmap := make([]Sensor, 0), make([]Point, 0), make(map[Point]Point)
	grid, gridB := make([]bool, 10000000), make([]bool, 10000000)

	for _, s := range input {
		var sx, sy, bx, by int
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		gridB[bx+5000000] = (by == 2000000)
		dist := manDist(sx, sy, bx, by)
		fill(sx, sy, dist*2+1, grid)
		top, right, bottom, left := Point{x: sx, y: sy - dist}, Point{x: sx + dist, y: sy}, Point{x: sx, y: sy + dist}, Point{x: sx - dist, y: sy}
		currentSensor := Sensor{top: top, right: right, bottom: bottom, left: left, self: Point{x: sx, y: sy}, dist: dist}
		currentBeacon := Point{x: bx, y: by}
		sensors, beacons = append(sensors, currentSensor), append(beacons, currentBeacon)
	}

	fmt.Println("Part 1:", count(grid, gridB))
	fmt.Println("Time elapsed:", time.Since(start))
	fmt.Println("\nPart 2 Runtime is a few minutes...")

	for _, s := range sensors {
		addAllDiagElements(dmap, s)
	}

	for i := range sensors {
		checkDiagonals(dmap, sensors[i], sensors)
	}

	fmt.Println("Time elapsed:", time.Since(start))
}

func manDist(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func addAllDiagElements(dmap map[Point]Point, s Sensor) map[Point]Point {
	dmap = addSingleDiagElements(s.bottom.x, s.right.x, s.bottom.y, s.bottom.y, dmap, false)
	dmap = addSingleDiagElements(s.top.x, s.right.x, s.bottom.y, s.top.y, dmap, true)
	dmap = addSingleDiagElements(s.left.x, s.top.x, s.bottom.y, s.left.y, dmap, false)
	dmap = addSingleDiagElements(s.left.x, s.bottom.x, s.bottom.y, s.left.y, dmap, true)
	return dmap
}

func addSingleDiagElements(x, y, z1, z2 int, dmap map[Point]Point, sign bool) map[Point]Point {
	count := 0
	if sign {
		for i := x; i <= y; i++ {
			if i >= 0 && i <= 4000000 && z1-count >= 0 && z2+count <= 4000000 {
				p := Point{x: i, y: z2 + count}
				dmap[p] = p
			}
			count++
		}
	} else {
		for i := x; i <= y; i++ {
			if i >= 0 && i <= 4000000 && z1-count >= 0 && z2-count <= 4000000 {
				p := Point{x: i, y: z2 - count}
				dmap[p] = p
			}
			count++
		}
	}
	return dmap
}

func checkDiagonals(dmap map[Point]Point, s Sensor, sc []Sensor) {
	checkDiagonal(s.bottom.x, s.right.x, s.bottom.y, 1, dmap, sc, false)
	checkDiagonal(s.top.x, s.right.x, s.top.y, 1, dmap, sc, true)
	checkDiagonal(s.left.x, s.top.x, s.left.y, -1, dmap, sc, false)
	checkDiagonal(s.left.x, s.bottom.x, s.left.y, -1, dmap, sc, true)
}

func checkDiagonal(x, y, z, o int, dmap map[Point]Point, sc []Sensor, sign bool) {
	count := 0
	for i := x; i <= y; i++ {
		if _, ok := dmap[Point{x: i + 2*o, y: z + count}]; ok {
			if _, ok := dmap[Point{x: i + o, y: z + count + 1}]; ok {
				if _, ok := dmap[Point{x: i + o, y: z + count - 1}]; ok {
					if _, ok := dmap[Point{x: i + o, y: z + count}]; ok {
					} else {
						if !isInside(Point{x: i + o, y: z + count}, sc) {
							fmt.Println("Part 2:", int64(((i+o)*4000000 + z + count)))
							os.Exit(0)
						}
					}
				}
			}
		}
		if sign {
			count++
		} else {
			count--
		}

	}
}

func isInside(p Point, sc []Sensor) bool {
	for _, s := range sc {
		if manDist(p.x, p.y, s.self.x, s.self.y) <= s.dist {
			return true
		}
	}
	return false
}

func fill(x int, y int, c int, grid []bool) []bool {
	offset := 2000000 - y
	from := (x - (c / 2)) + int(math.Abs(float64(offset)))
	to := (x + (c / 2)) - int(math.Abs(float64(offset)))

	if checkIfOnLine(y, (c / 2)) {
		for i := from; i <= to; i++ {
			grid[i+5000000] = true
		}
	}
	return grid
}

func checkIfOnLine(y int, d int) bool {
	if y-d <= 2000000 && y+d >= 2000000 {
		return true
	}
	return false
}

func count(grid []bool, gridB []bool) int {
	res, res2 := 0, 0

	for _, s := range grid {
		if s == true {
			res++
		}
	}

	for _, s := range gridB {
		if s == true {
			res2++
		}
	}
	return res - res2
}

type Sensor struct {
	top, right, bottom, left, self Point
	dist                           int
}

type Point struct {
	x, y int
}
