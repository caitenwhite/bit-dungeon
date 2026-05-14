package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &Game{}
	// testing crlf -> lf conversion

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
