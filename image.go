package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("40.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file.Name())
	img, kind, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(img.Bounds(),kind)

	fmt.Fprintln(os.Stderr, "input format: ", kind)
	file2, err := os.Create("50.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	jpeg.Encode(file2, img, &jpeg.Options{100})
	file2.Close()

}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
}
