package carbon

import "image/color"

func UiBackground() color.Color     { return CurrentTheme.UiBackground }
func Interactive1() color.Color     { return CurrentTheme.Interactive1 }
func Interactive2() color.Color     { return CurrentTheme.Interactive2 }
func Interactive3() color.Color     { return CurrentTheme.Interactive3 }
func Interactive4() color.Color     { return CurrentTheme.Interactive4 }
func Danger() color.Color           { return CurrentTheme.Danger }
func Ui1() color.Color              { return CurrentTheme.Ui1 }
func Ui2() color.Color              { return CurrentTheme.Ui2 }
func Ui3() color.Color              { return CurrentTheme.Ui3 }
func Ui4() color.Color              { return CurrentTheme.Ui4 }
func Ui5() color.Color              { return CurrentTheme.Ui5 }
func ButtonSeparator() color.Color  { return CurrentTheme.ButtonSeparator }
func Decorative() color.Color       { return CurrentTheme.Decorative }
func Text1() color.Color            { return CurrentTheme.Text1 }
func Text2() color.Color            { return CurrentTheme.Text2 }
func Text3() color.Color            { return CurrentTheme.Text3 }
func Text4() color.Color            { return CurrentTheme.Text4 }
func Text5() color.Color            { return CurrentTheme.Text5 }
func TextError() color.Color        { return CurrentTheme.TextError }
func Link1() color.Color            { return CurrentTheme.Link1 }
func InverseLink() color.Color      { return CurrentTheme.InverseLink }
func Icon1() color.Color            { return CurrentTheme.Icon1 }
func Icon2() color.Color            { return CurrentTheme.Icon2 }
func Icon3() color.Color            { return CurrentTheme.Icon3 }
func Field1() color.Color           { return CurrentTheme.Field1 }
func Field2() color.Color           { return CurrentTheme.Field2 }
func Inverse1() color.Color         { return CurrentTheme.Inverse1 }
func Inverse2() color.Color         { return CurrentTheme.Inverse2 }
func Support1() color.Color         { return CurrentTheme.Support1 }
func Support2() color.Color         { return CurrentTheme.Support2 }
func Support3() color.Color         { return CurrentTheme.Support3 }
func Support4() color.Color         { return CurrentTheme.Support4 }
func InverseSupport1() color.Color  { return CurrentTheme.InverseSupport1 }
func InverseSupport2() color.Color  { return CurrentTheme.InverseSupport2 }
func InverseSupport3() color.Color  { return CurrentTheme.InverseSupport3 }
func InverseSupport4() color.Color  { return CurrentTheme.InverseSupport4 }
func Overlay1() color.Color         { return CurrentTheme.Overlay1 }
func Focus() color.Color            { return CurrentTheme.Focus }
func InverseFocusUi() color.Color   { return CurrentTheme.InverseFocusUi }
func HoverPrimary() color.Color     { return CurrentTheme.HoverPrimary }
func HoverPrimaryText() color.Color { return CurrentTheme.HoverPrimaryText }
func HoverSecondary() color.Color   { return CurrentTheme.HoverSecondary }
func HoverTertiary() color.Color    { return CurrentTheme.HoverTertiary }
func HoverUi() color.Color          { return CurrentTheme.HoverUi }
func HoverLightUi() color.Color     { return CurrentTheme.HoverLightUi }
func HoverSelectedUi() color.Color  { return CurrentTheme.HoverSelectedUi }
func HoverDanger() color.Color      { return CurrentTheme.HoverDanger }
func HoverRow() color.Color         { return CurrentTheme.HoverRow }
func InverseHoverUi() color.Color   { return CurrentTheme.InverseHoverUi }
func ActivePrimary() color.Color    { return CurrentTheme.ActivePrimary }
func ActiveSecondary() color.Color  { return CurrentTheme.ActiveSecondary }
func ActiveTertiary() color.Color   { return CurrentTheme.ActiveTertiary }
func ActiveUi() color.Color         { return CurrentTheme.ActiveUi }
func ActiveLightUi() color.Color    { return CurrentTheme.ActiveLightUi }
func ActiveDanger() color.Color     { return CurrentTheme.ActiveDanger }
func SelectedUi() color.Color       { return CurrentTheme.SelectedUi }
func SelectedLightUi() color.Color  { return CurrentTheme.SelectedLightUi }
func Highlight() color.Color        { return CurrentTheme.Highlight }
func Skeleton1() color.Color        { return CurrentTheme.Skeleton1 }
func Skeleton2() color.Color        { return CurrentTheme.Skeleton2 }
func VisitedLink() color.Color      { return CurrentTheme.VisitedLink }
func Disabled1() color.Color        { return CurrentTheme.Disabled1 }
func Disabled2() color.Color        { return CurrentTheme.Disabled2 }
func Disabled3() color.Color        { return CurrentTheme.Disabled3 }

