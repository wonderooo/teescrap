package uploader

import (
	"reflect"
)

var colorMapping = map[string]string{
	"white":       "12",
	"black_white": "48",
	"black":       "1",
}

type ColorChoices struct {
	Tshirt         string `default:"white" json:"tshirt"`
	Hoodie         string `default:"black" json:"hoodie"`
	Tank           string `default:"white" json:"tank"`
	Crewneck       string `default:"white" json:"crewneck"`
	Longsleeve     string `default:"white" json:"longsleeve"`
	Baseball       string `default:"black_white" json:"baseball"`
	Kids           string `default:"white" json:"kids"`
	KidsHoodie     string `default:"black" json:"kidsHoodie"`
	KidsLongsleeve string `default:"black" json:"kidsLongsleeve"`
	BabyBody       string `default:"black" json:"babyBody"`
}

func (c *ColorChoices) setDefaults() {
	for i := 0; i < reflect.TypeOf(*c).NumField(); i++ {
		field := reflect.TypeOf(*c).Field(i)

		if value, ok := field.Tag.Lookup("default"); ok {
			switch field.Name {
			case "Tshirt":
				if c.Tshirt == "" {
					c.Tshirt = value
				}
				break
			case "Hoodie":
				if c.Hoodie == "" {
					c.Hoodie = value
				}
				break
			case "Tank":
				if c.Tank == "" {
					c.Tank = value
				}
				break
			case "Crewneck":
				if c.Crewneck == "" {
					c.Crewneck = value
				}
				break
			case "Longsleeve":
				if c.Longsleeve == "" {
					c.Longsleeve = value
				}
				break
			case "Baseball":
				if c.Baseball == "" {
					c.Baseball = value
				}
				break
			case "Kids":
				if c.Kids == "" {
					c.Kids = value
				}
				break
			case "KidsHoodie":
				if c.KidsHoodie == "" {
					c.KidsHoodie = value
				}
				break
			case "KidsLongsleeve":
				if c.KidsLongsleeve == "" {
					c.KidsLongsleeve = value
				}
				break
			case "BabyBody":
				if c.BabyBody == "" {
					c.BabyBody = value
				}
				break
			}
		}
	}
}
