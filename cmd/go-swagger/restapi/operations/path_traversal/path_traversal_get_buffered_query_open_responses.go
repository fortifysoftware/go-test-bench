// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PathTraversalGetBufferedQueryOpenOKCode is the HTTP code returned for type PathTraversalGetBufferedQueryOpenOK
const PathTraversalGetBufferedQueryOpenOKCode int = 200

/*
PathTraversalGetBufferedQueryOpenOK returns the rendered response as a string

swagger:response pathTraversalGetBufferedQueryOpenOK
*/
type PathTraversalGetBufferedQueryOpenOK struct {

	/*The response when succesful query happens
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPathTraversalGetBufferedQueryOpenOK creates PathTraversalGetBufferedQueryOpenOK with default headers values
func NewPathTraversalGetBufferedQueryOpenOK() *PathTraversalGetBufferedQueryOpenOK {

	return &PathTraversalGetBufferedQueryOpenOK{}
}

// WithPayload adds the payload to the path traversal get buffered query open o k response
func (o *PathTraversalGetBufferedQueryOpenOK) WithPayload(payload string) *PathTraversalGetBufferedQueryOpenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the path traversal get buffered query open o k response
func (o *PathTraversalGetBufferedQueryOpenOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PathTraversalGetBufferedQueryOpenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
PathTraversalGetBufferedQueryOpenDefault Error occured

swagger:response pathTraversalGetBufferedQueryOpenDefault
*/
type PathTraversalGetBufferedQueryOpenDefault struct {
	_statusCode int
}

// NewPathTraversalGetBufferedQueryOpenDefault creates PathTraversalGetBufferedQueryOpenDefault with default headers values
func NewPathTraversalGetBufferedQueryOpenDefault(code int) *PathTraversalGetBufferedQueryOpenDefault {
	if code <= 0 {
		code = 500
	}

	return &PathTraversalGetBufferedQueryOpenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the path traversal get buffered query open default response
func (o *PathTraversalGetBufferedQueryOpenDefault) WithStatusCode(code int) *PathTraversalGetBufferedQueryOpenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the path traversal get buffered query open default response
func (o *PathTraversalGetBufferedQueryOpenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PathTraversalGetBufferedQueryOpenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
