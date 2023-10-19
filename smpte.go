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
)

type Smpte struct {
	Element *gst.Element
}

func NewSmpte() (*Smpte, error) {
	e, err := gst.NewElement("smpte")
	if err != nil {
		return nil, err
	}
	return &Smpte{Element: e}, nil
}

func NewSmpteWithName(name string) (*Smpte, error) {
	e, err := gst.NewElementWithName("smpte", name)
	if err != nil {
		return nil, err
	}
	return &Smpte{Element: e}, nil
}

// ----- Properties -----

// border (int)
//
// The border width of the transition

func (e *Smpte) GetBorder() (int, error) {
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

func (e *Smpte) SetBorder(value int) error {
	return e.Element.SetProperty("border", value)
}

// depth (int)
//
// Depth of the mask in bits

func (e *Smpte) GetDepth() (int, error) {
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

func (e *Smpte) SetDepth(value int) error {
	return e.Element.SetProperty("depth", value)
}

// duration (uint64)
//
// Duration of the transition effect in nanoseconds

func (e *Smpte) GetDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Smpte) SetDuration(value uint64) error {
	return e.Element.SetProperty("duration", value)
}

// invert (bool)
//
// Invert transition mask

func (e *Smpte) GetInvert() (bool, error) {
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

func (e *Smpte) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

// type (GstSmptetransitionType)
//
// The type of transition to use

func (e *Smpte) GetType() (GstSmptetransitionType, error) {
	var value GstSmptetransitionType
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSmptetransitionType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSmptetransitionType")
	}
	return value, nil
}

func (e *Smpte) SetType(value GstSmptetransitionType) error {
	e.Element.SetArg("type", string(value))
	return nil
}

// ----- Constants -----

type GstSmptetransitionType string

const (
	GstSmptetransitionTypeBarWipeLr GstSmptetransitionType = "bar-wipe-lr" // bar-wipe-lr (1) – A bar moves from left to right
	GstSmptetransitionTypeBarWipeTb GstSmptetransitionType = "bar-wipe-tb" // bar-wipe-tb (2) – A bar moves from top to bottom
	GstSmptetransitionTypeBoxWipeTl GstSmptetransitionType = "box-wipe-tl" // box-wipe-tl (3) – A box expands from the upper-left corner to the lower-right corner
	GstSmptetransitionTypeBoxWipeTr GstSmptetransitionType = "box-wipe-tr" // box-wipe-tr (4) – A box expands from the upper-right corner to the lower-left corner
	GstSmptetransitionTypeBoxWipeBr GstSmptetransitionType = "box-wipe-br" // box-wipe-br (5) – A box expands from the lower-right corner to the upper-left corner
	GstSmptetransitionTypeBoxWipeBl GstSmptetransitionType = "box-wipe-bl" // box-wipe-bl (6) – A box expands from the lower-left corner to the upper-right corner
	GstSmptetransitionTypeFourBoxWipeCi GstSmptetransitionType = "four-box-wipe-ci" // four-box-wipe-ci (7) – A box shape expands from each of the four corners toward the center
	GstSmptetransitionTypeFourBoxWipeCo GstSmptetransitionType = "four-box-wipe-co" // four-box-wipe-co (8) – A box shape expands from the center of each quadrant toward the corners of each quadrant
	GstSmptetransitionTypeBarndoorV GstSmptetransitionType = "barndoor-v" // barndoor-v (21) – A central, vertical line splits and expands toward the left and right edges
	GstSmptetransitionTypeBarndoorH GstSmptetransitionType = "barndoor-h" // barndoor-h (22) – A central, horizontal line splits and expands toward the top and bottom edges
	GstSmptetransitionTypeBoxWipeTc GstSmptetransitionType = "box-wipe-tc" // box-wipe-tc (23) – A box expands from the top edge's midpoint to the bottom corners
	GstSmptetransitionTypeBoxWipeRc GstSmptetransitionType = "box-wipe-rc" // box-wipe-rc (24) – A box expands from the right edge's midpoint to the left corners
	GstSmptetransitionTypeBoxWipeBc GstSmptetransitionType = "box-wipe-bc" // box-wipe-bc (25) – A box expands from the bottom edge's midpoint to the top corners
	GstSmptetransitionTypeBoxWipeLc GstSmptetransitionType = "box-wipe-lc" // box-wipe-lc (26) – A box expands from the left edge's midpoint to the right corners
	GstSmptetransitionTypeDiagonalTl GstSmptetransitionType = "diagonal-tl" // diagonal-tl (41) – A diagonal line moves from the upper-left corner to the lower-right corner
	GstSmptetransitionTypeDiagonalTr GstSmptetransitionType = "diagonal-tr" // diagonal-tr (42) – A diagonal line moves from the upper right corner to the lower-left corner
	GstSmptetransitionTypeBowtieV GstSmptetransitionType = "bowtie-v" // bowtie-v (43) – Two wedge shapes slide in from the top and bottom edges toward the center
	GstSmptetransitionTypeBowtieH GstSmptetransitionType = "bowtie-h" // bowtie-h (44) – Two wedge shapes slide in from the left and right edges toward the center
	GstSmptetransitionTypeBarndoorDbl GstSmptetransitionType = "barndoor-dbl" // barndoor-dbl (45) – A diagonal line from the lower-left to upper-right corners splits and expands toward the opposite corners
	GstSmptetransitionTypeBarndoorDtl GstSmptetransitionType = "barndoor-dtl" // barndoor-dtl (46) – A diagonal line from upper-left to lower-right corners splits and expands toward the opposite corners
	GstSmptetransitionTypeMiscDiagonalDbd GstSmptetransitionType = "misc-diagonal-dbd" // misc-diagonal-dbd (47) – Four wedge shapes split from the center and retract toward the four edges
	GstSmptetransitionTypeMiscDiagonalDd GstSmptetransitionType = "misc-diagonal-dd" // misc-diagonal-dd (48) – A diamond connecting the four edge midpoints simultaneously contracts toward the center and expands toward the edges
	GstSmptetransitionTypeVeeD GstSmptetransitionType = "vee-d" // vee-d (61) – A wedge shape moves from top to bottom
	GstSmptetransitionTypeVeeL GstSmptetransitionType = "vee-l" // vee-l (62) – A wedge shape moves from right to left
	GstSmptetransitionTypeVeeU GstSmptetransitionType = "vee-u" // vee-u (63) – A wedge shape moves from bottom to top
	GstSmptetransitionTypeVeeR GstSmptetransitionType = "vee-r" // vee-r (64) – A wedge shape moves from left to right
	GstSmptetransitionTypeBarnveeD GstSmptetransitionType = "barnvee-d" // barnvee-d (65) – A 'V' shape extending from the bottom edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSmptetransitionTypeBarnveeL GstSmptetransitionType = "barnvee-l" // barnvee-l (66) – A 'V' shape extending from the left edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSmptetransitionTypeBarnveeU GstSmptetransitionType = "barnvee-u" // barnvee-u (67) – A 'V' shape extending from the top edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSmptetransitionTypeBarnveeR GstSmptetransitionType = "barnvee-r" // barnvee-r (68) – A 'V' shape extending from the right edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSmptetransitionTypeIrisRect GstSmptetransitionType = "iris-rect" // iris-rect (101) – A rectangle expands from the center.
	GstSmptetransitionTypeClockCw12 GstSmptetransitionType = "clock-cw12" // clock-cw12 (201) – A radial hand sweeps clockwise from the twelve o'clock position
	GstSmptetransitionTypeClockCw3 GstSmptetransitionType = "clock-cw3" // clock-cw3 (202) – A radial hand sweeps clockwise from the three o'clock position
	GstSmptetransitionTypeClockCw6 GstSmptetransitionType = "clock-cw6" // clock-cw6 (203) – A radial hand sweeps clockwise from the six o'clock position
	GstSmptetransitionTypeClockCw9 GstSmptetransitionType = "clock-cw9" // clock-cw9 (204) – A radial hand sweeps clockwise from the nine o'clock position
	GstSmptetransitionTypePinwheelTbv GstSmptetransitionType = "pinwheel-tbv" // pinwheel-tbv (205) – Two radial hands sweep clockwise from the twelve and six o'clock positions
	GstSmptetransitionTypePinwheelTbh GstSmptetransitionType = "pinwheel-tbh" // pinwheel-tbh (206) – Two radial hands sweep clockwise from the nine and three o'clock positions
	GstSmptetransitionTypePinwheelFb GstSmptetransitionType = "pinwheel-fb" // pinwheel-fb (207) – Four radial hands sweep clockwise
	GstSmptetransitionTypeFanCt GstSmptetransitionType = "fan-ct" // fan-ct (211) – A fan unfolds from the top edge, the fan axis at the center
	GstSmptetransitionTypeFanCr GstSmptetransitionType = "fan-cr" // fan-cr (212) – A fan unfolds from the right edge, the fan axis at the center
	GstSmptetransitionTypeDoublefanFov GstSmptetransitionType = "doublefan-fov" // doublefan-fov (213) – Two fans, their axes at the center, unfold from the top and bottom
	GstSmptetransitionTypeDoublefanFoh GstSmptetransitionType = "doublefan-foh" // doublefan-foh (214) – Two fans, their axes at the center, unfold from the left and right
	GstSmptetransitionTypeSinglesweepCwt GstSmptetransitionType = "singlesweep-cwt" // singlesweep-cwt (221) – A radial hand sweeps clockwise from the top edge's midpoint
	GstSmptetransitionTypeSinglesweepCwr GstSmptetransitionType = "singlesweep-cwr" // singlesweep-cwr (222) – A radial hand sweeps clockwise from the right edge's midpoint
	GstSmptetransitionTypeSinglesweepCwb GstSmptetransitionType = "singlesweep-cwb" // singlesweep-cwb (223) – A radial hand sweeps clockwise from the bottom edge's midpoint
	GstSmptetransitionTypeSinglesweepCwl GstSmptetransitionType = "singlesweep-cwl" // singlesweep-cwl (224) – A radial hand sweeps clockwise from the left edge's midpoint
	GstSmptetransitionTypeDoublesweepPv GstSmptetransitionType = "doublesweep-pv" // doublesweep-pv (225) – Two radial hands sweep clockwise and counter-clockwise from the top and bottom edges' midpoints
	GstSmptetransitionTypeDoublesweepPd GstSmptetransitionType = "doublesweep-pd" // doublesweep-pd (226) – Two radial hands sweep clockwise and counter-clockwise from the left and right edges' midpoints
	GstSmptetransitionTypeDoublesweepOv GstSmptetransitionType = "doublesweep-ov" // doublesweep-ov (227) – Two radial hands attached at the top and bottom edges' midpoints sweep from right to left
	GstSmptetransitionTypeDoublesweepOh GstSmptetransitionType = "doublesweep-oh" // doublesweep-oh (228) – Two radial hands attached at the left and right edges' midpoints sweep from top to bottom
	GstSmptetransitionTypeFanT GstSmptetransitionType = "fan-t" // fan-t (231) – A fan unfolds from the bottom, the fan axis at the top edge's midpoint
	GstSmptetransitionTypeFanR GstSmptetransitionType = "fan-r" // fan-r (232) – A fan unfolds from the left, the fan axis at the right edge's midpoint
	GstSmptetransitionTypeFanB GstSmptetransitionType = "fan-b" // fan-b (233) – A fan unfolds from the top, the fan axis at the bottom edge's midpoint
	GstSmptetransitionTypeFanL GstSmptetransitionType = "fan-l" // fan-l (234) – A fan unfolds from the right, the fan axis at the left edge's midpoint
	GstSmptetransitionTypeDoublefanFiv GstSmptetransitionType = "doublefan-fiv" // doublefan-fiv (235) – Two fans, their axes at the top and bottom, unfold from the center
	GstSmptetransitionTypeDoublefanFih GstSmptetransitionType = "doublefan-fih" // doublefan-fih (236) – Two fans, their axes at the left and right, unfold from the center
	GstSmptetransitionTypeSinglesweepCwtl GstSmptetransitionType = "singlesweep-cwtl" // singlesweep-cwtl (241) – A radial hand sweeps clockwise from the upper-left corner
	GstSmptetransitionTypeSinglesweepCwbl GstSmptetransitionType = "singlesweep-cwbl" // singlesweep-cwbl (242) – A radial hand sweeps counter-clockwise from the lower-left corner.
	GstSmptetransitionTypeSinglesweepCwbr GstSmptetransitionType = "singlesweep-cwbr" // singlesweep-cwbr (243) – A radial hand sweeps clockwise from the lower-right corner
	GstSmptetransitionTypeSinglesweepCwtr GstSmptetransitionType = "singlesweep-cwtr" // singlesweep-cwtr (244) – A radial hand sweeps counter-clockwise from the upper-right corner
	GstSmptetransitionTypeDoublesweepPdtl GstSmptetransitionType = "doublesweep-pdtl" // doublesweep-pdtl (245) – Two radial hands attached at the upper-left and lower-right corners sweep down and up
	GstSmptetransitionTypeDoublesweepPdbl GstSmptetransitionType = "doublesweep-pdbl" // doublesweep-pdbl (246) – Two radial hands attached at the lower-left and upper-right corners sweep down and up
	GstSmptetransitionTypeSaloondoorT GstSmptetransitionType = "saloondoor-t" // saloondoor-t (251) – Two radial hands attached at the upper-left and upper-right corners sweep down
	GstSmptetransitionTypeSaloondoorL GstSmptetransitionType = "saloondoor-l" // saloondoor-l (252) – Two radial hands attached at the upper-left and lower-left corners sweep to the right
	GstSmptetransitionTypeSaloondoorB GstSmptetransitionType = "saloondoor-b" // saloondoor-b (253) – Two radial hands attached at the lower-left and lower-right corners sweep up
	GstSmptetransitionTypeSaloondoorR GstSmptetransitionType = "saloondoor-r" // saloondoor-r (254) – Two radial hands attached at the upper-right and lower-right corners sweep to the left
	GstSmptetransitionTypeWindshieldR GstSmptetransitionType = "windshield-r" // windshield-r (261) – Two radial hands attached at the midpoints of the top and bottom halves sweep from right to left
	GstSmptetransitionTypeWindshieldU GstSmptetransitionType = "windshield-u" // windshield-u (262) – Two radial hands attached at the midpoints of the left and right halves sweep from top to bottom
	GstSmptetransitionTypeWindshieldV GstSmptetransitionType = "windshield-v" // windshield-v (263) – Two sets of radial hands attached at the midpoints of the top and bottom halves sweep from top to bottom and bottom to top
	GstSmptetransitionTypeWindshieldH GstSmptetransitionType = "windshield-h" // windshield-h (264) – Two sets of radial hands attached at the midpoints of the left and right halves sweep from left to right and right to left
)

