// source: gst-plugins-ugly

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Decodes DVD subtitles into AYUV video frames
type GstDvdSubDec struct {
	*gst.Element
}

func NewDvdSubDec() (*GstDvdSubDec, error) {
	e, err := gst.NewElement("dvdsubdec")
	if err != nil {
		return nil, err
	}
	return &GstDvdSubDec{Element: e}, nil
}

func NewDvdSubDecWithName(name string) (*GstDvdSubDec, error) {
	e, err := gst.NewElementWithName("dvdsubdec", name)
	if err != nil {
		return nil, err
	}
	return &GstDvdSubDec{Element: e}, nil
}

// Parses and packetizes DVD subtitle streams
type GstDvdSubParse struct {
	*gst.Element
}

func NewDvdSubParse() (*GstDvdSubParse, error) {
	e, err := gst.NewElement("dvdsubparse")
	if err != nil {
		return nil, err
	}
	return &GstDvdSubParse{Element: e}, nil
}

func NewDvdSubParseWithName(name string) (*GstDvdSubParse, error) {
	e, err := gst.NewElementWithName("dvdsubparse", name)
	if err != nil {
		return nil, err
	}
	return &GstDvdSubParse{Element: e}, nil
}
