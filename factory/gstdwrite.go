// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Overlays the current clock time on a video stream
type GstDWriteClockOverlay struct {
	*GstDWriteBaseOverlay
}

func NewDWriteClockOverlay() (*GstDWriteClockOverlay, error) {
	e, err := gst.NewElement("dwriteclockoverlay")
	if err != nil {
		return nil, err
	}
	return &GstDWriteClockOverlay{GstDWriteBaseOverlay: &GstDWriteBaseOverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewDWriteClockOverlayWithName(name string) (*GstDWriteClockOverlay, error) {
	e, err := gst.NewElementWithName("dwriteclockoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstDWriteClockOverlay{GstDWriteBaseOverlay: &GstDWriteBaseOverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Format to use for time and date value, as in strftime.
// Default: %%H:%%M:%%S
func (e *GstDWriteClockOverlay) SetTimeFormat(value string) error {
	return e.Element.SetProperty("time-format", value)
}

func (e *GstDWriteClockOverlay) GetTimeFormat() (string, error) {
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

// Attach subtitle metas on video buffers
type GstDWriteSubtitleMux struct {
	*GstAggregator
}

func NewDWriteSubtitleMux() (*GstDWriteSubtitleMux, error) {
	e, err := gst.NewElement("dwritesubtitlemux")
	if err != nil {
		return nil, err
	}
	return &GstDWriteSubtitleMux{GstAggregator: &GstAggregator{Element: e}}, nil
}

func NewDWriteSubtitleMuxWithName(name string) (*GstDWriteSubtitleMux, error) {
	e, err := gst.NewElementWithName("dwritesubtitlemux", name)
	if err != nil {
		return nil, err
	}
	return &GstDWriteSubtitleMux{GstAggregator: &GstAggregator{Element: e}}, nil
}

// Adds subtitle strings on top of a video buffer
type GstDWriteSubtitleOverlay struct {
	*gst.Bin
}

func NewDWriteSubtitleOverlay() (*GstDWriteSubtitleOverlay, error) {
	e, err := gst.NewElement("dwritesubtitleoverlay")
	if err != nil {
		return nil, err
	}
	return &GstDWriteSubtitleOverlay{Bin: &gst.Bin{Element: e}}, nil
}

func NewDWriteSubtitleOverlayWithName(name string) (*GstDWriteSubtitleOverlay, error) {
	e, err := gst.NewElementWithName("dwritesubtitleoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstDWriteSubtitleOverlay{Bin: &gst.Bin{Element: e}}, nil
}

// Enable color font, requires Windows 10 or newer
// Default: true
func (e *GstDWriteSubtitleOverlay) SetColorFont(value bool) error {
	return e.Element.SetProperty("color-font", value)
}

func (e *GstDWriteSubtitleOverlay) GetColorFont() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("color-font")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Font Style
// Default: normal (0)
func (e *GstDWriteSubtitleOverlay) SetFontStyle(value GstDWriteFontStyle) error {
	e.Element.SetArg("font-style", string(value))
	return nil
}

func (e *GstDWriteSubtitleOverlay) GetFontStyle() (GstDWriteFontStyle, error) {
	var value GstDWriteFontStyle
	var ok bool
	v, err := e.Element.GetProperty("font-style")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteFontStyle)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteFontStyle")
	}
	return value, nil
}

// Text to render

func (e *GstDWriteSubtitleOverlay) SetText(value string) error {
	return e.Element.SetProperty("text", value)
}

func (e *GstDWriteSubtitleOverlay) GetText() (string, error) {
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

// Whether to draw text
// Default: true
func (e *GstDWriteSubtitleOverlay) SetVisible(value bool) error {
	return e.Element.SetProperty("visible", value)
}

func (e *GstDWriteSubtitleOverlay) GetVisible() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("visible")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Duration after which to erase overlay when no cc data has arrived for the selected field, in nanoseconds unit
// Default: 18446744073709551615, Min: 16000000000, Max: 18446744073709551615
func (e *GstDWriteSubtitleOverlay) SetCcTimeout(value uint64) error {
	return e.Element.SetProperty("cc-timeout", value)
}

func (e *GstDWriteSubtitleOverlay) GetCcTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("cc-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Enable closed caption rendering
// Default: true
func (e *GstDWriteSubtitleOverlay) SetEnableCc(value bool) error {
	return e.Element.SetProperty("enable-cc", value)
}

func (e *GstDWriteSubtitleOverlay) GetEnableCc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-cc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Font Weight
// Default: normal (400)
func (e *GstDWriteSubtitleOverlay) SetFontWeight(value GstDWriteFontWeight) error {
	e.Element.SetArg("font-weight", string(value))
	return nil
}

func (e *GstDWriteSubtitleOverlay) GetFontWeight() (GstDWriteFontWeight, error) {
	var value GstDWriteFontWeight
	var ok bool
	v, err := e.Element.GetProperty("font-weight")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteFontWeight)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteFontWeight")
	}
	return value, nil
}

// Normalized X coordinate of text layout
// Default: 0.03, Min: 0, Max: 1
func (e *GstDWriteSubtitleOverlay) SetLayoutX(value float64) error {
	return e.Element.SetProperty("layout-x", value)
}

func (e *GstDWriteSubtitleOverlay) GetLayoutX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Remove caption meta from output buffers when closed caption rendering is enabled
// Default: false
func (e *GstDWriteSubtitleOverlay) SetRemoveCcMeta(value bool) error {
	return e.Element.SetProperty("remove-cc-meta", value)
}

func (e *GstDWriteSubtitleOverlay) GetRemoveCcMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-cc-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Text Alignment
// Default: center (2)
func (e *GstDWriteSubtitleOverlay) SetTextAlignment(value GstDWriteTextAlignment) error {
	e.Element.SetArg("text-alignment", string(value))
	return nil
}

func (e *GstDWriteSubtitleOverlay) GetTextAlignment() (GstDWriteTextAlignment, error) {
	var value GstDWriteTextAlignment
	var ok bool
	v, err := e.Element.GetProperty("text-alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteTextAlignment)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteTextAlignment")
	}
	return value, nil
}

// The closed caption field to render when available, (-1 = automatic)
// Default: -1, Min: -1, Max: 1
func (e *GstDWriteSubtitleOverlay) SetCcField(value int) error {
	return e.Element.SetProperty("cc-field", value)
}

func (e *GstDWriteSubtitleOverlay) GetCcField() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cc-field")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Font size to use
// Default: 20, Min: 0.1, Max: 1638
func (e *GstDWriteSubtitleOverlay) SetFontSize(value float32) error {
	return e.Element.SetProperty("font-size", value)
}

func (e *GstDWriteSubtitleOverlay) GetFontSize() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("font-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Normalized height of text layout
// Default: 0.92, Min: 0, Max: 1
func (e *GstDWriteSubtitleOverlay) SetLayoutHeight(value float64) error {
	return e.Element.SetProperty("layout-height", value)
}

func (e *GstDWriteSubtitleOverlay) GetLayoutHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Normalized Y coordinate of text layout
// Default: 0.03, Min: 0, Max: 1
func (e *GstDWriteSubtitleOverlay) SetLayoutY(value float64) error {
	return e.Element.SetProperty("layout-y", value)
}

func (e *GstDWriteSubtitleOverlay) GetLayoutY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Text outline color to use (big-endian ARGB)
// Default: -16777216, Min: 0, Max: -1
func (e *GstDWriteSubtitleOverlay) SetOutlineColor(value uint) error {
	return e.Element.SetProperty("outline-color", value)
}

func (e *GstDWriteSubtitleOverlay) GetOutlineColor() (uint, error) {
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

// Paragraph Alignment
// Default: far (1)
func (e *GstDWriteSubtitleOverlay) SetParagraphAlignment(value GstDWriteParagraphAlignment) error {
	e.Element.SetArg("paragraph-alignment", string(value))
	return nil
}

func (e *GstDWriteSubtitleOverlay) GetParagraphAlignment() (GstDWriteParagraphAlignment, error) {
	var value GstDWriteParagraphAlignment
	var ok bool
	v, err := e.Element.GetProperty("paragraph-alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteParagraphAlignment)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteParagraphAlignment")
	}
	return value, nil
}

// Shadow color to use (big-endian ARGB)
// Default: -2147483648, Min: 0, Max: -1
func (e *GstDWriteSubtitleOverlay) SetShadowColor(value uint) error {
	return e.Element.SetProperty("shadow-color", value)
}

func (e *GstDWriteSubtitleOverlay) GetShadowColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shadow-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Calculate font size to be equivalent to "font-size" at "reference-frame-size"
// Default: true
func (e *GstDWriteSubtitleOverlay) SetAutoResize(value bool) error {
	return e.Element.SetProperty("auto-resize", value)
}

func (e *GstDWriteSubtitleOverlay) GetAutoResize() (bool, error) {
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

// Background color to use (big-endian ARGB)
// Default: 0, Min: 0, Max: -1
func (e *GstDWriteSubtitleOverlay) SetBackgroundColor(value uint) error {
	return e.Element.SetProperty("background-color", value)
}

func (e *GstDWriteSubtitleOverlay) GetBackgroundColor() (uint, error) {
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

// Font family to use
// Default: MS Reference Sans Serif
func (e *GstDWriteSubtitleOverlay) SetFontFamily(value string) error {
	return e.Element.SetProperty("font-family", value)
}

func (e *GstDWriteSubtitleOverlay) GetFontFamily() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-family")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Font Stretch
// Default: normal (5)
func (e *GstDWriteSubtitleOverlay) SetFontStretch(value GstDWriteFontStretch) error {
	e.Element.SetArg("font-stretch", string(value))
	return nil
}

func (e *GstDWriteSubtitleOverlay) GetFontStretch() (GstDWriteFontStretch, error) {
	var value GstDWriteFontStretch
	var ok bool
	v, err := e.Element.GetProperty("font-stretch")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteFontStretch)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteFontStretch")
	}
	return value, nil
}

// Foreground color to use (big-endian ARGB)
// Default: -1, Min: 0, Max: -1
func (e *GstDWriteSubtitleOverlay) SetForegroundColor(value uint) error {
	return e.Element.SetProperty("foreground-color", value)
}

func (e *GstDWriteSubtitleOverlay) GetForegroundColor() (uint, error) {
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

// Normalized width of text layout
// Default: 0.92, Min: 0, Max: 1
func (e *GstDWriteSubtitleOverlay) SetLayoutWidth(value float64) error {
	return e.Element.SetProperty("layout-width", value)
}

func (e *GstDWriteSubtitleOverlay) GetLayoutWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adds text strings on top of a video buffer
type GstDWriteTextOverlay struct {
	*GstDWriteBaseOverlay
}

func NewDWriteTextOverlay() (*GstDWriteTextOverlay, error) {
	e, err := gst.NewElement("dwritetextoverlay")
	if err != nil {
		return nil, err
	}
	return &GstDWriteTextOverlay{GstDWriteBaseOverlay: &GstDWriteBaseOverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewDWriteTextOverlayWithName(name string) (*GstDWriteTextOverlay, error) {
	e, err := gst.NewElementWithName("dwritetextoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstDWriteTextOverlay{GstDWriteBaseOverlay: &GstDWriteBaseOverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The closed caption field to render when available, (-1 = automatic)
// Default: -1, Min: -1, Max: 1
func (e *GstDWriteTextOverlay) SetCcField(value int) error {
	return e.Element.SetProperty("cc-field", value)
}

func (e *GstDWriteTextOverlay) GetCcField() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cc-field")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Duration after which to erase overlay when no cc data has arrived for the selected field, in nanoseconds unit
// Default: 18446744073709551615, Min: 16000000000, Max: 18446744073709551615
func (e *GstDWriteTextOverlay) SetCcTimeout(value uint64) error {
	return e.Element.SetProperty("cc-timeout", value)
}

func (e *GstDWriteTextOverlay) GetCcTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("cc-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Enable closed caption rendering
// Default: true
func (e *GstDWriteTextOverlay) SetEnableCc(value bool) error {
	return e.Element.SetProperty("enable-cc", value)
}

func (e *GstDWriteTextOverlay) GetEnableCc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-cc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Remove caption meta from output buffers when closed caption rendering is enabled
// Default: false
func (e *GstDWriteTextOverlay) SetRemoveCcMeta(value bool) error {
	return e.Element.SetProperty("remove-cc-meta", value)
}

func (e *GstDWriteTextOverlay) GetRemoveCcMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-cc-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Overlays buffer time stamps on a video stream
type GstDWriteTimeOverlay struct {
	*GstDWriteBaseOverlay
}

func NewDWriteTimeOverlay() (*GstDWriteTimeOverlay, error) {
	e, err := gst.NewElement("dwritetimeoverlay")
	if err != nil {
		return nil, err
	}
	return &GstDWriteTimeOverlay{GstDWriteBaseOverlay: &GstDWriteBaseOverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewDWriteTimeOverlayWithName(name string) (*GstDWriteTimeOverlay, error) {
	e, err := gst.NewElementWithName("dwritetimeoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstDWriteTimeOverlay{GstDWriteBaseOverlay: &GstDWriteBaseOverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// When showing times as dates, the initial date from which time is counted, if not specified prime epoch is used (1900-01-01)

func (e *GstDWriteTimeOverlay) SetDatetimeEpoch(value interface{}) error {
	return e.Element.SetProperty("datetime-epoch", value)
}

func (e *GstDWriteTimeOverlay) GetDatetimeEpoch() (interface{}, error) {
	return e.Element.GetProperty("datetime-epoch")
}

// When showing times as dates, the format to render date and time in
// Default: %%F %%T
func (e *GstDWriteTimeOverlay) SetDatetimeFormat(value string) error {
	return e.Element.SetProperty("datetime-format", value)
}

func (e *GstDWriteTimeOverlay) GetDatetimeFormat() (string, error) {
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
func (e *GstDWriteTimeOverlay) SetReferenceTimestampCaps(value *gst.Caps) error {
	return e.Element.SetProperty("reference-timestamp-caps", value)
}

func (e *GstDWriteTimeOverlay) GetReferenceTimestampCaps() (*gst.Caps, error) {
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
func (e *GstDWriteTimeOverlay) SetShowTimesAsDates(value bool) error {
	return e.Element.SetProperty("show-times-as-dates", value)
}

func (e *GstDWriteTimeOverlay) GetShowTimesAsDates() (bool, error) {
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
func (e *GstDWriteTimeOverlay) SetTimeMode(value GstDWriteTimeOverlayTimeLine) error {
	e.Element.SetArg("time-mode", string(value))
	return nil
}

func (e *GstDWriteTimeOverlay) GetTimeMode() (GstDWriteTimeOverlayTimeLine, error) {
	var value GstDWriteTimeOverlayTimeLine
	var ok bool
	v, err := e.Element.GetProperty("time-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteTimeOverlayTimeLine)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteTimeOverlayTimeLine")
	}
	return value, nil
}

type GstDWriteTimeOverlayTimeLine string

const (
	GstDWriteTimeOverlayTimeLineBufferTime         GstDWriteTimeOverlayTimeLine = "buffer-time"          // buffer-time
	GstDWriteTimeOverlayTimeLineStreamTime         GstDWriteTimeOverlayTimeLine = "stream-time"          // stream-time
	GstDWriteTimeOverlayTimeLineRunningTime        GstDWriteTimeOverlayTimeLine = "running-time"         // running-time
	GstDWriteTimeOverlayTimeLineTimeCode           GstDWriteTimeOverlayTimeLine = "time-code"            // time-code
	GstDWriteTimeOverlayTimeLineElapsedRunningTime GstDWriteTimeOverlayTimeLine = "elapsed-running-time" // elapsed-running-time
	GstDWriteTimeOverlayTimeLineReferenceTimestamp GstDWriteTimeOverlayTimeLine = "reference-timestamp"  // reference-timestamp
	GstDWriteTimeOverlayTimeLineBufferCount        GstDWriteTimeOverlayTimeLine = "buffer-count"         // buffer-count
	GstDWriteTimeOverlayTimeLineBufferOffset       GstDWriteTimeOverlayTimeLine = "buffer-offset"        // buffer-offset
)

type GstDWriteBaseOverlay struct {
	*base.GstBaseTransform
}

// Font Style
// Default: normal (0)
func (e *GstDWriteBaseOverlay) SetFontStyle(value GstDWriteFontStyle) error {
	e.Element.SetArg("font-style", string(value))
	return nil
}

func (e *GstDWriteBaseOverlay) GetFontStyle() (GstDWriteFontStyle, error) {
	var value GstDWriteFontStyle
	var ok bool
	v, err := e.Element.GetProperty("font-style")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteFontStyle)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteFontStyle")
	}
	return value, nil
}

// Normalized Y coordinate of text layout
// Default: 0.03, Min: 0, Max: 1
func (e *GstDWriteBaseOverlay) SetLayoutY(value float64) error {
	return e.Element.SetProperty("layout-y", value)
}

func (e *GstDWriteBaseOverlay) GetLayoutY() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Paragraph Alignment
// Default: near (0)
func (e *GstDWriteBaseOverlay) SetParagraphAlignment(value GstDWriteParagraphAlignment) error {
	e.Element.SetArg("paragraph-alignment", string(value))
	return nil
}

func (e *GstDWriteBaseOverlay) GetParagraphAlignment() (GstDWriteParagraphAlignment, error) {
	var value GstDWriteParagraphAlignment
	var ok bool
	v, err := e.Element.GetProperty("paragraph-alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteParagraphAlignment)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteParagraphAlignment")
	}
	return value, nil
}

// Text to render

func (e *GstDWriteBaseOverlay) SetText(value string) error {
	return e.Element.SetProperty("text", value)
}

func (e *GstDWriteBaseOverlay) GetText() (string, error) {
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

// Text Alignment
// Default: leading (0)
func (e *GstDWriteBaseOverlay) SetTextAlignment(value GstDWriteTextAlignment) error {
	e.Element.SetArg("text-alignment", string(value))
	return nil
}

func (e *GstDWriteBaseOverlay) GetTextAlignment() (GstDWriteTextAlignment, error) {
	var value GstDWriteTextAlignment
	var ok bool
	v, err := e.Element.GetProperty("text-alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteTextAlignment)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteTextAlignment")
	}
	return value, nil
}

// Font family to use
// Default: MS Reference Sans Serif
func (e *GstDWriteBaseOverlay) SetFontFamily(value string) error {
	return e.Element.SetProperty("font-family", value)
}

func (e *GstDWriteBaseOverlay) GetFontFamily() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-family")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Font Weight
// Default: normal (400)
func (e *GstDWriteBaseOverlay) SetFontWeight(value GstDWriteFontWeight) error {
	e.Element.SetArg("font-weight", string(value))
	return nil
}

func (e *GstDWriteBaseOverlay) GetFontWeight() (GstDWriteFontWeight, error) {
	var value GstDWriteFontWeight
	var ok bool
	v, err := e.Element.GetProperty("font-weight")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteFontWeight)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteFontWeight")
	}
	return value, nil
}

// Normalized width of text layout
// Default: 0.92, Min: 0, Max: 1
func (e *GstDWriteBaseOverlay) SetLayoutWidth(value float64) error {
	return e.Element.SetProperty("layout-width", value)
}

func (e *GstDWriteBaseOverlay) GetLayoutWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Normalized X coordinate of text layout
// Default: 0.03, Min: 0, Max: 1
func (e *GstDWriteBaseOverlay) SetLayoutX(value float64) error {
	return e.Element.SetProperty("layout-x", value)
}

func (e *GstDWriteBaseOverlay) GetLayoutX() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Calculate font size to be equivalent to "font-size" at "reference-frame-size"
// Default: true
func (e *GstDWriteBaseOverlay) SetAutoResize(value bool) error {
	return e.Element.SetProperty("auto-resize", value)
}

func (e *GstDWriteBaseOverlay) GetAutoResize() (bool, error) {
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

// Background color to use (big-endian ARGB)
// Default: 0, Min: 0, Max: -1
func (e *GstDWriteBaseOverlay) SetBackgroundColor(value uint) error {
	return e.Element.SetProperty("background-color", value)
}

func (e *GstDWriteBaseOverlay) GetBackgroundColor() (uint, error) {
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

// Font Stretch
// Default: normal (5)
func (e *GstDWriteBaseOverlay) SetFontStretch(value GstDWriteFontStretch) error {
	e.Element.SetArg("font-stretch", string(value))
	return nil
}

func (e *GstDWriteBaseOverlay) GetFontStretch() (GstDWriteFontStretch, error) {
	var value GstDWriteFontStretch
	var ok bool
	v, err := e.Element.GetProperty("font-stretch")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDWriteFontStretch)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDWriteFontStretch")
	}
	return value, nil
}

// Shadow color to use (big-endian ARGB)
// Default: -2147483648, Min: 0, Max: -1
func (e *GstDWriteBaseOverlay) SetShadowColor(value uint) error {
	return e.Element.SetProperty("shadow-color", value)
}

func (e *GstDWriteBaseOverlay) GetShadowColor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shadow-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Whether to draw text
// Default: true
func (e *GstDWriteBaseOverlay) SetVisible(value bool) error {
	return e.Element.SetProperty("visible", value)
}

func (e *GstDWriteBaseOverlay) GetVisible() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("visible")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable color font, requires Windows 10 or newer
// Default: true
func (e *GstDWriteBaseOverlay) SetColorFont(value bool) error {
	return e.Element.SetProperty("color-font", value)
}

func (e *GstDWriteBaseOverlay) GetColorFont() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("color-font")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Font size to use
// Default: 24, Min: 0.1, Max: 1638
func (e *GstDWriteBaseOverlay) SetFontSize(value float32) error {
	return e.Element.SetProperty("font-size", value)
}

func (e *GstDWriteBaseOverlay) GetFontSize() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("font-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Foreground color to use (big-endian ARGB)
// Default: -1, Min: 0, Max: -1
func (e *GstDWriteBaseOverlay) SetForegroundColor(value uint) error {
	return e.Element.SetProperty("foreground-color", value)
}

func (e *GstDWriteBaseOverlay) GetForegroundColor() (uint, error) {
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

// Normalized height of text layout
// Default: 0.92, Min: 0, Max: 1
func (e *GstDWriteBaseOverlay) SetLayoutHeight(value float64) error {
	return e.Element.SetProperty("layout-height", value)
}

func (e *GstDWriteBaseOverlay) GetLayoutHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("layout-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Text outline color to use (big-endian ARGB)
// Default: -16777216, Min: 0, Max: -1
func (e *GstDWriteBaseOverlay) SetOutlineColor(value uint) error {
	return e.Element.SetProperty("outline-color", value)
}

func (e *GstDWriteBaseOverlay) GetOutlineColor() (uint, error) {
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

type GstDWriteFontStretch string

const (
	GstDWriteFontStretchUndefined      GstDWriteFontStretch = "undefined"       // DWRITE_FONT_STRETCH_UNDEFINED
	GstDWriteFontStretchUltraCondensed GstDWriteFontStretch = "ultra-condensed" // DWRITE_FONT_STRETCH_ULTRA_CONDENSED
	GstDWriteFontStretchExtraCondensed GstDWriteFontStretch = "extra-condensed" // DWRITE_FONT_STRETCH_EXTRA_CONDENSED
	GstDWriteFontStretchCondensed      GstDWriteFontStretch = "condensed"       // DWRITE_FONT_STRETCH_CONDENSED
	GstDWriteFontStretchSemiCondensed  GstDWriteFontStretch = "semi-condensed"  // DWRITE_FONT_STRETCH_SEMI_CONDENSED
	GstDWriteFontStretchNormal         GstDWriteFontStretch = "normal"          // DWRITE_FONT_STRETCH_NORMAL
	GstDWriteFontStretchMedium         GstDWriteFontStretch = "medium"          // DWRITE_FONT_STRETCH_MEDIUM
	GstDWriteFontStretchSemiExpanded   GstDWriteFontStretch = "semi-expanded"   // DWRITE_FONT_STRETCH_SEMI_EXPANDED
	GstDWriteFontStretchExpanded       GstDWriteFontStretch = "expanded"        // DWRITE_FONT_STRETCH_EXPANDED
	GstDWriteFontStretchExtraExpanded  GstDWriteFontStretch = "extra-expanded"  // DWRITE_FONT_STRETCH_EXTRA_EXPANDED
	GstDWriteFontStretchUltraExpanded  GstDWriteFontStretch = "ultra-expanded"  // DWRITE_FONT_STRETCH_ULTRA_EXPANDED
)

type GstDWriteFontStyle string

const (
	GstDWriteFontStyleNormal  GstDWriteFontStyle = "normal"  // DWRITE_FONT_STYLE_NORMAL
	GstDWriteFontStyleOblique GstDWriteFontStyle = "oblique" // DWRITE_FONT_STYLE_OBLIQUE
	GstDWriteFontStyleItalic  GstDWriteFontStyle = "italic"  // DWRITE_FONT_STYLE_ITALIC
)

type GstDWriteFontWeight string

const (
	GstDWriteFontWeightThin       GstDWriteFontWeight = "thin"        // DWRITE_FONT_WEIGHT_THIN
	GstDWriteFontWeightExtraLight GstDWriteFontWeight = "extra-light" // DWRITE_FONT_WEIGHT_EXTRA_LIGHT
	GstDWriteFontWeightUltraLight GstDWriteFontWeight = "ultra-light" // DWRITE_FONT_WEIGHT_ULTRA_LIGHT
	GstDWriteFontWeightLight      GstDWriteFontWeight = "light"       // DWRITE_FONT_WEIGHT_LIGHT
	GstDWriteFontWeightSemiLight  GstDWriteFontWeight = "semi-light"  // DWRITE_FONT_WEIGHT_SEMI_LIGHT
	GstDWriteFontWeightNormal     GstDWriteFontWeight = "normal"      // DWRITE_FONT_WEIGHT_NORMAL
	GstDWriteFontWeightRegular    GstDWriteFontWeight = "regular"     // DWRITE_FONT_WEIGHT_REGULAR
	GstDWriteFontWeightMedium     GstDWriteFontWeight = "medium"      // DWRITE_FONT_WEIGHT_MEDIUM
	GstDWriteFontWeightDemiBold   GstDWriteFontWeight = "demi-bold"   // DWRITE_FONT_WEIGHT_DEMI_BOLD
	GstDWriteFontWeightSemiBold   GstDWriteFontWeight = "semi-bold"   // DWRITE_FONT_WEIGHT_SEMI_BOLD
	GstDWriteFontWeightBold       GstDWriteFontWeight = "bold"        // DWRITE_FONT_WEIGHT_BOLD
	GstDWriteFontWeightExtraBold  GstDWriteFontWeight = "extra-bold"  // DWRITE_FONT_WEIGHT_EXTRA_BOLD
	GstDWriteFontWeightUltraBold  GstDWriteFontWeight = "ultra-bold"  // DWRITE_FONT_WEIGHT_ULTRA_BOLD
	GstDWriteFontWeightBlack      GstDWriteFontWeight = "black"       // DWRITE_FONT_WEIGHT_BLACK
	GstDWriteFontWeightHeavy      GstDWriteFontWeight = "heavy"       // DWRITE_FONT_WEIGHT_HEAVY
	GstDWriteFontWeightExtraBlack GstDWriteFontWeight = "extra-black" // DWRITE_FONT_WEIGHT_EXTRA_BLACK
	GstDWriteFontWeightUltraBlack GstDWriteFontWeight = "ultra-black" // DWRITE_FONT_WEIGHT_ULTRA_BLACK
)

type GstDWriteParagraphAlignment string

const (
	GstDWriteParagraphAlignmentNear   GstDWriteParagraphAlignment = "near"   // DWRITE_PARAGRAPH_ALIGNMENT_NEAR
	GstDWriteParagraphAlignmentFar    GstDWriteParagraphAlignment = "far"    // DWRITE_PARAGRAPH_ALIGNMENT_FAR
	GstDWriteParagraphAlignmentCenter GstDWriteParagraphAlignment = "center" // DWRITE_PARAGRAPH_ALIGNMENT_CENTER
)

type GstDWriteTextAlignment string

const (
	GstDWriteTextAlignmentLeading   GstDWriteTextAlignment = "leading"   // DWRITE_TEXT_ALIGNMENT_LEADING
	GstDWriteTextAlignmentTrailing  GstDWriteTextAlignment = "trailing"  // DWRITE_TEXT_ALIGNMENT_TRAILING
	GstDWriteTextAlignmentCenter    GstDWriteTextAlignment = "center"    // DWRITE_TEXT_ALIGNMENT_CENTER
	GstDWriteTextAlignmentJustified GstDWriteTextAlignment = "justified" // DWRITE_TEXT_ALIGNMENT_JUSTIFIED
)
