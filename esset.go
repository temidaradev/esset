package esset

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	_ "image/png"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func GetAsset(efs embed.FS, path string) *ebiten.Image {
	file, err := efs.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func GetMultiAssets(efs embed.FS, path string) []*ebiten.Image {
	matches, err := fs.Glob(efs, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = GetAsset(efs, match)
	}

	return images
}

var (
	fontFaceSource *text.GoTextFaceSource
)

func DrawText(screen *ebiten.Image, data []byte, str string, fontSize int, posX, posY float64, color color.Color) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource = s

	ffs := &text.GoTextFace{
		Source: s,
		Size:   float64(fontSize),
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(posX, posY)
	op.ColorScale.ScaleWithColor(color)
	op.LayoutOptions.LineSpacing = float64(fontSize)

	text.Draw(screen, str, ffs, op)
}
