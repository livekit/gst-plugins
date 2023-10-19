// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type Cameracalibrate struct {
	*base.GstBaseTransform
}

func NewCameracalibrate() (*Cameracalibrate, error) {
	e, err := gst.NewElement("cameracalibrate")
	if err != nil {
		return nil, err
	}
	return &Cameracalibrate{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCameracalibrateWithName(name string) (*Cameracalibrate, error) {
	e, err := gst.NewElementWithName("cameracalibrate", name)
	if err != nil {
		return nil, err
	}
	return &Cameracalibrate{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// aspect-ratio (float32)
//
// The aspect ratio

func (e *Cameracalibrate) GetAspectRatio() (float32, error) {
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

func (e *Cameracalibrate) SetAspectRatio(value float32) error {
	return e.Element.SetProperty("aspect-ratio", value)
}

// board-height (int)
//
// The board height in number of items (e.g. number of squares for chessboard)

func (e *Cameracalibrate) GetBoardHeight() (int, error) {
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

func (e *Cameracalibrate) SetBoardHeight(value int) error {
	return e.Element.SetProperty("board-height", value)
}

// board-width (int)
//
// The board width in number of items (e.g. number of squares for chessboard)

func (e *Cameracalibrate) GetBoardWidth() (int, error) {
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

func (e *Cameracalibrate) SetBoardWidth(value int) error {
	return e.Element.SetProperty("board-width", value)
}

// center-principal-point (bool)
//
// Fix the principal point at the center

func (e *Cameracalibrate) GetCenterPrincipalPoint() (bool, error) {
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

func (e *Cameracalibrate) SetCenterPrincipalPoint(value bool) error {
	return e.Element.SetProperty("center-principal-point", value)
}

// corner-sub-pixel (bool)
//
// Improve corner detection accuracy for chessboard

func (e *Cameracalibrate) GetCornerSubPixel() (bool, error) {
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

func (e *Cameracalibrate) SetCornerSubPixel(value bool) error {
	return e.Element.SetProperty("corner-sub-pixel", value)
}

// delay (int)
//
// Sampling periodicity in ms

func (e *Cameracalibrate) GetDelay() (int, error) {
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

func (e *Cameracalibrate) SetDelay(value int) error {
	return e.Element.SetProperty("delay", value)
}

// frame-count (int)
//
// The number of frames to use from the input for calibration

func (e *Cameracalibrate) GetFrameCount() (int, error) {
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

func (e *Cameracalibrate) SetFrameCount(value int) error {
	return e.Element.SetProperty("frame-count", value)
}

// pattern (GstCameraCalibrationPattern)
//
// One of the chessboard, circles, or asymmetric circle pattern

func (e *Cameracalibrate) GetPattern() (GstCameraCalibrationPattern, error) {
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

func (e *Cameracalibrate) SetPattern(value GstCameraCalibrationPattern) error {
	e.Element.SetArg("pattern", string(value))
	return nil
}

// settings (string)
//
// Camera correction parameters (opaque string of serialized OpenCV objects)

func (e *Cameracalibrate) GetSettings() (string, error) {
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

// show-corners (bool)
//
// Show corners

func (e *Cameracalibrate) GetShowCorners() (bool, error) {
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

func (e *Cameracalibrate) SetShowCorners(value bool) error {
	return e.Element.SetProperty("show-corners", value)
}

// square-size (float32)
//
// The size of a square in your defined unit (point, millimeter, etc.)

func (e *Cameracalibrate) GetSquareSize() (float32, error) {
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

func (e *Cameracalibrate) SetSquareSize(value float32) error {
	return e.Element.SetProperty("square-size", value)
}

// use-fisheye (bool)
//
// Use fisheye camera model for calibration

func (e *Cameracalibrate) GetUseFisheye() (bool, error) {
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

func (e *Cameracalibrate) SetUseFisheye(value bool) error {
	return e.Element.SetProperty("use-fisheye", value)
}

// zero-tangent-distorsion (bool)
//
// Assume zero tangential distortion

func (e *Cameracalibrate) GetZeroTangentDistorsion() (bool, error) {
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

func (e *Cameracalibrate) SetZeroTangentDistorsion(value bool) error {
	return e.Element.SetProperty("zero-tangent-distorsion", value)
}

// ----- Constants -----

type GstCameraCalibrationPattern string

const (
	GstCameraCalibrationPatternChessboard GstCameraCalibrationPattern = "chessboard" // chessboard (0) – Chessboard
	GstCameraCalibrationPatternCircleGrids GstCameraCalibrationPattern = "circle_grids" // circle_grids (1) – Circle Grids
	GstCameraCalibrationPatternAsymmetricCircleGrids GstCameraCalibrationPattern = "asymmetric_circle_grids" // asymmetric_circle_grids (2) – Asymmetric Circle Grids
)

