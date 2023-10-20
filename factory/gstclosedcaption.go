// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Combines GstVideoCaptionMeta with video input stream
type GstCCCombiner struct {
	*GstAggregator
}

func NewCCCombiner() (*GstCCCombiner, error) {
	e, err := gst.NewElement("cccombiner")
	if err != nil {
		return nil, err
	}
	return &GstCCCombiner{GstAggregator: &GstAggregator{Element: e}}, nil
}

func NewCCCombinerWithName(name string) (*GstCCCombiner, error) {
	e, err := gst.NewElementWithName("cccombiner", name)
	if err != nil {
		return nil, err
	}
	return &GstCCCombiner{GstAggregator: &GstAggregator{Element: e}}, nil
}

// Whether to output padding packets when schedule=true
// Default: true
func (e *GstCCCombiner) SetOutputPadding(value bool) error {
	return e.Element.SetProperty("output-padding", value)
}

func (e *GstCCCombiner) GetOutputPadding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Schedule caption buffers so that exactly one is output per video frame
// Default: true
func (e *GstCCCombiner) SetSchedule(value bool) error {
	return e.Element.SetProperty("schedule", value)
}

func (e *GstCCCombiner) GetSchedule() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schedule")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum number of buffers to queue for scheduling
// Default: 30, Min: 0, Max: -1
func (e *GstCCCombiner) SetMaxScheduled(value uint) error {
	return e.Element.SetProperty("max-scheduled", value)
}

func (e *GstCCCombiner) GetMaxScheduled() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-scheduled")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Converts Closed Captions between different formats
type GstCCConverter struct {
	*base.GstBaseTransform
}

