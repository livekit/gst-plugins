// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decode JPEG2000 streams
type GstOpenJPEGDec struct {
	*GstVideoDecoder
}

func NewOpenJPEGDec() (*GstOpenJPEGDec, error) {
	e, err := gst.NewElement("openjpegdec")
	if err != nil {
		return nil, err
	}
	return &GstOpenJPEGDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewOpenJPEGDecWithName(name string) (*GstOpenJPEGDec, error) {
	e, err := gst.NewElementWithName("openjpegdec", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenJPEGDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// Maximum number of worker threads to spawn used by openjpeg internally. (0 = no thread)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstOpenJPEGDec) SetMaxThreads(value int) error {
	return e.Element.SetProperty("max-threads", value)
}

func (e *GstOpenJPEGDec) GetMaxThreads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum number of worker threads to spawn according to the frame boundary. (0 = no thread)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstOpenJPEGDec) SetMaxSliceThreads(value int) error {
	return e.Element.SetProperty("max-slice-threads", value)
}

func (e *GstOpenJPEGDec) GetMaxSliceThreads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-slice-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encode JPEG2000 streams
type GstOpenJPEGEnc struct {
	*GstVideoEncoder
}

func NewOpenJPEGEnc() (*GstOpenJPEGEnc, error) {
	e, err := gst.NewElement("openjpegenc")
	if err != nil {
		return nil, err
	}
	return &GstOpenJPEGEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewOpenJPEGEncWithName(name string) (*GstOpenJPEGEnc, error) {
	e, err := gst.NewElementWithName("openjpegenc", name)
	if err != nil {
		return nil, err
	}
	return &GstOpenJPEGEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// Number of layers
// Default: 1, Min: 1, Max: 10
func (e *GstOpenJPEGEnc) SetNumLayers(value int) error {
	return e.Element.SetProperty("num-layers", value)
}

func (e *GstOpenJPEGEnc) GetNumLayers() (int, error) {
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

// Number of resolutions
// Default: 6, Min: 1, Max: 10
func (e *GstOpenJPEGEnc) SetNumResolutions(value int) error {
	return e.Element.SetProperty("num-resolutions", value)
}

func (e *GstOpenJPEGEnc) GetNumResolutions() (int, error) {
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

// Number of stripes for low latency encoding. (1 = low latency disabled)
// Default: 1, Min: 1, Max: 2147483647
func (e *GstOpenJPEGEnc) SetNumStripes(value int) error {
	return e.Element.SetProperty("num-stripes", value)
}

func (e *GstOpenJPEGEnc) GetNumStripes() (int, error) {
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

// Max number of simultaneous threads to encode stripe or frame, default: encode with streaming thread.
// Default: 0, Min: 0, Max: 2147483647
func (e *GstOpenJPEGEnc) SetNumThreads(value uint) error {
	return e.Element.SetProperty("num-threads", value)
}

func (e *GstOpenJPEGEnc) GetNumThreads() (uint, error) {
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

// Progression order
// Default: lrcp (0)
func (e *GstOpenJPEGEnc) SetProgressionOrder(value GstOpenJPEGEncProgressionOrder) error {
	e.Element.SetArg("progression-order", string(value))
	return nil
}

func (e *GstOpenJPEGEnc) GetProgressionOrder() (GstOpenJPEGEncProgressionOrder, error) {
	var value GstOpenJPEGEncProgressionOrder
	var ok bool
	v, err := e.Element.GetProperty("progression-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenJPEGEncProgressionOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenJPEGEncProgressionOrder")
	}
	return value, nil
}

// Tile Height
// Default: 0, Min: 0, Max: 2147483647
func (e *GstOpenJPEGEnc) SetTileHeight(value int) error {
	return e.Element.SetProperty("tile-height", value)
}

func (e *GstOpenJPEGEnc) GetTileHeight() (int, error) {
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

// Tile Offset Y
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstOpenJPEGEnc) SetTileOffsetY(value int) error {
	return e.Element.SetProperty("tile-offset-y", value)
}

func (e *GstOpenJPEGEnc) GetTileOffsetY() (int, error) {
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

// Tile Width
// Default: 0, Min: 0, Max: 2147483647
func (e *GstOpenJPEGEnc) SetTileWidth(value int) error {
	return e.Element.SetProperty("tile-width", value)
}

func (e *GstOpenJPEGEnc) GetTileWidth() (int, error) {
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

// Tile Offset X
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstOpenJPEGEnc) SetTileOffsetX(value int) error {
	return e.Element.SetProperty("tile-offset-x", value)
}

func (e *GstOpenJPEGEnc) GetTileOffsetX() (int, error) {
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

type GstOpenJPEGEncProgressionOrder string

const (
	GstOpenJPEGEncProgressionOrderLrcp GstOpenJPEGEncProgressionOrder = "lrcp" // LRCP
	GstOpenJPEGEncProgressionOrderRlcp GstOpenJPEGEncProgressionOrder = "rlcp" // RLCP
	GstOpenJPEGEncProgressionOrderRpcl GstOpenJPEGEncProgressionOrder = "rpcl" // RPCL
	GstOpenJPEGEncProgressionOrderPcrl GstOpenJPEGEncProgressionOrder = "pcrl" // PCRL
	GstOpenJPEGEncProgressionOrderCrpl GstOpenJPEGEncProgressionOrder = "crpl" // CPRL
)
