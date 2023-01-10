package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"time"
)

type Game struct {
	backgroundColor color.RGBA
}

func (g *Game) Update() error {
	second := uint8(time.Now().Second())
	g.backgroundColor = color.RGBA{3 * second, 0, 0, 0xff}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.backgroundColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 400
}

func NewGame() (*Game, error) {
	return &Game{}, nil
}

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Game1")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
