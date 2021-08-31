package httpclient

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	ContentType     = "Content-Type"
	JsonContentType = "application/json"
	FormContentType = "application/x-www-form-urlencoded"
)

type (
	HttpClient struct {
		client  *http.Client
		baseUrl string
		header  http.Header
		debug   bool
	}

	LogData struct {
		Method     string      `json:"method"`
		Url        string      `json:"url"`
		ReqHeader  http.Header `json:"req_header"`
		ReqBody    string      `json:"req_body"`
		StatusCode int         `json:"status_code"`
		RespHeader http.Header `json:"resp_header"`
		Error      error       `json:"error"`
	}
)

var (
	DefaultClient = New()
)

func New() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		header: make(http.Header),
	}
}

func (c *HttpClient) Base(path string) *HttpClient {
	c.baseUrl = path
	return c
}

func (c *HttpClient) SetDebug(debug bool) {
	c.debug = debug
}

func (c *HttpClient) SetTimeout(t time.Duration) {
	c.client.Timeout = t
}

func (c *HttpClient) SetHeader(key, value string) *HttpClient {
	c.header.Set(key, value)
	return c
}

func (c *HttpClient) ResetHeader() {
	c.header = make(http.Header)
}

func (c *HttpClient) Get(path string, data interface{}) (*Response, error) {
	rawUrl := parseReqUrl(c.baseUrl, path)
	reqUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	// urlValues
	urlValues, err := ParseUrlValues(data)
	if err != nil {
		return nil, err
	}
	// add query string
	err = parseQueryString(reqUrl, urlValues)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	parseHeader(req, c.header)
	// logData
	logData := &LogData{}
	return c.send(req, logData)
}

func (c *HttpClient) PostForm(path string, data interface{}) (*Response, error) {
	return c.Post(path, &FormParams{data})
}

func (c *HttpClient) PostJson(path string, data interface{}) (*Response, error) {
	return c.Post(path, &JsonParams{data})
}

func (c *HttpClient) PostUpload(path string, params *UploadParams) (*Response, error) {
	return c.Post(path, params)
}

func (c *HttpClient) Post(path string, params Params) (*Response, error) {
	rawUrl := parseReqUrl(c.baseUrl, path)
	reqUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	body, err := params.Body()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, reqUrl.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(ContentType, params.ContentType())
	parseHeader(req, c.header)
	// logData
	logData := &LogData{}
	if c.debug {
		logData.ReqBody = params.String()
	}
	return c.send(req, logData)
}

func (c *HttpClient) send(req *http.Request, logData *LogData) (*Response, error) {
	defer func() {
		if c.debug {
			logData.Method = req.Method
			logData.Url = req.URL.String()
			logData.ReqHeader = req.Header
			logByte, _ := json.MarshalIndent(logData, "", "  ")
			log.Printf("httpclient send:\n%s\n", string(logByte))
		}
	}()
	resp, err := c.client.Do(req)
	if err != nil {
		logData.Error = err
		return nil, err
	}
	if c.debug {
		logData.StatusCode = resp.StatusCode
		logData.RespHeader = resp.Header
	}
	newResp := NewResponse(resp)
	return newResp, nil
}

func parseHeader(req *http.Request, header http.Header) {
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}

func parseReqUrl(base, path string) string {
	baseUrl, baseErr := url.Parse(base)
	pathUrl, pathErr := url.Parse(path)
	if pathUrl.Host == "" {
		return base + path
	}
	var reqUrl = path
	if baseErr == nil && pathErr == nil {
		reqUrl = baseUrl.ResolveReference(pathUrl).String()
	}
	return reqUrl
}

func parseQueryString(reqUrl *url.URL, queryValues url.Values) error {
	urlValues, err := url.ParseQuery(reqUrl.RawQuery)
	if err != nil {
		return err
	}
	for key, values := range queryValues {
		for _, value := range values {
			urlValues.Add(key, value)
		}
	}
	reqUrl.RawQuery = urlValues.Encode()
	return nil
}

func Get(path string, data interface{}) (*Response, error) {
	return DefaultClient.Get(path, data)
}

func PostForm(path string, data interface{}) (*Response, error) {
	return DefaultClient.PostForm(path, data)
}

func PostJson(path string, data interface{}) (*Response, error) {
	return DefaultClient.PostJson(path, data)
}

func PostUpload(path string, params *UploadParams) (*Response, error) {
	return DefaultClient.PostUpload(path, params)
}

func SetDebug(debug bool) {
	DefaultClient.SetDebug(debug)
}

func SetHeader(key string, value string) *HttpClient {
	return DefaultClient.SetHeader(key, value)
}

func ResetHeader() {
	DefaultClient.ResetHeader()
}

func SetTimeout(t time.Duration) {
	DefaultClient.SetTimeout(t)
}
