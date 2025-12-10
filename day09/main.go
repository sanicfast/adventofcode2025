package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x int
	y int
}
type Edge struct {
	a Coord
	b Coord
}

func readInput() ([]Coord, []Edge, []Coord, []Edge) {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	tiles := make([]Coord, 0, len(lines))
	xVals := []int{}
	yVals := []int{}
	for _, line := range lines {
		f := strings.Split(line, ",")
		x, e1 := strconv.Atoi(f[0])
		y, e2 := strconv.Atoi(f[1])
		check(e1)
		check(e2)
		tiles = append(tiles, Coord{x, y})
		xVals = append(xVals, x)
		yVals = append(yVals, y)
	}
	xValsDedup := sortAndDedup(xVals)
	yValsDedup := sortAndDedup(yVals)

	tilesCompressed := make([]Coord, len(tiles))
	for i := range tilesCompressed {
		for j := range xValsDedup {
			if tiles[i].x == xValsDedup[j] {
				tilesCompressed[i].x = j * 2
			}
		}
		for j := range yValsDedup {
			if tiles[i].y == yValsDedup[j] {
				tilesCompressed[i].y = j * 2
			}
		}
	}

	edges := make([]Edge, 0, len(lines))
	for i := range tiles {
		edges = append(edges, Edge{tiles[i], tiles[(i+1)%len(tiles)]})
	}
	edgesCompressed := make([]Edge, 0, len(lines))
	for i := range tiles {
		edgesCompressed = append(edgesCompressed, Edge{tilesCompressed[i], tilesCompressed[(i+1)%len(tiles)]})
	}

	return tiles, edges, tilesCompressed, edgesCompressed
}

func sortAndDedup(slice []int) []int {
	sort.Ints(slice)
	j := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[j] {
			j++
			slice[j] = slice[i] // Move unique element to the next available position
		}
	}
	return slice[:j+1] // Return the slice up to the last unique element
}