var CurrentTheme = Gray100Theme()

type Theme struct {
    UiBackground     color.Color
    Interactive1     color.Color
    Interactive2     color.Color
    Interactive3     color.Color
    Interactive4     color.Color
    Danger           color.Color
    Ui1              color.Color
    Ui2              color.Color
    Ui3              color.Color
    Ui4              color.Color
    Ui5              color.Color
    ButtonSeparator  color.Color
    Decorative       color.Color
    Text1            color.Color
    Text2            color.Color
    Text3            color.Color
    Text4            color.Color
    Text5            color.Color
    TextError        color.Color
    Link1            color.Color
    InverseLink      color.Color
    Icon1            color.Color
    Icon2            color.Color
    Icon3            color.Color
    Field1           color.Color
    Field2           color.Color
    Inverse1         color.Color
    Inverse2         color.Color
    Support1         color.Color
    Support2         color.Color
    Support3         color.Color
    Support4         color.Color
    InverseSupport1  color.Color
    InverseSupport2  color.Color
    InverseSupport3  color.Color
    InverseSupport4  color.Color
    Overlay1         color.Color
    Focus            color.Color
    InverseFocusUi   color.Color
    HoverPrimary     color.Color
    HoverPrimaryText color.Color
    HoverSecondary   color.Color
    HoverTertiary    color.Color
    HoverUi          color.Color
    HoverLightUi     color.Color
    HoverSelectedUi  color.Color
    HoverDanger      color.Color
    HoverRow         color.Color
    InverseHoverUi   color.Color
    ActivePrimary    color.Color
    ActiveSecondary  color.Color
    ActiveTertiary   color.Color
    ActiveUi         color.Color
    ActiveLightUi    color.Color
    ActiveDanger     color.Color
    SelectedUi       color.Color
    SelectedLightUi  color.Color
    Highlight        color.Color
    Skeleton1        color.Color
    Skeleton2        color.Color
    VisitedLink      color.Color
    Disabled1        color.Color
    Disabled2        color.Color
    Disabled3        color.Color
}

func Gray100Theme() *Theme {
    return &Theme{
        UiBackground:     Gray100,
        Interactive1:     Blue60,
        Interactive2:     Gray60,
        Interactive3:     White,
        Interactive4:     Blue50,
        Danger:           Red60,
        Ui1:              Gray90,
        Ui2:              Gray80,
        Ui3:              Gray80,
        Ui4:              Gray60,
        Ui5:              Gray10,
        ButtonSeparator:  Gray100,
        Decorative:       Gray70,
        Text1:            Gray10,
        Text2:            Gray30,
        Text3:            Gray60,
        Text4:            White,
        Text5:            Gray50,
        TextError:        Red40,
        Link1:            Blue40,
        InverseLink:      Blue60,
        Icon1:            Gray10,
        Icon2:            Gray30,
        Icon3:            White,
        Field1:           Gray90,
        Field2:           Gray80,
        Inverse1:         Gray100,
        Inverse2:         Gray10,
        Support1:         Red50,
        Support2:         Green40,
        Support3:         Yellow30,
        Support4:         Blue50,
        InverseSupport1:  Red60,
        InverseSupport2:  Green50,
        InverseSupport3:  Yellow30,
        InverseSupport4:  Blue70,
        Overlay1:         color.RGBA{Gray100.R, Gray100.G, Gray100.B, 179},
        Focus:            White,
        InverseFocusUi:   Blue60,
        HoverPrimary:     Blue60Hover,
        HoverPrimaryText: Blue30,
        HoverSecondary:   Gray60Hover,
        HoverTertiary:    Gray10,
        HoverUi:          Gray90Hover,
        HoverLightUi:     Gray80Hover,
        HoverSelectedUi:  Gray80Hover,
        HoverDanger:      Red60Hover,
        HoverRow:         Gray90Hover,
        InverseHoverUi:   Gray10Hover,
        ActivePrimary:    Blue80,
        ActiveSecondary:  Gray80,
        ActiveTertiary:   Gray30,
        ActiveUi:         Gray70,
        ActiveLightUi:    Gray60,
        ActiveDanger:     Red80,
        SelectedUi:       Gray80,
        SelectedLightUi:  Gray70,
        Highlight:        Blue90,
        Skeleton1:        Gray80Hover,
        Skeleton2:        Gray80,
        VisitedLink:      Purple40,
        Disabled1:        Gray90,
        Disabled2:        Gray70,
        Disabled3:        Gray60,
    }
}
