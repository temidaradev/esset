[![Go Reference](https://pkg.go.dev/badge/github.com/temidaradev/esset.svg)](https://pkg.go.dev/github.com/temidaradev/esset)

# Esset

Esset, ebitengine için bast bir asset yükleyici.

# Kullanım

Önce `go get github.com/temidaradev/esset` ile go modülünüze ekleyin. Eğer yoksa bir asset klasörü oluşturun. Klasörü oluşturduktan sonra .png varlıklarını bu klasöre koyun ve `assets.go` dosyasını oluşturun. Bundan sonra

```
//go:embed *
var asset embed.FS
```

bu embed ifadesini içe aktarma kısmından sonra ekleyin. Artık esset'i asset yükleyici olarak kullanabilirsiniz. Sizden öncelikle embed ifadesi ve ardından `"asset.png"` olmak üzere 2 parametre ister.

## GetAsset

İşte bir örnek: `var Idle = esset.GetAsset(assets, "path/to/your/asset.png")`

## GetMultipleAssets

Önemli olan bir klasör oluşturmak ve her bir tile öğesini (.png) şu şekilde koymaktır

<img src="../resources/image.png" height="400">

ve sonra bu işlevi şu şekilde kolayca kullanabilirsiniz: `var Tile = esset.GetMultipleAssets(assets, "path/to/your/*.png")` 1'den fazla resim seçtiğiniz için `*ebiten.Image`'imiz şu şekilde dizine göre seçebileceğiniz bir slice'dır: `TileComponent := asset.Tile[0]` veya bu klasörden rastgele bir varlık almanız gerekiyorsa şu şekilde yapabilirsiniz: `TileRandom := asset.Tile[rand.Intn(len(assets.Tile))]`

## DrawText

Yazı tipleri için yazı tiplerini ayrı ayrı gömmelisiniz ve bunun gibi bir `text.Face` değişkeni eklemelisiniz:

```
//go:embed path/to/your/font.ttf
var MyFont []byte
var FontFaceS text.Face
var FontFaceM text.Face
var FontFaceM text.Face
```

Daha sonra yazı tipini her seferinde yüklememek için `GetFont()` fonksiyonunu ana dosyanızda bulunan `init()` fonksiyonuna (`Game{}` içerir) şu şekilde koymalısınız:

```
func init() {
assets.FontFaceS, _ = esset.GetFont(assets.MyFont, 48)
}
```

`GetFont()` fonksiyonunu ayarladıktan sonra `DrawText()` fonksiyonunu kullanmaya hazırsınız

Bunun için özel bir DrawOptions oluşturmanıza gerek yok. Sadece (screen, text, fontSize, posX, posY, text.Face, color)

`esset.DrawText` fonksiyonunu şu şekilde kullanabilirsiniz: `esset.DrawText(screen, "naber", 16, 100, 50, assets.FontFaceS, color.White)`

Çok teşekkürler Kaynak desteği için [@m110](https://github.com/m110) <3
