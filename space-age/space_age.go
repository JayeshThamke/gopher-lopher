package space

import (
	"fmt"
	"strconv"
)

// Planet -
type Planet string

// PlanetMeta -
type PlanetMeta struct {
	planet      Planet
	description string
	seconds     float32 // seconds in a year of planet
	period      float32
}

func planetMetaBuilder(p Planet) PlanetMeta {
	var pmeta PlanetMeta
	switch p {
	case "Mercury":
		pmeta = PlanetMeta{planet: p, period: 0.2408467}
	case "Venus":
		pmeta = PlanetMeta{planet: p, period: 0.61519726}
	case "Earth":
		pmeta = PlanetMeta{planet: p, period: 1.0}
	case "Mars":
		pmeta = PlanetMeta{planet: p, period: 1.8808158}
	case "Jupiter":
		pmeta = PlanetMeta{planet: p, period: 11.862615}
	case "Saturn":
		pmeta = PlanetMeta{planet: p, period: 29.447498}
	case "Uranus":
		pmeta = PlanetMeta{planet: p, period: 84.016846}
	case "Neptune":
		pmeta = PlanetMeta{planet: p, period: 164.79132}
	default:
		pmeta = PlanetMeta{planet: "UnknownPlanet", period: 0}
	}
	updateSecondsInYear(&pmeta)
	return pmeta
}

func updateSecondsInYear(pmeta *PlanetMeta) {
	const earthYearInSeconds = 31557600
	pmeta.seconds = pmeta.period * earthYearInSeconds
}

// Age returns age on planet in years
func Age(ageInSeconds float64, p Planet) float64 {
	planetMeta := planetMetaBuilder(p)
	age := ageInSeconds / float64(planetMeta.seconds)
	strAge := fmt.Sprintf("%.2f", age)
	twoPrecisionAge, _ := strconv.ParseFloat(strAge, 32)
	return twoPrecisionAge
}
