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
    mutex       sync.Mutex
)

type Icon struct {
    svg   *oksvg.SvgIcon
}

func ParseIcon(filename string) Icon {
    mutex.Lock()
    defer mutex.Unlock()
    path := strings.TrimSuffix(IconFolder, "/") + "/" + filename
    if icon, ok := parsedIcons[path]; ok {
        return Icon{svg: icon}
    } else {
        icon, err := oksvg.ReadIcon(path)
        if err != nil {
            panic(err)
        }
        parsedIcons[path] = icon
        return Icon{svg: icon}
    }
}

func (icon Icon) Draw(d draw.Image, r image.Rectangle, c color.Color) {
    alpha := image.NewAlpha(image.Rect(0, 0, int(icon.svg.ViewBox.W), int(icon.svg.ViewBox.H)))
    scanner := rasterx.NewScannerGV(alpha.Bounds().Dx(), alpha.Bounds().Dy(), alpha, alpha.Bounds())
    icon.svg.Draw(rasterx.NewDasher(alpha.Bounds().Dx(), alpha.Bounds().Dy(), scanner), 0.5)
    fixAlpha(alpha)
    resized := resize.Resize(uint(r.Dx()), uint(r.Dy()), alpha, resize.Bicubic)
    draw.DrawMask(d, r, image.NewUniform(c), image.Point{}, resized, image.Point{}, draw.Over)
}

func fixAlpha(alpha *image.Alpha) {
    for x := 0; x < alpha.Bounds().Dx(); x++ {
        for y := 0; y < alpha.Bounds().Dy(); y++ {
            a := uint16(alpha.AlphaAt(x, y).A)
            alpha.SetAlpha(x, y, color.Alpha{A: uint8((a - 127) * 255 / 64)})
        }
    }
}
