package gowordle

import (
	// "github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
	// "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	// "golang.org/x/image/font/opentype"
	"image/color"
)

const (
	title  string = "GoWordle"
	width  int    = 435
	height int    = 600
	rows   int    = 6
	cols   int    = 5
)

var (
	fontSize        int = 32
	mplusNormalFont font.Face
	bkg             = color.White
	lightgrey       = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	grey            = color.RGBA{0x77, 0x7c, 0x7e, 0xff}
	yellow          = color.RGBA{0xcd, 0xb3, 0x5d, 0xff}
	green           = color.RGBA{0x60, 0xa6, 0x65, 0xff}
	fontColor       = color.Black
	edge            = false
	alphabet        = "qwertyuiopasdfghjklzxcvbnm"
)
