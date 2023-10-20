// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// An ASCII art videosink
type GstAASink struct {
	*GstVideoSink
}

func NewAASink() (*GstAASink, error) {
	e, err := gst.NewElement("aasink")
	if err != nil {
		return nil, err
	}
	return &GstAASink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewAASinkWithName(name string) (*GstAASink, error) {
	e, err := gst.NewElementWithName("aasink", name)
	if err != nil {
		return nil, err
	}
	return &GstAASink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// brightness
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstAASink) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstAASink) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// contrast
// Default: 16, Min: -2147483648, Max: 2147483647
func (e *GstAASink) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstAASink) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// gamma
// Default: 1, Min: 0, Max: 5
func (e *GstAASink) SetGamma(value float32) error {
	return e.Element.SetProperty("gamma", value)
}

func (e *GstAASink) GetGamma() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gamma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// inversion
// Default: false
func (e *GstAASink) SetInversion(value bool) error {
	return e.Element.SetProperty("inversion", value)
}

func (e *GstAASink) GetInversion() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("inversion")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// randomval
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstAASink) SetRandomval(value int) error {
	return e.Element.SetProperty("randomval", value)
}

func (e *GstAASink) GetRandomval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("randomval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// dither
// Default: no-dithering (0)
func (e *GstAASink) SetDither(value GstAASinkDitherers) error {
	e.Element.SetArg("dither", string(value))
	return nil
}

func (e *GstAASink) GetDither() (GstAASinkDitherers, error) {
	var value GstAASinkDitherers
	var ok bool
	v, err := e.Element.GetProperty("dither")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAASinkDitherers)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAASinkDitherers")
	}
	return value, nil
}

// driver
// Default: x11 (0)
func (e *GstAASink) SetDriver(value GstAASinkDrivers) error {
	e.Element.SetArg("driver", string(value))
	return nil
}

func (e *GstAASink) GetDriver() (GstAASinkDrivers, error) {
	var value GstAASinkDrivers
	var ok bool
	v, err := e.Element.GetProperty("driver")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAASinkDrivers)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAASinkDrivers")
	}
	return value, nil
}

// frame time
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstAASink) GetFrameTime() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frame-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// frames displayed
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstAASink) GetFramesDisplayed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frames-displayed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// height
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstAASink) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstAASink) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// width
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstAASink) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstAASink) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ASCII art effect
type GstAATv struct {
	*GstVideoFilter
}

func NewAATv() (*GstAATv, error) {
	e, err := gst.NewElement("aatv")
	if err != nil {
		return nil, err
	}
	return &GstAATv{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAATvWithName(name string) (*GstAATv, error) {
	e, err := gst.NewElementWithName("aatv", name)
	if err != nil {
		return nil, err
	}
	return &GstAATv{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Add noise to more closely approximate gray levels.
// Default: no-dithering (0)
func (e *GstAATv) SetDither(value GstAATvDitherers) error {
	e.Element.SetArg("dither", string(value))
	return nil
}

func (e *GstAATv) GetDither() (GstAATvDitherers, error) {
	var value GstAATvDitherers
	var ok bool
	v, err := e.Element.GetProperty("dither")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAATvDitherers)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAATvDitherers")
	}
	return value, nil
}

// Maximum frame delay between rain motion
// Default: 3, Min: 0, Max: 2147483647
func (e *GstAATv) SetRainDelayMax(value int) error {
	return e.Element.SetProperty("rain-delay-max", value)
}

func (e *GstAATv) GetRainDelayMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rain-delay-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum frame delay between rain motion
// Default: 0, Min: 0, Max: 2147483647
func (e *GstAATv) SetRainDelayMin(value int) error {
	return e.Element.SetProperty("rain-delay-min", value)
}

func (e *GstAATv) GetRainDelayMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rain-delay-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Automatically adjust brightness based on the previous frame's foreground pixel fill percentage
// Default: true
func (e *GstAATv) SetBrightnessAuto(value bool) error {
	return e.Element.SetProperty("brightness-auto", value)
}

func (e *GstAATv) GetBrightnessAuto() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("brightness-auto")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Automatically sets color-rain-bold, color-rain-normal, and color-rain-dim with progressively dimmer values (big-endian ARGB).
// Default: -16711936, Min: 0, Max: -1
func (e *GstAATv) SetColorRain(value uint) error {
	return e.Element.SetProperty("color-rain", value)
}

func (e *GstAATv) GetColorRain() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-rain")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Sets the brightest color to use for foreground ASCII text rain overlays (big-endian ARGB).
// Default: -16711936, Min: 0, Max: -1
func (e *GstAATv) SetColorRainBold(value uint) error {
	return e.Element.SetProperty("color-rain-bold", value)
}

func (e *GstAATv) GetColorRainBold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-rain-bold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Sets the normal brightness color to use for foreground ASCII text (big-endian ARGB).
// Default: -8421505, Min: 0, Max: -1
func (e *GstAATv) SetColorTextNormal(value uint) error {
	return e.Element.SetProperty("color-text-normal", value)
}

func (e *GstAATv) GetColorTextNormal() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-text-normal")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Contrast
// Default: 0, Min: 0, Max: 255
func (e *GstAATv) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstAATv) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Width of the ASCII canvas
// Default: 80, Min: 0, Max: 2147483647
func (e *GstAATv) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstAATv) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Brightness
// Default: 0, Min: -255, Max: 255
func (e *GstAATv) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstAATv) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum target foreground pixel fill percentage for automatic brightness control
// Default: 0.3, Min: 0, Max: 1
func (e *GstAATv) SetBrightnessMin(value float32) error {
	return e.Element.SetProperty("brightness-min", value)
}

