package os

import (
	"strings"
)

type Platform int

const (
	Linux Platform = iota
	Windows
	Darwin
)

func (plat Platform) String() string {
	switch plat {
	case Linux:
		return "linux"
	case Windows:
		return "windows"
	case Darwin:
		return "darwin"
	default:
		panic("Platform not found")
	}
}

func PlatformFromString(plat string) Platform {
	plat = strings.ToLower(plat)
	switch plat {
	case "linux":
		return Linux
	case "windows":
		return Windows
	case "darwin":
		return Darwin
	default:
		panic("Platform not found")
	}
}
