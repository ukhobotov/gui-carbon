package carbon

import (
    "image"
    "image/color"
    "image/draw"
)

type Shadow struct {
    Offset image.Point
    Blur   int
    Color  color.Color
}

var (
    Raised       = Shadow{Offset: image.Pt(0, 1), Blur: 2, Color: ShadowColor}
    Overlay      = Shadow{Offset: image.Pt(0, 4), Blur: 8, Color: ShadowColor}
    StickyNav    = Shadow{Offset: image.Pt(0, 6), Blur: 12, Color: ShadowColor}
    TemporaryNav = Shadow{Offset: image.Pt(0, 8), Blur: 16, Color: ShadowColor}
    PopOut       = Shadow{Offset: image.Pt(0, 12), Blur: 24, Color: ShadowColor}
)

func DrawShadow(dst draw.Image, r image.Rectangle, s Shadow) image.Rectangle {
    R := r.Add(s.Offset)
    draw.Draw(dst, R, image.NewUniform(s.Color), image.Point{}, draw.Over)
    
    draw.DrawMask(
        dst, image.Rect(R.Min.X-s.Blur, R.Min.Y-s.Blur, R.Min.X, R.Min.Y), image.NewUniform(s.Color), image.Point{},
        RadGrad{image.Pt(s.Blur, s.Blur), s.Blur}, image.Point{}, draw.Over,
    )
    draw.DrawMask(
        dst, image.Rect(R.Max.X, R.Min.Y-s.Blur, R.Max.X+s.Blur, R.Min.Y), image.NewUniform(s.Color), image.Point{},
        RadGrad{image.Pt(0, s.Blur), s.Blur}, image.Point{}, draw.Over,
    )
    draw.DrawMask(
        dst, image.Rect(R.Max.X, R.Max.Y, R.Max.X+s.Blur, R.Max.Y+s.Blur), image.NewUniform(s.Color), image.Point{},
        RadGrad{image.Pt(0, 0), s.Blur}, image.Point{}, draw.Over,
    )
    draw.DrawMask(
        dst, image.Rect(R.Min.X-s.Blur, R.Max.Y, R.Min.X, R.Max.Y+s.Blur), image.NewUniform(s.Color), image.Point{},
        RadGrad{image.Pt(s.Blur, 0), s.Blur}, image.Point{}, draw.Over,
    )
    
    draw.DrawMask(
        dst, image.Rect(R.Min.X, R.Min.Y-s.Blur, R.Max.X, R.Min.Y), image.NewUniform(s.Color), image.Point{},
        LinGrad{image.Pt(0, s.Blur), image.Pt(0, 0)}, image.Point{}, draw.Over,
    )
    draw.DrawMask(
        dst, image.Rect(R.Max.X, R.Min.Y, R.Max.X+s.Blur, R.Max.Y), image.NewUniform(s.Color), image.Point{},
        LinGrad{image.Pt(0, 0), image.Pt(s.Blur, 0)}, image.Point{}, draw.Over,
    )
    draw.DrawMask(
        dst, image.Rect(R.Min.X, R.Max.Y, R.Max.X, R.Max.Y+s.Blur), image.NewUniform(s.Color), image.Point{},
        LinGrad{image.Pt(0, 0), image.Pt(0, s.Blur)}, image.Point{}, draw.Over,
    )
    draw.DrawMask(
        dst, image.Rect(R.Min.X-s.Blur, R.Min.Y, R.Min.X, R.Max.Y), image.NewUniform(s.Color), image.Point{},
        LinGrad{image.Pt(s.Blur, 0), image.Pt(0, 0)}, image.Point{}, draw.Over,
    )
    
    return image.Rect(R.Min.X-s.Blur, R.Min.Y-s.Blur, R.Max.X+s.Blur, R.Max.Y+s.Blur)
}

type LinGrad struct {
    Opaque, Transparent image.Point
}

func (l LinGrad) ColorModel() color.Model { return color.AlphaModel }

func (l LinGrad) Bounds() image.Rectangle { return image.Rect(-1e9, -1e9, 1e9, 1e9) }

func (l LinGrad) At(x, y int) color.Color {
    
    u := l.Transparent.Sub(l.Opaque)
    v := image.Pt(x, y).Sub(l.Opaque)
    vu := u.X*v.X + u.Y*v.Y
    u2 := u.X*u.X + u.Y*u.Y
    
    if vu <= 0 {
        return color.Alpha{A: 255}
    }
    if vu >= u2 {
        return color.Alpha{}
    }
    return color.Alpha{A: uint8(255 - 255*vu/u2)}
}

type RadGrad struct {
    Opaque image.Point
    Radius int
}

func (r RadGrad) ColorModel() color.Model { return color.AlphaModel }

func (r RadGrad) Bounds() image.Rectangle { return image.Rect(-1e9, -1e9, 1e9, 1e9) }

func (r RadGrad) At(x, y int) color.Color {
    v := image.Pt(x, y).Sub(r.Opaque)
    v2 := v.X*v.X + v.Y*v.Y
    if v2 >= r.Radius*r.Radius {
        return color.Alpha{}
    }
    return color.Alpha{A: uint8(255 - 255*intSqrt(v2)/r.Radius)}
}

func intSqrt(x int) int {
    if x < 0 {
        panic("negative argument")
    }
    if x < 2 {
        return x
    }
    inf := intSqrt(x>>2) << 1
    sup := inf + 1
    if sup*sup > x {
        return inf
    }
    return sup
}
