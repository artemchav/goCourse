package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	height = 500
	width  = 700
)

func main() {
	bgColor := color.RGBA{100, 155, 100, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	drawVerticalLine(rectImg, 150, height, red)
	drawVerticalLine(rectImg, 200, height, black)
	drawHorizontalLine(rectImg, 200, width, blue)
	drawHorizontalLine(rectImg, 400, width, white)

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}

func drawVerticalLine(img *image.RGBA, x, height int, color color.RGBA) {
	for i := 0; i < height; i++ {
		img.Set(x, i, color)
	}
}

func drawHorizontalLine(img *image.RGBA, y, width int, color color.RGBA) {
	for i := 0; i < width; i++ {
		img.Set(i, y, color)
	}
}