func (e *GstAATv) GetBrightnessMin() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("brightness-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Set the direction of raindrops
// Default: none (0)
func (e *GstAATv) SetRainMode(value GstAATvRainModes) error {
	e.Element.SetArg("rain-mode", string(value))
	return nil
}

func (e *GstAATv) GetRainMode() (GstAATvRainModes, error) {
	var value GstAATvRainModes
	var ok bool
	v, err := e.Element.GetProperty("rain-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAATvRainModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAATvRainModes")
	}
	return value, nil
}

// Color to use as the background for the ASCII text (big-endian ARGB).
// Default: -16777216, Min: 0, Max: -1
func (e *GstAATv) SetColorBackground(value uint) error {
	return e.Element.SetProperty("color-background", value)
}

func (e *GstAATv) GetColorBackground() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-background")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Sets the normal brightness color to use for foreground ASCII text rain overlays (big-endian ARGB).
// Default: -16744704, Min: 0, Max: -1
func (e *GstAATv) SetColorRainNormal(value uint) error {
	return e.Element.SetProperty("color-rain-normal", value)
}

func (e *GstAATv) GetColorRainNormal() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-rain-normal")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Sets the brightest color to use for foreground ASCII text (big-endian ARGB).
// Default: -1, Min: 0, Max: -1
func (e *GstAATv) SetColorTextBold(value uint) error {
	return e.Element.SetProperty("color-text-bold", value)
}

func (e *GstAATv) GetColorTextBold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-text-bold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Gamma correction
// Default: 1, Min: 0, Max: 5
func (e *GstAATv) SetGamma(value float32) error {
	return e.Element.SetProperty("gamma", value)
}

func (e *GstAATv) GetGamma() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("gamma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Actual calculated foreground pixel fill percentage
// Default: 0.35, Min: 0, Max: 1
func (e *GstAATv) GetBrightnessActual() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("brightness-actual")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Sets the dimmest brightness color to use for foreground ASCII text (big-endian ARGB).
// Default: -12632257, Min: 0, Max: -1
func (e *GstAATv) SetColorTextDim(value uint) error {
	return e.Element.SetProperty("color-text-dim", value)
}

func (e *GstAATv) GetColorTextDim() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-text-dim")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum length of a rain
// Default: 30, Min: 0, Max: 2147483647
func (e *GstAATv) SetRainLengthMax(value int) error {
	return e.Element.SetProperty("rain-length-max", value)
}

func (e *GstAATv) GetRainLengthMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rain-length-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Adds a random value in the range (-randomval/2,ranomval/2) to each pixel during rendering
// Default: 0, Min: 0, Max: 255
func (e *GstAATv) SetRandomval(value int) error {
	return e.Element.SetProperty("randomval", value)
}

func (e *GstAATv) GetRandomval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("randomval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sets the dimmest brightness color to use for foreground ASCII text rain overlays (big-endian ARGB).
// Default: -16761088, Min: 0, Max: -1
func (e *GstAATv) SetColorRainDim(value uint) error {
	return e.Element.SetProperty("color-rain-dim", value)
}

func (e *GstAATv) GetColorRainDim() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-rain-dim")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Percentage chance for a raindrop to spawn
// Default: 0.2, Min: 0, Max: 1
func (e *GstAATv) SetRainSpawnRate(value float32) error {
	return e.Element.SetProperty("rain-spawn-rate", value)
}

func (e *GstAATv) GetRainSpawnRate() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rain-spawn-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Height of the ASCII canvas
// Default: 24, Min: 0, Max: 2147483647
func (e *GstAATv) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstAATv) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum length of a rain
// Default: 4, Min: 0, Max: 2147483647
func (e *GstAATv) SetRainLengthMin(value int) error {
	return e.Element.SetProperty("rain-length-min", value)
}

func (e *GstAATv) GetRainLengthMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rain-length-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum target foreground pixel fill percentage for automatic brightness control
// Default: 0.4, Min: 0, Max: 1
func (e *GstAATv) SetBrightnessMax(value float32) error {
	return e.Element.SetProperty("brightness-max", value)
}

func (e *GstAATv) GetBrightnessMax() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("brightness-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Automatically sets color-test-bold, color-text-normal, and color-text-dim with progressively dimmer values (big-endian ARGB).
// Default: -1, Min: 0, Max: -1
func (e *GstAATv) SetColorText(value uint) error {
	return e.Element.SetProperty("color-text", value)
}

func (e *GstAATv) GetColorText() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color-text")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// AAlib Font
// Default: Standard-vga-8x8-font (0)
func (e *GstAATv) SetFont(value GstAATvFonts) error {
	e.Element.SetArg("font", string(value))
	return nil
}

func (e *GstAATv) GetFont() (GstAATvFonts, error) {
	var value GstAATvFonts
	var ok bool
	v, err := e.Element.GetProperty("font")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAATvFonts)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAATvFonts")
	}
	return value, nil
}

