package carbon

import (
    "sync"
    
    "github.com/golang/freetype/truetype"
)

// FontLoader is a function loading the font
type FontLoader func() *truetype.Font

// Load just calls the FontLoader
func (l FontLoader) Load() *truetype.Font {
    return l()
}

func parseFont(bytes []byte, err error) (ttf *truetype.Font) {
    if err != nil {
        panic(err)
    }
    ttf, err = truetype.Parse(bytes)
    if err != nil {
        panic(err)
    }
    return
}

// ParseFontLater takes the font reader and returns the FontLoader function,
// that parses the font only at its first call.
func ParseFontLater(reader func() ([]byte, error)) FontLoader {
    var (
        once sync.Once
        ttf  *truetype.Font
    )
    return func() *truetype.Font {
        once.Do(func() {
            ttf = parseFont(reader())
            reader = nil
        })
        return ttf
    }
}

var (
    // Built in regular font, parsed using bindata
    PlexSansRegular = ParseFontLater(ibmplexsansRegularTtfBytes)
    // Built in bold font, parsed using bindata
    PlexSansSemiBold = ParseFontLater(ibmplexsansSemiboldTtfBytes)
    // Built in light font, parsed using bindata
    PlexSansLight = ParseFontLater(ibmplexsansLightTtfBytes)
    // Built in monospace font, parsed using bindata
    PlexMonoRegular = ParseFontLater(ibmplexmonoRegularTtfBytes)
)
