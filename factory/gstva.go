// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// VA-API based H.264 video decoder
type GstVaH264Dec struct {
	*GstH264Decoder
}

func NewVaH264Dec() (*GstVaH264Dec, error) {
	e, err := gst.NewElement("vah264dec")
	if err != nil {
		return nil, err
	}
	return &GstVaH264Dec{GstH264Decoder: &GstH264Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVaH264DecWithName(name string) (*GstVaH264Dec, error) {
	e, err := gst.NewElementWithName("vah264dec", name)
	if err != nil {
		return nil, err
	}
	return &GstVaH264Dec{GstH264Decoder: &GstH264Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// DRM device path
// Default: NULL
func (e *GstVaH264Dec) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// VA-API based H.264 video encoder
type GstVaH264Enc struct {
	*GstVaBaseEnc
}

func NewVaH264Enc() (*GstVaH264Enc, error) {
	e, err := gst.NewElement("vah264enc")
	if err != nil {
		return nil, err
	}
	return &GstVaH264Enc{GstVaBaseEnc: &GstVaBaseEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

func NewVaH264EncWithName(name string) (*GstVaH264Enc, error) {
	e, err := gst.NewElementWithName("vah264enc", name)
	if err != nil {
		return nil, err
	}
	return &GstVaH264Enc{GstVaBaseEnc: &GstVaBaseEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}}, nil
}

// The target usage to control and balance the encoding speed/quality
// Default: 4, Min: 1, Max: 7
func (e *GstVaH264Enc) SetTargetUsage(value uint) error {
	return e.Element.SetProperty("target-usage", value)
}

func (e *GstVaH264Enc) GetTargetUsage() (uint, error) {
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

// Insert AU (Access Unit) delimeter for each frame
// Default: false
func (e *GstVaH264Enc) SetAud(value bool) error {
	return e.Element.SetProperty("aud", value)
}

func (e *GstVaH264Enc) GetAud() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("aud")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The desired bitrate expressed in kbps (0: auto-calculate)
// Default: 0, Min: 0, Max: 2048000
func (e *GstVaH264Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstVaH264Enc) GetBitrate() (uint, error) {
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

// Maximum quantizer value for each frame
// Default: 51, Min: 0, Max: 51
func (e *GstVaH264Enc) SetMaxQp(value uint) error {
	return e.Element.SetProperty("max-qp", value)
}

func (e *GstVaH264Enc) GetMaxQp() (uint, error) {
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

// Macroblock level Bitrate Control. It is not compatible with CQP
// Default: disabled (0)
func (e *GstVaH264Enc) SetMbbrc(value interface{}) error {
	return e.Element.SetProperty("mbbrc", value)
}

func (e *GstVaH264Enc) GetMbbrc() (interface{}, error) {
	return e.Element.GetProperty("mbbrc")
}

// The percentage for 'target bitrate'/'maximum bitrate' (Only in VBR)
// Default: 66, Min: 50, Max: 100
func (e *GstVaH264Enc) SetTargetPercentage(value uint) error {
	return e.Element.SetProperty("target-percentage", value)
}

func (e *GstVaH264Enc) GetTargetPercentage() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-percentage")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable CABAC entropy coding mode
// Default: true
func (e *GstVaH264Enc) SetCabac(value bool) error {
	return e.Element.SetProperty("cabac", value)
}

func (e *GstVaH264Enc) GetCabac() (bool, error) {
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

// Insert CEA-708 Closed Captions
// Default: true
func (e *GstVaH264Enc) SetCcInsert(value bool) error {
	return e.Element.SetProperty("cc-insert", value)
}

func (e *GstVaH264Enc) GetCcInsert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cc-insert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The desired max CPB size in Kb (0: auto-calculate)
// Default: 0, Min: 0, Max: 2048000
func (e *GstVaH264Enc) SetCpbSize(value uint) error {
	return e.Element.SetProperty("cpb-size", value)
}

func (e *GstVaH264Enc) GetCpbSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cpb-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Force the number of I frames insertion within one GOP, not including the first IDR frame
// Default: 0, Min: 0, Max: 1023
func (e *GstVaH264Enc) SetIFrames(value uint) error {
	return e.Element.SetProperty("i-frames", value)
}

func (e *GstVaH264Enc) GetIFrames() (uint, error) {
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

// The quantizer value for B frame. This is available only in CQP mode
// Default: 26, Min: 0, Max: 51
func (e *GstVaH264Enc) SetQpb(value uint) error {
	return e.Element.SetProperty("qpb", value)
}

func (e *GstVaH264Enc) GetQpb() (uint, error) {
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

// The maximal distance between two keyframes. It decides the size of GOP (0: auto-calculate)
// Default: 0, Min: 0, Max: 1024
func (e *GstVaH264Enc) SetKeyIntMax(value uint) error {
	return e.Element.SetProperty("key-int-max", value)
}

func (e *GstVaH264Enc) GetKeyIntMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("key-int-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum quantizer value for each frame
// Default: 1, Min: 0, Max: 51
func (e *GstVaH264Enc) SetMinQp(value uint) error {
	return e.Element.SetProperty("min-qp", value)
}

func (e *GstVaH264Enc) GetMinQp() (uint, error) {
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

// The quantizer value for I frame. In CQP mode, it specifies the QP of I frame, in other mode, it specifies the init QP of all frames
// Default: 26, Min: 0, Max: 51
func (e *GstVaH264Enc) SetQpi(value uint) error {
	return e.Element.SetProperty("qpi", value)
}

func (e *GstVaH264Enc) GetQpi() (uint, error) {
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

// Number of reference frames, including both the forward and the backward
// Default: 3, Min: 0, Max: 16
func (e *GstVaH264Enc) SetRefFrames(value uint) error {
	return e.Element.SetProperty("ref-frames", value)
}

func (e *GstVaH264Enc) GetRefFrames() (uint, error) {
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

// Enable the trellis quantization method
// Default: false
func (e *GstVaH264Enc) SetTrellis(value bool) error {
	return e.Element.SetProperty("trellis", value)
}

func (e *GstVaH264Enc) GetTrellis() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of B frames between I and P reference frames
// Default: 0, Min: 0, Max: 31
func (e *GstVaH264Enc) SetBFrames(value uint) error {
	return e.Element.SetProperty("b-frames", value)
}

func (e *GstVaH264Enc) GetBFrames() (uint, error) {
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

// Enable the b-pyramid reference structure in the GOP
// Default: false
func (e *GstVaH264Enc) SetBPyramid(value bool) error {
	return e.Element.SetProperty("b-pyramid", value)
}

func (e *GstVaH264Enc) GetBPyramid() (bool, error) {
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

// Enable adaptive use of 8x8 transforms in I-frames
// Default: true
func (e *GstVaH264Enc) SetDct8x8(value bool) error {
	return e.Element.SetProperty("dct8x8", value)
}

func (e *GstVaH264Enc) GetDct8x8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dct8x8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of slices per frame
// Default: 1, Min: 1, Max: 200
func (e *GstVaH264Enc) SetNumSlices(value uint) error {
	return e.Element.SetProperty("num-slices", value)
}

func (e *GstVaH264Enc) GetNumSlices() (uint, error) {
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

// The quantizer value for P frame. This is available only in CQP mode
// Default: 26, Min: 0, Max: 51
func (e *GstVaH264Enc) SetQpp(value uint) error {
	return e.Element.SetProperty("qpp", value)
}

func (e *GstVaH264Enc) GetQpp() (uint, error) {
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

// VA-API based VP8 video decoder
type GstVaVp8Dec struct {
	*GstVp8Decoder
}

func NewVaVp8Dec() (*GstVaVp8Dec, error) {
	e, err := gst.NewElement("vavp8dec")
	if err != nil {
		return nil, err
	}
	return &GstVaVp8Dec{GstVp8Decoder: &GstVp8Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVaVp8DecWithName(name string) (*GstVaVp8Dec, error) {
	e, err := gst.NewElementWithName("vavp8dec", name)
	if err != nil {
		return nil, err
	}
	return &GstVaVp8Dec{GstVp8Decoder: &GstVp8Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// DRM device path
// Default: NULL
func (e *GstVaVp8Dec) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// VA-API based VP9 video decoder
type GstVaVp9Dec struct {
	*GstVp9Decoder
}

func NewVaVp9Dec() (*GstVaVp9Dec, error) {
	e, err := gst.NewElement("vavp9dec")
	if err != nil {
		return nil, err
	}
	return &GstVaVp9Dec{GstVp9Decoder: &GstVp9Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVaVp9DecWithName(name string) (*GstVaVp9Dec, error) {
	e, err := gst.NewElementWithName("vavp9dec", name)
	if err != nil {
		return nil, err
	}
	return &GstVaVp9Dec{GstVp9Decoder: &GstVp9Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// DRM device path
// Default: NULL
func (e *GstVaVp9Dec) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// VA-API based Mpeg2 video decoder
type GstVaMpeg2Dec struct {
	*GstMpeg2Decoder
}

func NewVaMpeg2Dec() (*GstVaMpeg2Dec, error) {
	e, err := gst.NewElement("vampeg2dec")
	if err != nil {
		return nil, err
	}
	return &GstVaMpeg2Dec{GstMpeg2Decoder: &GstMpeg2Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVaMpeg2DecWithName(name string) (*GstVaMpeg2Dec, error) {
	e, err := gst.NewElementWithName("vampeg2dec", name)
	if err != nil {
		return nil, err
	}
	return &GstVaMpeg2Dec{GstMpeg2Decoder: &GstMpeg2Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// DRM device path
// Default: NULL
func (e *GstVaMpeg2Dec) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// VA-API based video postprocessor
type GstVaPostProc struct {
	*GstVaBaseTransform
}

func NewVaPostProc() (*GstVaPostProc, error) {
	e, err := gst.NewElement("vapostproc")
	if err != nil {
		return nil, err
	}
	return &GstVaPostProc{GstVaBaseTransform: &GstVaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVaPostProcWithName(name string) (*GstVaPostProc, error) {
	e, err := gst.NewElementWithName("vapostproc", name)
	if err != nil {
		return nil, err
	}
	return &GstVaPostProc{GstVaBaseTransform: &GstVaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Add black borders if necessary to keep the display aspect ratio
// Default: false
func (e *GstVaPostProc) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

func (e *GstVaPostProc) GetAddBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Forces passing buffers through the postprocessor
// Default: false
func (e *GstVaPostProc) SetDisablePassthrough(value bool) error {
	return e.Element.SetProperty("disable-passthrough", value)
}

func (e *GstVaPostProc) GetDisablePassthrough() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("disable-passthrough")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Scale method to use
// Default: default (0)
func (e *GstVaPostProc) SetScaleMethod(value GstVaScaleMethod) error {
	e.Element.SetArg("scale-method", string(value))
	return nil
}

func (e *GstVaPostProc) GetScaleMethod() (GstVaScaleMethod, error) {
	var value GstVaScaleMethod
	var ok bool
	v, err := e.Element.GetProperty("scale-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVaScaleMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVaScaleMethod")
	}
	return value, nil
}

// VA-API based AV1 video decoder
type GstVaAV1Dec struct {
	*GstAV1Decoder
}

func NewVaAV1Dec() (*GstVaAV1Dec, error) {
	e, err := gst.NewElement("vaav1dec")
	if err != nil {
		return nil, err
	}
	return &GstVaAV1Dec{GstAV1Decoder: &GstAV1Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVaAV1DecWithName(name string) (*GstVaAV1Dec, error) {
	e, err := gst.NewElementWithName("vaav1dec", name)
	if err != nil {
		return nil, err
	}
	return &GstVaAV1Dec{GstAV1Decoder: &GstAV1Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// DRM device path
// Default: NULL
func (e *GstVaAV1Dec) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// VA-API based video compositor
type GstVaCompositor struct {
	*GstVideoAggregator
}

func NewVaCompositor() (*GstVaCompositor, error) {
	e, err := gst.NewElement("vacompositor")
	if err != nil {
		return nil, err
	}
	return &GstVaCompositor{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewVaCompositorWithName(name string) (*GstVaCompositor, error) {
	e, err := gst.NewElementWithName("vacompositor", name)
	if err != nil {
		return nil, err
	}
	return &GstVaCompositor{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// DRM device path
// Default: NULL
func (e *GstVaCompositor) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Scale method to use
// Default: default (0)
func (e *GstVaCompositor) SetScaleMethod(value GstVaScaleMethod) error {
	e.Element.SetArg("scale-method", string(value))
	return nil
}

func (e *GstVaCompositor) GetScaleMethod() (GstVaScaleMethod, error) {
	var value GstVaScaleMethod
	var ok bool
	v, err := e.Element.GetProperty("scale-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVaScaleMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVaScaleMethod")
	}
	return value, nil
}

// VA-API based deinterlacer
type GstVaDeinterlace struct {
	*GstVaBaseTransform
}

func NewVaDeinterlace() (*GstVaDeinterlace, error) {
	e, err := gst.NewElement("vadeinterlace")
	if err != nil {
		return nil, err
	}
	return &GstVaDeinterlace{GstVaBaseTransform: &GstVaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVaDeinterlaceWithName(name string) (*GstVaDeinterlace, error) {
	e, err := gst.NewElementWithName("vadeinterlace", name)
	if err != nil {
		return nil, err
	}
	return &GstVaDeinterlace{GstVaBaseTransform: &GstVaBaseTransform{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Deinterlace Method
// Default: bob (1)
func (e *GstVaDeinterlace) SetMethod(value GstVaDeinterlaceMethods) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstVaDeinterlace) GetMethod() (GstVaDeinterlaceMethods, error) {
	var value GstVaDeinterlaceMethods
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVaDeinterlaceMethods)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVaDeinterlaceMethods")
	}
	return value, nil
}

// VA-API based H.265 video decoder
type GstVaH265Dec struct {
	*GstH265Decoder
}

func NewVaH265Dec() (*GstVaH265Dec, error) {
	e, err := gst.NewElement("vah265dec")
	if err != nil {
		return nil, err
	}
	return &GstVaH265Dec{GstH265Decoder: &GstH265Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVaH265DecWithName(name string) (*GstVaH265Dec, error) {
	e, err := gst.NewElementWithName("vah265dec", name)
	if err != nil {
		return nil, err
	}
	return &GstVaH265Dec{GstH265Decoder: &GstH265Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// DRM device path
// Default: NULL
func (e *GstVaH265Dec) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// VA-API based JPEG image decoder
type GstVaJpegDec struct {
	*GstJpegDecoder
}

func NewVaJpegDec() (*GstVaJpegDec, error) {
	e, err := gst.NewElement("vajpegdec")
	if err != nil {
		return nil, err
	}
	return &GstVaJpegDec{GstJpegDecoder: &GstJpegDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

func NewVaJpegDecWithName(name string) (*GstVaJpegDec, error) {
	e, err := gst.NewElementWithName("vajpegdec", name)
	if err != nil {
		return nil, err
	}
	return &GstVaJpegDec{GstJpegDecoder: &GstJpegDecoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}, nil
}

// DRM device path
// Default: NULL
func (e *GstVaJpegDec) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstJpegDecoder struct {
	*GstVideoDecoder
}

type GstVaBaseEnc struct {
	*GstVideoEncoder
}

// DRM device path
// Default: NULL
func (e *GstVaBaseEnc) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstVaBaseTransform struct {
	*base.GstBaseTransform
}

// DRM device path
// Default: NULL
func (e *GstVaBaseTransform) GetDevicePath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstVaDeinterlaceMethods string

const (
	GstVaDeinterlaceMethodsBob         GstVaDeinterlaceMethods = "bob"         // Bob: Interpolating missing lines by using the adjacent lines.
	GstVaDeinterlaceMethodsAdaptive    GstVaDeinterlaceMethods = "adaptive"    // Adaptive: Interpolating missing lines by using spatial/temporal references.
	GstVaDeinterlaceMethodsCompensated GstVaDeinterlaceMethods = "compensated" // Compensation: Recreating missing lines by using motion vector.
)

type GstVaScaleMethod string

const (
	GstVaScaleMethodDefault GstVaScaleMethod = "default" // Default scaling method
	GstVaScaleMethodFast    GstVaScaleMethod = "fast"    // Fast scaling method
	GstVaScaleMethodHq      GstVaScaleMethod = "hq"      // High quality scaling method
)
