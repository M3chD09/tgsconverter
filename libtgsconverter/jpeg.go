package libtgsconverter

import (
	"bytes"
	"image"
	"image/jpeg"
)

type tojpeg struct {
	result []byte
}

func (to_jpeg *tojpeg) init(w uint, h uint, options ConverterOptions) {
}

func (to_jpeg *tojpeg) SupportsAnimation() bool {
	return false
}

func (to_jpeg *tojpeg) AddFrame(image *image.RGBA, fps uint) error {
	var data []byte
	w := bytes.NewBuffer(data)
	if err := jpeg.Encode(w, image, nil); err != nil {
		return err
	}
	to_jpeg.result = w.Bytes()
	return nil
}

func (to_jpeg *tojpeg) Result() []byte {
	return to_jpeg.result
}
