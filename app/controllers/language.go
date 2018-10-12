package controllers

import (
	"github.com/xDarkicex/easycare_server/helpers"
)

// Translations holds data struct for  multi langauge support
type Translations map[string]interface{}

func (t Translations) JSON(path string) map[string]interface{} {
	return helpers.ReadJSON(path, t)
}

func Translated(path string) Translations {
	T := Translations{}
	T.JSON(path)
	return T
}

// LoadLanguage function returns proper langage from langauge json
func LoadLanguage(language string) Translations {
	switch lang := language; lang {
	case "english":
		return Translated("english")
	case "spanish":
		return Translated("english")
	case "italian":
		return Translated("english")
	case "greek":
		return Translated("english")
	case "russian":
		return Translated("english")
	case "portugese":
		return Translated("english")
	case "french":
		return Translated("english")
	case "german":
		return Translated("english")
	default:
		return Translated("english")
	}
}
