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

type Motioncells struct {
	*base.GstBaseTransform
}

func NewMotioncells() (*Motioncells, error) {
	e, err := gst.NewElement("motioncells")
	if err != nil {
		return nil, err
	}
	return &Motioncells{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewMotioncellsWithName(name string) (*Motioncells, error) {
	e, err := gst.NewElementWithName("motioncells", name)
	if err != nil {
		return nil, err
	}
	return &Motioncells{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// calculatemotion (bool)
//
// Toggles motion calculation. If FALSE, this filter does nothing

func (e *Motioncells) GetCalculatemotion() (bool, error) {
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

func (e *Motioncells) SetCalculatemotion(value bool) error {
	return e.Element.SetProperty("calculatemotion", value)
}

// cellscolor (string)
//
// Color for motion cells in R,G,B format. Max per channel is 255

func (e *Motioncells) GetCellscolor() (string, error) {
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

func (e *Motioncells) SetCellscolor(value string) error {
	return e.Element.SetProperty("cellscolor", value)
}

// datafile (string)
//
// Location of motioncells data file (empty string means no saving)

func (e *Motioncells) GetDatafile() (string, error) {
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

func (e *Motioncells) SetDatafile(value string) error {
	return e.Element.SetProperty("datafile", value)
}

// datafileextension (string)
//
// Extension of datafile

func (e *Motioncells) GetDatafileextension() (string, error) {
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

func (e *Motioncells) SetDatafileextension(value string) error {
	return e.Element.SetProperty("datafileextension", value)
}

// display (bool)
//
// Toggle display of motion cells on current frame

func (e *Motioncells) GetDisplay() (bool, error) {
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

func (e *Motioncells) SetDisplay(value bool) error {
	return e.Element.SetProperty("display", value)
}

// gap (int)
//
// Interval in seconds after which motion is considered finished and a motion finished bus message is posted.

func (e *Motioncells) GetGap() (int, error) {
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

func (e *Motioncells) SetGap(value int) error {
	return e.Element.SetProperty("gap", value)
}

// gridx (int)
//
// Number of horizontal grid cells.

func (e *Motioncells) GetGridx() (int, error) {
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

func (e *Motioncells) SetGridx(value int) error {
	return e.Element.SetProperty("gridx", value)
}

// gridy (int)
//
// Number of vertical grid cells.

func (e *Motioncells) GetGridy() (int, error) {
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

func (e *Motioncells) SetGridy(value int) error {
	return e.Element.SetProperty("gridy", value)
}

// minimummotionframes (int)
//
// Minimum number of motion frames triggering a motion event

func (e *Motioncells) GetMinimummotionframes() (int, error) {
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

func (e *Motioncells) SetMinimummotionframes(value int) error {
	return e.Element.SetProperty("minimummotionframes", value)
}

// motioncellsidx (string)
//
// Describe a cell with its line and column idx separated with ":". Pass multiple cells as a comma-separated list

func (e *Motioncells) GetMotioncellsidx() (string, error) {
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

func (e *Motioncells) SetMotioncellsidx(value string) error {
	return e.Element.SetProperty("motioncellsidx", value)
}

// motioncellthickness (int)
//
// Motion Cell Border Thickness. Set to -1 to fill motion cell

func (e *Motioncells) GetMotioncellthickness() (int, error) {
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

func (e *Motioncells) SetMotioncellthickness(value int) error {
	return e.Element.SetProperty("motioncellthickness", value)
}

// motionmaskcellspos (string)
//
// Describe a cell with its line and column idx separated with ":". Pass multiple cells as a comma-separated list

func (e *Motioncells) GetMotionmaskcellspos() (string, error) {
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

func (e *Motioncells) SetMotionmaskcellspos(value string) error {
	return e.Element.SetProperty("motionmaskcellspos", value)
}

// motionmaskcoords (string)
//
// Describe a region with its upper left and lower right x, y coordinates separated with ":". Pass multiple regions as a comma-separated list

func (e *Motioncells) GetMotionmaskcoords() (string, error) {
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

func (e *Motioncells) SetMotionmaskcoords(value string) error {
	return e.Element.SetProperty("motionmaskcoords", value)
}

// postallmotion (bool)
//
// Post bus messages for every motion frame or just motion start and motion stop

func (e *Motioncells) GetPostallmotion() (bool, error) {
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

func (e *Motioncells) SetPostallmotion(value bool) error {
	return e.Element.SetProperty("postallmotion", value)
}

// postnomotion (int)
//
// If non 0, post a no_motion event on the bus if no motion is detected for the given number of seconds

func (e *Motioncells) GetPostnomotion() (int, error) {
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

func (e *Motioncells) SetPostnomotion(value int) error {
	return e.Element.SetProperty("postnomotion", value)
}

// sensitivity (float64)
//
// Motion detection sensitivity.

func (e *Motioncells) GetSensitivity() (float64, error) {
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

func (e *Motioncells) SetSensitivity(value float64) error {
	return e.Element.SetProperty("sensitivity", value)
}

// threshold (float64)
//
// Threshold value for motion. Filter detects motion when at least this fraction of the cells have moved

func (e *Motioncells) GetThreshold() (float64, error) {
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

func (e *Motioncells) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

// usealpha (bool)
//
// Toggle usage of alpha blending on frames with motion cells

func (e *Motioncells) GetUsealpha() (bool, error) {
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

func (e *Motioncells) SetUsealpha(value bool) error {
	return e.Element.SetProperty("usealpha", value)
}

