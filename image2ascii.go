package main

import "github.com/qeesung/image2ascii/convert"

func ImageToAscii(imgPath string, w int, h int) string {
	convertOptions := convert.DefaultOptions
	convertOptions.Ratio = 5
	convertOptions.FitScreen = true
	convertOptions.FixedWidth = w
	convertOptions.FixedHeight = h
	convertOptions.Colored = false

	converter := convert.NewImageConverter()
	return converter.ImageFile2ASCIIString(imgPath, &convertOptions)
}
