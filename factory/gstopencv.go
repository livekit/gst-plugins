// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Performs template matching on videos and images, providing detected positions via bus messages.
type GstTemplateMatch struct {
	*GstOpencvVideoFilter
}

func NewTemplateMatch() (*GstTemplateMatch, error) {
	e, err := gst.NewElement("templatematch")
	if err != nil {
		return nil, err
	}
	return &GstTemplateMatch{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewTemplateMatchWithName(name string) (*GstTemplateMatch, error) {
	e, err := gst.NewElementWithName("templatematch", name)
	if err != nil {
		return nil, err
	}
	return &GstTemplateMatch{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Sets whether the detected template should be highlighted in the output
// Default: true
func (e *GstTemplateMatch) SetDisplay(value bool) error {
	return e.Element.SetProperty("display", value)
}

func (e *GstTemplateMatch) GetDisplay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Specifies the way the template must be compared with image regions. 0=SQDIFF, 1=SQDIFF_NORMED, 2=CCOR, 3=CCOR_NORMED, 4=CCOEFF, 5=CCOEFF_NORMED.
// Default: 3, Min: 0, Max: 5
func (e *GstTemplateMatch) SetMethod(value int) error {
	return e.Element.SetProperty("method", value)
}

func (e *GstTemplateMatch) GetMethod() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Filename of template image
// Default: NULL
func (e *GstTemplateMatch) SetTemplate(value string) error {
	return e.Element.SetProperty("template", value)
}

func (e *GstTemplateMatch) GetTemplate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("template")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Performs camera calibration by having it point at a chessboard pattern using upstream/downstream cameraundistort
type GstCameraCalibrate struct {
	*GstOpencvVideoFilter
}

func NewCameraCalibrate() (*GstCameraCalibrate, error) {
	e, err := gst.NewElement("cameracalibrate")
	if err != nil {
		return nil, err
	}
	return &GstCameraCalibrate{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCameraCalibrateWithName(name string) (*GstCameraCalibrate, error) {
	e, err := gst.NewElementWithName("cameracalibrate", name)
	if err != nil {
		return nil, err
	}
	return &GstCameraCalibrate{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Camera correction parameters (opaque string of serialized OpenCV objects)
// Default: NULL
func (e *GstCameraCalibrate) GetSettings() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("settings")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Show corners
// Default: true
func (e *GstCameraCalibrate) SetShowCorners(value bool) error {
	return e.Element.SetProperty("show-corners", value)
}

func (e *GstCameraCalibrate) GetShowCorners() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-corners")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The board height in number of items (e.g. number of squares for chessboard)
// Default: 6, Min: 1, Max: 2147483647
func (e *GstCameraCalibrate) SetBoardHeight(value int) error {
	return e.Element.SetProperty("board-height", value)
}

func (e *GstCameraCalibrate) GetBoardHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("board-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The board width in number of items (e.g. number of squares for chessboard)
// Default: 9, Min: 1, Max: 2147483647
func (e *GstCameraCalibrate) SetBoardWidth(value int) error {
	return e.Element.SetProperty("board-width", value)
}

func (e *GstCameraCalibrate) GetBoardWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("board-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sampling periodicity in ms
// Default: 350, Min: 0, Max: 2147483647
func (e *GstCameraCalibrate) SetDelay(value int) error {
	return e.Element.SetProperty("delay", value)
}

func (e *GstCameraCalibrate) GetDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The number of frames to use from the input for calibration
// Default: 25, Min: 1, Max: 2147483647
func (e *GstCameraCalibrate) SetFrameCount(value int) error {
	return e.Element.SetProperty("frame-count", value)
}

func (e *GstCameraCalibrate) GetFrameCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frame-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// One of the chessboard, circles, or asymmetric circle pattern
// Default: chessboard (0)
func (e *GstCameraCalibrate) SetPattern(value GstCameraCalibrationPattern) error {
	e.Element.SetArg("pattern", string(value))
	return nil
}

func (e *GstCameraCalibrate) GetPattern() (GstCameraCalibrationPattern, error) {
	var value GstCameraCalibrationPattern
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCameraCalibrationPattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCameraCalibrationPattern")
	}
	return value, nil
}

// The size of a square in your defined unit (point, millimeter, etc.)
// Default: 50, Min: 0, Max: 3.40282e+38
func (e *GstCameraCalibrate) SetSquareSize(value float32) error {
	return e.Element.SetProperty("square-size", value)
}

func (e *GstCameraCalibrate) GetSquareSize() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("square-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Use fisheye camera model for calibration
// Default: false
func (e *GstCameraCalibrate) SetUseFisheye(value bool) error {
	return e.Element.SetProperty("use-fisheye", value)
}

func (e *GstCameraCalibrate) GetUseFisheye() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-fisheye")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Assume zero tangential distortion
// Default: false
func (e *GstCameraCalibrate) SetZeroTangentDistorsion(value bool) error {
	return e.Element.SetProperty("zero-tangent-distorsion", value)
}

func (e *GstCameraCalibrate) GetZeroTangentDistorsion() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("zero-tangent-distorsion")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The aspect ratio
// Default: 1, Min: 0, Max: 3.40282e+38
func (e *GstCameraCalibrate) SetAspectRatio(value float32) error {
	return e.Element.SetProperty("aspect-ratio", value)
}

func (e *GstCameraCalibrate) GetAspectRatio() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Fix the principal point at the center
// Default: false
func (e *GstCameraCalibrate) SetCenterPrincipalPoint(value bool) error {
	return e.Element.SetProperty("center-principal-point", value)
}

func (e *GstCameraCalibrate) GetCenterPrincipalPoint() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("center-principal-point")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Improve corner detection accuracy for chessboard
// Default: true
func (e *GstCameraCalibrate) SetCornerSubPixel(value bool) error {
	return e.Element.SetProperty("corner-sub-pixel", value)
}

func (e *GstCameraCalibrate) GetCornerSubPixel() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("corner-sub-pixel")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Applies cvEqualizeHist OpenCV function to the image
type GstCvEqualizeHist struct {
	*GstOpencvVideoFilter
}

func NewCvEqualizeHist() (*GstCvEqualizeHist, error) {
	e, err := gst.NewElement("cvequalizehist")
	if err != nil {
		return nil, err
	}
	return &GstCvEqualizeHist{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCvEqualizeHistWithName(name string) (*GstCvEqualizeHist, error) {
	e, err := gst.NewElementWithName("cvequalizehist", name)
	if err != nil {
		return nil, err
	}
	return &GstCvEqualizeHist{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Applies cvLaplace OpenCV function to the image
type GstCvLaplace struct {
	*GstOpencvVideoFilter
}

func NewCvLaplace() (*GstCvLaplace, error) {
	e, err := gst.NewElement("cvlaplace")
	if err != nil {
		return nil, err
	}
	return &GstCvLaplace{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCvLaplaceWithName(name string) (*GstCvLaplace, error) {
	e, err := gst.NewElementWithName("cvlaplace", name)
	if err != nil {
		return nil, err
	}
	return &GstCvLaplace{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Size of the extended Laplace Kernel (1, 3, 5 or 7)
// Default: 3, Min: 1, Max: 7
func (e *GstCvLaplace) SetApertureSize(value int) error {
	return e.Element.SetProperty("aperture-size", value)
}

func (e *GstCvLaplace) GetApertureSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("aperture-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sets whether the detected edges should be used as a mask on the original input or not
// Default: true
func (e *GstCvLaplace) SetMask(value bool) error {
	return e.Element.SetProperty("mask", value)
}

func (e *GstCvLaplace) GetMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Scale factor
// Default: 1, Min: 0, Max: 1.79769e+308
func (e *GstCvLaplace) SetScale(value float64) error {
	return e.Element.SetProperty("scale", value)
}

func (e *GstCvLaplace) GetScale() (float64, error) {
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

// Value added to the scaled source array elements
// Default: 0, Min: 0, Max: 1.79769e+308
func (e *GstCvLaplace) SetShift(value float64) error {
	return e.Element.SetProperty("shift", value)
}

func (e *GstCvLaplace) GetShift() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shift")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Applies cvSobel OpenCV function to the image
type GstCvSobel struct {
	*GstOpencvVideoFilter
}

func NewCvSobel() (*GstCvSobel, error) {
	e, err := gst.NewElement("cvsobel")
	if err != nil {
		return nil, err
	}
	return &GstCvSobel{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCvSobelWithName(name string) (*GstCvSobel, error) {
	e, err := gst.NewElementWithName("cvsobel", name)
	if err != nil {
		return nil, err
	}
	return &GstCvSobel{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Size of the extended Sobel Kernel (1, 3, 5 or 7)
// Default: 3, Min: 1, Max: 7
func (e *GstCvSobel) SetApertureSize(value int) error {
	return e.Element.SetProperty("aperture-size", value)
}

func (e *GstCvSobel) GetApertureSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("aperture-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sets whether the detected derivative edges should be used as a mask on the original input or not
// Default: true
func (e *GstCvSobel) SetMask(value bool) error {
	return e.Element.SetProperty("mask", value)
}

func (e *GstCvSobel) GetMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Order of the derivative x
// Default: 1, Min: -1, Max: 2147483647
func (e *GstCvSobel) SetXOrder(value int) error {
	return e.Element.SetProperty("x-order", value)
}

func (e *GstCvSobel) GetXOrder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("x-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Order of the derivative y
// Default: 0, Min: -1, Max: 2147483647
func (e *GstCvSobel) SetYOrder(value int) error {
	return e.Element.SetProperty("y-order", value)
}

func (e *GstCvSobel) GetYOrder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("y-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Multiscale retinex for colour image enhancement
type GstRetinex struct {
	*GstOpencvVideoFilter
}

func NewRetinex() (*GstRetinex, error) {
	e, err := gst.NewElement("retinex")
	if err != nil {
		return nil, err
	}
	return &GstRetinex{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewRetinexWithName(name string) (*GstRetinex, error) {
	e, err := gst.NewElementWithName("retinex", name)
	if err != nil {
		return nil, err
	}
	return &GstRetinex{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Gain
// Default: 128, Min: 0, Max: 2147483647
func (e *GstRetinex) SetGain(value int) error {
	return e.Element.SetProperty("gain", value)
}

func (e *GstRetinex) GetGain() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Retinex method to use
// Default: basic (0)
func (e *GstRetinex) SetMethod(value GstRetinexMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstRetinex) GetMethod() (GstRetinexMethod, error) {
	var value GstRetinexMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRetinexMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRetinexMethod")
	}
	return value, nil
}

// Offset
// Default: 128, Min: 0, Max: 2147483647
func (e *GstRetinex) SetOffset(value int) error {
	return e.Element.SetProperty("offset", value)
}

func (e *GstRetinex) GetOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Amount of gaussian filters (scales) used in multiscale retinex
// Default: 3, Min: 1, Max: 4
func (e *GstRetinex) SetScales(value int) error {
	return e.Element.SetProperty("scales", value)
}

func (e *GstRetinex) GetScales() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("scales")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sigma
// Default: 14, Min: 0, Max: 1.79769e+308
func (e *GstRetinex) SetSigma(value float64) error {
	return e.Element.SetProperty("sigma", value)
}

func (e *GstRetinex) GetSigma() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sigma")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Performs camera undistort
type GstCameraUndistort struct {
	*GstOpencvVideoFilter
}

func NewCameraUndistort() (*GstCameraUndistort, error) {
	e, err := gst.NewElement("cameraundistort")
	if err != nil {
		return nil, err
	}
	return &GstCameraUndistort{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCameraUndistortWithName(name string) (*GstCameraUndistort, error) {
	e, err := gst.NewElementWithName("cameraundistort", name)
	if err != nil {
		return nil, err
	}
	return &GstCameraUndistort{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Apply camera corrections
// Default: true
func (e *GstCameraUndistort) SetUndistort(value bool) error {
	return e.Element.SetProperty("undistort", value)
}

func (e *GstCameraUndistort) GetUndistort() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("undistort")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Show all pixels (1), only valid ones (0) or something in between
// Default: 0, Min: 0, Max: 1
func (e *GstCameraUndistort) SetAlpha(value float32) error {
	return e.Element.SetProperty("alpha", value)
}

func (e *GstCameraUndistort) GetAlpha() (float32, error) {
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

// Camera correction parameters (opaque string of serialized OpenCV objects)
// Default: NULL
func (e *GstCameraUndistort) SetSettings(value string) error {
	return e.Element.SetProperty("settings", value)
}

func (e *GstCameraUndistort) GetSettings() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("settings")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Applies cvDilate OpenCV function to the image
type GstCvDilate struct {
	*GstCvDilateErode
}

func NewCvDilate() (*GstCvDilate, error) {
	e, err := gst.NewElement("cvdilate")
	if err != nil {
		return nil, err
	}
	return &GstCvDilate{GstCvDilateErode: &GstCvDilateErode{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewCvDilateWithName(name string) (*GstCvDilate, error) {
	e, err := gst.NewElementWithName("cvdilate", name)
	if err != nil {
		return nil, err
	}
	return &GstCvDilate{GstCvDilateErode: &GstCvDilateErode{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Performs face detection on videos and images, providing detected positions via bus messages
type GstFaceDetect struct {
	*GstOpencvVideoFilter
}

func NewFaceDetect() (*GstFaceDetect, error) {
	e, err := gst.NewElement("facedetect")
	if err != nil {
		return nil, err
	}
	return &GstFaceDetect{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewFaceDetectWithName(name string) (*GstFaceDetect, error) {
	e, err := gst.NewElementWithName("facedetect", name)
	if err != nil {
		return nil, err
	}
	return &GstFaceDetect{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// When send update bus messages, if at all
// Default: every_frame (0)
func (e *GstFaceDetect) SetUpdates(value GstFaceDetectUpdates) error {
	e.Element.SetArg("updates", string(value))
	return nil
}

func (e *GstFaceDetect) GetUpdates() (GstFaceDetectUpdates, error) {
	var value GstFaceDetectUpdates
	var ok bool
	v, err := e.Element.GetProperty("updates")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFaceDetectUpdates)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFaceDetectUpdates")
	}
	return value, nil
}

// Sets whether the detected faces should be highlighted in the output
// Default: true
func (e *GstFaceDetect) SetDisplay(value bool) error {
	return e.Element.SetProperty("display", value)
}

func (e *GstFaceDetect) GetDisplay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimum area width to be recognized as a face
// Default: 30, Min: 0, Max: 2147483647
func (e *GstFaceDetect) SetMinSizeWidth(value int) error {
	return e.Element.SetProperty("min-size-width", value)
}

func (e *GstFaceDetect) GetMinSizeWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-size-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Location of Haar cascade file to use for mouth detection
// Default: /usr/share/OpenCV/haarcascades/haarcascade_mcs_mouth.xml
func (e *GstFaceDetect) SetMouthProfile(value string) error {
	return e.Element.SetProperty("mouth-profile", value)
}

func (e *GstFaceDetect) GetMouthProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mouth-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Location of Haar cascade file to use for face detection
// Default: /usr/share/OpenCV/haarcascades/haarcascade_frontalface_default.xml
func (e *GstFaceDetect) SetProfile(value string) error {
	return e.Element.SetProperty("profile", value)
}

func (e *GstFaceDetect) GetProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Minimum image average standard deviation: on images with standard deviation lesser than this value facedetection will not be performed. Setting this property help to save cpu and reduce false positives not performing face detection on images with little changes
// Default: 0, Min: 0, Max: 255
func (e *GstFaceDetect) SetMinStddev(value int) error {
	return e.Element.SetProperty("min-stddev", value)
}

func (e *GstFaceDetect) GetMinStddev() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-stddev")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Location of Haar cascade file to use for nose detection
// Default: /usr/share/OpenCV/haarcascades/haarcascade_mcs_nose.xml
func (e *GstFaceDetect) SetNoseProfile(value string) error {
	return e.Element.SetProperty("nose-profile", value)
}

func (e *GstFaceDetect) GetNoseProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("nose-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Factor by which the frame is scaled after each object scan
// Default: 1.25, Min: 1.1, Max: 10
func (e *GstFaceDetect) SetScaleFactor(value float64) error {
	return e.Element.SetProperty("scale-factor", value)
}

func (e *GstFaceDetect) GetScaleFactor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Location of Haar cascade file to use for eye-pair detection
// Default: /usr/share/OpenCV/haarcascades/haarcascade_mcs_eyepair_small.xml
func (e *GstFaceDetect) SetEyesProfile(value string) error {
	return e.Element.SetProperty("eyes-profile", value)
}

func (e *GstFaceDetect) GetEyesProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("eyes-profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Flags to cvHaarDetectObjects
// Default: do-canny-pruning
func (e *GstFaceDetect) SetFlags(value GstOpencvFaceDetectFlags) error {
	e.Element.SetArg("flags", fmt.Sprint(value))
	return nil
}

func (e *GstFaceDetect) GetFlags() (GstOpencvFaceDetectFlags, error) {
	var value GstOpencvFaceDetectFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpencvFaceDetectFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpencvFaceDetectFlags")
	}
	return value, nil
}

// Minimum number (minus 1) of neighbor rectangles that makes up an object
// Default: 3, Min: 0, Max: 2147483647
func (e *GstFaceDetect) SetMinNeighbors(value int) error {
	return e.Element.SetProperty("min-neighbors", value)
}

func (e *GstFaceDetect) GetMinNeighbors() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-neighbors")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum area height to be recognized as a face
// Default: 30, Min: 0, Max: 2147483647
func (e *GstFaceDetect) SetMinSizeHeight(value int) error {
	return e.Element.SetProperty("min-size-height", value)
}

func (e *GstFaceDetect) GetMinSizeHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-size-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Runs Grabcut algorithm on input alpha. Values: BG=0, FG=1, PR_BG=2, PR_FGD=3; NOTE: larger values of alpha (notably 255) are interpreted as PR_FGD too.
// IN CASE OF no alpha mask input (all 0's or all 1's), the 'face' downstream event is used to create a bbox of PR_FG elements.
// IF nothing is present, then nothing is done.
type GstGrabcut struct {
	*GstOpencvVideoFilter
}

func NewGrabcut() (*GstGrabcut, error) {
	e, err := gst.NewElement("grabcut")
	if err != nil {
		return nil, err
	}
	return &GstGrabcut{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewGrabcutWithName(name string) (*GstGrabcut, error) {
	e, err := gst.NewElementWithName("grabcut", name)
	if err != nil {
		return nil, err
	}
	return &GstGrabcut{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Grow factor for the face bounding box, if present
// Default: 1.6, Min: 1, Max: 4
func (e *GstGrabcut) SetScale(value float32) error {
	return e.Element.SetProperty("scale", value)
}

func (e *GstGrabcut) GetScale() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// If true, the output RGB is overwritten with the segmented foreground. Alpha channel same as normal case
// Default: false
func (e *GstGrabcut) SetTestMode(value bool) error {
	return e.Element.SetProperty("test-mode", value)
}

func (e *GstGrabcut) GetTestMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("test-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Write text on the top of video
type GstOpencvTextOverlay struct {
	*GstOpencvVideoFilter
}

func NewOpencvTextOverlay() (*GstOpencvTextOverlay, error) {
	e, err := gst.NewElement("opencvtextoverlay")
	if err != nil {
		return nil, err
	}
	return &GstOpencvTextOverlay{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewOpencvTextOverlayWithName(name string) (*GstOpencvTextOverlay, error) {
	e, err := gst.NewElementWithName("opencvtextoverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstOpencvTextOverlay{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Sets the color -G
// Default: 0, Min: 0, Max: 255
func (e *GstOpencvTextOverlay) SetColorG(value int) error {
	return e.Element.SetProperty("colorG", value)
}

func (e *GstOpencvTextOverlay) GetColorG() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("colorG")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Text to be display.
// Default: Opencv Text Overlay
func (e *GstOpencvTextOverlay) SetText(value string) error {
	return e.Element.SetProperty("text", value)
}

func (e *GstOpencvTextOverlay) GetText() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("text")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Sets the Thickness of Font
// Default: 2, Min: 0, Max: 2147483647
func (e *GstOpencvTextOverlay) SetThickness(value int) error {
	return e.Element.SetProperty("thickness", value)
}

func (e *GstOpencvTextOverlay) GetThickness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("thickness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sets the Vertical position
// Default: 50, Min: 0, Max: 2147483647
func (e *GstOpencvTextOverlay) SetYpos(value int) error {
	return e.Element.SetProperty("ypos", value)
}

func (e *GstOpencvTextOverlay) GetYpos() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ypos")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sets the color -B
// Default: 0, Min: 0, Max: 255
func (e *GstOpencvTextOverlay) SetColorB(value int) error {
	return e.Element.SetProperty("colorB", value)
}

func (e *GstOpencvTextOverlay) GetColorB() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("colorB")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sets the color -R
// Default: 0, Min: 0, Max: 255
func (e *GstOpencvTextOverlay) SetColorR(value int) error {
	return e.Element.SetProperty("colorR", value)
}

func (e *GstOpencvTextOverlay) GetColorR() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("colorR")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sets the height of fonts
// Default: 1, Min: 1, Max: 5
func (e *GstOpencvTextOverlay) SetHeight(value float64) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstOpencvTextOverlay) GetHeight() (float64, error) {
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

// Sets the width of fonts
// Default: 1, Min: 1, Max: 5
func (e *GstOpencvTextOverlay) SetWidth(value float64) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstOpencvTextOverlay) GetWidth() (float64, error) {
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

// Sets the Horizontal position
// Default: 50, Min: 0, Max: 2147483647
func (e *GstOpencvTextOverlay) SetXpos(value int) error {
	return e.Element.SetProperty("xpos", value)
}

func (e *GstOpencvTextOverlay) GetXpos() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("xpos")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Performs object tracking on videos and stores it in video buffer metadata.
type GstCVTracker struct {
	*GstOpencvVideoFilter
}

func NewCVTracker() (*GstCVTracker, error) {
	e, err := gst.NewElement("cvtracker")
	if err != nil {
		return nil, err
	}
	return &GstCVTracker{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCVTrackerWithName(name string) (*GstCVTracker, error) {
	e, err := gst.NewElementWithName("cvtracker", name)
	if err != nil {
		return nil, err
	}
	return &GstCVTracker{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Algorithm for tracking objects
// Default: MedianFlow (3)
func (e *GstCVTracker) SetAlgorithm(value GstOpenCVTrackerAlgorithm) error {
	e.Element.SetArg("algorithm", string(value))
	return nil
}

func (e *GstCVTracker) GetAlgorithm() (GstOpenCVTrackerAlgorithm, error) {
	var value GstOpenCVTrackerAlgorithm
	var ok bool
	v, err := e.Element.GetProperty("algorithm")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpenCVTrackerAlgorithm)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpenCVTrackerAlgorithm")
	}
	return value, nil
}

// Draw rectangle around tracked object
// Default: true
func (e *GstCVTracker) SetDrawRect(value bool) error {
	return e.Element.SetProperty("draw-rect", value)
}

func (e *GstCVTracker) GetDrawRect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-rect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Track object box's initial height
// Default: 50, Min: 0, Max: -1
func (e *GstCVTracker) SetObjectInitialHeight(value uint) error {
	return e.Element.SetProperty("object-initial-height", value)
}

func (e *GstCVTracker) GetObjectInitialHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("object-initial-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Track object box's initial width
// Default: 50, Min: 0, Max: -1
func (e *GstCVTracker) SetObjectInitialWidth(value uint) error {
	return e.Element.SetProperty("object-initial-width", value)
}

func (e *GstCVTracker) GetObjectInitialWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("object-initial-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Track object box's initial X coordinate
// Default: 50, Min: 0, Max: -1
func (e *GstCVTracker) SetObjectInitialX(value uint) error {
	return e.Element.SetProperty("object-initial-x", value)
}

func (e *GstCVTracker) GetObjectInitialX() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("object-initial-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Track object box's initial Y coordinate
// Default: 50, Min: 0, Max: -1
func (e *GstCVTracker) SetObjectInitialY(value uint) error {
	return e.Element.SetProperty("object-initial-y", value)
}

func (e *GstCVTracker) GetObjectInitialY() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("object-initial-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Calculates the stereo disparity map from two (sequences of) rectified and aligned stereo images
type GstDisparity struct {
	*gst.Element
}

func NewDisparity() (*GstDisparity, error) {
	e, err := gst.NewElement("disparity")
	if err != nil {
		return nil, err
	}
	return &GstDisparity{Element: e}, nil
}

func NewDisparityWithName(name string) (*GstDisparity, error) {
	e, err := gst.NewElementWithName("disparity", name)
	if err != nil {
		return nil, err
	}
	return &GstDisparity{Element: e}, nil
}

// Stereo matching method to use
// Default: sgbm (1)
func (e *GstDisparity) SetMethod(value GstDisparityMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstDisparity) GetMethod() (GstDisparityMethod, error) {
	var value GstDisparityMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDisparityMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDisparityMethod")
	}
	return value, nil
}

// Performs hand gesture detection on videos, providing detected hand positions via bus message and navigation event, and deals with hand gesture events
type GstHanddetect struct {
	*GstOpencvVideoFilter
}

func NewHanddetect() (*GstHanddetect, error) {
	e, err := gst.NewElement("handdetect")
	if err != nil {
		return nil, err
	}
	return &GstHanddetect{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewHanddetectWithName(name string) (*GstHanddetect, error) {
	e, err := gst.NewElementWithName("handdetect", name)
	if err != nil {
		return nil, err
	}
	return &GstHanddetect{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Location of HAAR cascade file (palm gesture)
// Default: @0@/palm.xml
func (e *GstHanddetect) SetProfilePalm(value string) error {
	return e.Element.SetProperty("profile-palm", value)
}

func (e *GstHanddetect) GetProfilePalm() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("profile-palm")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// HEIGHT of left-top pointer in region of interest
// Gestures in the defined region of interest will emit messages
// Default: 0, Min: 0, Max: 2147483647
func (e *GstHanddetect) SetROIHEIGHT(value int) error {
	return e.Element.SetProperty("ROI-HEIGHT", value)
}

func (e *GstHanddetect) GetROIHEIGHT() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ROI-HEIGHT")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// WIDTH of left-top pointer in region of interest
// Gestures in the defined region of interest will emit messages
// Default: 0, Min: 0, Max: 2147483647
func (e *GstHanddetect) SetROIWIDTH(value int) error {
	return e.Element.SetProperty("ROI-WIDTH", value)
}

func (e *GstHanddetect) GetROIWIDTH() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ROI-WIDTH")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// X of left-top pointer in region of interest
// Gestures in the defined region of interest will emit messages
// Default: 0, Min: 0, Max: 2147483647
func (e *GstHanddetect) SetROIX(value int) error {
	return e.Element.SetProperty("ROI-X", value)
}

func (e *GstHanddetect) GetROIX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ROI-X")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Y of left-top pointer in region of interest
// Gestures in the defined region of interest will emit messages
// Default: 0, Min: 0, Max: 2147483647
func (e *GstHanddetect) SetROIY(value int) error {
	return e.Element.SetProperty("ROI-Y", value)
}

func (e *GstHanddetect) GetROIY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ROI-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Whether the detected hands are highlighted in output frame
// Default: true
func (e *GstHanddetect) SetDisplay(value bool) error {
	return e.Element.SetProperty("display", value)
}

func (e *GstHanddetect) GetDisplay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Location of HAAR cascade file (fist gesture)
// Default: @0@/fist.xml
func (e *GstHanddetect) SetProfileFist(value string) error {
	return e.Element.SetProperty("profile-fist", value)
}

func (e *GstHanddetect) GetProfileFist() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("profile-fist")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Create a Foregound/Background mask applying a particular algorithm
type GstSegmentation struct {
	*GstOpencvVideoFilter
}

func NewSegmentation() (*GstSegmentation, error) {
	e, err := gst.NewElement("segmentation")
	if err != nil {
		return nil, err
	}
	return &GstSegmentation{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewSegmentationWithName(name string) (*GstSegmentation, error) {
	e, err := gst.NewElementWithName("segmentation", name)
	if err != nil {
		return nil, err
	}
	return &GstSegmentation{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Speed with which a motionless foreground pixel would become background (inverse of number of frames)
// Default: 0.01, Min: 0, Max: 1
func (e *GstSegmentation) SetLearningRate(value float32) error {
	return e.Element.SetProperty("learning-rate", value)
}

func (e *GstSegmentation) GetLearningRate() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("learning-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Segmentation method to use
// Default: mog2 (2)
func (e *GstSegmentation) SetMethod(value GstSegmentationMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstSegmentation) GetMethod() (GstSegmentationMethod, error) {
	var value GstSegmentationMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSegmentationMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSegmentationMethod")
	}
	return value, nil
}

// If true, the output RGB is overwritten with the calculated foreground (white color)
// Default: false
func (e *GstSegmentation) SetTestMode(value bool) error {
	return e.Element.SetProperty("test-mode", value)
}

func (e *GstSegmentation) GetTestMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("test-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Performs non-parametric skin detection on input
type GstSkinDetect struct {
	*GstOpencvVideoFilter
}

func NewSkinDetect() (*GstSkinDetect, error) {
	e, err := gst.NewElement("skindetect")
	if err != nil {
		return nil, err
	}
	return &GstSkinDetect{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewSkinDetectWithName(name string) (*GstSkinDetect, error) {
	e, err := gst.NewElementWithName("skindetect", name)
	if err != nil {
		return nil, err
	}
	return &GstSkinDetect{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Method to use
// Default: hsv (0)
func (e *GstSkinDetect) SetMethod(value GstSkindetectMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstSkinDetect) GetMethod() (GstSkindetectMethod, error) {
	var value GstSkindetectMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSkindetectMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSkindetectMethod")
	}
	return value, nil
}

// Apply opening-closing to skin detection to extract large, significant blobs
// Default: true
func (e *GstSkinDetect) SetPostprocess(value bool) error {
	return e.Element.SetProperty("postprocess", value)
}

func (e *GstSkinDetect) GetPostprocess() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("postprocess")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Performs motion detection on videos and images, providing detected motion cells index via bus messages
type GstMotioncells struct {
	*GstOpencvVideoFilter
}

func NewMotioncells() (*GstMotioncells, error) {
	e, err := gst.NewElement("motioncells")
	if err != nil {
		return nil, err
	}
	return &GstMotioncells{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewMotioncellsWithName(name string) (*GstMotioncells, error) {
	e, err := gst.NewElementWithName("motioncells", name)
	if err != nil {
		return nil, err
	}
	return &GstMotioncells{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Number of horizontal grid cells.
// Default: 10, Min: 8, Max: 32
func (e *GstMotioncells) SetGridx(value int) error {
	return e.Element.SetProperty("gridx", value)
}

func (e *GstMotioncells) GetGridx() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gridx")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Describe a region with its upper left and lower right x, y coordinates separated with ":". Pass multiple regions as a comma-separated list

func (e *GstMotioncells) SetMotionmaskcoords(value string) error {
	return e.Element.SetProperty("motionmaskcoords", value)
}

func (e *GstMotioncells) GetMotionmaskcoords() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("motionmaskcoords")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Post bus messages for every motion frame or just motion start and motion stop
// Default: false
func (e *GstMotioncells) SetPostallmotion(value bool) error {
	return e.Element.SetProperty("postallmotion", value)
}

func (e *GstMotioncells) GetPostallmotion() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("postallmotion")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// If non 0, post a no_motion event on the bus if no motion is detected for the given number of seconds
// Default: 0, Min: 0, Max: 180
func (e *GstMotioncells) SetPostnomotion(value int) error {
	return e.Element.SetProperty("postnomotion", value)
}

func (e *GstMotioncells) GetPostnomotion() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("postnomotion")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Location of motioncells data file (empty string means no saving)
// Default: NULL
func (e *GstMotioncells) SetDatafile(value string) error {
	return e.Element.SetProperty("datafile", value)
}

func (e *GstMotioncells) GetDatafile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("datafile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Minimum number of motion frames triggering a motion event
// Default: 1, Min: 1, Max: 60
func (e *GstMotioncells) SetMinimummotionframes(value int) error {
	return e.Element.SetProperty("minimummotionframes", value)
}

func (e *GstMotioncells) GetMinimummotionframes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("minimummotionframes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Describe a cell with its line and column idx separated with ":". Pass multiple cells as a comma-separated list

func (e *GstMotioncells) SetMotioncellsidx(value string) error {
	return e.Element.SetProperty("motioncellsidx", value)
}

func (e *GstMotioncells) GetMotioncellsidx() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("motioncellsidx")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Describe a cell with its line and column idx separated with ":". Pass multiple cells as a comma-separated list

func (e *GstMotioncells) SetMotionmaskcellspos(value string) error {
	return e.Element.SetProperty("motionmaskcellspos", value)
}

func (e *GstMotioncells) GetMotionmaskcellspos() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("motionmaskcellspos")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Toggle usage of alpha blending on frames with motion cells
// Default: true
func (e *GstMotioncells) SetUsealpha(value bool) error {
	return e.Element.SetProperty("usealpha", value)
}

func (e *GstMotioncells) GetUsealpha() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("usealpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Color for motion cells in R,G,B format. Max per channel is 255
// Default: 255,255,0
func (e *GstMotioncells) SetCellscolor(value string) error {
	return e.Element.SetProperty("cellscolor", value)
}

func (e *GstMotioncells) GetCellscolor() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("cellscolor")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Toggle display of motion cells on current frame
// Default: true
func (e *GstMotioncells) SetDisplay(value bool) error {
	return e.Element.SetProperty("display", value)
}

func (e *GstMotioncells) GetDisplay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Interval in seconds after which motion is considered finished and a motion finished bus message is posted.
// Default: 5, Min: 1, Max: 60
func (e *GstMotioncells) SetGap(value int) error {
	return e.Element.SetProperty("gap", value)
}

func (e *GstMotioncells) GetGap() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gap")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Motion detection sensitivity.
// Default: 0.5, Min: 0, Max: 1
func (e *GstMotioncells) SetSensitivity(value float64) error {
	return e.Element.SetProperty("sensitivity", value)
}

func (e *GstMotioncells) GetSensitivity() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sensitivity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Threshold value for motion. Filter detects motion when at least this fraction of the cells have moved
// Default: 0.01, Min: 0, Max: 1
func (e *GstMotioncells) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstMotioncells) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Toggles motion calculation. If FALSE, this filter does nothing
// Default: true
func (e *GstMotioncells) SetCalculatemotion(value bool) error {
	return e.Element.SetProperty("calculatemotion", value)
}

func (e *GstMotioncells) GetCalculatemotion() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("calculatemotion")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extension of datafile
// Default: vamc
func (e *GstMotioncells) SetDatafileextension(value string) error {
	return e.Element.SetProperty("datafileextension", value)
}

func (e *GstMotioncells) GetDatafileextension() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("datafileextension")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Number of vertical grid cells.
// Default: 10, Min: 8, Max: 32
func (e *GstMotioncells) SetGridy(value int) error {
	return e.Element.SetProperty("gridy", value)
}

func (e *GstMotioncells) GetGridy() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("gridy")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Motion Cell Border Thickness. Set to -1 to fill motion cell
// Default: 1, Min: -1, Max: 5
func (e *GstMotioncells) SetMotioncellthickness(value int) error {
	return e.Element.SetProperty("motioncellthickness", value)
}

func (e *GstMotioncells) GetMotioncellthickness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("motioncellthickness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Applies cvErode OpenCV function to the image
type GstCvErode struct {
	*GstCvDilateErode
}

func NewCvErode() (*GstCvErode, error) {
	e, err := gst.NewElement("cverode")
	if err != nil {
		return nil, err
	}
	return &GstCvErode{GstCvDilateErode: &GstCvDilateErode{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

func NewCvErodeWithName(name string) (*GstCvErode, error) {
	e, err := gst.NewElementWithName("cverode", name)
	if err != nil {
		return nil, err
	}
	return &GstCvErode{GstCvDilateErode: &GstCvDilateErode{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}}, nil
}

// Applies cvSmooth OpenCV function to the image
type GstCvSmooth struct {
	*GstOpencvVideoFilter
}

func NewCvSmooth() (*GstCvSmooth, error) {
	e, err := gst.NewElement("cvsmooth")
	if err != nil {
		return nil, err
	}
	return &GstCvSmooth{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewCvSmoothWithName(name string) (*GstCvSmooth, error) {
	e, err := gst.NewElementWithName("cvsmooth", name)
	if err != nil {
		return nil, err
	}
	return &GstCvSmooth{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// If type is gaussian, this means the standard deviation.If type is bilateral, this means the color-sigma. If zero, Default values are used.
// Default: 0, Min: 0, Max: 1.79769e+308
func (e *GstCvSmooth) SetColor(value float64) error {
	return e.Element.SetProperty("color", value)
}

func (e *GstCvSmooth) GetColor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("color")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Height of the area to blur (in pixels).
// Default: 2147483647, Min: 0, Max: 2147483647
func (e *GstCvSmooth) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

func (e *GstCvSmooth) GetHeight() (int, error) {
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

// The gaussian kernel height (must be positive and odd).
// Default: 3, Min: 0, Max: 2147483647
func (e *GstCvSmooth) SetKernelHeight(value int) error {
	return e.Element.SetProperty("kernel-height", value)
}

func (e *GstCvSmooth) GetKernelHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kernel-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Starting x position for blur (in pixels).
// Default: 0, Min: 0, Max: 2147483647
func (e *GstCvSmooth) SetPositionX(value int) error {
	return e.Element.SetProperty("position-x", value)
}

func (e *GstCvSmooth) GetPositionX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("position-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Starting y position for blur (in pixels).
// Default: 0, Min: 0, Max: 2147483647
func (e *GstCvSmooth) SetPositionY(value int) error {
	return e.Element.SetProperty("position-y", value)
}

func (e *GstCvSmooth) GetPositionY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("position-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Only used in bilateral type, means the spatial-sigma.
// Default: 0, Min: 0, Max: 1.79769e+308
func (e *GstCvSmooth) SetSpatial(value float64) error {
	return e.Element.SetProperty("spatial", value)
}

func (e *GstCvSmooth) GetSpatial() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("spatial")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Smooth Type
// Default: gaussian (2)
func (e *GstCvSmooth) SetType(value GstCvSmoothTypeType) error {
	e.Element.SetArg("type", string(value))
	return nil
}

func (e *GstCvSmooth) GetType() (GstCvSmoothTypeType, error) {
	var value GstCvSmoothTypeType
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCvSmoothTypeType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCvSmoothTypeType")
	}
	return value, nil
}

// The gaussian kernel width (must be positive and odd).If type is median, this means the aperture linear size.Check OpenCV docs: http://docs.opencv.org/2.4/modules/imgproc/doc/filtering.htm
// Default: 3, Min: 1, Max: 2147483647
func (e *GstCvSmooth) SetKernelWidth(value int) error {
	return e.Element.SetProperty("kernel-width", value)
}

func (e *GstCvSmooth) GetKernelWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("kernel-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Width of the area to blur (in pixels).
// Default: 2147483647, Min: 0, Max: 2147483647
func (e *GstCvSmooth) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

func (e *GstCvSmooth) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Dewarp fisheye images
type GstDewarp struct {
	*GstOpencvVideoFilter
}

func NewDewarp() (*GstDewarp, error) {
	e, err := gst.NewElement("dewarp")
	if err != nil {
		return nil, err
	}
	return &GstDewarp{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewDewarpWithName(name string) (*GstDewarp, error) {
	e, err := gst.NewElementWithName("dewarp", name)
	if err != nil {
		return nil, err
	}
	return &GstDewarp{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Correction factor for remapping on x axis. A correction is needed if the fisheye image is not inside a circle
// Default: 1, Min: 0.1, Max: 10
func (e *GstDewarp) SetXRemapCorrection(value float64) error {
	return e.Element.SetProperty("x-remap-correction", value)
}

func (e *GstDewarp) GetXRemapCorrection() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("x-remap-correction")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Y axis center of the fisheye image
// Default: 0.5, Min: 0, Max: 1
func (e *GstDewarp) SetYCenter(value float64) error {
	return e.Element.SetProperty("y-center", value)
}

func (e *GstDewarp) GetYCenter() (float64, error) {
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

// Correction factor for remapping on y axis. A correction is needed if the fisheye image is not inside a circle
// Default: 1, Min: 0.1, Max: 10
func (e *GstDewarp) SetYRemapCorrection(value float64) error {
	return e.Element.SetProperty("y-remap-correction", value)
}

func (e *GstDewarp) GetYRemapCorrection() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("y-remap-correction")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// How to display the dewarped image
// Default: single-panorama (0)
func (e *GstDewarp) SetDisplayMode(value GstDewarpDisplayMode) error {
	e.Element.SetArg("display-mode", string(value))
	return nil
}

func (e *GstDewarp) GetDisplayMode() (GstDewarpDisplayMode, error) {
	var value GstDewarpDisplayMode
	var ok bool
	v, err := e.Element.GetProperty("display-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDewarpDisplayMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDewarpDisplayMode")
	}
	return value, nil
}

// Inner radius of the fisheye image donut. If outer radius <= inner radius the element will work in passthrough mode
// Default: 0, Min: 0, Max: 1
func (e *GstDewarp) SetInnerRadius(value float64) error {
	return e.Element.SetProperty("inner-radius", value)
}

func (e *GstDewarp) GetInnerRadius() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("inner-radius")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Interpolation method to use
// Default: bilinear (1)
func (e *GstDewarp) SetInterpolationMethod(value GstDewarpInterpolationMode) error {
	e.Element.SetArg("interpolation-method", string(value))
	return nil
}

func (e *GstDewarp) GetInterpolationMethod() (GstDewarpInterpolationMode, error) {
	var value GstDewarpInterpolationMode
	var ok bool
	v, err := e.Element.GetProperty("interpolation-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDewarpInterpolationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDewarpInterpolationMode")
	}
	return value, nil
}

// Outer radius of the fisheye image donut. If outer radius <= inner radius the element will work in passthrough mode
// Default: 0, Min: 0, Max: 1
func (e *GstDewarp) SetOuterRadius(value float64) error {
	return e.Element.SetProperty("outer-radius", value)
}

func (e *GstDewarp) GetOuterRadius() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("outer-radius")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// X axis center of the fisheye image
// Default: 0.5, Min: 0, Max: 1
func (e *GstDewarp) SetXCenter(value float64) error {
	return e.Element.SetProperty("x-center", value)
}

func (e *GstDewarp) GetXCenter() (float64, error) {
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

// Performs canny edge detection on videos and images.
type GstEdgeDetect struct {
	*GstOpencvVideoFilter
}

func NewEdgeDetect() (*GstEdgeDetect, error) {
	e, err := gst.NewElement("edgedetect")
	if err != nil {
		return nil, err
	}
	return &GstEdgeDetect{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewEdgeDetectWithName(name string) (*GstEdgeDetect, error) {
	e, err := gst.NewElementWithName("edgedetect", name)
	if err != nil {
		return nil, err
	}
	return &GstEdgeDetect{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Sets whether the detected edges should be used as a mask on the original input or not
// Default: true
func (e *GstEdgeDetect) SetMask(value bool) error {
	return e.Element.SetProperty("mask", value)
}

func (e *GstEdgeDetect) GetMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Threshold value for canny edge detection
// Default: 50, Min: 0, Max: 1000
func (e *GstEdgeDetect) SetThreshold1(value int) error {
	return e.Element.SetProperty("threshold1", value)
}

func (e *GstEdgeDetect) GetThreshold1() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("threshold1")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Second threshold value for canny edge detection
// Default: 150, Min: 0, Max: 1000
func (e *GstEdgeDetect) SetThreshold2(value int) error {
	return e.Element.SetProperty("threshold2", value)
}

func (e *GstEdgeDetect) GetThreshold2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("threshold2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Aperture size for Sobel operator (Must be either 3, 5 or 7
// Default: 3, Min: 3, Max: 7
func (e *GstEdgeDetect) SetAperture(value int) error {
	return e.Element.SetProperty("aperture", value)
}

func (e *GstEdgeDetect) GetAperture() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("aperture")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Blurs faces in images and videos
type GstFaceBlur struct {
	*GstOpencvVideoFilter
}

func NewFaceBlur() (*GstFaceBlur, error) {
	e, err := gst.NewElement("faceblur")
	if err != nil {
		return nil, err
	}
	return &GstFaceBlur{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

func NewFaceBlurWithName(name string) (*GstFaceBlur, error) {
	e, err := gst.NewElementWithName("faceblur", name)
	if err != nil {
		return nil, err
	}
	return &GstFaceBlur{GstOpencvVideoFilter: &GstOpencvVideoFilter{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}}, nil
}

// Minimum window height size
// Default: 30, Min: 0, Max: 2147483647
func (e *GstFaceBlur) SetMinSizeHeight(value int) error {
	return e.Element.SetProperty("min-size-height", value)
}

func (e *GstFaceBlur) GetMinSizeHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-size-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum window width size
// Default: 30, Min: 0, Max: 2147483647
func (e *GstFaceBlur) SetMinSizeWidth(value int) error {
	return e.Element.SetProperty("min-size-width", value)
}

func (e *GstFaceBlur) GetMinSizeWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-size-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Location of Haar cascade file to use for face blurion
// Default: /usr/share/OpenCV/haarcascades/haarcascade_frontalface_default.xml
func (e *GstFaceBlur) SetProfile(value string) error {
	return e.Element.SetProperty("profile", value)
}

func (e *GstFaceBlur) GetProfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Factor by which the windows is scaled after each scan
// Default: 1.25, Min: 1.1, Max: 10
func (e *GstFaceBlur) SetScaleFactor(value float64) error {
	return e.Element.SetProperty("scale-factor", value)
}

func (e *GstFaceBlur) GetScaleFactor() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Flags to cvHaarDetectObjects
// Default: do-canny-pruning
func (e *GstFaceBlur) SetFlags(value GstOpencvFaceBlurFlags) error {
	e.Element.SetArg("flags", fmt.Sprint(value))
	return nil
}

func (e *GstFaceBlur) GetFlags() (GstOpencvFaceBlurFlags, error) {
	var value GstOpencvFaceBlurFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpencvFaceBlurFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpencvFaceBlurFlags")
	}
	return value, nil
}

// Minimum number (minus 1) of neighbor rectangles that makes up an object
// Default: 3, Min: 0, Max: 2147483647
func (e *GstFaceBlur) SetMinNeighbors(value int) error {
	return e.Element.SetProperty("min-neighbors", value)
}

func (e *GstFaceBlur) GetMinNeighbors() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-neighbors")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstCameraCalibrationPattern string

const (
	GstCameraCalibrationPatternChessboard            GstCameraCalibrationPattern = "chessboard"              // Chessboard
	GstCameraCalibrationPatternCircleGrids           GstCameraCalibrationPattern = "circle_grids"            // Circle Grids
	GstCameraCalibrationPatternAsymmetricCircleGrids GstCameraCalibrationPattern = "asymmetric_circle_grids" // Asymmetric Circle Grids
)

type GstCvDilateErode struct {
	*GstOpencvVideoFilter
}

// Number of iterations to run the algorithm
// Default: 1, Min: 1, Max: 2147483647
func (e *GstCvDilateErode) SetIterations(value int) error {
	return e.Element.SetProperty("iterations", value)
}

func (e *GstCvDilateErode) GetIterations() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("iterations")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstDewarpDisplayMode string

const (
	GstDewarpDisplayModeSinglePanorama GstDewarpDisplayMode = "single-panorama" // Single panorama image
	GstDewarpDisplayModeDoublePanorama GstDewarpDisplayMode = "double-panorama" // Dewarped image is split in two images displayed one below the other
	GstDewarpDisplayModeQuadView       GstDewarpDisplayMode = "quad-view"       // Dewarped image is split in four images dysplayed as a quad view
)

type GstDewarpInterpolationMode string

const (
	GstDewarpInterpolationModeNearest  GstDewarpInterpolationMode = "nearest"  // A nearest-neighbor interpolation
	GstDewarpInterpolationModeBilinear GstDewarpInterpolationMode = "bilinear" // A bilinear interpolation
	GstDewarpInterpolationModeBicubic  GstDewarpInterpolationMode = "bicubic"  // A bicubic interpolation over 4x4 pixel neighborhood
	GstDewarpInterpolationModeLanczos  GstDewarpInterpolationMode = "Lanczos"  // A Lanczos interpolation over 8x8 pixel neighborhood
)

type GstOpenCVTrackerAlgorithm string

const (
	GstOpenCVTrackerAlgorithmBoosting   GstOpenCVTrackerAlgorithm = "Boosting"   // the Boosting tracker
	GstOpenCVTrackerAlgorithmCSRT       GstOpenCVTrackerAlgorithm = "CSRT"       // the CSRT tracker
	GstOpenCVTrackerAlgorithmKCF        GstOpenCVTrackerAlgorithm = "KCF"        // the KCF (Kernelized Correlation Filter) tracker
	GstOpenCVTrackerAlgorithmMedianFlow GstOpenCVTrackerAlgorithm = "MedianFlow" // the Median Flow tracker
	GstOpenCVTrackerAlgorithmMIL        GstOpenCVTrackerAlgorithm = "MIL"        // the MIL tracker
	GstOpenCVTrackerAlgorithmMOSSE      GstOpenCVTrackerAlgorithm = "MOSSE"      // the MOSSE (Minimum Output Sum of Squared Error) tracker
	GstOpenCVTrackerAlgorithmTLD        GstOpenCVTrackerAlgorithm = "TLD"        // the TLD (Tracking, learning and detection) tracker
)

type GstSkindetectMethod string

const (
	GstSkindetectMethodHsv GstSkindetectMethod = "hsv" // Classic HSV thresholding
	GstSkindetectMethodRgb GstSkindetectMethod = "rgb" // Normalised-RGB colorspace thresholding
)

type GstCvSmoothTypeType string

const (
	GstCvSmoothTypeTypeBlur      GstCvSmoothTypeType = "blur"      // CV Blur
	GstCvSmoothTypeTypeGaussian  GstCvSmoothTypeType = "gaussian"  // CV Gaussian
	GstCvSmoothTypeTypeMedian    GstCvSmoothTypeType = "median"    // CV Median
	GstCvSmoothTypeTypeBilateral GstCvSmoothTypeType = "bilateral" // CV Bilateral
)

type GstDisparityMethod string

const (
	GstDisparityMethodSbm  GstDisparityMethod = "sbm"  // Global block matching algorithm
	GstDisparityMethodSgbm GstDisparityMethod = "sgbm" // Semi-global block matching algorithm
)

type GstFaceDetectUpdates string

const (
	GstFaceDetectUpdatesEveryFrame GstFaceDetectUpdates = "every_frame" // Send update messages on every frame
	GstFaceDetectUpdatesOnChange   GstFaceDetectUpdates = "on_change"   // Send messages when a new face is detected or one is not anymore detected
	GstFaceDetectUpdatesOnFace     GstFaceDetectUpdates = "on_face"     // Send messages whenever a face is detected
	GstFaceDetectUpdatesNone       GstFaceDetectUpdates = "none"        // Send no messages update
)

type GstOpencvFaceBlurFlags int

const (
	GstOpencvFaceBlurFlagsDoCannyPruning GstOpencvFaceBlurFlags = 0x00000001 // Do Canny edge detection to discard some regions
)

type GstOpencvFaceDetectFlags int

const (
	GstOpencvFaceDetectFlagsDoCannyPruning GstOpencvFaceDetectFlags = 0x00000001 // Do Canny edge detection to discard some regions
)

type GstRetinexMethod string

const (
	GstRetinexMethodBasic      GstRetinexMethod = "basic"      // Basic retinex restoration
	GstRetinexMethodMultiscale GstRetinexMethod = "multiscale" // Mutiscale retinex restoration
)

type GstSegmentationMethod string

const (
	GstSegmentationMethodCodebook GstSegmentationMethod = "codebook" // Codebook-based segmentation (Bradski2008)
	GstSegmentationMethodMog      GstSegmentationMethod = "mog"      // Mixture-of-Gaussians segmentation (Bowden2001)
	GstSegmentationMethodMog2     GstSegmentationMethod = "mog2"     // Mixture-of-Gaussians segmentation (Zivkovic2004)
)
