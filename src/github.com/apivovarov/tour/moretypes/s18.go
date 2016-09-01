package moretypes

func Pic(dx, dy int) [][]uint8 {

	p := make([][]uint8, dy)
	const wh = 255

	for i := 0; i < dy; i++ {
		r := make([]uint8, dx)
		p[i] = r

		c1 := float32(i)
		step := (wh - c1) / float32(dx)
		for j := 0; j < dx; j++ {
			c := int(float32(c1) + step*float32(j))
			if c > 255 {
				c = 255
			}
			r[j] = uint8(c)
		}
	}

	return p
}
