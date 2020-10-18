package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	Space bool
    Height uint64
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.Space = ebiten.IsKeyPressed(ebiten.KeySpace)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Space {
        g.Height++
	}

    if g.Height != 0 && !g.Space {
        g.Height--
    }

    ebitenutil.DebugPrint(screen, strconv.FormatUint(g.Height, 10))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)

	ebiten.SetWindowTitle("ebiten keyboard")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
