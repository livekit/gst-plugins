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

type Facedetect struct {
	*base.GstBaseTransform
}

func NewFacedetect() (*Facedetect, error) {
	e, err := gst.NewElement("facedetect")
	if err != nil {
		return nil, err
	}
	return &Facedetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFacedetectWithName(name string) (*Facedetect, error) {
	e, err := gst.NewElementWithName("facedetect", name)
	if err != nil {
		return nil, err
	}
	return &Facedetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// display (bool)
//
// Sets whether the detected faces should be highlighted in the output

func (e *Facedetect) GetDisplay() (bool, error) {
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

func (e *Facedetect) SetDisplay(value bool) error {
	return e.Element.SetProperty("display", value)
}

// eyes-profile (string)
//
// Location of Haar cascade file to use for eye-pair detection

func (e *Facedetect) GetEyesProfile() (string, error) {
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

func (e *Facedetect) SetEyesProfile(value string) error {
	return e.Element.SetProperty("eyes-profile", value)
}

// flags (GstOpencvFaceDetectFlags)
//
// Flags to cvHaarDetectObjects

func (e *Facedetect) GetFlags() (GstOpencvFaceDetectFlags, error) {
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

func (e *Facedetect) SetFlags(value GstOpencvFaceDetectFlags) error {
	e.Element.SetArg("flags", string(value))
	return nil
}

// min-neighbors (int)
//
// Minimum number (minus 1) of neighbor rectangles that makes up an object

func (e *Facedetect) GetMinNeighbors() (int, error) {
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

func (e *Facedetect) SetMinNeighbors(value int) error {
	return e.Element.SetProperty("min-neighbors", value)
}

// min-size-height (int)
//
// Minimum area height to be recognized as a face

func (e *Facedetect) GetMinSizeHeight() (int, error) {
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

func (e *Facedetect) SetMinSizeHeight(value int) error {
	return e.Element.SetProperty("min-size-height", value)
}

// min-size-width (int)
//
// Minimum area width to be recognized as a face

func (e *Facedetect) GetMinSizeWidth() (int, error) {
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

func (e *Facedetect) SetMinSizeWidth(value int) error {
	return e.Element.SetProperty("min-size-width", value)
}

// min-stddev (int)
//
// Minimum image average standard deviation: on images with standard deviation lesser than this value facedetection will not be performed. Setting this property help to save cpu and reduce false positives not performing face detection on images with little changes

func (e *Facedetect) GetMinStddev() (int, error) {
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

func (e *Facedetect) SetMinStddev(value int) error {
	return e.Element.SetProperty("min-stddev", value)
}

// mouth-profile (string)
//
// Location of Haar cascade file to use for mouth detection

func (e *Facedetect) GetMouthProfile() (string, error) {
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

func (e *Facedetect) SetMouthProfile(value string) error {
	return e.Element.SetProperty("mouth-profile", value)
}

// nose-profile (string)
//
// Location of Haar cascade file to use for nose detection

func (e *Facedetect) GetNoseProfile() (string, error) {
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

func (e *Facedetect) SetNoseProfile(value string) error {
	return e.Element.SetProperty("nose-profile", value)
}

// profile (string)
//
// Location of Haar cascade file to use for face detection

func (e *Facedetect) GetProfile() (string, error) {
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

func (e *Facedetect) SetProfile(value string) error {
	return e.Element.SetProperty("profile", value)
}

// scale-factor (float64)
//
// Factor by which the frame is scaled after each object scan

func (e *Facedetect) GetScaleFactor() (float64, error) {
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

func (e *Facedetect) SetScaleFactor(value float64) error {
	return e.Element.SetProperty("scale-factor", value)
}

// updates (GstFaceDetectUpdates)
//
// When send update bus messages, if at all

func (e *Facedetect) GetUpdates() (GstFaceDetectUpdates, error) {
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

func (e *Facedetect) SetUpdates(value GstFaceDetectUpdates) error {
	e.Element.SetArg("updates", string(value))
	return nil
}

// ----- Constants -----

type GstFaceDetectUpdates string

const (
	GstFaceDetectUpdatesEveryFrame GstFaceDetectUpdates = "every_frame" // every_frame (0) – Send update messages on every frame
	GstFaceDetectUpdatesOnChange GstFaceDetectUpdates = "on_change" // on_change (1) – Send messages when a new face is detected or one is not anymore detected
	GstFaceDetectUpdatesOnFace GstFaceDetectUpdates = "on_face" // on_face (2) – Send messages whenever a face is detected
	GstFaceDetectUpdatesNone GstFaceDetectUpdates = "none" // none (3) – Send no messages update
)

type GstOpencvFaceDetectFlags string

const (
	GstOpencvFaceDetectFlagsDoCannyPruning GstOpencvFaceDetectFlags = "do-canny-pruning" // do-canny-pruning (0x00000001) – Do Canny edge detection to discard some regions
)

