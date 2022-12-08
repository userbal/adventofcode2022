package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	height  int
	visible bool
}

const testInput = `30373
25512
65332
33549
35390`

// 30373
// 25512
// 65332
// 33549
// 35390

// expected top
// ttttt
// fttff
// tfftf
// fffft
// ffftf

//expected left
// tfftf
// ttfff
// tffff
// tftft
// ttftf

func convert(input string) *[][]tree {
	data := strings.Split(input, "")
	trees := [][]tree{}

	fmt.Printf("len(data): %v\n", len(data))

	row := []tree{}
	for _, v := range data {
		if v == "\n" {
			trees = append(trees, row)
			row = []tree{}
		} else {
			h, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			row = append(row, tree{height: h})
		}

	}
	trees = append(trees, row)

	return &trees
}

func checkAll(trees *[][]tree) {
	xlen := len((*trees)[0])
	ylen := len(*trees)

	// left
	for i := 0; i < ylen; i++ {
		highest := -1
		for j := 0; j < xlen; j++ {
			if (*trees)[i][j].height > highest {
				highest = (*trees)[i][j].height
				(*trees)[i][j].visible = true
			}
		}
	}
	// right
	for i := 0; i < ylen; i++ {
		highest := -1
		for j := xlen - 1; j >= 0; j-- {
			if (*trees)[i][j].height > highest {
				highest = (*trees)[i][j].height
				(*trees)[i][j].visible = true
			}
		}
	}

	// top
	for x := 0; x < xlen; x++ {
		highest := -1
		for y := 0; y < ylen; y++ {
			if (*trees)[y][x].height > highest {
				highest = (*trees)[y][x].height
				(*trees)[y][x].visible = true
			}

		}
	}
	// // bottom
	for x := 0; x < xlen; x++ {
		highest := -1
		for y := ylen - 1; y > 0; y-- {
			if (*trees)[y][x].height > highest {
				highest = (*trees)[y][x].height
				(*trees)[y][x].visible = true
			}

		}
	}
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	trees := convert(string(data))

	checkAll(trees)

	visible := 0
	for _, row := range *trees {
		for _, tree := range row {
			if tree.visible {
				visible += 1
			} else {
			}
		}
	}

	fmt.Printf("visible: %v\n", visible)
}
