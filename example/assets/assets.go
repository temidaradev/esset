package assets

import (
	"embed"

	"github.com/temidaradev/esset"
)

//go:embed *
var assets embed.FS

var Idle = esset.GetSingleAsset(assets, "mainchar.png")
var Left = esset.GetSingleAsset(assets, "left.png")
var Right = esset.GetSingleAsset(assets, "right.png")
