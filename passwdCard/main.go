package main

import (
	"fmt"
	"github.com/ActiveState/golor"
	"math/rand"
	"strings"
)

type PasswordCard struct {
	number  int64
	symbols bool
	digits  bool
	grid    [HEIGHT][WIDTH]string
}

const (
	HEADER_CHARS               = "■□▲△○●★☂☀☁☹☺♠♣♥♦♫€¥£$!?¡¿⊙◐◩�"
	DIGITS                     = "0123456789"
	DIGITS_AND_LETTERS         = "0123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	DIGITS_LETTERS_AND_SYMBOLS = "0123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ@#$%&*<>?€+{}[]()/\\"
	WIDTH                      = 29
	HEIGHT                     = 9
)

func (p *PasswordCard) random() func(int) int {
	rand.Seed(p.number)
	return func(x int) int {
		return rand.Intn(x)
	}
}

func (p *PasswordCard) getGrid() [HEIGHT][WIDTH]string {
	var empty_grid = new([HEIGHT][WIDTH]string)
	// new pointer
	// TODO daha iyi çözüm var mı?
	if p.grid == *empty_grid {
		p.genereateGrid()
	}
	return p.grid
}

func (p *PasswordCard) genereateGrid() {
	rnd := p.random()
	headerCharsArray := strings.Split(HEADER_CHARS, "")
	p.shuffle(headerCharsArray, rnd)

	headerChars := new([WIDTH]string)
	copy(headerChars[:], headerCharsArray[:])

	for i, v := range headerChars {
		p.grid[0][i] = v
	}

	if p.digits {
		halfHeight := 1 + ((HEIGHT - 1) / 2)
		for y := 1; y < halfHeight; y++ {
			for x := 0; x < WIDTH; x++ {
				if p.symbols && (x%2) == 0 {
					p.grid[y][x] = string(DIGITS_LETTERS_AND_SYMBOLS[rnd(len(DIGITS_LETTERS_AND_SYMBOLS))])
				} else {
					p.grid[y][x] = string(DIGITS_AND_LETTERS[rnd(len(DIGITS_AND_LETTERS))])
				}
			}
		}
		for y := halfHeight; y < HEIGHT; y++ {
			for x := 0; x < WIDTH; x++ {
				p.grid[y][x] = string(DIGITS[rnd(10)])
			}
		}
	} else {
		for y := 1; y < HEIGHT; y++ {
			for x := 0; x < WIDTH; x++ {
				if p.symbols && (x%2) == 0 {
					p.grid[y][x] = string(DIGITS_LETTERS_AND_SYMBOLS[rnd(len(DIGITS_LETTERS_AND_SYMBOLS))])
				} else {
					p.grid[y][x] = string(DIGITS_AND_LETTERS[rnd(len(DIGITS_AND_LETTERS))])

				}
			}
		}
	}
}

func (p *PasswordCard) shuffle(list []string, random func(int) int) {
	for i := len(list); i > 1; i-- {
		a := random(i)
		b := i - 1
		list[b], list[a] = list[a], list[b]
	}
}

func (p *PasswordCard) printGrid() {
	colors := []int{
		golor.WHITE,
		golor.WHITE,
		golor.GRAY,
		golor.RED,
		golor.GREEN,
		golor.YELLOW,
		golor.BLUE,
		golor.MAGENTA,
		golor.CYAN,
	}

	fmt.Printf("\nCard id: %d\n", p.number)
	for i, r := range p.getGrid() {
		ro := strings.Join(r[:], "")
		fmt.Println(golor.Colorize(ro, -1, colors[i%len(colors)]))
	}
}

func main() {
	var pds = PasswordCard{number: 19, digits: true, symbols: true}
	pds.printGrid()
	var pd = PasswordCard{number: 19, digits: true}
	pd.printGrid()
	var ps = PasswordCard{number: 19, symbols: true}
	ps.printGrid()
	var p = PasswordCard{number: 19}
	p.printGrid()
}
