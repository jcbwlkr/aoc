package main

import (
	"fmt"

	"github.com/cheggaaa/pb"
)

func main() {
	var (
		b        = NewBoard()
		commands = ParseCommands("input.txt")
		g1       = NewGIF("lights_1.gif")
		g2       = NewGIF("lights_2.gif")
	)

	fmt.Println("Processing commands")
	bar := pb.StartNew(len(commands))
	for _, cmd := range commands {
		for _, p := range cmd.Range() {
			b[p].TakeAction(cmd.Action)
		}

		g1.AddImage(b.ImageBW())
		g2.AddImage(b.ImageColor())
		bar.Increment()
	}

	bar.FinishPrint("Wrapping up")

	g1.Encode()
	g2.Encode()

	fmt.Println("Lights on at the end", b.LitCount()) // 569999 for me
	fmt.Println("Total Brightness", b.Brightness())   // 17836115 for me
}
