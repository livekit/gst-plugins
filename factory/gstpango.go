// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Adds text strings on top of a video buffer
type GstTextOverlay struct {
	*GstBaseTextOverlay
}

func NewTextOverlay() (*GstTextOverlay, error) {
	e, err := gst.NewElement("textoverlay")
	if err != nil {
		return nil, err
	}
	return &GstTextOverlay{GstBaseTextOverlay: &GstBaseTextOverlay{Element: e}}, nil
}

func NewTextOverlayWithName(name string) (*GstTextOverlay, error) {
	e, err := gst.NewElementWithName("textoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstTextOverlay{GstBaseTextOverlay: &GstBaseTextOverlay{Element: e}}, nil
}

// Renders a text string to an image bitmap
type GstTextRender struct {
	*gst.Element
}

func NewTextRender() (*GstTextRender, error) {
	e, err := gst.NewElement("textrender")
	if err != nil {
		return nil, err
	}
	return &GstTextRender{Element: e}, nil
}

func NewTextRenderWithName(name string) (*GstTextRender, error) {
	e, err := gst.NewElementWithName("textrender", name)
	if err != nil {
		return nil, err
	}
	return &GstTextRender{Element: e}, nil
}

// Pango font description of font to be used for rendering. See documentation of pango_font_description_from_string for syntax.

func (e *GstTextRender) SetFontDesc(value string) error {
	return e.Element.SetProperty("font-desc", value)
}

// Horizontal alignment of the text
// Default: center (1)
func (e *GstTextRender) SetHalignment(value GstTextRenderHAlign) error {
	e.Element.SetArg("halignment", string(value))
	return nil
}

func (e *GstTextRender) GetHalignment() (GstTextRenderHAlign, error) {
	var value GstTextRenderHAlign
	var ok bool
	v, err := e.Element.GetProperty("halignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTextRenderHAlign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTextRenderHAlign")
	}
	return value, nil
}

// Alignment of text lines relative to each other.
// Default: center (1)
func (e *GstTextRender) SetLineAlignment(value GstTextRenderLineAlign) error {
	e.Element.SetArg("line-alignment", string(value))
	return nil
}

func (e *GstTextRender) GetLineAlignment() (GstTextRenderLineAlign, error) {
	var value GstTextRenderLineAlign
	var ok bool
	v, err := e.Element.GetProperty("line-alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTextRenderLineAlign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTextRenderLineAlign")
	}
	return value, nil
}

// Vertical alignment of the text
// Default: baseline (0)
func (e *GstTextRender) SetValignment(value GstTextRenderVAlign) error {
	e.Element.SetArg("valignment", string(value))
	return nil
}

func (e *GstTextRender) GetValignment() (GstTextRenderVAlign, error) {
	var value GstTextRenderVAlign
	var ok bool
	v, err := e.Element.GetProperty("valignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTextRenderVAlign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTextRenderVAlign")
	}
	return value, nil
}

// Horizontal paddding when using left/right alignment
// Default: 25, Min: 0, Max: 2147483647
func (e *GstTextRender) SetXpad(value int) error {
	return e.Element.SetProperty("xpad", value)
}

func (e *GstTextRender) GetXpad() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("xpad")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Vertical padding when using top/bottom alignment
// Default: 25, Min: 0, Max: 2147483647
func (e *GstTextRender) SetYpad(value int) error {
	return e.Element.SetProperty("ypad", value)
}

func (e *GstTextRender) GetYpad() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ypad")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Overlays buffer time stamps on a video stream
type GstTimeOverlay struct {
	*GstBaseTextOverlay
}

func NewTimeOverlay() (*GstTimeOverlay, error) {
	e, err := gst.NewElement("timeoverlay")
	if err != nil {
		return nil, err
	}
	return &GstTimeOverlay{GstBaseTextOverlay: &GstBaseTextOverlay{Element: e}}, nil
}

