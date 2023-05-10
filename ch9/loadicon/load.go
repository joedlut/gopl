package main

import (
	"image"
	"sync"
)

var mu sync.Mutex

var icons map[string]image.Image

func loadIcons() {
	icons := make(map[string]image.Image)
	icons["1.png"] = loadIcon("1.png")
	icons["2.png"] = loadIcon("2.png")
}

func Icon(name string) image.Image {
	mu.Lock()
	defer mu.UnLock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}
