package Any

import (
	"math/rand"
)

func String() string {
	return ""
}

func Country() string {
	return countries[randomIndex(len(countries))]
}

func Month() string {
	return month[randomIndex(len(month))]
}

func UsState() string {
	return usstates[randomIndex(len(usstates))]
}

func DayOftheWeek() string {
	return daysoftheweek[randomIndex(len(daysoftheweek))]
}

func SolarSystemPlanet() string {
	return solarsystemplanets[randomIndex(len(solarsystemplanets))]
}

func randomIndex(length int) int {
	var randSource = rand.NewSource(int64(length))
	random := rand.New(randSource)
	return random.Intn(length)
}
