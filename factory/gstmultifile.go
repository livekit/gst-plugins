// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Convenience bin that muxes incoming streams into multiple time/size limited files
type GstSplitMuxSink struct {
	*gst.Bin
}

func NewSplitMuxSink() (*GstSplitMuxSink, error) {
	e, err := gst.NewElement("splitmuxsink")
	if err != nil {
		return nil, err
	}
	return &GstSplitMuxSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewSplitMuxSinkWithName(name string) (*GstSplitMuxSink, error) {
	e, err := gst.NewElementWithName("splitmuxsink", name)
	if err != nil {
		return nil, err
	}
	return &GstSplitMuxSink{Bin: &gst.Bin{Element: e}}, nil
}

// Maximum difference in timecode between first and last frame. Separator is assumed to be ":" everywhere (e.g. 01:00:00:00). Will only be effective if a timecode track is present.
// Default: NULL
func (e *GstSplitMuxSink) SetMaxSizeTimecode(value string) error {
	return e.Element.SetProperty("max-size-timecode", value)
}

func (e *GstSplitMuxSink) GetMaxSizeTimecode() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("max-size-timecode")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Request a keyframe every max-size-time ns to try splitting at that point. Needs max-size-bytes to be 0 in order to be effective.
// Default: false
func (e *GstSplitMuxSink) SetSendKeyframeRequests(value bool) error {
	return e.Element.SetProperty("send-keyframe-requests", value)
}

func (e *GstSplitMuxSink) GetSendKeyframeRequests() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-keyframe-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The sink element (or element chain) to use (NULL = default filesink). Valid only for async-finalize = FALSE

func (e *GstSplitMuxSink) SetSink(value *gst.Element) error {
	return e.Element.SetProperty("sink", value)
}

func (e *GstSplitMuxSink) GetSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// The sink preset to use. Valid only for async-finalize = TRUE
// Default: NULL
func (e *GstSplitMuxSink) SetSinkPreset(value string) error {
	return e.Element.SetProperty("sink-preset", value)
}

func (e *GstSplitMuxSink) GetSinkPreset() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("sink-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Finalize each fragment asynchronously and start a new one
// Default: false
func (e *GstSplitMuxSink) SetAsyncFinalize(value bool) error {
	return e.Element.SetProperty("async-finalize", value)
}

func (e *GstSplitMuxSink) GetAsyncFinalize() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async-finalize")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Format string pattern for the location of the files to write (e.g. video%%05d.mp4)
// Default: NULL
func (e *GstSplitMuxSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstSplitMuxSink) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Maximum number of files to keep on disk. Once the maximum is reached,old files start to be deleted to make room for new ones.
// Default: 0, Min: 0, Max: -1
func (e *GstSplitMuxSink) SetMaxFiles(value uint) error {
	return e.Element.SetProperty("max-files", value)
}

func (e *GstSplitMuxSink) GetMaxFiles() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-files")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extra size overhead of muxing (0.02 = 2%%)
// Default: 0.02, Min: 0, Max: 1
func (e *GstSplitMuxSink) SetMuxOverhead(value float64) error {
	return e.Element.SetProperty("mux-overhead", value)
}

func (e *GstSplitMuxSink) GetMuxOverhead() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("mux-overhead")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The muxer element to use (NULL = default mp4mux). Valid only for async-finalize = FALSE

func (e *GstSplitMuxSink) SetMuxer(value *gst.Element) error {
	return e.Element.SetProperty("muxer", value)
}

func (e *GstSplitMuxSink) GetMuxer() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("muxer")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// A GstStructure specifies the mapping from splitmuxsink sink pads to muxer pads

func (e *GstSplitMuxSink) SetMuxerPadMap(value *gst.Structure) error {
	return e.Element.SetProperty("muxer-pad-map", value)
}

func (e *GstSplitMuxSink) GetMuxerPadMap() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("muxer-pad-map")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// The sink element factory to use (default = filesink). Valid only for async-finalize = TRUE
// Default: filesink
func (e *GstSplitMuxSink) SetSinkFactory(value string) error {
	return e.Element.SetProperty("sink-factory", value)
}

func (e *GstSplitMuxSink) GetSinkFactory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("sink-factory")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Max. amount of data per file (in bytes, 0=disable)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstSplitMuxSink) SetMaxSizeBytes(value uint64) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

func (e *GstSplitMuxSink) GetMaxSizeBytes() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Max. amount of time per file (in ns, 0=disable)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstSplitMuxSink) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstSplitMuxSink) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The muxer element factory to use (default = mp4mux). Valid only for async-finalize = TRUE
// Default: mp4mux
func (e *GstSplitMuxSink) SetMuxerFactory(value string) error {
	return e.Element.SetProperty("muxer-factory", value)
}

