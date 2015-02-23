package hqx

import "testing"
import "log"
import "os"
import "image/png"

const (
	sample_file = "sample.png"
	width       = 128
	height      = 64
)

func Test_hqx1(t *testing.T) {
	file, err := os.Create("identical.png")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	file1, err := os.Open(sample_file)
	if err != nil {
		log.Println(err)
	}
	defer file1.Close()

	img, _ := png.Decode(file1)
	data := img2data(img)
	newimg := data2img(data, width, height)

	png.Encode(file, newimg)
}

func Test_hqx2(t *testing.T) {
	file, err := os.Create("sample_2x.png")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	file1, err := os.Open(sample_file)
	if err != nil {
		log.Println(err)
	}
	defer file1.Close()

	img, _ := png.Decode(file1)
	newimg := Hq2x(img)

	png.Encode(file, newimg)
}

func Test_hqx3(t *testing.T) {
	file, err := os.Create("sample_3x.png")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	file1, err := os.Open(sample_file)
	if err != nil {
		log.Println(err)
	}
	defer file1.Close()

	img, _ := png.Decode(file1)
	newimg := Hq3x(img)

	png.Encode(file, newimg)
}

func Test_hqx4(t *testing.T) {
	file, err := os.Create("sample_4x.png")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	file1, err := os.Open(sample_file)
	if err != nil {
		log.Println(err)
	}
	defer file1.Close()

	img, _ := png.Decode(file1)
	newimg := Hq4x(img)

	png.Encode(file, newimg)
}
