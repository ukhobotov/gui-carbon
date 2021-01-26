package carbon

import (
    "image"
    "image/color"
    "image/draw"
    "strings"
    
    "github.com/golang/freetype"
    "golang.org/x/image/math/fixed"
)

func DrawText(d draw.Image, r image.Rectangle, s string, c color.Color, ts *TextStyle) {
    if ts == nil {
        ts = BodyShort16()
    }
    
    clip := image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Max.Y)
    
    ctx := freetype.NewContext()
    ctx.SetFontSize(ts.Size)
    ctx.SetFont(ts.Font.Load())
    ctx.SetSrc(image.NewUniform(c))
    ctx.SetDst(d)
    ctx.SetClip(clip)
    
    p := fixed.P(r.Min.X, r.Min.Y+int(ts.Size))
    words := strings.Split(s, " ")
    pointClip := fixed.R(clip.Min.X, clip.Min.Y, clip.Max.X, clip.Max.Y)
    
    for _, word := range words {
        
        var err error
        for _, s := range word {
            p, err = ctx.DrawString(string(s), p)
            if err != nil {
                panic(err)
            }
            p.X += ts.LetterSpacing
        }
        
        if p.In(pointClip) {
            p, err = ctx.DrawString(" ", p)
            if err != nil {
                panic(err)
            }
            p.X += ts.LetterSpacing
        } else {
            p.X = fixed.I(r.Min.X)
            p.Y += fixed.I(ts.LineHeight)
        }
        
        if !p.In(pointClip) {
            break
        }
    }
}

type TextStyle struct {
    Size          float64
    LineHeight    int
    Font          FontLoader
    LetterSpacing fixed.Int26_6
}

func Code12() *TextStyle {
    return &TextStyle{Size: 12, LineHeight: 16, Font: PlexMonoRegular, LetterSpacing: 32}
}
func Code14() *TextStyle {
    return &TextStyle{Size: 14, LineHeight: 20, Font: PlexMonoRegular, LetterSpacing: 32}
}
func Label() *TextStyle {
    return &TextStyle{Size: 12, LineHeight: 16, Font: PlexSansRegular, LetterSpacing: 32}
}
func HelperText() *TextStyle {
    return &TextStyle{Size: 12, LineHeight: 16, Font: PlexSansRegular, LetterSpacing: 32}
}
func BodyShort14() *TextStyle {
    return &TextStyle{Size: 14, LineHeight: 18, Font: PlexSansRegular, LetterSpacing: 16}
}
func BodyLong14() *TextStyle {
    return &TextStyle{Size: 14, LineHeight: 20, Font: PlexSansRegular, LetterSpacing: 16}
}
func BodyShort16() *TextStyle {
    return &TextStyle{Size: 16, LineHeight: 22, Font: PlexSansRegular}
}
func BodyLong16() *TextStyle {
    return &TextStyle{Size: 16, LineHeight: 24, Font: PlexSansRegular}
}
func ProductiveHeading14() *TextStyle {
    return &TextStyle{Size: 14, LineHeight: 28, Font: PlexSansSemiBold, LetterSpacing: 16}
}
func ProductiveHeading16() *TextStyle {
    return &TextStyle{Size: 16, LineHeight: 22, Font: PlexSansSemiBold}
}
func ProductiveHeading20() *TextStyle {
    return &TextStyle{Size: 20, LineHeight: 28, Font: PlexSansRegular}
}
func ProductiveHeading28() *TextStyle {
    return &TextStyle{Size: 28, LineHeight: 36, Font: PlexSansRegular}
}
func ProductiveHeading32() *TextStyle {
    return &TextStyle{Size: 32, LineHeight: 40, Font: PlexSansRegular}
}
func ProductiveHeading42() *TextStyle {
    return &TextStyle{Size: 42, LineHeight: 50, Font: PlexSansLight}
}
func ProductiveHeading54() *TextStyle {
    return &TextStyle{Size: 54, LineHeight: 64, Font: PlexSansLight}
}
