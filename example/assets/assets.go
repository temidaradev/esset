package assets

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/temidaradev/esset/v2"
)

//go:embed *
var assets embed.FS

//go:embed font/OpenSans-Medium.ttf
var MyFont []byte
var FontFaceS text.Face
var FontFaceM text.Face
var FontFaceL text.Face

var Idle = esset.GetAsset(assets, "mainchar.png")
var Left = esset.GetAsset(assets, "left.png")
var Right = esset.GetAsset(assets, "right.png")

var GopherTile = esset.GetMultiAssets(assets, "gopherTile/*.png")