func NewTimeOverlayWithName(name string) (*GstTimeOverlay, error) {
	e, err := gst.NewElementWithName("timeoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstTimeOverlay{GstBaseTextOverlay: &GstBaseTextOverlay{Element: e}}, nil
}

// When showing times as dates, the initial date from which time is counted, if not specified prime epoch is used (1900-01-01)

func (e *GstTimeOverlay) SetDatetimeEpoch(value interface{}) error {
	return e.Element.SetProperty("datetime-epoch", value)
}

func (e *GstTimeOverlay) GetDatetimeEpoch() (interface{}, error) {
	return e.Element.GetProperty("datetime-epoch")
}

// When showing times as dates, the format to render date and time in
// Default: %%F %%T
func (e *GstTimeOverlay) SetDatetimeFormat(value string) error {
	return e.Element.SetProperty("datetime-format", value)
}

func (e *GstTimeOverlay) GetDatetimeFormat() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("datetime-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Caps to use for the reference timestamp time mode
// Default: timestamp/x-ntp
func (e *GstTimeOverlay) SetReferenceTimestampCaps(value *gst.Caps) error {
	return e.Element.SetProperty("reference-timestamp-caps", value)
}

func (e *GstTimeOverlay) GetReferenceTimestampCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("reference-timestamp-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Whether to display times, counted from datetime-epoch, as dates
// Default: false
func (e *GstTimeOverlay) SetShowTimesAsDates(value bool) error {
	return e.Element.SetProperty("show-times-as-dates", value)
}

func (e *GstTimeOverlay) GetShowTimesAsDates() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-times-as-dates")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// What time to show
// Default: buffer-time (0)
func (e *GstTimeOverlay) SetTimeMode(value GstTimeOverlayTimeLine) error {
	e.Element.SetArg("time-mode", string(value))
	return nil
}

func (e *GstTimeOverlay) GetTimeMode() (GstTimeOverlayTimeLine, error) {
	var value GstTimeOverlayTimeLine
	var ok bool
	v, err := e.Element.GetProperty("time-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTimeOverlayTimeLine)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTimeOverlayTimeLine")
	}
	return value, nil
}

// Overlays the current clock time on a video stream
type GstClockOverlay struct {
	*GstBaseTextOverlay
}

func NewClockOverlay() (*GstClockOverlay, error) {
	e, err := gst.NewElement("clockoverlay")
	if err != nil {
		return nil, err
	}
	return &GstClockOverlay{GstBaseTextOverlay: &GstBaseTextOverlay{Element: e}}, nil
}

func NewClockOverlayWithName(name string) (*GstClockOverlay, error) {
	e, err := gst.NewElementWithName("clockoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstClockOverlay{GstBaseTextOverlay: &GstBaseTextOverlay{Element: e}}, nil
}

// Format to use for time and date value, as in strftime.
// Default: %%H:%%M:%%S
func (e *GstClockOverlay) SetTimeFormat(value string) error {
	return e.Element.SetProperty("time-format", value)
}

func (e *GstClockOverlay) GetTimeFormat() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("time-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstBaseTextOverlay struct {
	*gst.Element
}

// Resulting height of font rendering
// Default: 1, Min: 0, Max: 2147483647
func (e *GstBaseTextOverlay) GetTextHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("text-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Resulting X position of font rendering.
// Default: 0, Min: -2147483647, Max: 2147483647
func (e *GstBaseTextOverlay) GetTextX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("text-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Shift Y position up or down. Unit is pixels.
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstBaseTextOverlay) SetDeltay(value int) error {
	return e.Element.SetProperty("deltay", value)
}

func (e *GstBaseTextOverlay) GetDeltay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("deltay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Whether to draw outline
// Default: true
func (e *GstBaseTextOverlay) SetDrawOutline(value bool) error {
	return e.Element.SetProperty("draw-outline", value)
}

func (e *GstBaseTextOverlay) GetDrawOutline() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-outline")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Vertical position when using clamped position alignment
// Default: 0.5, Min: 0, Max: 1
func (e *GstBaseTextOverlay) SetYpos(value float64) error {
	return e.Element.SetProperty("ypos", value)
}

func (e *GstBaseTextOverlay) GetYpos() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("ypos")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Whether to draw shadow
// Default: true
func (e *GstBaseTextOverlay) SetDrawShadow(value bool) error {
	return e.Element.SetProperty("draw-shadow", value)
}

func (e *GstBaseTextOverlay) GetDrawShadow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-shadow")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Alignment of text lines relative to each other.
// Default: center (1)
func (e *GstBaseTextOverlay) SetLineAlignment(value GstBaseTextOverlayLineAlign) error {
	e.Element.SetArg("line-alignment", string(value))
	return nil
}

func (e *GstBaseTextOverlay) GetLineAlignment() (GstBaseTextOverlayLineAlign, error) {
	var value GstBaseTextOverlayLineAlign
	var ok bool
	v, err := e.Element.GetProperty("line-alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstBaseTextOverlayLineAlign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstBaseTextOverlayLineAlign")
	}
	return value, nil
}

// Text to be display.

func (e *GstBaseTextOverlay) SetText(value string) error {
	return e.Element.SetProperty("text", value)
}

func (e *GstBaseTextOverlay) GetText() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("text")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Horizontal position when using absolute alignment
// Default: 0.5, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstBaseTextOverlay) SetXAbsolute(value float64) error {
	return e.Element.SetProperty("x-absolute", value)
}

func (e *GstBaseTextOverlay) GetXAbsolute() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-absolute")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Automatically adjust font size to screen-size.
// Default: true
func (e *GstBaseTextOverlay) SetAutoResize(value bool) error {
	return e.Element.SetProperty("auto-resize", value)
}

func (e *GstBaseTextOverlay) GetAutoResize() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-resize")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Color to use for outline the text (big-endian ARGB).
// Default: -16777216, Min: 0, Max: -1
func (e *GstBaseTextOverlay) SetOutlineColor(value uint) error {
	return e.Element.SetProperty("outline-color", value)
}

func (e *GstBaseTextOverlay) GetOutlineColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("outline-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Whether to shade the background under the text area
// Default: false
func (e *GstBaseTextOverlay) SetShadedBackground(value bool) error {
	return e.Element.SetProperty("shaded-background", value)
}

func (e *GstBaseTextOverlay) GetShadedBackground() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("shaded-background")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to wait for subtitles
// Default: true
func (e *GstBaseTextOverlay) SetWaitText(value bool) error {
	return e.Element.SetProperty("wait-text", value)
}

func (e *GstBaseTextOverlay) GetWaitText() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-text")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to wrap the text and if so how.
// Default: wordchar (2)
func (e *GstBaseTextOverlay) SetWrapMode(value GstBaseTextOverlayWrapMode) error {
	e.Element.SetArg("wrap-mode", string(value))
	return nil
}

func (e *GstBaseTextOverlay) GetWrapMode() (GstBaseTextOverlayWrapMode, error) {
	var value GstBaseTextOverlayWrapMode
	var ok bool
	v, err := e.Element.GetProperty("wrap-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstBaseTextOverlayWrapMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstBaseTextOverlayWrapMode")
	}
	return value, nil
}

// Scale text to compensate for and avoid distortion by subsequent video scaling.
// Default: none (0)
func (e *GstBaseTextOverlay) SetScaleMode(value GstBaseTextOverlayScaleMode) error {
	e.Element.SetArg("scale-mode", string(value))
	return nil
}

func (e *GstBaseTextOverlay) GetScaleMode() (GstBaseTextOverlayScaleMode, error) {
	var value GstBaseTextOverlayScaleMode
	var ok bool
	v, err := e.Element.GetProperty("scale-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstBaseTextOverlayScaleMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstBaseTextOverlayScaleMode")
	}
	return value, nil
}

// Pixel aspect ratio of video scale to compensate for in user scale-mode
// Default: 1/1, Min: 1/100, Max: 100/1
func (e *GstBaseTextOverlay) SetScalePixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("scale-pixel-aspect-ratio", value)
}

func (e *GstBaseTextOverlay) GetScalePixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("scale-pixel-aspect-ratio")
}

