package main

func checkVisible(trees *[][]Tree) {
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
