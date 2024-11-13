[![Go Reference](https://pkg.go.dev/badge/github.com/temidaradev/esset.svg)](https://pkg.go.dev/github.com/temidaradev/esset)

# Esset

Esset is a basic asset implementer for ebitengine.

# Usage

First `go get github.com/temidaradev/esset` and create an assets folder if you don't have. After creating folder put .png assets into that folder and create `assets.go`. After this add

```
//go:embed *
var assets embed.FS
```

this embed statement after import part. Now you can use esset as you asset implementer. It wants 2 parameters from you firstly embed statement and then your `"asset.png"`.

## GetAsset

Here is an example: `var Idle = esset.GetAsset(assets, "path/to/your/asset.png")`

## GetMultipleAssets

Important thing is create a folder and put every single tile item (.png) like this 

<img src="resources/image.png" height="400"> 

and then you can use that function easily like this: `var Tile = esset.GetMultipleAssets(assets, "path/to/your/*.png")` Because of you are selecting more than 1 image our `*ebiten.Image` is a slice you can select by index like this: `TileComponent := assets.Tile[0]` or if you need to get random asset from that folder you can do like this: `TileRandom := assets.Tile[rand.Intn(len(assets.Tile))]`

## UseFont

For fonts you have to embed fonts seperataly like this:

```
//go:embed font/OpenSans-Medium.ttf
var MyFont []byte
```

After that you should create a `&text.DrawOptions{}` in your `Draw()` func like this:

```
opF := &text.DrawOptions{}
opF.GeoM.Translate(x, y)
opF.ColorScale.ScaleWithColor(color.White)
```

After that you can use `esset.UseFont` func like this: `esset.UseFont(screen, assets.MyFont, "wassup", 24, opF)`


Much thanks to @m110 for source support <3