// Vertical padding when using top/bottom alignment
// Default: 25, Min: 0, Max: 2147483647
func (e *GstBaseTextOverlay) SetYpad(value int) error {
	return e.Element.SetProperty("ypad", value)
}

func (e *GstBaseTextOverlay) GetYpad() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ypad")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Color to use for text (big-endian ARGB).
// Default: -1, Min: 0, Max: -1
func (e *GstBaseTextOverlay) SetColor(value uint) error {
	return e.Element.SetProperty("color", value)
}

func (e *GstBaseTextOverlay) GetColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Shading value to apply if shaded-background is true
// Default: 80, Min: 1, Max: 255
func (e *GstBaseTextOverlay) SetShadingValue(value uint) error {
	return e.Element.SetProperty("shading-value", value)
}

func (e *GstBaseTextOverlay) GetShadingValue() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shading-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Horizontal position when using clamped position alignment
// Default: 0.5, Min: 0, Max: 1
func (e *GstBaseTextOverlay) SetXpos(value float64) error {
	return e.Element.SetProperty("xpos", value)
}

func (e *GstBaseTextOverlay) GetXpos() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("xpos")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Vertical Render.
// Default: false
func (e *GstBaseTextOverlay) SetVerticalRender(value bool) error {
	return e.Element.SetProperty("vertical-render", value)
}

