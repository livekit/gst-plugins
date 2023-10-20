// source: gst-plugins-bad

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Use luma from an opaque stream as alpha plane on another
type GstAlphaCombine struct {
	*gst.Element
}

func NewAlphaCombine() (*GstAlphaCombine, error) {
	e, err := gst.NewElement("alphacombine")
	if err != nil {
		return nil, err
	}
	return &GstAlphaCombine{Element: e}, nil
}

func NewAlphaCombineWithName(name string) (*GstAlphaCombine, error) {
	e, err := gst.NewElementWithName("alphacombine", name)
	if err != nil {
		return nil, err
	}
	return &GstAlphaCombine{Element: e}, nil
}

// Extract and expose as a stream the CODEC alpha.
type GstCodecAlphaDemux struct {
	*gst.Element
}

func NewCodecAlphaDemux() (*GstCodecAlphaDemux, error) {
	e, err := gst.NewElement("codecalphademux")
	if err != nil {
		return nil, err
	}
	return &GstCodecAlphaDemux{Element: e}, nil
}

func NewCodecAlphaDemuxWithName(name string) (*GstCodecAlphaDemux, error) {
	e, err := gst.NewElementWithName("codecalphademux", name)
	if err != nil {
		return nil, err
	}
	return &GstCodecAlphaDemux{Element: e}, nil
}

// Wrapper bin to decode VP8 with alpha stream.
type GstVp8AlphaDecodeBin struct {
	*GstAlphaDecodeBin
}

func NewVp8AlphaDecodeBin() (*GstVp8AlphaDecodeBin, error) {
	e, err := gst.NewElement("vp8alphadecodebin")
	if err != nil {
		return nil, err
	}
	return &GstVp8AlphaDecodeBin{GstAlphaDecodeBin: &GstAlphaDecodeBin{Bin: &gst.Bin{Element: e}}}, nil
}

func NewVp8AlphaDecodeBinWithName(name string) (*GstVp8AlphaDecodeBin, error) {
	e, err := gst.NewElementWithName("vp8alphadecodebin", name)
	if err != nil {
		return nil, err
	}
	return &GstVp8AlphaDecodeBin{GstAlphaDecodeBin: &GstAlphaDecodeBin{Bin: &gst.Bin{Element: e}}}, nil
}

// Wrapper bin to decode VP9 with alpha stream.
type GstVp9AlphaDecodeBin struct {
	*GstAlphaDecodeBin
}

func NewVp9AlphaDecodeBin() (*GstVp9AlphaDecodeBin, error) {
	e, err := gst.NewElement("vp9alphadecodebin")
	if err != nil {
		return nil, err
	}
	return &GstVp9AlphaDecodeBin{GstAlphaDecodeBin: &GstAlphaDecodeBin{Bin: &gst.Bin{Element: e}}}, nil
}

func NewVp9AlphaDecodeBinWithName(name string) (*GstVp9AlphaDecodeBin, error) {
	e, err := gst.NewElementWithName("vp9alphadecodebin", name)
	if err != nil {
		return nil, err
	}
	return &GstVp9AlphaDecodeBin{GstAlphaDecodeBin: &GstAlphaDecodeBin{Bin: &gst.Bin{Element: e}}}, nil
}

type GstAlphaDecodeBin struct {
	*gst.Bin
}
