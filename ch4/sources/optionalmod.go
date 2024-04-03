package main

import "fmt"

// version 1.0

//type FinishedHouse struct {
//	style                  int
//	centralAirConditioning bool
//	floorMaterial          string
//	wallMaterial           string
//}
//
//func NewFinishedHouse(style int, centralAirConditioning bool,
//	floorMaterial string, wallMaterial string) *FinishedHouse {
//
//	h := &FinishedHouse{
//		style:                  style,
//		centralAirConditioning: centralAirConditioning,
//		floorMaterial:          floorMaterial,
//		wallMaterial:           wallMaterial,
//	}
//
//	return h
//}

// version 2.0

//type FinishedHouse struct {
//	style                  int
//	centralAirConditioning bool
//	floorMaterial          string
//	wallMaterial           string
//}
//
//type Options struct {
//	Style                  int
//	CentralAirConditioning bool
//	FloorMaterial          string
//	WallMaterial           string
//}
//
//func NewFinishedHouse(options *Options) *FinishedHouse {
//	var (
//		style                  = 0
//		centralAirConditioning = true
//		floorMaterial          = "wood"
//		wallMaterial           = "paper"
//	)
//
//	if options != nil {
//		style = options.Style
//		centralAirConditioning = options.CentralAirConditioning
//		floorMaterial = options.FloorMaterial
//		wallMaterial = options.WallMaterial
//	}
//
//	h := &FinishedHouse{
//		style:                  style,
//		centralAirConditioning: centralAirConditioning,
//		floorMaterial:          floorMaterial,
//		wallMaterial:           wallMaterial,
//	}
//
//	return h
//}

// version 3.0

type FinishedHouse struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

type Option func(*FinishedHouse)

func NewFinishedHouse(options ...Option) *FinishedHouse {
	h := &FinishedHouse{
		style:                  0,
		centralAirConditioning: true,
		floorMaterial:          "wood",
		wallMaterial:           "paper",
	}

	for _, option := range options {
		option(h)
	}

	return h
}

func WithStyle(style int) Option {
	return func(h *FinishedHouse) {
		h.style = style
	}
}

func WithFloorMaterial(material string) Option {
	return func(h *FinishedHouse) {
		h.floorMaterial = material
	}
}

func WithWallMaterial(material string) Option {
	return func(h *FinishedHouse) {
		h.wallMaterial = material
	}
}

func WithCentralAirConditioning(cac bool) Option {
	return func(h *FinishedHouse) {
		h.centralAirConditioning = cac
	}
}

func main() {
	//fmt.Printf("%+v\n", NewFinishedHouse(0, true, "wood", "paper"))

	house := NewFinishedHouse(WithStyle(1))
	fmt.Printf("%+v\n", house)
}
