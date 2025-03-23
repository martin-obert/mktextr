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

// BuildCompleteTaskPayload builds the payload for the mktextr completeTask
// endpoint from CLI flags.
func BuildCompleteTaskPayload(mktextrCompleteTaskBody string, mktextrCompleteTaskTaskID string) (*mktextr.TaskCompletionPayload, error) {
	var err error
	var body CompleteTaskRequestBody
	{
		err = json.Unmarshal([]byte(mktextrCompleteTaskBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"texture\": \"RGljdGEgcXVhZSB2ZWxpdCB2b2x1cHRhdGVzIGRvbG9yLg==\"\n   }'")
		}
		if body.Texture == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("texture", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var taskID string
	{
		taskID = mktextrCompleteTaskTaskID
	}
	v := &mktextr.TaskCompletionPayload{
		Texture: body.Texture,
	}
	v.TaskID = taskID

	return v, nil
}
