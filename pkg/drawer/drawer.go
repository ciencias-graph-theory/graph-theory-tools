package drawer

import (
	"math"

	"github.com/ciencias-graph-theory/graph-theory-tools/internal/fileutils/svg"
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/fileutils/writer"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

type vertex struct {
	cx    float64
	cy    float64
	color string
}

type edge struct {
	u int
	v int
}

func Circular(g graph.Graph, width int, height int) {

	verticesMap := make(map[int]vertex)
	edgesSlice := []edge{}

	w := float64(width)
	h := float64(height)
	s := svg.NewSvg(width+40, height+40)
	d := math.Min(w, h)

	cx := 10 + (w / 2)
	cy := 10 + (h / 2)
	r := d / 2

	//angleDivision := 360.0 / float64(g.Order()) // fix graph order function

	l, _ := g.List()
	angleDivision := 360.0 / float64(len(l))
	angle := 0.0

	if list, err := g.List(); err != graph.NilAdjacencyList {
		for v, neighbours := range list {
			x := cx - (r * math.Sin(math.Pi*2*angle/360))
			y := cy - (r * math.Cos(math.Pi*2*angle/360))
			angle += angleDivision
			verticesMap[v] = vertex{x, y, "white"}
			for _, u := range neighbours {
				edgesSlice = append(edgesSlice, edge{v, u})
			}
		}
		for _, e := range edgesSlice {
			u := verticesMap[e.u]
			v := verticesMap[e.v]
			x1 := u.cx
			y1 := u.cy
			x2 := v.cx
			y2 := v.cy
			s.DrawLine(x1, y1, x2, y2, "black", 1)
		}
		for _, info := range verticesMap {
			s.DrawCircle(info.cx, info.cy, 5, "black", 1, info.color)
		}

	}

	data := []byte(s.Content())
	writer.Write("test.svg", data)
}
