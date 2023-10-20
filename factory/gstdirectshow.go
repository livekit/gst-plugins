// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// DirectShow Microsoft ISO MPEG-4 version 1 Decoder Wrapper
type DshowvdecMsmpeg41 struct {
	*gst.Element
}

func NewDshowvdecMsmpeg41() (*DshowvdecMsmpeg41, error) {
	e, err := gst.NewElement("dshowvdec_msmpeg41")
	if err != nil {
		return nil, err
	}
	return &DshowvdecMsmpeg41{Element: e}, nil
}

func NewDshowvdecMsmpeg41WithName(name string) (*DshowvdecMsmpeg41, error) {
	e, err := gst.NewElementWithName("dshowvdec_msmpeg41", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecMsmpeg41{Element: e}, nil
}

// DirectShow XVID Video Decoder Wrapper
type DshowvdecXvid struct {
	*gst.Element
}

func NewDshowvdecXvid() (*DshowvdecXvid, error) {
	e, err := gst.NewElement("dshowvdec_xvid")
	if err != nil {
		return nil, err
	}
	return &DshowvdecXvid{Element: e}, nil
}

func NewDshowvdecXvidWithName(name string) (*DshowvdecXvid, error) {
	e, err := gst.NewElementWithName("dshowvdec_xvid", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecXvid{Element: e}, nil
}

// DirectShow DIVX 4.0 Video Decoder Wrapper
type DshowvdecDivx4 struct {
	*gst.Element
}

func NewDshowvdecDivx4() (*DshowvdecDivx4, error) {
	e, err := gst.NewElement("dshowvdec_divx4")
	if err != nil {
		return nil, err
	}
	return &DshowvdecDivx4{Element: e}, nil
}

func NewDshowvdecDivx4WithName(name string) (*DshowvdecDivx4, error) {
	e, err := gst.NewElementWithName("dshowvdec_divx4", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecDivx4{Element: e}, nil
}

// DirectShow Windows Media Audio 8 Decoder Wrapper
type DshowadecWma2 struct {
	*gst.Element
}

func NewDshowadecWma2() (*DshowadecWma2, error) {
	e, err := gst.NewElement("dshowadec_wma2")
	if err != nil {
		return nil, err
	}
	return &DshowadecWma2{Element: e}, nil
}

func NewDshowadecWma2WithName(name string) (*DshowadecWma2, error) {
	e, err := gst.NewElementWithName("dshowadec_wma2", name)
	if err != nil {
		return nil, err
	}
	return &DshowadecWma2{Element: e}, nil
}

// DirectShow Windows Media Audio Voice v9 Decoder Wrapper
type DshowadecWms struct {
	*gst.Element
}

func NewDshowadecWms() (*DshowadecWms, error) {
	e, err := gst.NewElement("dshowadec_wms")
	if err != nil {
		return nil, err
	}
	return &DshowadecWms{Element: e}, nil
}

func NewDshowadecWmsWithName(name string) (*DshowadecWms, error) {
	e, err := gst.NewElementWithName("dshowadec_wms", name)
	if err != nil {
		return nil, err
	}
	return &DshowadecWms{Element: e}, nil
}

// DirectShow MPEG Layer 1,2 Audio Decoder Wrapper
type DshowadecMpeg12 struct {
	*gst.Element
}

func NewDshowadecMpeg12() (*DshowadecMpeg12, error) {
	e, err := gst.NewElement("dshowadec_mpeg_1_2")
	if err != nil {
		return nil, err
	}
	return &DshowadecMpeg12{Element: e}, nil
}

func NewDshowadecMpeg12WithName(name string) (*DshowadecMpeg12, error) {
	e, err := gst.NewElementWithName("dshowadec_mpeg_1_2", name)
	if err != nil {
		return nil, err
	}
	return &DshowadecMpeg12{Element: e}, nil
}

// DirectShow Windows Media Video 9 Advanced Decoder Wrapper
type DshowvdecWmva struct {
	*gst.Element
}

func NewDshowvdecWmva() (*DshowvdecWmva, error) {
	e, err := gst.NewElement("dshowvdec_wmva")
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmva{Element: e}, nil
}

func NewDshowvdecWmvaWithName(name string) (*DshowvdecWmva, error) {
	e, err := gst.NewElementWithName("dshowvdec_wmva", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmva{Element: e}, nil
}

// DirectShow Windows Media VC1 video Decoder Wrapper
type DshowvdecWvc1 struct {
	*gst.Element
}

func NewDshowvdecWvc1() (*DshowvdecWvc1, error) {
	e, err := gst.NewElement("dshowvdec_wvc1")
	if err != nil {
		return nil, err
	}
	return &DshowvdecWvc1{Element: e}, nil
}

func NewDshowvdecWvc1WithName(name string) (*DshowvdecWvc1, error) {
	e, err := gst.NewElementWithName("dshowvdec_wvc1", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecWvc1{Element: e}, nil
}

// Display data using a DirectShow video renderer
type GstDshowVideoSink struct {
	*GstVideoSink
}

func NewDshowVideoSink() (*GstDshowVideoSink, error) {
	e, err := gst.NewElement("dshowvideosink")
	if err != nil {
		return nil, err
	}
	return &GstDshowVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewDshowVideoSinkWithName(name string) (*GstDshowVideoSink, error) {
	e, err := gst.NewElementWithName("dshowvideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstDshowVideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: false
func (e *GstDshowVideoSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstDshowVideoSink) GetForceAspectRatio() (bool, error) {
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

// Use full-screen mode (not available when using XOverlay)
// Default: false
func (e *GstDshowVideoSink) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

func (e *GstDshowVideoSink) GetFullscreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fullscreen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Force usage of specific DirectShow renderer (EVR, VMR9 or VMR7)
// Default: NULL
func (e *GstDshowVideoSink) SetRenderer(value string) error {
	return e.Element.SetProperty("renderer", value)
}

func (e *GstDshowVideoSink) GetRenderer() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("renderer")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// DirectShow Microsoft ISO MPEG-4 version 1.1 Decoder Wrapper
type DshowvdecMsmpeg4 struct {
	*gst.Element
}

func NewDshowvdecMsmpeg4() (*DshowvdecMsmpeg4, error) {
	e, err := gst.NewElement("dshowvdec_msmpeg4")
	if err != nil {
		return nil, err
	}
	return &DshowvdecMsmpeg4{Element: e}, nil
}

func NewDshowvdecMsmpeg4WithName(name string) (*DshowvdecMsmpeg4, error) {
	e, err := gst.NewElementWithName("dshowvdec_msmpeg4", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecMsmpeg4{Element: e}, nil
}

// Receive data from a directshow audio capture graph
type GstDshowAudioSrc struct {
	*GstAudioSrc
}

func NewDshowAudioSrc() (*GstDshowAudioSrc, error) {
	e, err := gst.NewElement("dshowaudiosrc")
	if err != nil {
		return nil, err
	}
	return &GstDshowAudioSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

func NewDshowAudioSrcWithName(name string) (*GstDshowAudioSrc, error) {
	e, err := gst.NewElementWithName("dshowaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDshowAudioSrc{GstAudioSrc: &GstAudioSrc{GstAudioBaseSrc: &GstAudioBaseSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}}}, nil
}

// Directshow device reference (classID/name)
// Default: NULL
func (e *GstDshowAudioSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstDshowAudioSrc) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Index of the enumerated audio device
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDshowAudioSrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

func (e *GstDshowAudioSrc) GetDeviceIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Human-readable name of the sound device
// Default: NULL
func (e *GstDshowAudioSrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

func (e *GstDshowAudioSrc) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// DirectShow Cinepack Decoder Wrapper
type DshowvdecCinepak struct {
	*gst.Element
}

func NewDshowvdecCinepak() (*DshowvdecCinepak, error) {
	e, err := gst.NewElement("dshowvdec_cinepak")
	if err != nil {
		return nil, err
	}
	return &DshowvdecCinepak{Element: e}, nil
}

func NewDshowvdecCinepakWithName(name string) (*DshowvdecCinepak, error) {
	e, err := gst.NewElementWithName("dshowvdec_cinepak", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecCinepak{Element: e}, nil
}

// DirectShow DIVX 3.0 Video Decoder Wrapper
type DshowvdecDivx3 struct {
	*gst.Element
}

func NewDshowvdecDivx3() (*DshowvdecDivx3, error) {
	e, err := gst.NewElement("dshowvdec_divx3")
	if err != nil {
		return nil, err
	}
	return &DshowvdecDivx3{Element: e}, nil
}

func NewDshowvdecDivx3WithName(name string) (*DshowvdecDivx3, error) {
	e, err := gst.NewElementWithName("dshowvdec_divx3", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecDivx3{Element: e}, nil
}

// DirectShow Windows Media Video 7 Decoder Wrapper
type DshowvdecWmv1 struct {
	*gst.Element
}

func NewDshowvdecWmv1() (*DshowvdecWmv1, error) {
	e, err := gst.NewElement("dshowvdec_wmv1")
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmv1{Element: e}, nil
}

func NewDshowvdecWmv1WithName(name string) (*DshowvdecWmv1, error) {
	e, err := gst.NewElementWithName("dshowvdec_wmv1", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmv1{Element: e}, nil
}

// DirectShow Windows Media Video 9 Decoder Wrapper
type DshowvdecWmv3 struct {
	*gst.Element
}

func NewDshowvdecWmv3() (*DshowvdecWmv3, error) {
	e, err := gst.NewElement("dshowvdec_wmv3")
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmv3{Element: e}, nil
}

func NewDshowvdecWmv3WithName(name string) (*DshowvdecWmv3, error) {
	e, err := gst.NewElementWithName("dshowvdec_wmv3", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmv3{Element: e}, nil
}

// DirectShow Windows Media Audio 9 Professional Decoder Wrapper
type DshowadecWma3 struct {
	*gst.Element
}

func NewDshowadecWma3() (*DshowadecWma3, error) {
	e, err := gst.NewElement("dshowadec_wma3")
	if err != nil {
		return nil, err
	}
	return &DshowadecWma3{Element: e}, nil
}

func NewDshowadecWma3WithName(name string) (*DshowadecWma3, error) {
	e, err := gst.NewElementWithName("dshowadec_wma3", name)
	if err != nil {
		return nil, err
	}
	return &DshowadecWma3{Element: e}, nil
}

// DirectShow DIVX 5.0 Video Decoder Wrapper
type DshowvdecDivx5 struct {
	*gst.Element
}

func NewDshowvdecDivx5() (*DshowvdecDivx5, error) {
	e, err := gst.NewElement("dshowvdec_divx5")
	if err != nil {
		return nil, err
	}
	return &DshowvdecDivx5{Element: e}, nil
}

func NewDshowvdecDivx5WithName(name string) (*DshowvdecDivx5, error) {
	e, err := gst.NewElementWithName("dshowvdec_divx5", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecDivx5{Element: e}, nil
}

// DirectShow Microsoft ISO MPEG-4 version 3 Decoder Wrapper
type DshowvdecMsmpeg43 struct {
	*gst.Element
}

func NewDshowvdecMsmpeg43() (*DshowvdecMsmpeg43, error) {
	e, err := gst.NewElement("dshowvdec_msmpeg43")
	if err != nil {
		return nil, err
	}
	return &DshowvdecMsmpeg43{Element: e}, nil
}

func NewDshowvdecMsmpeg43WithName(name string) (*DshowvdecMsmpeg43, error) {
	e, err := gst.NewElementWithName("dshowvdec_msmpeg43", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecMsmpeg43{Element: e}, nil
}

// DirectShow Windows Media Video 8 Decoder Wrapper
type DshowvdecWmv2 struct {
	*gst.Element
}

func NewDshowvdecWmv2() (*DshowvdecWmv2, error) {
	e, err := gst.NewElement("dshowvdec_wmv2")
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmv2{Element: e}, nil
}

func NewDshowvdecWmv2WithName(name string) (*DshowvdecWmv2, error) {
	e, err := gst.NewElementWithName("dshowvdec_wmv2", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmv2{Element: e}, nil
}

// Receive data from a directshow video capture graph
type GstDshowVideoSrc struct {
	*base.GstPushSrc
}

func NewDshowVideoSrc() (*GstDshowVideoSrc, error) {
	e, err := gst.NewElement("dshowvideosrc")
	if err != nil {
		return nil, err
	}
	return &GstDshowVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewDshowVideoSrcWithName(name string) (*GstDshowVideoSrc, error) {
	e, err := gst.NewElementWithName("dshowvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDshowVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Directshow device path (@..classID/name)
// Default: NULL
func (e *GstDshowVideoSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstDshowVideoSrc) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Index of the enumerated video device
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDshowVideoSrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

func (e *GstDshowVideoSrc) GetDeviceIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Human-readable name of the video device
// Default: NULL
func (e *GstDshowVideoSrc) SetDeviceName(value string) error {
	return e.Element.SetProperty("device-name", value)
}

func (e *GstDshowVideoSrc) GetDeviceName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// DirectShow MPEG Layer 3 Audio Decoder Wrapper
type DshowadecMp3 struct {
	*gst.Element
}

func NewDshowadecMp3() (*DshowadecMp3, error) {
	e, err := gst.NewElement("dshowadec_mp3")
	if err != nil {
		return nil, err
	}
	return &DshowadecMp3{Element: e}, nil
}

func NewDshowadecMp3WithName(name string) (*DshowadecMp3, error) {
	e, err := gst.NewElementWithName("dshowadec_mp3", name)
	if err != nil {
		return nil, err
	}
	return &DshowadecMp3{Element: e}, nil
}

// DirectShow Windows Media Audio 9 Lossless Decoder Wrapper
type DshowadecWma4 struct {
	*gst.Element
}

func NewDshowadecWma4() (*DshowadecWma4, error) {
	e, err := gst.NewElement("dshowadec_wma4")
	if err != nil {
		return nil, err
	}
	return &DshowadecWma4{Element: e}, nil
}

func NewDshowadecWma4WithName(name string) (*DshowadecWma4, error) {
	e, err := gst.NewElementWithName("dshowadec_wma4", name)
	if err != nil {
		return nil, err
	}
	return &DshowadecWma4{Element: e}, nil
}

// DirectShow MPEG-1 Video Decoder Wrapper
type DshowvdecMpeg1 struct {
	*gst.Element
}

func NewDshowvdecMpeg1() (*DshowvdecMpeg1, error) {
	e, err := gst.NewElement("dshowvdec_mpeg1")
	if err != nil {
		return nil, err
	}
	return &DshowvdecMpeg1{Element: e}, nil
}

func NewDshowvdecMpeg1WithName(name string) (*DshowvdecMpeg1, error) {
	e, err := gst.NewElementWithName("dshowvdec_mpeg1", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecMpeg1{Element: e}, nil
}

// DirectShow MPEG-4 Video Decoder Wrapper
type DshowvdecMpeg4 struct {
	*gst.Element
}

func NewDshowvdecMpeg4() (*DshowvdecMpeg4, error) {
	e, err := gst.NewElement("dshowvdec_mpeg4")
	if err != nil {
		return nil, err
	}
	return &DshowvdecMpeg4{Element: e}, nil
}

func NewDshowvdecMpeg4WithName(name string) (*DshowvdecMpeg4, error) {
	e, err := gst.NewElementWithName("dshowvdec_mpeg4", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecMpeg4{Element: e}, nil
}

// DirectShow Microsoft ISO MPEG-4 version 2 Decoder Wrapper
type DshowvdecMsmpeg42 struct {
	*gst.Element
}

func NewDshowvdecMsmpeg42() (*DshowvdecMsmpeg42, error) {
	e, err := gst.NewElement("dshowvdec_msmpeg42")
	if err != nil {
		return nil, err
	}
	return &DshowvdecMsmpeg42{Element: e}, nil
}

func NewDshowvdecMsmpeg42WithName(name string) (*DshowvdecMsmpeg42, error) {
	e, err := gst.NewElementWithName("dshowvdec_msmpeg42", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecMsmpeg42{Element: e}, nil
}

// DirectShow Windows Media Audio 7 Decoder Wrapper
type DshowadecWma1 struct {
	*gst.Element
}

func NewDshowadecWma1() (*DshowadecWma1, error) {
	e, err := gst.NewElement("dshowadec_wma1")
	if err != nil {
		return nil, err
	}
	return &DshowadecWma1{Element: e}, nil
}

func NewDshowadecWma1WithName(name string) (*DshowadecWma1, error) {
	e, err := gst.NewElementWithName("dshowadec_wma1", name)
	if err != nil {
		return nil, err
	}
	return &DshowadecWma1{Element: e}, nil
}

// DirectShow Windows Media Video 9 Image Decoder Wrapper
type DshowvdecWmvp struct {
	*gst.Element
}

func NewDshowvdecWmvp() (*DshowvdecWmvp, error) {
	e, err := gst.NewElement("dshowvdec_wmvp")
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmvp{Element: e}, nil
}

func NewDshowvdecWmvpWithName(name string) (*DshowvdecWmvp, error) {
	e, err := gst.NewElementWithName("dshowvdec_wmvp", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecWmvp{Element: e}, nil
}
