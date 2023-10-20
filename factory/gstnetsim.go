// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// An element that simulates network jitter, packet loss and packet duplication
type GstNetSim struct {
	*gst.Element
}

func NewNetSim() (*GstNetSim, error) {
	e, err := gst.NewElement("netsim")
	if err != nil {
		return nil, err
	}
	return &GstNetSim{Element: e}, nil
}

func NewNetSimWithName(name string) (*GstNetSim, error) {
	e, err := gst.NewElementWithName("netsim", name)
	if err != nil {
		return nil, err
	}
	return &GstNetSim{Element: e}, nil
}

// The Probability a buffer is delayed
// Default: 0, Min: 0, Max: 1
func (e *GstNetSim) SetDelayProbability(value float32) error {
	return e.Element.SetProperty("delay-probability", value)
}

func (e *GstNetSim) GetDelayProbability() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("delay-probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Drop the next n packets
// Default: 0, Min: 0, Max: -1
func (e *GstNetSim) SetDropPackets(value uint) error {
	return e.Element.SetProperty("drop-packets", value)
}

func (e *GstNetSim) GetDropPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("drop-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The Probability a buffer is dropped
// Default: 0, Min: 0, Max: 1
func (e *GstNetSim) SetDropProbability(value float32) error {
	return e.Element.SetProperty("drop-probability", value)
}

func (e *GstNetSim) GetDropProbability() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("drop-probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The maximum delay (inclusive) in ms to apply to buffers
// Default: 400, Min: -2147483648, Max: 2147483647
func (e *GstNetSim) SetMaxDelay(value int) error {
	return e.Element.SetProperty("max-delay", value)
}

func (e *GstNetSim) GetMaxDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The minimum delay in ms to apply to buffers
// Default: 200, Min: -2147483648, Max: 2147483647
func (e *GstNetSim) SetMinDelay(value int) error {
	return e.Element.SetProperty("min-delay", value)
}

func (e *GstNetSim) GetMinDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// When delaying packets, are they allowed to be reordered or not
// Default: true
func (e *GstNetSim) SetAllowReordering(value bool) error {
	return e.Element.SetProperty("allow-reordering", value)
}

func (e *GstNetSim) GetAllowReordering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-reordering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Distribution for the amount of delay
// Default: uniform (0)
func (e *GstNetSim) SetDelayDistribution(value GstNetSimDistribution) error {
	e.Element.SetArg("delay-distribution", string(value))
	return nil
}

func (e *GstNetSim) GetDelayDistribution() (GstNetSimDistribution, error) {
	var value GstNetSimDistribution
	var ok bool
	v, err := e.Element.GetProperty("delay-distribution")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNetSimDistribution)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNetSimDistribution")
	}
	return value, nil
}

// The maximum number of kilobits to let through per second (-1 = unlimited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstNetSim) SetMaxKbps(value int) error {
	return e.Element.SetProperty("max-kbps", value)
}

func (e *GstNetSim) GetMaxKbps() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-kbps")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The Probability a buffer is duplicated
// Default: 0, Min: 0, Max: 1
func (e *GstNetSim) SetDuplicateProbability(value float32) error {
	return e.Element.SetProperty("duplicate-probability", value)
}

func (e *GstNetSim) GetDuplicateProbability() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("duplicate-probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// The size of the token bucket, related to burstiness resilience (-1 = unlimited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstNetSim) SetMaxBucketSize(value int) error {
	return e.Element.SetProperty("max-bucket-size", value)
}

func (e *GstNetSim) GetMaxBucketSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-bucket-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstNetSimDistribution string

const (
	GstNetSimDistributionUniform GstNetSimDistribution = "uniform" // uniform
	GstNetSimDistributionNormal  GstNetSimDistribution = "normal"  // normal
	GstNetSimDistributionGamma   GstNetSimDistribution = "gamma"   // gamma
)
