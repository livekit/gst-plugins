// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// decode raw theora streams to raw YUV video
type GstTheoraDec struct {
	*GstVideoDecoder
}

func NewTheoraDec() (*GstTheoraDec, error) {
	e, err := gst.NewElement("theoradec")
	if err != nil {
		return nil, err
	}
	return &GstTheoraDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewTheoraDecWithName(name string) (*GstTheoraDec, error) {
	e, err := gst.NewElementWithName("theoradec", name)
	if err != nil {
		return nil, err
	}
	return &GstTheoraDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// Sets the bitstream breakdown visualization mode. Values influence the width of the bit usage bars to show
// Default: 0, Min: 0, Max: 255
func (e *GstTheoraDec) SetVisualizeBitUsage(value int) error {
	return e.Element.SetProperty("visualize-bit-usage", value)
}

func (e *GstTheoraDec) GetVisualizeBitUsage() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("visualize-bit-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Show macroblock mode selection overlaid on image. Value gives a mask for macroblock (MB) modes to show
// Default: 0, Min: 0, Max: 65535
func (e *GstTheoraDec) SetVisualizeMacroblockModes(value int) error {
	return e.Element.SetProperty("visualize-macroblock-modes", value)
}

func (e *GstTheoraDec) GetVisualizeMacroblockModes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("visualize-macroblock-modes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Show motion vector selection overlaid on image. Value gives a mask for motion vector (MV) modes to show
// Default: 0, Min: 0, Max: 65535
func (e *GstTheoraDec) SetVisualizeMotionVectors(value int) error {
	return e.Element.SetProperty("visualize-motion-vectors", value)
}

func (e *GstTheoraDec) GetVisualizeMotionVectors() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("visualize-motion-vectors")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Show adaptive quantization mode selection overlaid on image. Value gives a mask for quantization (QI) modes to show
// Default: 0, Min: 0, Max: 65535
func (e *GstTheoraDec) SetVisualizeQuantizationModes(value int) error {
	return e.Element.SetProperty("visualize-quantization-modes", value)
}

func (e *GstTheoraDec) GetVisualizeQuantizationModes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("visualize-quantization-modes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// encode raw YUV video to a theora stream
type GstTheoraEnc struct {
	*GstVideoEncoder
}

func NewTheoraEnc() (*GstTheoraEnc, error) {
	e, err := gst.NewElement("theoraenc")
	if err != nil {
		return nil, err
	}
	return &GstTheoraEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewTheoraEncWithName(name string) (*GstTheoraEnc, error) {
	e, err := gst.NewElementWithName("theoraenc", name)
	if err != nil {
		return nil, err
	}
	return &GstTheoraEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// Keyframe frequency
// Default: 64, Min: 1, Max: 32768
func (e *GstTheoraEnc) SetKeyframeFreq(value int) error {
	return e.Element.SetProperty("keyframe-freq", value)
}

func (e *GstTheoraEnc) GetKeyframeFreq() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyframe-freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Multipass cache file
// Default: NULL
func (e *GstTheoraEnc) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

func (e *GstTheoraEnc) GetMultipassCacheFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("multipass-cache-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Single pass or first/second pass
// Default: single-pass (0)
func (e *GstTheoraEnc) SetMultipassMode(value GstTheoraEncMultipassMode) error {
	e.Element.SetArg("multipass-mode", string(value))
	return nil
}

func (e *GstTheoraEnc) GetMultipassMode() (GstTheoraEncMultipassMode, error) {
	var value GstTheoraEncMultipassMode
	var ok bool
	v, err := e.Element.GetProperty("multipass-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTheoraEncMultipassMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTheoraEncMultipassMode")
	}
	return value, nil
}

// Disables non-VP3 compatible features
// Default: false
func (e *GstTheoraEnc) SetVp3Compatible(value bool) error {
	return e.Element.SetProperty("vp3-compatible", value)
}

func (e *GstTheoraEnc) GetVp3Compatible() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vp3-compatible")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable capping of bit reservoir overflows
// Default: true
func (e *GstTheoraEnc) SetCapOverflow(value bool) error {
	return e.Element.SetProperty("cap-overflow", value)
}

func (e *GstTheoraEnc) GetCapOverflow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cap-overflow")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable capping of bit reservoir underflows
// Default: false
func (e *GstTheoraEnc) SetCapUnderflow(value bool) error {
	return e.Element.SetProperty("cap-underflow", value)
}

func (e *GstTheoraEnc) GetCapUnderflow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cap-underflow")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Force keyframe every N frames
// Default: 64, Min: 1, Max: 32768
func (e *GstTheoraEnc) SetKeyframeForce(value int) error {
	return e.Element.SetProperty("keyframe-force", value)
}

func (e *GstTheoraEnc) GetKeyframeForce() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyframe-force")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Video quality
// Default: 48, Min: 0, Max: 63
func (e *GstTheoraEnc) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstTheoraEnc) GetQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sets the size of the rate control buffer, in units of frames.  The default value of 0 instructs the encoder to automatically select an appropriate value
// Default: 0, Min: 0, Max: 1000
func (e *GstTheoraEnc) SetRateBuffer(value int) error {
	return e.Element.SetProperty("rate-buffer", value)
}

func (e *GstTheoraEnc) GetRateBuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rate-buffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Controls the amount of motion vector searching done while encoding
// Default: 1, Min: 0, Max: 3
func (e *GstTheoraEnc) SetSpeedLevel(value int) error {
	return e.Element.SetProperty("speed-level", value)
}

func (e *GstTheoraEnc) GetSpeedLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("speed-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Compressed video bitrate (kbps)
// Default: 0, Min: 0, Max: 16777215
func (e *GstTheoraEnc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstTheoraEnc) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Allow or disallow frame dropping
// Default: true
func (e *GstTheoraEnc) SetDropFrames(value bool) error {
	return e.Element.SetProperty("drop-frames", value)
}

func (e *GstTheoraEnc) GetDropFrames() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Automatic keyframe detection
// Default: true
func (e *GstTheoraEnc) SetKeyframeAuto(value bool) error {
	return e.Element.SetProperty("keyframe-auto", value)
}

func (e *GstTheoraEnc) GetKeyframeAuto() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keyframe-auto")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// parse raw theora streams
type GstTheoraParse struct {
	*gst.Element
}

func NewTheoraParse() (*GstTheoraParse, error) {
	e, err := gst.NewElement("theoraparse")
	if err != nil {
		return nil, err
	}
	return &GstTheoraParse{Element: e}, nil
}

func NewTheoraParseWithName(name string) (*GstTheoraParse, error) {
	e, err := gst.NewElementWithName("theoraparse", name)
	if err != nil {
		return nil, err
	}
	return &GstTheoraParse{Element: e}, nil
}

type GstTheoraEncMultipassMode string

const (
	GstTheoraEncMultipassModeSinglePass GstTheoraEncMultipassMode = "single-pass" // Single pass
	GstTheoraEncMultipassModeFirstPass  GstTheoraEncMultipassMode = "first-pass"  // First pass
	GstTheoraEncMultipassModeSecondPass GstTheoraEncMultipassMode = "second-pass" // Second pass
)
