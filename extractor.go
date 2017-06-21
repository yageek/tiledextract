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

// Process a tilemap file
func (e *Extractor) Process(r io.Reader, outputPath string) error {
	_, err := e.extracts(r)
	if err != nil {
		return err
	}
	return nil

}
func (e *Extractor) extracts(r io.Reader) (*TileSet, error) {
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

func (e *Extractor) convert(set *TileSet, inputPath, outputPath string) error {

	if err := e.createFinal(outputPath); err != nil {
		return errors.Wrap(err, "Impossible to create destination directory")
	}
	inputPath = filepath.Join(inputPath, set.Image.Source)
	source, err := os.Open(inputPath)
	if err != nil {
		return errors.Wrap(err, "Impossible to read the tmx file")
	}
	defer source.Close()

	m, _, err := image.Decode(source)
	if err != nil {
		return errors.Wrap(err, "Impossible to decode image")
	}
	fmt.Printf("Tileset: %+v \n", set)

	for tile := 0; tile < set.TileCount; tile++ {
		c := tile % set.ColumnsCount
		r := tile / set.ColumnsCount
		w := c*(set.TileWidth+set.Spacing) + set.Margin
		h := r*(set.TileHeight+set.Spacing) + set.Margin

		dp := image.Pt(w, h)
		subRectangle := image.Rectangle{dp, dp.Add(image.Pt(set.TileWidth, set.TileHeight))}
		fmt.Printf("Getting (row: %d, colum: %d),: %+v \n", r, c, subRectangle)
		newImage := e.crop(m, subRectangle)

		outputImagePath := filepath.Join(outputPath, fmt.Sprintf("%d.png", tile))
		fmt.Printf("OutputPath: %v \n", outputImagePath)
		file, err := os.Create(outputImagePath)
		if err != nil {
			fmt.Printf("Error creating file: %v \n", err)
			continue
		}
		defer file.Close()

		if err := png.Encode(file, newImage); err != nil {
			fmt.Printf("Error: %v \n", err)
			continue
		}

		fmt.Println("Success")

	}
	return nil
}

func (e *Extractor) crop(src image.Image, sub image.Rectangle) image.Image {
	sample := image.NewNRGBA(sub)
	draw.Draw(sample, sub, src, sub.Min, draw.Src)
	return sample
}

func (e *Extractor) createFinal(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.Mkdir(path, 0777); err != nil {
			return err
		}
	}
	return nil
}
