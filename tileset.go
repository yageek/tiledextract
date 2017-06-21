package tiledextract

import "encoding/xml"

// TileSet represents a Tiled XML
type TileSet struct {
	XMLName      xml.Name `xml:"tileset"`
	Name         string   `xml:"name,attr"`
	TileWidth    int      `xml:"tilewidth,attr"`
	TileHeight   int      `xml:"tileheight,attr"`
	Spacing      int      `xml:"spacing,attr"`
	Margin       int      `xml:"margin,attr"`
	TileCount    int      `xml:"tilecount,attr"`
	ColumnsCount int      `xml:"columns,attr"`
	Image        Image
}

// Image is a source image inside the TileSet
type Image struct {
	XMLName xml.Name `xml:"image"`
	Source  string   `xml:"source,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}
