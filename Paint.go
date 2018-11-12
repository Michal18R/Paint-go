package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"


	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	// poniżej parametry okienka painta
	Width  = 500
	Height = 400
	// zmienne pod przypisanie przeźroczystości
	c10 = 0x40
	c11 = 0xc0
	c12 = 0xff
)

// przypisywanie zmiennych
var (
	brush      *ebiten.Image
	whiteImage *ebiten.Image
	COLOR      [4]float64
        
	// przypisanie do zmiennych dany kolor (1.0 to inaczej 255 w skali rgb)
	WHITE  = [4]float64{1.0, 1.0, 1.0, 1.0}
	BLACK  = [4]float64{0, 0, 0, 1.0}
	BLUE   = [4]float64{0, 0, 1.0, 1.0}
	RED    = [4]float64{1.0, 0, 0, 1.0}
	GREEN  = [4]float64{0, 1.0, 0, 1.0}
	YELLOW = [4]float64{1.0, 1.0, 0, 1.0}
	PURPLE = [4]float64{1.0, 0, 1.0, 1.0}
)

func init() {
	// tworzymy wyglad pędzla
	pixels := []uint8{

		c10, c11, c10,
		c11, c12, c11,
		c10, c11, c10,
	}
	// rysuje faktycznie pędzel
	brush, _ = ebiten.NewImageFromImage(&image.Alpha{
		Pix:    pixels,
		Stride: 3,
		Rect:   image.Rect(0, 0, 3, 3),
	}, ebiten.FilterDefault)

	// _ ignoruje argument (drugi argument jest ewentualny błąd)
	//tworzymy biały prostokąt na screenie
	whiteImage, _ = ebiten.NewImage(Width, Height, ebiten.FilterDefault)
	//wypełniamy prostokąt kolorem
	whiteImage.Fill(color.White)
	COLOR = WHITE
}

//narysowanie nowych kropek pędzlem na whiteImage
func draw(canvas *ebiten.Image, x, y int) {
	//wyciągamy paramety rysowania
	op := &ebiten.DrawImageOptions{}
	//ustalamy pozycję rysowania
	op.GeoM.Translate(float64(x), float64(y))
	//ustawiamy kolor rysowania
	op.ColorM.Scale(COLOR[0], COLOR[1], COLOR[2], COLOR[3])
	//faktyczne rysowanie kropki na whiteImage
	canvas.DrawImage(brush, op)
}

//sprawdzamy który przycisk jest nacisniety i wraz z nim przypisujemy działanie
func button(screen *ebiten.Image) {
	if ebiten.IsKeyPressed(42) && ebiten.IsKeyPressed(28) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("save"))
//		err := image.Encode(w, whiteImage)
	}

	// przypisujemy kolory do posczegolnych przyciskow
	if ebiten.IsKeyPressed(1) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("BLACK"))
		COLOR = BLACK
	}
	if ebiten.IsKeyPressed(2) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("BLUE"))
		COLOR = BLUE
	}
	if ebiten.IsKeyPressed(3) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("RED"))
		COLOR = RED
	}
	if ebiten.IsKeyPressed(4) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("GREEN"))
		COLOR = GREEN
	}
	if ebiten.IsKeyPressed(5) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("YELLOW"))
		COLOR = YELLOW
	}
	if ebiten.IsKeyPressed(6) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("PURPLE"))
		COLOR = PURPLE
	}
}

func update(screen *ebiten.Image) error {

	// pozycja kursora wykorzystująca bibliotekę ebiten
	mousex, mousey := ebiten.CursorPosition()

	//sprawdzamy czy przycisk lewej myszki jest naciskany jak tak to rysuje
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		draw(whiteImage, mousex, mousey)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// wyświetlanie białęgo prostokątu
	screen.DrawImage(whiteImage, nil)

	button(screen)


	return nil
}

func main() {
	if err := ebiten.Run(update, Width, Height, 2, "Paint w go"); err != nil {
		log.Fatal(err)
	}
}
