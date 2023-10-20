// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// MPEG2 video encoder based on Intel Media SDK
type GstMsdkMPEG2Enc struct {
	*GstMsdkEnc
}

func NewMsdkMPEG2Enc() (*GstMsdkMPEG2Enc, error) {
	e, err := gst.NewElement("msdkmpeg2enc")
	if err != nil {
		return nil, err
	}
	return &GstMsdkMPEG2Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMsdkMPEG2EncWithName(name string) (*GstMsdkMPEG2Enc, error) {
	e, err := gst.NewElementWithName("msdkmpeg2enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkMPEG2Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// H264 video encoder based on Intel Media SDK
type GstMsdkH264Enc struct {
	*GstMsdkEnc
}

func NewMsdkH264Enc() (*GstMsdkH264Enc, error) {
	e, err := gst.NewElement("msdkh264enc")
	if err != nil {
		return nil, err
	}
	return &GstMsdkH264Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMsdkH264EncWithName(name string) (*GstMsdkH264Enc, error) {
	e, err := gst.NewElementWithName("msdkh264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkH264Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Maximum quantizer scale for I/P/B frames
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH264Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

func (e *GstMsdkH264Enc) GetMaxQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum quantizer scale for I frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH264Enc) SetMaxQpI(value uint) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstMsdkH264Enc) GetMaxQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Miminum quantizer scale for I frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH264Enc) SetMinQpI(value uint) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstMsdkH264Enc) GetMinQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable P-Pyramid Reference structure
// Default: false
func (e *GstMsdkH264Enc) SetPPyramid(value bool) error {
	return e.Element.SetProperty("p-pyramid", value)
}

func (e *GstMsdkH264Enc) GetPPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("p-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Down sampling mode in look ahead bitrate control
// Default: default (0)
func (e *GstMsdkH264Enc) SetRcLookaheadDs(value GstMsdkEncRCLookAheadDownsampling) error {
	e.Element.SetArg("rc-lookahead-ds", string(value))
	return nil
}

func (e *GstMsdkH264Enc) GetRcLookaheadDs() (GstMsdkEncRCLookAheadDownsampling, error) {
	var value GstMsdkEncRCLookAheadDownsampling
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead-ds")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncRCLookAheadDownsampling)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncRCLookAheadDownsampling")
	}
	return value, nil
}

// Set frame packing mode for Stereoscopic content
// Default: none (-1)
func (e *GstMsdkH264Enc) SetFramePacking(value GstMsdkH264EncFramePacking) error {
	e.Element.SetArg("frame-packing", string(value))
	return nil
}

func (e *GstMsdkH264Enc) GetFramePacking() (GstMsdkH264EncFramePacking, error) {
	var value GstMsdkH264EncFramePacking
	var ok bool
	v, err := e.Element.GetProperty("frame-packing")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkH264EncFramePacking)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkH264EncFramePacking")
	}
	return value, nil
}

// Set intra refresh type
// Default: no (0)
func (e *GstMsdkH264Enc) SetIntraRefreshType(value GstMsdkEncIntraRefreshType) error {
	e.Element.SetArg("intra-refresh-type", string(value))
	return nil
}

func (e *GstMsdkH264Enc) GetIntraRefreshType() (GstMsdkEncIntraRefreshType, error) {
	var value GstMsdkEncIntraRefreshType
	var ok bool
	v, err := e.Element.GetProperty("intra-refresh-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncIntraRefreshType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncIntraRefreshType")
	}
	return value, nil
}

// Maximum quantizer scale for B frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH264Enc) SetMaxQpB(value uint) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstMsdkH264Enc) GetMaxQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum slice size in bytes (if enabled MSDK will ignore the control over num-slices)
// Default: 0, Min: 0, Max: 4294967295
func (e *GstMsdkH264Enc) SetMaxSliceSize(value uint) error {
	return e.Element.SetProperty("max-slice-size", value)
}

func (e *GstMsdkH264Enc) GetMaxSliceSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-slice-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Miminum quantizer scale for B frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH264Enc) SetMinQpB(value uint) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstMsdkH264Enc) GetMinQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable B-Pyramid Reference structure
// Default: false
func (e *GstMsdkH264Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

func (e *GstMsdkH264Enc) GetBPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Option of disable deblocking idc
// Default: 0, Min: 0, Max: 2
func (e *GstMsdkH264Enc) SetDblkIdc(value uint) error {
	return e.Element.SetProperty("dblk-idc", value)
}

func (e *GstMsdkH264Enc) GetDblkIdc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dblk-idc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Miminum quantizer scale for P frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH264Enc) SetMinQpP(value uint) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstMsdkH264Enc) GetMinQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Insert picture timing SEI with pic_struct syntax
// Default: true
func (e *GstMsdkH264Enc) SetPicTimingSei(value bool) error {
	return e.Element.SetProperty("pic-timing-sei", value)
}

func (e *GstMsdkH264Enc) GetPicTimingSei() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pic-timing-sei")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable CABAC entropy coding
// Default: true
func (e *GstMsdkH264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

func (e *GstMsdkH264Enc) GetCabac() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cabac")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable low power mode (DEPRECATED, use tune instead)
// Default: false
func (e *GstMsdkH264Enc) SetLowPower(value bool) error {
	return e.Element.SetProperty("low-power", value)
}

func (e *GstMsdkH264Enc) GetLowPower() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-power")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum quantizer scale for P frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH264Enc) SetMaxQpP(value uint) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstMsdkH264Enc) GetMaxQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimal quantizer scale for I/P/B frames
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH264Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

func (e *GstMsdkH264Enc) GetMinQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable Trellis Quantization
// Default: None (0x00000000)
func (e *GstMsdkH264Enc) SetTrellis(value GstMsdkEncTrellisQuantization) error {
	e.Element.SetArg("trellis", fmt.Sprint(value))
	return nil
}

func (e *GstMsdkH264Enc) GetTrellis() (GstMsdkEncTrellisQuantization, error) {
	var value GstMsdkEncTrellisQuantization
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncTrellisQuantization)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncTrellisQuantization")
	}
	return value, nil
}

