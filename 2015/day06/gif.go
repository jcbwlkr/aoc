package main

import (
	"image"
	"image/gif"
	"log"
	"os"
)

// GIF embeds a gif.GIF and adds some useful fields and methods.
type GIF struct {
	Name string

	gif.GIF
}

// NewGIF creates a GIF with defaults.
func NewGIF(name string) *GIF {
	return &GIF{
		Name: name,
		GIF: gif.GIF{
			LoopCount: 1,
			Image:     []*image.Paletted{},
			Delay:     []int{},
		},
	}
}

// AddImage adds a frame to the GIF's animation.
func (g *GIF) AddImage(img *image.Paletted) {
	g.GIF.Image = append(g.GIF.Image, img)
	g.GIF.Delay = append(g.GIF.Delay, 2) // 2 100ths of a second
}

// Encode encodes the GIF as an actual gif image and writes it to disk.
func (g *GIF) Encode() {
	file, err := os.Create(g.Name)
	if err != nil {
		log.Fatalln(err)
	}

	if err := gif.EncodeAll(file, &g.GIF); err != nil {
		log.Fatal(err)
	}
}
