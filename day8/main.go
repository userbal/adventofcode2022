package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
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

func Convert(input string) *[][]Tree {
	data := strings.Split(input, "")
	trees := [][]Tree{}

	fmt.Printf("len(data): %v\n", len(data))

	row := []Tree{}
	for _, v := range data {
		if v == "\n" {
			trees = append(trees, row)
			row = []Tree{}
		} else {
			h, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			row = append(row, Tree{height: h})
		}

	}
	trees = append(trees, row)

	return &trees
}

// func main() {
// 	data, err := os.ReadFile("data.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	trees := Convert(string(data))

// 	checkAll(trees)

// 	visible := 0
// 	for _, row := range *trees {
// 		for _, tree := range row {
// 			if tree.visible {
// 				visible += 1
// 			} else {
// 			}
// 		}
// 	}

// 	fmt.Printf("visible: %v\n", visible)
// }

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	trees := Convert(string(data))

	// trees := Convert(testInput2)

	x, y, score := findBestTree(trees)
	fmt.Printf("x: %v, y: %v \n", x, y)
	fmt.Printf("score: %v\n", score)
	fmt.Println((*trees)[4][1].height)
}

// func main() {
// 	trees := Convert(testInput2)
// 	score := calcViewableTreeScore(trees, 2, 3)
// 	fmt.Printf("score: %v\n", score)
// }
