// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Overlay Qrcodes over each buffer with buffer information and custom data
type GstDebugQROverlay struct {
	*GstBaseQROverlay
}

func NewDebugQROverlay() (*GstDebugQROverlay, error) {
	e, err := gst.NewElement("debugqroverlay")
	if err != nil {
		return nil, err
	}
	return &GstDebugQROverlay{GstBaseQROverlay: &GstBaseQROverlay{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewDebugQROverlayWithName(name string) (*GstDebugQROverlay, error) {
	e, err := gst.NewElementWithName("debugqroverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstDebugQROverlay{GstBaseQROverlay: &GstBaseQROverlay{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// List of comma separated values that the extra data value will be  cycled from at each interval, example array structure : "240,480,720,960,1200,1440,1680,1920"
// Default: NULL
func (e *GstDebugQROverlay) SetExtraDataArray(value string) error {
	return e.Element.SetProperty("extra-data-array", value)
}

func (e *GstDebugQROverlay) GetExtraDataArray() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("extra-data-array")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Extra data append into the Qrcode at the first buffer of each  interval
// Default: 60, Min: 0, Max: 9223372036854775807
func (e *GstDebugQROverlay) SetExtraDataIntervalBuffers(value int64) error {
	return e.Element.SetProperty("extra-data-interval-buffers", value)
}

func (e *GstDebugQROverlay) GetExtraDataIntervalBuffers() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("extra-data-interval-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Json key name for extra append data
// Default: NULL
func (e *GstDebugQROverlay) SetExtraDataName(value string) error {
	return e.Element.SetProperty("extra-data-name", value)
}

func (e *GstDebugQROverlay) GetExtraDataName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("extra-data-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Numbers of consecutive buffers that the extra data will be inserted  (counting the first buffer)
// Default: 1, Min: 0, Max: 9223372036854775807
func (e *GstDebugQROverlay) SetExtraDataSpanBuffers(value int64) error {
	return e.Element.SetProperty("extra-data-span-buffers", value)
}

func (e *GstDebugQROverlay) GetExtraDataSpanBuffers() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("extra-data-span-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Overlay Qrcodes over each buffer with data passed in
type GstQROverlay struct {
	*GstBaseQROverlay
}

func NewQROverlay() (*GstQROverlay, error) {
	e, err := gst.NewElement("qroverlay")
	if err != nil {
		return nil, err
	}
	return &GstQROverlay{GstBaseQROverlay: &GstBaseQROverlay{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewQROverlayWithName(name string) (*GstQROverlay, error) {
	e, err := gst.NewElementWithName("qroverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstQROverlay{GstBaseQROverlay: &GstBaseQROverlay{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Data to write in the QRCode to be overlaid
// Default: NULL
func (e *GstQROverlay) SetData(value string) error {
	return e.Element.SetProperty("data", value)
}

func (e *GstQROverlay) GetData() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("data")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstBaseQROverlay struct {
	*GstVideoFilter
}

// Pixel size of each Qrcode pixel
// Default: 3, Min: 1, Max: 100
func (e *GstBaseQROverlay) SetPixelSize(value float32) error {
	return e.Element.SetProperty("pixel-size", value)
}

func (e *GstBaseQROverlay) GetPixelSize() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pixel-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// qrcode-error-correction
// Default: Approx 15%% (1)
func (e *GstBaseQROverlay) SetQrcodeErrorCorrection(value GstQrcodeOverlayCorrection) error {
	e.Element.SetArg("qrcode-error-correction", string(value))
	return nil
}

func (e *GstBaseQROverlay) GetQrcodeErrorCorrection() (GstQrcodeOverlayCorrection, error) {
	var value GstQrcodeOverlayCorrection
	var ok bool
	v, err := e.Element.GetProperty("qrcode-error-correction")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQrcodeOverlayCorrection)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQrcodeOverlayCorrection")
	}
	return value, nil
}

// X position (in percent of the width)
// Default: 50, Min: 0, Max: 100
func (e *GstBaseQROverlay) SetX(value float32) error {
	return e.Element.SetProperty("x", value)
}

func (e *GstBaseQROverlay) GetX() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Y position (in percent of the height)
// Default: 50, Min: 0, Max: 100
func (e *GstBaseQROverlay) SetY(value float32) error {
	return e.Element.SetProperty("y", value)
}

func (e *GstBaseQROverlay) GetY() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

type GstQrcodeOverlayCorrection string

const (
	GstQrcodeOverlayCorrectionApprox7  GstQrcodeOverlayCorrection = "Approx 7%%"  // Level L
	GstQrcodeOverlayCorrectionApprox15 GstQrcodeOverlayCorrection = "Approx 15%%" // Level M
	GstQrcodeOverlayCorrectionApprox25 GstQrcodeOverlayCorrection = "Approx 25%%" // Level Q
	GstQrcodeOverlayCorrectionApprox30 GstQrcodeOverlayCorrection = "Approx 30%%" // Level H
)
