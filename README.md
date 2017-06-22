# tiledextract

A tool helping you to extract the tiles of a [Tiled](https://github.com/bjorn/tiled) tilemap.

## Install

```shell
go get -u -v github.com/yageek/tiledextract
```

## Usage

To extract tiles from an `input.tmx` file to the `output` directory:

```shell
tiledextract -i input.tmx -o output
```

This will create a bunch of tiles images inside `output`