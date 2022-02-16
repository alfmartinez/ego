package terrain

type Grid struct {
	tiles map[struct {
		x int
		y int
	}]Tile
}

func CreateGrid(width int, height int) *Grid {
	tiles := make(map[struct {
		x int
		y int
	}]Tile)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tiles[struct {
				x int
				y int
			}{x, y}] = Tile{}
		}
	}
	return &Grid{tiles}

}
