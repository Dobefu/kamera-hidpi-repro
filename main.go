package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/setanarut/kamera/v2"
)

const (
	GAME_WIDTH  = 640
	GAME_HEIGHT = 360
)

type Game struct {
	camera kamera.Camera
	img    *ebiten.Image
}

func (g *Game) Init() {
	g.camera = *kamera.NewCamera(GAME_WIDTH/2, GAME_HEIGHT/2, GAME_WIDTH, GAME_HEIGHT)

	g.img = ebiten.NewImage(GAME_WIDTH, GAME_HEIGHT)
	g.img.Fill(color.Gray{Y: 100})
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.camera.Draw(g.img, &ebiten.DrawImageOptions{}, screen)
}

func (g *Game) Layout(
	outsideWidth int,
	outsideHeight int,
) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	screen.DrawImage(offscreen, nil)
}

func main() {
	ebiten.SetWindowSize(GAME_WIDTH, GAME_HEIGHT)

	game := &Game{}
	game.Init()

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}
