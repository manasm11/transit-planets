package main

import (
	"fmt"
	"strings"
)

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

func (pn PlanetName) ShortName() string {
	switch pn {
	case As:
		return "As"
	case Su:
		return "Su"
	case Mo:
		return "Mo"
	case Me:
		return "Me"
	case Ve:
		return "Ve"
	case Ma:
		return "Ma"
	case Ju:
		return "Ju"
	case Ra:
		return "Ra"
	case Ke:
		return "Ke"
	case Sa:
		return "Sa"
	}
	panic("Incorrect Planet Name")
}

type Planet struct {
	name    PlanetName
	houseNo int
}

func (p Planet) IsValid() bool {
	switch p.name {
	case As, Su, Mo, Me, Ve, Ma, Ju, Ra, Ke, Sa:
	default:
		return false
	}

	if p.houseNo < 1 || p.houseNo > 12 {
		return false
	}
	return true
}

type Chart struct {
	planets []Planet
	asc     Zodiac
}

func (c Chart) String() string {
	var chart string = chartString
	for _, v := range c.planets {
		chart = strings.Replace(chart,
			fmt.Sprintf("__H%X__", v.houseNo),
			fmt.Sprintf("%-6s", v.name.ShortName()),
			1)
	}

	for hn := range 12 {
		hn++
		chart = strings.ReplaceAll(chart,
			fmt.Sprintf("__H%X__", hn),
			"      ", // 6 blank spaces
		)
	}

	for zn := range 12 {
		chart = strings.ReplaceAll(chart,
			fmt.Sprintf("Z%X", zn+1),
			fmt.Sprintf("%-2d", (int(c.asc)+zn-1)%12+1),
		)

	}

	return chart
}

func main() {
	manas := Chart{}
	manas.asc = Leo

	manas.planets = []Planet{
		{Su, 4}, {Ra, 1}, {Ma, 2},
		{Me, 4}, {Ve, 4}, {Mo, 5},
		{Ke, 7}, {Ju, 7}, {Sa, 9},
	}

	fmt.Println("Lagna Chart", manas.String())

}

func ensureChars(s string, k int) string {
	if len(s) < k {
		return s + strings.Repeat(" ", k-len(s))
	}
	return s
}

var chartString = `
+-----------------------------------------------------------------------------+
|\                                     ^                                     /|
|  \    __H2__  __H2__ __H2__ __H2__ /   \  __HC__  __HC__  __HC__  __HC__ /  |
|    \                             /       \                             /    |
|      \     __H2__   __H2__     /           \     __HC__   __HC__     /      |
|  __H3__\                     /               \                     / __HB__ |
|          \   __H2__ __H2__ /       __H1__      \      __HC__     /          |
|  __H3__    \             /                       \             /   __HB__   |
|              \         /                           \  __HC__ /              |
|  __H3__ __H3__ \  Z2 /     __H1__  __H1__  __H1__    \  ZC / __HB__  __HB__ |
|               Z3 \_/                                   \_/ ZB               |
|  __H3__ __H3__   / \                                   / \   __HB__  __HB__ |
|                /     \     __H1__  __H1__  __H1__    /     \                |
|  __H3__      /         \                           /         \    __HB__    |
|            /             \                       /             \            |
|  __H3__  /     __H4__      \       __H1__      /     __HA__      \   __HB__ |
|        /                     \               /                     \        |
|      /                         \           /                         \      |
|    /      __H4__    __H4__       \   Z1  /      __HA__    __HA__       \    |
|  /                                 \   /                                 \  |
|<     __H4__    __H4__    __H4__  Z4  X  ZA  __HA__    __HA__    __HA__     >|
|  \                                 /   \                                 /  |
|    \                             /       \                             /    |
|      \    __H4__    __H4__     /     Z7    \    __HA__    __HA__     /      |
|  __H5__\                     /               \                     / __H9__ |
|          \                 /       __H7__      \                 /          |
|  __H5__    \             /                       \             /   __H9__   |
|              \         /                           \         /              |
|  __H5__ __H5__ \     /     __H7__  __H7__  __H7__    \     / __H9__  __H9__ |
|               Z5 \_/                                   \_/ Z9               |
|  __H5__  __H5__  / \                                   / \   __H9__  __H9__ |
|                /  Z6 \     __H7__  __H7__  __H7__    /  Z8 \                |
|  __H5__      /  __H6__ \                           /  __H8__ \     __H9__   |
|            /             \                       /             \            |
|  __H5__  /   __H6__ __H6__ \       __H7__      /   __H8__ __H8__ \   __H9__ |
|        /                     \               /                     \        |
|      /    __H6__     __H6__    \           /    __H8__     __H8__    \      |
|    /                             \       /                             \    |
|  /    __H6__   __H6__   __H6__     \   /    __H8__   __H8__   __H8__     \  |
|/                                     V                                     \|
+-----------------------------------------------------------------------------+
`

