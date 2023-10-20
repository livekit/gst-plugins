// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Upload data over FILE protocol using libcurl
type GstCurlFileSink struct {
	*GstCurlBaseSink
}

func NewCurlFileSink() (*GstCurlFileSink, error) {
	e, err := gst.NewElement("curlfilesink")
	if err != nil {
		return nil, err
	}
	return &GstCurlFileSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewCurlFileSinkWithName(name string) (*GstCurlFileSink, error) {
	e, err := gst.NewElementWithName("curlfilesink", name)
	if err != nil {
		return nil, err
	}
	return &GstCurlFileSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Attempt to create missing directory included in the path
// Default: false
func (e *GstCurlFileSink) SetCreateDirs(value bool) error {
	return e.Element.SetProperty("create-dirs", value)
}

func (e *GstCurlFileSink) GetCreateDirs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("create-dirs")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Upload data over FTP protocol using libcurl
type GstCurlFtpSink struct {
	*GstCurlTlsSink
}

func NewCurlFtpSink() (*GstCurlFtpSink, error) {
	e, err := gst.NewElement("curlftpsink")
	if err != nil {
		return nil, err
	}
	return &GstCurlFtpSink{GstCurlTlsSink: &GstCurlTlsSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewCurlFtpSinkWithName(name string) (*GstCurlFtpSink, error) {
	e, err := gst.NewElementWithName("curlftpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstCurlFtpSink{GstCurlTlsSink: &GstCurlTlsSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// Attempt to create missing directory included in the path
// Default: false
func (e *GstCurlFtpSink) SetCreateDirs(value bool) error {
	return e.Element.SetProperty("create-dirs", value)
}

func (e *GstCurlFtpSink) GetCreateDirs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("create-dirs")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use a temporary file name when uploading a a file. When the transfer is complete,            this temporary file is renamed to the final file name. This is useful for ensuring           that remote systems do not read a partially uploaded file
// Default: false
func (e *GstCurlFtpSink) SetCreateTmpFile(value bool) error {
	return e.Element.SetProperty("create-tmp-file", value)
}

func (e *GstCurlFtpSink) GetCreateTmpFile() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("create-tmp-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable the use of the EPSV command when doing passive FTP transfers
// Default: false
func (e *GstCurlFtpSink) SetEpsvMode(value bool) error {
	return e.Element.SetProperty("epsv-mode", value)
}

func (e *GstCurlFtpSink) GetEpsvMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("epsv-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The PORT instruction tells the remote server to connect to the IP address
// Default: NULL
func (e *GstCurlFtpSink) SetFtpPort(value string) error {
	return e.Element.SetProperty("ftp-port", value)
}

func (e *GstCurlFtpSink) GetFtpPort() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ftp-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Filename pattern to use when generating a temporary filename for uploads
// Default: NULL
func (e *GstCurlFtpSink) SetTempFileName(value string) error {
	return e.Element.SetProperty("temp-file-name", value)
}

func (e *GstCurlFtpSink) GetTempFileName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-file-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Upload data over HTTP/HTTPS protocol using libcurl
type GstCurlHttpSink struct {
	*GstCurlTlsSink
}

func NewCurlHttpSink() (*GstCurlHttpSink, error) {
	e, err := gst.NewElement("curlhttpsink")
	if err != nil {
		return nil, err
	}
	return &GstCurlHttpSink{GstCurlTlsSink: &GstCurlTlsSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewCurlHttpSinkWithName(name string) (*GstCurlHttpSink, error) {
	e, err := gst.NewElementWithName("curlhttpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstCurlHttpSink{GstCurlTlsSink: &GstCurlTlsSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// HTTP proxy server port
// Default: 3128, Min: 0, Max: 2147483647
func (e *GstCurlHttpSink) SetProxyPort(value int) error {
	return e.Element.SetProperty("proxy-port", value)
}

func (e *GstCurlHttpSink) GetProxyPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("proxy-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Proxy user name to use for proxy authentication
// Default: NULL
func (e *GstCurlHttpSink) SetProxyUser(value string) error {
	return e.Element.SetProperty("proxy-user", value)
}

func (e *GstCurlHttpSink) GetProxyUser() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-user")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Use the Content-Length HTTP header instead of Transfer-Encoding header
// Default: false
func (e *GstCurlHttpSink) SetUseContentLength(value bool) error {
	return e.Element.SetProperty("use-content-length", value)
}

func (e *GstCurlHttpSink) GetUseContentLength() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-content-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Content Type to use for the Content-Type header. If not set, detected mime type will be used
// Default: NULL
func (e *GstCurlHttpSink) SetContentType(value string) error {
	return e.Element.SetProperty("content-type", value)
}

func (e *GstCurlHttpSink) GetContentType() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("content-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// HTTP proxy server URI
// Default: NULL
func (e *GstCurlHttpSink) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

func (e *GstCurlHttpSink) GetProxy() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Proxy user password to use for proxy authentication
// Default: NULL
func (e *GstCurlHttpSink) SetProxyPasswd(value string) error {
	return e.Element.SetProperty("proxy-passwd", value)
}

func (e *GstCurlHttpSink) GetProxyPasswd() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-passwd")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Receiver data as a client over a network via HTTP using cURL
type GstCurlHttpSrc struct {
	*base.GstPushSrc
}

func NewCurlHttpSrc() (*GstCurlHttpSrc, error) {
	e, err := gst.NewElement("curlhttpsrc")
	if err != nil {
		return nil, err
	}
	return &GstCurlHttpSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewCurlHttpSrcWithName(name string) (*GstCurlHttpSrc, error) {
	e, err := gst.NewElementWithName("curlhttpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstCurlHttpSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// The preferred HTTP protocol version
// Default: 2.0 (2)
func (e *GstCurlHttpSrc) SetHttpVersion(value GstCurlHttpVersionType) error {
	e.Element.SetArg("http-version", string(value))
	return nil
}

func (e *GstCurlHttpSrc) GetHttpVersion() (GstCurlHttpVersionType, error) {
	var value GstCurlHttpVersionType
	var ok bool
	v, err := e.Element.GetProperty("http-version")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCurlHttpVersionType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCurlHttpVersionType")
	}
	return value, nil
}

// Maximum number of concurrent connections allowed for HTTP/1.x
// Default: 255, Min: 1, Max: 255
func (e *GstCurlHttpSrc) SetMaxConnections(value uint) error {
	return e.Element.SetProperty("max-connections", value)
}

func (e *GstCurlHttpSrc) GetMaxConnections() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-connections")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum number of concurrent connections allowed per proxy for HTTP/1.x
// Default: 30, Min: 1, Max: 60
func (e *GstCurlHttpSrc) SetMaxConnectionsPerProxy(value uint) error {
	return e.Element.SetProperty("max-connections-per-proxy", value)
}

func (e *GstCurlHttpSrc) GetMaxConnectionsPerProxy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-connections-per-proxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// HTTP proxy URI password for authentication
// Default: NULL
func (e *GstCurlHttpSrc) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

func (e *GstCurlHttpSrc) GetProxyPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// URI of resource requested
// Default: GStreamer curlhttpsrc libcurl/<curl-version>
func (e *GstCurlHttpSrc) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

func (e *GstCurlHttpSrc) GetUserAgent() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-agent")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// HTTP location URI password for authentication
// Default: NULL
func (e *GstCurlHttpSrc) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

func (e *GstCurlHttpSrc) GetUserPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Allow compressed content encodings
// Default: false
func (e *GstCurlHttpSrc) SetCompress(value bool) error {
	return e.Element.SetProperty("compress", value)
}

func (e *GstCurlHttpSrc) GetCompress() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("compress")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extra headers to append to the HTTP request
// Default: request-headers;
func (e *GstCurlHttpSrc) SetExtraHeaders(value *gst.Structure) error {
	return e.Element.SetProperty("extra-headers", value)
}

func (e *GstCurlHttpSrc) GetExtraHeaders() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("extra-headers")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// URI of HTTP proxy server
// Default: NULL
func (e *GstCurlHttpSrc) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

func (e *GstCurlHttpSrc) GetProxy() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Location of an SSL CA file to use for checking SSL certificates
// Default: NULL
func (e *GstCurlHttpSrc) SetSslCaFile(value string) error {
	return e.Element.SetProperty("ssl-ca-file", value)
}

func (e *GstCurlHttpSrc) GetSslCaFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssl-ca-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// List of HTTP Cookies

func (e *GstCurlHttpSrc) SetCookies(value interface{}) error {
	return e.Element.SetProperty("cookies", value)
}

func (e *GstCurlHttpSrc) GetCookies() (interface{}, error) {
	return e.Element.GetProperty("cookies")
}

// Maximum amount of time to keep-alive HTTP connections
// Default: 30, Min: 2, Max: 60
func (e *GstCurlHttpSrc) SetMaxConnectionTime(value uint) error {
	return e.Element.SetProperty("max-connection-time", value)
}

func (e *GstCurlHttpSrc) GetMaxConnectionTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-connection-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum number of connections allowed per server for HTTP/1.x
// Default: 5, Min: 1, Max: 60
func (e *GstCurlHttpSrc) SetMaxConnectionsPerServer(value uint) error {
	return e.Element.SetProperty("max-connections-per-server", value)
}

func (e *GstCurlHttpSrc) GetMaxConnectionsPerServer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-connections-per-server")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Strict SSL certificate checking
// Default: true
func (e *GstCurlHttpSrc) SetSslStrict(value bool) error {
	return e.Element.SetProperty("ssl-strict", value)
}

func (e *GstCurlHttpSrc) GetSslStrict() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssl-strict")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Value in seconds before timeout a blocking request (0 = no timeout)
// Default: 0, Min: 0, Max: 3600
func (e *GstCurlHttpSrc) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstCurlHttpSrc) GetTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// HTTP location URI user id for authentication
// Default: NULL
func (e *GstCurlHttpSrc) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

func (e *GstCurlHttpSrc) GetUserId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Toggle keep-alive for connection reuse.
// Default: true
func (e *GstCurlHttpSrc) SetKeepAlive(value bool) error {
	return e.Element.SetProperty("keep-alive", value)
}

func (e *GstCurlHttpSrc) GetKeepAlive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keep-alive")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// URI of resource to read
// Default: NULL
func (e *GstCurlHttpSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstCurlHttpSrc) GetLocation() (string, error) {
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

// HTTP proxy URI user id for authentication
// Default: NULL
func (e *GstCurlHttpSrc) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

func (e *GstCurlHttpSrc) GetProxyId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Maximum number of retries until giving up (-1=infinite)
// Default: -1, Min: -1, Max: 9999
func (e *GstCurlHttpSrc) SetRetries(value int) error {
	return e.Element.SetProperty("retries", value)
}

func (e *GstCurlHttpSrc) GetRetries() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("retries")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Allow HTTP Redirections (HTTP Status Code 300 series)
// Default: true
func (e *GstCurlHttpSrc) SetAutomaticRedirect(value bool) error {
	return e.Element.SetProperty("automatic-redirect", value)
}

func (e *GstCurlHttpSrc) GetAutomaticRedirect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("automatic-redirect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum number of permitted redirections. -1 is unlimited.
// Default: -1, Min: -1, Max: 255
func (e *GstCurlHttpSrc) SetMaxRedirect(value int) error {
	return e.Element.SetProperty("max-redirect", value)
}

func (e *GstCurlHttpSrc) GetMaxRedirect() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-redirect")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Upload data over the SFTP protocol using libcurl
type GstCurlSftpSink struct {
	*GstCurlSshSink
}

func NewCurlSftpSink() (*GstCurlSftpSink, error) {
	e, err := gst.NewElement("curlsftpsink")
	if err != nil {
		return nil, err
	}
	return &GstCurlSftpSink{GstCurlSshSink: &GstCurlSshSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewCurlSftpSinkWithName(name string) (*GstCurlSftpSink, error) {
	e, err := gst.NewElementWithName("curlsftpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstCurlSftpSink{GstCurlSshSink: &GstCurlSshSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// Timestamp offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstCurlSftpSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstCurlSftpSink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Sink Statistics
// Default: application/x-gst-base-sink-stats, average-rate=(double)0, dropped=(guint64)0, rendered=(guint64)0;
func (e *GstCurlSftpSink) GetStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstCurlSftpSink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

func (e *GstCurlSftpSink) GetMaxLateness() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-lateness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// User password to use for server authentication
// Default: NULL
func (e *GstCurlSftpSink) SetPasswd(value string) error {
	return e.Element.SetProperty("passwd", value)
}

func (e *GstCurlSftpSink) GetPasswd() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("passwd")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Sync on the clock
// Default: true
func (e *GstCurlSftpSink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstCurlSftpSink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// MD5 checksum (32 hexadecimal digits, case-insensitive) of the remote host's public key
// Default: NULL
func (e *GstCurlSftpSink) SetSshHostPubkeyMd5(value string) error {
	return e.Element.SetProperty("ssh-host-pubkey-md5", value)
}

func (e *GstCurlSftpSink) GetSshHostPubkeyMd5() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-host-pubkey-md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The passphrase used to protect the SSH private key file
// Default: NULL
func (e *GstCurlSftpSink) SetSshKeyPassphrase(value string) error {
	return e.Element.SetProperty("ssh-key-passphrase", value)
}

func (e *GstCurlSftpSink) GetSshKeyPassphrase() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-key-passphrase")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Go asynchronously to PAUSED
// Default: true
func (e *GstCurlSftpSink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

func (e *GstCurlSftpSink) GetAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Attempt to create missing directories
// Default: false
func (e *GstCurlSftpSink) SetCreateDirs(value bool) error {
	return e.Element.SetProperty("create-dirs", value)
}

func (e *GstCurlSftpSink) GetCreateDirs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("create-dirs")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The complete path & filename of the SSH private key file
// Default: NULL
func (e *GstCurlSftpSink) SetSshPrivKeyfile(value string) error {
	return e.Element.SetProperty("ssh-priv-keyfile", value)
}

func (e *GstCurlSftpSink) GetSshPrivKeyfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-priv-keyfile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The last sample received in the sink

func (e *GstCurlSftpSink) GetLastSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// Additional render delay of the sink in nanoseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstCurlSftpSink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

func (e *GstCurlSftpSink) GetRenderDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("render-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// SSH authentication method to authenticate on the SSH/SFTP server
// Default: none (0)
func (e *GstCurlSftpSink) SetSshAuthType(value GstCurlSshAuthType) error {
	e.Element.SetArg("ssh-auth-type", string(value))
	return nil
}

func (e *GstCurlSftpSink) GetSshAuthType() (GstCurlSshAuthType, error) {
	var value GstCurlSshAuthType
	var ok bool
	v, err := e.Element.GetProperty("ssh-auth-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCurlSshAuthType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCurlSshAuthType")
	}
	return value, nil
}

// The complete path & filename of the SSH 'known_hosts' file
// Default: NULL
func (e *GstCurlSftpSink) SetSshKnownhosts(value string) error {
	return e.Element.SetProperty("ssh-knownhosts", value)
}

func (e *GstCurlSftpSink) GetSshKnownhosts() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-knownhosts")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Generate Quality-of-Service events upstream
// Default: false
func (e *GstCurlSftpSink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstCurlSftpSink) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The time to keep between rendered buffers (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstCurlSftpSink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

func (e *GstCurlSftpSink) GetThrottleTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("throttle-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Number of seconds waiting to write before timeout
// Default: 30, Min: 0, Max: 2147483647
func (e *GstCurlSftpSink) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstCurlSftpSink) GetTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Size in bytes to pull per buffer (0 = default)
// Default: 4096, Min: 0, Max: -1
func (e *GstCurlSftpSink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstCurlSftpSink) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum processing time for a buffer in nanoseconds
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstCurlSftpSink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstCurlSftpSink) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Accept an unknown remote public host key
// Default: false
func (e *GstCurlSftpSink) SetSshAcceptUnknownhost(value bool) error {
	return e.Element.SetProperty("ssh-accept-unknownhost", value)
}

func (e *GstCurlSftpSink) GetSshAcceptUnknownhost() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssh-accept-unknownhost")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// User name to use for server authentication
// Default: NULL
func (e *GstCurlSftpSink) SetUser(value string) error {
	return e.Element.SetProperty("user", value)
}

func (e *GstCurlSftpSink) GetUser() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// URI location to write to
// Default: localhost:5555
func (e *GstCurlSftpSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstCurlSftpSink) GetLocation() (string, error) {
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

// The maximum bits per second to render (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstCurlSftpSink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstCurlSftpSink) GetMaxBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Enable the last-sample property
// Default: true
func (e *GstCurlSftpSink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

func (e *GstCurlSftpSink) GetEnableLastSample() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The base file name for the uploaded images
// Default: NULL
func (e *GstCurlSftpSink) SetFileName(value string) error {
	return e.Element.SetProperty("file-name", value)
}

func (e *GstCurlSftpSink) GetFileName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("file-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The complete path & filename of the SSH public key file
// Default: NULL
func (e *GstCurlSftpSink) SetSshPubKeyfile(value string) error {
	return e.Element.SetProperty("ssh-pub-keyfile", value)
}

func (e *GstCurlSftpSink) GetSshPubKeyfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-pub-keyfile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Quality of Service, differentiated services code point (0 default)
// Default: 0, Min: 0, Max: 63
func (e *GstCurlSftpSink) SetQosDscp(value int) error {
	return e.Element.SetProperty("qos-dscp", value)
}

func (e *GstCurlSftpSink) GetQosDscp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qos-dscp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Upload data over SMTP protocol using libcurl
type GstCurlSmtpSink struct {
	*GstCurlTlsSink
}

func NewCurlSmtpSink() (*GstCurlSmtpSink, error) {
	e, err := gst.NewElement("curlsmtpsink")
	if err != nil {
		return nil, err
	}
	return &GstCurlSmtpSink{GstCurlTlsSink: &GstCurlTlsSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewCurlSmtpSinkWithName(name string) (*GstCurlSmtpSink, error) {
	e, err := gst.NewElementWithName("curlsmtpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstCurlSmtpSink{GstCurlTlsSink: &GstCurlTlsSink{GstCurlBaseSink: &GstCurlBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// Single address that the given mail should get sent to
// Default: NULL
func (e *GstCurlSmtpSink) SetMailRcpt(value string) error {
	return e.Element.SetProperty("mail-rcpt", value)
}

func (e *GstCurlSmtpSink) GetMailRcpt() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mail-rcpt")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Message body
// Default: NULL
func (e *GstCurlSmtpSink) SetMessageBody(value string) error {
	return e.Element.SetProperty("message-body", value)
}

func (e *GstCurlSmtpSink) GetMessageBody() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("message-body")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Number attachments to send
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstCurlSmtpSink) SetNbrAttachments(value int) error {
	return e.Element.SetProperty("nbr-attachments", value)
}

func (e *GstCurlSmtpSink) GetNbrAttachments() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("nbr-attachments")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// URL POP used for authentication
// Default: NULL
func (e *GstCurlSmtpSink) SetPopLocation(value string) error {
	return e.Element.SetProperty("pop-location", value)
}

func (e *GstCurlSmtpSink) GetPopLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pop-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Use SSL/TLS for the connection
// Default: false
func (e *GstCurlSmtpSink) SetUseSsl(value bool) error {
	return e.Element.SetProperty("use-ssl", value)
}

func (e *GstCurlSmtpSink) GetUseSsl() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-ssl")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The mime type of the body of the request
// Default: NULL
func (e *GstCurlSmtpSink) SetContentType(value string) error {
	return e.Element.SetProperty("content-type", value)
}

func (e *GstCurlSmtpSink) GetContentType() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("content-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Single address that the given mail should get sent from
// Default: NULL
func (e *GstCurlSmtpSink) SetMailFrom(value string) error {
	return e.Element.SetProperty("mail-from", value)
}

func (e *GstCurlSmtpSink) GetMailFrom() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mail-from")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// User password to use for POP server authentication
// Default: NULL
func (e *GstCurlSmtpSink) SetPopPasswd(value string) error {
	return e.Element.SetProperty("pop-passwd", value)
}

func (e *GstCurlSmtpSink) GetPopPasswd() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pop-passwd")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// User name to use for POP server authentication
// Default: NULL
func (e *GstCurlSmtpSink) SetPopUser(value string) error {
	return e.Element.SetProperty("pop-user", value)
}

func (e *GstCurlSmtpSink) GetPopUser() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pop-user")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Mail subject
// Default: NULL
func (e *GstCurlSmtpSink) SetSubject(value string) error {
	return e.Element.SetProperty("subject", value)
}

func (e *GstCurlSmtpSink) GetSubject() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subject")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstCurlBaseSink struct {
	*base.GstBaseSink
}

// The base file name for the uploaded images
// Default: NULL
func (e *GstCurlBaseSink) SetFileName(value string) error {
	return e.Element.SetProperty("file-name", value)
}

func (e *GstCurlBaseSink) GetFileName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("file-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// URI location to write to
// Default: localhost:5555
func (e *GstCurlBaseSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstCurlBaseSink) GetLocation() (string, error) {
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

// User password to use for server authentication
// Default: NULL
func (e *GstCurlBaseSink) SetPasswd(value string) error {
	return e.Element.SetProperty("passwd", value)
}

func (e *GstCurlBaseSink) GetPasswd() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("passwd")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Quality of Service, differentiated services code point (0 default)
// Default: 0, Min: 0, Max: 63
func (e *GstCurlBaseSink) SetQosDscp(value int) error {
	return e.Element.SetProperty("qos-dscp", value)
}

func (e *GstCurlBaseSink) GetQosDscp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qos-dscp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of seconds waiting to write before timeout
// Default: 30, Min: 0, Max: 2147483647
func (e *GstCurlBaseSink) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstCurlBaseSink) GetTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// User name to use for server authentication
// Default: NULL
func (e *GstCurlBaseSink) SetUser(value string) error {
	return e.Element.SetProperty("user", value)
}

func (e *GstCurlBaseSink) GetUser() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstCurlHttpVersionType string

const (
	GstCurlHttpVersionType10 GstCurlHttpVersionType = "1.0" // HTTP Version 1.0
	GstCurlHttpVersionType11 GstCurlHttpVersionType = "1.1" // HTTP Version 1.1
	GstCurlHttpVersionType20 GstCurlHttpVersionType = "2.0" // HTTP Version 2.0
)

type GstCurlSshAuthType string

const (
	GstCurlSshAuthTypeNone     GstCurlSshAuthType = "none"     // Not allowed
	GstCurlSshAuthTypePubkey   GstCurlSshAuthType = "pubkey"   // Public/private key files
	GstCurlSshAuthTypePassword GstCurlSshAuthType = "password" // Password authentication
)

type GstCurlSshSink struct {
	*GstCurlBaseSink
}

// Accept an unknown remote public host key
// Default: false
func (e *GstCurlSshSink) SetSshAcceptUnknownhost(value bool) error {
	return e.Element.SetProperty("ssh-accept-unknownhost", value)
}

func (e *GstCurlSshSink) GetSshAcceptUnknownhost() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssh-accept-unknownhost")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// SSH authentication method to authenticate on the SSH/SFTP server
// Default: none (0)
func (e *GstCurlSshSink) SetSshAuthType(value GstCurlSshAuthType) error {
	e.Element.SetArg("ssh-auth-type", string(value))
	return nil
}

func (e *GstCurlSshSink) GetSshAuthType() (GstCurlSshAuthType, error) {
	var value GstCurlSshAuthType
	var ok bool
	v, err := e.Element.GetProperty("ssh-auth-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCurlSshAuthType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCurlSshAuthType")
	}
	return value, nil
}

// MD5 checksum (32 hexadecimal digits, case-insensitive) of the remote host's public key
// Default: NULL
func (e *GstCurlSshSink) SetSshHostPubkeyMd5(value string) error {
	return e.Element.SetProperty("ssh-host-pubkey-md5", value)
}

func (e *GstCurlSshSink) GetSshHostPubkeyMd5() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-host-pubkey-md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The passphrase used to protect the SSH private key file
// Default: NULL
func (e *GstCurlSshSink) SetSshKeyPassphrase(value string) error {
	return e.Element.SetProperty("ssh-key-passphrase", value)
}

func (e *GstCurlSshSink) GetSshKeyPassphrase() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-key-passphrase")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The complete path & filename of the SSH 'known_hosts' file
// Default: NULL
func (e *GstCurlSshSink) SetSshKnownhosts(value string) error {
	return e.Element.SetProperty("ssh-knownhosts", value)
}

func (e *GstCurlSshSink) GetSshKnownhosts() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-knownhosts")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The complete path & filename of the SSH private key file
// Default: NULL
func (e *GstCurlSshSink) SetSshPrivKeyfile(value string) error {
	return e.Element.SetProperty("ssh-priv-keyfile", value)
}

func (e *GstCurlSshSink) GetSshPrivKeyfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-priv-keyfile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The complete path & filename of the SSH public key file
// Default: NULL
func (e *GstCurlSshSink) SetSshPubKeyfile(value string) error {
	return e.Element.SetProperty("ssh-pub-keyfile", value)
}

func (e *GstCurlSshSink) GetSshPubKeyfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-pub-keyfile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstCurlTlsSink struct {
	*GstCurlBaseSink
}

// CA certificate to use in order to verify the peer
// Default: NULL
func (e *GstCurlTlsSink) SetCaCert(value string) error {
	return e.Element.SetProperty("ca-cert", value)
}

func (e *GstCurlTlsSink) GetCaCert() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ca-cert")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// CA directory path to use in order to verify the peer
// Default: NULL
func (e *GstCurlTlsSink) SetCaPath(value string) error {
	return e.Element.SetProperty("ca-path", value)
}

func (e *GstCurlTlsSink) GetCaPath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ca-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// OpenSSL crypto engine to use for cipher operations
// Default: NULL
func (e *GstCurlTlsSink) SetCryptoEngine(value string) error {
	return e.Element.SetProperty("crypto-engine", value)
}

func (e *GstCurlTlsSink) GetCryptoEngine() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("crypto-engine")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Allow curl to perform insecure SSL connections
// Default: true
func (e *GstCurlTlsSink) SetInsecure(value bool) error {
	return e.Element.SetProperty("insecure", value)
}

func (e *GstCurlTlsSink) GetInsecure() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("insecure")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
