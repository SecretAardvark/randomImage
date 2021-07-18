package main

//TODO: Minimum POC: generate one random image on request of one type.

import (
	"fmt"
	"image/color"
	"math/rand"
	"net/http"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./html")))
	http.HandleFunc("/image", serveImage)
	http.ListenAndServe(":3000", nil)
	fmt.Println("Starting a server on port :3000")

}

func serveImage(w http.ResponseWriter, r *http.Request) {
	w.Write(rshapesImage())
}

func rshapesImage() []byte {
	rand.Seed(time.Now().Unix())
	canvas := generativeart.NewCanva(500, 500)
	canvas.SetBackground(common.Mintcream)
	canvas.FillBackground()
	canvas.SetColorSchema([]color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	})
	canvas.Draw(arts.NewRandomShape(150))
	image, _ := canvas.ToBytes()
	return image
}
