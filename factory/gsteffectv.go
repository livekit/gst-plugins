// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Apply edge detect on video
type GstEdgeTV struct {
	*GstVideoFilter
}

func NewEdgeTV() (*GstEdgeTV, error) {
	e, err := gst.NewElement("edgetv")
	if err != nil {
		return nil, err
	}
	return &GstEdgeTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewEdgeTVWithName(name string) (*GstEdgeTV, error) {
	e, err := gst.NewElementWithName("edgetv", name)
	if err != nil {
		return nil, err
	}
	return &GstEdgeTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Motion dissolver
type GstQuarkTV struct {
	*GstVideoFilter
}

func NewQuarkTV() (*GstQuarkTV, error) {
	e, err := gst.NewElement("quarktv")
	if err != nil {
		return nil, err
	}
	return &GstQuarkTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewQuarkTVWithName(name string) (*GstQuarkTV, error) {
	e, err := gst.NewElementWithName("quarktv", name)
	if err != nil {
		return nil, err
	}
	return &GstQuarkTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Number of planes
// Default: 16, Min: 1, Max: 64
func (e *GstQuarkTV) SetPlanes(value int) error {
	return e.Element.SetProperty("planes", value)
}

func (e *GstQuarkTV) GetPlanes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("planes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// motion-enlightment effect
type GstRadioacTV struct {
	*GstVideoFilter
}

func NewRadioacTV() (*GstRadioacTV, error) {
	e, err := gst.NewElement("radioactv")
	if err != nil {
		return nil, err
	}
	return &GstRadioacTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewRadioacTVWithName(name string) (*GstRadioacTV, error) {
	e, err := gst.NewElementWithName("radioactv", name)
	if err != nil {
		return nil, err
	}
	return &GstRadioacTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Snapshot interval (in strobe mode)
// Default: 3, Min: 0, Max: 2147483647
func (e *GstRadioacTV) SetInterval(value uint) error {
	return e.Element.SetProperty("interval", value)
}

func (e *GstRadioacTV) GetInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Mode
// Default: normal (0)
func (e *GstRadioacTV) SetMode(value GstRadioacTVMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstRadioacTV) GetMode() (GstRadioacTVMode, error) {
	var value GstRadioacTVMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRadioacTVMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRadioacTVMode")
	}
	return value, nil
}

// Trigger (in trigger mode)
// Default: false
func (e *GstRadioacTV) SetTrigger(value bool) error {
	return e.Element.SetProperty("trigger", value)
}

func (e *GstRadioacTV) GetTrigger() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("trigger")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Color
// Default: white (3)
func (e *GstRadioacTV) SetColor(value GstRadioacTVColor) error {
	e.Element.SetArg("color", string(value))
	return nil
}

func (e *GstRadioacTV) GetColor() (GstRadioacTVColor, error) {
	var value GstRadioacTVColor
	var ok bool
	v, err := e.Element.GetProperty("color")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRadioacTVColor)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRadioacTVColor")
	}
	return value, nil
}

// A video waveform monitor for each line of video processed
type GstRevTV struct {
	*GstVideoFilter
}

func NewRevTV() (*GstRevTV, error) {
	e, err := gst.NewElement("revtv")
	if err != nil {
		return nil, err
	}
	return &GstRevTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewRevTVWithName(name string) (*GstRevTV, error) {
	e, err := gst.NewElementWithName("revtv", name)
	if err != nil {
		return nil, err
	}
	return &GstRevTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Delay in frames between updates
// Default: 1, Min: 1, Max: 100
func (e *GstRevTV) SetDelay(value int) error {
	return e.Element.SetProperty("delay", value)
}

func (e *GstRevTV) GetDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Control gain
// Default: 50, Min: 1, Max: 200
func (e *GstRevTV) SetGain(value int) error {
	return e.Element.SetProperty("gain", value)
}

func (e *GstRevTV) GetGain() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Control line spacing
// Default: 6, Min: 1, Max: 100
func (e *GstRevTV) SetLinespace(value int) error {
	return e.Element.SetProperty("linespace", value)
}

func (e *GstRevTV) GetLinespace() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("linespace")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// RippleTV does ripple mark effect on the video input
type GstRippleTV struct {
	*GstVideoFilter
}

func NewRippleTV() (*GstRippleTV, error) {
	e, err := gst.NewElement("rippletv")
	if err != nil {
		return nil, err
	}
	return &GstRippleTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewRippleTVWithName(name string) (*GstRippleTV, error) {
	e, err := gst.NewElementWithName("rippletv", name)
	if err != nil {
		return nil, err
	}
	return &GstRippleTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Mode
// Default: motion-detection (0)
func (e *GstRippleTV) SetMode(value GstRippleTVMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstRippleTV) GetMode() (GstRippleTVMode, error) {
	var value GstRippleTVMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRippleTVMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRippleTVMode")
	}
	return value, nil
}

// Reset all current ripples
// Default: false
func (e *GstRippleTV) SetReset(value bool) error {
	return e.Element.SetProperty("reset", value)
}

// Oh behave, ShagedelicTV makes images shagadelic!
type GstShagadelicTV struct {
	*GstVideoFilter
}

func NewShagadelicTV() (*GstShagadelicTV, error) {
	e, err := gst.NewElement("shagadelictv")
	if err != nil {
		return nil, err
	}
	return &GstShagadelicTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewShagadelicTVWithName(name string) (*GstShagadelicTV, error) {
	e, err := gst.NewElementWithName("shagadelictv", name)
	if err != nil {
		return nil, err
	}
	return &GstShagadelicTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// AgingTV adds age to video input using scratches and dust
type GstAgingTV struct {
	*GstVideoFilter
}

func NewAgingTV() (*GstAgingTV, error) {
	e, err := gst.NewElement("agingtv")
	if err != nil {
		return nil, err
	}
	return &GstAgingTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAgingTVWithName(name string) (*GstAgingTV, error) {
	e, err := gst.NewElementWithName("agingtv", name)
	if err != nil {
		return nil, err
	}
	return &GstAgingTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Color Aging
// Default: true
func (e *GstAgingTV) SetColorAging(value bool) error {
	return e.Element.SetProperty("color-aging", value)
}

func (e *GstAgingTV) GetColorAging() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("color-aging")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Dusts
// Default: true
func (e *GstAgingTV) SetDusts(value bool) error {
	return e.Element.SetProperty("dusts", value)
}

func (e *GstAgingTV) GetDusts() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dusts")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Pits
// Default: true
func (e *GstAgingTV) SetPits(value bool) error {
	return e.Element.SetProperty("pits", value)
}

func (e *GstAgingTV) GetPits() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pits")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of scratch lines
// Default: 7, Min: 0, Max: 20
func (e *GstAgingTV) SetScratchLines(value uint) error {
	return e.Element.SetProperty("scratch-lines", value)
}

func (e *GstAgingTV) GetScratchLines() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("scratch-lines")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// 'Dices' the screen up into many small squares
type GstDiceTV struct {
	*GstVideoFilter
}

func NewDiceTV() (*GstDiceTV, error) {
	e, err := gst.NewElement("dicetv")
	if err != nil {
		return nil, err
	}
	return &GstDiceTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewDiceTVWithName(name string) (*GstDiceTV, error) {
	e, err := gst.NewElementWithName("dicetv", name)
	if err != nil {
		return nil, err
	}
	return &GstDiceTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The size of the Squares
// Default: 4, Min: 0, Max: 5
func (e *GstDiceTV) SetSquareBits(value int) error {
	return e.Element.SetProperty("square-bits", value)
}

func (e *GstDiceTV) GetSquareBits() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("square-bits")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// A loopback alpha blending effector with rotating and scaling
type GstVertigoTV struct {
	*GstVideoFilter
}

func NewVertigoTV() (*GstVertigoTV, error) {
	e, err := gst.NewElement("vertigotv")
	if err != nil {
		return nil, err
	}
	return &GstVertigoTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVertigoTVWithName(name string) (*GstVertigoTV, error) {
	e, err := gst.NewElementWithName("vertigotv", name)
	if err != nil {
		return nil, err
	}
	return &GstVertigoTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Control the speed of movement
// Default: 0.02, Min: 0.01, Max: 100
func (e *GstVertigoTV) SetSpeed(value float32) error {
	return e.Element.SetProperty("speed", value)
}

func (e *GstVertigoTV) GetSpeed() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Control the rate of zooming
// Default: 1.01, Min: 1.01, Max: 1.1
func (e *GstVertigoTV) SetZoomSpeed(value float32) error {
	return e.Element.SetProperty("zoom-speed", value)
}

func (e *GstVertigoTV) GetZoomSpeed() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("zoom-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// WarpTV does realtime goo'ing of the video input
type GstWarpTV struct {
	*GstVideoFilter
}

func NewWarpTV() (*GstWarpTV, error) {
	e, err := gst.NewElement("warptv")
	if err != nil {
		return nil, err
	}
	return &GstWarpTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewWarpTVWithName(name string) (*GstWarpTV, error) {
	e, err := gst.NewElementWithName("warptv", name)
	if err != nil {
		return nil, err
	}
	return &GstWarpTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Optical art meets real-time video effect
type GstOpTV struct {
	*GstVideoFilter
}

func NewOpTV() (*GstOpTV, error) {
	e, err := gst.NewElement("optv")
	if err != nil {
		return nil, err
	}
	return &GstOpTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewOpTVWithName(name string) (*GstOpTV, error) {
	e, err := gst.NewElementWithName("optv", name)
	if err != nil {
		return nil, err
	}
	return &GstOpTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Mode
// Default: maelstrom (0)
func (e *GstOpTV) SetMode(value GstOpTVMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstOpTV) GetMode() (GstOpTVMode, error) {
	var value GstOpTVMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpTVMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpTVMode")
	}
	return value, nil
}

// Effect speed
// Default: 16, Min: -2147483648, Max: 2147483647
func (e *GstOpTV) SetSpeed(value int) error {
	return e.Element.SetProperty("speed", value)
}

func (e *GstOpTV) GetSpeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Luma threshold
// Default: 60, Min: 0, Max: 2147483647
func (e *GstOpTV) SetThreshold(value uint) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstOpTV) GetThreshold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// StreakTV makes after images of moving objects
type GstStreakTV struct {
	*GstVideoFilter
}

func NewStreakTV() (*GstStreakTV, error) {
	e, err := gst.NewElement("streaktv")
	if err != nil {
		return nil, err
	}
	return &GstStreakTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewStreakTVWithName(name string) (*GstStreakTV, error) {
	e, err := gst.NewElementWithName("streaktv", name)
	if err != nil {
		return nil, err
	}
	return &GstStreakTV{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Feedback
// Default: false
func (e *GstStreakTV) SetFeedback(value bool) error {
	return e.Element.SetProperty("feedback", value)
}

func (e *GstStreakTV) GetFeedback() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("feedback")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstOpTVMode string

const (
	GstOpTVModeMaelstrom         GstOpTVMode = "maelstrom"          // Maelstrom
	GstOpTVModeRadiation         GstOpTVMode = "radiation"          // Radiation
	GstOpTVModeHorizontalStripes GstOpTVMode = "horizontal-stripes" // Horizontal Stripes
	GstOpTVModeVerticalStripes   GstOpTVMode = "vertical-stripes"   // Vertical Stripes
)

type GstRadioacTVColor string

const (
	GstRadioacTVColorRed   GstRadioacTVColor = "red"   // Red
	GstRadioacTVColorGreen GstRadioacTVColor = "green" // Green
	GstRadioacTVColorBlue  GstRadioacTVColor = "blue"  // Blue
	GstRadioacTVColorWhite GstRadioacTVColor = "white" // White
)

type GstRadioacTVMode string

const (
	GstRadioacTVModeNormal  GstRadioacTVMode = "normal"  // Normal
	GstRadioacTVModeStrobe1 GstRadioacTVMode = "strobe1" // Strobe 1
	GstRadioacTVModeStrobe2 GstRadioacTVMode = "strobe2" // Strobe 2
	GstRadioacTVModeTrigger GstRadioacTVMode = "trigger" // Trigger
)

type GstRippleTVMode string

const (
	GstRippleTVModeMotionDetection GstRippleTVMode = "motion-detection" // Motion Detection
	GstRippleTVModeRain            GstRippleTVMode = "rain"             // Rain
)
