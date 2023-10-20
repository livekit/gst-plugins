// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// A Direct3D11 compositor
type GstD3D11Compositor struct {
	*GstVideoAggregator
}

func NewD3D11Compositor() (*GstD3D11Compositor, error) {
	e, err := gst.NewElement("d3d11compositor")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Compositor{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewD3D11CompositorWithName(name string) (*GstD3D11Compositor, error) {
	e, err := gst.NewElementWithName("d3d11compositor", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Compositor{GstVideoAggregator: &GstVideoAggregator{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// Adapter index for creating device (-1 for default)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstD3D11Compositor) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

func (e *GstD3D11Compositor) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Background type
// Default: checker (0)
func (e *GstD3D11Compositor) SetBackground(value GstD3D11CompositorBackground) error {
	e.Element.SetArg("background", string(value))
	return nil
}

func (e *GstD3D11Compositor) GetBackground() (GstD3D11CompositorBackground, error) {
	var value GstD3D11CompositorBackground
	var ok bool
	v, err := e.Element.GetProperty("background")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11CompositorBackground)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11CompositorBackground")
	}
	return value, nil
}

// Avoid timing out waiting for inactive pads
// Default: false
func (e *GstD3D11Compositor) SetIgnoreInactivePads(value bool) error {
	return e.Element.SetProperty("ignore-inactive-pads", value)
}

func (e *GstD3D11Compositor) GetIgnoreInactivePads() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-inactive-pads")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Performs resizing, colorspace conversion, cropping, and flipping/rotating using Direct3D11
type GstD3D11Convert struct {
	*GstD3D11BaseConvert
}

func NewD3D11Convert() (*GstD3D11Convert, error) {
	e, err := gst.NewElement("d3d11convert")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Convert{GstD3D11BaseConvert: &GstD3D11BaseConvert{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewD3D11ConvertWithName(name string) (*GstD3D11Convert, error) {
	e, err := gst.NewElementWithName("d3d11convert", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Convert{GstD3D11BaseConvert: &GstD3D11BaseConvert{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Add black borders if necessary to keep the display aspect ratio
// Default: true
func (e *GstD3D11Convert) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

func (e *GstD3D11Convert) GetAddBorders() (bool, error) {
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

// Border color to use in ARGB64 format
// Default: 18446462598732840960, Min: 0, Max: 18446744073709551615
func (e *GstD3D11Convert) SetBorderColor(value uint64) error {
	return e.Element.SetProperty("border-color", value)
}

func (e *GstD3D11Convert) GetBorderColor() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("border-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Output alpha mode to be applied
// Default: unspecified (0)
func (e *GstD3D11Convert) SetDestAlphaMode(value GstD3D11AlphaMode) error {
	e.Element.SetArg("dest-alpha-mode", string(value))
	return nil
}

func (e *GstD3D11Convert) GetDestAlphaMode() (GstD3D11AlphaMode, error) {
	var value GstD3D11AlphaMode
	var ok bool
	v, err := e.Element.GetProperty("dest-alpha-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11AlphaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11AlphaMode")
	}
	return value, nil
}

// Gamma conversion mode
// Default: none (0)
func (e *GstD3D11Convert) SetGammaMode(value interface{}) error {
	return e.Element.SetProperty("gamma-mode", value)
}

func (e *GstD3D11Convert) GetGammaMode() (interface{}, error) {
	return e.Element.GetProperty("gamma-mode")
}

// Primaries conversion mode
// Default: none (0)
func (e *GstD3D11Convert) SetPrimariesMode(value interface{}) error {
	return e.Element.SetProperty("primaries-mode", value)
}

func (e *GstD3D11Convert) GetPrimariesMode() (interface{}, error) {
	return e.Element.GetProperty("primaries-mode")
}

// Applied input alpha mode
// Default: unspecified (0)
func (e *GstD3D11Convert) SetSrcAlphaMode(value GstD3D11AlphaMode) error {
	e.Element.SetArg("src-alpha-mode", string(value))
	return nil
}

func (e *GstD3D11Convert) GetSrcAlphaMode() (GstD3D11AlphaMode, error) {
	var value GstD3D11AlphaMode
	var ok bool
	v, err := e.Element.GetProperty("src-alpha-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11AlphaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11AlphaMode")
	}
	return value, nil
}

// Downloads Direct3D11 texture memory into system memory
type GstD3D11Download struct {
	*GstD3D11BaseFilter
}

func NewD3D11Download() (*GstD3D11Download, error) {
	e, err := gst.NewElement("d3d11download")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Download{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewD3D11DownloadWithName(name string) (*GstD3D11Download, error) {
	e, err := gst.NewElementWithName("d3d11download", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Download{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Direct3D11/DXVA based H.265 video decoder
type GstD3D11H265Dec struct {
	*GstDxvaH265Decoder
}

func NewD3D11H265Dec() (*GstD3D11H265Dec, error) {
	e, err := gst.NewElement("d3d11h265dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D11H265Dec{GstDxvaH265Decoder: &GstDxvaH265Decoder{GstH265Decoder: &GstH265Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D11H265DecWithName(name string) (*GstD3D11H265Dec, error) {
	e, err := gst.NewElementWithName("d3d11h265dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11H265Dec{GstDxvaH265Decoder: &GstDxvaH265Decoder{GstH265Decoder: &GstH265Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D11H265Dec) GetAdapterLuid() (int64, error) {
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

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11H265Dec) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Vendor ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11H265Dec) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Sends Direct3D11 shared handle to peer d3d11ipcsrc elements
type GstD3D11IpcSink struct {
	*base.GstBaseSink
}

func NewD3D11IpcSink() (*GstD3D11IpcSink, error) {
	e, err := gst.NewElement("d3d11ipcsink")
	if err != nil {
		return nil, err
	}
	return &GstD3D11IpcSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewD3D11IpcSinkWithName(name string) (*GstD3D11IpcSink, error) {
	e, err := gst.NewElementWithName("d3d11ipcsink", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11IpcSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// DXGI adapter index (-1 for default)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstD3D11IpcSink) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

func (e *GstD3D11IpcSink) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minumum number of buffers
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11IpcSink) SetMinBufferSize(value uint) error {
	return e.Element.SetProperty("min-buffer-size", value)
}

func (e *GstD3D11IpcSink) GetMinBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The name of Win32 named pipe to communicate with clients. Validation of the pipe name is caller's responsibility
// Default: \\.\pipe\gst.d3d11.ipc
func (e *GstD3D11IpcSink) SetPipeName(value string) error {
	return e.Element.SetProperty("pipe-name", value)
}

func (e *GstD3D11IpcSink) GetPipeName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pipe-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Receives Direct3D11 shared handle from the d3d11ipcsink element
type GstD3D11IpcSrc struct {
	*base.GstBaseSrc
}

func NewD3D11IpcSrc() (*GstD3D11IpcSrc, error) {
	e, err := gst.NewElement("d3d11ipcsrc")
	if err != nil {
		return nil, err
	}
	return &GstD3D11IpcSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewD3D11IpcSrcWithName(name string) (*GstD3D11IpcSrc, error) {
	e, err := gst.NewElementWithName("d3d11ipcsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11IpcSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Memory I/O mode to use
// Default: copy (0)
func (e *GstD3D11IpcSrc) SetIoMode(value GstD3D11IpcIOMode) error {
	e.Element.SetArg("io-mode", string(value))
	return nil
}

func (e *GstD3D11IpcSrc) GetIoMode() (GstD3D11IpcIOMode, error) {
	var value GstD3D11IpcIOMode
	var ok bool
	v, err := e.Element.GetProperty("io-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11IpcIOMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11IpcIOMode")
	}
	return value, nil
}

// The name of Win32 named pipe to communicate with clients. Validation of the pipe name is caller's responsibility
// Default: \\.\pipe\gst.d3d11.ipc
func (e *GstD3D11IpcSrc) SetPipeName(value string) error {
	return e.Element.SetProperty("pipe-name", value)
}

func (e *GstD3D11IpcSrc) GetPipeName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pipe-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Maximum processing time for a buffer in nanoseconds
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstD3D11IpcSrc) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstD3D11IpcSrc) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Connection timeout in seconds (0 = never timeout)
// Default: 5, Min: 0, Max: 2147483647
func (e *GstD3D11IpcSrc) SetConnectionTimeout(value uint) error {
	return e.Element.SetProperty("connection-timeout", value)
}

func (e *GstD3D11IpcSrc) GetConnectionTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("connection-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Direct3D11/DXVA based AV1 video decoder
type GstD3D11AV1Dec struct {
	*GstDxvaAV1Decoder
}

func NewD3D11AV1Dec() (*GstD3D11AV1Dec, error) {
	e, err := gst.NewElement("d3d11av1dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D11AV1Dec{GstDxvaAV1Decoder: &GstDxvaAV1Decoder{GstAV1Decoder: &GstAV1Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D11AV1DecWithName(name string) (*GstD3D11AV1Dec, error) {
	e, err := gst.NewElementWithName("d3d11av1dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11AV1Dec{GstDxvaAV1Decoder: &GstDxvaAV1Decoder{GstAV1Decoder: &GstAV1Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D11AV1Dec) GetAdapterLuid() (int64, error) {
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

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11AV1Dec) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Vendor ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11AV1Dec) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Color conversion using Direct3D11
type GstD3D11ColorConvert struct {
	*GstD3D11BaseConvert
}

func NewD3D11ColorConvert() (*GstD3D11ColorConvert, error) {
	e, err := gst.NewElement("d3d11colorconvert")
	if err != nil {
		return nil, err
	}
	return &GstD3D11ColorConvert{GstD3D11BaseConvert: &GstD3D11BaseConvert{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewD3D11ColorConvertWithName(name string) (*GstD3D11ColorConvert, error) {
	e, err := gst.NewElementWithName("d3d11colorconvert", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11ColorConvert{GstD3D11BaseConvert: &GstD3D11BaseConvert{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Gamma conversion mode
// Default: none (0)
func (e *GstD3D11ColorConvert) SetGammaMode(value interface{}) error {
	return e.Element.SetProperty("gamma-mode", value)
}

func (e *GstD3D11ColorConvert) GetGammaMode() (interface{}, error) {
	return e.Element.GetProperty("gamma-mode")
}

// Primaries conversion mode
// Default: none (0)
func (e *GstD3D11ColorConvert) SetPrimariesMode(value interface{}) error {
	return e.Element.SetProperty("primaries-mode", value)
}

func (e *GstD3D11ColorConvert) GetPrimariesMode() (interface{}, error) {
	return e.Element.GetProperty("primaries-mode")
}

// Applied input alpha mode
// Default: unspecified (0)
func (e *GstD3D11ColorConvert) SetSrcAlphaMode(value GstD3D11AlphaMode) error {
	e.Element.SetArg("src-alpha-mode", string(value))
	return nil
}

func (e *GstD3D11ColorConvert) GetSrcAlphaMode() (GstD3D11AlphaMode, error) {
	var value GstD3D11AlphaMode
	var ok bool
	v, err := e.Element.GetProperty("src-alpha-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11AlphaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11AlphaMode")
	}
	return value, nil
}

// Output alpha mode to be applied
// Default: unspecified (0)
func (e *GstD3D11ColorConvert) SetDestAlphaMode(value GstD3D11AlphaMode) error {
	e.Element.SetArg("dest-alpha-mode", string(value))
	return nil
}

func (e *GstD3D11ColorConvert) GetDestAlphaMode() (GstD3D11AlphaMode, error) {
	var value GstD3D11AlphaMode
	var ok bool
	v, err := e.Element.GetProperty("dest-alpha-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11AlphaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11AlphaMode")
	}
	return value, nil
}

// Captures desktop screen
type GstD3D11ScreenCaptureSrc struct {
	*base.GstBaseSrc
}

func NewD3D11ScreenCaptureSrc() (*GstD3D11ScreenCaptureSrc, error) {
	e, err := gst.NewElement("d3d11screencapturesrc")
	if err != nil {
		return nil, err
	}
	return &GstD3D11ScreenCaptureSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewD3D11ScreenCaptureSrcWithName(name string) (*GstD3D11ScreenCaptureSrc, error) {
	e, err := gst.NewElementWithName("d3d11screencapturesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11ScreenCaptureSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Capture API to use
// Default: dxgi (0)
func (e *GstD3D11ScreenCaptureSrc) SetCaptureApi(value GstD3D11ScreenCaptureAPI) error {
	e.Element.SetArg("capture-api", string(value))
	return nil
}

func (e *GstD3D11ScreenCaptureSrc) GetCaptureApi() (GstD3D11ScreenCaptureAPI, error) {
	var value GstD3D11ScreenCaptureAPI
	var ok bool
	v, err := e.Element.GetProperty("capture-api")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11ScreenCaptureAPI)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11ScreenCaptureAPI")
	}
	return value, nil
}

// Zero-based index for monitor to capture (-1 = primary monitor)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstD3D11ScreenCaptureSrc) SetMonitorIndex(value int) error {
	return e.Element.SetProperty("monitor-index", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetMonitorIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("monitor-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Show border lines to capture area when WGC mode is selected
// Default: false
func (e *GstD3D11ScreenCaptureSrc) SetShowBorder(value bool) error {
	return e.Element.SetProperty("show-border", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetShowBorder() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-border")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// A HWND handle of window to capture
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstD3D11ScreenCaptureSrc) SetWindowHandle(value uint64) error {
	return e.Element.SetProperty("window-handle", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetWindowHandle() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("window-handle")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// A HMONITOR handle of monitor to capture
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstD3D11ScreenCaptureSrc) SetMonitorHandle(value uint64) error {
	return e.Element.SetProperty("monitor-handle", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetMonitorHandle() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("monitor-handle")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Whether to show mouse cursor
// Default: false
func (e *GstD3D11ScreenCaptureSrc) SetShowCursor(value bool) error {
	return e.Element.SetProperty("show-cursor", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetShowCursor() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-cursor")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Window capture mode to use if "window-handle" is set
// Default: default (0)
func (e *GstD3D11ScreenCaptureSrc) SetWindowCaptureMode(value GstD3D11WindowCaptureMode) error {
	e.Element.SetArg("window-capture-mode", string(value))
	return nil
}

func (e *GstD3D11ScreenCaptureSrc) GetWindowCaptureMode() (GstD3D11WindowCaptureMode, error) {
	var value GstD3D11WindowCaptureMode
	var ok bool
	v, err := e.Element.GetProperty("window-capture-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11WindowCaptureMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11WindowCaptureMode")
	}
	return value, nil
}

// DXGI Adapter index for creating device when WGC mode is selected (-1 for default)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstD3D11ScreenCaptureSrc) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Height of screen capture area (0 = maximum)
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11ScreenCaptureSrc) SetCropHeight(value uint) error {
	return e.Element.SetProperty("crop-height", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetCropHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Width of screen capture area (0 = maximum)
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11ScreenCaptureSrc) SetCropWidth(value uint) error {
	return e.Element.SetProperty("crop-width", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetCropWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Horizontal coordinate of top left corner for the screen capture area
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11ScreenCaptureSrc) SetCropX(value uint) error {
	return e.Element.SetProperty("crop-x", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetCropX() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Vertical coordinate of top left corner for the screen capture area
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11ScreenCaptureSrc) SetCropY(value uint) error {
	return e.Element.SetProperty("crop-y", value)
}

func (e *GstD3D11ScreenCaptureSrc) GetCropY() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// A Direct3D11 based videosink
type GstD3D11VideoSink struct {
	*GstVideoSink
}

func NewD3D11VideoSink() (*GstD3D11VideoSink, error) {
	e, err := gst.NewElement("d3d11videosink")
	if err != nil {
		return nil, err
	}
	return &GstD3D11VideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewD3D11VideoSinkWithName(name string) (*GstD3D11VideoSink, error) {
	e, err := gst.NewElementWithName("d3d11videosink", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11VideoSink{GstVideoSink: &GstVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Adapter index for creating device (-1 for default)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstD3D11VideoSink) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

func (e *GstD3D11VideoSink) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Ignored when "fullscreen-toggle-mode" does not include "property"
// Default: false
func (e *GstD3D11VideoSink) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

func (e *GstD3D11VideoSink) GetFullscreen() (bool, error) {
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

// Rotate method to use
// Default: identity (0)
func (e *GstD3D11VideoSink) SetRotateMethod(value interface{}) error {
	return e.Element.SetProperty("rotate-method", value)
}

func (e *GstD3D11VideoSink) GetRotateMethod() (interface{}, error) {
	return e.Element.GetProperty("rotate-method")
}

// Gamma conversion mode
// Default: none (0)
func (e *GstD3D11VideoSink) SetGammaMode(value interface{}) error {
	return e.Element.SetProperty("gamma-mode", value)
}

func (e *GstD3D11VideoSink) GetGammaMode() (interface{}, error) {
	return e.Element.GetProperty("gamma-mode")
}

// Primaries conversion mode
// Default: none (0)
func (e *GstD3D11VideoSink) SetPrimariesMode(value interface{}) error {
	return e.Element.SetProperty("primaries-mode", value)
}

func (e *GstD3D11VideoSink) GetPrimariesMode() (interface{}, error) {
	return e.Element.GetProperty("primaries-mode")
}

// Swapchain display format
// Default: unknown (0)
func (e *GstD3D11VideoSink) SetDisplayFormat(value GstD3D11VideoSinkDisplayFormat) error {
	e.Element.SetArg("display-format", string(value))
	return nil
}

func (e *GstD3D11VideoSink) GetDisplayFormat() (GstD3D11VideoSinkDisplayFormat, error) {
	var value GstD3D11VideoSinkDisplayFormat
	var ok bool
	v, err := e.Element.GetProperty("display-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11VideoSinkDisplayFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11VideoSinkDisplayFormat")
	}
	return value, nil
}

// Draw on user provided shared texture instead of window. When enabled, user can pass application's own texture to sink by using "draw" action signal on "begin-draw" signal handler, so that sink can draw video data on application's texture. Supported texture formats for user texture are DXGI_FORMAT_R8G8B8A8_UNORM, DXGI_FORMAT_B8G8R8A8_UNORM, and DXGI_FORMAT_R10G10B10A2_UNORM.
// Default: false
func (e *GstD3D11VideoSink) SetDrawOnSharedTexture(value bool) error {
	return e.Element.SetProperty("draw-on-shared-texture", value)
}

func (e *GstD3D11VideoSink) GetDrawOnSharedTexture() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-on-shared-texture")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Emits present signal
// Default: false
func (e *GstD3D11VideoSink) SetEmitPresent(value bool) error {
	return e.Element.SetProperty("emit-present", value)
}

func (e *GstD3D11VideoSink) GetEmitPresent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("emit-present")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, navigation events are sent upstream
// Default: true
func (e *GstD3D11VideoSink) SetEnableNavigationEvents(value bool) error {
	return e.Element.SetProperty("enable-navigation-events", value)
}

func (e *GstD3D11VideoSink) GetEnableNavigationEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-navigation-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, scaling will respect original aspect ratio
// Default: true
func (e *GstD3D11VideoSink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

func (e *GstD3D11VideoSink) GetForceAspectRatio() (bool, error) {
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

// Full screen toggle mode used to trigger fullscreen mode change
// Default: none
func (e *GstD3D11VideoSink) SetFullscreenToggleMode(value GstD3D11WindowFullscreenToggleMode) error {
	e.Element.SetArg("fullscreen-toggle-mode", fmt.Sprint(value))
	return nil
}

func (e *GstD3D11VideoSink) GetFullscreenToggleMode() (GstD3D11WindowFullscreenToggleMode, error) {
	var value GstD3D11WindowFullscreenToggleMode
	var ok bool
	v, err := e.Element.GetProperty("fullscreen-toggle-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11WindowFullscreenToggleMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11WindowFullscreenToggleMode")
	}
	return value, nil
}

// The render rectangle ('<x, y, width, height>')

func (e *GstD3D11VideoSink) SetRenderRectangle(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("render-rectangle", value)
}

// Direct3D11/DXVA based VP9 video decoder
type GstD3D11Vp9Dec struct {
	*GstDxvaVp9Decoder
}

func NewD3D11Vp9Dec() (*GstD3D11Vp9Dec, error) {
	e, err := gst.NewElement("d3d11vp9dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Vp9Dec{GstDxvaVp9Decoder: &GstDxvaVp9Decoder{GstVp9Decoder: &GstVp9Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D11Vp9DecWithName(name string) (*GstD3D11Vp9Dec, error) {
	e, err := gst.NewElementWithName("d3d11vp9dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Vp9Dec{GstDxvaVp9Decoder: &GstDxvaVp9Decoder{GstVp9Decoder: &GstVp9Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Vendor ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Vp9Dec) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D11Vp9Dec) GetAdapterLuid() (int64, error) {
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

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Vp9Dec) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Provides application renderable Direct3D11 render target view
type GstD3D11Overlay struct {
	*GstD3D11BaseFilter
}

func NewD3D11Overlay() (*GstD3D11Overlay, error) {
	e, err := gst.NewElement("d3d11overlay")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Overlay{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewD3D11OverlayWithName(name string) (*GstD3D11Overlay, error) {
	e, err := gst.NewElementWithName("d3d11overlay", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Overlay{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Resizes video using Direct3D11
type GstD3D11Scale struct {
	*GstD3D11BaseConvert
}

func NewD3D11Scale() (*GstD3D11Scale, error) {
	e, err := gst.NewElement("d3d11scale")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Scale{GstD3D11BaseConvert: &GstD3D11BaseConvert{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewD3D11ScaleWithName(name string) (*GstD3D11Scale, error) {
	e, err := gst.NewElementWithName("d3d11scale", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Scale{GstD3D11BaseConvert: &GstD3D11BaseConvert{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Add black borders if necessary to keep the display aspect ratio
// Default: true
func (e *GstD3D11Scale) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

func (e *GstD3D11Scale) GetAddBorders() (bool, error) {
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

// Border color to use in ARGB64 format
// Default: 18446462598732840960, Min: 0, Max: 18446744073709551615
func (e *GstD3D11Scale) SetBorderColor(value uint64) error {
	return e.Element.SetProperty("border-color", value)
}

func (e *GstD3D11Scale) GetBorderColor() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("border-color")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// A Direct3D11 based deinterlacer bin
type GstD3D11DeinterlaceBin struct {
	*gst.Bin
}

func NewD3D11DeinterlaceBin() (*GstD3D11DeinterlaceBin, error) {
	e, err := gst.NewElement("d3d11deinterlace")
	if err != nil {
		return nil, err
	}
	return &GstD3D11DeinterlaceBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewD3D11DeinterlaceBinWithName(name string) (*GstD3D11DeinterlaceBin, error) {
	e, err := gst.NewElementWithName("d3d11deinterlace", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11DeinterlaceBin{Bin: &gst.Bin{Element: e}}, nil
}

// DXGI Adapter index for creating device
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11DeinterlaceBin) GetAdapter() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11DeinterlaceBin) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Deinterlace Method. Use can set multiple methods as a flagset and element will select one of method automatically. If deinterlacing device failed to deinterlace with given mode, fallback might happen by the device
// Default: mocomp+adaptive+bob+blend
func (e *GstD3D11DeinterlaceBin) SetMethod(value GstD3D11DeinterlaceMethod) error {
	e.Element.SetArg("method", fmt.Sprint(value))
	return nil
}

func (e *GstD3D11DeinterlaceBin) GetMethod() (GstD3D11DeinterlaceMethod, error) {
	var value GstD3D11DeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11DeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11DeinterlaceMethod")
	}
	return value, nil
}

// Handle Quality-of-Service events
// Default: false
func (e *GstD3D11DeinterlaceBin) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstD3D11DeinterlaceBin) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Set of supported deinterlace methods by device
// Default: mocomp+adaptive+bob+blend
func (e *GstD3D11DeinterlaceBin) GetSupportedMethods() (GstD3D11DeinterlaceMethod, error) {
	var value GstD3D11DeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("supported-methods")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11DeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11DeinterlaceMethod")
	}
	return value, nil
}

// DXGI Vendor ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11DeinterlaceBin) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// A Direct3D11 based deinterlacer
type GstD3D11Deinterlace struct {
	*base.GstBaseTransform
}

func NewD3D11Deinterlace() (*GstD3D11Deinterlace, error) {
	e, err := gst.NewElement("d3d11deinterlaceelement")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Deinterlace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewD3D11DeinterlaceWithName(name string) (*GstD3D11Deinterlace, error) {
	e, err := gst.NewElementWithName("d3d11deinterlaceelement", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Deinterlace{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// DXGI Adapter index for creating device
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Deinterlace) GetAdapter() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Deinterlace) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Deinterlace Method. Use can set multiple methods as a flagset and element will select one of method automatically. If deinterlacing device failed to deinterlace with given mode, fallback might happen by the device
// Default: mocomp+adaptive+bob+blend
func (e *GstD3D11Deinterlace) SetMethod(value GstD3D11DeinterlaceMethod) error {
	e.Element.SetArg("method", fmt.Sprint(value))
	return nil
}

func (e *GstD3D11Deinterlace) GetMethod() (GstD3D11DeinterlaceMethod, error) {
	var value GstD3D11DeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11DeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11DeinterlaceMethod")
	}
	return value, nil
}

// Set of supported deinterlace methods by device
// Default: mocomp+adaptive+bob+blend
func (e *GstD3D11Deinterlace) GetSupportedMethods() (GstD3D11DeinterlaceMethod, error) {
	var value GstD3D11DeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("supported-methods")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11DeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11DeinterlaceMethod")
	}
	return value, nil
}

// DXGI Vendor ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Deinterlace) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Direct3D11/DXVA based VP8 video decoder
type GstD3D11Vp8Dec struct {
	*GstDxvaVp8Decoder
}

func NewD3D11Vp8Dec() (*GstD3D11Vp8Dec, error) {
	e, err := gst.NewElement("d3d11vp8dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Vp8Dec{GstDxvaVp8Decoder: &GstDxvaVp8Decoder{GstVp8Decoder: &GstVp8Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D11Vp8DecWithName(name string) (*GstD3D11Vp8Dec, error) {
	e, err := gst.NewElementWithName("d3d11vp8dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Vp8Dec{GstDxvaVp8Decoder: &GstDxvaVp8Decoder{GstVp8Decoder: &GstVp8Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D11Vp8Dec) GetAdapterLuid() (int64, error) {
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

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Vp8Dec) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Vendor ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Vp8Dec) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Creates a test video stream
type GstD3D11TestSrc struct {
	*base.GstBaseSrc
}

func NewD3D11TestSrc() (*GstD3D11TestSrc, error) {
	e, err := gst.NewElement("d3d11testsrc")
	if err != nil {
		return nil, err
	}
	return &GstD3D11TestSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewD3D11TestSrcWithName(name string) (*GstD3D11TestSrc, error) {
	e, err := gst.NewElementWithName("d3d11testsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11TestSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// DXGI Adapter index (-1 for any device)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstD3D11TestSrc) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

func (e *GstD3D11TestSrc) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Global alpha value to use
// Default: 1, Min: 0, Max: 1
func (e *GstD3D11TestSrc) SetAlpha(value float32) error {
	return e.Element.SetProperty("alpha", value)
}

func (e *GstD3D11TestSrc) GetAlpha() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// alpha mode to use
// Default: unspecified (0)
func (e *GstD3D11TestSrc) SetAlphaMode(value GstD3D11AlphaMode) error {
	e.Element.SetArg("alpha-mode", string(value))
	return nil
}

func (e *GstD3D11TestSrc) GetAlphaMode() (GstD3D11AlphaMode, error) {
	var value GstD3D11AlphaMode
	var ok bool
	v, err := e.Element.GetProperty("alpha-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11AlphaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11AlphaMode")
	}
	return value, nil
}

// Whether to act as a live source
// Default: false
func (e *GstD3D11TestSrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstD3D11TestSrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Type of test pattern to generate
// Default: smpte (0)
func (e *GstD3D11TestSrc) SetPattern(value GstD3D11TestSrcPattern) error {
	e.Element.SetArg("pattern", string(value))
	return nil
}

func (e *GstD3D11TestSrc) GetPattern() (GstD3D11TestSrcPattern, error) {
	var value GstD3D11TestSrcPattern
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11TestSrcPattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11TestSrcPattern")
	}
	return value, nil
}

// Uploads data into Direct3D11 texture memory
type GstD3D11Upload struct {
	*GstD3D11BaseFilter
}

func NewD3D11Upload() (*GstD3D11Upload, error) {
	e, err := gst.NewElement("d3d11upload")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Upload{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewD3D11UploadWithName(name string) (*GstD3D11Upload, error) {
	e, err := gst.NewElementWithName("d3d11upload", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Upload{GstD3D11BaseFilter: &GstD3D11BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Direct3D11/DXVA based H.264 video decoder
type GstD3D11H264Dec struct {
	*GstDxvaH264Decoder
}

func NewD3D11H264Dec() (*GstD3D11H264Dec, error) {
	e, err := gst.NewElement("d3d11h264dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D11H264Dec{GstDxvaH264Decoder: &GstDxvaH264Decoder{GstH264Decoder: &GstH264Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D11H264DecWithName(name string) (*GstD3D11H264Dec, error) {
	e, err := gst.NewElementWithName("d3d11h264dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11H264Dec{GstDxvaH264Decoder: &GstDxvaH264Decoder{GstH264Decoder: &GstH264Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D11H264Dec) GetAdapterLuid() (int64, error) {
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

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11H264Dec) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Vendor ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11H264Dec) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Direct3D11/DXVA based MPEG2 video decoder
type GstD3D11Mpeg2Dec struct {
	*GstDxvaMpeg2Decoder
}

func NewD3D11Mpeg2Dec() (*GstD3D11Mpeg2Dec, error) {
	e, err := gst.NewElement("d3d11mpeg2dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D11Mpeg2Dec{GstDxvaMpeg2Decoder: &GstDxvaMpeg2Decoder{GstMpeg2Decoder: &GstMpeg2Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D11Mpeg2DecWithName(name string) (*GstD3D11Mpeg2Dec, error) {
	e, err := gst.NewElementWithName("d3d11mpeg2dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D11Mpeg2Dec{GstDxvaMpeg2Decoder: &GstDxvaMpeg2Decoder{GstMpeg2Decoder: &GstMpeg2Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Vendor ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Mpeg2Dec) GetVendorId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("vendor-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D11Mpeg2Dec) GetAdapterLuid() (int64, error) {
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

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D11Mpeg2Dec) GetDeviceId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstD3D11BaseFilter struct {
	*base.GstBaseTransform
}

// Adapter index for creating device (-1 for default)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstD3D11BaseFilter) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

func (e *GstD3D11BaseFilter) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstD3D11CompositorSizingPolicy string

const (
	GstD3D11CompositorSizingPolicyNone            GstD3D11CompositorSizingPolicy = "none"              // None: Image is scaled to fill configured destination rectangle without padding or keeping the aspect ratio
	GstD3D11CompositorSizingPolicyKeepAspectRatio GstD3D11CompositorSizingPolicy = "keep-aspect-ratio" // Keep Aspect Ratio: Image is scaled to fit destination rectangle specified by GstD3D11CompositorPad:{xpos, ypos, width, height} with preserved aspect ratio. Resulting image will be centered in the destination rectangle with padding if necessary
)

type GstD3D11VideoSinkDisplayFormat string

const (
	GstD3D11VideoSinkDisplayFormatUnknown          GstD3D11VideoSinkDisplayFormat = "unknown"           // DXGI_FORMAT_UNKNOWN
	GstD3D11VideoSinkDisplayFormatR10g10b10a2Unorm GstD3D11VideoSinkDisplayFormat = "r10g10b10a2-unorm" // DXGI_FORMAT_R10G10B10A2_UNORM
	GstD3D11VideoSinkDisplayFormatR8g8b8a8Unorm    GstD3D11VideoSinkDisplayFormat = "r8g8b8a8-unorm"    // DXGI_FORMAT_R8G8B8A8_UNORM
	GstD3D11VideoSinkDisplayFormatB8g8r8a8Unorm    GstD3D11VideoSinkDisplayFormat = "b8g8r8a8-unorm"    // DXGI_FORMAT_B8G8R8A8_UNORM
)

type GstD3D11CompositorOperator string

const (
	GstD3D11CompositorOperatorSource GstD3D11CompositorOperator = "source" // Source
	GstD3D11CompositorOperatorOver   GstD3D11CompositorOperator = "over"   // Over
)

type GstD3D11SamplingMethod string

const (
	GstD3D11SamplingMethodNearestNeighbour   GstD3D11SamplingMethod = "nearest-neighbour"   // Nearest Neighbour
	GstD3D11SamplingMethodBilinear           GstD3D11SamplingMethod = "bilinear"            // Bilinear
	GstD3D11SamplingMethodLinearMinification GstD3D11SamplingMethod = "linear-minification" // Linear minification, point magnification
)

type GstD3D11AlphaMode string

const (
	GstD3D11AlphaModeUnspecified   GstD3D11AlphaMode = "unspecified"   // Unspecified
	GstD3D11AlphaModePremultiplied GstD3D11AlphaMode = "premultiplied" // Premultiplied
	GstD3D11AlphaModeStraight      GstD3D11AlphaMode = "straight"      // Straight
)

type GstD3D11CompositorBackground string

const (
	GstD3D11CompositorBackgroundChecker     GstD3D11CompositorBackground = "checker"     // Checker pattern
	GstD3D11CompositorBackgroundBlack       GstD3D11CompositorBackground = "black"       // Black
	GstD3D11CompositorBackgroundWhite       GstD3D11CompositorBackground = "white"       // White
	GstD3D11CompositorBackgroundTransparent GstD3D11CompositorBackground = "transparent" // Transparent Background to enable further compositing
)

type GstD3D11DeinterlaceMethod int

const (
	GstD3D11DeinterlaceMethodBlend    GstD3D11DeinterlaceMethod = 0x00000001 // Blend: Blending top/bottom field pictures into one frame. Framerate will be preserved (e.g., 60i -> 30p)
	GstD3D11DeinterlaceMethodBob      GstD3D11DeinterlaceMethod = 0x00000002 // Bob: Interpolating missing lines by using the adjacent lines. Framerate will be doubled (e,g, 60i -> 60p)
	GstD3D11DeinterlaceMethodAdaptive GstD3D11DeinterlaceMethod = 0x00000004 // Adaptive: Interpolating missing lines by using spatial/temporal references. Framerate will be doubled (e,g, 60i -> 60p)
	GstD3D11DeinterlaceMethodMocomp   GstD3D11DeinterlaceMethod = 0x00000008 // Motion Compensation: Recreating missing lines by using motion vector. Framerate will be doubled (e,g, 60i -> 60p)
)

type GstD3D11ScreenCaptureAPI string

const (
	GstD3D11ScreenCaptureAPIDxgi GstD3D11ScreenCaptureAPI = "dxgi" // DXGI Desktop Duplication
	GstD3D11ScreenCaptureAPIWgc  GstD3D11ScreenCaptureAPI = "wgc"  // Windows Graphics Capture
)

type GstD3D11TestSrcPattern string

const (
	GstD3D11TestSrcPatternSmpte     GstD3D11TestSrcPattern = "smpte"      // SMPTE 100%% color bars
	GstD3D11TestSrcPatternSnow      GstD3D11TestSrcPattern = "snow"       // Random (television snow)
	GstD3D11TestSrcPatternBlack     GstD3D11TestSrcPattern = "black"      // 100%% Black
	GstD3D11TestSrcPatternWhite     GstD3D11TestSrcPattern = "white"      // 100%% White
	GstD3D11TestSrcPatternRed       GstD3D11TestSrcPattern = "red"        // Red
	GstD3D11TestSrcPatternGreen     GstD3D11TestSrcPattern = "green"      // Green
	GstD3D11TestSrcPatternBlue      GstD3D11TestSrcPattern = "blue"       // Blue
	GstD3D11TestSrcPatternCheckers1 GstD3D11TestSrcPattern = "checkers-1" // Checkers 1px
	GstD3D11TestSrcPatternCheckers2 GstD3D11TestSrcPattern = "checkers-2" // Checkers 2px
	GstD3D11TestSrcPatternCheckers4 GstD3D11TestSrcPattern = "checkers-4" // Checkers 4px
	GstD3D11TestSrcPatternCheckers8 GstD3D11TestSrcPattern = "checkers-8" // Checkers 8px
	GstD3D11TestSrcPatternCircular  GstD3D11TestSrcPattern = "circular"   // Circular
	GstD3D11TestSrcPatternBlink     GstD3D11TestSrcPattern = "blink"      // Blink
	GstD3D11TestSrcPatternBall      GstD3D11TestSrcPattern = "ball"       // Moving ball
)

type GstD3D11WindowFullscreenToggleMode int

const (
	GstD3D11WindowFullscreenToggleModeNone     GstD3D11WindowFullscreenToggleMode = 0x00000000 // GST_D3D11_WINDOW_FULLSCREEN_TOGGLE_MODE_NONE
	GstD3D11WindowFullscreenToggleModeAltEnter GstD3D11WindowFullscreenToggleMode = 0x00000002 // GST_D3D11_WINDOW_FULLSCREEN_TOGGLE_MODE_ALT_ENTER
	GstD3D11WindowFullscreenToggleModeProperty GstD3D11WindowFullscreenToggleMode = 0x00000004 // GST_D3D11_WINDOW_FULLSCREEN_TOGGLE_MODE_PROPERTY
)

type GstD3D11BaseConvert struct {
	*GstD3D11BaseFilter
}

// Method used for sampling
// Default: bilinear (1)
func (e *GstD3D11BaseConvert) SetMethod(value GstD3D11SamplingMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstD3D11BaseConvert) GetMethod() (GstD3D11SamplingMethod, error) {
	var value GstD3D11SamplingMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstD3D11SamplingMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstD3D11SamplingMethod")
	}
	return value, nil
}

type GstD3D11IpcIOMode string

const (
	GstD3D11IpcIOModeCopy   GstD3D11IpcIOMode = "copy"   // Copy remote texture
	GstD3D11IpcIOModeImport GstD3D11IpcIOMode = "import" // Import remote texture
)

type GstD3D11WindowCaptureMode string

const (
	GstD3D11WindowCaptureModeDefault GstD3D11WindowCaptureMode = "default" // Capture entire window area
	GstD3D11WindowCaptureModeClient  GstD3D11WindowCaptureMode = "client"  // Capture client area
)
