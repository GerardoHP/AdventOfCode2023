package Days

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	maxRed     = 12
	maxGreen   = 13
	maxBlue    = 14
	gameLabel  = "Game "
	redLabel   = " red"
	greenLabel = " green"
	blueLabel  = " blue"
)

type game struct {
	number     int
	isPossible bool
	tosses     []toss
	original   string
}

type toss struct {
	red   int
	green int
	blue  int
}

func Day2(gamesStr []string) int {
	games := make([]game, 0)
	for _, game := range gamesStr {
		_, g := newGame(game)
		games = append(games, *g)
	}

	total := 0
	for _, v := range games {
		if v.Validate() {
			total += v.number
		} else {
			fmt.Println(v.original)
		}
	}

	return total
}

func (game *game) Validate() bool {
	valid := true
	for _, v := range game.tosses {
		valid = valid && (v.red <= maxRed && v.green <= maxGreen && v.blue <= maxBlue)
	}

	game.isPossible = valid
	return game.isPossible
}

func newGame(str string) (error, *game) {
	gm := strings.Split(str, ":")
	var err error
	if len(gm) != 2 {
		return errors.New("wrong format"), nil
	}

	g := game{original: str}
	g.number, err = strconv.Atoi(strings.ReplaceAll(gm[0], gameLabel, ""))
	if err != nil {
		return err, nil
	}

	throws := strings.Split(gm[1], ";")
	for _, throw := range throws {
		cubes := strings.Split(throw, ",")
		t := toss{}
		for _, cube := range cubes {
			if strings.Contains(cube, redLabel) {
				t.red, _ = strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(cube, redLabel, "")))
				continue
			}

			if strings.Contains(cube, greenLabel) {
				t.green, _ = strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(cube, greenLabel, "")))
				continue
			}

			if strings.Contains(cube, blueLabel) {
				t.blue, _ = strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(cube, blueLabel, "")))
				continue
			}
		}

		g.tosses = append(g.tosses, t)
	}

	return nil, &g
}