func getArea(c1, c2 Coord) int {

	loX, hiX := minMaxInt(c1.x, c2.x)
	loY, hiY := minMaxInt(c1.y, c2.y)

	x := hiX - loX + 1
	y := hiY - loY + 1
	area := x * y
	if area < 0 {
		return -area
	}
	return area
}
func biggestSquare(tiles []Coord) int {
	maxArea := 0
	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := getArea(tiles[i], tiles[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func minMaxInt(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

// check if the point is on the line segment including the ends
// assuming no diagonal edges!
func isPointInEdge(point Coord, e Edge) bool {
	if e.a.x == e.b.x { // horizontal
		if e.a.x != point.x {
			return false
		}
		loY, hiY := minMaxInt(e.a.y, e.b.y)
		if loY <= point.y && point.y <= hiY {
			return true
		}
	} else { // vertical
		if e.a.y != point.y {
			return false
		}
		loX, hiX := minMaxInt(e.a.x, e.b.x)
		if loX <= point.x && point.x <= hiX {
			return true
		}
	}
	return false
}

// check if we're inside by casting a ray to the left (x negative to zero)
// if we're inside any edge, return true
// if the sum of intersections is odd, we're in the polygon, even we're out
func isPointinPolygon(point Coord, edges []Edge) bool {
	intersectionSum := 0
	for _, e := range edges {
		// check if we're to the right of both edge points
		if point.x < e.a.x && point.x < e.b.x {
			continue
		}
		if isPointInEdge(point, e) {
			return true
		}

		if e.a.y == e.b.y { // horizontal
			if point.y == e.a.y {
				continue
			}
		} else { // vertical
			loY, hiY := minMaxInt(e.a.y, e.b.y)

			if loY < point.y && point.y <= hiY {
				intersectionSum += 1
			}
		}
	}
	return intersectionSum%2 == 1
}

func plot(edges []Edge) {
	maxX, maxY := 0, 0
	for _, e := range edges {
		if e.a.x > maxX {
			maxX = e.a.x
		}
		if e.b.y > maxY {
			maxY = e.a.y
		}
	}
	display := make([][]int, maxY+4)
	for i := range display {
		display[i] = make([]int, maxX+4)
	}

	for _, e := range edges {
		if e.a.x == e.b.x {
			lo, hi := minMaxInt(e.a.y, e.b.y)
			for y := lo; y <= hi; y++ {
				display[y][e.a.x] += 1
			}
		} else {
			lo, hi := minMaxInt(e.a.x, e.b.x)
			for x := lo; x <= hi; x++ {
				display[e.a.y][x] += 1
			}
		}
	}

	display2 := make([][]int, maxY+4)
	for y := range maxY + 4 {
		display2[y] = make([]int, maxX+4)
		for x := range maxX + 4 {
			if isPointinPolygon(Coord{x, y}, edges) {
				display2[y][x] = 1
			}
		}
	}

	for i := len(display) - 1; i >= 0; i-- {
		fmt.Println(display[i])
	}

	for i := len(display2) - 1; i >= 0; i-- {
		fmt.Println(display2[i])
	}
	for i := range display {
		for j := range display[0] {
			if display[i][j] == 0 {
				display[i][j] = 255
			} else if display[i][j] == 2 {
				display[i][j] = 150
			} else {
				display[i][j] = 0
			}
			if display2[i][j] == 0 {
				display2[i][j] = 255
			} else {
				display2[i][j] = 0
			}

		}
	}
	makeimage(display, "input.png")
	makeimage(display2, "output.png")

}

func test() {
	testp := Coord{1, 5}
	teste1 := Edge{Coord{0, 5}, Coord{5, 5}}
	teste1a := Edge{Coord{5, 5}, Coord{0, 5}}
	teste2 := Edge{Coord{0, 4}, Coord{5, 4}}
	teste3 := Edge{Coord{0, 0}, Coord{0, 5}}
	teste4 := Edge{Coord{1, 0}, Coord{1, 5}}
	fmt.Println("isPointInEdge")
	fmt.Println(true == isPointInEdge(testp, teste1))
	fmt.Println(true == isPointInEdge(testp, teste1a))
	fmt.Println(false == isPointInEdge(testp, teste2))
	fmt.Println(false == isPointInEdge(testp, teste3))
	fmt.Println(true == isPointInEdge(testp, teste4))
	fmt.Println("isPointinPolygon")
	myPolygon := []Edge{
		{Coord{1, 1}, Coord{1, 8}},
		{Coord{1, 8}, Coord{3, 8}},
		{Coord{3, 8}, Coord{3, 10}},
		{Coord{3, 10}, Coord{5, 10}},
		{Coord{5, 10}, Coord{5, 8}},
		{Coord{5, 8}, Coord{8, 8}},
		{Coord{8, 8}, Coord{8, 5}},
		{Coord{8, 5}, Coord{6, 5}},
		{Coord{6, 5}, Coord{6, 3}},
		{Coord{6, 3}, Coord{8, 3}},
		{Coord{8, 3}, Coord{8, 1}},
		{Coord{8, 1}, Coord{4, 1}},
		{Coord{4, 1}, Coord{4, 6}},
		{Coord{4, 6}, Coord{2, 6}},
		{Coord{2, 6}, Coord{2, 1}},
		{Coord{2, 1}, Coord{1, 1}},
	}
	testp5 := Coord{1, 1}
	testp6 := Coord{1, 2}
	testp7 := Coord{1, 8}
	testp8 := Coord{5, 1}
	testp9 := Coord{4, 8}
	testp10 := Coord{9, 5} //false
	testp11 := Coord{9, 4} //false
	testp12 := Coord{11, 10}
	testp13 := Coord{11, 1}
	plot(myPolygon)

	fmt.Println(isPointinPolygon(testp5, myPolygon))
	fmt.Println(isPointinPolygon(testp6, myPolygon))
	fmt.Println(isPointinPolygon(testp7, myPolygon))
	fmt.Println(isPointinPolygon(testp8, myPolygon))
	fmt.Println(isPointinPolygon(testp9, myPolygon))
	fmt.Println(false == isPointinPolygon(testp10, myPolygon))
	fmt.Println(false == isPointinPolygon(testp11, myPolygon))
	fmt.Println(false == isPointinPolygon(testp12, myPolygon))
	fmt.Println(false == isPointinPolygon(testp13, myPolygon))

}

func isAllRedOrGreen(corner1, corner2 Coord, edges []Edge) bool {
	loX, hiX := minMaxInt(corner1.x, corner2.x)
	loY, hiY := minMaxInt(corner1.y, corner2.y)
	cCheck1 := isPointinPolygon(Coord{corner1.x, corner2.y}, edges)
	cCheck2 := isPointinPolygon(Coord{corner2.x, corner1.y}, edges)
	// check the other corners to see if we don't have to check thoroughly
	if !(cCheck1 && cCheck2) {
		return false
	}
	// quick scan along the verticals
	for y := loY; y <= hiY; y += 1500 {
		if !isPointinPolygon(Coord{loX, y}, edges) {
			return false
		}
		if !isPointinPolygon(Coord{hiX, y}, edges) {
			return false
		}
	}

	// check the whole outline of the box to see if it's inbounds
	for y := loY; y <= hiY; y++ { // left side
		if !isPointinPolygon(Coord{loX, y}, edges) {
			return false
		}
		if !isPointinPolygon(Coord{hiX, y}, edges) { // right side
			return false
		}
	}
	for x := loX; x <= hiX; x++ { // bottom and top
		if !isPointinPolygon(Coord{x, loY}, edges) { //bottom
			return false
		} else if !isPointinPolygon(Coord{x, hiY}, edges) { //top
			return false
		}
	}

	return true

}

func biggestRedGreenSquare(tiles []Coord, edges []Edge) int {
	maxArea := 0
	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := getArea(tiles[i], tiles[j])
			if area > maxArea {
				rectGood := isAllRedOrGreen(tiles[i], tiles[j], edges)
				if rectGood {
					// fmt.Println(tiles[i], tiles[j], maxArea, area)
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func makeimage(pixelData [][]int, output string) {
	// Example list of lists of integers representing pixel values (0-255 for grayscale)
	// pixelData := [][]int{
	// 	{0, 50, 100, 150, 200},
	// 	{20, 70, 120, 170, 220},
	// 	{40, 90, 140, 190, 240},
	// 	{60, 110, 160, 210, 255},
	// }

	// Get dimensions
	height := len(pixelData)
	width := 0
	if height > 0 {
		width = len(pixelData[0])
	}

	// Create a new grayscale image
	img := image.NewGray(image.Rect(0, 0, width, height))

	// Populate the image pixels
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Convert int to uint8 for grayscale value
			grayValue := uint8(pixelData[y][x])
			img.SetGray(x, y, color.Gray{Y: grayValue})
		}
	}

	// Save the image to a file
	outputFile, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}
}

func biggestRedGreenSquareCompress(tiles []Coord, tilesCompressed []Coord, edges []Edge) int {
	maxArea := 0

	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := getArea(tiles[i], tiles[j])
			if area > maxArea {
				rectGood := isAllRedOrGreen(tilesCompressed[i], tilesCompressed[j], edges)
				if rectGood {
					// fmt.Println(tiles[i], tiles[j], maxArea, area)
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func main() {
	startTime := time.Now()
	tiles, edges, tilesCompressed, edgesCompressed := readInput()

	p1Time := time.Now()
	part1 := biggestSquare(tiles)
	fmt.Println("Day 9, Part 1:", part1, time.Since(p1Time))
	// test()
	// plot(edgesCompressed)

	_ = edges
	// p2Time := time.Now()
	// part2 := biggestRedGreenSquare(tiles, edges)
	// fmt.Println("Day 9, Part 2:", part2, time.Since(p2Time))
	// 2s

	p2_2Time := time.Now()
	part2_2 := biggestRedGreenSquareCompress(tiles, tilesCompressed, edgesCompressed)
	fmt.Println("Day 9, Part 2:", part2_2, time.Since(p2_2Time))
	// 650ms

	fmt.Println(time.Since(startTime))

}
