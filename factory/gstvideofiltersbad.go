// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Visualize differences between adjacent video frames
type GstVideoDiff struct {
	*GstVideoFilter
}

func NewVideoDiff() (*GstVideoDiff, error) {
	e, err := gst.NewElement("videodiff")
	if err != nil {
		return nil, err
	}
	return &GstVideoDiff{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewVideoDiffWithName(name string) (*GstVideoDiff, error) {
	e, err := gst.NewElementWithName("videodiff", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoDiff{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Overlays zebra striping on overexposed areas of video
type GstZebraStripe struct {
	*GstVideoFilter
}

func NewZebraStripe() (*GstZebraStripe, error) {
	e, err := gst.NewElement("zebrastripe")
	if err != nil {
		return nil, err
	}
	return &GstZebraStripe{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewZebraStripeWithName(name string) (*GstZebraStripe, error) {
	e, err := gst.NewElementWithName("zebrastripe", name)
	if err != nil {
		return nil, err
	}
	return &GstZebraStripe{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Threshold above which the video is striped
// Default: 90, Min: 0, Max: 100
func (e *GstZebraStripe) SetThreshold(value int) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstZebraStripe) GetThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Detects scene changes in video
type GstSceneChange struct {
	*GstVideoFilter
}

func NewSceneChange() (*GstSceneChange, error) {
	e, err := gst.NewElement("scenechange")
	if err != nil {
		return nil, err
	}
	return &GstSceneChange{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewSceneChangeWithName(name string) (*GstSceneChange, error) {
	e, err := gst.NewElementWithName("scenechange", name)
	if err != nil {
		return nil, err
	}
	return &GstSceneChange{GstVideoFilter: &GstVideoFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}
