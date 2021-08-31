package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"strings"
)

type (
	Params interface {
		ContentType() string
		String() string
		Body() (io.Reader, error)
	}
)

// JsonParams
type JsonParams struct {
	payload interface{}
}

func (p *JsonParams) ContentType() string {
	return JsonContentType
}

func (p *JsonParams) String() string {
	b, _ := json.Marshal(p.payload)
	return string(b)
}

func (p *JsonParams) Body() (io.Reader, error) {
	if v, ok := p.payload.(string); ok {
		return strings.NewReader(v), nil
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(p.payload); err != nil {
		return nil, err
	}
	return buf, nil
}

// FormParams
type FormParams struct {
	payload interface{}
}

func (p *FormParams) ContentType() string {
	return FormContentType
}

func (p *FormParams) String() string {
	urlValues, _ := ParseUrlValues(p.payload)
	return urlValues.Encode()
}

func (p *FormParams) Body() (io.Reader, error) {
	if v, ok := p.payload.(string); ok {
		return strings.NewReader(v), nil
	}
	urlValues, err := ParseUrlValues(p.payload)
	if err != nil {
		return nil, err
	}
	buf := strings.NewReader(urlValues.Encode())
	return buf, nil
}

// UploadParams
type UploadParams struct {
	FieldName string
	FileName  string
	Reader    io.Reader
	Data      map[string]string
	writer    *multipart.Writer
}

func (p *UploadParams) ContentType() string {
	return p.writer.FormDataContentType()
}

func (p *UploadParams) String() string {
	return p.FieldName + ":" + p.FileName
}

func (p *UploadParams) Body() (io.Reader, error) {
	body := &bytes.Buffer{}
	p.writer = multipart.NewWriter(body)
	fileWriter, err := p.writer.CreateFormFile(p.FieldName, p.FileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, p.Reader)
	if err != nil {
		return nil, err
	}
	if len(p.Data) > 0 {
		for k, v := range p.Data {
			p.writer.WriteField(k, v)
		}
	}
	p.writer.Close()
	return body, nil
}