func (e *GstBaseTextOverlay) GetVerticalRender() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vertical-render")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Vertical position when using absolute alignment
// Default: 0.5, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstBaseTextOverlay) SetYAbsolute(value float64) error {
	return e.Element.SetProperty("y-absolute", value)
}

func (e *GstBaseTextOverlay) GetYAbsolute() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-absolute")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Horizontal alignment of the text
// Default: center (1)
func (e *GstBaseTextOverlay) SetHalignment(value GstBaseTextOverlayHAlign) error {
	e.Element.SetArg("halignment", string(value))
	return nil
}

func (e *GstBaseTextOverlay) GetHalignment() (GstBaseTextOverlayHAlign, error) {
	var value GstBaseTextOverlayHAlign
	var ok bool
	v, err := e.Element.GetProperty("halignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstBaseTextOverlayHAlign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstBaseTextOverlayHAlign")
	}
	return value, nil
}

// Whether to render the text string
// Default: false
func (e *GstBaseTextOverlay) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstBaseTextOverlay) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Resulting width of font rendering
// Default: 1, Min: 0, Max: 2147483647
func (e *GstBaseTextOverlay) GetTextWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("text-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Resulting Y position of font rendering.
// Default: 0, Min: -2147483647, Max: 2147483647
func (e *GstBaseTextOverlay) GetTextY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("text-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Vertical alignment of the text
// Default: baseline (0)
func (e *GstBaseTextOverlay) SetValignment(value GstBaseTextOverlayVAlign) error {
	e.Element.SetArg("valignment", string(value))
	return nil
}

func (e *GstBaseTextOverlay) GetValignment() (GstBaseTextOverlayVAlign, error) {
	var value GstBaseTextOverlayVAlign
	var ok bool
	v, err := e.Element.GetProperty("valignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstBaseTextOverlayVAlign)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstBaseTextOverlayVAlign")
	}
	return value, nil
}

// Shift X position to the left or to the right. Unit is pixels.
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstBaseTextOverlay) SetDeltax(value int) error {
	return e.Element.SetProperty("deltax", value)
}

func (e *GstBaseTextOverlay) GetDeltax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("deltax")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Pango font description of font to be used for rendering. See documentation of pango_font_description_from_string for syntax.

func (e *GstBaseTextOverlay) SetFontDesc(value string) error {
	return e.Element.SetProperty("font-desc", value)
}

func (e *GstBaseTextOverlay) GetFontDesc() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-desc")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Horizontal paddding when using left/right alignment
// Default: 25, Min: 0, Max: 2147483647
func (e *GstBaseTextOverlay) SetXpad(value int) error {
	return e.Element.SetProperty("xpad", value)
}

func (e *GstBaseTextOverlay) GetXpad() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("xpad")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstBaseTextOverlayLineAlign string

const (
	GstBaseTextOverlayLineAlignLeft   GstBaseTextOverlayLineAlign = "left"   // left
	GstBaseTextOverlayLineAlignCenter GstBaseTextOverlayLineAlign = "center" // center
	GstBaseTextOverlayLineAlignRight  GstBaseTextOverlayLineAlign = "right"  // right
)

type GstBaseTextOverlayVAlign string

