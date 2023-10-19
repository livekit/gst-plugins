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

type Avfvideosrc struct {
	*base.GstPushSrc
}

func NewAvfvideosrc() (*Avfvideosrc, error) {
	e, err := gst.NewElement("avfvideosrc")
	if err != nil {
		return nil, err
	}
	return &Avfvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewAvfvideosrcWithName(name string) (*Avfvideosrc, error) {
	e, err := gst.NewElementWithName("avfvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &Avfvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// capture-screen (bool)
//
// Enable screen capture functionality

func (e *Avfvideosrc) GetCaptureScreen() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("capture-screen")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Avfvideosrc) SetCaptureScreen(value bool) error {
	return e.Element.SetProperty("capture-screen", value)
}

// capture-screen-cursor (bool)
//
// Enable cursor capture while capturing screen

func (e *Avfvideosrc) GetCaptureScreenCursor() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("capture-screen-cursor")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Avfvideosrc) SetCaptureScreenCursor(value bool) error {
	return e.Element.SetProperty("capture-screen-cursor", value)
}

// capture-screen-mouse-clicks (bool)
//
// Enable mouse clicks capture while capturing screen

func (e *Avfvideosrc) GetCaptureScreenMouseClicks() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("capture-screen-mouse-clicks")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Avfvideosrc) SetCaptureScreenMouseClicks(value bool) error {
	return e.Element.SetProperty("capture-screen-mouse-clicks", value)
}

// device-index (int)
//
// The zero-based device index

func (e *Avfvideosrc) GetDeviceIndex() (int, error) {
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

func (e *Avfvideosrc) SetDeviceIndex(value int) error {
	return e.Element.SetProperty("device-index", value)
}

// device-name (string)
//
// The name of the currently opened capture device

func (e *Avfvideosrc) GetDeviceName() (string, error) {
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

// device-type (GstAvfvideoSourceDeviceType)
//
// The general type of a video capture device

func (e *Avfvideosrc) GetDeviceType() (GstAvfvideoSourceDeviceType, error) {
	var value GstAvfvideoSourceDeviceType
	var ok bool
	v, err := e.Element.GetProperty("device-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAvfvideoSourceDeviceType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAvfvideoSourceDeviceType")
	}
	return value, nil
}

func (e *Avfvideosrc) SetDeviceType(value GstAvfvideoSourceDeviceType) error {
	e.Element.SetArg("device-type", string(value))
	return nil
}

// do-stats (bool)
//
// Enable logging of statistics

func (e *Avfvideosrc) GetDoStats() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Avfvideosrc) SetDoStats(value bool) error {
	return e.Element.SetProperty("do-stats", value)
}

// fps (int)
//
// Last measured framerate, if statistics are enabled

func (e *Avfvideosrc) GetFps() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fps")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// orientation (GstAvfvideoSourceOrientation)
//
// The orientation of the video

func (e *Avfvideosrc) GetOrientation() (GstAvfvideoSourceOrientation, error) {
	var value GstAvfvideoSourceOrientation
	var ok bool
	v, err := e.Element.GetProperty("orientation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAvfvideoSourceOrientation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAvfvideoSourceOrientation")
	}
	return value, nil
}

func (e *Avfvideosrc) SetOrientation(value GstAvfvideoSourceOrientation) error {
	e.Element.SetArg("orientation", string(value))
	return nil
}

// position (GstAvfvideoSourcePosition)
//
// The position of the capture device (front or back-facing)

func (e *Avfvideosrc) GetPosition() (GstAvfvideoSourcePosition, error) {
	var value GstAvfvideoSourcePosition
	var ok bool
	v, err := e.Element.GetProperty("position")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAvfvideoSourcePosition)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAvfvideoSourcePosition")
	}
	return value, nil
}

func (e *Avfvideosrc) SetPosition(value GstAvfvideoSourcePosition) error {
	e.Element.SetArg("position", string(value))
	return nil
}

// ----- Constants -----

type GstAvfvideoSourceOrientation string

const (
	GstAvfvideoSourceOrientationPortrait GstAvfvideoSourceOrientation = "portrait" // portrait (1) – Indicates that video should be oriented vertically, top at the top.
	GstAvfvideoSourceOrientationPortratUpsideDown GstAvfvideoSourceOrientation = "portrat-upside-down" // portrat-upside-down (2) – Indicates that video should be oriented vertically, top at the bottom.
	GstAvfvideoSourceOrientationLandscapeRight GstAvfvideoSourceOrientation = "landscape-right" // landscape-right (3) – Indicates that video should be oriented horizontally, top on the left.
	GstAvfvideoSourceOrientationLandscapeLeft GstAvfvideoSourceOrientation = "landscape-left" // landscape-left (4) – Indicates that video should be oriented horizontally, top on the right.
	GstAvfvideoSourceOrientationDefault GstAvfvideoSourceOrientation = "default" // default (0) – Default
)

type GstAvfvideoSourcePosition string

const (
	GstAvfvideoSourcePositionFront GstAvfvideoSourcePosition = "front" // front (1) – Front-facing camera
	GstAvfvideoSourcePositionBack GstAvfvideoSourcePosition = "back" // back (2) – Back-facing camera
	GstAvfvideoSourcePositionDefault GstAvfvideoSourcePosition = "default" // default (0) – Default
)

type GstAvfvideoSourceDeviceType string

const (
	GstAvfvideoSourceDeviceTypeWideAngle GstAvfvideoSourceDeviceType = "wide-angle" // wide-angle (1) – A built-in wide angle camera. These devices are suitable for general purpose use.
	GstAvfvideoSourceDeviceTypeTelephoto GstAvfvideoSourceDeviceType = "telephoto" // telephoto (2) – A built-in camera device with a longer focal length than a wide-angle camera.
	GstAvfvideoSourceDeviceTypeDual GstAvfvideoSourceDeviceType = "dual" // dual (3) – A dual camera device, combining built-in wide-angle and telephoto cameras that work together as a single capture device.
	GstAvfvideoSourceDeviceTypeDefault GstAvfvideoSourceDeviceType = "default" // default (0) – Default
)

