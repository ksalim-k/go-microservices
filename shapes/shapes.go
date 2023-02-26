package shapes

func RectangleProps(length, breadth float64) (float64, float64) {
	area := length * breadth
	perimeter := 2 * (length + breadth)

	return area, perimeter
}
