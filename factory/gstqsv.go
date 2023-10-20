// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Intel Quick Sync Video H.264 Decoder
type GstQsvH264Dec struct {
	*GstQsvDecoder
}

func NewQsvH264Dec() (*GstQsvH264Dec, error) {
	e, err := gst.NewElement("qsvh264dec")
	if err != nil {
		return nil, err
	}
	return &GstQsvH264Dec{GstQsvDecoder: &GstQsvDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewQsvH264DecWithName(name string) (*GstQsvH264Dec, error) {
	e, err := gst.NewElementWithName("qsvh264dec", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvH264Dec{GstQsvDecoder: &GstQsvDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Intel Quick Sync Video H.265 Encoder
type GstQsvH265Enc struct {
	*GstQsvEncoder
}

func NewQsvH265Enc() (*GstQsvH265Enc, error) {
	e, err := gst.NewElement("qsvh265enc")
	if err != nil {
		return nil, err
	}
	return &GstQsvH265Enc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewQsvH265EncWithName(name string) (*GstQsvH265Enc, error) {
	e, err := gst.NewElementWithName("qsvh265enc", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvH265Enc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Closed Caption Insert mode. Only CEA-708 RAW format is supported for now
// Default: insert (0)
func (e *GstQsvH265Enc) SetCcInsert(value GstQsvH265EncSeiInsertMode) error {
	e.Element.SetArg("cc-insert", string(value))
	return nil
}

func (e *GstQsvH265Enc) GetCcInsert() (GstQsvH265EncSeiInsertMode, error) {
	var value GstQsvH265EncSeiInsertMode
	var ok bool
	v, err := e.Element.GetProperty("cc-insert")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH265EncSeiInsertMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH265EncSeiInsertMode")
	}
	return value, nil
}

// Minimum allowed QP value for I-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetMinQpI(value uint) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstQsvH265Enc) GetMinQpI() (uint, error) {
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

// Number of B frames between I and P frames
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH265Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstQsvH265Enc) GetBFrames() (uint, error) {
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

// Target bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")
// Default: 2000, Min: 0, Max: 2147483647
func (e *GstQsvH265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstQsvH265Enc) GetBitrate() (uint, error) {
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

// Constant quantizer for I frames (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstQsvH265Enc) GetQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of reference frames (0: unspecified)
// Default: 2, Min: 0, Max: 16
func (e *GstQsvH265Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

func (e *GstQsvH265Enc) GetRefFrames() (uint, error) {
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

// IDR-frame interval in terms of I-frames. 0: only first I-frame is is an IDR frame, 1: every I-frame is an IDR frame, N: "N - 1" I-frames are inserted between IDR-frames
// Default: 1, Min: 0, Max: 65535
func (e *GstQsvH265Enc) SetIdrInterval(value uint) error {
	return e.Element.SetProperty("idr-interval", value)
}

func (e *GstQsvH265Enc) GetIdrInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("idr-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum allowed QP value for B-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetMaxQpB(value uint) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstQsvH265Enc) GetMaxQpB() (uint, error) {
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

// Intelligent Constant Quality for "icq" rate-control (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetIcqQuality(value uint) error {
	return e.Element.SetProperty("icq-quality", value)
}

func (e *GstQsvH265Enc) GetIcqQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("icq-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")
// Default: 0, Min: 0, Max: 2147483647
func (e *GstQsvH265Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstQsvH265Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum allowed QP value for B-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetMinQpB(value uint) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstQsvH265Enc) GetMinQpB() (uint, error) {
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

// Constant quantizer for B frames (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetQpB(value uint) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstQsvH265Enc) GetQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Constant quantizer for P frames (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstQsvH265Enc) GetQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Allow NAL HRD non-conformant stream
// Default: false
func (e *GstQsvH265Enc) SetDisableHrdConformance(value bool) error {
	return e.Element.SetProperty("disable-hrd-conformance", value)
}

func (e *GstQsvH265Enc) GetDisableHrdConformance() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("disable-hrd-conformance")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of pictures within a GOP (0: unspecified)
// Default: 30, Min: 0, Max: 65535
func (e *GstQsvH265Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstQsvH265Enc) GetGopSize() (uint, error) {
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

// Minimum allowed QP value for P-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetMinQpP(value uint) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstQsvH265Enc) GetMinQpP() (uint, error) {
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

// Quality level used for "qvbr" rate-control mode (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetQvbrQuality(value uint) error {
	return e.Element.SetProperty("qvbr-quality", value)
}

func (e *GstQsvH265Enc) GetQvbrQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qvbr-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (2)
func (e *GstQsvH265Enc) SetRateControl(value GstQsvH265EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstQsvH265Enc) GetRateControl() (GstQsvH265EncRateControl, error) {
	var value GstQsvH265EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH265EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH265EncRateControl")
	}
	return value, nil
}

// Maximum allowed QP value for I-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetMaxQpI(value uint) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstQsvH265Enc) GetMaxQpI() (uint, error) {
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

// Maximum allowed QP value for P-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH265Enc) SetMaxQpP(value uint) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstQsvH265Enc) GetMaxQpP() (uint, error) {
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

// Intel Quick Sync Video JPEG Encoder
type GstQsvJpegEnc struct {
	*GstQsvEncoder
}

func NewQsvJpegEnc() (*GstQsvJpegEnc, error) {
	e, err := gst.NewElement("qsvjpegenc")
	if err != nil {
		return nil, err
	}
	return &GstQsvJpegEnc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewQsvJpegEncWithName(name string) (*GstQsvJpegEnc, error) {
	e, err := gst.NewElementWithName("qsvjpegenc", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvJpegEnc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Encoding quality, 100 for best quality
// Default: 85, Min: 1, Max: 100
func (e *GstQsvJpegEnc) SetQuality(value uint) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstQsvJpegEnc) GetQuality() (uint, error) {
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

// Intel Quick Sync Video VP9 Decoder
type GstQsvVP9Dec struct {
	*GstQsvDecoder
}

func NewQsvVP9Dec() (*GstQsvVP9Dec, error) {
	e, err := gst.NewElement("qsvvp9dec")
	if err != nil {
		return nil, err
	}
	return &GstQsvVP9Dec{GstQsvDecoder: &GstQsvDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewQsvVP9DecWithName(name string) (*GstQsvVP9Dec, error) {
	e, err := gst.NewElementWithName("qsvvp9dec", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvVP9Dec{GstQsvDecoder: &GstQsvDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Intel Quick Sync Video VP9 Encoder
type GstQsvVP9Enc struct {
	*GstQsvEncoder
}

func NewQsvVP9Enc() (*GstQsvVP9Enc, error) {
	e, err := gst.NewElement("qsvvp9enc")
	if err != nil {
		return nil, err
	}
	return &GstQsvVP9Enc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewQsvVP9EncWithName(name string) (*GstQsvVP9Enc, error) {
	e, err := gst.NewElementWithName("qsvvp9enc", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvVP9Enc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Maximum bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvVP9Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstQsvVP9Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Constant quantizer for I frames (0: default)
// Default: 0, Min: 0, Max: 255
func (e *GstQsvVP9Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstQsvVP9Enc) GetQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Constant quantizer for P frames (0: default)
// Default: 0, Min: 0, Max: 255
func (e *GstQsvVP9Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstQsvVP9Enc) GetQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (2)
func (e *GstQsvVP9Enc) SetRateControl(value GstQsvVP9EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstQsvVP9Enc) GetRateControl() (GstQsvVP9EncRateControl, error) {
	var value GstQsvVP9EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvVP9EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvVP9EncRateControl")
	}
	return value, nil
}

// Number of reference frames (0: unspecified)
// Default: 1, Min: 0, Max: 3
func (e *GstQsvVP9Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

func (e *GstQsvVP9Enc) GetRefFrames() (uint, error) {
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

// Target bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")
// Default: 2000, Min: 0, Max: 65535
func (e *GstQsvVP9Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstQsvVP9Enc) GetBitrate() (uint, error) {
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

// Number of pictures within a GOP (0: unspecified)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstQsvVP9Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstQsvVP9Enc) GetGopSize() (uint, error) {
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

// Intelligent Constant Quality for "icq" rate-control (0: default)
// Default: 0, Min: 0, Max: 255
func (e *GstQsvVP9Enc) SetIcqQuality(value uint) error {
	return e.Element.SetProperty("icq-quality", value)
}

func (e *GstQsvVP9Enc) GetIcqQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("icq-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Intel Quick Sync Video AV1 Encoder
type GstQsvAV1Enc struct {
	*GstQsvEncoder
}

func NewQsvAV1Enc() (*GstQsvAV1Enc, error) {
	e, err := gst.NewElement("qsvav1enc")
	if err != nil {
		return nil, err
	}
	return &GstQsvAV1Enc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewQsvAV1EncWithName(name string) (*GstQsvAV1Enc, error) {
	e, err := gst.NewElementWithName("qsvav1enc", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvAV1Enc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Constant quantizer for I frames (0: default)
// Default: 0, Min: 0, Max: 255
func (e *GstQsvAV1Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstQsvAV1Enc) GetQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Constant quantizer for P frames (0: default)
// Default: 0, Min: 0, Max: 255
func (e *GstQsvAV1Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstQsvAV1Enc) GetQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (2)
func (e *GstQsvAV1Enc) SetRateControl(value GstQsvAV1EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstQsvAV1Enc) GetRateControl() (GstQsvAV1EncRateControl, error) {
	var value GstQsvAV1EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvAV1EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvAV1EncRateControl")
	}
	return value, nil
}

// Number of reference frames (0: unspecified)
// Default: 1, Min: 0, Max: 3
func (e *GstQsvAV1Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

func (e *GstQsvAV1Enc) GetRefFrames() (uint, error) {
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

// Target bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")
// Default: 2000, Min: 0, Max: 65535
func (e *GstQsvAV1Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstQsvAV1Enc) GetBitrate() (uint, error) {
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

// Number of pictures within a GOP (0: unspecified)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstQsvAV1Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstQsvAV1Enc) GetGopSize() (uint, error) {
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

// Maximum bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp" and "icq")
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvAV1Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstQsvAV1Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Intel Quick Sync Video H.264 Encoder
type GstQsvH264Enc struct {
	*GstQsvEncoder
}

func NewQsvH264Enc() (*GstQsvH264Enc, error) {
	e, err := gst.NewElement("qsvh264enc")
	if err != nil {
		return nil, err
	}
	return &GstQsvH264Enc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewQsvH264EncWithName(name string) (*GstQsvH264Enc, error) {
	e, err := gst.NewElementWithName("qsvh264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvH264Enc{GstQsvEncoder: &GstQsvEncoder{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// Number of slices for B frame
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetNumSliceB(value uint) error {
	return e.Element.SetProperty("num-slice-b", value)
}

func (e *GstQsvH264Enc) GetNumSliceB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slice-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Constant quantizer for P frames (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetQpP(value uint) error {
	return e.Element.SetProperty("qp-p", value)
}

func (e *GstQsvH264Enc) GetQpP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Target bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp", "icq", and "la_icq")
// Default: 2000, Min: 0, Max: 2147483647
func (e *GstQsvH264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstQsvH264Enc) GetBitrate() (uint, error) {
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

// Maximum encoded I frame size in bytes, used for VBR based bitrate control modes
// Default: 0, Min: 0, Max: -1
func (e *GstQsvH264Enc) SetMaxFrameSizeI(value uint) error {
	return e.Element.SetProperty("max-frame-size-i", value)
}

func (e *GstQsvH264Enc) GetMaxFrameSizeI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-frame-size-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum slice size in bytes. If this parameter is specified other controls over number of slices are ignored
// Default: 0, Min: 0, Max: -1
func (e *GstQsvH264Enc) SetMaxSliceSize(value uint) error {
	return e.Element.SetProperty("max-slice-size", value)
}

func (e *GstQsvH264Enc) GetMaxSliceSize() (uint, error) {
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

// Number of slices for P frame
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetNumSliceP(value uint) error {
	return e.Element.SetProperty("num-slice-p", value)
}

func (e *GstQsvH264Enc) GetNumSliceP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slice-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum bitrate in kbit/sec, Ignored when selected rate-control mode is constant QP variants (i.e., "cqp", "icq", and "la_icq")
// Default: 0, Min: 0, Max: 2147483647
func (e *GstQsvH264Enc) SetMaxBitrate(value uint) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstQsvH264Enc) GetMaxBitrate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum encoded P and B frame size in bytes, used for VBR based bitrate control modes. "max-frame-size-i" must be non-zero, otherwise this propert will be ignored
// Default: 0, Min: 0, Max: -1
func (e *GstQsvH264Enc) SetMaxFrameSizeP(value uint) error {
	return e.Element.SetProperty("max-frame-size-p", value)
}

func (e *GstQsvH264Enc) GetMaxFrameSizeP() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-frame-size-p")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Closed Caption Insert mode. Only CEA-708 RAW format is supported for now
// Default: insert (0)
func (e *GstQsvH264Enc) SetCcInsert(value GstQsvH264EncSeiInsertMode) error {
	e.Element.SetArg("cc-insert", string(value))
	return nil
}

func (e *GstQsvH264Enc) GetCcInsert() (GstQsvH264EncSeiInsertMode, error) {
	var value GstQsvH264EncSeiInsertMode
	var ok bool
	v, err := e.Element.GetProperty("cc-insert")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH264EncSeiInsertMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH264EncSeiInsertMode")
	}
	return value, nil
}

// Number of slices in each video frame
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetNumSlice(value uint) error {
	return e.Element.SetProperty("num-slice", value)
}

func (e *GstQsvH264Enc) GetNumSlice() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slice")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Constant quantizer for B frames (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetQpB(value uint) error {
	return e.Element.SetProperty("qp-b", value)
}

func (e *GstQsvH264Enc) GetQpB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of reference frames (0: unspecified)
// Default: 2, Min: 0, Max: 16
func (e *GstQsvH264Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

func (e *GstQsvH264Enc) GetRefFrames() (uint, error) {
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

// AVBR Accuracy in the unit of tenth of percent
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetAvbrAccuracy(value uint) error {
	return e.Element.SetProperty("avbr-accuracy", value)
}

func (e *GstQsvH264Enc) GetAvbrAccuracy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("avbr-accuracy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of B frames between I and P frames
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstQsvH264Enc) GetBFrames() (uint, error) {
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

// Maximum allowed QP value for P-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetMaxQpP(value uint) error {
	return e.Element.SetProperty("max-qp-p", value)
}

func (e *GstQsvH264Enc) GetMaxQpP() (uint, error) {
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

// Minimum allowed QP value for B-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetMinQpB(value uint) error {
	return e.Element.SetProperty("min-qp-b", value)
}

func (e *GstQsvH264Enc) GetMinQpB() (uint, error) {
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

// Minimum allowed QP value for P-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetMinQpP(value uint) error {
	return e.Element.SetProperty("min-qp-p", value)
}

func (e *GstQsvH264Enc) GetMinQpP() (uint, error) {
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

// Constant quantizer for I frames (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetQpI(value uint) error {
	return e.Element.SetProperty("qp-i", value)
}

func (e *GstQsvH264Enc) GetQpI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qp-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Trellis quantization mode
// Default: unknown
func (e *GstQsvH264Enc) SetTrellis(value GstQsvH264Trellis) error {
	e.Element.SetArg("trellis", fmt.Sprint(value))
	return nil
}

func (e *GstQsvH264Enc) GetTrellis() (GstQsvH264Trellis, error) {
	var value GstQsvH264Trellis
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH264Trellis)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH264Trellis")
	}
	return value, nil
}

// Allow NAL HRD non-conformant stream
// Default: false
func (e *GstQsvH264Enc) SetDisableHrdConformance(value bool) error {
	return e.Element.SetProperty("disable-hrd-conformance", value)
}

func (e *GstQsvH264Enc) GetDisableHrdConformance() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("disable-hrd-conformance")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum encoded frame size in bytes, used for VBR based bitrate control modes
// Default: 0, Min: 0, Max: -1
func (e *GstQsvH264Enc) SetMaxFrameSize(value uint) error {
	return e.Element.SetProperty("max-frame-size", value)
}

func (e *GstQsvH264Enc) GetMaxFrameSize() (uint, error) {
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

// Downsampling method in look-ahead rate control
// Default: unknown (0)
func (e *GstQsvH264Enc) SetRcLookaheadDs(value GstQsvH264EncRCLookAheadDS) error {
	e.Element.SetArg("rc-lookahead-ds", string(value))
	return nil
}

func (e *GstQsvH264Enc) GetRcLookaheadDs() (GstQsvH264EncRCLookAheadDS, error) {
	var value GstQsvH264EncRCLookAheadDS
	var ok bool
	v, err := e.Element.GetProperty("rc-lookahead-ds")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH264EncRCLookAheadDS)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH264EncRCLookAheadDS")
	}
	return value, nil
}

// Enables CABAC entropy coding
// Default: unknown (0)
func (e *GstQsvH264Enc) SetCabac(value GstQsvCodingOption) error {
	e.Element.SetArg("cabac", string(value))
	return nil
}

func (e *GstQsvH264Enc) GetCabac() (GstQsvCodingOption, error) {
	var value GstQsvCodingOption
	var ok bool
	v, err := e.Element.GetProperty("cabac")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvCodingOption)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvCodingOption")
	}
	return value, nil
}

// Maximum allowed QP value for I-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetMaxQpI(value uint) error {
	return e.Element.SetProperty("max-qp-i", value)
}

func (e *GstQsvH264Enc) GetMaxQpI() (uint, error) {
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

// Intelligent Constant Quality for "icq" rate-control (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetIcqQuality(value uint) error {
	return e.Element.SetProperty("icq-quality", value)
}

func (e *GstQsvH264Enc) GetIcqQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("icq-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of slices for I frame
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetNumSliceI(value uint) error {
	return e.Element.SetProperty("num-slice-i", value)
}

func (e *GstQsvH264Enc) GetNumSliceI() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-slice-i")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// AVBR Convergence in the unit of 100 frames
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetAvbrConvergence(value uint) error {
	return e.Element.SetProperty("avbr-convergence", value)
}

func (e *GstQsvH264Enc) GetAvbrConvergence() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("avbr-convergence")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum allowed QP value for I-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetMinQpI(value uint) error {
	return e.Element.SetProperty("min-qp-i", value)
}

func (e *GstQsvH264Enc) GetMinQpI() (uint, error) {
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

// Maximum allowed QP value for B-frame types (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetMaxQpB(value uint) error {
	return e.Element.SetProperty("max-qp-b", value)
}

func (e *GstQsvH264Enc) GetMaxQpB() (uint, error) {
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

// Quality level used for "qvbr" rate-control mode (0: default)
// Default: 0, Min: 0, Max: 51
func (e *GstQsvH264Enc) SetQvbrQuality(value uint) error {
	return e.Element.SetProperty("qvbr-quality", value)
}

func (e *GstQsvH264Enc) GetQvbrQuality() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("qvbr-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Rate Control Method
// Default: vbr (2)
func (e *GstQsvH264Enc) SetRateControl(value GstQsvH264EncRateControl) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstQsvH264Enc) GetRateControl() (GstQsvH264EncRateControl, error) {
	var value GstQsvH264EncRateControl
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQsvH264EncRateControl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQsvH264EncRateControl")
	}
	return value, nil
}

// Number of frames to look ahead for Rate Control, used for "la_vbr", "la_icq", and "la_hrd" rate-control modes
// Default: 10, Min: 10, Max: 100
func (e *GstQsvH264Enc) SetRcLookahead(value uint) error {
	return e.Element.SetProperty("rc-lookahead", value)
}

func (e *GstQsvH264Enc) GetRcLookahead() (uint, error) {
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

// Number of pictures within a GOP (0: unspecified)
// Default: 30, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetGopSize(value uint) error {
	return e.Element.SetProperty("gop-size", value)
}

func (e *GstQsvH264Enc) GetGopSize() (uint, error) {
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

// IDR-frame interval in terms of I-frames. 0: every I-frame is an IDR frame, N: "N" I-frames are inserted between IDR-frames
// Default: 0, Min: 0, Max: 65535
func (e *GstQsvH264Enc) SetIdrInterval(value uint) error {
	return e.Element.SetProperty("idr-interval", value)
}

func (e *GstQsvH264Enc) GetIdrInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("idr-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Intel Quick Sync Video H.265 Decoder
type GstQsvH265Dec struct {
	*GstQsvDecoder
}

func NewQsvH265Dec() (*GstQsvH265Dec, error) {
	e, err := gst.NewElement("qsvh265dec")
	if err != nil {
		return nil, err
	}
	return &GstQsvH265Dec{GstQsvDecoder: &GstQsvDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewQsvH265DecWithName(name string) (*GstQsvH265Dec, error) {
	e, err := gst.NewElementWithName("qsvh265dec", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvH265Dec{GstQsvDecoder: &GstQsvDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// Intel Quick Sync Video JPEG Decoder
type GstQsvJpegDec struct {
	*GstQsvDecoder
}

func NewQsvJpegDec() (*GstQsvJpegDec, error) {
	e, err := gst.NewElement("qsvjpegdec")
	if err != nil {
		return nil, err
	}
	return &GstQsvJpegDec{GstQsvDecoder: &GstQsvDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewQsvJpegDecWithName(name string) (*GstQsvJpegDec, error) {
	e, err := gst.NewElementWithName("qsvjpegdec", name)
	if err != nil {
		return nil, err
	}
	return &GstQsvJpegDec{GstQsvDecoder: &GstQsvDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

type GstQsvH264EncSeiInsertMode string

const (
	GstQsvH264EncSeiInsertModeInsert        GstQsvH264EncSeiInsertMode = "insert"          // Insert SEI
	GstQsvH264EncSeiInsertModeInsertAndDrop GstQsvH264EncSeiInsertMode = "insert-and-drop" // Insert SEI and remove corresponding meta from output buffer
	GstQsvH264EncSeiInsertModeDisabled      GstQsvH264EncSeiInsertMode = "disabled"        // Disable SEI insertion
)

type GstQsvH264Trellis int

const (
	GstQsvH264TrellisUnknown GstQsvH264Trellis = 0x00000000 // Unknown
	GstQsvH264TrellisOff     GstQsvH264Trellis = 0x00000001 // Disable for all frame types
	GstQsvH264TrellisI       GstQsvH264Trellis = 0x00000002 // Enable for I frames
	GstQsvH264TrellisP       GstQsvH264Trellis = 0x00000004 // Enable for P frames
	GstQsvH264TrellisB       GstQsvH264Trellis = 0x00000008 // Enable for B frames
)

type GstQsvH265EncRateControl string

const (
	GstQsvH265EncRateControlCbr  GstQsvH265EncRateControl = "cbr"  // Constant Bitrate
	GstQsvH265EncRateControlVbr  GstQsvH265EncRateControl = "vbr"  // Variable Bitrate
	GstQsvH265EncRateControlCqp  GstQsvH265EncRateControl = "cqp"  // Constant Quantizer
	GstQsvH265EncRateControlIcq  GstQsvH265EncRateControl = "icq"  // Intelligent CQP
	GstQsvH265EncRateControlVcm  GstQsvH265EncRateControl = "vcm"  // Video Conferencing Mode (Non HRD compliant)
	GstQsvH265EncRateControlQvbr GstQsvH265EncRateControl = "qvbr" // VBR with CQP
)

type GstQsvH265EncSeiInsertMode string

const (
	GstQsvH265EncSeiInsertModeInsert        GstQsvH265EncSeiInsertMode = "insert"          // Insert SEI
	GstQsvH265EncSeiInsertModeInsertAndDrop GstQsvH265EncSeiInsertMode = "insert-and-drop" // Insert SEI and remove corresponding meta from output buffer
	GstQsvH265EncSeiInsertModeDisabled      GstQsvH265EncSeiInsertMode = "disabled"        // Disable SEI insertion
)

type GstQsvAV1EncRateControl string

const (
	GstQsvAV1EncRateControlCbr GstQsvAV1EncRateControl = "cbr" // Constant Bitrate
	GstQsvAV1EncRateControlVbr GstQsvAV1EncRateControl = "vbr" // Variable Bitrate
	GstQsvAV1EncRateControlCqp GstQsvAV1EncRateControl = "cqp" // Constant Quantizer
)

type GstQsvCodingOption string

const (
	GstQsvCodingOptionUnknown GstQsvCodingOption = "unknown" // Unknown
	GstQsvCodingOptionOn      GstQsvCodingOption = "on"      // On
	GstQsvCodingOptionOff     GstQsvCodingOption = "off"     // Off
)

type GstQsvH264EncRateControl string

const (
	GstQsvH264EncRateControlCbr   GstQsvH264EncRateControl = "cbr"    // Constant Bitrate
	GstQsvH264EncRateControlVbr   GstQsvH264EncRateControl = "vbr"    // Variable Bitrate
	GstQsvH264EncRateControlCqp   GstQsvH264EncRateControl = "cqp"    // Constant Quantizer
	GstQsvH264EncRateControlAvbr  GstQsvH264EncRateControl = "avbr"   // Average Variable Bitrate
	GstQsvH264EncRateControlLaVbr GstQsvH264EncRateControl = "la-vbr" // VBR with look ahead (Non HRD compliant)
	GstQsvH264EncRateControlIcq   GstQsvH264EncRateControl = "icq"    // Intelligent CQP
	GstQsvH264EncRateControlVcm   GstQsvH264EncRateControl = "vcm"    // Video Conferencing Mode (Non HRD compliant)
	GstQsvH264EncRateControlLaIcq GstQsvH264EncRateControl = "la-icq" // Intelligent CQP with LA (Non HRD compliant)
	GstQsvH264EncRateControlLaHrd GstQsvH264EncRateControl = "la-hrd" // HRD compliant LA
	GstQsvH264EncRateControlQvbr  GstQsvH264EncRateControl = "qvbr"   // VBR with CQP
)

type GstQsvVP9EncRateControl string

const (
	GstQsvVP9EncRateControlCbr GstQsvVP9EncRateControl = "cbr" // Constant Bitrate
	GstQsvVP9EncRateControlVbr GstQsvVP9EncRateControl = "vbr" // Variable Bitrate
	GstQsvVP9EncRateControlCqp GstQsvVP9EncRateControl = "cqp" // Constant Quantizer
	GstQsvVP9EncRateControlIcq GstQsvVP9EncRateControl = "icq" // Intelligent CQP
)

type GstQsvDecoder struct {
	*GstVideoDecoder
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstQsvDecoder) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

type GstQsvEncoder struct {
	*GstVideoEncoder
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstQsvEncoder) GetAdapterLuid() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("adapter-luid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Enables low-latency encoding
// Default: false
func (e *GstQsvEncoder) SetLowLatency(value bool) error {
	return e.Element.SetProperty("low-latency", value)
}

func (e *GstQsvEncoder) GetLowLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("low-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// 1: Best quality, 4: Balanced, 7: Best speed
// Default: 4, Min: 1, Max: 7
func (e *GstQsvEncoder) SetTargetUsage(value uint) error {
	return e.Element.SetProperty("target-usage", value)
}

func (e *GstQsvEncoder) GetTargetUsage() (uint, error) {
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

type GstQsvH264EncRCLookAheadDS string

const (
	GstQsvH264EncRCLookAheadDSUnknown GstQsvH264EncRCLookAheadDS = "unknown" // Unknown
	GstQsvH264EncRCLookAheadDSOff     GstQsvH264EncRCLookAheadDS = "off"     // Do not use down sampling
	GstQsvH264EncRCLookAheadDS2x      GstQsvH264EncRCLookAheadDS = "2x"      // Down sample frames two times before estimation
	GstQsvH264EncRCLookAheadDS4x      GstQsvH264EncRCLookAheadDS = "4x"      // Down sample frames four times before estimation
)
