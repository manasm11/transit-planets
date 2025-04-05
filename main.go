package main

type Zodiac int

const (
	_ = iota
	Ari
	Tau
	Gem
	Can
	Leo
	Vir
	Lib
	Sco
	Sag
	Cap
	Aqu
	Pis
)

func (z Zodiac) String() string {
	return [...]string{"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo",
		"Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces"}[z]
}

type PlanetName string

const (
	As PlanetName = "Ascendent"
	Su PlanetName = "Sun"
	Mo PlanetName = "Moon"
	Me PlanetName = "Mercury"
	Ve PlanetName = "Venus"
	Ma PlanetName = "Mars"
	Ju PlanetName = "Jupiter"
	Ra PlanetName = "Rahu"
	Ke PlanetName = "Ketu"
	Sa PlanetName = "Saturn"
)

type Planet struct {
	name    PlanetName
	houseNo int
	zodiac  Zodiac
}

func (p Planet) IsValid() bool {
	switch p.name {
	case As, Su, Mo, Me, Ve, Ma, Ju, Ra, Ke, Sa:
	default:
		return false
	}

	switch p.zodiac {
	case Ari, Tau, Gem, Can, Leo, Vir, Lib, Sco, Sag, Cap, Aqu, Pis:
	default:
		return false
	}

	if p.houseNo < 1 || p.houseNo > 12 {
		return false
	}
	return true
}

type Chart map[PlanetName]Planet

func main() {

}

