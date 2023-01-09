package svg

import "fmt"

const svgHeader = "<svg xmlns=\"http://www.w3.org/2000/svg\" %s>%s"

type Svg struct {
	content string
}

func NewSvg(width int, height int) *Svg {
	dimensions := fmt.Sprintf("width=\"%d\" height=\"%d\"", width, height)
	emptySvg := fmt.Sprintf(svgHeader, dimensions, "\n %s \n</svg>")
	return &Svg{
		content: emptySvg,
	}
}

func (s *Svg) Content() string {
	return s.content
}

func (s *Svg) DrawCircle(cx float64, cy float64, r float64, stroke string,
	strokeWidth int, fill string) {
	circle := fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" stroke=\"%s\" stroke-width=\"%d\" fill=\"%s\"/>",
		cx, cy, r, stroke, strokeWidth, fill)
	s.content = fmt.Sprintf(s.content, circle+"\n%s")
}

func (s *Svg) DrawLine(x1 float64, y1 float64, x2 float64, y2 float64,
	stroke string, strokeWidth int) {
	line := fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" stroke=\"%s\" stroke-width=\"%d\"/>",
		x1, y1, x2, y2, stroke, strokeWidth)
	s.content = fmt.Sprintf(s.content, line+"\n%s")

}
