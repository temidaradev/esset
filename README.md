[![Go Reference](https://pkg.go.dev/badge/github.com/temidaradev/esset/v2.svg)](https://pkg.go.dev/github.com/temidaradev/esset/v2)

# Esset

[日本語](i18n/README_jp.md) [Türkçe](i18n/README_tr.md) [Slovenčina](i18n/README_sk.md)

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

## DrawText

For fonts you have to embed fonts seperataly and you should add a `text.Face` variable like this:

```
//go:embed path/to/your/font.ttf
var MyFont []byte
var FontFaceS text.Face
var FontFaceM text.Face
var FontFaceM text.Face
```

Then for not loading the font each time you should put `GetFont()` function into `init()` function which is in your main file (contains `Game{}`) like this:

```
func init() {
	assets.FontFaceS, _ = esset.GetFont(assets.MyFont, 48)
}
```

After setting up the `GetFont()` function you are ready to use the `DrawText()` function

No need to create a special DrawOptions for this. Just (screen, text, posX, posY, text.Face, color)

You can use `esset.DrawText` func like this: `esset.DrawText(screen, "wassup", 100, 50, assets.FontFaceS, color.White)`

Much thanks to [@m110](https://github.com/m110) for source support <3
