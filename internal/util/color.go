package util

import (
	"github.com/AvraamMavridis/randomcolor"
	"github.com/go-playground/colors"
)

func RandomColor() string {
	return randomcolor.GetRandomColorInHex()
}

func IsColorValid(color string) bool {
	_, err := colors.Parse(color)
	return err == nil
}

func MustParseColor(color string) string {
	c, err := colors.Parse(color)
	if err != nil {
		panic(err)
	}
	return c.ToHEX().String()
}

func MustColorToRGB(color string) string {
	c, err := colors.Parse(color)
	if err != nil {
		panic(err)
	}
	return c.ToRGB().String()
}
