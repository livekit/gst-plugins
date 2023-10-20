// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Parses and frames MPEG-1 and MPEG-2 elementary video streams
type GstMpegvParse struct {
	*GstBaseParse
}

func NewMpegvParse() (*GstMpegvParse, error) {
	e, err := gst.NewElement("mpegvideoparse")
	if err != nil {
		return nil, err
	}
	return &GstMpegvParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewMpegvParseWithName(name string) (*GstMpegvParse, error) {
	e, err := gst.NewElementWithName("mpegvideoparse", name)
	if err != nil {
		return nil, err
	}
	return &GstMpegvParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Drop data until valid configuration data is received either in the stream or through caps
// Default: true
func (e *GstMpegvParse) SetDrop(value bool) error {
	return e.Element.SetProperty("drop", value)
}

func (e *GstMpegvParse) GetDrop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Split frame when encountering GOP
// Default: false
func (e *GstMpegvParse) SetGopSplit(value bool) error {
	return e.Element.SetProperty("gop-split", value)
}

func (e *GstMpegvParse) GetGopSplit() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gop-split")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Parses VC1 streams
type GstVC1Parse struct {
	*GstBaseParse
}

func NewVC1Parse() (*GstVC1Parse, error) {
	e, err := gst.NewElement("vc1parse")
	if err != nil {
		return nil, err
	}
	return &GstVC1Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewVC1ParseWithName(name string) (*GstVC1Parse, error) {
	e, err := gst.NewElementWithName("vc1parse", name)
	if err != nil {
		return nil, err
	}
	return &GstVC1Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Parses H.263 streams
type GstH263Parse struct {
	*GstBaseParse
}

func NewH263Parse() (*GstH263Parse, error) {
	e, err := gst.NewElement("h263parse")
	if err != nil {
		return nil, err
	}
	return &GstH263Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewH263ParseWithName(name string) (*GstH263Parse, error) {
	e, err := gst.NewElementWithName("h263parse", name)
	if err != nil {
		return nil, err
	}
	return &GstH263Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Parses H.264 streams
type GstH264Parse struct {
	*GstBaseParse
}

func NewH264Parse() (*GstH264Parse, error) {
	e, err := gst.NewElement("h264parse")
	if err != nil {
		return nil, err
	}
	return &GstH264Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewH264ParseWithName(name string) (*GstH264Parse, error) {
	e, err := gst.NewElementWithName("h264parse", name)
	if err != nil {
		return nil, err
	}
	return &GstH264Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Send SPS and PPS Insertion Interval in seconds (sprop parameter sets will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)
// Default: 0, Min: -1, Max: 3600
func (e *GstH264Parse) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstH264Parse) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Update time code values in Picture Timing SEI if GstVideoTimeCodeMeta is attached to incoming buffer and also Picture Timing SEI exists in the bitstream. To make this property work, SPS must contain VUI and pic_struct_present_flag of VUI must be non-zero
// Default: false
func (e *GstH264Parse) SetUpdateTimecode(value bool) error {
	return e.Element.SetProperty("update-timecode", value)
}

func (e *GstH264Parse) GetUpdateTimecode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("update-timecode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Parses H.265 streams
type GstH265Parse struct {
	*GstBaseParse
}

func NewH265Parse() (*GstH265Parse, error) {
	e, err := gst.NewElement("h265parse")
	if err != nil {
		return nil, err
	}
	return &GstH265Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewH265ParseWithName(name string) (*GstH265Parse, error) {
	e, err := gst.NewElementWithName("h265parse", name)
	if err != nil {
		return nil, err
	}
	return &GstH265Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Send VPS, SPS and PPS Insertion Interval in seconds (sprop parameter sets will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)
// Default: 0, Min: -1, Max: 3600
func (e *GstH265Parse) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstH265Parse) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Parses JPEG 2000 files
type GstJPEG2000Parse struct {
	*GstBaseParse
}

func NewJPEG2000Parse() (*GstJPEG2000Parse, error) {
	e, err := gst.NewElement("jpeg2000parse")
	if err != nil {
		return nil, err
	}
	return &GstJPEG2000Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewJPEG2000ParseWithName(name string) (*GstJPEG2000Parse, error) {
	e, err := gst.NewElementWithName("jpeg2000parse", name)
	if err != nil {
		return nil, err
	}
	return &GstJPEG2000Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Parses MPEG-4 Part 2 elementary video streams
type GstMpeg4VParse struct {
	*GstBaseParse
}

func NewMpeg4VParse() (*GstMpeg4VParse, error) {
	e, err := gst.NewElement("mpeg4videoparse")
	if err != nil {
		return nil, err
	}
	return &GstMpeg4VParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewMpeg4VParseWithName(name string) (*GstMpeg4VParse, error) {
	e, err := gst.NewElementWithName("mpeg4videoparse", name)
	if err != nil {
		return nil, err
	}
	return &GstMpeg4VParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Send Configuration Insertion Interval in seconds (configuration headers will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)
// Default: 0, Min: -1, Max: 3600
func (e *GstMpeg4VParse) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstMpeg4VParse) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Drop data until valid configuration data is received either in the stream or through caps
// Default: true
func (e *GstMpeg4VParse) SetDrop(value bool) error {
	return e.Element.SetProperty("drop", value)
}

func (e *GstMpeg4VParse) GetDrop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Parses AV1 streams
type GstAV1Parse struct {
	*GstBaseParse
}

func NewAV1Parse() (*GstAV1Parse, error) {
	e, err := gst.NewElement("av1parse")
	if err != nil {
		return nil, err
	}
	return &GstAV1Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewAV1ParseWithName(name string) (*GstAV1Parse, error) {
	e, err := gst.NewElementWithName("av1parse", name)
	if err != nil {
		return nil, err
	}
	return &GstAV1Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Parses Dirac streams
type GstDiracParse struct {
	*GstBaseParse
}

func NewDiracParse() (*GstDiracParse, error) {
	e, err := gst.NewElement("diracparse")
	if err != nil {
		return nil, err
	}
	return &GstDiracParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewDiracParseWithName(name string) (*GstDiracParse, error) {
	e, err := gst.NewElementWithName("diracparse", name)
	if err != nil {
		return nil, err
	}
	return &GstDiracParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Parses PNG files
type GstPngParse struct {
	*GstBaseParse
}

func NewPngParse() (*GstPngParse, error) {
	e, err := gst.NewElement("pngparse")
	if err != nil {
		return nil, err
	}
	return &GstPngParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewPngParseWithName(name string) (*GstPngParse, error) {
	e, err := gst.NewElementWithName("pngparse", name)
	if err != nil {
		return nil, err
	}
	return &GstPngParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Parses VP9 streams
type GstVp9Parse struct {
	*GstBaseParse
}

func NewVp9Parse() (*GstVp9Parse, error) {
	e, err := gst.NewElement("vp9parse")
	if err != nil {
		return nil, err
	}
	return &GstVp9Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewVp9ParseWithName(name string) (*GstVp9Parse, error) {
	e, err := gst.NewElementWithName("vp9parse", name)
	if err != nil {
		return nil, err
	}
	return &GstVp9Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}
