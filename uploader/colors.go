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
	tshirt         string `default:"white"`
	hoodie         string `default:"black"`
	tank           string `default:"white"`
	crewneck       string `default:"white"`
	longsleeve     string `default:"white"`
	baseball       string `default:"black_white"`
	kids           string `default:"white"`
	kidsHoodie     string `default:"black"`
	kidsLongsleeve string `default:"black"`
	babyBody       string `default:"black"`
}

func (c *ColorChoices) setDefaults() {
	for i := 0; i < reflect.TypeOf(*c).NumField(); i++ {
		field := reflect.TypeOf(*c).Field(i)

		if value, ok := field.Tag.Lookup("default"); ok {
			switch field.Name {
			case "tshirt":
				if c.tshirt == "" { c.tshirt = value }
				break
			case "hoodie":
				if c.hoodie == "" { c.hoodie = value }
				break
			case "tank":
				if c.tank == "" { c.tank = value }
				break
			case "crewneck":
				if c.crewneck == "" { c.crewneck = value }
				break
			case "longsleeve":
				if c.longsleeve == "" { c.longsleeve = value }
				break
			case "baseball":
				if c.baseball == "" { c.baseball = value }
				break
			case "kids":
				if c.kids == "" { c.kids = value }
				break
			case "kidsHoodie":
				if c.kidsHoodie == "" { c.kidsHoodie = value }
				break
			case "kidsLongsleeve":
				if c.kidsLongsleeve == "" { c.kidsLongsleeve = value }
				break
			case "babyBody":
				if c.babyBody == "" { c.babyBody = value }
				break
			}
		}
	 }
}