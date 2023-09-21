package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/eiannone/keyboard"
	"github.com/nsf/termbox-go"
)

type screen struct {
	xMax       int
	yMax       int
	aX         int
	aY         int
	bX         int
	bY         int
	deltaBAX   int
	deltaBAY   int
	a          float64 // y=a*x+b
	b          int     // y=a*x+b
	y          int     // y=a*x+b
	x          int     // y=a*x+b
	colorIndex uint8
	char       rune
	fontColor  int
}

func main() {
	a := 13
	var p *int
	p = &a
	b := *p
	//var b int = int(*p)
	fmt.Printf("adres: %p\n%d\n", p, b)
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press a key (press 'q' to quit):")
	wait4key()
	err = termbox.Init()
	if err != nil {
		log.Println("Terminal initailization errer", err)
		panic(err)
	}
	defer termbox.Close()

	var console screen
	console.xMax, console.yMax = termbox.Size()
	console.aX = console.xMax / 2
	console.aY = console.yMax / 2
	console.bX = 10
	console.bY = 20

	//termbox.SetCell(console.aX, console.aY, 'A', termbox.ColorRed, termbox.ColorDefault)
	//termbox.SetCell(console.bX, console.bY, 'B', termbox.ColorGreen, termbox.ColorDefault)
	for {

		console.aX = rand.Intn(console.xMax)
		console.aY = rand.Intn(console.yMax)
		console.bX = rand.Intn(console.xMax)
		console.bY = rand.Intn(console.yMax)
		console.fontColor = rand.Intn(15)
		console.draw()
		funkcja := fmt.Sprintf("y=%x*x+%d; A(%d,%d) and (%d,%d)", console.a, console.b, console.aX, console.aY, console.bX, console.bY)
		for i, v := range funkcja {
			termbox.SetCell(i, 1, v, termbox.ColorDefault, termbox.ColorDefault)
		}

		//termbox.SetCell(console.aX, console.aY, 'A', termbox.Attribute(console.fontColor), termbox.ColorDefault)
		//termbox.SetCell(console.bX, console.bY, 'B', termbox.Attribute(console.fontColor), termbox.ColorDefault)
		termbox.Flush()
		wait4key()
	}
}

func wait4key() {
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc || char == 'q' {
			// Exit the loop when 'q' or the Escape key is pressed
			break
		}
		break
	}
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func (s *screen) draw() {
	s.deltaBAX = s.bX - s.aX
	s.deltaBAY = s.bY - s.aY
	s.a = float64(s.deltaBAY) / float64(s.deltaBAX)
	s.b = s.aY - int(float64(s.aX)*s.a)
	if abs(s.deltaBAY) > abs(s.deltaBAX) {

		if s.aY > s.bY {
			for pointY := s.bY; pointY <= s.aY; pointY++ {
				pointX := int(float64(pointY-s.b) / s.a)
				termbox.SetCell(pointX, pointY, '█', termbox.Attribute(s.fontColor), termbox.ColorDefault)
			}

		} else {
			for pointY := s.aY; pointY <= s.bY; pointY++ {
				pointX := int(float64((pointY - s.b)) / s.a)
				termbox.SetCell(pointX, pointY, '█', termbox.Attribute(s.fontColor), termbox.ColorDefault)
			}

		}

	} else {
		if s.aX > s.bX {
			for pointX := s.bX; pointX <= s.aX; pointX++ {
				pointY := int(float64(pointX)*s.a) + s.b
				termbox.SetCell(pointX, pointY, '█', termbox.Attribute(s.fontColor), termbox.ColorDefault)
			}
		} else {
			for pointX := s.aX; pointX <= s.bX; pointX++ {
				pointY := int(float64(pointX)*s.a) + s.b
				termbox.SetCell(pointX, pointY, '█', termbox.Attribute(s.fontColor), termbox.ColorDefault)
			}
		}

	}

}
