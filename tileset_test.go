package tiledextract

import (
	"encoding/xml"
	"testing"
)

func TestExtract(t *testing.T) {

	raw := `<?xml version="1.0" encoding="UTF-8"?>
		<tileset name="dwqdwdqwd" tilewidth="50" tileheight="50" spacing="1" margin="1" tilecount="15" columns="5">
 		<image source="tmw_desert_spacing.png" width="265" height="199"/>
		</tileset>
`
	tile := TileSet{}
	if err := xml.Unmarshal([]byte(raw), &tile); err != nil {
		t.Errorf("The tiled set should be correctly decoded: %v", err)
	}

	if tile.Name != "dwqdwdqwd" {
		t.Errorf("The name should be correcly decoded")
	}

	if tile.TileWidth != 50 || tile.TileHeight != 50 {
		t.Errorf("Tile sizeshould be correctly decoded")
	}

	if tile.Spacing != 1 || tile.Margin != 1 {
		t.Errorf("Spacing and margin should be correctly decoded")
	}

	if tile.TileCount != 15 || tile.ColumnsCount != 5 {
		t.Errorf("Count and columns should be correctly decoded")
	}

	image := tile.Image
	if image.Source != "tmw_desert_spacing.png" || image.Height != 199 || image.Width != 265 {
		t.Errorf("The image should be correctly decoded")
	}
}
