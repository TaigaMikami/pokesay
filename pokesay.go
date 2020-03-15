package pokesay

import (
	"github.com/TaigaMikami/pokesay/balloon"
	"github.com/TaigaMikami/pokesay/img2xterm"
	"image"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Say(options ...Option) (string, error) {
	poke, err := NewPoke(options...)
	if err != nil {
		return "", err
	}

	pokemon, err := poke.GetPoke()
	if err != nil {
		return "", err
	}

	balloon, err := poke.GetBalloon()
	if err != nil {
		return "", err
	}

	return balloon + pokemon, nil
}

func (poke *Poke) GetPoke() (string, error) {
	file, _ := os.Open(poke.typ)
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}
	pokemon := img2xterm.Img2xterm(img)
	return pokemon, nil
}

func (poke *Poke) GetBalloon() (string, error) {
	var thoughts string
	if poke.thinking {
		thoughts = `         o
          o
`
	} else {
		thoughts = `         \
          \
`
	}

	var s []string
	s = append(s, poke.phrase)
	inputs := balloon.ReadInput(s)
	width := balloon.MaxWidth(inputs)
	messages := balloon.SetPadding(inputs, width)
	balloon := balloon.ConstructBalloon(messages, width)
	return balloon + thoughts, nil
}

func pickPoke() string {
	pokes := AssetNames()
	n := len(pokes)
	return pokes[rand.Intn(n)]
}
