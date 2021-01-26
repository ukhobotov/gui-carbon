package carbon

import (
    "image"
    "image/color"
    "image/draw"
    
    "github.com/srwiley/oksvg"
    "github.com/ukhobotov/gui"
    "github.com/ukhobotov/gui/win"
)

func TextButton(env gui.Env, bs *ButtonStyle, text string, act chan<- gui.Event) {
    Button(env, bs, text, nil, act)
}

func IconButton(env gui.Env, bs *ButtonStyle, icon *oksvg.SvgIcon, act chan<- gui.Event) {
    Button(env, bs, "", icon, act)
}

func Button(env gui.Env, bs *ButtonStyle, text string, icon *oksvg.SvgIcon, act chan<- gui.Event) {
    if bs == nil {
        panic("launching button with a nil style")
    }
    var (
        r image.Rectangle
        b = buttonData{
            text:  text,
            icon:  icon,
            style: bs,
        }
    )
    
    in := make(chan gui.Event)
    go gui.StartQueue(in, act)
    
    for event := range env.Events() {
        switch e := event.(type) {
        case gui.Resize:
            r = e.Bounds()
            env.Draw() <- drawButton(r, b)
        case win.MoMove:
            switch {
            case e.In(r) && !b.hovered:
                in <- event
                b.hovered = true
                env.Draw() <- drawButton(r, b)
            case !e.In(r) && b.hovered:
                b.hovered = false
                b.pressed = false
                env.Draw() <- drawButton(r, b)
            }
        case win.MoDown:
            switch {
            case e.In(r):
                in <- event
                b.pressed = true
                b.focused = true
                env.Draw() <- drawButton(r, b)
            case !e.In(r) && b.focused:
                b.focused = false
                env.Draw() <- drawButton(r, b)
            }
        case win.MoUp:
            if e.In(r) && b.pressed {
                in <- event
                b.pressed = false
                env.Draw() <- drawButton(r, b)
            }
        }
    }
}

type buttonData struct {
    text  string
    icon  *oksvg.SvgIcon
    style *ButtonStyle
    
    hovered, pressed, focused, disabled bool
}

func drawButton(r image.Rectangle, b buttonData) Drawer {
    
    return func(d draw.Image) image.Rectangle {
        switch {
        case b.disabled:
            drawButtonState(d, r, b, b.style.Disabled)
            return r
        case b.pressed:
            drawButtonState(d, r, b, b.style.Active)
        case b.hovered:
            drawButtonState(d, r, b, b.style.Hover)
        default:
            drawButtonState(d, r, b, b.style.Default)
        }
        if b.focused {
            drawButtonState(d, r, b, b.style.Focus)
        }
        return r
    }
}

type ButtonStyle struct {
    Default, Hover, Active, Focus, Disabled *ButtonStateStyle
    TextStyle                               *TextStyle
    IconSize                                int
}

type ButtonStateStyle struct {
    Background                       color.Color
    Border                           Border
    InsetColor, TextColor, IconColor color.Color
}

func drawButtonState(d draw.Image, r image.Rectangle, b buttonData, s *ButtonStateStyle) {
    if s.Background != nil {
        DrawRect(d, r, s.Background)
    }
    if s.Border != (Border{}) {
        DrawBorder(d, r, s.Border)
    }
    if s.InsetColor != nil {
        DrawBorder(
            d, image.Rect(
                r.Min.X+s.Border.Width, r.Min.Y+s.Border.Width,
                r.Max.X-s.Border.Width, r.Max.Y-s.Border.Width,
            ), Border{Color: s.InsetColor},
        )
    }
    if b.text != "" && s.TextColor != nil && b.style.TextStyle != nil {
        marginLeft := 16
        DrawText(d, image.Rect(
            r.Min.X+marginLeft, (r.Min.Y+r.Max.Y-b.style.TextStyle.LineHeight)/2,
            r.Max.X, (r.Min.Y+r.Max.Y+b.style.TextStyle.LineHeight)/2,
        ), b.text, s.TextColor, b.style.TextStyle)
    }
    if b.icon != nil && s.IconColor != nil {
        if b.text != "" {
            const marginRight = 16
            DrawIcon(d, image.Rect(
                r.Max.X-marginRight-b.style.IconSize, (r.Min.Y+r.Max.Y)/2-b.style.IconSize/2,
                r.Max.X-marginRight, (r.Min.Y+r.Max.Y)/2+b.style.IconSize/2,
            ), s.IconColor, b.icon)
        } else {
            DrawIcon(d, image.Rect(
                (r.Min.X+r.Max.X)/2-b.style.IconSize/2, (r.Min.Y+r.Max.Y)/2-b.style.IconSize/2,
                (r.Min.X+r.Max.X)/2+b.style.IconSize/2, (r.Min.Y+r.Max.Y)/2+b.style.IconSize/2,
            ), s.IconColor, b.icon)
        }
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
            Border:     Border{Color: Focus(), Width: 2},
            InsetColor: UiBackground(),
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
            Border:     Border{Color: Focus(), Width: 2},
            InsetColor: UiBackground(),
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
            InsetColor: UiBackground(),
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
            Border:     Border{Focus(), 2},
            InsetColor: UiBackground(),
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
    }
}
