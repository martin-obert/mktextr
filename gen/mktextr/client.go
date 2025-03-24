// Code generated by goa v3.20.0, DO NOT EDIT.
//
// mktextr client
//
// Command:
// $ goa gen mktextr/design

package mktextr

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "mktextr" service client.
type Client struct {
	GetTextureByIDEndpoint          goa.Endpoint
	GetTextureByCoordinatesEndpoint goa.Endpoint
	CompleteTaskEndpoint            goa.Endpoint
}

// NewClient initializes a "mktextr" service client given the endpoints.
func NewClient(getTextureByID, getTextureByCoordinates, completeTask goa.Endpoint) *Client {
	return &Client{
		GetTextureByIDEndpoint:          getTextureByID,
		GetTextureByCoordinatesEndpoint: getTextureByCoordinates,
		CompleteTaskEndpoint:            completeTask,
	}
}

// GetTextureByID calls the "getTextureById" endpoint of the "mktextr" service.
func (c *Client) GetTextureByID(ctx context.Context, p *GetTextureByIDPayload) (res *TextureReferencePayload, err error) {
	var ires any
	ires, err = c.GetTextureByIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TextureReferencePayload), nil
}

// GetTextureByCoordinates calls the "getTextureByCoordinates" endpoint of the
// "mktextr" service.
func (c *Client) GetTextureByCoordinates(ctx context.Context, p *GetTextureByCoordinatesPayload) (res *TextureReferencePayload, err error) {
	var ires any
	ires, err = c.GetTextureByCoordinatesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TextureReferencePayload), nil
}

// CompleteTask calls the "completeTask" endpoint of the "mktextr" service.
func (c *Client) CompleteTask(ctx context.Context, p *TaskCompletionPayload) (err error) {
	_, err = c.CompleteTaskEndpoint(ctx, p)
	return
}
