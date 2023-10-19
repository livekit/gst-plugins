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

type Smptealpha struct {
	*base.GstBaseTransform
}

func NewSmptealpha() (*Smptealpha, error) {
	e, err := gst.NewElement("smptealpha")
	if err != nil {
		return nil, err
	}
	return &Smptealpha{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSmptealphaWithName(name string) (*Smptealpha, error) {
	e, err := gst.NewElementWithName("smptealpha", name)
	if err != nil {
		return nil, err
	}
	return &Smptealpha{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// border (int)
//
// The border width of the transition

func (e *Smptealpha) GetBorder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("border")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Smptealpha) SetBorder(value int) error {
	return e.Element.SetProperty("border", value)
}

// depth (int)
//
// Depth of the mask in bits

func (e *Smptealpha) GetDepth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Smptealpha) SetDepth(value int) error {
	return e.Element.SetProperty("depth", value)
}

// invert (bool)
//
// Set to TRUE to invert the transition mask (ie. flip it horizontally).

func (e *Smptealpha) GetInvert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Smptealpha) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

// position (float64)
//
// Position of the transition effect

func (e *Smptealpha) GetPosition() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("position")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Smptealpha) SetPosition(value float64) error {
	return e.Element.SetProperty("position", value)
}

// type (GstSmptealphaTransitionType)
//
// The type of transition to use

func (e *Smptealpha) GetType() (GstSmptealphaTransitionType, error) {
	var value GstSmptealphaTransitionType
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSmptealphaTransitionType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSmptealphaTransitionType")
	}
	return value, nil
}

func (e *Smptealpha) SetType(value GstSmptealphaTransitionType) error {
	e.Element.SetArg("type", string(value))
	return nil
}

// ----- Constants -----

type GstSmptealphaTransitionType string

