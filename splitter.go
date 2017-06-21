package tiledextract

import (
	"image"
	"os"
	"path/filepath"

	_ "golang.org/x/image/bmp"
)

// Extractor extracts elements from a tiled set
type Extractor struct {
}

func (e *Extractor) Extracts(set TileSet, path string) error {

	if err := e.createFinal(path); err != nil {
		return err
	}

	tmxDir := filepath.Dir(set.Image.Source)

	source, err := os.Open(tmxDir)
	if err != nil {
		return err
	}
	defer source.Close()

	m, _, err := image.Decode(source)
	if err != nil {
		return err
	}

	for c := 0; c < set.ColumnsCount; c++ {
		for r := 0; r < set.TileCount/set.ColumnsCount; r++ {
			subRectangle := image.Rect(c*set.TileWidth, r*set.TileHeight, set.TileWidth, set.TileHeight)
			sample := image.NewRGBA(subRectangle)
		}
	}

	return nil
}

func (e *Extractor) writeImage(r *image.RGBA) {

}

func (e *Extractor) createFinal(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {

		if err := os.Mkdir(path, 777); err != nil {
			return err
		}
	}
}
