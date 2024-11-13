package assets

import (
	"embed"

	"github.com/temidaradev/esset"
)

//go:embed *
var assets embed.FS

var Idle = esset.GetAsset(assets, "mainchar.png")
var Left = esset.GetAsset(assets, "left.png")
var Right = esset.GetAsset(assets, "right.png")
