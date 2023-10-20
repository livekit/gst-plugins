// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Warps the picture into an arc shaped form
type GstCircle struct {
	*GstCircleGeometricTransform
}

func NewCircle() (*GstCircle, error) {
	e, err := gst.NewElement("circle")
	if err != nil {
		return nil, err
	}
	return &GstCircle{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewCircleWithName(name string) (*GstCircle, error) {
	e, err := gst.NewElementWithName("circle", name)
	if err != nil {
		return nil, err
	}
	return &GstCircle{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Angle at which the arc starts in radians
// Default: 0, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstCircle) SetAngle(value float64) error {
	return e.Element.SetProperty("angle", value)
}

func (e *GstCircle) GetAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Height of the arc
// Default: 20, Min: 0, Max: 2147483647
func (e *GstCircle) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstCircle) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Length of the arc in radians
// Default: 3.14159, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstCircle) SetSpreadAngle(value float64) error {
	return e.Element.SetProperty("spread-angle", value)
}

func (e *GstCircle) GetSpreadAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("spread-angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Applies 'kaleidoscope' geometric transform to the image
type GstKaleidoscope struct {
	*GstCircleGeometricTransform
}

func NewKaleidoscope() (*GstKaleidoscope, error) {
	e, err := gst.NewElement("kaleidoscope")
	if err != nil {
		return nil, err
	}
	return &GstKaleidoscope{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewKaleidoscopeWithName(name string) (*GstKaleidoscope, error) {
	e, err := gst.NewElementWithName("kaleidoscope", name)
	if err != nil {
		return nil, err
	}
	return &GstKaleidoscope{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// primary angle in radians of the kaleidoscope effect
// Default: 0, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstKaleidoscope) SetAngle(value float64) error {
	return e.Element.SetProperty("angle", value)
}

func (e *GstKaleidoscope) GetAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// secondary angle in radians of the kaleidoscope effect
// Default: 0, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstKaleidoscope) SetAngle2(value float64) error {
	return e.Element.SetProperty("angle2", value)
}

func (e *GstKaleidoscope) GetAngle2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("angle2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Number of sides of the kaleidoscope
// Default: 3, Min: 2, Max: 2147483647
func (e *GstKaleidoscope) SetSides(value int) error {
	return e.Element.SetProperty("sides", value)
}

func (e *GstKaleidoscope) GetSides() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sides")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Split the image into two halves and reflect one over each other
type GstMirror struct {
	*GstGeometricTransform
}

func NewMirror() (*GstMirror, error) {
	e, err := gst.NewElement("mirror")
	if err != nil {
		return nil, err
	}
	return &GstMirror{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewMirrorWithName(name string) (*GstMirror, error) {
	e, err := gst.NewElementWithName("mirror", name)
	if err != nil {
		return nil, err
	}
	return &GstMirror{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// How to split the video frame and which side reflect
// Default: left (0)
func (e *GstMirror) SetMode(value GstMirrorMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstMirror) GetMode() (GstMirrorMode, error) {
	var value GstMirrorMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMirrorMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMirrorMode")
	}
	return value, nil
}

// Apply a 2D perspective transform
type GstPerspective struct {
	*GstGeometricTransform
}

func NewPerspective() (*GstPerspective, error) {
	e, err := gst.NewElement("perspective")
	if err != nil {
		return nil, err
	}
	return &GstPerspective{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewPerspectiveWithName(name string) (*GstPerspective, error) {
	e, err := gst.NewElementWithName("perspective", name)
	if err != nil {
		return nil, err
	}
	return &GstPerspective{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Matrix of dimension 3x3 to use in the 2D transform, passed as an array of 9 elements in row-major order

func (e *GstPerspective) SetMatrix(value interface{}) error {
	return e.Element.SetProperty("matrix", value)
}

func (e *GstPerspective) GetMatrix() (interface{}, error) {
	return e.Element.GetProperty("matrix")
}

// Rotates the picture by an arbitrary angle
type GstRotate struct {
	*GstGeometricTransform
}

func NewRotate() (*GstRotate, error) {
	e, err := gst.NewElement("rotate")
	if err != nil {
		return nil, err
	}
	return &GstRotate{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewRotateWithName(name string) (*GstRotate, error) {
	e, err := gst.NewElementWithName("rotate", name)
	if err != nil {
		return nil, err
	}
	return &GstRotate{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Angle by which the picture is rotated, in radians
// Default: 0, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstRotate) SetAngle(value float64) error {
	return e.Element.SetProperty("angle", value)
}

func (e *GstRotate) GetAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Adds a protuberance in the center point
type GstBulge struct {
	*GstCircleGeometricTransform
}

func NewBulge() (*GstBulge, error) {
	e, err := gst.NewElement("bulge")
	if err != nil {
		return nil, err
	}
	return &GstBulge{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewBulgeWithName(name string) (*GstBulge, error) {
	e, err := gst.NewElementWithName("bulge", name)
	if err != nil {
		return nil, err
	}
	return &GstBulge{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Zoom of the bulge effect
// Default: 3, Min: 1, Max: 100
func (e *GstBulge) SetZoom(value float64) error {
	return e.Element.SetProperty("zoom", value)
}

func (e *GstBulge) GetZoom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Applies a marbling effect to the image
type GstMarble struct {
	*GstGeometricTransform
}

func NewMarble() (*GstMarble, error) {
	e, err := gst.NewElement("marble")
	if err != nil {
		return nil, err
	}
	return &GstMarble{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewMarbleWithName(name string) (*GstMarble, error) {
	e, err := gst.NewElementWithName("marble", name)
	if err != nil {
		return nil, err
	}
	return &GstMarble{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Amount of effect
// Default: 1, Min: 0, Max: 1
func (e *GstMarble) SetAmount(value float64) error {
	return e.Element.SetProperty("amount", value)
}

func (e *GstMarble) GetAmount() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Turbulence of the effect
// Default: 4, Min: 0, Max: 1
func (e *GstMarble) SetTurbulence(value float64) error {
	return e.Element.SetProperty("turbulence", value)
}

func (e *GstMarble) GetTurbulence() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("turbulence")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X scale of the texture
// Default: 4, Min: 0, Max: 1.79769e+308
func (e *GstMarble) SetXScale(value float64) error {
	return e.Element.SetProperty("x-scale", value)
}

func (e *GstMarble) GetXScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y scale of the texture
// Default: 4, Min: 0, Max: 1.79769e+308
func (e *GstMarble) SetYScale(value float64) error {
	return e.Element.SetProperty("y-scale", value)
}

func (e *GstMarble) GetYScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Applies 'sphere' geometric transform to the image
type GstSphere struct {
	*GstCircleGeometricTransform
}

func NewSphere() (*GstSphere, error) {
	e, err := gst.NewElement("sphere")
	if err != nil {
		return nil, err
	}
	return &GstSphere{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewSphereWithName(name string) (*GstSphere, error) {
	e, err := gst.NewElementWithName("sphere", name)
	if err != nil {
		return nil, err
	}
	return &GstSphere{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// refraction index
// Default: 1.5, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstSphere) SetRefraction(value float64) error {
	return e.Element.SetProperty("refraction", value)
}

func (e *GstSphere) GetRefraction() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("refraction")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Light tunnel effect
type GstTunnel struct {
	*GstCircleGeometricTransform
}

func NewTunnel() (*GstTunnel, error) {
	e, err := gst.NewElement("tunnel")
	if err != nil {
		return nil, err
	}
	return &GstTunnel{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewTunnelWithName(name string) (*GstTunnel, error) {
	e, err := gst.NewElementWithName("tunnel", name)
	if err != nil {
		return nil, err
	}
	return &GstTunnel{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Twists the image from the center out
type GstTwirl struct {
	*GstCircleGeometricTransform
}

func NewTwirl() (*GstTwirl, error) {
	e, err := gst.NewElement("twirl")
	if err != nil {
		return nil, err
	}
	return &GstTwirl{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewTwirlWithName(name string) (*GstTwirl, error) {
	e, err := gst.NewElementWithName("twirl", name)
	if err != nil {
		return nil, err
	}
	return &GstTwirl{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// This is the angle in radians by which pixels at the nearest edge of the image will move
// Default: 3.14159, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstTwirl) SetAngle(value float64) error {
	return e.Element.SetProperty("angle", value)
}

func (e *GstTwirl) GetAngle() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Diffuses the image by moving its pixels in random directions
type GstDiffuse struct {
	*GstGeometricTransform
}

func NewDiffuse() (*GstDiffuse, error) {
	e, err := gst.NewElement("diffuse")
	if err != nil {
		return nil, err
	}
	return &GstDiffuse{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewDiffuseWithName(name string) (*GstDiffuse, error) {
	e, err := gst.NewElementWithName("diffuse", name)
	if err != nil {
		return nil, err
	}
	return &GstDiffuse{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Scale of the texture
// Default: 4, Min: 1, Max: 1.79769e+308
func (e *GstDiffuse) SetScale(value float64) error {
	return e.Element.SetProperty("scale", value)
}

func (e *GstDiffuse) GetScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Simulate a fisheye lens by zooming on the center of the image and compressing the edges
type GstFisheye struct {
	*GstGeometricTransform
}

func NewFisheye() (*GstFisheye, error) {
	e, err := gst.NewElement("fisheye")
	if err != nil {
		return nil, err
	}
	return &GstFisheye{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewFisheyeWithName(name string) (*GstFisheye, error) {
	e, err := gst.NewElementWithName("fisheye", name)
	if err != nil {
		return nil, err
	}
	return &GstFisheye{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Applies 'pinch' geometric transform to the image
type GstPinch struct {
	*GstCircleGeometricTransform
}

func NewPinch() (*GstPinch, error) {
	e, err := gst.NewElement("pinch")
	if err != nil {
		return nil, err
	}
	return &GstPinch{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewPinchWithName(name string) (*GstPinch, error) {
	e, err := gst.NewElementWithName("pinch", name)
	if err != nil {
		return nil, err
	}
	return &GstPinch{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// intensity of the pinch effect
// Default: 0.5, Min: -1, Max: 1
func (e *GstPinch) SetIntensity(value float64) error {
	return e.Element.SetProperty("intensity", value)
}

func (e *GstPinch) GetIntensity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("intensity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Stretch the image in a circle around the center point
type GstStretch struct {
	*GstCircleGeometricTransform
}

func NewStretch() (*GstStretch, error) {
	e, err := gst.NewElement("stretch")
	if err != nil {
		return nil, err
	}
	return &GstStretch{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewStretchWithName(name string) (*GstStretch, error) {
	e, err := gst.NewElementWithName("stretch", name)
	if err != nil {
		return nil, err
	}
	return &GstStretch{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Intensity of the stretch effect
// Default: 0.5, Min: 0, Max: 1
func (e *GstStretch) SetIntensity(value float64) error {
	return e.Element.SetProperty("intensity", value)
}

func (e *GstStretch) GetIntensity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("intensity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Creates a water ripple effect on the image
type GstWaterRipple struct {
	*GstCircleGeometricTransform
}

func NewWaterRipple() (*GstWaterRipple, error) {
	e, err := gst.NewElement("waterripple")
	if err != nil {
		return nil, err
	}
	return &GstWaterRipple{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewWaterRippleWithName(name string) (*GstWaterRipple, error) {
	e, err := gst.NewElementWithName("waterripple", name)
	if err != nil {
		return nil, err
	}
	return &GstWaterRipple{GstCircleGeometricTransform: &GstCircleGeometricTransform{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// phase
// Default: 0, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstWaterRipple) SetPhase(value float64) error {
	return e.Element.SetProperty("phase", value)
}

func (e *GstWaterRipple) GetPhase() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("phase")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// wavelength
// Default: 16, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstWaterRipple) SetWavelength(value float64) error {
	return e.Element.SetProperty("wavelength", value)
}

func (e *GstWaterRipple) GetWavelength() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("wavelength")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// amplitude
// Default: 10, Min: -1.79769e+308, Max: 1.79769e+308
func (e *GstWaterRipple) SetAmplitude(value float64) error {
	return e.Element.SetProperty("amplitude", value)
}

func (e *GstWaterRipple) GetAmplitude() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("amplitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Distort center part of the image into a square
type GstSquare struct {
	*GstGeometricTransform
}

func NewSquare() (*GstSquare, error) {
	e, err := gst.NewElement("square")
	if err != nil {
		return nil, err
	}
	return &GstSquare{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewSquareWithName(name string) (*GstSquare, error) {
	e, err := gst.NewElementWithName("square", name)
	if err != nil {
		return nil, err
	}
	return &GstSquare{GstGeometricTransform: &GstGeometricTransform{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Height of the square, relative to the frame height
// Default: 0.5, Min: 0, Max: 1
func (e *GstSquare) SetHeight(value float64) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstSquare) GetHeight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Width of the square, relative to the frame width
// Default: 0.5, Min: 0, Max: 1
func (e *GstSquare) SetWidth(value float64) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstSquare) GetWidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Zoom amount in the center region
// Default: 2, Min: 1, Max: 100
func (e *GstSquare) SetZoom(value float64) error {
	return e.Element.SetProperty("zoom", value)
}

func (e *GstSquare) GetZoom() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

type GstMirrorMode string

const (
	GstMirrorModeLeft   GstMirrorMode = "left"   // Split horizontally and reflect left into right
	GstMirrorModeRight  GstMirrorMode = "right"  // Split horizontally and reflect right into left
	GstMirrorModeTop    GstMirrorMode = "top"    // Split vertically and reflect top into bottom
	GstMirrorModeBottom GstMirrorMode = "bottom" // Split vertically and reflect bottom into top
)

type GstCircleGeometricTransform struct {
	*GstGeometricTransform
}

// radius of the circle_geometric_transform effect
// Default: 0.35, Min: 0, Max: 1
func (e *GstCircleGeometricTransform) SetRadius(value float64) error {
	return e.Element.SetProperty("radius", value)
}

func (e *GstCircleGeometricTransform) GetRadius() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("radius")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X axis center of the circle_geometric_transform effect
// Default: 0.5, Min: 0, Max: 1
func (e *GstCircleGeometricTransform) SetXCenter(value float64) error {
	return e.Element.SetProperty("x-center", value)
}

func (e *GstCircleGeometricTransform) GetXCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y axis center of the circle_geometric_transform effect
// Default: 0.5, Min: 0, Max: 1
func (e *GstCircleGeometricTransform) SetYCenter(value float64) error {
	return e.Element.SetProperty("y-center", value)
}

func (e *GstCircleGeometricTransform) GetYCenter() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-center")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

type GstGeometricTransform struct {
	*GstVideoFilter
}

// What to do with off edge pixels
// Default: ignore (0)
func (e *GstGeometricTransform) SetOffEdgePixels(value GstGeometricTransformOffEdgesPixelsMethod) error {
	e.Element.SetArg("off-edge-pixels", string(value))
	return nil
}

func (e *GstGeometricTransform) GetOffEdgePixels() (GstGeometricTransformOffEdgesPixelsMethod, error) {
	var value GstGeometricTransformOffEdgesPixelsMethod
	var ok bool
	v, err := e.Element.GetProperty("off-edge-pixels")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGeometricTransformOffEdgesPixelsMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGeometricTransformOffEdgesPixelsMethod")
	}
	return value, nil
}

type GstGeometricTransformOffEdgesPixelsMethod string

const (
	GstGeometricTransformOffEdgesPixelsMethodIgnore GstGeometricTransformOffEdgesPixelsMethod = "ignore" // Ignore
	GstGeometricTransformOffEdgesPixelsMethodClamp  GstGeometricTransformOffEdgesPixelsMethod = "clamp"  // Clamp
	GstGeometricTransformOffEdgesPixelsMethodWrap   GstGeometricTransformOffEdgesPixelsMethod = "wrap"   // Wrap
)
