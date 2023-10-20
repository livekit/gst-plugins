// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Adaptive Multi-Rate Wideband audio encoder
type GstVoAmrWbEnc struct {
	*GstAudioEncoder
}

func NewVoAmrWbEnc() (*GstVoAmrWbEnc, error) {
	e, err := gst.NewElement("voamrwbenc")
	if err != nil {
		return nil, err
	}
	return &GstVoAmrWbEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewVoAmrWbEncWithName(name string) (*GstVoAmrWbEnc, error) {
	e, err := gst.NewElementWithName("voamrwbenc", name)
	if err != nil {
		return nil, err
	}
	return &GstVoAmrWbEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Encoding Band Mode (Kbps)
// Default: MR660 (0)
func (e *GstVoAmrWbEnc) SetBandMode(value GstVoAmrWbEncBandMode) error {
	e.Element.SetArg("band-mode", string(value))
	return nil
}

func (e *GstVoAmrWbEnc) GetBandMode() (GstVoAmrWbEncBandMode, error) {
	var value GstVoAmrWbEncBandMode
	var ok bool
	v, err := e.Element.GetProperty("band-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVoAmrWbEncBandMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVoAmrWbEncBandMode")
	}
	return value, nil
}

type GstVoAmrWbEncBandMode string

const (
	GstVoAmrWbEncBandModeMR660  GstVoAmrWbEncBandMode = "MR660"  // MR660
	GstVoAmrWbEncBandModeMR885  GstVoAmrWbEncBandMode = "MR885"  // MR885
	GstVoAmrWbEncBandModeMR1265 GstVoAmrWbEncBandMode = "MR1265" // MR1265
	GstVoAmrWbEncBandModeMR1425 GstVoAmrWbEncBandMode = "MR1425" // MR1425
	GstVoAmrWbEncBandModeMR1585 GstVoAmrWbEncBandMode = "MR1585" // MR1585
	GstVoAmrWbEncBandModeMR1825 GstVoAmrWbEncBandMode = "MR1825" // MR1825
	GstVoAmrWbEncBandModeMR1985 GstVoAmrWbEncBandMode = "MR1985" // MR1985
	GstVoAmrWbEncBandModeMR2305 GstVoAmrWbEncBandMode = "MR2305" // MR2305
	GstVoAmrWbEncBandModeMR2385 GstVoAmrWbEncBandMode = "MR2385" // MR2385
	GstVoAmrWbEncBandModeMRDTX  GstVoAmrWbEncBandMode = "MRDTX"  // MRDTX
)
