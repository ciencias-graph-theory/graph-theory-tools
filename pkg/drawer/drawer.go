package drawer

import (
	"math"

	"github.com/ciencias-graph-theory/graph-theory-tools/internal/fileutils/svg"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

// vertex struct to represent vertex drawings.
type vertex struct {
	cx    float64
	cy    float64
	color string
}

// newVertex initializes a vertex drawing. Receives the coordinates of the center
// of the vertex.
func newVertex(cx float64, cy float64) *vertex {
	return &vertex{
		cx:    cx,
		cy:    cy,
		color: "white",
	}
}

// edge struct to represent edge drawings.
type edge struct {
	u     int
	v     int
	curve bool
	color string
}

// newEdge initializes an edge drawing. Receives the ends of the edge and a
// boolean that indicates if the edge needs to be curved.
func newEdge(u int, v int, curve bool) *edge {
	return &edge{
		u:     u,
		v:     v,
		curve: curve,
		color: "black",
	}
}

// Implementation of circular graph drawing algorithm. Receives a graph and
// returns the Svg object with the graph's SVG representation.
func Circular(g graph.Graph, width int, height int) *svg.Svg {

	verticesMap := make(map[int]vertex)
	edgesSlice := []edge{}

	w := float64(width)
	h := float64(height)
	s := svg.NewSvg(width+40, height+40)
	d := math.Min(w, h)
	cx := 20 + (w / 2)
	cy := 20 + (h / 2)
	r := d / 2

	order := float64(g.Order())
	vertexRadius := math.Sin(float64(math.Pi)/(2.0*order)) * r
	// Limits vertex size
	if vertexRadius > d*0.05 {
		vertexRadius = d * 0.05
	}
	angleDivision := 360.0 / order
	angle := 0.0

	innerAreaRadius := r - 1.5*vertexRadius
	//s.DrawCircle(cx, cy, innerAreaRadius, "red", 1, "white")
	if matrix, err := g.Matrix(); err != graph.NilAdjacencyMatrix {
		for i, v := range matrix {
			x := cx - (r * math.Sin(math.Pi*2*angle/360))
			y := cy - (r * math.Cos(math.Pi*2*angle/360))
			angle += angleDivision
			verticesMap[i] = *newVertex(x, y)
			for j, n := range v {
				if n == 1 {
					if math.Abs(float64(i-j)) == 1 {
						edgesSlice = append(edgesSlice, *newEdge(i, j, false))
					} else {
						edgesSlice = append(edgesSlice, *newEdge(i, j, true))
					}
				}
			}
		}
	} else if list, err := g.List(); err != graph.NilAdjacencyMatrix {
		for v, neighbours := range list {
			x := cx - (r * math.Sin(math.Pi*2*angle/360))
			y := cy - (r * math.Cos(math.Pi*2*angle/360))
			angle += angleDivision
			verticesMap[v] = *newVertex(x, y)
			for _, u := range neighbours {
				if math.Abs(float64(u-v)) == 1 {
					edgesSlice = append(edgesSlice, *newEdge(u, v, false))
				} else {
					edgesSlice = append(edgesSlice, *newEdge(u, v, true))
				}
			}
		}
	}
	for _, e := range edgesSlice {
		u := verticesMap[e.u]
		v := verticesMap[e.v]
		x1 := u.cx
		y1 := u.cy
		x2 := v.cx
		y2 := v.cy

		a := y1 - y2
		b := x2 - x1
		c := x1*y2 - x2*y1
		dist := math.Abs(a*cx+b*cy+c) / math.Sqrt(a*a+b*b)

		dx := (x1+x2)/2 - cx
		dy := (y1+y2)/2 - cy

		qx := cx + dx*(innerAreaRadius/r)
		qy := cy + dy*(innerAreaRadius/r)

		if dist >= innerAreaRadius && e.curve {
			s.DrawCurve(x1, y1, qx, qy, x2, y2, e.color, 1)
		} else {
			s.DrawLine(x1, y1, x2, y2, e.color, 1)
		}
	}
	for _, info := range verticesMap {
		s.DrawCircle(info.cx, info.cy, vertexRadius, "black", 1, info.color)
	}
	return s
}
