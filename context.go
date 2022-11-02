package cc4echo

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"

	"github.com/labstack/echo/v4"
)

type Context struct {
	mu sync.Mutex
	ec echo.Context
}

func New(c echo.Context) echo.Context {
	return &Context{ec: c}
}

// Request returns `*http.Request`.
func (c *Context) Request() *http.Request {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Request()
}

// SetRequest sets `*http.Request`.
func (c *Context) SetRequest(r *http.Request) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.SetRequest(r)
}

// SetResponse sets `*Response`.
func (c *Context) SetResponse(r *echo.Response) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.SetResponse(r)
}

// Response returns `*Response`.
func (c *Context) Response() *echo.Response {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Response()
}

// IsTLS returns true if HTTP connection is TLS otherwise false.
func (c *Context) IsTLS() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.IsTLS()
}

// IsWebSocket returns true if HTTP connection is WebSocket otherwise false.
func (c *Context) IsWebSocket() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.IsWebSocket()
}

// Scheme returns the HTTP protocol scheme, `http` or `https`.
func (c *Context) Scheme() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Scheme()
}

// RealIP returns the client's network address based on `X-Forwarded-For`
// or `X-Real-IP` request header.
// The behavior can be configured using `Echo#IPExtractor`.
func (c *Context) RealIP() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.RealIP()
}

// Path returns the registered path for the handler.
func (c *Context) Path() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Path()
}

// SetPath sets the registered path for the handler.
func (c *Context) SetPath(p string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.SetPath(p)
}

// Param returns path parameter by name.
func (c *Context) Param(name string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Param(name)
}

// ParamNames returns path parameter names.
func (c *Context) ParamNames() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.ParamNames()
}

// SetParamNames sets path parameter names.
func (c *Context) SetParamNames(names ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.SetParamNames(names...)
}

// ParamValues returns path parameter values.
func (c *Context) ParamValues() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.ParamValues()
}

// SetParamValues sets path parameter values.
func (c *Context) SetParamValues(values ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.SetParamValues(values...)
}

// QueryParam returns the query param for the provided name.
func (c *Context) QueryParam(name string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.QueryParam(name)
}

// QueryParams returns the query parameters as `url.Values`.
func (c *Context) QueryParams() url.Values {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.QueryParams()
}

// QueryString returns the URL query string.
func (c *Context) QueryString() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.QueryString()
}

// FormValue returns the form field value for the provided name.
func (c *Context) FormValue(name string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.FormValue(name)
}

// FormParams returns the form parameters as `url.Values`.
func (c *Context) FormParams() (url.Values, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.FormParams()
}

// FormFile returns the multipart form file for the provided name.
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.FormFile(name)
}

// MultipartForm returns the multipart form.
func (c *Context) MultipartForm() (*multipart.Form, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.MultipartForm()
}

// Cookie returns the named cookie provided in the request.
func (c *Context) Cookie(name string) (*http.Cookie, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Cookie(name)
}

// SetCookie adds a `Set-Cookie` header in HTTP response.
func (c *Context) SetCookie(cookie *http.Cookie) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.SetCookie(cookie)
}

// Cookies returns the HTTP cookies sent with the request.
func (c *Context) Cookies() []*http.Cookie {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Cookies()
}

// Get retrieves data from the context.
func (c *Context) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Get(key)
}

// Set saves data in the context.
func (c *Context) Set(key string, val interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.Set(key, val)
}

// Bind binds the request body into provided type `i`. The default binder
// does it based on Content-Type header.
func (c *Context) Bind(i interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Bind(i)
}

// Validate validates provided `i`. It is usually called after `Context#Bind()`.
// Validator must be registered using `Echo#Validator`.
func (c *Context) Validate(i interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Validate(i)
}

// Render renders a template with data and sends a text/html response with status
// code. Renderer must be registered using `Echo.Renderer`.
func (c *Context) Render(code int, name string, data interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Render(code, name, data)
}

// HTML sends an HTTP response with status code.
func (c *Context) HTML(code int, html string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.HTML(code, html)
}

// HTMLBlob sends an HTTP blob response with status code.
func (c *Context) HTMLBlob(code int, b []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.HTMLBlob(code, b)
}

// String sends a string response with status code.
func (c *Context) String(code int, s string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.String(code, s)
}

// JSON sends a JSON response with status code.
func (c *Context) JSON(code int, i interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.JSON(code, i)
}

// JSONPretty sends a pretty-print JSON with status code.
func (c *Context) JSONPretty(code int, i interface{}, indent string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.JSONPretty(code, i, indent)
}

// JSONBlob sends a JSON blob response with status code.
func (c *Context) JSONBlob(code int, b []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.JSONBlob(code, b)
}

// JSONP sends a JSONP response with status code. It uses `callback` to construct
// the JSONP payload.
func (c *Context) JSONP(code int, callback string, i interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.JSONP(code, callback, i)
}

// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
// to construct the JSONP payload.
func (c *Context) JSONPBlob(code int, callback string, b []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.JSONPBlob(code, callback, b)
}

// XML sends an XML response with status code.
func (c *Context) XML(code int, i interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.XML(code, i)
}

// XMLPretty sends a pretty-print XML with status code.
func (c *Context) XMLPretty(code int, i interface{}, indent string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.XMLPretty(code, i, indent)
}

// XMLBlob sends an XML blob response with status code.
func (c *Context) XMLBlob(code int, b []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.XMLBlob(code, b)
}

// Blob sends a blob response with status code and content type.
func (c *Context) Blob(code int, contentType string, b []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Blob(code, contentType, b)
}

// Stream sends a streaming response with status code and content type.
func (c *Context) Stream(code int, contentType string, r io.Reader) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Stream(code, contentType, r)
}

// File sends a response with the content of the file.
func (c *Context) File(file string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.File(file)
}

// Attachment sends a response as attachment, prompting client to save the
// file.
func (c *Context) Attachment(file string, name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Attachment(file, name)
}

// Inline sends a response as inline, opening the file in the browser.
func (c *Context) Inline(file string, name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Inline(file, name)
}

// NoContent sends a response with no body and a status code.
func (c *Context) NoContent(code int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.NoContent(code)
}

// Redirect redirects the request to a provided URL with status code.
func (c *Context) Redirect(code int, url string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Redirect(code, url)
}

// Error invokes the registered HTTP error handler. Generally used by middleware.
func (c *Context) Error(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.Error(err)
}

// Handler returns the matched handler by router.
func (c *Context) Handler() echo.HandlerFunc {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Handler()
}

// SetHandler sets the matched handler by router.
func (c *Context) SetHandler(h echo.HandlerFunc) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.SetHandler(h)
}

// Logger returns the `Logger` instance.
func (c *Context) Logger() echo.Logger {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Logger()
}

// SetLogger Set the logger
func (c *Context) SetLogger(l echo.Logger) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.SetLogger(l)
}

// Echo returns the `Echo` instance.
func (c *Context) Echo() *echo.Echo {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ec.Echo()
}

// Reset resets the context after request completes. It must be called along
// with `Echo#AcquireContext()` and `Echo#ReleaseContext()`.
// See `Echo#ServeHTTP()`
func (c *Context) Reset(r *http.Request, w http.ResponseWriter) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ec.Reset(r, w)
}
