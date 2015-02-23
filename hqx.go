package hqx

// #include "hqx.h"
import "C"

import (
	"image"
)

func Hq2x(img image.Image) image.Image {
	return hqx(img, 2)
}

func Hq3x(img image.Image) image.Image {
	return hqx(img, 3)
}

func Hq4x(img image.Image) image.Image {
	return hqx(img, 4)
}

func hqx(img image.Image, scaleBy int) image.Image {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	srcData := img2data(img)
	destData := make([]uint32, width*scaleBy*height*scaleBy)
	sp := (*C.uint32_t)(&srcData[0])
	dp := (*C.uint32_t)(&destData[0])
	c_w := C.int(width)
	c_h := C.int(height)

	C.hqxInit()
	switch scaleBy {
	case 2:
		C.hq2x_32(sp, dp, c_w, c_h)
	case 3:
		C.hq3x_32(sp, dp, c_w, c_h)
	case 4:
		C.hq4x_32(sp, dp, c_w, c_h)
	}
	return data2img(destData, width*scaleBy, height*scaleBy)
}

func img2data(img image.Image) []uint32 {
	bounds := img.Bounds()
	data := make([]uint32, bounds.Max.X*bounds.Max.Y)
	pixel_idx := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// RGBA -> BGRA
			data[pixel_idx] = ((r >> 8) << 8) | ((g >> 8) << 16) | ((b >> 8) << 24) | (a >> 8)
			pixel_idx++
		}
	}
	return data
}

func data2img(data []uint32, width, height int) image.Image {
	imgout := image.NewRGBA(image.Rect(0, 0, width, height))

	rgb := [4]uint32{}

	pixel_idx := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel_raw := data[pixel_idx]
			// BRGA -> RGBA
			rgb[2] = (pixel_raw & 0xff000000) >> 24 //blue
			rgb[1] = (pixel_raw & 0x00ff0000) >> 16 //green
			rgb[0] = (pixel_raw & 0x0000ff00) >> 8  //red
			rgb[3] = pixel_raw & 0x000000ff         //alpha
			for i, c := range rgb {
				imgout.Pix[pixel_idx*4+i] = uint8(c)
			}
			pixel_idx++
		}
	}
	return imgout
}