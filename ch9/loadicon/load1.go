package load

import (
	"image"
	"sync"
)

//var mu sync.Mutex
var loadIconsOnce sync.Once

var icons map[string]image.Image

func loadIcons() {
	icons := make(map[string]image.Image)
	icons["1.png"] = loadIcon("1.png")
	icons["2.png"] = loadIcon("2.png")
}

//test
func Icon(name string) image.Image {
	loadIconsOnce.do(loadIcons)
	return icons[name]
}
