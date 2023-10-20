// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes packets with SCTP
type GstSctpDec struct {
	*gst.Element
}

func NewSctpDec() (*GstSctpDec, error) {
	e, err := gst.NewElement("sctpdec")
	if err != nil {
		return nil, err
	}
	return &GstSctpDec{Element: e}, nil
}

func NewSctpDecWithName(name string) (*GstSctpDec, error) {
	e, err := gst.NewElementWithName("sctpdec", name)
	if err != nil {
		return nil, err
	}
	return &GstSctpDec{Element: e}, nil
}

// Local sctp port for the sctp association. The remote port is configured via the GstSctpEnc element.
// Default: 0, Min: 0, Max: 65535
func (e *GstSctpDec) SetLocalSctpPort(value uint) error {
	return e.Element.SetProperty("local-sctp-port", value)
}

func (e *GstSctpDec) GetLocalSctpPort() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("local-sctp-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Every encoder/decoder pair should have the same, unique, sctp-association-id. This value must be set before any pads are requested.
// Default: 1, Min: 0, Max: 65535
func (e *GstSctpDec) SetSctpAssociationId(value uint) error {
	return e.Element.SetProperty("sctp-association-id", value)
}

func (e *GstSctpDec) GetSctpAssociationId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sctp-association-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Encodes packets with SCTP
type GstSctpEnc struct {
	*gst.Element
}

func NewSctpEnc() (*GstSctpEnc, error) {
	e, err := gst.NewElement("sctpenc")
	if err != nil {
		return nil, err
	}
	return &GstSctpEnc{Element: e}, nil
}

func NewSctpEncWithName(name string) (*GstSctpEnc, error) {
	e, err := gst.NewElementWithName("sctpenc", name)
	if err != nil {
		return nil, err
	}
	return &GstSctpEnc{Element: e}, nil
}

// Sctp remote sctp port for the sctp association. The local port is configured via the GstSctpDec element.
// Default: 0, Min: 0, Max: 65535
func (e *GstSctpEnc) SetRemoteSctpPort(value uint) error {
	return e.Element.SetProperty("remote-sctp-port", value)
}

func (e *GstSctpEnc) GetRemoteSctpPort() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("remote-sctp-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Every encoder/decoder pair should have the same, unique, sctp-association-id. This value must be set before any pads are requested.
// Default: 1, Min: 0, Max: -1
func (e *GstSctpEnc) SetSctpAssociationId(value uint) error {
	return e.Element.SetProperty("sctp-association-id", value)
}

func (e *GstSctpEnc) GetSctpAssociationId() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sctp-association-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// When set to TRUE, a sequenced, reliable, connection-based connection is used.When TRUE the partial reliability parameters of the channel are ignored.
// Default: false
func (e *GstSctpEnc) SetUseSockStream(value bool) error {
	return e.Element.SetProperty("use-sock-stream", value)
}

func (e *GstSctpEnc) GetUseSockStream() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-sock-stream")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
