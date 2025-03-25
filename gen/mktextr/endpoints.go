// Code generated by goa v3.20.0, DO NOT EDIT.
//
// mktextr endpoints
//
// Command:
// $ goa gen mktextr/design

package mktextr

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "mktextr" service endpoints.
type Endpoints struct {
	GetTaskQueue            goa.Endpoint
	GetTextureByCoordinates goa.Endpoint
	CompleteTask            goa.Endpoint
}

// NewEndpoints wraps the methods of the "mktextr" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetTaskQueue:            NewGetTaskQueueEndpoint(s),
		GetTextureByCoordinates: NewGetTextureByCoordinatesEndpoint(s),
		CompleteTask:            NewCompleteTaskEndpoint(s),
	}
}

// Use applies the given middleware to all the "mktextr" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetTaskQueue = m(e.GetTaskQueue)
	e.GetTextureByCoordinates = m(e.GetTextureByCoordinates)
	e.CompleteTask = m(e.CompleteTask)
}

// NewGetTaskQueueEndpoint returns an endpoint function that calls the method
// "GetTaskQueue" of service "mktextr".
func NewGetTaskQueueEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.GetTaskQueue(ctx)
	}
}

// NewGetTextureByCoordinatesEndpoint returns an endpoint function that calls
// the method "getTextureByCoordinates" of service "mktextr".
func NewGetTextureByCoordinatesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetTextureByCoordinatesPayload)
		return s.GetTextureByCoordinates(ctx, p)
	}
}

// NewCompleteTaskEndpoint returns an endpoint function that calls the method
// "completeTask" of service "mktextr".
func NewCompleteTaskEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CompleteTaskPayload)
		return nil, s.CompleteTask(ctx, p)
	}
}
