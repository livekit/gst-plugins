// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Apply the standard SMPTE transitions on video images
type GstSMPTE struct {
	*gst.Element
}

func NewSMPTE() (*GstSMPTE, error) {
	e, err := gst.NewElement("smpte")
	if err != nil {
		return nil, err
	}
	return &GstSMPTE{Element: e}, nil
}

func NewSMPTEWithName(name string) (*GstSMPTE, error) {
	e, err := gst.NewElementWithName("smpte", name)
	if err != nil {
		return nil, err
	}
	return &GstSMPTE{Element: e}, nil
}

// The border width of the transition
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSMPTE) SetBorder(value int) error {
	return e.Element.SetProperty("border", value)
}

func (e *GstSMPTE) GetBorder() (int, error) {
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

// Depth of the mask in bits
// Default: 16, Min: 1, Max: 24
func (e *GstSMPTE) SetDepth(value int) error {
	return e.Element.SetProperty("depth", value)
}

func (e *GstSMPTE) GetDepth() (int, error) {
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

// Duration of the transition effect in nanoseconds
// Default: 1000000000, Min: 0, Max: 18446744073709551615
func (e *GstSMPTE) SetDuration(value uint64) error {
	return e.Element.SetProperty("duration", value)
}

func (e *GstSMPTE) GetDuration() (uint64, error) {
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

// Invert transition mask
// Default: false
func (e *GstSMPTE) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

func (e *GstSMPTE) GetInvert() (bool, error) {
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

// The type of transition to use
// Default: bar-wipe-lr (1)
func (e *GstSMPTE) SetType(value GstSMPTETransitionType) error {
	e.Element.SetArg("type", string(value))
	return nil
}

func (e *GstSMPTE) GetType() (GstSMPTETransitionType, error) {
	var value GstSMPTETransitionType
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSMPTETransitionType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSMPTETransitionType")
	}
	return value, nil
}

// Apply the standard SMPTE transitions as alpha on video images
type GstSMPTEAlpha struct {
	*GstVideoFilter
}

func NewSMPTEAlpha() (*GstSMPTEAlpha, error) {
	e, err := gst.NewElement("smptealpha")
	if err != nil {
		return nil, err
	}
	return &GstSMPTEAlpha{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewSMPTEAlphaWithName(name string) (*GstSMPTEAlpha, error) {
	e, err := gst.NewElementWithName("smptealpha", name)
	if err != nil {
		return nil, err
	}
	return &GstSMPTEAlpha{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// The border width of the transition
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSMPTEAlpha) SetBorder(value int) error {
	return e.Element.SetProperty("border", value)
}

func (e *GstSMPTEAlpha) GetBorder() (int, error) {
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

// Depth of the mask in bits
// Default: 16, Min: 1, Max: 24
func (e *GstSMPTEAlpha) SetDepth(value int) error {
	return e.Element.SetProperty("depth", value)
}

func (e *GstSMPTEAlpha) GetDepth() (int, error) {
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

// Invert transition mask
// Default: false
func (e *GstSMPTEAlpha) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

func (e *GstSMPTEAlpha) GetInvert() (bool, error) {
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

// Position of the transition effect
// Default: 0, Min: 0, Max: 1
func (e *GstSMPTEAlpha) SetPosition(value float64) error {
	return e.Element.SetProperty("position", value)
}

func (e *GstSMPTEAlpha) GetPosition() (float64, error) {
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

// The type of transition to use
// Default: bar-wipe-lr (1)
func (e *GstSMPTEAlpha) SetType(value GstSMPTEAlphaTransitionType) error {
	e.Element.SetArg("type", string(value))
	return nil
}

func (e *GstSMPTEAlpha) GetType() (GstSMPTEAlphaTransitionType, error) {
	var value GstSMPTEAlphaTransitionType
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSMPTEAlphaTransitionType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSMPTEAlphaTransitionType")
	}
	return value, nil
}

type GstSMPTEAlphaTransitionType string

const (
	GstSMPTEAlphaTransitionTypeBarWipeLr       GstSMPTEAlphaTransitionType = "bar-wipe-lr"       // A bar moves from left to right
	GstSMPTEAlphaTransitionTypeBarWipeTb       GstSMPTEAlphaTransitionType = "bar-wipe-tb"       // A bar moves from top to bottom
	GstSMPTEAlphaTransitionTypeBoxWipeTl       GstSMPTEAlphaTransitionType = "box-wipe-tl"       // A box expands from the upper-left corner to the lower-right corner
	GstSMPTEAlphaTransitionTypeBoxWipeTr       GstSMPTEAlphaTransitionType = "box-wipe-tr"       // A box expands from the upper-right corner to the lower-left corner
	GstSMPTEAlphaTransitionTypeBoxWipeBr       GstSMPTEAlphaTransitionType = "box-wipe-br"       // A box expands from the lower-right corner to the upper-left corner
	GstSMPTEAlphaTransitionTypeBoxWipeBl       GstSMPTEAlphaTransitionType = "box-wipe-bl"       // A box expands from the lower-left corner to the upper-right corner
	GstSMPTEAlphaTransitionTypeFourBoxWipeCi   GstSMPTEAlphaTransitionType = "four-box-wipe-ci"  // A box shape expands from each of the four corners toward the center
	GstSMPTEAlphaTransitionTypeFourBoxWipeCo   GstSMPTEAlphaTransitionType = "four-box-wipe-co"  // A box shape expands from the center of each quadrant toward the corners of each quadrant
	GstSMPTEAlphaTransitionTypeBarndoorV       GstSMPTEAlphaTransitionType = "barndoor-v"        // A central, vertical line splits and expands toward the left and right edges
	GstSMPTEAlphaTransitionTypeBarndoorH       GstSMPTEAlphaTransitionType = "barndoor-h"        // A central, horizontal line splits and expands toward the top and bottom edges
	GstSMPTEAlphaTransitionTypeBoxWipeTc       GstSMPTEAlphaTransitionType = "box-wipe-tc"       // A box expands from the top edge's midpoint to the bottom corners
	GstSMPTEAlphaTransitionTypeBoxWipeRc       GstSMPTEAlphaTransitionType = "box-wipe-rc"       // A box expands from the right edge's midpoint to the left corners
	GstSMPTEAlphaTransitionTypeBoxWipeBc       GstSMPTEAlphaTransitionType = "box-wipe-bc"       // A box expands from the bottom edge's midpoint to the top corners
	GstSMPTEAlphaTransitionTypeBoxWipeLc       GstSMPTEAlphaTransitionType = "box-wipe-lc"       // A box expands from the left edge's midpoint to the right corners
	GstSMPTEAlphaTransitionTypeDiagonalTl      GstSMPTEAlphaTransitionType = "diagonal-tl"       // A diagonal line moves from the upper-left corner to the lower-right corner
	GstSMPTEAlphaTransitionTypeDiagonalTr      GstSMPTEAlphaTransitionType = "diagonal-tr"       // A diagonal line moves from the upper right corner to the lower-left corner
	GstSMPTEAlphaTransitionTypeBowtieV         GstSMPTEAlphaTransitionType = "bowtie-v"          // Two wedge shapes slide in from the top and bottom edges toward the center
	GstSMPTEAlphaTransitionTypeBowtieH         GstSMPTEAlphaTransitionType = "bowtie-h"          // Two wedge shapes slide in from the left and right edges toward the center
	GstSMPTEAlphaTransitionTypeBarndoorDbl     GstSMPTEAlphaTransitionType = "barndoor-dbl"      // A diagonal line from the lower-left to upper-right corners splits and expands toward the opposite corners
	GstSMPTEAlphaTransitionTypeBarndoorDtl     GstSMPTEAlphaTransitionType = "barndoor-dtl"      // A diagonal line from upper-left to lower-right corners splits and expands toward the opposite corners
	GstSMPTEAlphaTransitionTypeMiscDiagonalDbd GstSMPTEAlphaTransitionType = "misc-diagonal-dbd" // Four wedge shapes split from the center and retract toward the four edges
	GstSMPTEAlphaTransitionTypeMiscDiagonalDd  GstSMPTEAlphaTransitionType = "misc-diagonal-dd"  // A diamond connecting the four edge midpoints simultaneously contracts toward the center and expands toward the edges
	GstSMPTEAlphaTransitionTypeVeeD            GstSMPTEAlphaTransitionType = "vee-d"             // A wedge shape moves from top to bottom
	GstSMPTEAlphaTransitionTypeVeeL            GstSMPTEAlphaTransitionType = "vee-l"             // A wedge shape moves from right to left
	GstSMPTEAlphaTransitionTypeVeeU            GstSMPTEAlphaTransitionType = "vee-u"             // A wedge shape moves from bottom to top
	GstSMPTEAlphaTransitionTypeVeeR            GstSMPTEAlphaTransitionType = "vee-r"             // A wedge shape moves from left to right
	GstSMPTEAlphaTransitionTypeBarnveeD        GstSMPTEAlphaTransitionType = "barnvee-d"         // A 'V' shape extending from the bottom edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSMPTEAlphaTransitionTypeBarnveeL        GstSMPTEAlphaTransitionType = "barnvee-l"         // A 'V' shape extending from the left edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSMPTEAlphaTransitionTypeBarnveeU        GstSMPTEAlphaTransitionType = "barnvee-u"         // A 'V' shape extending from the top edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSMPTEAlphaTransitionTypeBarnveeR        GstSMPTEAlphaTransitionType = "barnvee-r"         // A 'V' shape extending from the right edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSMPTEAlphaTransitionTypeIrisRect        GstSMPTEAlphaTransitionType = "iris-rect"         // A rectangle expands from the center.
	GstSMPTEAlphaTransitionTypeClockCw12       GstSMPTEAlphaTransitionType = "clock-cw12"        // A radial hand sweeps clockwise from the twelve o'clock position
	GstSMPTEAlphaTransitionTypeClockCw3        GstSMPTEAlphaTransitionType = "clock-cw3"         // A radial hand sweeps clockwise from the three o'clock position
	GstSMPTEAlphaTransitionTypeClockCw6        GstSMPTEAlphaTransitionType = "clock-cw6"         // A radial hand sweeps clockwise from the six o'clock position
	GstSMPTEAlphaTransitionTypeClockCw9        GstSMPTEAlphaTransitionType = "clock-cw9"         // A radial hand sweeps clockwise from the nine o'clock position
	GstSMPTEAlphaTransitionTypePinwheelTbv     GstSMPTEAlphaTransitionType = "pinwheel-tbv"      // Two radial hands sweep clockwise from the twelve and six o'clock positions
	GstSMPTEAlphaTransitionTypePinwheelTbh     GstSMPTEAlphaTransitionType = "pinwheel-tbh"      // Two radial hands sweep clockwise from the nine and three o'clock positions
	GstSMPTEAlphaTransitionTypePinwheelFb      GstSMPTEAlphaTransitionType = "pinwheel-fb"       // Four radial hands sweep clockwise
	GstSMPTEAlphaTransitionTypeFanCt           GstSMPTEAlphaTransitionType = "fan-ct"            // A fan unfolds from the top edge, the fan axis at the center
	GstSMPTEAlphaTransitionTypeFanCr           GstSMPTEAlphaTransitionType = "fan-cr"            // A fan unfolds from the right edge, the fan axis at the center
	GstSMPTEAlphaTransitionTypeDoublefanFov    GstSMPTEAlphaTransitionType = "doublefan-fov"     // Two fans, their axes at the center, unfold from the top and bottom
	GstSMPTEAlphaTransitionTypeDoublefanFoh    GstSMPTEAlphaTransitionType = "doublefan-foh"     // Two fans, their axes at the center, unfold from the left and right
	GstSMPTEAlphaTransitionTypeSinglesweepCwt  GstSMPTEAlphaTransitionType = "singlesweep-cwt"   // A radial hand sweeps clockwise from the top edge's midpoint
	GstSMPTEAlphaTransitionTypeSinglesweepCwr  GstSMPTEAlphaTransitionType = "singlesweep-cwr"   // A radial hand sweeps clockwise from the right edge's midpoint
	GstSMPTEAlphaTransitionTypeSinglesweepCwb  GstSMPTEAlphaTransitionType = "singlesweep-cwb"   // A radial hand sweeps clockwise from the bottom edge's midpoint
	GstSMPTEAlphaTransitionTypeSinglesweepCwl  GstSMPTEAlphaTransitionType = "singlesweep-cwl"   // A radial hand sweeps clockwise from the left edge's midpoint
	GstSMPTEAlphaTransitionTypeDoublesweepPv   GstSMPTEAlphaTransitionType = "doublesweep-pv"    // Two radial hands sweep clockwise and counter-clockwise from the top and bottom edges' midpoints
	GstSMPTEAlphaTransitionTypeDoublesweepPd   GstSMPTEAlphaTransitionType = "doublesweep-pd"    // Two radial hands sweep clockwise and counter-clockwise from the left and right edges' midpoints
	GstSMPTEAlphaTransitionTypeDoublesweepOv   GstSMPTEAlphaTransitionType = "doublesweep-ov"    // Two radial hands attached at the top and bottom edges' midpoints sweep from right to left
	GstSMPTEAlphaTransitionTypeDoublesweepOh   GstSMPTEAlphaTransitionType = "doublesweep-oh"    // Two radial hands attached at the left and right edges' midpoints sweep from top to bottom
	GstSMPTEAlphaTransitionTypeFanT            GstSMPTEAlphaTransitionType = "fan-t"             // A fan unfolds from the bottom, the fan axis at the top edge's midpoint
	GstSMPTEAlphaTransitionTypeFanR            GstSMPTEAlphaTransitionType = "fan-r"             // A fan unfolds from the left, the fan axis at the right edge's midpoint
	GstSMPTEAlphaTransitionTypeFanB            GstSMPTEAlphaTransitionType = "fan-b"             // A fan unfolds from the top, the fan axis at the bottom edge's midpoint
	GstSMPTEAlphaTransitionTypeFanL            GstSMPTEAlphaTransitionType = "fan-l"             // A fan unfolds from the right, the fan axis at the left edge's midpoint
	GstSMPTEAlphaTransitionTypeDoublefanFiv    GstSMPTEAlphaTransitionType = "doublefan-fiv"     // Two fans, their axes at the top and bottom, unfold from the center
	GstSMPTEAlphaTransitionTypeDoublefanFih    GstSMPTEAlphaTransitionType = "doublefan-fih"     // Two fans, their axes at the left and right, unfold from the center
	GstSMPTEAlphaTransitionTypeSinglesweepCwtl GstSMPTEAlphaTransitionType = "singlesweep-cwtl"  // A radial hand sweeps clockwise from the upper-left corner
	GstSMPTEAlphaTransitionTypeSinglesweepCwbl GstSMPTEAlphaTransitionType = "singlesweep-cwbl"  // A radial hand sweeps counter-clockwise from the lower-left corner.
	GstSMPTEAlphaTransitionTypeSinglesweepCwbr GstSMPTEAlphaTransitionType = "singlesweep-cwbr"  // A radial hand sweeps clockwise from the lower-right corner
	GstSMPTEAlphaTransitionTypeSinglesweepCwtr GstSMPTEAlphaTransitionType = "singlesweep-cwtr"  // A radial hand sweeps counter-clockwise from the upper-right corner
	GstSMPTEAlphaTransitionTypeDoublesweepPdtl GstSMPTEAlphaTransitionType = "doublesweep-pdtl"  // Two radial hands attached at the upper-left and lower-right corners sweep down and up
	GstSMPTEAlphaTransitionTypeDoublesweepPdbl GstSMPTEAlphaTransitionType = "doublesweep-pdbl"  // Two radial hands attached at the lower-left and upper-right corners sweep down and up
	GstSMPTEAlphaTransitionTypeSaloondoorT     GstSMPTEAlphaTransitionType = "saloondoor-t"      // Two radial hands attached at the upper-left and upper-right corners sweep down
	GstSMPTEAlphaTransitionTypeSaloondoorL     GstSMPTEAlphaTransitionType = "saloondoor-l"      // Two radial hands attached at the upper-left and lower-left corners sweep to the right
	GstSMPTEAlphaTransitionTypeSaloondoorB     GstSMPTEAlphaTransitionType = "saloondoor-b"      // Two radial hands attached at the lower-left and lower-right corners sweep up
	GstSMPTEAlphaTransitionTypeSaloondoorR     GstSMPTEAlphaTransitionType = "saloondoor-r"      // Two radial hands attached at the upper-right and lower-right corners sweep to the left
	GstSMPTEAlphaTransitionTypeWindshieldR     GstSMPTEAlphaTransitionType = "windshield-r"      // Two radial hands attached at the midpoints of the top and bottom halves sweep from right to left
	GstSMPTEAlphaTransitionTypeWindshieldU     GstSMPTEAlphaTransitionType = "windshield-u"      // Two radial hands attached at the midpoints of the left and right halves sweep from top to bottom
	GstSMPTEAlphaTransitionTypeWindshieldV     GstSMPTEAlphaTransitionType = "windshield-v"      // Two sets of radial hands attached at the midpoints of the top and bottom halves sweep from top to bottom and bottom to top
	GstSMPTEAlphaTransitionTypeWindshieldH     GstSMPTEAlphaTransitionType = "windshield-h"      // Two sets of radial hands attached at the midpoints of the left and right halves sweep from left to right and right to left
)

type GstSMPTETransitionType string

const (
	GstSMPTETransitionTypeBarWipeLr       GstSMPTETransitionType = "bar-wipe-lr"       // A bar moves from left to right
	GstSMPTETransitionTypeBarWipeTb       GstSMPTETransitionType = "bar-wipe-tb"       // A bar moves from top to bottom
	GstSMPTETransitionTypeBoxWipeTl       GstSMPTETransitionType = "box-wipe-tl"       // A box expands from the upper-left corner to the lower-right corner
	GstSMPTETransitionTypeBoxWipeTr       GstSMPTETransitionType = "box-wipe-tr"       // A box expands from the upper-right corner to the lower-left corner
	GstSMPTETransitionTypeBoxWipeBr       GstSMPTETransitionType = "box-wipe-br"       // A box expands from the lower-right corner to the upper-left corner
	GstSMPTETransitionTypeBoxWipeBl       GstSMPTETransitionType = "box-wipe-bl"       // A box expands from the lower-left corner to the upper-right corner
	GstSMPTETransitionTypeFourBoxWipeCi   GstSMPTETransitionType = "four-box-wipe-ci"  // A box shape expands from each of the four corners toward the center
	GstSMPTETransitionTypeFourBoxWipeCo   GstSMPTETransitionType = "four-box-wipe-co"  // A box shape expands from the center of each quadrant toward the corners of each quadrant
	GstSMPTETransitionTypeBarndoorV       GstSMPTETransitionType = "barndoor-v"        // A central, vertical line splits and expands toward the left and right edges
	GstSMPTETransitionTypeBarndoorH       GstSMPTETransitionType = "barndoor-h"        // A central, horizontal line splits and expands toward the top and bottom edges
	GstSMPTETransitionTypeBoxWipeTc       GstSMPTETransitionType = "box-wipe-tc"       // A box expands from the top edge's midpoint to the bottom corners
	GstSMPTETransitionTypeBoxWipeRc       GstSMPTETransitionType = "box-wipe-rc"       // A box expands from the right edge's midpoint to the left corners
	GstSMPTETransitionTypeBoxWipeBc       GstSMPTETransitionType = "box-wipe-bc"       // A box expands from the bottom edge's midpoint to the top corners
	GstSMPTETransitionTypeBoxWipeLc       GstSMPTETransitionType = "box-wipe-lc"       // A box expands from the left edge's midpoint to the right corners
	GstSMPTETransitionTypeDiagonalTl      GstSMPTETransitionType = "diagonal-tl"       // A diagonal line moves from the upper-left corner to the lower-right corner
	GstSMPTETransitionTypeDiagonalTr      GstSMPTETransitionType = "diagonal-tr"       // A diagonal line moves from the upper right corner to the lower-left corner
	GstSMPTETransitionTypeBowtieV         GstSMPTETransitionType = "bowtie-v"          // Two wedge shapes slide in from the top and bottom edges toward the center
	GstSMPTETransitionTypeBowtieH         GstSMPTETransitionType = "bowtie-h"          // Two wedge shapes slide in from the left and right edges toward the center
	GstSMPTETransitionTypeBarndoorDbl     GstSMPTETransitionType = "barndoor-dbl"      // A diagonal line from the lower-left to upper-right corners splits and expands toward the opposite corners
	GstSMPTETransitionTypeBarndoorDtl     GstSMPTETransitionType = "barndoor-dtl"      // A diagonal line from upper-left to lower-right corners splits and expands toward the opposite corners
	GstSMPTETransitionTypeMiscDiagonalDbd GstSMPTETransitionType = "misc-diagonal-dbd" // Four wedge shapes split from the center and retract toward the four edges
	GstSMPTETransitionTypeMiscDiagonalDd  GstSMPTETransitionType = "misc-diagonal-dd"  // A diamond connecting the four edge midpoints simultaneously contracts toward the center and expands toward the edges
	GstSMPTETransitionTypeVeeD            GstSMPTETransitionType = "vee-d"             // A wedge shape moves from top to bottom
	GstSMPTETransitionTypeVeeL            GstSMPTETransitionType = "vee-l"             // A wedge shape moves from right to left
	GstSMPTETransitionTypeVeeU            GstSMPTETransitionType = "vee-u"             // A wedge shape moves from bottom to top
	GstSMPTETransitionTypeVeeR            GstSMPTETransitionType = "vee-r"             // A wedge shape moves from left to right
	GstSMPTETransitionTypeBarnveeD        GstSMPTETransitionType = "barnvee-d"         // A 'V' shape extending from the bottom edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSMPTETransitionTypeBarnveeL        GstSMPTETransitionType = "barnvee-l"         // A 'V' shape extending from the left edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSMPTETransitionTypeBarnveeU        GstSMPTETransitionType = "barnvee-u"         // A 'V' shape extending from the top edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSMPTETransitionTypeBarnveeR        GstSMPTETransitionType = "barnvee-r"         // A 'V' shape extending from the right edge's midpoint to the opposite corners contracts toward the center and expands toward the edges
	GstSMPTETransitionTypeIrisRect        GstSMPTETransitionType = "iris-rect"         // A rectangle expands from the center.
	GstSMPTETransitionTypeClockCw12       GstSMPTETransitionType = "clock-cw12"        // A radial hand sweeps clockwise from the twelve o'clock position
	GstSMPTETransitionTypeClockCw3        GstSMPTETransitionType = "clock-cw3"         // A radial hand sweeps clockwise from the three o'clock position
	GstSMPTETransitionTypeClockCw6        GstSMPTETransitionType = "clock-cw6"         // A radial hand sweeps clockwise from the six o'clock position
	GstSMPTETransitionTypeClockCw9        GstSMPTETransitionType = "clock-cw9"         // A radial hand sweeps clockwise from the nine o'clock position
	GstSMPTETransitionTypePinwheelTbv     GstSMPTETransitionType = "pinwheel-tbv"      // Two radial hands sweep clockwise from the twelve and six o'clock positions
	GstSMPTETransitionTypePinwheelTbh     GstSMPTETransitionType = "pinwheel-tbh"      // Two radial hands sweep clockwise from the nine and three o'clock positions
	GstSMPTETransitionTypePinwheelFb      GstSMPTETransitionType = "pinwheel-fb"       // Four radial hands sweep clockwise
	GstSMPTETransitionTypeFanCt           GstSMPTETransitionType = "fan-ct"            // A fan unfolds from the top edge, the fan axis at the center
	GstSMPTETransitionTypeFanCr           GstSMPTETransitionType = "fan-cr"            // A fan unfolds from the right edge, the fan axis at the center
	GstSMPTETransitionTypeDoublefanFov    GstSMPTETransitionType = "doublefan-fov"     // Two fans, their axes at the center, unfold from the top and bottom
	GstSMPTETransitionTypeDoublefanFoh    GstSMPTETransitionType = "doublefan-foh"     // Two fans, their axes at the center, unfold from the left and right
	GstSMPTETransitionTypeSinglesweepCwt  GstSMPTETransitionType = "singlesweep-cwt"   // A radial hand sweeps clockwise from the top edge's midpoint
	GstSMPTETransitionTypeSinglesweepCwr  GstSMPTETransitionType = "singlesweep-cwr"   // A radial hand sweeps clockwise from the right edge's midpoint
	GstSMPTETransitionTypeSinglesweepCwb  GstSMPTETransitionType = "singlesweep-cwb"   // A radial hand sweeps clockwise from the bottom edge's midpoint
	GstSMPTETransitionTypeSinglesweepCwl  GstSMPTETransitionType = "singlesweep-cwl"   // A radial hand sweeps clockwise from the left edge's midpoint
	GstSMPTETransitionTypeDoublesweepPv   GstSMPTETransitionType = "doublesweep-pv"    // Two radial hands sweep clockwise and counter-clockwise from the top and bottom edges' midpoints
	GstSMPTETransitionTypeDoublesweepPd   GstSMPTETransitionType = "doublesweep-pd"    // Two radial hands sweep clockwise and counter-clockwise from the left and right edges' midpoints
	GstSMPTETransitionTypeDoublesweepOv   GstSMPTETransitionType = "doublesweep-ov"    // Two radial hands attached at the top and bottom edges' midpoints sweep from right to left
	GstSMPTETransitionTypeDoublesweepOh   GstSMPTETransitionType = "doublesweep-oh"    // Two radial hands attached at the left and right edges' midpoints sweep from top to bottom
	GstSMPTETransitionTypeFanT            GstSMPTETransitionType = "fan-t"             // A fan unfolds from the bottom, the fan axis at the top edge's midpoint
	GstSMPTETransitionTypeFanR            GstSMPTETransitionType = "fan-r"             // A fan unfolds from the left, the fan axis at the right edge's midpoint
	GstSMPTETransitionTypeFanB            GstSMPTETransitionType = "fan-b"             // A fan unfolds from the top, the fan axis at the bottom edge's midpoint
	GstSMPTETransitionTypeFanL            GstSMPTETransitionType = "fan-l"             // A fan unfolds from the right, the fan axis at the left edge's midpoint
	GstSMPTETransitionTypeDoublefanFiv    GstSMPTETransitionType = "doublefan-fiv"     // Two fans, their axes at the top and bottom, unfold from the center
	GstSMPTETransitionTypeDoublefanFih    GstSMPTETransitionType = "doublefan-fih"     // Two fans, their axes at the left and right, unfold from the center
	GstSMPTETransitionTypeSinglesweepCwtl GstSMPTETransitionType = "singlesweep-cwtl"  // A radial hand sweeps clockwise from the upper-left corner
	GstSMPTETransitionTypeSinglesweepCwbl GstSMPTETransitionType = "singlesweep-cwbl"  // A radial hand sweeps counter-clockwise from the lower-left corner.
	GstSMPTETransitionTypeSinglesweepCwbr GstSMPTETransitionType = "singlesweep-cwbr"  // A radial hand sweeps clockwise from the lower-right corner
	GstSMPTETransitionTypeSinglesweepCwtr GstSMPTETransitionType = "singlesweep-cwtr"  // A radial hand sweeps counter-clockwise from the upper-right corner
	GstSMPTETransitionTypeDoublesweepPdtl GstSMPTETransitionType = "doublesweep-pdtl"  // Two radial hands attached at the upper-left and lower-right corners sweep down and up
	GstSMPTETransitionTypeDoublesweepPdbl GstSMPTETransitionType = "doublesweep-pdbl"  // Two radial hands attached at the lower-left and upper-right corners sweep down and up
	GstSMPTETransitionTypeSaloondoorT     GstSMPTETransitionType = "saloondoor-t"      // Two radial hands attached at the upper-left and upper-right corners sweep down
	GstSMPTETransitionTypeSaloondoorL     GstSMPTETransitionType = "saloondoor-l"      // Two radial hands attached at the upper-left and lower-left corners sweep to the right
	GstSMPTETransitionTypeSaloondoorB     GstSMPTETransitionType = "saloondoor-b"      // Two radial hands attached at the lower-left and lower-right corners sweep up
	GstSMPTETransitionTypeSaloondoorR     GstSMPTETransitionType = "saloondoor-r"      // Two radial hands attached at the upper-right and lower-right corners sweep to the left
	GstSMPTETransitionTypeWindshieldR     GstSMPTETransitionType = "windshield-r"      // Two radial hands attached at the midpoints of the top and bottom halves sweep from right to left
	GstSMPTETransitionTypeWindshieldU     GstSMPTETransitionType = "windshield-u"      // Two radial hands attached at the midpoints of the left and right halves sweep from top to bottom
	GstSMPTETransitionTypeWindshieldV     GstSMPTETransitionType = "windshield-v"      // Two sets of radial hands attached at the midpoints of the top and bottom halves sweep from top to bottom and bottom to top
	GstSMPTETransitionTypeWindshieldH     GstSMPTETransitionType = "windshield-h"      // Two sets of radial hands attached at the midpoints of the left and right halves sweep from left to right and right to left
)
