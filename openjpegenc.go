// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type Openjpegenc struct {
	Element *gst.Element
}

func NewOpenjpegenc() (*Openjpegenc, error) {
	e, err := gst.NewElement("openjpegenc")
	if err != nil {
		return nil, err
	}
	return &Openjpegenc{Element: e}, nil
}

func NewOpenjpegencWithName(name string) (*Openjpegenc, error) {
	e, err := gst.NewElementWithName("openjpegenc", name)
	if err != nil {
		return nil, err
	}
	return &Openjpegenc{Element: e}, nil
}

// ----- Properties -----

// num-layers (int)
//
// Number of layers

func (e *Openjpegenc) GetNumLayers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-layers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegenc) SetNumLayers(value int) error {
	return e.Element.SetProperty("num-layers", value)
}

// num-resolutions (int)
//
// Number of resolutions

func (e *Openjpegenc) GetNumResolutions() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-resolutions")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegenc) SetNumResolutions(value int) error {
	return e.Element.SetProperty("num-resolutions", value)
}

// num-stripes (int)
//
// Number of stripes to use for low latency encoding . (1 = low latency disabled)

func (e *Openjpegenc) GetNumStripes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-stripes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegenc) SetNumStripes(value int) error {
	return e.Element.SetProperty("num-stripes", value)
}

// num-threads (uint)
//
// Max number of simultaneous threads to encode stripes, default: encode with streaming thread

func (e *Openjpegenc) GetNumThreads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Openjpegenc) SetNumThreads(value uint) error {
	return e.Element.SetProperty("num-threads", value)
}

// progression-order (GstOpenJpegencProgressionOrder)
//
// Progression order

func (e *Openjpegenc) GetProgressionOrder() (GstOpenJpegencProgressionOrder, error) {
	var value GstOpenJpegencProgressionOrder
	var ok bool
	v, err := e.Element.GetProperty("progression-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenJpegencProgressionOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenJpegencProgressionOrder")
	}
	return value, nil
}

func (e *Openjpegenc) SetProgressionOrder(value GstOpenJpegencProgressionOrder) error {
	e.Element.SetArg("progression-order", string(value))
	return nil
}

// tile-height (int)
//
// Tile Height

func (e *Openjpegenc) GetTileHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tile-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegenc) SetTileHeight(value int) error {
	return e.Element.SetProperty("tile-height", value)
}

// tile-offset-x (int)
//
// Tile Offset X

func (e *Openjpegenc) GetTileOffsetX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tile-offset-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegenc) SetTileOffsetX(value int) error {
	return e.Element.SetProperty("tile-offset-x", value)
}

// tile-offset-y (int)
//
// Tile Offset Y

func (e *Openjpegenc) GetTileOffsetY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tile-offset-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegenc) SetTileOffsetY(value int) error {
	return e.Element.SetProperty("tile-offset-y", value)
}

// tile-width (int)
//
// Tile Width

func (e *Openjpegenc) GetTileWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tile-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Openjpegenc) SetTileWidth(value int) error {
	return e.Element.SetProperty("tile-width", value)
}

// ----- Constants -----

type GstOpenJpegencProgressionOrder string

const (
	GstOpenJpegencProgressionOrderLrcp GstOpenJpegencProgressionOrder = "lrcp" // lrcp (0) – LRCP
	GstOpenJpegencProgressionOrderRlcp GstOpenJpegencProgressionOrder = "rlcp" // rlcp (1) – RLCP
	GstOpenJpegencProgressionOrderRpcl GstOpenJpegencProgressionOrder = "rpcl" // rpcl (2) – RPCL
	GstOpenJpegencProgressionOrderPcrl GstOpenJpegencProgressionOrder = "pcrl" // pcrl (3) – PCRL
	GstOpenJpegencProgressionOrderCrpl GstOpenJpegencProgressionOrder = "crpl" // crpl (4) – CPRL
)

