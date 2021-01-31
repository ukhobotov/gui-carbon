package carbon

import (
    "image"
    "image/color"
    "image/draw"
    
    "github.com/ukhobotov/gui"
    "github.com/ukhobotov/gui/win"
)

func TextButton(env gui.Env, bs *ButtonStyle, text string, click chan<- win.MoUp) {
    Button(env, bs, text, Icon{}, click)
}

func IconButton(env gui.Env, bs *ButtonStyle, icon Icon, click chan<- win.MoUp) {
    Button(env, bs, "", icon, click)
}

func Button(env gui.Env, bs *ButtonStyle, text string, icon Icon, click chan<- win.MoUp) {
    if bs == nil {
        panic("launching button with a nil style")
    }
    var (
        r     image.Rectangle
        state State
    )
    
    for event := range env.Events() {
        switch e := event.(type) {
        case gui.Resize:
            r = e.Bounds()
            env.Draw() <- drawButton(r, text, icon, *bs, state)
        case win.MoMove:
            switch {
            case e.In(r) && state&Hover == 0:
                state |= Hover
                env.Draw() <- drawButton(r, text, icon, *bs, state)
            case !e.In(r) && state&Hover != 0:
                state &^= Hover | Active
                env.Draw() <- drawButton(r, text, icon, *bs, state)
            }
        case win.MoDown:
            switch {
            case e.In(r):
                state |= Active | Focused
                env.Draw() <- drawButton(r, text, icon, *bs, state)
            case !e.In(r) && state&Focused != 0:
                state &^= Focused
                env.Draw() <- drawButton(r, text, icon, *bs, state)
            }
        case win.MoUp:
            if e.In(r) && state&Active != 0 {
                click <- e
                state &^= Active
                env.Draw() <- drawButton(r, text, icon, *bs, state)
            }
        }
    }
}

type ButtonStyle struct {
    Default, Hover, Active, Focus, Disabled *ButtonStateStyle
    TextStyle                               *TextStyle
    IconSize                                int
}

func drawButton(r image.Rectangle, text string, icon Icon, s ButtonStyle, state State) Drawer {
    if state&Disabled != 0 {
        return drawButtonState(r, text, icon, *s.Disabled, *s.TextStyle, s.IconSize)
    }
    var base Drawer
    switch {
    case state&Active != 0:
        base = drawButtonState(r, text, icon, *s.Active, *s.TextStyle, s.IconSize)
    case state&Hover != 0:
        base = drawButtonState(r, text, icon, *s.Hover, *s.TextStyle, s.IconSize)
    default:
        base = drawButtonState(r, text, icon, *s.Default, *s.TextStyle, s.IconSize)
    }
    if state&Focused != 0 {
        focus := drawButtonState(r, text, icon, *s.Focus, *s.TextStyle, s.IconSize)
        return func(d draw.Image) image.Rectangle {
            base(d)
            focus(d)
            return r
        }
    }
    return base
}

type ButtonStateStyle struct {
    Background                  color.Color
    Border                      Border
    Inset, TextColor, IconColor color.Color
}

func drawButtonState(r image.Rectangle, text string, icon Icon, s ButtonStateStyle, ts TextStyle, is int) Drawer {
    return func(d draw.Image) image.Rectangle {
        if s.Background != nil {
            DrawRect(d, r, s.Background)
        }
        if s.Border != (Border{}) {
            DrawBorder(d, r, s.Border)
        }
        if s.Inset != nil {
            DrawBorder(
                d, image.Rect(
                    r.Min.X+s.Border.Width, r.Min.Y+s.Border.Width,
                    r.Max.X-s.Border.Width, r.Max.Y-s.Border.Width,
                ), Border{Color: s.Inset},
            )
        }
        if text != "" && s.TextColor != nil {
            marginLeft := 16
            DrawText(d, image.Rect(
                r.Min.X+marginLeft, (r.Min.Y+r.Max.Y-ts.LineHeight)/2,
                r.Max.X, (r.Min.Y+r.Max.Y+ts.LineHeight)/2,
            ), text, s.TextColor, ts)
        }
        if icon != (Icon{}) && s.IconColor != nil {
            if text != "" {
                const marginRight = 16
                icon.Draw(d, image.Rect(
                    r.Max.X-marginRight-is, (r.Min.Y+r.Max.Y)/2-is/2,
                    r.Max.X-marginRight, (r.Min.Y+r.Max.Y)/2+is/2,
                ), s.IconColor)
            } else {
                icon.Draw(d, image.Rect(
                    (r.Min.X+r.Max.X)/2-is/2, (r.Min.Y+r.Max.Y)/2-is/2,
                    (r.Min.X+r.Max.X)/2+is/2, (r.Min.Y+r.Max.Y)/2+is/2,
                ), s.IconColor)
            }
        }
        return r
    }
}

