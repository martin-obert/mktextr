// Code generated by goa v3.20.0, DO NOT EDIT.
//
// mktextr HTTP client CLI support package
//
// Command:
// $ goa gen mktextr/design

package client

import (
	"encoding/json"
	"fmt"
	mktextr "mktextr/gen/mktextr"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetTextureByIDPayload builds the payload for the mktextr getTextureById
// endpoint from CLI flags.
func BuildGetTextureByIDPayload(mktextrGetTextureByIDID string) (*mktextr.GetTextureByIDPayload, error) {
	var id string
	{
		id = mktextrGetTextureByIDID
	}
	v := &mktextr.GetTextureByIDPayload{}
	v.ID = id

	return v, nil
}

// BuildGetTextureByCoordinatesPayload builds the payload for the mktextr
// getTextureByCoordinates endpoint from CLI flags.
func BuildGetTextureByCoordinatesPayload(mktextrGetTextureByCoordinatesWorldID string, mktextrGetTextureByCoordinatesX string, mktextrGetTextureByCoordinatesY string) (*mktextr.GetTextureByCoordinatesPayload, error) {
	var err error
	var worldID string
	{
		worldID = mktextrGetTextureByCoordinatesWorldID
	}
	var x int
	{
		var v int64
		v, err = strconv.ParseInt(mktextrGetTextureByCoordinatesX, 10, strconv.IntSize)
		x = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for x, must be INT")
		}
	}
	var y int
	{
		var v int64
		v, err = strconv.ParseInt(mktextrGetTextureByCoordinatesY, 10, strconv.IntSize)
		y = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for y, must be INT")
		}
	}
	v := &mktextr.GetTextureByCoordinatesPayload{}
	v.WorldID = worldID
	v.X = x
	v.Y = y

	return v, nil
}

// BuildCompleteTaskPayload builds the payload for the mktextr completeTask
// endpoint from CLI flags.
func BuildCompleteTaskPayload(mktextrCompleteTaskBody string, mktextrCompleteTaskTaskID string) (*mktextr.CompleteTaskPayload, error) {
	var err error
	var body CompleteTaskRequestBody
	{
		err = json.Unmarshal([]byte(mktextrCompleteTaskBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"file\": \"TmloaWwgcmVydW0gcXVpIGV0IGRvbG9yIHByYWVzZW50aXVtIGxhYm9ydW0u\",\n      \"filename\": \"Et earum.\"\n   }'")
		}
		if body.File == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("file", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var taskID string
	{
		taskID = mktextrCompleteTaskTaskID
	}
	v := &mktextr.CompleteTaskPayload{
		File:     body.File,
		Filename: body.Filename,
	}
	v.TaskID = taskID

	return v, nil
}
