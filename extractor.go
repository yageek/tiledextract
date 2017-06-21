package tiledextract

import (
	"encoding/xml"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Extractor extracts elements from a tiled set
type Extractor struct {
}

func (e *Extractor) Extracts(r io.Reader) (*TileSet, error) {
	decoder := xml.NewDecoder(r)

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}

		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == "tileset" {
				tileSet := &TileSet{}
				err := decoder.DecodeElement(tileSet, &t)
				return tileSet, err
			}
		}
	}
	return nil, io.EOF
}

func (e *Extractor) Convert(set TileSet, outputPath string) error {

	if err := e.createFinal(outputPath); err != nil {
		return errors.Wrap(err, "Impossible to create destination directory")
	}
	println("InputPath:", set.Image.Source)
	source, err := os.Open(set.Image.Source)
	if err != nil {
		return errors.Wrap(err, "Impossible to read tiled file")
	}
	defer source.Close()

	m, _, err := image.Decode(source)
	if err != nil {
		return errors.Wrap(err, "Impossible to decode image")
	}

	for c := 0; c < set.ColumnsCount; c++ {
		for r := 0; r < set.TileCount/set.ColumnsCount; r++ {

			subRectangle := image.Rect(c*set.TileWidth, r*set.TileHeight, set.TileWidth, set.TileHeight)
			newImage := e.crop(m, subRectangle)

			file, err := os.Open(filepath.Join(outputPath, fmt.Sprintf("%d_%d.png", r, c)))
			if err != nil {
				print(err)
				continue
			}
			defer file.Close()
			if err := png.Encode(file, newImage); err != nil {
				print(err)
			}
		}
	}
	return nil
}

func (e *Extractor) crop(src image.Image, sub image.Rectangle) image.Image {
	sample := image.NewRGBA(sub)
	draw.Draw(sample, sub, src, sub.Min, draw.Src)
	return sample
}

func (e *Extractor) createFinal(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {

		if err := os.Mkdir(path, 777); err != nil {
			return err
		}
	}
	return nil
}
