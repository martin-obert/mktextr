// Code generated by goa v3.20.0, DO NOT EDIT.
//
// mktextr client HTTP transport
//
// Command:
// $ goa gen mktextr/design

package client

import (
	"context"
	"mime/multipart"
	mktextr "mktextr/gen/mktextr"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the mktextr service endpoint HTTP clients.
type Client struct {
	// GetTextureByID Doer is the HTTP client used to make requests to the
	// getTextureById endpoint.
	GetTextureByIDDoer goahttp.Doer

	// GetTextureByCoordinates Doer is the HTTP client used to make requests to the
	// getTextureByCoordinates endpoint.
	GetTextureByCoordinatesDoer goahttp.Doer

	// CompleteTask Doer is the HTTP client used to make requests to the
	// completeTask endpoint.
	CompleteTaskDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// MktextrCompleteTaskEncoderFunc is the type to encode multipart request for
// the "mktextr" service "completeTask" endpoint.
type MktextrCompleteTaskEncoderFunc func(*multipart.Writer, *mktextr.CompleteTaskPayload) error

// NewClient instantiates HTTP clients for all the mktextr service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetTextureByIDDoer:          doer,
		GetTextureByCoordinatesDoer: doer,
		CompleteTaskDoer:            doer,
		RestoreResponseBody:         restoreBody,
		scheme:                      scheme,
		host:                        host,
		decoder:                     dec,
		encoder:                     enc,
	}
}

// GetTextureByID returns an endpoint that makes HTTP requests to the mktextr
// service getTextureById server.
func (c *Client) GetTextureByID() goa.Endpoint {
	var (
		decodeResponse = DecodeGetTextureByIDResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetTextureByIDRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetTextureByIDDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mktextr", "getTextureById", err)
		}
		return decodeResponse(resp)
	}
}

// GetTextureByCoordinates returns an endpoint that makes HTTP requests to the
// mktextr service getTextureByCoordinates server.
func (c *Client) GetTextureByCoordinates() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetTextureByCoordinatesRequest(c.encoder)
		decodeResponse = DecodeGetTextureByCoordinatesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetTextureByCoordinatesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetTextureByCoordinatesDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mktextr", "getTextureByCoordinates", err)
		}
		return decodeResponse(resp)
	}
}

// CompleteTask returns an endpoint that makes HTTP requests to the mktextr
// service completeTask server.
func (c *Client) CompleteTask(mktextrCompleteTaskEncoderFn MktextrCompleteTaskEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodeCompleteTaskRequest(NewMktextrCompleteTaskEncoder(mktextrCompleteTaskEncoderFn))
		decodeResponse = DecodeCompleteTaskResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCompleteTaskRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CompleteTaskDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mktextr", "completeTask", err)
		}
		return decodeResponse(resp)
	}
}
