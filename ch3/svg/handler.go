package svg

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func getParams(v url.Values) (map[string]int, map[string]float64) {
	params := v

	// set defaults
	xyrange := 30.0 // axis ranges (-xyrange..+xyrange)

	intMap := map[string]int{
		"width":  600,
		"height": 320,
		"cells":  100,
	}

	for name := range intMap {
		if len(params[name]) > 0 {
			num, err := strconv.Atoi(params[name][0])
			if err != nil {
				fmt.Printf("Bad param: %s=%s\n", name, params[name][0])
				fmt.Println(err)
				continue
			}
			intMap[name] = num
		}
	}

	floatMap := map[string]float64{
		"xyrange": xyrange,
		"xyscale": float64(intMap["width"]) / 2 / xyrange,
		"zscale":  float64(intMap["height"]) * 0.4,
		"angle":   math.Pi / 6,
	}

	for name := range floatMap {
		if len(params[name]) > 0 {
			num, err := strconv.Atoi(params[name][0])
			if err != nil {
				fmt.Printf("Bad param: %s=%s\n", name, params[name][0])
				fmt.Println(err)
			}
			floatMap[name] = float64(num)
		}
	}
	return intMap, floatMap
}

// Handler echoes the Path component of the requested URL.
func Handler(w http.ResponseWriter, r *http.Request) {

	intParams, floatParams := getParams(r.URL.Query())

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(
		[]byte(
			svg(
				intParams["width"],
				intParams["height"],
				intParams["cells"],
				floatParams["xyrange"],
				floatParams["xyscale"],
				floatParams["zscale"],
				floatParams["angle"])))
}

func svg(width int, height int, cells int, xyrange float64, xyscale float64, zscale float64, angle float64) string {
	var xmlResponse string

	var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

	xmlResponse += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(
				i+1, j, xyrange, sin30, cos30, xyscale, zscale, cells, width, height)
			if err != nil {
				continue
			}
			bx, by, err := corner(
				i, j, xyrange, sin30, cos30, xyscale, zscale, cells, width, height)
			if err != nil {
				continue
			}
			cx, cy, err := corner(
				i, j+1, xyrange, sin30, cos30, xyscale, zscale, cells, width, height)
			if err != nil {
				continue
			}
			dx, dy, err := corner(
				i+1, j+1, xyrange, sin30, cos30, xyscale, zscale, cells, width, height)
			if err != nil {
				continue
			}

			xmlResponse += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	xmlResponse += fmt.Sprintf("</svg>")

	return xmlResponse
}

func corner(i, j int, xyrange, sin30, cos30, xyscale, zscale float64, cells, width, height int) (float64, float64, error) {
	var err error
	// Find point (x,y) at corner of cell (i,j).
	x := float64(xyrange) * (float64(i)/float64(cells) - 0.5)
	y := float64(xyrange) * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	// Ex 3.1
	z := f(x, y)
	if z == math.NaN() || z == math.Inf(1) || z == math.Inf(-1) {
		err = errors.New("invalid number, skipping polygon")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, err
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Tan(r)
}