func (e *GstSplitMuxSink) GetMuxerFactory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("muxer-factory")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The muxer element properties to use. Example: {properties,boolean-prop=true,string-prop="hi"}. Valid only for async-finalize = TRUE

func (e *GstSplitMuxSink) SetMuxerProperties(value *gst.Structure) error {
	return e.Element.SetProperty("muxer-properties", value)
}

func (e *GstSplitMuxSink) GetMuxerProperties() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("muxer-properties")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Reset the muxer after each segment. Disabling this will not work for most muxers.
// Default: true
func (e *GstSplitMuxSink) SetResetMuxer(value bool) error {
	return e.Element.SetProperty("reset-muxer", value)
}

func (e *GstSplitMuxSink) GetResetMuxer() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reset-muxer")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Start value of fragment index.
// Default: 0, Min: 0, Max: 2147483647
func (e *GstSplitMuxSink) SetStartIndex(value int) error {
	return e.Element.SetProperty("start-index", value)
}

func (e *GstSplitMuxSink) GetStartIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("start-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Check if muxers support robust muxing via the reserved-max-duration and reserved-duration-remaining properties and use them if so. (Only present on qtmux and mp4mux for now). splitmuxsink may then also  create new fragments if the reserved header space is about to overflow. Note that for mp4mux and qtmux, reserved-moov-update-period must be set manually by the app to a non-zero value for robust muxing to have an effect.
// Default: false
func (e *GstSplitMuxSink) SetUseRobustMuxing(value bool) error {
	return e.Element.SetProperty("use-robust-muxing", value)
}

func (e *GstSplitMuxSink) GetUseRobustMuxing() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-robust-muxing")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Allow non-reference streams to be that many ns before the reference stream
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstSplitMuxSink) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

