// Code generated by goa v3.20.0, DO NOT EDIT.
//
// mktextr HTTP server encoders and decoders
//
// Command:
// $ goa gen mktextr/design

package server

import (
	"context"
	"errors"
	mktextr "mktextr/gen/mktextr"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetTaskQueueResponse returns an encoder for responses returned by the
// mktextr GetTaskQueue endpoint.
func EncodeGetTaskQueueResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*mktextr.GetTaskQueueResult)
		enc := encoder(ctx, w)
		body := NewGetTaskQueueResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetTextureByCoordinatesResponse returns an encoder for responses
// returned by the mktextr getTextureByCoordinates endpoint.
func EncodeGetTextureByCoordinatesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*mktextr.GetTextureByCoordinatesResponse)
		if res.TextureSetState != nil && *res.TextureSetState == "ready" {
			enc := encoder(ctx, w)
			body := NewGetTextureByCoordinatesOKResponseBody(res)
			w.WriteHeader(http.StatusOK)
			return enc.Encode(body)
		}
		if res.TextureSetState != nil && *res.TextureSetState == "processing" {
			enc := encoder(ctx, w)
			body := NewGetTextureByCoordinatesPartialContentResponseBody(res)
			w.WriteHeader(http.StatusPartialContent)
			return enc.Encode(body)
		}
		enc := encoder(ctx, w)
		body := NewGetTextureByCoordinatesBadRequestResponseBody(res)
		w.WriteHeader(http.StatusBadRequest)
		return enc.Encode(body)
	}
}

// DecodeGetTextureByCoordinatesRequest returns a decoder for requests sent to
// the mktextr getTextureByCoordinates endpoint.
func DecodeGetTextureByCoordinatesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			worldID string
			x       int
			y       int
			err     error
		)
		qp := r.URL.Query()
		worldID = qp.Get("worldId")
		if worldID == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("worldId", "query string"))
		}
		{
			xRaw := qp.Get("x")
			if xRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("x", "query string"))
			}
			v, err2 := strconv.ParseInt(xRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("x", xRaw, "integer"))
			}
			x = int(v)
		}
		{
			yRaw := qp.Get("y")
			if yRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("y", "query string"))
			}
			v, err2 := strconv.ParseInt(yRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("y", yRaw, "integer"))
			}
			y = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetTextureByCoordinatesPayload(worldID, x, y)

		return payload, nil
	}
}

// EncodeCompleteTaskResponse returns an encoder for responses returned by the
// mktextr completeTask endpoint.
func EncodeCompleteTaskResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeCompleteTaskRequest returns a decoder for requests sent to the mktextr
// completeTask endpoint.
func DecodeCompleteTaskRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var payload *mktextr.CompleteTaskPayload
		if err := decoder(r).Decode(&payload); err != nil {
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewMktextrCompleteTaskDecoder returns a decoder to decode the multipart
// request for the "mktextr" service "completeTask" endpoint.
func NewMktextrCompleteTaskDecoder(mux goahttp.Muxer, mktextrCompleteTaskDecoderFn MktextrCompleteTaskDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v any) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(**mktextr.CompleteTaskPayload)
			if err := mktextrCompleteTaskDecoderFn(mr, p); err != nil {
				return err
			}

			var (
				taskID string

				params = mux.Vars(r)
			)
			taskID = params["taskId"]
			(*p).TaskID = taskID
			return nil
		})
	}
}