type GstAATvDitherers string

const (
	GstAATvDitherersNoDithering             GstAATvDitherers = "no-dithering"              // no dithering
	GstAATvDitherersErrorDistribution       GstAATvDitherers = "error-distribution"        // error-distribution
	GstAATvDitherersFloydSteelbergDithering GstAATvDitherers = "floyd-steelberg-dithering" // floyd-steelberg dithering
)

type GstAATvFonts string

const (
	GstAATvFontsStandardVga8x8Font     GstAATvFonts = "Standard-vga-8x8-font"      // vga8
	GstAATvFontsStandardVga8x9Font     GstAATvFonts = "Standard-vga-8x9-font"      // vga9
	GstAATvFontsStandardMdavga8x14Font GstAATvFonts = "Standard-mda/vga-8x14-font" // mda14
	GstAATvFontsStandardVga8x16Font    GstAATvFonts = "Standard-vga-8x16-font"     // vga16
	GstAATvFontsX8x13Font              GstAATvFonts = "X-8x13-font"                // X8x13
	GstAATvFontsX8x13boldFont          GstAATvFonts = "X-8x13bold-font"            // X8x13bold
	GstAATvFontsStandardX8x16Font      GstAATvFonts = "Standard-X-8x16-font"       // X8x16
	GstAATvFontsLineFont8x8            GstAATvFonts = "line-Font-8x8"              // line
	GstAATvFontsFont8x8FromVgagl       GstAATvFonts = "Font-8x8-from-vgagl"        // vgagl8
	GstAATvFontsAdobeCourier           GstAATvFonts = "Adobe-courier"              // courier
)

type GstAATvRainModes string

const (
	GstAATvRainModesNone  GstAATvRainModes = "none"  // No Rain
	GstAATvRainModesDown  GstAATvRainModes = "down"  // Rain Down
	GstAATvRainModesUp    GstAATvRainModes = "up"    // Rain Up
	GstAATvRainModesLeft  GstAATvRainModes = "left"  // Rain Left
	GstAATvRainModesRight GstAATvRainModes = "right" // Rain Right
)

type GstAASinkDitherers string

const (
	GstAASinkDitherersNoDithering             GstAASinkDitherers = "no-dithering"              // no dithering
	GstAASinkDitherersErrorDistribution       GstAASinkDitherers = "error-distribution"        // error-distribution
	GstAASinkDitherersFloydSteelbergDithering GstAASinkDitherers = "floyd-steelberg-dithering" // floyd-steelberg dithering
)

type GstAASinkDrivers string

const (
	GstAASinkDriversX11    GstAASinkDrivers = "x11"    // X11 driver 1.1
	GstAASinkDriversLinux  GstAASinkDrivers = "linux"  // Linux pc console driver 1.0
	GstAASinkDriversSlang  GstAASinkDrivers = "slang"  // Slang driver 1.0
	GstAASinkDriversCurses GstAASinkDrivers = "curses" // Curses driver 1.0
	GstAASinkDriversStdout GstAASinkDrivers = "stdout" // Standard output driver
	GstAASinkDriversStderr GstAASinkDrivers = "stderr" // Standard error driver
)