// Encoder tuning option
// Default: auto (0)
func (e *GstMsdkH264Enc) SetTune(value GstMsdkEncTuneMode) error {
	e.Element.SetArg("tune", string(value))
	return nil
}

func (e *GstMsdkH264Enc) GetTune() (GstMsdkEncTuneMode, error) {
	var value GstMsdkEncTuneMode
	var ok bool
	v, err := e.Element.GetProperty("tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncTuneMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncTuneMode")
	}
	return value, nil
}

// MJPEG video encoder based on Intel Media SDK
type GstMsdkMJPEGEnc struct {
	*GstMsdkEnc
}

func NewMsdkMJPEGEnc() (*GstMsdkMJPEGEnc, error) {
	e, err := gst.NewElement("msdkmjpegenc")
	if err != nil {
		return nil, err
	}
	return &GstMsdkMJPEGEnc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMsdkMJPEGEncWithName(name string) (*GstMsdkMJPEGEnc, error) {
	e, err := gst.NewElementWithName("msdkmjpegenc", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkMJPEGEnc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Quality of encoding
// Default: 85, Min: 0, Max: 100
func (e *GstMsdkMJPEGEnc) SetQuality(value uint) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstMsdkMJPEGEnc) GetQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// MPEG2 video decoder based on Intel Media SDK
type GstMsdkMPEG2Dec struct {
	*GstMsdkDec
}

func NewMsdkMPEG2Dec() (*GstMsdkMPEG2Dec, error) {
	e, err := gst.NewElement("msdkmpeg2dec")
	if err != nil {
		return nil, err
	}
	return &GstMsdkMPEG2Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewMsdkMPEG2DecWithName(name string) (*GstMsdkMPEG2Dec, error) {
	e, err := gst.NewElementWithName("msdkmpeg2dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkMPEG2Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Decoded frames output order
// Default: display (0)
func (e *GstMsdkMPEG2Dec) SetOutputOrder(value GstMsdkDecOutputOrder) error {
	e.Element.SetArg("output-order", string(value))
	return nil
}

func (e *GstMsdkMPEG2Dec) GetOutputOrder() (GstMsdkDecOutputOrder, error) {
	var value GstMsdkDecOutputOrder
	var ok bool
	v, err := e.Element.GetProperty("output-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkDecOutputOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkDecOutputOrder")
	}
	return value, nil
}

// VP9 video decoder based on Intel Media SDK
type GstMsdkVP9Dec struct {
	*GstMsdkDec
}

func NewMsdkVP9Dec() (*GstMsdkVP9Dec, error) {
	e, err := gst.NewElement("msdkvp9dec")
	if err != nil {
		return nil, err
	}
	return &GstMsdkVP9Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewMsdkVP9DecWithName(name string) (*GstMsdkVP9Dec, error) {
	e, err := gst.NewElementWithName("msdkvp9dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkVP9Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Decoded frames output order
// Default: display (0)
func (e *GstMsdkVP9Dec) SetOutputOrder(value GstMsdkDecOutputOrder) error {
	e.Element.SetArg("output-order", string(value))
	return nil
}

func (e *GstMsdkVP9Dec) GetOutputOrder() (GstMsdkDecOutputOrder, error) {
	var value GstMsdkDecOutputOrder
	var ok bool
	v, err := e.Element.GetProperty("output-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkDecOutputOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkDecOutputOrder")
	}
	return value, nil
}

// VP9 video encoder based on Intel Media SDK
type GstMsdkVP9Enc struct {
	*GstMsdkEnc
}

func NewMsdkVP9Enc() (*GstMsdkVP9Enc, error) {
	e, err := gst.NewElement("msdkvp9enc")
	if err != nil {
		return nil, err
	}
	return &GstMsdkVP9Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMsdkVP9EncWithName(name string) (*GstMsdkVP9Enc, error) {
	e, err := gst.NewElementWithName("msdkvp9enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkVP9Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// H264 video decoder based on Intel Media SDK
type GstMsdkH264Dec struct {
	*GstMsdkDec
}

func NewMsdkH264Dec() (*GstMsdkH264Dec, error) {
	e, err := gst.NewElement("msdkh264dec")
	if err != nil {
		return nil, err
	}
	return &GstMsdkH264Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewMsdkH264DecWithName(name string) (*GstMsdkH264Dec, error) {
	e, err := gst.NewElementWithName("msdkh264dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkH264Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Decoded frames output order
// Default: display (0)
func (e *GstMsdkH264Dec) SetOutputOrder(value GstMsdkDecOutputOrder) error {
	e.Element.SetArg("output-order", string(value))
	return nil
}

func (e *GstMsdkH264Dec) GetOutputOrder() (GstMsdkDecOutputOrder, error) {
	var value GstMsdkDecOutputOrder
	var ok bool
	v, err := e.Element.GetProperty("output-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkDecOutputOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkDecOutputOrder")
	}
	return value, nil
}

// H265 video encoder based on Intel Media SDK
type GstMsdkH265Enc struct {
	*GstMsdkEnc
}

func NewMsdkH265Enc() (*GstMsdkH265Enc, error) {
	e, err := gst.NewElement("msdkh265enc")
	if err != nil {
		return nil, err
	}
	return &GstMsdkH265Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMsdkH265EncWithName(name string) (*GstMsdkH265Enc, error) {
	e, err := gst.NewElementWithName("msdkh265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkH265Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Set intra refresh type
// Default: no (0)
func (e *GstMsdkH265Enc) SetIntraRefreshType(value GstMsdkEncIntraRefreshType) error {
	e.Element.SetArg("intra-refresh-type", string(value))
	return nil
}

func (e *GstMsdkH265Enc) GetIntraRefreshType() (GstMsdkEncIntraRefreshType, error) {
	var value GstMsdkEncIntraRefreshType
	var ok bool
	v, err := e.Element.GetProperty("intra-refresh-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncIntraRefreshType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncIntraRefreshType")
	}
	return value, nil
}

// Maximum quantizer scale for B frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH265Enc) SetMaxQpB(value uint) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstMsdkH265Enc) GetMaxQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum quantizer scale for P frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH265Enc) SetMaxQpP(value uint) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstMsdkH265Enc) GetMaxQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Transform Skip option
// Default: auto (0)
func (e *GstMsdkH265Enc) SetTransformSkip(value GstMsdkEncTransformSkip) error {
	e.Element.SetArg("transform-skip", string(value))
	return nil
}

func (e *GstMsdkH265Enc) GetTransformSkip() (GstMsdkEncTransformSkip, error) {
	var value GstMsdkEncTransformSkip
	var ok bool
	v, err := e.Element.GetProperty("transform-skip")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncTransformSkip)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncTransformSkip")
	}
	return value, nil
}

// Maximum quantizer scale for I/P/B frames
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH265Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

func (e *GstMsdkH265Enc) GetMaxQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Miminum quantizer scale for I frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH265Enc) SetMinQpI(value uint) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstMsdkH265Enc) GetMinQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Miminum quantizer scale for P frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH265Enc) SetMinQpP(value uint) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstMsdkH265Enc) GetMinQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// number of rows for tiled encoding
// Default: 1, Min: 1, Max: 8192
func (e *GstMsdkH265Enc) SetNumTileRows(value uint) error {
	return e.Element.SetProperty("num-tile-rows", value)
}

func (e *GstMsdkH265Enc) GetNumTileRows() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-tile-rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Option of disable deblocking idc
// Default: 0, Min: 0, Max: 2
func (e *GstMsdkH265Enc) SetDblkIdc(value uint) error {
	return e.Element.SetProperty("dblk-idc", value)
}

func (e *GstMsdkH265Enc) GetDblkIdc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dblk-idc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimal quantizer scale for I/P/B frames
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH265Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

func (e *GstMsdkH265Enc) GetMinQp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable P-Pyramid Reference structure
// Default: false
func (e *GstMsdkH265Enc) SetPPyramid(value bool) error {
	return e.Element.SetProperty("p-pyramid", value)
}

func (e *GstMsdkH265Enc) GetPPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("p-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Insert picture timing SEI with pic_struct syntax
// Default: true
func (e *GstMsdkH265Enc) SetPicTimingSei(value bool) error {
	return e.Element.SetProperty("pic-timing-sei", value)
}

func (e *GstMsdkH265Enc) GetPicTimingSei() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pic-timing-sei")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable B-Pyramid Reference structure
// Default: false
func (e *GstMsdkH265Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

func (e *GstMsdkH265Enc) GetBPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable low power mode (DEPRECATED, use tune instead)
// Default: false
func (e *GstMsdkH265Enc) SetLowPower(value bool) error {
	return e.Element.SetProperty("low-power", value)
}

func (e *GstMsdkH265Enc) GetLowPower() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-power")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum quantizer scale for I frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH265Enc) SetMaxQpI(value uint) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstMsdkH265Enc) GetMaxQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum slice size in bytes (if enabled MSDK will ignore the control over num-slices)
// Default: 0, Min: 0, Max: 4294967295
func (e *GstMsdkH265Enc) SetMaxSliceSize(value uint) error {
	return e.Element.SetProperty("max-slice-size", value)
}

func (e *GstMsdkH265Enc) GetMaxSliceSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-slice-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Miminum quantizer scale for B frame
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkH265Enc) SetMinQpB(value uint) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstMsdkH265Enc) GetMinQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// number of columns for tiled encoding
// Default: 1, Min: 1, Max: 8192
func (e *GstMsdkH265Enc) SetNumTileCols(value uint) error {
	return e.Element.SetProperty("num-tile-cols", value)
}

func (e *GstMsdkH265Enc) GetNumTileCols() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-tile-cols")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Encoder tuning option
// Default: auto (0)
func (e *GstMsdkH265Enc) SetTune(value GstMsdkEncTuneMode) error {
	e.Element.SetArg("tune", string(value))
	return nil
}

func (e *GstMsdkH265Enc) GetTune() (GstMsdkEncTuneMode, error) {
	var value GstMsdkEncTuneMode
	var ok bool
	v, err := e.Element.GetProperty("tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncTuneMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncTuneMode")
	}
	return value, nil
}

// MJPEG video decoder based on Intel Media SDK
type GstMsdkMJPEGDec struct {
	*GstMsdkDec
}

func NewMsdkMJPEGDec() (*GstMsdkMJPEGDec, error) {
	e, err := gst.NewElement("msdkmjpegdec")
	if err != nil {
		return nil, err
	}
	return &GstMsdkMJPEGDec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewMsdkMJPEGDecWithName(name string) (*GstMsdkMJPEGDec, error) {
	e, err := gst.NewElementWithName("msdkmjpegdec", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkMJPEGDec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// AV1 video decoder based on Intel Media SDK
type GstMsdkAV1Dec struct {
	*GstMsdkDec
}

func NewMsdkAV1Dec() (*GstMsdkAV1Dec, error) {
	e, err := gst.NewElement("msdkav1dec")
	if err != nil {
		return nil, err
	}
	return &GstMsdkAV1Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewMsdkAV1DecWithName(name string) (*GstMsdkAV1Dec, error) {
	e, err := gst.NewElementWithName("msdkav1dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkAV1Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// VP8 video decoder based on Intel Media SDK
type GstMsdkVP8Dec struct {
	*GstMsdkDec
}

func NewMsdkVP8Dec() (*GstMsdkVP8Dec, error) {
	e, err := gst.NewElement("msdkvp8dec")
	if err != nil {
		return nil, err
	}
	return &GstMsdkVP8Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewMsdkVP8DecWithName(name string) (*GstMsdkVP8Dec, error) {
	e, err := gst.NewElementWithName("msdkvp8dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkVP8Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Decoded frames output order
// Default: display (0)
func (e *GstMsdkVP8Dec) SetOutputOrder(value GstMsdkDecOutputOrder) error {
	e.Element.SetArg("output-order", string(value))
	return nil
}

func (e *GstMsdkVP8Dec) GetOutputOrder() (GstMsdkDecOutputOrder, error) {
	var value GstMsdkDecOutputOrder
	var ok bool
	v, err := e.Element.GetProperty("output-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkDecOutputOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkDecOutputOrder")
	}
	return value, nil
}

// A MediaSDK Video Postprocessing Filter
type GstMsdkVPP struct {
	*base.GstBaseTransform
}

func NewMsdkVPP() (*GstMsdkVPP, error) {
	e, err := gst.NewElement("msdkvpp")
	if err != nil {
		return nil, err
	}
	return &GstMsdkVPP{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewMsdkVPPWithName(name string) (*GstMsdkVPP, error) {
	e, err := gst.NewElementWithName("msdkvpp", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkVPP{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// The Contrast of the video
// Default: 1, Min: 0, Max: 10
func (e *GstMsdkVPP) SetContrast(value float32) error {
	return e.Element.SetProperty("contrast", value)
}

func (e *GstMsdkVPP) GetContrast() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The factor of detail/edge enhancement filter algorithm
// Default: 0, Min: 0, Max: 100
func (e *GstMsdkVPP) SetDetail(value uint) error {
	return e.Element.SetProperty("detail", value)
}

func (e *GstMsdkVPP) GetDetail() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("detail")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Depth of asynchronous pipeline
// Default: 1, Min: 1, Max: 1
func (e *GstMsdkVPP) SetAsyncDepth(value uint) error {
	return e.Element.SetProperty("async-depth", value)
}

func (e *GstMsdkVPP) GetAsyncDepth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("async-depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Pixels to crop at right
// Default: 0, Min: 0, Max: 65535
func (e *GstMsdkVPP) SetCropRight(value uint) error {
	return e.Element.SetProperty("crop-right", value)
}

func (e *GstMsdkVPP) GetCropRight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstMsdkVPP) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstMsdkVPP) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable HDR to SDR tone mapping
// Default: false
func (e *GstMsdkVPP) SetHdrToneMapping(value bool) error {
	return e.Element.SetProperty("hdr-tone-mapping", value)
}

func (e *GstMsdkVPP) GetHdrToneMapping() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hdr-tone-mapping")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The Brightness of the video
// Default: 0, Min: -100, Max: 100
func (e *GstMsdkVPP) SetBrightness(value float32) error {
	return e.Element.SetProperty("brightness", value)
}

func (e *GstMsdkVPP) GetBrightness() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Pixels to crop at bottom
// Default: 0, Min: 0, Max: 65535
func (e *GstMsdkVPP) SetCropBottom(value uint) error {
	return e.Element.SetProperty("crop-bottom", value)
}

func (e *GstMsdkVPP) GetCropBottom() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Pixels to crop at top
// Default: 0, Min: 0, Max: 65535
func (e *GstMsdkVPP) SetCropTop(value uint) error {
	return e.Element.SetProperty("crop-top", value)
}

func (e *GstMsdkVPP) GetCropTop() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Deinterlace method to use
// Default: bob (1)
func (e *GstMsdkVPP) SetDeinterlaceMethod(value GstMsdkVPPDeinterlaceMethod) error {
	e.Element.SetArg("deinterlace-method", string(value))
	return nil
}

func (e *GstMsdkVPP) GetDeinterlaceMethod() (GstMsdkVPPDeinterlaceMethod, error) {
	var value GstMsdkVPPDeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("deinterlace-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVPPDeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVPPDeinterlaceMethod")
	}
	return value, nil
}

// Denoising Factor
// Default: 0, Min: 0, Max: 100
func (e *GstMsdkVPP) SetDenoise(value uint) error {
	return e.Element.SetProperty("denoise", value)
}

func (e *GstMsdkVPP) GetDenoise() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("denoise")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The Framerate Control Alogorithm to use
// Default: none (0)
func (e *GstMsdkVPP) SetFrcAlgorithm(value GstMsdkVPPFrcAlgorithm) error {
	e.Element.SetArg("frc-algorithm", string(value))
	return nil
}

func (e *GstMsdkVPP) GetFrcAlgorithm() (GstMsdkVPPFrcAlgorithm, error) {
	var value GstMsdkVPPFrcAlgorithm
	var ok bool
	v, err := e.Element.GetProperty("frc-algorithm")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVPPFrcAlgorithm)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVPPFrcAlgorithm")
	}
	return value, nil
}

// Enable hardware VPP
// Default: true
func (e *GstMsdkVPP) SetHardware(value bool) error {
	return e.Element.SetProperty("hardware", value)
}

func (e *GstMsdkVPP) GetHardware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hardware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The hue of the video
// Default: 0, Min: -180, Max: 180
func (e *GstMsdkVPP) SetHue(value float32) error {
	return e.Element.SetProperty("hue", value)
}

func (e *GstMsdkVPP) GetHue() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Rotation Angle (DEPRECATED, use video-direction instead)
// Default: 0 (0)
func (e *GstMsdkVPP) SetRotation(value GstMsdkVPPRotation) error {
	e.Element.SetArg("rotation", string(value))
	return nil
}

func (e *GstMsdkVPP) GetRotation() (GstMsdkVPPRotation, error) {
	var value GstMsdkVPPRotation
	var ok bool
	v, err := e.Element.GetProperty("rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVPPRotation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVPPRotation")
	}
	return value, nil
}

// The Scaling mode to use
// Default: disable (0)
func (e *GstMsdkVPP) SetScalingMode(value GstMsdkVPPScalingMode) error {
	e.Element.SetArg("scaling-mode", string(value))
	return nil
}

func (e *GstMsdkVPP) GetScalingMode() (GstMsdkVPPScalingMode, error) {
	var value GstMsdkVPPScalingMode
	var ok bool
	v, err := e.Element.GetProperty("scaling-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVPPScalingMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVPPScalingMode")
	}
	return value, nil
}

// Video direction: rotation and flipping, it will override both mirroring & rotation properties if set explicitly
// Default: identity (0)
func (e *GstMsdkVPP) SetVideoDirection(value interface{}) error {
	return e.Element.SetProperty("video-direction", value)
}

func (e *GstMsdkVPP) GetVideoDirection() (interface{}, error) {
	return e.Element.GetProperty("video-direction")
}

// Pixels to crop at left
// Default: 0, Min: 0, Max: 65535
func (e *GstMsdkVPP) SetCropLeft(value uint) error {
	return e.Element.SetProperty("crop-left", value)
}

func (e *GstMsdkVPP) GetCropLeft() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Deinterlace mode to use
// Default: auto (0)
func (e *GstMsdkVPP) SetDeinterlaceMode(value GstMsdkVPPDeinterlaceMode) error {
	e.Element.SetArg("deinterlace-mode", string(value))
	return nil
}

func (e *GstMsdkVPP) GetDeinterlaceMode() (GstMsdkVPPDeinterlaceMode, error) {
	var value GstMsdkVPPDeinterlaceMode
	var ok bool
	v, err := e.Element.GetProperty("deinterlace-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVPPDeinterlaceMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVPPDeinterlaceMode")
	}
	return value, nil
}

// The Mirroring type (DEPRECATED, use video-direction instead)
// Default: disable (0)
func (e *GstMsdkVPP) SetMirroring(value GstMsdkVPPMirroring) error {
	e.Element.SetArg("mirroring", string(value))
	return nil
}

func (e *GstMsdkVPP) GetMirroring() (GstMsdkVPPMirroring, error) {
	var value GstMsdkVPPMirroring
	var ok bool
	v, err := e.Element.GetProperty("mirroring")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVPPMirroring)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVPPMirroring")
	}
	return value, nil
}

// The Saturation of the video
// Default: 1, Min: 0, Max: 10
func (e *GstMsdkVPP) SetSaturation(value float32) error {
	return e.Element.SetProperty("saturation", value)
}

func (e *GstMsdkVPP) GetSaturation() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// AV1 video encoder based on Intel oneVPL
type GstMsdkAV1Enc struct {
	*GstMsdkEnc
}

func NewMsdkAV1Enc() (*GstMsdkAV1Enc, error) {
	e, err := gst.NewElement("msdkav1enc")
	if err != nil {
		return nil, err
	}
	return &GstMsdkAV1Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewMsdkAV1EncWithName(name string) (*GstMsdkAV1Enc, error) {
	e, err := gst.NewElementWithName("msdkav1enc", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkAV1Enc{GstMsdkEnc: &GstMsdkEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Enable B-Pyramid Reference structure
// Default: false
func (e *GstMsdkAV1Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

func (e *GstMsdkAV1Enc) GetBPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("b-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// number of columns for tiled encoding
// Default: 1, Min: 1, Max: 64
func (e *GstMsdkAV1Enc) SetNumTileCols(value uint) error {
	return e.Element.SetProperty("num-tile-cols", value)
}

func (e *GstMsdkAV1Enc) GetNumTileCols() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-tile-cols")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// number of rows for tiled encoding
// Default: 1, Min: 1, Max: 64
func (e *GstMsdkAV1Enc) SetNumTileRows(value uint) error {
	return e.Element.SetProperty("num-tile-rows", value)
}

func (e *GstMsdkAV1Enc) GetNumTileRows() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-tile-rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable P-Pyramid Reference structure
// Default: false
func (e *GstMsdkAV1Enc) SetPPyramid(value bool) error {
	return e.Element.SetProperty("p-pyramid", value)
}

func (e *GstMsdkAV1Enc) GetPPyramid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("p-pyramid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// H265 video decoder based on Intel Media SDK
type GstMsdkH265Dec struct {
	*GstMsdkDec
}

func NewMsdkH265Dec() (*GstMsdkH265Dec, error) {
	e, err := gst.NewElement("msdkh265dec")
	if err != nil {
		return nil, err
	}
	return &GstMsdkH265Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewMsdkH265DecWithName(name string) (*GstMsdkH265Dec, error) {
	e, err := gst.NewElementWithName("msdkh265dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkH265Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Decoded frames output order
// Default: display (0)
func (e *GstMsdkH265Dec) SetOutputOrder(value GstMsdkDecOutputOrder) error {
	e.Element.SetArg("output-order", string(value))
	return nil
}

func (e *GstMsdkH265Dec) GetOutputOrder() (GstMsdkDecOutputOrder, error) {
	var value GstMsdkDecOutputOrder
	var ok bool
	v, err := e.Element.GetProperty("output-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkDecOutputOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkDecOutputOrder")
	}
	return value, nil
}

// VC1/WMV video decoder based on Intel Media SDK
type GstMsdkVC1Dec struct {
	*GstMsdkDec
}

func NewMsdkVC1Dec() (*GstMsdkVC1Dec, error) {
	e, err := gst.NewElement("msdkvc1dec")
	if err != nil {
		return nil, err
	}
	return &GstMsdkVC1Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewMsdkVC1DecWithName(name string) (*GstMsdkVC1Dec, error) {
	e, err := gst.NewElementWithName("msdkvc1dec", name)
	if err != nil {
		return nil, err
	}
	return &GstMsdkVC1Dec{GstMsdkDec: &GstMsdkDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Decoded frames output order
// Default: display (0)
func (e *GstMsdkVC1Dec) SetOutputOrder(value GstMsdkDecOutputOrder) error {
	e.Element.SetArg("output-order", string(value))
	return nil
}

func (e *GstMsdkVC1Dec) GetOutputOrder() (GstMsdkDecOutputOrder, error) {
	var value GstMsdkDecOutputOrder
	var ok bool
	v, err := e.Element.GetProperty("output-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkDecOutputOrder)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkDecOutputOrder")
	}
	return value, nil
}

type GstMsdkEncAdaptiveI string

const (
	GstMsdkEncAdaptiveIAuto GstMsdkEncAdaptiveI = "auto" // SDK desides what to do
	GstMsdkEncAdaptiveIOff  GstMsdkEncAdaptiveI = "off"  // Disable Adaptive I frame insertion
	GstMsdkEncAdaptiveIOn   GstMsdkEncAdaptiveI = "on"   // Enable Aaptive I frame insertion
)

type GstMsdkEncTrellisQuantization int

const (
	GstMsdkEncTrellisQuantizationNone GstMsdkEncTrellisQuantization = 0x00000000 // Disable for all frames
	GstMsdkEncTrellisQuantizationI    GstMsdkEncTrellisQuantization = 0x00000002 // Enable for I frames
	GstMsdkEncTrellisQuantizationP    GstMsdkEncTrellisQuantization = 0x00000004 // Enable for P frames
	GstMsdkEncTrellisQuantizationB    GstMsdkEncTrellisQuantization = 0x00000008 // Enable for B frames
)

type GstMsdkEncTuneMode string

const (
	GstMsdkEncTuneModeAuto     GstMsdkEncTuneMode = "auto"      // Auto
	GstMsdkEncTuneModeNone     GstMsdkEncTuneMode = "none"      // None
	GstMsdkEncTuneModeLowPower GstMsdkEncTuneMode = "low-power" // Low power mode
)

type GstMsdkVPPDeinterlaceMethod string

const (
	GstMsdkVPPDeinterlaceMethodNone          GstMsdkVPPDeinterlaceMethod = "none"             // Disable deinterlacing
	GstMsdkVPPDeinterlaceMethodBob           GstMsdkVPPDeinterlaceMethod = "bob"              // Bob deinterlacing
	GstMsdkVPPDeinterlaceMethodAdvanced      GstMsdkVPPDeinterlaceMethod = "advanced"         // Advanced deinterlacing (Motion adaptive)
	GstMsdkVPPDeinterlaceMethodAdvancedNoRef GstMsdkVPPDeinterlaceMethod = "advanced-no-ref " // Advanced deinterlacing mode without using of reference frames
	GstMsdkVPPDeinterlaceMethodAdvancedScd   GstMsdkVPPDeinterlaceMethod = "advanced-scd "    // Advanced deinterlacing mode with scene change detection
	GstMsdkVPPDeinterlaceMethodFieldWeave    GstMsdkVPPDeinterlaceMethod = "field-weave"      // Field weaving
)

type GstMsdkVPPDeinterlaceMode string

const (
	GstMsdkVPPDeinterlaceModeAuto       GstMsdkVPPDeinterlaceMode = "auto"       // Auto detection
	GstMsdkVPPDeinterlaceModeInterlaced GstMsdkVPPDeinterlaceMode = "interlaced" // Force deinterlacing
	GstMsdkVPPDeinterlaceModeDisabled   GstMsdkVPPDeinterlaceMode = "disabled"   // Never deinterlace
)

type GstMsdkVPPScalingMode string

const (
	GstMsdkVPPScalingModeDisable  GstMsdkVPPScalingMode = "disable"  // Default Scaling
	GstMsdkVPPScalingModeLowpower GstMsdkVPPScalingMode = "lowpower" // Lowpower Scaling
	GstMsdkVPPScalingModeQuality  GstMsdkVPPScalingMode = "quality"  // High Quality Scaling
)

type GstMsdkEncAdaptiveB string

const (
	GstMsdkEncAdaptiveBAuto GstMsdkEncAdaptiveB = "auto" // SDK desides what to do
	GstMsdkEncAdaptiveBOff  GstMsdkEncAdaptiveB = "off"  // Disable Adaptive B-Frame insertion
	GstMsdkEncAdaptiveBOn   GstMsdkEncAdaptiveB = "on"   // Enable Aaptive B-Frame insertion
)

type GstMsdkEncMbBitrateControl string

const (
	GstMsdkEncMbBitrateControlAuto GstMsdkEncMbBitrateControl = "auto" // SDK desides what to do
	GstMsdkEncMbBitrateControlOff  GstMsdkEncMbBitrateControl = "off"  // Disable Macroblock level bit rate control
	GstMsdkEncMbBitrateControlOn   GstMsdkEncMbBitrateControl = "on"   // Enable Macroblock level bit rate control
)

type GstMsdkH264EncFramePacking string

const (
	GstMsdkH264EncFramePackingNone       GstMsdkH264EncFramePacking = "none"         // None (default)
	GstMsdkH264EncFramePackingSideBySide GstMsdkH264EncFramePacking = "side-by-side" // Side by Side
	GstMsdkH264EncFramePackingTopBottom  GstMsdkH264EncFramePacking = "top-bottom"   // Top Bottom
)

type GstMsdkDecOutputOrder string

const (
	GstMsdkDecOutputOrderDisplay GstMsdkDecOutputOrder = "display" // Output frames in Display order
	GstMsdkDecOutputOrderDecoded GstMsdkDecOutputOrder = "decoded" // Output frames in Decoded order
)

type GstMsdkEncRateControl string

const (
	GstMsdkEncRateControlCbr   GstMsdkEncRateControl = "cbr"    // Constant Bitrate
	GstMsdkEncRateControlVbr   GstMsdkEncRateControl = "vbr"    // Variable Bitrate
	GstMsdkEncRateControlCqp   GstMsdkEncRateControl = "cqp"    // Constant Quantizer
	GstMsdkEncRateControlAvbr  GstMsdkEncRateControl = "avbr"   // Average Bitrate
	GstMsdkEncRateControlLaVbr GstMsdkEncRateControl = "la_vbr" // VBR with look ahead (Non HRD compliant)
	GstMsdkEncRateControlIcq   GstMsdkEncRateControl = "icq"    // Intelligent CQP
	GstMsdkEncRateControlVcm   GstMsdkEncRateControl = "vcm"    // Video Conferencing Mode (Non HRD compliant)
	GstMsdkEncRateControlLaIcq GstMsdkEncRateControl = "la_icq" // Intelligent CQP with LA (Non HRD compliant)
	GstMsdkEncRateControlLaHrd GstMsdkEncRateControl = "la_hrd" // HRD compliant LA
	GstMsdkEncRateControlQvbr  GstMsdkEncRateControl = "qvbr"   // VBR with CQP
)

type GstMsdkEncTransformSkip string

const (
	GstMsdkEncTransformSkipAuto GstMsdkEncTransformSkip = "auto" // SDK desides what to do
	GstMsdkEncTransformSkipOff  GstMsdkEncTransformSkip = "off"  // transform_skip_enabled_flag will be set to 0 in PPS
	GstMsdkEncTransformSkipOn   GstMsdkEncTransformSkip = "on"   // transform_skip_enabled_flag will be set to 1 in PPS
)

type GstMsdkDec struct {
	*GstVideoDecoder
}

// Depth of asynchronous pipeline
// Default: 1, Min: 1, Max: 20
func (e *GstMsdkDec) SetAsyncDepth(value uint) error {
	return e.Element.SetProperty("async-depth", value)
}

func (e *GstMsdkDec) GetAsyncDepth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("async-depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable hardware decoders
// Default: true
func (e *GstMsdkDec) SetHardware(value bool) error {
	return e.Element.SetProperty("hardware", value)
}

func (e *GstMsdkDec) GetHardware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hardware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstMsdkEnc struct {
	*GstVideoEncoder
}

// The AVBR Accuracy in the unit of tenth of percent
// Default: 0, Min: 0, Max: 65535
func (e *GstMsdkEnc) SetAccuracy(value uint) error {
	return e.Element.SetProperty("accuracy", value)
}

func (e *GstMsdkEnc) GetAccuracy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("accuracy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Adaptive B-Frame Insertion control
// Default: off (32)
func (e *GstMsdkEnc) SetBAdapt(value GstMsdkEncAdaptiveB) error {
	e.Element.SetArg("b-adapt", string(value))
	return nil
}

func (e *GstMsdkEnc) GetBAdapt() (GstMsdkEncAdaptiveB, error) {
	var value GstMsdkEncAdaptiveB
	var ok bool
	v, err := e.Element.GetProperty("b-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncAdaptiveB)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncAdaptiveB")
	}
	return value, nil
}

// Number of B frames between I and P frames
// Default: 0, Min: 0, Max: 2147483647
func (e *GstMsdkEnc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstMsdkEnc) GetBFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("b-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// GOP Size
// Default: 256, Min: 0, Max: 2147483647
func (e *GstMsdkEnc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstMsdkEnc) GetGopSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("gop-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum bitrate(kbit/sec) at which data enters Video Buffering Verifier (0: auto-calculate)
// Default: 0, Min: 0, Max: 2048000
func (e *GstMsdkEnc) SetMaxVbvBitrate(value uint) error {
	return e.Element.SetProperty("max-vbv-bitrate", value)
}

func (e *GstMsdkEnc) GetMaxVbvBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-vbv-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Depth of asynchronous pipeline
// Default: 4, Min: 0, Max: 20
func (e *GstMsdkEnc) SetAsyncDepth(value uint) error {
	return e.Element.SetProperty("async-depth", value)
}

func (e *GstMsdkEnc) GetAsyncDepth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("async-depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The properties for the external coding

func (e *GstMsdkEnc) SetExtCodingProps(value *gst.Structure) error {
	return e.Element.SetProperty("ext-coding-props", value)
}

func (e *GstMsdkEnc) GetExtCodingProps() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("ext-coding-props")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Constant quantizer for P frames (0 unlimited)
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkEnc) SetQpp(value uint) error {
	return e.Element.SetProperty("qpp", value)
}

func (e *GstMsdkEnc) GetQpp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qpp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of frames to look ahead for Rate control
// Default: 10, Min: 10, Max: 100
func (e *GstMsdkEnc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstMsdkEnc) GetRcLookahead() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of reference frames
// Default: 1, Min: 0, Max: 2147483647
func (e *GstMsdkEnc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

func (e *GstMsdkEnc) GetRefFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ref-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Bitrate in kbit/sec
// Default: 2048, Min: 1, Max: 2048000
func (e *GstMsdkEnc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstMsdkEnc) GetBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The AVBR Convergence in the unit of 100 frames
// Default: 0, Min: 0, Max: 65535
func (e *GstMsdkEnc) SetConvergence(value uint) error {
	return e.Element.SetProperty("convergence", value)
}

func (e *GstMsdkEnc) GetConvergence() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("convergence")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable hardware encoders
// Default: true
func (e *GstMsdkEnc) SetHardware(value bool) error {
	return e.Element.SetProperty("hardware", value)
}

func (e *GstMsdkEnc) GetHardware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hardware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of I frames between IDR frames
// Default: 0, Min: 0, Max: 2147483647
func (e *GstMsdkEnc) SetIFrames(value uint) error {
	return e.Element.SetProperty("i-frames", value)
}

func (e *GstMsdkEnc) GetIFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("i-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate control method
// Default: cbr (1)
func (e *GstMsdkEnc) SetRateControl(value GstMsdkEncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstMsdkEnc) GetRateControl() (GstMsdkEncRateControl, error) {
	var value GstMsdkEncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncRateControl")
	}
	return value, nil
}

// Constant quantizer for I frames (0 unlimited). Also used as ICQQuality or QVBRQuality for different RateControl methods
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkEnc) SetQpi(value uint) error {
	return e.Element.SetProperty("qpi", value)
}

func (e *GstMsdkEnc) GetQpi() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qpi")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// 1: Best quality, 4: Balanced, 7: Best speed
// Default: 4, Min: 1, Max: 7
func (e *GstMsdkEnc) SetTargetUsage(value uint) error {
	return e.Element.SetProperty("target-usage", value)
}

func (e *GstMsdkEnc) GetTargetUsage() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Adaptive I-Frame Insertion control
// Default: off (32)
func (e *GstMsdkEnc) SetIAdapt(value GstMsdkEncAdaptiveI) error {
	e.Element.SetArg("i-adapt", string(value))
	return nil
}

func (e *GstMsdkEnc) GetIAdapt() (GstMsdkEncAdaptiveI, error) {
	var value GstMsdkEncAdaptiveI
	var ok bool
	v, err := e.Element.GetProperty("i-adapt")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncAdaptiveI)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncAdaptiveI")
	}
	return value, nil
}

// Maximum possible size (in kbyte) of any compressed frames (0: auto-calculate)
// Default: 0, Min: 0, Max: 65535
func (e *GstMsdkEnc) SetMaxFrameSize(value uint) error {
	return e.Element.SetProperty("max-frame-size", value)
}

func (e *GstMsdkEnc) GetMaxFrameSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-frame-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Macroblock level bitrate control
// Default: off (32)
func (e *GstMsdkEnc) SetMbbrc(value GstMsdkEncMbBitrateControl) error {
	e.Element.SetArg("mbbrc", string(value))
	return nil
}

func (e *GstMsdkEnc) GetMbbrc() (GstMsdkEncMbBitrateControl, error) {
	var value GstMsdkEncMbBitrateControl
	var ok bool
	v, err := e.Element.GetProperty("mbbrc")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkEncMbBitrateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkEncMbBitrateControl")
	}
	return value, nil
}

// Number of slices per frame, Zero tells the encoder to choose any slice partitioning allowed by the codec standard
// Default: 0, Min: 0, Max: 2147483647
func (e *GstMsdkEnc) SetNumSlices(value uint) error {
	return e.Element.SetProperty("num-slices", value)
}

func (e *GstMsdkEnc) GetNumSlices() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slices")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Constant quantizer for B frames (0 unlimited)
// Default: 0, Min: 0, Max: 51
func (e *GstMsdkEnc) SetQpb(value uint) error {
	return e.Element.SetProperty("qpb", value)
}

func (e *GstMsdkEnc) GetQpb() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qpb")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstMsdkEncIntraRefreshType string

const (
	GstMsdkEncIntraRefreshTypeNo         GstMsdkEncIntraRefreshType = "no"         // No (default)
	GstMsdkEncIntraRefreshTypeVertical   GstMsdkEncIntraRefreshType = "vertical"   // Vertical
	GstMsdkEncIntraRefreshTypeHorizontal GstMsdkEncIntraRefreshType = "horizontal" // Horizontal
	GstMsdkEncIntraRefreshTypeSlice      GstMsdkEncIntraRefreshType = "slice"      // Slice
)

type GstMsdkEncRCLookAheadDownsampling string

const (
	GstMsdkEncRCLookAheadDownsamplingDefault GstMsdkEncRCLookAheadDownsampling = "default" // SDK desides what to do
	GstMsdkEncRCLookAheadDownsamplingOff     GstMsdkEncRCLookAheadDownsampling = "off"     // No downsampling
	GstMsdkEncRCLookAheadDownsampling2x      GstMsdkEncRCLookAheadDownsampling = "2x"      // Down sample 2-times before estimation
	GstMsdkEncRCLookAheadDownsampling4x      GstMsdkEncRCLookAheadDownsampling = "4x"      // Down sample 4-times before estimation
)

type GstMsdkVPPFrcAlgorithm string

const (
	GstMsdkVPPFrcAlgorithmNone                    GstMsdkVPPFrcAlgorithm = "none"                      // No FrameRate Control algorithm
	GstMsdkVPPFrcAlgorithmPreserveTs              GstMsdkVPPFrcAlgorithm = "preserve-ts"               // Frame dropping/repetition, Preserve timestamp
	GstMsdkVPPFrcAlgorithmDistributeTs            GstMsdkVPPFrcAlgorithm = "distribute-ts"             // Frame dropping/repetition, Distribute timestamp
	GstMsdkVPPFrcAlgorithmInterpolate             GstMsdkVPPFrcAlgorithm = "interpolate"               // Frame interpolation
	GstMsdkVPPFrcAlgorithmInterpolatePreserveTs   GstMsdkVPPFrcAlgorithm = "interpolate-preserve-ts"   // Frame interpolation, Preserve timestamp
	GstMsdkVPPFrcAlgorithmInterpolateDistributeTs GstMsdkVPPFrcAlgorithm = "interpolate-distribute-ts" // Frame interpolation, Distribute timestamp
)

type GstMsdkVPPMirroring string

const (
	GstMsdkVPPMirroringDisable    GstMsdkVPPMirroring = "disable"    // Disable mirroring
	GstMsdkVPPMirroringHorizontal GstMsdkVPPMirroring = "horizontal" // Horizontal Mirroring
	GstMsdkVPPMirroringVertical   GstMsdkVPPMirroring = "vertical"   // Vertical Mirroring
)

type GstMsdkVPPRotation string

const (
	GstMsdkVPPRotation0   GstMsdkVPPRotation = "0"   // Unrotated mode
	GstMsdkVPPRotation90  GstMsdkVPPRotation = "90"  // Rotated by 90°
	GstMsdkVPPRotation180 GstMsdkVPPRotation = "180" // Rotated by 180°
	GstMsdkVPPRotation270 GstMsdkVPPRotation = "270" // Rotated by 270°
)
