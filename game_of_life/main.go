package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 200
	height = 200
)

type Universe struct {
	grid [][]bool
	rng  *rand.Rand
}

func NewUniverse(rng *rand.Rand) Universe {
	u := Universe{
		grid: make([][]bool, height),
		rng:  rng,
	}
	for i := range u.grid {
		u.grid[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.grid[u.rng.Intn(height)][u.rng.Intn(width)] = true
	}
}

func (u Universe) String() string {
	var b byte
	str := ""
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b = ' '
			if u.grid[i][j] {
				b = '*'
			}
			str += string(b)
		}
		str += "\n"
	}
	return str
}

func (u Universe) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i != 0 || j != 0) && u.grid[(x+i+height)%height][(y+j+width)%width] {
				alive++
			}
		}
	}
	return alive == 3 || alive == 2 && u.grid[x][y]
}

func (u Universe) Step() Universe {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	nu := NewUniverse(rng)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nu.grid[y][x] = u.Next(x, y)
		}
	}
	return nu
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	u := NewUniverse(rng)
	u.Seed()
	for i := 0; i < 100; i++ {
		fmt.Print("\x0c", u.String())
		u = u.Step()
		time.Sleep(time.Second)
	}
}
