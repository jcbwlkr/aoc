package main

import (
	"fmt"

	"github.com/cheggaaa/pb"
)

func main() {
	var (
		b         = NewBoard()
		commands  = ParseCommands("input.txt")
		doMakeGif = true
		g1        = NewGIF("lights_1.gif")
		g2        = NewGIF("lights_2.gif")
	)

	commands = commands[0:30]
	fmt.Println("Processing commands")
	bar := pb.StartNew(len(commands))
	for _, cmd := range commands {
		for _, p := range cmd.Range() {
			b[p].TakeAction(cmd.Action)
		}

		if doMakeGif {
			g1.AddImage(b.ImageBW())
			g2.AddImage(b.ImageColor())
		}
		bar.Increment()
	}

	bar.FinishPrint("Wrapping up")

	if doMakeGif {
		g1.Encode()
		g2.Encode()
	}

	fmt.Println("Lights on at the end", b.LitCount()) // 569999 for me
	fmt.Println("Total Brightness", b.Brightness())   // 17836115 for me
}
