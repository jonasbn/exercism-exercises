// REF: https://exercism.io/my/solutions/273bceb127a14252bbcfff9845b8eaa9

package space

import "math"

type Planet string

const seconds_in_earth_year = 60 * 60 * 24 * 365.25

var planets = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// rounding technique from: Algorithms to Go
// REF: https://yourbasic.org/golang/round-float-2-decimal-places/
func Age(age float64, planet Planet) float64 {
	return math.Round(((age/seconds_in_earth_year)/planets[planet])*100) / 100
}
