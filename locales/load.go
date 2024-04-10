package locales

import (
	"embed"
	"fmt"

	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"
)

//go:embed po
var PoFiles embed.FS

//go:embed po/en.po
var EnglishPo []byte

const readError string = `
A language is not available/does not exist in the configuration.
The available ones are:
- Spanish
- English
`

var Languages = []string{"Español", "English"}

func read(file string) []byte {
	data, err := PoFiles.ReadFile(file)
	if err != nil {
		err = zenity.Error(readError)
		if err != nil {
			fmt.Println(err)
		}
		return EnglishPo
	}
	return data
}

func GetPo(lang string) *gotext.Po {
	po := gotext.NewPo()
	po.Parse(GetParser(lang))
	return po
}

func GetParser(lang string) []byte {
	return read("po/" + lang + ".po")
}