const (
	GstSmptealphaTransitionTypeBarWipeLr GstSmptealphaTransitionType = "bar-wipe-lr" // bar-wipe-lr (1) – A bar moves from left to right
	GstSmptealphaTransitionTypeBarWipeTb GstSmptealphaTransitionType = "bar-wipe-tb" // bar-wipe-tb (2) – A bar moves from top to bottom
	GstSmptealphaTransitionTypeBoxWipeTl GstSmptealphaTransitionType = "box-wipe-tl" // box-wipe-tl (3) – A box expands from the upper-left corner to the lower-right corner
	GstSmptealphaTransitionTypeBoxWipeTr GstSmptealphaTransitionType = "box-wipe-tr" // box-wipe-tr (4) – A box expands from the upper-right corner to the lower-left corner
	GstSmptealphaTransitionTypeBoxWipeBr GstSmptealphaTransitionType = "box-wipe-br" // box-wipe-br (5) – A box expands from the lower-right corner to the upper-left corner
	GstSmptealphaTransitionTypeBoxWipeBl GstSmptealphaTransitionType = "box-wipe-bl" // box-wipe-bl (6) – A box expands from the lower-left corner to the upper-right corner
	GstSmptealphaTransitionTypeFourBoxWipeCi GstSmptealphaTransitionType = "four-box-wipe-ci" // four-box-wipe-ci (7) – A box shape expands from each of the four corners toward the center
	GstSmptealphaTransitionTypeFourBoxWipeCo GstSmptealphaTransitionType = "four-box-wipe-co" // four-box-wipe-co (8) – A box shape expands from the center of each quadrant toward the corners of each quadrant
	GstSmptealphaTransitionTypeBarndoorV GstSmptealphaTransitionType = "barndoor-v" // barndoor-v (21) – A central, vertical line splits and expands toward the left and right edges
	GstSmptealphaTransitionTypeBarndoorH GstSmptealphaTransitionType = "barndoor-h" // barndoor-h (22) – A central, horizontal line splits and expands toward the top and bottom edges
	GstSmptealphaTransitionTypeBoxWipeTc GstSmptealphaTransitionType = "box-wipe-tc" // box-wipe-tc (23) – A box expands from the top edge's midpoint to the bottom corners
	GstSmptealphaTransitionTypeBoxWipeRc GstSmptealphaTransitionType = "box-wipe-rc" // box-wipe-rc (24) – A box expands from the right edge's midpoint to the left corners
	GstSmptealphaTransitionTypeBoxWipeBc GstSmptealphaTransitionType = "box-wipe-bc" // box-wipe-bc (25) – A box expands from the bottom edge's midpoint to the top corners
	GstSmptealphaTransitionTypeBoxWipeLc GstSmptealphaTransitionType = "box-wipe-lc" // box-wipe-lc (26) – A box expands from the left edge's midpoint to the right corners
	GstSmptealphaTransitionTypeDiagonalTl GstSmptealphaTransitionType = "diagonal-tl" // diagonal-tl (41) – A diagonal line moves from the upper-left corner to the lower-right corner
	GstSmptealphaTransitionTypeDiagonalTr GstSmptealphaTransitionType = "diagonal-tr" // diagonal-tr (42) – A diagonal line moves from the upper right corner to the lower-left corner
	GstSmptealphaTransitionTypeBowtieV GstSmptealphaTransitionType = "bowtie-v" // bowtie-v (43) – Two wedge shapes slide in from the top and bottom edges toward the center
	GstSmptealphaTransitionTypeBowtieH GstSmptealphaTransitionType = "bowtie-h" // bowtie-h (44) – Two wedge shapes slide in from the left and right edges toward the center
	GstSmptealphaTransitionTypeBarndoorDbl GstSmptealphaTransitionType = "barndoor-dbl" // barndoor-dbl (45) – A diagonal line from the lower-left to upper-right corners splits and expands toward the opposite corners
	GstSmptealphaTransitionTypeBarndoorDtl GstSmptealphaTransitionType = "barndoor-dtl" // barndoor-dtl (46) – A diagonal line from upper-left to lower-right corners splits and expands toward the opposite corners
	GstSmptealphaTransitionTypeMiscDiagonalDbd GstSmptealphaTransitionType = "misc-diagonal-dbd" // misc-diagonal-dbd (47) – Four wedge shapes split from the center and retract toward the four edges
	GstSmptealphaTransitionTypeMiscDiagonalDd GstSmptealphaTransitionType = "misc-diagonal-dd" // misc-diagonal-dd (48) – A diamond connecting the four edge midpoints simultaneously contracts toward the center and expands toward the edges
	GstSmptealphaTransitionTypeVeeD GstSmptealphaTransitionType = "vee-d" // vee-d (61) – A wedge shape moves from top to bottom
	GstSmptealphaTransitionTypeVeeL GstSmptealphaTransitionType = "vee-l" // vee-l (62) – A wedge shape moves from right to left
	GstSmptealphaTransitionTypeVeeU GstSmptealphaTransitionType = "vee-u" // vee-u (63) – A wedge shape moves from bottom to top
	GstSmptealphaTransitionTypeVeeR GstSmptealphaTransitionType = "vee-r" // vee-r (64) – A wedge shape moves from left to right
	GstSmptealphaTransitionTypeBarnveeD GstSmptealphaTransitionType = "barnvee-d" // barnvee-d (65) – A 'V' shape extending from the bottom edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSmptealphaTransitionTypeBarnveeL GstSmptealphaTransitionType = "barnvee-l" // barnvee-l (66) – A 'V' shape extending from the left edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSmptealphaTransitionTypeBarnveeU GstSmptealphaTransitionType = "barnvee-u" // barnvee-u (67) – A 'V' shape extending from the top edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSmptealphaTransitionTypeBarnveeR GstSmptealphaTransitionType = "barnvee-r" // barnvee-r (68) – A 'V' shape extending from the right edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSmptealphaTransitionTypeIrisRect GstSmptealphaTransitionType = "iris-rect" // iris-rect (101) – A rectangle expands from the center.
	GstSmptealphaTransitionTypeClockCw12 GstSmptealphaTransitionType = "clock-cw12" // clock-cw12 (201) – A radial hand sweeps clockwise from the twelve o'clock position
	GstSmptealphaTransitionTypeClockCw3 GstSmptealphaTransitionType = "clock-cw3" // clock-cw3 (202) – A radial hand sweeps clockwise from the three o'clock position
	GstSmptealphaTransitionTypeClockCw6 GstSmptealphaTransitionType = "clock-cw6" // clock-cw6 (203) – A radial hand sweeps clockwise from the six o'clock position
	GstSmptealphaTransitionTypeClockCw9 GstSmptealphaTransitionType = "clock-cw9" // clock-cw9 (204) – A radial hand sweeps clockwise from the nine o'clock position
	GstSmptealphaTransitionTypePinwheelTbv GstSmptealphaTransitionType = "pinwheel-tbv" // pinwheel-tbv (205) – Two radial hands sweep clockwise from the twelve and six o'clock positions
	GstSmptealphaTransitionTypePinwheelTbh GstSmptealphaTransitionType = "pinwheel-tbh" // pinwheel-tbh (206) – Two radial hands sweep clockwise from the nine and three o'clock positions
	GstSmptealphaTransitionTypePinwheelFb GstSmptealphaTransitionType = "pinwheel-fb" // pinwheel-fb (207) – Four radial hands sweep clockwise
	GstSmptealphaTransitionTypeFanCt GstSmptealphaTransitionType = "fan-ct" // fan-ct (211) – A fan unfolds from the top edge, the fan axis at the center
	GstSmptealphaTransitionTypeFanCr GstSmptealphaTransitionType = "fan-cr" // fan-cr (212) – A fan unfolds from the right edge, the fan axis at the center
	GstSmptealphaTransitionTypeDoublefanFov GstSmptealphaTransitionType = "doublefan-fov" // doublefan-fov (213) – Two fans, their axes at the center, unfold from the top and bottom
	GstSmptealphaTransitionTypeDoublefanFoh GstSmptealphaTransitionType = "doublefan-foh" // doublefan-foh (214) – Two fans, their axes at the center, unfold from the left and right
	GstSmptealphaTransitionTypeSinglesweepCwt GstSmptealphaTransitionType = "singlesweep-cwt" // singlesweep-cwt (221) – A radial hand sweeps clockwise from the top edge's midpoint
	GstSmptealphaTransitionTypeSinglesweepCwr GstSmptealphaTransitionType = "singlesweep-cwr" // singlesweep-cwr (222) – A radial hand sweeps clockwise from the right edge's midpoint
	GstSmptealphaTransitionTypeSinglesweepCwb GstSmptealphaTransitionType = "singlesweep-cwb" // singlesweep-cwb (223) – A radial hand sweeps clockwise from the bottom edge's midpoint
	GstSmptealphaTransitionTypeSinglesweepCwl GstSmptealphaTransitionType = "singlesweep-cwl" // singlesweep-cwl (224) – A radial hand sweeps clockwise from the left edge's midpoint
	GstSmptealphaTransitionTypeDoublesweepPv GstSmptealphaTransitionType = "doublesweep-pv" // doublesweep-pv (225) – Two radial hands sweep clockwise and counter-clockwise from the top and bottom edges' midpoints
	GstSmptealphaTransitionTypeDoublesweepPd GstSmptealphaTransitionType = "doublesweep-pd" // doublesweep-pd (226) – Two radial hands sweep clockwise and counter-clockwise from the left and right edges' midpoints
	GstSmptealphaTransitionTypeDoublesweepOv GstSmptealphaTransitionType = "doublesweep-ov" // doublesweep-ov (227) – Two radial hands attached at the top and bottom edges' midpoints sweep from right to left
	GstSmptealphaTransitionTypeDoublesweepOh GstSmptealphaTransitionType = "doublesweep-oh" // doublesweep-oh (228) – Two radial hands attached at the left and right edges' midpoints sweep from top to bottom
	GstSmptealphaTransitionTypeFanT GstSmptealphaTransitionType = "fan-t" // fan-t (231) – A fan unfolds from the bottom, the fan axis at the top edge's midpoint
	GstSmptealphaTransitionTypeFanR GstSmptealphaTransitionType = "fan-r" // fan-r (232) – A fan unfolds from the left, the fan axis at the right edge's midpoint
	GstSmptealphaTransitionTypeFanB GstSmptealphaTransitionType = "fan-b" // fan-b (233) – A fan unfolds from the top, the fan axis at the bottom edge's midpoint
	GstSmptealphaTransitionTypeFanL GstSmptealphaTransitionType = "fan-l" // fan-l (234) – A fan unfolds from the right, the fan axis at the left edge's midpoint
	GstSmptealphaTransitionTypeDoublefanFiv GstSmptealphaTransitionType = "doublefan-fiv" // doublefan-fiv (235) – Two fans, their axes at the top and bottom, unfold from the center
	GstSmptealphaTransitionTypeDoublefanFih GstSmptealphaTransitionType = "doublefan-fih" // doublefan-fih (236) – Two fans, their axes at the left and right, unfold from the center
	GstSmptealphaTransitionTypeSinglesweepCwtl GstSmptealphaTransitionType = "singlesweep-cwtl" // singlesweep-cwtl (241) – A radial hand sweeps clockwise from the upper-left corner
	GstSmptealphaTransitionTypeSinglesweepCwbl GstSmptealphaTransitionType = "singlesweep-cwbl" // singlesweep-cwbl (242) – A radial hand sweeps counter-clockwise from the lower-left corner.
	GstSmptealphaTransitionTypeSinglesweepCwbr GstSmptealphaTransitionType = "singlesweep-cwbr" // singlesweep-cwbr (243) – A radial hand sweeps clockwise from the lower-right corner
	GstSmptealphaTransitionTypeSinglesweepCwtr GstSmptealphaTransitionType = "singlesweep-cwtr" // singlesweep-cwtr (244) – A radial hand sweeps counter-clockwise from the upper-right corner
	GstSmptealphaTransitionTypeDoublesweepPdtl GstSmptealphaTransitionType = "doublesweep-pdtl" // doublesweep-pdtl (245) – Two radial hands attached at the upper-left and lower-right corners sweep down and up
	GstSmptealphaTransitionTypeDoublesweepPdbl GstSmptealphaTransitionType = "doublesweep-pdbl" // doublesweep-pdbl (246) – Two radial hands attached at the lower-left and upper-right corners sweep down and up
	GstSmptealphaTransitionTypeSaloondoorT GstSmptealphaTransitionType = "saloondoor-t" // saloondoor-t (251) – Two radial hands attached at the upper-left and upper-right corners sweep down
	GstSmptealphaTransitionTypeSaloondoorL GstSmptealphaTransitionType = "saloondoor-l" // saloondoor-l (252) – Two radial hands attached at the upper-left and lower-left corners sweep to the right
	GstSmptealphaTransitionTypeSaloondoorB GstSmptealphaTransitionType = "saloondoor-b" // saloondoor-b (253) – Two radial hands attached at the lower-left and lower-right corners sweep up
	GstSmptealphaTransitionTypeSaloondoorR GstSmptealphaTransitionType = "saloondoor-r" // saloondoor-r (254) – Two radial hands attached at the upper-right and lower-right corners sweep to the left
	GstSmptealphaTransitionTypeWindshieldR GstSmptealphaTransitionType = "windshield-r" // windshield-r (261) – Two radial hands attached at the midpoints of the top and bottom halves sweep from right to left
	GstSmptealphaTransitionTypeWindshieldU GstSmptealphaTransitionType = "windshield-u" // windshield-u (262) – Two radial hands attached at the midpoints of the left and right halves sweep from top to bottom
	GstSmptealphaTransitionTypeWindshieldV GstSmptealphaTransitionType = "windshield-v" // windshield-v (263) – Two sets of radial hands attached at the midpoints of the top and bottom halves sweep from top to bottom and bottom to top
	GstSmptealphaTransitionTypeWindshieldH GstSmptealphaTransitionType = "windshield-h" // windshield-h (264) – Two sets of radial hands attached at the midpoints of the left and right halves sweep from left to right and right to left
)