func NewCCConverter() (*GstCCConverter, error) {
	e, err := gst.NewElement("ccconverter")
	if err != nil {
		return nil, err
	}
	return &GstCCConverter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCCConverterWithName(name string) (*GstCCConverter, error) {
	e, err := gst.NewElementWithName("ccconverter", name)
	if err != nil {
		return nil, err
	}
	return &GstCCConverter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Select which CDP sections to store in CDP packets
// Default: cc-svc-info+cc-data+time-code
func (e *GstCCConverter) SetCdpMode(value GstCCConverterCDPMode) error {
	e.Element.SetArg("cdp-mode", fmt.Sprint(value))
	return nil
}

func (e *GstCCConverter) GetCdpMode() (GstCCConverterCDPMode, error) {
	var value GstCCConverterCDPMode
	var ok bool
	v, err := e.Element.GetProperty("cdp-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCCConverterCDPMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCCConverterCDPMode")
	}
	return value, nil
}

// Extract GstVideoCaptionMeta from input stream
type GstCCExtractor struct {
	*gst.Element
}

func NewCCExtractor() (*GstCCExtractor, error) {
	e, err := gst.NewElement("ccextractor")
	if err != nil {
		return nil, err
	}
	return &GstCCExtractor{Element: e}, nil
}

func NewCCExtractorWithName(name string) (*GstCCExtractor, error) {
	e, err := gst.NewElementWithName("ccextractor", name)
	if err != nil {
		return nil, err
	}
	return &GstCCExtractor{Element: e}, nil
}

// Remove caption meta from outgoing video buffers
// Default: false
func (e *GstCCExtractor) SetRemoveCaptionMeta(value bool) error {
	return e.Element.SetProperty("remove-caption-meta", value)
}

func (e *GstCCExtractor) GetRemoveCaptionMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-caption-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Combines raw 608 streams
type GstCea608Mux struct {
	*GstAggregator
}

func NewCea608Mux() (*GstCea608Mux, error) {
	e, err := gst.NewElement("cea608mux")
	if err != nil {
		return nil, err
	}
	return &GstCea608Mux{GstAggregator: &GstAggregator{Element: e}}, nil
}

func NewCea608MuxWithName(name string) (*GstCea608Mux, error) {
	e, err := gst.NewElementWithName("cea608mux", name)
	if err != nil {
		return nil, err
	}
	return &GstCea608Mux{GstAggregator: &GstAggregator{Element: e}}, nil
}

// Extract line21 CC from SD video streams
type GstLine21Decoder struct {
	*GstVideoFilter
}

func NewLine21Decoder() (*GstLine21Decoder, error) {
	e, err := gst.NewElement("line21decoder")
	if err != nil {
		return nil, err
	}
	return &GstLine21Decoder{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewLine21DecoderWithName(name string) (*GstLine21Decoder, error) {
	e, err := gst.NewElementWithName("line21decoder", name)
	if err != nil {
		return nil, err
	}
	return &GstLine21Decoder{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Control whether and how detected CC meta should be inserted in the list of existing CC meta on a frame (if any).
// Default: add (0)
func (e *GstLine21Decoder) SetMode(value GstLine21DecoderMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstLine21Decoder) GetMode() (GstLine21DecoderMode, error) {
	var value GstLine21DecoderMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstLine21DecoderMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstLine21DecoderMode")
	}
	return value, nil
}

// Whether line 21 decoding should only be attempted when the input resolution matches NTSC
// Default: false
func (e *GstLine21Decoder) SetNtscOnly(value bool) error {
	return e.Element.SetProperty("ntsc-only", value)
}

func (e *GstLine21Decoder) GetNtscOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ntsc-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Inject line21 CC in SD video streams
type GstLine21Encoder struct {
	*GstVideoFilter
}

func NewLine21Encoder() (*GstLine21Encoder, error) {
	e, err := gst.NewElement("line21encoder")
	if err != nil {
		return nil, err
	}
	return &GstLine21Encoder{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewLine21EncoderWithName(name string) (*GstLine21Encoder, error) {
	e, err := gst.NewElementWithName("line21encoder", name)
	if err != nil {
		return nil, err
	}
	return &GstLine21Encoder{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Remove encoded caption meta from outgoing video buffers
// Default: false
func (e *GstLine21Encoder) SetRemoveCaptionMeta(value bool) error {
	return e.Element.SetProperty("remove-caption-meta", value)
}

func (e *GstLine21Encoder) GetRemoveCaptionMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-caption-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Decode cea608/cea708 data and overlay on proper position of a video buffer
type GstCeaCcOverlay struct {
	*gst.Element
}

func NewCeaCcOverlay() (*GstCeaCcOverlay, error) {
	e, err := gst.NewElement("cc708overlay")
	if err != nil {
		return nil, err
	}
	return &GstCeaCcOverlay{Element: e}, nil
}

func NewCeaCcOverlayWithName(name string) (*GstCeaCcOverlay, error) {
	e, err := gst.NewElementWithName("cc708overlay", name)
	if err != nil {
		return nil, err
	}
	return &GstCeaCcOverlay{Element: e}, nil
}

// Pango font description of font to be used for rendering.
// See documentation of pango_font_description_from_string for syntax.
// this will override closed caption stream specified font style/pen size.
// Default: NULL
func (e *GstCeaCcOverlay) SetFontDesc(value string) error {
	return e.Element.SetProperty("font-desc", value)
}

func (e *GstCeaCcOverlay) GetFontDesc() (string, error) {
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

// Service number. Service 1 is designated as the Primary Caption Service, Service 2 is the Secondary Language Service.
// Default: 1, Min: -1, Max: 63
func (e *GstCeaCcOverlay) SetServiceNumber(value int) error {
	return e.Element.SetProperty("service-number", value)
}

func (e *GstCeaCcOverlay) GetServiceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("service-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Whether to render the text string
// Default: false
func (e *GstCeaCcOverlay) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstCeaCcOverlay) GetSilent() (bool, error) {
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

// Window's Horizontal position
// Default: center (1)
func (e *GstCeaCcOverlay) SetWindowHPos(value GstCeaCcOverlayWinHPos) error {
	e.Element.SetArg("window-h-pos", string(value))
	return nil
}

func (e *GstCeaCcOverlay) GetWindowHPos() (GstCeaCcOverlayWinHPos, error) {
	var value GstCeaCcOverlayWinHPos
	var ok bool
	v, err := e.Element.GetProperty("window-h-pos")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCeaCcOverlayWinHPos)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCeaCcOverlayWinHPos")
	}
	return value, nil
}

type GstCCConverterCDPMode int

const (
	GstCCConverterCDPModeTimeCode  GstCCConverterCDPMode = 0x00000001 // Store time code information in CDP packets
	GstCCConverterCDPModeCcData    GstCCConverterCDPMode = 0x00000002 // Store CC data in CDP packets
	GstCCConverterCDPModeCcSvcInfo GstCCConverterCDPMode = 0x00000004 // Store CC service information in CDP packets
)

type GstCeaCcOverlayWinHPos string

const (
	GstCeaCcOverlayWinHPosLeft   GstCeaCcOverlayWinHPos = "left"   // left
	GstCeaCcOverlayWinHPosCenter GstCeaCcOverlayWinHPos = "center" // center
	GstCeaCcOverlayWinHPosRight  GstCeaCcOverlayWinHPos = "right"  // right
	GstCeaCcOverlayWinHPosAuto   GstCeaCcOverlayWinHPos = "auto"   // auto
)

type GstLine21DecoderMode string

const (
	GstLine21DecoderModeAdd     GstLine21DecoderMode = "add"     // add new CC meta on top of other CC meta, if any
	GstLine21DecoderModeDrop    GstLine21DecoderMode = "drop"    // ignore CC if a CC meta was already present
	GstLine21DecoderModeReplace GstLine21DecoderMode = "replace" // replace existing CC meta
)
