package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Possible actions
const (
	TurnOn  = "turn on"
	TurnOff = "turn off"
	Toggle  = "toggle"
)

type command struct {
	Action         string
	StartX, StartY int
	EndX, EndY     int
}

func main() {
	var (
		board     = makeBoard()
		commands  = getCommands()
		images    = make([]*image.Paletted, len(commands))
		r         = image.Rect(0, 0, 1000, 1000)
		doMakeGif = true
	)

	fmt.Print("Generating frames from commands")
	for i, cmd := range commands {
		fmt.Print(".")

		for x := cmd.StartX; x <= cmd.EndX; x++ {
			for y := cmd.StartY; y <= cmd.EndY; y++ {
				switch cmd.Action {
				case TurnOn:
					board[x][y] = true
				case TurnOff:
					board[x][y] = false
				case Toggle:
					board[x][y] = !board[x][y]
				}
			}
		}

		if doMakeGif {
			img := image.NewPaletted(r, palette.Plan9)
			for x, col := range board {
				for y, val := range col {
					if val {
						img.Set(x, y, color.White)
					}
				}
			}
			images[i] = img
		}
	}
	fmt.Print("\n")

	if doMakeGif {
		// Create the delay times for each frame. Use 2 100ths of a second per frame.
		times := make([]int, len(images))
		for i := range times {
			times[i] = 2
		}

		g := gif.GIF{
			Image:     images,
			Delay:     times,
			LoopCount: 1,
		}

		outfile, err := os.Create("lights.gif")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Encoding GIF")
		if err := gif.EncodeAll(outfile, &g); err != nil {
			log.Fatal(err)
		}

		fmt.Println("GIF created")
	}

	count := 0
	for _, col := range board {
		for _, val := range col {
			if val {
				count++
			}
		}
	}
	fmt.Println("Lights on at the end", count)
}

func makeBoard() [][]bool {
	board := make([][]bool, 1000)
	for i := range board {
		board[i] = make([]bool, 1000)
	}

	return board
}

func getCommands() []command {
	re := regexp.MustCompile("(turn on|turn off|toggle) ([0-9]*),([0-9]*) through ([0-9]*),([0-9]*)")

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var commands []command
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		cmd := scanner.Text()
		m := re.FindStringSubmatch(cmd)
		commands = append(commands, command{
			Action: m[1],
			StartX: mustAtoi(m[2]),
			StartY: mustAtoi(m[3]),
			EndX:   mustAtoi(m[4]),
			EndY:   mustAtoi(m[5]),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return commands
}

func mustAtoi(a string) (i int) {
	var err error
	i, err = strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return i
}
