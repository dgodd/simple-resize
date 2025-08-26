package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // Import this to register the JPEG decoder
	"image/png"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sunset-beach.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Define the image dimensions
	outWidth := 20
	outHeight := int(float64(height) / float64(width) * float64(outWidth))
	stepWidth, stepHeight := width/outWidth, height/outHeight
	fmt.Println("Original image dimensions:", width, "x", height)
	fmt.Println("Output image dimensions:", outWidth, "x", outHeight)
	fmt.Println("Steps:", stepWidth, "x", stepHeight)

	// Create a new RGBA image (supports transparency)
	// The image.Rect defines the bounds of the image
	outImg := image.NewRGBA(image.Rect(0, 0, outWidth, outHeight))

	for y := range outHeight {
		for x := range outWidth {
			c := img.At(x*stepWidth, y*stepHeight)
			r, g, b, _ := c.RGBA() // RGBA values are in the range [0, 65535]
			// Process pixel data (e.g., convert to 0-255 range by dividing by 257)
			r8, g8, b8 := uint8(r/257), uint8(g/257), uint8(b/257)

			// TODO: Average of the area
			avgRGB := color.RGBA{r8, g8, b8, 255}
			outImg.Set(x, y, avgRGB)
		}
	}

	// Set individual pixel colors (e.g., a green dot)
	green := color.RGBA{0, 255, 0, 255}
	outImg.Set(4, 4, green)

	// Create a file to save the image
	file, err = os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Encode and save the image as PNG
	err = png.Encode(file, outImg)
	if err != nil {
		log.Fatal(err)
	}
}
