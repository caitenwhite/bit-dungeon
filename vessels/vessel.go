package vessels

import (
	"embed"
	"image"
	_ "image/png"

	//"github.com/caitenwhite/bit-dungeon/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

func MustLoadSprite(file string) *ebiten.Image {
	f, err := assets.Open(file)
	if err != nil {
		// TBD better error handler
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
