package factory

import (
	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

type RTPSession interface{}
type GstX265Tune string

type GstTagMux struct {
	*gst.Element
}

type GstTagDemux struct {
	*gst.Element
}

type GstRTPBasePayload struct {
	*gst.Element
}

type GstRTPBaseDepayload struct {
	*gst.Element
}

type GstRTPBaseAudioPayload struct {
	*GstRTPBasePayload
}

type GstRTPHeaderExtension struct {
	*gst.Element
}

type GstBaseCameraSrc struct {
	*gst.Bin
}

type GstVulkanVideoFilter struct {
	*base.GstBaseTransform
}

type GstAdaptiveDemux struct {
	*gst.Bin
}

type GstOpencvVideoFilter struct {
	*GstVideoFilter
}

type GstDxvaAV1Decoder struct {
	*GstAV1Decoder
}

type GstAV1Decoder struct {
	*GstVideoDecoder
}

type GstDxvaH264Decoder struct {
	*GstH264Decoder
}

type GstH264Decoder struct {
	*GstVideoDecoder
}

type GstDxvaH265Decoder struct {
	*GstH265Decoder
}

type GstH265Decoder struct {
	*GstVideoDecoder
}

type GstDxvaVp9Decoder struct {
	*GstVp9Decoder
}

type GstVp9Decoder struct {
	*GstVideoDecoder
}

type GstDxvaVp8Decoder struct {
	*GstVp8Decoder
}

type GstVp8Decoder struct {
	*GstVideoDecoder
}

type GstDxvaMpeg2Decoder struct {
	*GstMpeg2Decoder
}

type GstMpeg2Decoder struct {
	*GstVideoDecoder
}

type GstNonstreamAudioDecoder struct {
	*gst.Element
}

// Base

// gstreamer/subprojects/gstreamer/libs/gst/base/gstaggregator.c
type GstAggregator struct {
	*gst.Element
}

// gstreamer/subprojects/gstreamer/libs/gst/base/gstbaseparse.c
type GstBaseParse struct {
	*gst.Element
}

// Audio

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudioaggregator.c
type GstAudioAggregator struct {
	*GstAggregator
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudiobasesink.c
type GstAudioBaseSink struct {
	*base.GstBaseSink
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudiobasesrc.c
type GstAudioBaseSrc struct {
	*base.GstPushSrc
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudiocdsrc.c
type GstAudioCdSrc struct {
	*base.GstPushSrc
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudiodecoder.c
type GstAudioDecoder struct {
	*gst.Element
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudioencoder.c
type GstAudioEncoder struct {
	*gst.Element
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudiofilter.c
type GstAudioFilter struct {
	*base.GstBaseTransform
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudiosink.c
type GstAudioSink struct {
	*GstAudioBaseSink
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/audio/gstaudiosrc.c
type GstAudioSrc struct {
	*GstAudioBaseSrc
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/pbutils/gstaudiovisualizer.c
type GstAudioVisualizer struct {
	*gst.Element
}

// Video

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/video/gstvideoaggregator.c
type GstVideoAggregator struct {
	*GstAggregator
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/video/gstvideodecoder.c
type GstVideoDecoder struct {
	*gst.Element
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/video/gstvideoencoder.c
type GstVideoEncoder struct {
	*gst.Element
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/video/gstvideofilter.c
type GstVideoFilter struct {
	*base.GstBaseTransform
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/video/gstvideosink.c
type GstVideoSink struct {
	*base.GstBaseSink
}

// GL

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/gl/gstglbasefilter.c
type GstGLBaseFilter struct {
	*base.GstBaseTransform
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/gl/gstglbasemixer.c
type GstGLBaseMixer struct {
	*GstVideoAggregator
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/gl/gstglbasesrc.c
type GstGLBaseSrc struct {
	*base.GstPushSrc
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/gl/gstglfilter.c
type GstGLFilter struct {
	*GstGLBaseFilter
}

// gstreamer/subprojects/gst-plugins-base/gst-libs/gst/gl/gstglmixer.c
type GstGLMixer struct {
	*GstGLBaseMixer
}
