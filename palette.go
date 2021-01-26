package carbon

import (
    "fmt"
    "image/color"
)

var (
    Black       = color.RGBA{A: 255}
    White       = color.RGBA{R: 255, G: 255, B: 255, A: 255}
    Transparent = color.RGBA{}
    ShadowColor = color.RGBA{A: 26}
    
    Red100 = hex("2d0709")
    Red90  = hex("520408")
    Red80  = hex("750e13")
    Red70  = hex("a2191f")
    Red60  = hex("da1e28")
    Red50  = hex("fa4d56")
    Red40  = hex("ff8389")
    Red30  = hex("ffb3b8")
    Red20  = hex("ffd7d9")
    Red10  = hex("fff1f1")
    
    Orange100 = hex("231000")
    Orange90  = hex("3e1a00")
    Orange80  = hex("5e2900")
    Orange70  = hex("8a3800")
    Orange60  = hex("ba4e00")
    Orange50  = hex("eb6200")
    Orange40  = hex("ff832b")
    Orange30  = hex("ffb784")
    Orange20  = hex("ffd9be")
    Orange10  = hex("ffd9be")
    
    Yellow100 = hex("1c1500")
    Yellow90  = hex("302400")
    Yellow80  = hex("483700")
    Yellow70  = hex("684e00")
    Yellow60  = hex("8e6a00")
    Yellow50  = hex("b28600")
    Yellow40  = hex("d2a106")
    Yellow30  = hex("f1c21b")
    Yellow20  = hex("fddc69")
    Yellow10  = hex("fcf4d6")
    
    Green100 = hex("071908")
    Green90  = hex("022d0d")
    Green80  = hex("044317")
    Green70  = hex("0e6027")
    Green60  = hex("198038")
    Green50  = hex("24a148")
    Green40  = hex("42be65")
    Green30  = hex("6fdc8c")
    Green20  = hex("a7f0ba")
    Green10  = hex("defbe6")
    
    Teal100 = hex("081a1c")
    Teal90  = hex("022b30")
    Teal80  = hex("004144")
    Teal70  = hex("005d5d")
    Teal60  = hex("007d79")
    Teal50  = hex("009d9a")
    Teal40  = hex("08bdba")
    Teal30  = hex("3ddbd9")
    Teal20  = hex("9ef0f0")
    Teal10  = hex("d9fbfb")
    
    Cyan100 = hex("061727")
    Cyan90  = hex("012749")
    Cyan80  = hex("003a6d")
    Cyan70  = hex("00539a")
    Cyan60  = hex("0072c3")
    Cyan50  = hex("1192e8")
    Cyan40  = hex("33b1ff")
    Cyan30  = hex("82cfff")
    Cyan20  = hex("bae6ff")
    Cyan10  = hex("e5f6ff")
    
    Blue100 = hex("001141")
    Blue90  = hex("001d6c")
    Blue80  = hex("002d9c")
    Blue70  = hex("0043ce")
    Blue60  = hex("0f62fe")
    Blue50  = hex("4589ff")
    Blue40  = hex("78a9ff")
    Blue30  = hex("a6c8ff")
    Blue20  = hex("d0e2ff")
    Blue10  = hex("edf5ff")
    
    Purple100 = hex("1c0f30")
    Purple90  = hex("31135e")
    Purple80  = hex("491d8b")
    Purple70  = hex("6929c4")
    Purple60  = hex("8a3ffc")
    Purple50  = hex("a56eff")
    Purple40  = hex("be95ff")
    Purple30  = hex("d4bbff")
    Purple20  = hex("e8daff")
    Purple10  = hex("f6f2ff")
    
    Magenta100 = hex("2a0a18")
    Magenta90  = hex("510224")
    Magenta80  = hex("740937")
    Magenta70  = hex("9f1853")
    Magenta60  = hex("d12771")
    Magenta50  = hex("ee5396")
    Magenta40  = hex("ff7eb6")
    Magenta30  = hex("ffafd2")
    Magenta20  = hex("ffd6e8")
    Magenta10  = hex("fff0f7")
    
    Gray100 = hex("161616")
    Gray90  = hex("262626")
    Gray80  = hex("393939")
    Gray70  = hex("525252")
    Gray60  = hex("6f6f6f")
    Gray50  = hex("8d8d8d")
    Gray40  = hex("a8a8a8")
    Gray30  = hex("c6c6c6")
    Gray20  = hex("e0e0e0")
    Gray10  = hex("f4f4f4")
    
    // CoolGray100 = hex("121619")
    // CoolGray90  = hex("21272a")
    // CoolGray80  = hex("343a3f")
    // CoolGray70  = hex("4d5358")
    // CoolGray60  = hex("697077")
    
    Blue60Hover = hex("0353e9")
    Red60Hover  = hex("ba1b23")
    Gray90Hover = hex("353535")
    Gray80Hover = hex("4c4c4c")
    Gray60Hover = hex("606060")
    Gray10Hover = hex("e5e5e5")
)

func hex(hex string) color.RGBA {
    if hex[0] == '#' {
        hex = hex[1:]
    }
    if len(hex) != 6 {
        panic(fmt.Errorf("wrong color format: %s", hex))
    }
    var bytes [6]byte
    for i, r := range hex {
        b := byte(r)
        switch {
        case '0' <= b && b <= '9':
            bytes[i] = b - '0'
        case 'a' <= b && b <= 'f':
            bytes[i] = 10 + b - 'a'
        default:
            panic(fmt.Errorf("wrong color format: %s", hex))
        }
    }
    return color.RGBA{
        R: bytes[0]<<4 + bytes[1],
        G: bytes[2]<<4 + bytes[3],
        B: bytes[4]<<4 + bytes[5],
        A: 0xff,
    }
}
