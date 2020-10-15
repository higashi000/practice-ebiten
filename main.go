package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	H bool
	J bool
	K bool
	L bool
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.H = ebiten.IsKeyPressed(ebiten.KeyH)
	g.J = ebiten.IsKeyPressed(ebiten.KeyJ)
	g.K = ebiten.IsKeyPressed(ebiten.KeyK)
	g.L = ebiten.IsKeyPressed(ebiten.KeyL)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.H {
		ebitenutil.DebugPrintAt(screen, "H", 0, 0)
	}
	if g.J {
		ebitenutil.DebugPrintAt(screen, "J", 0, 20)
	}
	if g.K {
		ebitenutil.DebugPrintAt(screen, "K", 0, 40)
	}
	if g.L {
		ebitenutil.DebugPrintAt(screen, "L", 0, 60)
	}
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
