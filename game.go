package main

import (
	"github.com/hajimehoshi/ebiten/v2"
       // "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	speed = 3.0
)

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	playerPosition Vector
}

func (g *Game) Update() error {
	// create switch case for movement

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerPosition.Y += speed
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerPosition.Y -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerPosition.X += speed
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerPosition.X -= speed
	}

        return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	screen.DrawImage(PlayerSprite, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
        return 320, 240
}
