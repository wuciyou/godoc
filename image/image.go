package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	// "image/png"
	"bufio"
	"net/http"
	"os"
)

func main() {
	fmt.Println("wuciyou")
	// f1()
	// f2()
	// f3()
	f4()
	// f5()
}

func f5() {
	f, err := os.Open("./testdata/code.jpg")
	defer f.Close()
	checkErr(err)
	f1, err := os.Create("./testdata/code/code.jpg")
	defer f1.Close()
	checkErr(err)

	img, err := jpeg.Decode(f)
	checkErr(err)
	var minPoint, maxPoint image.Point
	minPoint, maxPoint = img.Bounds().Min, img.Bounds().Max
	rgba := image.NewRGBA(image.Rectangle{Min: minPoint, Max: maxPoint})

	// var cut_rectangle image.Rectangle

	for x := minPoint.X; x <= maxPoint.X; x++ {
		for y := minPoint.Y; y <= maxPoint.Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			fmt.Print(r)
			fmt.Print(" ")
			rgba.Set(x, y, img.At(x, y))
		}
		fmt.Println("")
	}
	jpeg.Encode(f1, rgba, &jpeg.Options{100})
	fmt.Println("ok")
}

func f4() {
	f, err := os.Open("./testdata/code.jpg")
	defer f.Close()
	checkErr(err)
	f1, err := os.Create("./testdata/code/code.jpg")
	defer f1.Close()
	checkErr(err)

	img, err := jpeg.Decode(f)
	checkErr(err)
	var minPoint, maxPoint image.Point
	minPoint, maxPoint = img.Bounds().Min, img.Bounds().Max
	rgba := image.NewRGBA(image.Rectangle{Min: minPoint, Max: maxPoint})

	// var cut_rectangle image.Rectangle

	for x := minPoint.X; x <= maxPoint.X; x++ {
		for y := minPoint.Y; y <= maxPoint.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()

			if ((r + g + b + a) % 255) > 120 {
				var okcount, allcount int
				for x1 := x - 1; x1 <= x+1; x1++ {
					for y1 := y - 1; y1 <= y+1; y1++ {
						r1, g1, b1, a1 := img.At(x1, y1).RGBA()
						allcount++
						if ((r1 + g1 + b1 + a1) % 255) > 120 {
							okcount++
						}
					}
				}
				fmt.Println(allcount)
				fmt.Println(okcount)
				if (allcount - okcount) > (allcount / 10) {
					rgba.Set(x, y, image.White)
				} else {
					rgba.Set(x, y, image.Black)
				}

			} else {
				rgba.Set(x, y, image.White)
			}
		}
	}
	jpeg.Encode(f1, rgba, &jpeg.Options{100})
	fmt.Println("ok")
}

func f3() {
	resp, err := http.Get("http://admin.daili.33m.com/admin-login/changeVerifyCode?width=500&height=200")
	checkErr(err)
	f, err := os.Create("./testdata/code.jpg")
	defer f.Close()
	checkErr(err)

	img, err := jpeg.Decode(resp.Body)
	checkErr(err)
	jpeg.Encode(f, img, &jpeg.Options{100})

	// reader := bufio.NewReader(resp.Body)
	// fmt.Println(resp.Body)

}
func f2() {
	f, err := os.Open("./testdata/videos.jpg")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	imgTextFile, err := os.Create("./testdata/video.txt")
	defer imgTextFile.Close()
	if err != nil {
		panic(err)
	}

	// img, err := png.Decode(f)
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	var minPoint image.Point

	var maxPoint image.Point
	minPoint = img.Bounds().Min

	maxPoint = img.Bounds().Max
	var str = ""
	for x := 0; x <= maxPoint.X; x++ {
		for y := 0; y <= maxPoint.Y; y++ {
			r, g, b, _ := img.At(y, x).RGBA()
			str += string(((r + g + b) % 26) + 65)
		}
		str += "\n"
	}
	imgTextFile.WriteString(str)
	fmt.Println(minPoint)
	fmt.Println(maxPoint)
}

func f1() {
	f, err := os.Open("./testdata/video-001.progressive.jpeg")

	defer f.Close()
	if err != nil {
		panic(err)
	}
	createfile, err := os.Create("./testdata/video.jpg")
	defer createfile.Close()
	if err != nil {
		panic(err)
	}

	img, errs := jpeg.Decode(bufio.NewReader(f))

	if errs != nil {
		panic(errs)
	}

	var x0, y0, x1, y1 int
	x0, y0, x1, y1 = 0, 0, 200, 200
	nrgba := image.NewRGBA(image.Rect(x0, y0, x1, y1))
	imgRgba := color.RGBA{}
	for i := 0; i <= x1; i++ {
		for y := 0; y < y1; y++ {
			imgRgba.B = uint8(y)
			imgRgba.G = uint8(i)
			imgRgba.R = uint8(i)
			nrgba.Set(i, y, imgRgba)
		}
		// img.ColorModel()
	}
	imgcolor := img.At(110, 10)
	fmt.Println(imgcolor.RGBA())
	//

	// jpeg.Encode(img, nrgba, &jpeg.Options{100})
	jpeg.Encode(createfile, nrgba, &jpeg.Options{70})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
