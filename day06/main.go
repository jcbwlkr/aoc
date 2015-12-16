package main

import "fmt"

func main() {
	var (
		b         = NewBoard()
		commands  = ParseCommands("input.txt")
		doMakeGif = false
		g1        = NewGIF("lights_1.gif")
		//g2        = NewGIF("lights_2.gif")
	)

	for _, cmd := range commands {
		for _, p := range cmd.Range() {
			b[p].TakeAction(cmd.Action)
		}

		if doMakeGif {
			g1.AddImage(b.Image())
		}
	}

	if doMakeGif {
		g1.Encode()
	}

	fmt.Println("Lights on at the end", b.LitCount()) // 569999 for me
	fmt.Println("Total Brightness", b.Brightness())   // 17836115 for me
}
