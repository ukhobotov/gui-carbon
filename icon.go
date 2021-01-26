package carbon

import (
    "image"
    "image/color"
    "image/draw"
    "strings"
    "sync"
    
    "github.com/nfnt/resize"
    "github.com/srwiley/oksvg"
    "github.com/srwiley/rasterx"
)

var (
    IconFolder  = "."
    parsedIcons = map[string]*oksvg.SvgIcon{}
    mu          sync.Mutex
)

func ParseIcon(filename string) *oksvg.SvgIcon {
    mu.Lock()
    defer mu.Unlock()
    path := strings.TrimSuffix(IconFolder, "/") + "/" + filename
    if icon, ok := parsedIcons[path]; ok {
        return icon
    } else {
        icon, err := oksvg.ReadIcon(path)
        if err != nil {
            panic(err)
        }
        parsedIcons[path] = icon
        return icon
    }
}

func DrawIcon(d draw.Image, r image.Rectangle, c color.Color, icon *oksvg.SvgIcon) {
    alpha := image.NewAlpha(image.Rect(0, 0, int(icon.ViewBox.W), int(icon.ViewBox.H)))
    scanner := rasterx.NewScannerGV(alpha.Bounds().Dx(), alpha.Bounds().Dy(), alpha, alpha.Bounds())
    icon.Draw(rasterx.NewDasher(alpha.Bounds().Dx(), alpha.Bounds().Dy(), scanner), 0.5)
    fixAlpha(alpha)
    resized := resize.Resize(uint(r.Dx()), uint(r.Dy()), alpha, resize.Bicubic)
    draw.DrawMask(d, r, image.NewUniform(c), image.Point{}, resized, image.Point{}, draw.Over)
}

func fixAlpha(alpha *image.Alpha) {
    for x := 0; x < alpha.Bounds().Dx(); x++ {
        for y := 0; y < alpha.Bounds().Dy(); y++ {
            a := uint16(alpha.AlphaAt(x, y).A)
            alpha.SetAlpha(x, y, color.Alpha{uint8((a - 127) * 255 / 64)})
        }
    }
}
