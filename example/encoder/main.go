package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"reflect"
	"time"

	libjpeg "github.com/bluvec/go-libjpeg/jpeg"
	"github.com/disintegration/imaging"
)

func encodelibjpeg(src image.Image) []byte {
	start := time.Now()
	buf, err := libjpeg.EncodeToBytes(src, &libjpeg.EncoderOptions{Quality: 90})
	if err != nil {
		log.Println("(imencode failed)", err)
		return nil
	}

	fmt.Println(time.Since(start))

	return buf
}

func main() {
	f, err := os.Open("../../test/images/kinkaku.jpg")
	if err != nil {
		log.Fatal(err)
	}
	buf := io.Reader(f)
	img, err := libjpeg.Decode(buf, &libjpeg.DecoderOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reflect.TypeOf(img))
	bound := img.Bounds()
	thumb := imaging.Resize(img, bound.Max.X/3, bound.Max.Y/3, imaging.Lanczos)
	log.Println(reflect.TypeOf(thumb))
	_ = encodelibjpeg(thumb)
	fmt.Println("ok")
}
