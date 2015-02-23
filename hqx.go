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
	width := bounds.Max.X
	height := bounds.Max.Y

	data := make([]uint32, width*height)

	{
		idx := 0
		for y := bounds.Min.Y; y < height; y++ {
			for x := bounds.Min.X; x < width; x++ {
				r, g, b, a := img.At(x, y).RGBA()

				//[0, 0xffff] -> [0, 0xff]
				r = narrowcolor(r)
				g = narrowcolor(g)
				b = narrowcolor(b)
				a = narrowcolor(a)

				// RGBA -> BGRA
				data[idx] = (r << 8) | (g << 16) | (b << 24) | a
				idx++
			}
		}
	}
	return data
}

func data2img(data []uint32, width, height int) image.Image {

	imgout := image.NewRGBA(image.Rect(0, 0, width, height))

	rgb := [4]uint32{}

	{
		idx := 0
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pixel_raw := data[idx]
				// BGRA -> RGBA
				rgb[2] = (pixel_raw & 0xff000000) >> 24 //blue
				rgb[1] = (pixel_raw & 0x00ff0000) >> 16 //green
				rgb[0] = (pixel_raw & 0x0000ff00) >> 8  //red
				rgb[3] = pixel_raw & 0x000000ff         //alpha

				for i, c := range rgb {
					imgout.Pix[idx*4+i] = uint8(c)
				}
				idx++
			}
		}
	}
	return imgout
}

func narrowcolor(c uint32) uint32 {
	if c >= 0xff00 {
		return c >> 8
	} else {
		return ((c & 0xff00) + (((c & 0x00ff) >> 7) << 8)) >> 8
	}
	return 0
}
