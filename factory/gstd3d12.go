// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Direct3D12/DXVA based AV1 video decoder
type GstD3D12AV1Dec struct {
	*GstDxvaAV1Decoder
}

func NewD3D12AV1Dec() (*GstD3D12AV1Dec, error) {
	e, err := gst.NewElement("d3d12av1dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D12AV1Dec{GstDxvaAV1Decoder: &GstDxvaAV1Decoder{GstAV1Decoder: &GstAV1Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D12AV1DecWithName(name string) (*GstD3D12AV1Dec, error) {
	e, err := gst.NewElementWithName("d3d12av1dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D12AV1Dec{GstDxvaAV1Decoder: &GstDxvaAV1Decoder{GstAV1Decoder: &GstAV1Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D12AV1Dec) GetAdapterLuid() (int64, error) {
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
func (e *GstD3D12AV1Dec) GetDeviceId() (uint, error) {
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
func (e *GstD3D12AV1Dec) GetVendorId() (uint, error) {
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

// Downloads Direct3D12 texture memory into system memory
type GstD3D12Download struct {
	*GstD3D12BaseFilter
}

func NewD3D12Download() (*GstD3D12Download, error) {
	e, err := gst.NewElement("d3d12download")
	if err != nil {
		return nil, err
	}
	return &GstD3D12Download{GstD3D12BaseFilter: &GstD3D12BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewD3D12DownloadWithName(name string) (*GstD3D12Download, error) {
	e, err := gst.NewElementWithName("d3d12download", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D12Download{GstD3D12BaseFilter: &GstD3D12BaseFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Direct3D12/DXVA based H.264 video decoder
type GstD3D12H264Dec struct {
	*GstDxvaH264Decoder
}

func NewD3D12H264Dec() (*GstD3D12H264Dec, error) {
	e, err := gst.NewElement("d3d12h264dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D12H264Dec{GstDxvaH264Decoder: &GstDxvaH264Decoder{GstH264Decoder: &GstH264Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D12H264DecWithName(name string) (*GstD3D12H264Dec, error) {
	e, err := gst.NewElementWithName("d3d12h264dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D12H264Dec{GstDxvaH264Decoder: &GstDxvaH264Decoder{GstH264Decoder: &GstH264Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D12H264Dec) GetAdapterLuid() (int64, error) {
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
func (e *GstD3D12H264Dec) GetDeviceId() (uint, error) {
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
func (e *GstD3D12H264Dec) GetVendorId() (uint, error) {
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

// Direct3D12/DXVA based H.265 video decoder
type GstD3D12H265Dec struct {
	*GstDxvaH265Decoder
}

func NewD3D12H265Dec() (*GstD3D12H265Dec, error) {
	e, err := gst.NewElement("d3d12h265dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D12H265Dec{GstDxvaH265Decoder: &GstDxvaH265Decoder{GstH265Decoder: &GstH265Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D12H265DecWithName(name string) (*GstD3D12H265Dec, error) {
	e, err := gst.NewElementWithName("d3d12h265dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D12H265Dec{GstDxvaH265Decoder: &GstDxvaH265Decoder{GstH265Decoder: &GstH265Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Device ID
// Default: 0, Min: 0, Max: -1
func (e *GstD3D12H265Dec) GetDeviceId() (uint, error) {
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
func (e *GstD3D12H265Dec) GetVendorId() (uint, error) {
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
func (e *GstD3D12H265Dec) GetAdapterLuid() (int64, error) {
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

// Direct3D12/DXVA based VP9 video decoder
type GstD3D12Vp9Dec struct {
	*GstDxvaVp9Decoder
}

func NewD3D12Vp9Dec() (*GstD3D12Vp9Dec, error) {
	e, err := gst.NewElement("d3d12vp9dec")
	if err != nil {
		return nil, err
	}
	return &GstD3D12Vp9Dec{GstDxvaVp9Decoder: &GstDxvaVp9Decoder{GstVp9Decoder: &GstVp9Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

func NewD3D12Vp9DecWithName(name string) (*GstD3D12Vp9Dec, error) {
	e, err := gst.NewElementWithName("d3d12vp9dec", name)
	if err != nil {
		return nil, err
	}
	return &GstD3D12Vp9Dec{GstDxvaVp9Decoder: &GstDxvaVp9Decoder{GstVp9Decoder: &GstVp9Decoder{GstVideoDecoder: &GstVideoDecoder{Element: e}}}}, nil
}

// DXGI Adapter LUID (Locally Unique Identifier) of created device
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstD3D12Vp9Dec) GetAdapterLuid() (int64, error) {
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
func (e *GstD3D12Vp9Dec) GetDeviceId() (uint, error) {
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
func (e *GstD3D12Vp9Dec) GetVendorId() (uint, error) {
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

type GstD3D12BaseFilter struct {
	*base.GstBaseTransform
}
