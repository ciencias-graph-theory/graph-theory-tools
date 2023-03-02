package svg

import (
	"fmt"
)

// String constant for the header of the SVG file.
const svgHeader = "<svg xmlns=\"http://www.w3.org/2000/svg\" %s>%s"

// A Svg struct represents a SVG file, with all its content described by a
// single string.
type Svg struct {
	content string
}

// NewSvg initializes an empty SVG image.
func NewSvg(width int, height int) *Svg {
	dimensions := fmt.Sprintf("width=\"%d\" height=\"%d\"", width, height)
	emptySvg := fmt.Sprintf(svgHeader, dimensions, "\n %s \n</svg>")
	return &Svg{
		content: emptySvg,
	}
}

// Content returns the current content of the SVG image as a string.
func (s *Svg) Content() string {
	return s.content
}

// DrawCircle adds a circle element to the SVG image.
func (s *Svg) DrawCircle(cx float64, cy float64, r float64, stroke string,
	strokeWidth int, fill string) {
	circle := fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" stroke=\"%s\" stroke-width=\"%d\" fill=\"%s\"/>",
		cx, cy, r, stroke, strokeWidth, fill)
	s.content = fmt.Sprintf(s.content, circle+"\n%s")
}

// DrawLine adds a line element to the SVG image.
func (s *Svg) DrawLine(x1 float64, y1 float64, x2 float64, y2 float64,
	stroke string, strokeWidth int) {
	line := fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" stroke=\"%s\" stroke-width=\"%d\"/>",
		x1, y1, x2, y2, stroke, strokeWidth)
	s.content = fmt.Sprintf(s.content, line+"\n%s")

}

// DrawCurve adds a curve (path) element to the SVG image.
func (s *Svg) DrawCurve(x1 float64, y1 float64, qx float64, qy float64,
	x2 float64, y2 float64, stroke string, strokeWidth int) {
	line := fmt.Sprintf("<path d=\" M %f %f Q %f %f %f %f \" stroke=\"%s\" stroke-width=\"%d\" fill=\"transparent\"/>",
		x1, y1, qx, qy, x2, y2, stroke, strokeWidth)
	s.content = fmt.Sprintf(s.content, line+"\n%s")
}