func PrimaryButton() *ButtonStyle {
    return &ButtonStyle{
        Default: &ButtonStateStyle{
            Background: Interactive1(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Hover: &ButtonStateStyle{
            Background: HoverPrimary(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Active: &ButtonStateStyle{
            Background: ActivePrimary(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Focus: &ButtonStateStyle{
            Border: Border{Color: Focus(), Width: 2},
            Inset:  UiBackground(),
        },
        Disabled: &ButtonStateStyle{
            Background: Disabled2(),
            TextColor:  Disabled3(),
            IconColor:  Disabled3(),
        },
        TextStyle: BodyShort14(),
        IconSize:  16,
    }
}

func SecondaryButton() *ButtonStyle {
    return &ButtonStyle{
        Default: &ButtonStateStyle{
            Background: Interactive2(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Hover: &ButtonStateStyle{
            Background: HoverSecondary(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Active: &ButtonStateStyle{
            Background: ActiveSecondary(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Focus: &ButtonStateStyle{
            Border: Border{Color: Focus(), Width: 2},
            Inset:  UiBackground(),
        },
        Disabled: &ButtonStateStyle{
            Background: Disabled2(),
            TextColor:  Disabled3(),
            IconColor:  Disabled3(),
        },
        TextStyle: BodyShort14(),
        IconSize:  16,
    }
}

func TertiaryButton(background color.Color) *ButtonStyle {
    return &ButtonStyle{
        Default: &ButtonStateStyle{
            Background: background,
            Border:     Border{Color: Interactive3()},
            TextColor:  Interactive3(),
            IconColor:  Interactive3(),
        },
        Hover: &ButtonStateStyle{
            Background: HoverTertiary(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Active: &ButtonStateStyle{
            Background: ActiveTertiary(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Focus: &ButtonStateStyle{
            Background: Interactive3(),
            Border:     Border{Color: Focus(), Width: 2},
            Inset:      UiBackground(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Disabled: &ButtonStateStyle{
            Background: background,
            Border:     Border{Color: Disabled2()},
            TextColor:  Disabled2(),
            IconColor:  Disabled2(),
        },
        TextStyle: BodyShort14(),
        IconSize:  16,
    }
}

func GhostButton(background color.Color) *ButtonStyle {
    return &ButtonStyle{
        Default: &ButtonStateStyle{
            Background: background,
            TextColor:  Link1(),
            IconColor:  Link1(),
        },
        Hover: &ButtonStateStyle{
            Background: HoverUi(),
            TextColor:  Link1(),
            IconColor:  Link1(),
        },
        Active: &ButtonStateStyle{
            Background: ActiveUi(),
            TextColor:  Link1(),
            IconColor:  Link1(),
        },
        Focus: &ButtonStateStyle{
            Border: Border{Focus(), 2},
        },
        Disabled: &ButtonStateStyle{
            Background: background,
            TextColor:  Disabled2(),
            IconColor:  Disabled2(),
        },
        TextStyle: BodyShort14(),
        IconSize:  16,
    }
}

func DangerButton() *ButtonStyle {
    return &ButtonStyle{
        Default: &ButtonStateStyle{
            Background: Support1(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Hover: &ButtonStateStyle{
            Background: HoverDanger(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Active: &ButtonStateStyle{
            Background: ActiveDanger(),
            TextColor:  Text4(),
            IconColor:  Icon3(),
        },
        Focus: &ButtonStateStyle{
            Border: Border{Focus(), 2},
            Inset:  UiBackground(),
        },
        Disabled: &ButtonStateStyle{
            Background: Disabled2(),
            TextColor:  Disabled3(),
            IconColor:  Disabled3(),
        },
        TextStyle: BodyShort14(),
        IconSize:  16,
    }
}

func UiButton() *ButtonStyle {
    return &ButtonStyle{
        Default: &ButtonStateStyle{
            Background: Ui1(),
            TextColor:  Text1(),
            IconColor:  Icon1(),
        },
        Hover: &ButtonStateStyle{
            Background: HoverUi(),
            TextColor:  Text1(),
            IconColor:  Icon1(),
        },
        Active: &ButtonStateStyle{
            Background: ActiveUi(),
            TextColor:  Text1(),
            IconColor:  Icon1(),
        },
        Focus: &ButtonStateStyle{
            Border: Border{Color: Focus(), Width: 2},
        },
        Disabled: &ButtonStateStyle{
            Background: Disabled2(),
            TextColor:  Disabled3(),
            IconColor:  Disabled3(),
        },
        TextStyle: BodyShort14(),
        IconSize:  20,
    }
}

func HeaderButton() *ButtonStyle {
    return &ButtonStyle{
        Default: &ButtonStateStyle{
            Background: UiBackground(),
            TextColor:  Text1(),
            IconColor:  Icon1(),
        },
        Hover: &ButtonStateStyle{
            Background: HoverUi(),
            TextColor:  Text1(),
            IconColor:  Icon1(),
        },
        Active: &ButtonStateStyle{
            Background: ActiveUi(),
            TextColor:  Text1(),
            IconColor:  Icon1(),
        },
        Focus: &ButtonStateStyle{
            Border: Border{Color: Focus(), Width: 2},
        },
        Disabled: &ButtonStateStyle{
            Background: Disabled2(),
            TextColor:  Disabled3(),
            IconColor:  Disabled3(),
        },
        TextStyle: BodyShort14(),
        IconSize:  20,
    }
}