func (e *GstSplitMuxSink) GetAlignmentThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("alignment-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The muxer preset to use. Valid only for async-finalize = TRUE
// Default: NULL
func (e *GstSplitMuxSink) SetMuxerPreset(value string) error {
	return e.Element.SetProperty("muxer-preset", value)
}

func (e *GstSplitMuxSink) GetMuxerPreset() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("muxer-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The sink element properties to use. Example: {properties,boolean-prop=true,string-prop="hi"}. Valid only for async-finalize = TRUE

func (e *GstSplitMuxSink) SetSinkProperties(value *gst.Structure) error {
	return e.Element.SetProperty("sink-properties", value)
}

func (e *GstSplitMuxSink) GetSinkProperties() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("sink-properties")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Source that reads a set of files created by splitmuxsink
type GstSplitMuxSrc struct {
	*gst.Bin
}

func NewSplitMuxSrc() (*GstSplitMuxSrc, error) {
	e, err := gst.NewElement("splitmuxsrc")
	if err != nil {
		return nil, err
	}
	return &GstSplitMuxSrc{Bin: &gst.Bin{Element: e}}, nil
}

func NewSplitMuxSrcWithName(name string) (*GstSplitMuxSrc, error) {
	e, err := gst.NewElementWithName("splitmuxsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstSplitMuxSrc{Bin: &gst.Bin{Element: e}}, nil
}

// Glob pattern for the location of the files to read
// Default: NULL
func (e *GstSplitMuxSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstSplitMuxSrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Create a video stream from a sequence of image files
type GstImageSequenceSrc struct {
	*base.GstPushSrc
}

func NewImageSequenceSrc() (*GstImageSequenceSrc, error) {
	e, err := gst.NewElement("imagesequencesrc")
	if err != nil {
		return nil, err
	}
	return &GstImageSequenceSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewImageSequenceSrcWithName(name string) (*GstImageSequenceSrc, error) {
	e, err := gst.NewElementWithName("imagesequencesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstImageSequenceSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Pattern to create file names of input files.  File names are created by calling sprintf() with the pattern and the current index.
// Default: %%05d
func (e *GstImageSequenceSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstImageSequenceSrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Start value of index.  The initial value of index can be set either by setting index or start-index.  When the end of the loop is reached, the index will be set to the value start-index.
// Default: 0, Min: 0, Max: 2147483647
func (e *GstImageSequenceSrc) SetStartIndex(value int) error {
	return e.Element.SetProperty("start-index", value)
}

func (e *GstImageSequenceSrc) GetStartIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("start-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Stop value of index.  The special value -1 means no stop.
// Default: -1, Min: -1, Max: 2147483647
func (e *GstImageSequenceSrc) SetStopIndex(value int) error {
	return e.Element.SetProperty("stop-index", value)
}

func (e *GstImageSequenceSrc) GetStopIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("stop-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The output framerate.
// Default: 30/1, Min: 1/1, Max: 2147483647/1
func (e *GstImageSequenceSrc) SetFramerate(value interface{}) error {
	return e.Element.SetProperty("framerate", value)
}

// Write buffers to a sequentially named set of files
type GstMultiFileSink struct {
	*base.GstBaseSink
}

func NewMultiFileSink() (*GstMultiFileSink, error) {
	e, err := gst.NewElement("multifilesink")
	if err != nil {
		return nil, err
	}
	return &GstMultiFileSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewMultiFileSinkWithName(name string) (*GstMultiFileSink, error) {
	e, err := gst.NewElementWithName("multifilesink", name)
	if err != nil {
		return nil, err
	}
	return &GstMultiFileSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Location of the file to write
// Default: %%05d
func (e *GstMultiFileSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstMultiFileSink) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Minimum distance between keyframes to start a new file
// Default: 10000000000, Min: 0, Max: 18446744073709551615
func (e *GstMultiFileSink) SetMinKeyframeDistance(value uint64) error {
	return e.Element.SetProperty("min-keyframe-distance", value)
}

func (e *GstMultiFileSink) GetMinKeyframeDistance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-keyframe-distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Post a message for each file with information of the buffer
// Default: false
func (e *GstMultiFileSink) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

func (e *GstMultiFileSink) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to aggregate GOPs and process them as a whole without splitting
// Default: false
func (e *GstMultiFileSink) SetAggregateGops(value bool) error {
	return e.Element.SetProperty("aggregate-gops", value)
}

func (e *GstMultiFileSink) GetAggregateGops() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("aggregate-gops")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Index to use with location property to create file names.  The index is incremented by one for each buffer written.
// Default: 0, Min: 0, Max: 2147483647
func (e *GstMultiFileSink) SetIndex(value int) error {
	return e.Element.SetProperty("index", value)
}

func (e *GstMultiFileSink) GetIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum file duration before starting a new file in max-duration mode (in nanoseconds)
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstMultiFileSink) SetMaxFileDuration(value uint64) error {
	return e.Element.SetProperty("max-file-duration", value)
}

func (e *GstMultiFileSink) GetMaxFileDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-file-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum file size before starting a new file in max-size mode
// Default: 2147483648, Min: 0, Max: 18446744073709551615
func (e *GstMultiFileSink) SetMaxFileSize(value uint64) error {
	return e.Element.SetProperty("max-file-size", value)
}

func (e *GstMultiFileSink) GetMaxFileSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-file-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum number of files to keep on disk. Once the maximum is reached,old files start to be deleted to make room for new ones.
// Default: 0, Min: 0, Max: -1
func (e *GstMultiFileSink) SetMaxFiles(value uint) error {
	return e.Element.SetProperty("max-files", value)
}

func (e *GstMultiFileSink) GetMaxFiles() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-files")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// When to start a new file
// Default: buffer (0)
func (e *GstMultiFileSink) SetNextFile(value GstMultiFileSinkNext) error {
	e.Element.SetArg("next-file", string(value))
	return nil
}

func (e *GstMultiFileSink) GetNextFile() (GstMultiFileSinkNext, error) {
	var value GstMultiFileSinkNext
	var ok bool
	v, err := e.Element.GetProperty("next-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMultiFileSinkNext)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMultiFileSinkNext")
	}
	return value, nil
}

// Read a sequentially named set of files into buffers
type GstMultiFileSrc struct {
	*base.GstPushSrc
}

func NewMultiFileSrc() (*GstMultiFileSrc, error) {
	e, err := gst.NewElement("multifilesrc")
	if err != nil {
		return nil, err
	}
	return &GstMultiFileSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewMultiFileSrcWithName(name string) (*GstMultiFileSrc, error) {
	e, err := gst.NewElementWithName("multifilesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstMultiFileSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Pattern to create file names of input files.  File names are created by calling sprintf() with the pattern and the current index.
// Default: %%05d
func (e *GstMultiFileSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstMultiFileSrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Whether to repeat from the beginning when all files have been read.
// Default: false
func (e *GstMultiFileSrc) SetLoop(value bool) error {
	return e.Element.SetProperty("loop", value)
}

func (e *GstMultiFileSrc) GetLoop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("loop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Start value of index.  The initial value of index can be set either by setting index or start-index.  When the end of the loop is reached, the index will be set to the value start-index.
// Default: 0, Min: 0, Max: 2147483647
func (e *GstMultiFileSrc) SetStartIndex(value int) error {
	return e.Element.SetProperty("start-index", value)
}

func (e *GstMultiFileSrc) GetStartIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("start-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Stop value of index.  The special value -1 means no stop.
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMultiFileSrc) SetStopIndex(value int) error {
	return e.Element.SetProperty("stop-index", value)
}

func (e *GstMultiFileSrc) GetStopIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("stop-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Caps describing the format of the data.

func (e *GstMultiFileSrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstMultiFileSrc) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Index to use with location property to create file names.  The index is incremented by one for each buffer read.
// Default: 0, Min: 0, Max: 2147483647
func (e *GstMultiFileSrc) SetIndex(value int) error {
	return e.Element.SetProperty("index", value)
}

func (e *GstMultiFileSrc) GetIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Read a sequentially named set of files as if it was one large file
type GstSplitFileSrc struct {
	*base.GstBaseSrc
}

func NewSplitFileSrc() (*GstSplitFileSrc, error) {
	e, err := gst.NewElement("splitfilesrc")
	if err != nil {
		return nil, err
	}
	return &GstSplitFileSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewSplitFileSrcWithName(name string) (*GstSplitFileSrc, error) {
	e, err := gst.NewElementWithName("splitfilesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstSplitFileSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Wildcard pattern to match file names of the input files. If the location is an absolute path or contains directory components, only the base file name part will be considered for pattern matching. The results will be sorted.
// Default: NULL
func (e *GstSplitFileSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstSplitFileSrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstMultiFileSinkNext string

const (
	GstMultiFileSinkNextBuffer       GstMultiFileSinkNext = "buffer"         // New file for each buffer
	GstMultiFileSinkNextDiscont      GstMultiFileSinkNext = "discont"        // New file after each discontinuity
	GstMultiFileSinkNextKeyFrame     GstMultiFileSinkNext = "key-frame"      // New file at each key frame (Useful for MPEG-TS segmenting)
	GstMultiFileSinkNextKeyUnitEvent GstMultiFileSinkNext = "key-unit-event" // New file after a force key unit event
	GstMultiFileSinkNextMaxSize      GstMultiFileSinkNext = "max-size"       // New file when the configured maximum file size would be exceeded with the next buffer or buffer list
	GstMultiFileSinkNextMaxDuration  GstMultiFileSinkNext = "max-duration"   // New file when the configured maximum file duration would be exceeded with the next buffer or buffer list
)
