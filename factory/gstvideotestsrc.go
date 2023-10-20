// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Creates a test video stream
type GstVideoTestSrc struct {
	*base.GstPushSrc
}

func NewVideoTestSrc() (*GstVideoTestSrc, error) {
	e, err := gst.NewElement("videotestsrc")
	if err != nil {
		return nil, err
	}
	return &GstVideoTestSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewVideoTestSrcWithName(name string) (*GstVideoTestSrc, error) {
	e, err := gst.NewElementWithName("videotestsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoTestSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Scroll image number of pixels per frame (positive is scroll to the left)
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetHorizontalSpeed(value int) error {
	return e.Element.SetProperty("horizontal-speed", value)
}

func (e *GstVideoTestSrc) GetHorizontalSpeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("horizontal-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate 1st order t phase, for generating phase rotation as a function of time
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKt(value int) error {
	return e.Element.SetProperty("kt", value)
}

func (e *GstVideoTestSrc) GetKt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate 1st order x phase, for generating constant horizontal frequencies
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKx(value int) error {
	return e.Element.SetProperty("kx", value)
}

func (e *GstVideoTestSrc) GetKx() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kx")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate 2nd order y phase, normailsed to ky2/256 cycles per vertical pixel at height/2 from origin
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKy2(value int) error {
	return e.Element.SetProperty("ky2", value)
}

func (e *GstVideoTestSrc) GetKy2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ky2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Background color to use (big-endian ARGB)
// Default: -16777216, Min: 0, Max: -1
func (e *GstVideoTestSrc) SetBackgroundColor(value uint) error {
	return e.Element.SetProperty("background-color", value)
}

func (e *GstVideoTestSrc) GetBackgroundColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("background-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// For pattern=ball, invert colors every second.
// Default: false
func (e *GstVideoTestSrc) SetFlip(value bool) error {
	return e.Element.SetProperty("flip", value)
}

func (e *GstVideoTestSrc) GetFlip() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("flip")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to act as a live source
// Default: false
func (e *GstVideoTestSrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstVideoTestSrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Zoneplate x*y product phase
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKxy(value int) error {
	return e.Element.SetProperty("kxy", value)
}

func (e *GstVideoTestSrc) GetKxy() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate 1st order y phase, for generating content vertical frequencies
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKy(value int) error {
	return e.Element.SetProperty("ky", value)
}

func (e *GstVideoTestSrc) GetKy() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ky")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate 2nd order products y offset
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetYoffset(value int) error {
	return e.Element.SetProperty("yoffset", value)
}

func (e *GstVideoTestSrc) GetYoffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("yoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// An offset added to timestamps set on buffers (in ns)
// Default: 0, Min: 0, Max: 9223372036854775807
func (e *GstVideoTestSrc) SetTimestampOffset(value int64) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

func (e *GstVideoTestSrc) GetTimestampOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// For pattern=ball, which counter defines the position of the ball.
// Default: frames (0)
func (e *GstVideoTestSrc) SetAnimationMode(value GstVideoTestSrcAnimationMode) error {
	e.Element.SetArg("animation-mode", string(value))
	return nil
}

func (e *GstVideoTestSrc) GetAnimationMode() (GstVideoTestSrcAnimationMode, error) {
	var value GstVideoTestSrcAnimationMode
	var ok bool
	v, err := e.Element.GetProperty("animation-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoTestSrcAnimationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoTestSrcAnimationMode")
	}
	return value, nil
}

// Foreground color to use (big-endian ARGB)
// Default: -1, Min: 0, Max: -1
func (e *GstVideoTestSrc) SetForegroundColor(value uint) error {
	return e.Element.SetProperty("foreground-color", value)
}

func (e *GstVideoTestSrc) GetForegroundColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("foreground-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Zoneplate zero order phase, for generating plain fields or phase offsets
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetK0(value int) error {
	return e.Element.SetProperty("k0", value)
}

func (e *GstVideoTestSrc) GetK0() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("k0")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate 2nd order x phase, normalised to kx2/256 cycles per horizontal pixel at width/2 from origin
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKx2(value int) error {
	return e.Element.SetProperty("kx2", value)
}

func (e *GstVideoTestSrc) GetKx2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kx2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate x*t product phase, normalised to kxy/256 cycles per vertical pixel at width/2 from origin
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKxt(value int) error {
	return e.Element.SetProperty("kxt", value)
}

func (e *GstVideoTestSrc) GetKxt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kxt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// For pattern=ball, what motion the ball does
// Default: wavy (0)
func (e *GstVideoTestSrc) SetMotion(value GstVideoTestSrcMotionType) error {
	e.Element.SetArg("motion", string(value))
	return nil
}

func (e *GstVideoTestSrc) GetMotion() (GstVideoTestSrcMotionType, error) {
	var value GstVideoTestSrcMotionType
	var ok bool
	v, err := e.Element.GetProperty("motion")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoTestSrcMotionType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoTestSrcMotionType")
	}
	return value, nil
}

// Type of test pattern to generate
// Default: smpte (0)
func (e *GstVideoTestSrc) SetPattern(value GstVideoTestSrcPattern) error {
	e.Element.SetArg("pattern", string(value))
	return nil
}

func (e *GstVideoTestSrc) GetPattern() (GstVideoTestSrcPattern, error) {
	var value GstVideoTestSrcPattern
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoTestSrcPattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoTestSrcPattern")
	}
	return value, nil
}

// Zoneplate 2nd order products x offset
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetXoffset(value int) error {
	return e.Element.SetProperty("xoffset", value)
}

func (e *GstVideoTestSrc) GetXoffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("xoffset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate 2nd order t phase, t*t/256 cycles per picture
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKt2(value int) error {
	return e.Element.SetProperty("kt2", value)
}

func (e *GstVideoTestSrc) GetKt2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kt2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Zoneplate y*t product phase
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstVideoTestSrc) SetKyt(value int) error {
	return e.Element.SetProperty("kyt", value)
}

func (e *GstVideoTestSrc) GetKyt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kyt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstVideoTestSrcAnimationMode string

const (
	GstVideoTestSrcAnimationModeFrames      GstVideoTestSrcAnimationMode = "frames"       // frame count
	GstVideoTestSrcAnimationModeWallTime    GstVideoTestSrcAnimationMode = "wall-time"    // wall clock time
	GstVideoTestSrcAnimationModeRunningTime GstVideoTestSrcAnimationMode = "running-time" // running time
)

type GstVideoTestSrcMotionType string

const (
	GstVideoTestSrcMotionTypeWavy   GstVideoTestSrcMotionType = "wavy"   // Ball waves back and forth, up and down
	GstVideoTestSrcMotionTypeSweep  GstVideoTestSrcMotionType = "sweep"  // 1 revolution per second
	GstVideoTestSrcMotionTypeHsweep GstVideoTestSrcMotionType = "hsweep" // 1/2 revolution per second, then reset to top
)

type GstVideoTestSrcPattern string

const (
	GstVideoTestSrcPatternSmpte           GstVideoTestSrcPattern = "smpte"             // SMPTE 100%% color bars
	GstVideoTestSrcPatternSnow            GstVideoTestSrcPattern = "snow"              // Random (television snow)
	GstVideoTestSrcPatternBlack           GstVideoTestSrcPattern = "black"             // 100%% Black
	GstVideoTestSrcPatternWhite           GstVideoTestSrcPattern = "white"             // 100%% White
	GstVideoTestSrcPatternRed             GstVideoTestSrcPattern = "red"               // Red
	GstVideoTestSrcPatternGreen           GstVideoTestSrcPattern = "green"             // Green
	GstVideoTestSrcPatternBlue            GstVideoTestSrcPattern = "blue"              // Blue
	GstVideoTestSrcPatternCheckers1       GstVideoTestSrcPattern = "checkers-1"        // Checkers 1px
	GstVideoTestSrcPatternCheckers2       GstVideoTestSrcPattern = "checkers-2"        // Checkers 2px
	GstVideoTestSrcPatternCheckers4       GstVideoTestSrcPattern = "checkers-4"        // Checkers 4px
	GstVideoTestSrcPatternCheckers8       GstVideoTestSrcPattern = "checkers-8"        // Checkers 8px
	GstVideoTestSrcPatternCircular        GstVideoTestSrcPattern = "circular"          // Circular
	GstVideoTestSrcPatternBlink           GstVideoTestSrcPattern = "blink"             // Blink
	GstVideoTestSrcPatternSmpte75         GstVideoTestSrcPattern = "smpte75"           // SMPTE 75%% color bars
	GstVideoTestSrcPatternZonePlate       GstVideoTestSrcPattern = "zone-plate"        // Zone plate
	GstVideoTestSrcPatternGamut           GstVideoTestSrcPattern = "gamut"             // Gamut checkers
	GstVideoTestSrcPatternChromaZonePlate GstVideoTestSrcPattern = "chroma-zone-plate" // Chroma zone plate
	GstVideoTestSrcPatternSolidColor      GstVideoTestSrcPattern = "solid-color"       // Solid color
	GstVideoTestSrcPatternBall            GstVideoTestSrcPattern = "ball"              // Moving ball
	GstVideoTestSrcPatternSmpte100        GstVideoTestSrcPattern = "smpte100"          // SMPTE 100%% color bars
	GstVideoTestSrcPatternBar             GstVideoTestSrcPattern = "bar"               // Bar
	GstVideoTestSrcPatternPinwheel        GstVideoTestSrcPattern = "pinwheel"          // Pinwheel
	GstVideoTestSrcPatternSpokes          GstVideoTestSrcPattern = "spokes"            // Spokes
	GstVideoTestSrcPatternGradient        GstVideoTestSrcPattern = "gradient"          // Gradient
	GstVideoTestSrcPatternColors          GstVideoTestSrcPattern = "colors"            // Colors
	GstVideoTestSrcPatternSmpteRp219      GstVideoTestSrcPattern = "smpte-rp-219"      // SMPTE test pattern, RP 219 conformant
)
