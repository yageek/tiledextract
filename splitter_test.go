package tiledextract

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestSplitter(t *testing.T) {

	file, _ := os.Open("resources/TileMap.tmx")
	defer file.Close()
	tile := TileSet{}

	xml.NewDecoder(file).Decode(&tile)

	ext := &Extractor{}
	err := ext.Extracts(tile, "output")
	if err != nil {
		t.Errorf("Should not have failed: %v", err)
	}
}