const (
	GstBaseTextOverlayVAlignBaseline                        GstBaseTextOverlayVAlign = "baseline"                            // baseline
	GstBaseTextOverlayVAlignBottom                          GstBaseTextOverlayVAlign = "bottom"                              // bottom
	GstBaseTextOverlayVAlignTop                             GstBaseTextOverlayVAlign = "top"                                 // top
	GstBaseTextOverlayVAlignAbsolutePositionClampedToCanvas GstBaseTextOverlayVAlign = "Absolute position clamped to canvas" // position
	GstBaseTextOverlayVAlignCenter                          GstBaseTextOverlayVAlign = "center"                              // center
	GstBaseTextOverlayVAlignAbsolutePosition                GstBaseTextOverlayVAlign = "Absolute position"                   // absolute
)

type GstTextRenderHAlign string

const (
	GstTextRenderHAlignLeft   GstTextRenderHAlign = "left"   // left
	GstTextRenderHAlignCenter GstTextRenderHAlign = "center" // center
	GstTextRenderHAlignRight  GstTextRenderHAlign = "right"  // right
)

type GstTimeOverlayTimeLine string

const (
	GstTimeOverlayTimeLineBufferTime         GstTimeOverlayTimeLine = "buffer-time"          // buffer-time
	GstTimeOverlayTimeLineStreamTime         GstTimeOverlayTimeLine = "stream-time"          // stream-time
	GstTimeOverlayTimeLineRunningTime        GstTimeOverlayTimeLine = "running-time"         // running-time
	GstTimeOverlayTimeLineTimeCode           GstTimeOverlayTimeLine = "time-code"            // time-code
	GstTimeOverlayTimeLineElapsedRunningTime GstTimeOverlayTimeLine = "elapsed-running-time" // elapsed-running-time
	GstTimeOverlayTimeLineReferenceTimestamp GstTimeOverlayTimeLine = "reference-timestamp"  // reference-timestamp
	GstTimeOverlayTimeLineBufferCount        GstTimeOverlayTimeLine = "buffer-count"         // buffer-count
	GstTimeOverlayTimeLineBufferOffset       GstTimeOverlayTimeLine = "buffer-offset"        // buffer-offset
)

type GstBaseTextOverlayHAlign string

const (
	GstBaseTextOverlayHAlignLeft                            GstBaseTextOverlayHAlign = "left"                                // left
	GstBaseTextOverlayHAlignCenter                          GstBaseTextOverlayHAlign = "center"                              // center
	GstBaseTextOverlayHAlignRight                           GstBaseTextOverlayHAlign = "right"                               // right
	GstBaseTextOverlayHAlignAbsolutePositionClampedToCanvas GstBaseTextOverlayHAlign = "Absolute position clamped to canvas" // position
	GstBaseTextOverlayHAlignAbsolutePosition                GstBaseTextOverlayHAlign = "Absolute position"                   // absolute
)

type GstBaseTextOverlayScaleMode string

const (
	GstBaseTextOverlayScaleModeNone    GstBaseTextOverlayScaleMode = "none"    // none
	GstBaseTextOverlayScaleModePar     GstBaseTextOverlayScaleMode = "par"     // par
	GstBaseTextOverlayScaleModeDisplay GstBaseTextOverlayScaleMode = "display" // display
	GstBaseTextOverlayScaleModeUser    GstBaseTextOverlayScaleMode = "user"    // user
)

type GstBaseTextOverlayWrapMode string

const (
	GstBaseTextOverlayWrapModeNone     GstBaseTextOverlayWrapMode = "none"     // none
	GstBaseTextOverlayWrapModeWord     GstBaseTextOverlayWrapMode = "word"     // word
	GstBaseTextOverlayWrapModeChar     GstBaseTextOverlayWrapMode = "char"     // char
	GstBaseTextOverlayWrapModeWordchar GstBaseTextOverlayWrapMode = "wordchar" // wordchar
)

type GstTextRenderLineAlign string

const (
	GstTextRenderLineAlignLeft   GstTextRenderLineAlign = "left"   // left
	GstTextRenderLineAlignCenter GstTextRenderLineAlign = "center" // center
	GstTextRenderLineAlignRight  GstTextRenderLineAlign = "right"  // right
)

type GstTextRenderVAlign string

const (
	GstTextRenderVAlignBaseline GstTextRenderVAlign = "baseline" // baseline
	GstTextRenderVAlignBottom   GstTextRenderVAlign = "bottom"   // bottom
	GstTextRenderVAlignTop      GstTextRenderVAlign = "top"      // top
)
