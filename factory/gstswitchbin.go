// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Switch between sub-pipelines (paths) based on input caps
type GstSwitchBin struct {
	*gst.Bin
}

func NewSwitchBin() (*GstSwitchBin, error) {
	e, err := gst.NewElement("switchbin")
	if err != nil {
		return nil, err
	}
	return &GstSwitchBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewSwitchBinWithName(name string) (*GstSwitchBin, error) {
	e, err := gst.NewElementWithName("switchbin", name)
	if err != nil {
		return nil, err
	}
	return &GstSwitchBin{Bin: &gst.Bin{Element: e}}, nil
}

// Currently selected path
// Default: -1, Min: 0, Max: -1
func (e *GstSwitchBin) GetCurrentPath() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of paths
// Default: 0, Min: 0, Max: -2
func (e *GstSwitchBin) SetNumPaths(value uint) error {
	return e.Element.SetProperty("num-paths", value)
}

func (e *GstSwitchBin) GetNumPaths() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-paths")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}
