package main

import (
	"image/color"
	"image"
	"os"
	"image/draw"
	"image/png"


  	"github.com/vova616/screenshot"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"


)

var(


	COLOR color.Color

)

func savefile(canvas *pixelgl.Canvas, pos pixel.Vec){


	img, _ := screenshot.CaptureScreen() // zrobienie screenshota do img
	screen := image.Image(img) // stworzenie obiektu typu Image z screenshota

  winScreen := image.NewRGBA(image.Rect(0, 0, 1024, 768)) // stworzenie obrazek, w który wrysowujesz kawałek screenshota zawierający okno programu
/*
  winscreen -> w co ma wrysowywac
  pozycja w której rysujemy fragment w obrazku docelowym , 0,0 to pozycja lewy dół a 1024 i 768 to wielkość pola rysowania
	screen -> źródło
	następny argument odpowiada za ustalenie punktu z którego zaczynamy pobierać dane ( współrzędne punktu pobieramy z getpos)
	działa?
*/
	draw.Draw(winScreen, image.Rect(0, 0, 1024, 768), screen, image.Pt(int(pos.X), int(pos.Y)), draw.Src)

	// otwarcie lub stworzenie (jeśli nie istnieje) pliku
	f, _:= os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0655)
	// co jest po defer zostanie wykonane po zakończeniu metody w której defer został użyty
	defer f.Close() // zamknięcie pliku
	png.Encode(f, winScreen) // do pliku f zapisuje dane z obrazka winscreen

}

func run(){
	// tworzymy okno
	cfg := pixelgl.WindowConfig{
		Title:     "Paint_go",
		// rozmiar
		Bounds:    pixel.R(0, 0, 1024, 768),
		// użytkownik może zmieniać rozmiar okna
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// nadanie koloru okna nie "płótna"
	win.Clear(colornames.White)

  imd := imdraw.New(nil)

  COLOR = colornames.Black

  // stworzenie samego płótna
	canvas := pixelgl.NewCanvas(win.Bounds())

  // główna pętla programu
	for !win.Closed(){

		if win.JustPressed(pixelgl.Key1){
			COLOR = colornames.Blue

   		 }

		if win.JustPressed(pixelgl.Key2){
			COLOR = colornames.Green

		}

    		if win.JustPressed(pixelgl.Key3){
     			 COLOR = colornames.Violet

    		}

   		 if win.JustPressed(pixelgl.Key4){
    			  COLOR = colornames.Yellow

    		}

    		if win.JustPressed(pixelgl.Key5){
   			   COLOR = colornames.Red

   		 }

		if win.JustPressed(pixelgl.KeyS){
      			savefile(canvas, win.GetPos())

		}

		// pętla kiedy jest nacisniety przycisk myszy
		if win.Pressed(pixelgl.MouseButtonLeft){
			// pozycja kursora w samym oknie
			pos := win.MousePosition()
			// wyczyszczenie parametrów rysowania przed samym narysowaniem
			imd.Clear()
     			 // kolor pędzla
			imd.Color = COLOR
      			//pozycja kursora
			imd.Push(pixel.V(pos.X, pos.Y))
			//rysowanie samego pedzla wraz z rozmiarem -> 5
			imd.Circle(5,0)
		}

		// rysowanie "płótna" na oknie
		imd.Draw(canvas)
		canvas.Draw(win, pixel.IM.Moved(canvas.Bounds().Center()))
  		win.Update()
		}

}


func main() {
	pixelgl.Run(run)
	}
