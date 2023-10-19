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

type Frei0RFilterOpencvfacedetect struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterOpencvfacedetect() (*Frei0RFilterOpencvfacedetect, error) {
	e, err := gst.NewElement("frei0r-filter-opencvfacedetect")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterOpencvfacedetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterOpencvfacedetectWithName(name string) (*Frei0RFilterOpencvfacedetect, error) {
	e, err := gst.NewElementWithName("frei0r-filter-opencvfacedetect", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterOpencvfacedetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha (float64)
//
// The alpha channel value for the shapes

func (e *Frei0RFilterOpencvfacedetect) GetAlpha() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetAlpha(value float64) error {
	return e.Element.SetProperty("alpha", value)
}

// antialias (bool)
//
// Draw with antialiasing

func (e *Frei0RFilterOpencvfacedetect) GetAntialias() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("antialias")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetAntialias(value bool) error {
	return e.Element.SetProperty("antialias", value)
}

// classifier (string)
//
// Full path to the XML pattern model for recognition; look in /usr/share/opencv/haarcascades

func (e *Frei0RFilterOpencvfacedetect) GetClassifier() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("classifier")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetClassifier(value string) error {
	return e.Element.SetProperty("classifier", value)
}

// color-1-b (float32)
//
// The color of the first object

func (e *Frei0RFilterOpencvfacedetect) GetColor1B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-1-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor1B(value float32) error {
	return e.Element.SetProperty("color-1-b", value)
}

// color-1-g (float32)
//
// The color of the first object

func (e *Frei0RFilterOpencvfacedetect) GetColor1G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-1-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor1G(value float32) error {
	return e.Element.SetProperty("color-1-g", value)
}

// color-1-r (float32)
//
// The color of the first object

func (e *Frei0RFilterOpencvfacedetect) GetColor1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-1-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor1R(value float32) error {
	return e.Element.SetProperty("color-1-r", value)
}

// color-2-b (float32)
//
// The color of the second object

func (e *Frei0RFilterOpencvfacedetect) GetColor2B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-2-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor2B(value float32) error {
	return e.Element.SetProperty("color-2-b", value)
}

// color-2-g (float32)
//
// The color of the second object

func (e *Frei0RFilterOpencvfacedetect) GetColor2G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-2-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor2G(value float32) error {
	return e.Element.SetProperty("color-2-g", value)
}

// color-2-r (float32)
//
// The color of the second object

func (e *Frei0RFilterOpencvfacedetect) GetColor2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-2-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor2R(value float32) error {
	return e.Element.SetProperty("color-2-r", value)
}

// color-3-b (float32)
//
// The color of the third object

func (e *Frei0RFilterOpencvfacedetect) GetColor3B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-3-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor3B(value float32) error {
	return e.Element.SetProperty("color-3-b", value)
}

// color-3-g (float32)
//
// The color of the third object

func (e *Frei0RFilterOpencvfacedetect) GetColor3G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-3-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor3G(value float32) error {
	return e.Element.SetProperty("color-3-g", value)
}

// color-3-r (float32)
//
// The color of the third object

func (e *Frei0RFilterOpencvfacedetect) GetColor3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-3-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor3R(value float32) error {
	return e.Element.SetProperty("color-3-r", value)
}

// color-4-b (float32)
//
// The color of the fourth object

func (e *Frei0RFilterOpencvfacedetect) GetColor4B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-4-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor4B(value float32) error {
	return e.Element.SetProperty("color-4-b", value)
}

// color-4-g (float32)
//
// The color of the fourth object

func (e *Frei0RFilterOpencvfacedetect) GetColor4G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-4-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor4G(value float32) error {
	return e.Element.SetProperty("color-4-g", value)
}

// color-4-r (float32)
//
// The color of the fourth object

func (e *Frei0RFilterOpencvfacedetect) GetColor4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-4-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor4R(value float32) error {
	return e.Element.SetProperty("color-4-r", value)
}

// color-5-b (float32)
//
// The color of the fifth object

func (e *Frei0RFilterOpencvfacedetect) GetColor5B() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-5-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor5B(value float32) error {
	return e.Element.SetProperty("color-5-b", value)
}

// color-5-g (float32)
//
// The color of the fifth object

func (e *Frei0RFilterOpencvfacedetect) GetColor5G() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-5-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor5G(value float32) error {
	return e.Element.SetProperty("color-5-g", value)
}

// color-5-r (float32)
//
// The color of the fifth object

func (e *Frei0RFilterOpencvfacedetect) GetColor5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("color-5-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetColor5R(value float32) error {
	return e.Element.SetProperty("color-5-r", value)
}

// neighbors (float64)
//
// Minimum number of rectangles that makes up an object, divided by 100

func (e *Frei0RFilterOpencvfacedetect) GetNeighbors() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("neighbors")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetNeighbors(value float64) error {
	return e.Element.SetProperty("neighbors", value)
}

// recheck (float64)
//
// How often to detect an object in number of frames, divided by 1000

func (e *Frei0RFilterOpencvfacedetect) GetRecheck() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("recheck")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetRecheck(value float64) error {
	return e.Element.SetProperty("recheck", value)
}

// scale (float64)
//
// Down scale the image prior detection

func (e *Frei0RFilterOpencvfacedetect) GetScale() (float64, error) {
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

func (e *Frei0RFilterOpencvfacedetect) SetScale(value float64) error {
	return e.Element.SetProperty("scale", value)
}

// search-scale (float64)
//
// The search window scale factor, divided by 10

func (e *Frei0RFilterOpencvfacedetect) GetSearchScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("search-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetSearchScale(value float64) error {
	return e.Element.SetProperty("search-scale", value)
}

// shape (float64)
//
// The shape to draw: 0=circle, 0.1=ellipse, 0.2=rectangle, 1=random

func (e *Frei0RFilterOpencvfacedetect) GetShape() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("shape")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetShape(value float64) error {
	return e.Element.SetProperty("shape", value)
}

// smallest (float64)
//
// Minimum window size in pixels, divided by 1000

func (e *Frei0RFilterOpencvfacedetect) GetSmallest() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("smallest")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetSmallest(value float64) error {
	return e.Element.SetProperty("smallest", value)
}

// stroke (float64)
//
// Line width, divided by 100, or fill if 0

func (e *Frei0RFilterOpencvfacedetect) GetStroke() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("stroke")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetStroke(value float64) error {
	return e.Element.SetProperty("stroke", value)
}

// threads (float64)
//
// How many threads to use divided by 100; 0 uses CPU count

func (e *Frei0RFilterOpencvfacedetect) GetThreads() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterOpencvfacedetect) SetThreads(value float64) error {
	return e.Element.SetProperty("threads", value)
}

