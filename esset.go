package esset

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetSingleAsset(fs embed.FS, name string) *ebiten.Image {
	file, err := fs.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
