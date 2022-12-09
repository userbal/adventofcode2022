package main

const (
	Up    string = "up"
	Down         = "down"
	Left         = "left"
	Right        = "right"
)

const testInput2 = `30373
25512
65332
33549
35390`

func calcViewableTreesUp(trees *[][]Tree, x, y, height int) int {
	// look for next tree, if out of range, or not visible return 0
	if y-1 < 0 {
		return 0
	}
	nextTree := (*trees)[y-1][x]
	if nextTree.height >= height {
		return 1
	}
	return 1 + calcViewableTreesUp(trees, x, y-1, height)
}

func calcViewableTreesDown(trees *[][]Tree, x, y, height int) int {
	// look for next tree, if out of range, or not visible return 0

	if y+1 > len(*trees)-1 {
		return 0
	}

	nextTree := (*trees)[y+1][x]
	if nextTree.height >= height {
		return 1
	}
	return 1 + calcViewableTreesDown(trees, x, y+1, height)
}

func calcViewableTreesLeft(trees *[][]Tree, x, y, height int) int {
	// look for next tree, if out of range, or not visible return 0
	if x-1 < 0 {
		return 0
	}
	nextTree := (*trees)[y][x-1]
	if nextTree.height >= height {
		return 1
	}
	return 1 + calcViewableTreesLeft(trees, x-1, y, height)
}

func calcViewableTreesRight(trees *[][]Tree, x, y, height int) int {
	// look for next tree, if out of range, or not visible return 0
	if x+1 > len(*trees)-1 {
		return 0
	}
	nextTree := (*trees)[y][x+1]
	if nextTree.height >= height {
		return 1
	}
	return 1 + calcViewableTreesRight(trees, x+1, y, height)
}

func calcViewableTreeScore(trees *[][]Tree, x, y int) int {
	up := calcViewableTreesUp(trees, x, y, (*trees)[y][x].height)
	down := calcViewableTreesDown(trees, x, y, (*trees)[y][x].height)
	left := calcViewableTreesLeft(trees, x, y, (*trees)[y][x].height)
	right := calcViewableTreesRight(trees, x, y, (*trees)[y][x].height)
	return up * down * left * right
}

func findBestTree(trees *[][]Tree) (int, int, int) {
	highScore := 0
	bestX := 0
	bestY := 0
	for y, row := range *trees {
		for x := range row {
			score := calcViewableTreeScore(trees, x, y)
			if score > highScore {
				highScore = score
				bestX = x
				bestY = y
			}
		}
	}
	return bestX, bestY, highScore
}
