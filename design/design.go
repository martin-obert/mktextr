package design

import (
	. "goa.design/goa/v3/dsl"
)

type TextureType int

const (
	TaskCodeFieldName           = "texture_set_state"
	TextureSetStateProcessing   = "processing"
	TextureSetStateReady        = "ready"
	TextureSetSubStateRendering = "rendering"

	TextureTypeBaseMap TextureType = iota + 1
	TextureTypeContourMap
)

var _ = Service("mktextr", func() {
	Description("Texture store")
	Method("GetTaskQueue", func() {
		HTTP(func() {
			GET("/tasks")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("tasks", ArrayOf(String))
		})
	})
	Method("getTextureByCoordinates", func() {

		HTTP(func() {
			GET("/textures")

			// Query parameters for pagination
			Param("worldId", String, "WorldId")
			Param("x", Int, "Texture X")
			Param("y", Int, "Texture Y")

			Response(StatusOK, func() {
				Tag(TaskCodeFieldName, TextureSetStateReady)
			})
			Response(StatusPartialContent, func() {
				Tag(TaskCodeFieldName, TextureSetStateProcessing)
			})
			Response(StatusBadRequest)
		})
		Payload(func() {
			Field(1, "x", Int, "Texture X")
			Field(1, "y", Int, "Texture y")
			Field(1, "worldId", String, "WorldId")
			Required("x", "y", "worldId")
		})
		Result(GetTextureByCoordinatesResponse)
	})

	Method("completeTask", func() {
		Payload(func() {
			Field(1, "file", Bytes, "The file to upload", func() {
				Meta("struct:tag:encoding", "form")
			})
			Field(2, "extension", String, "ID of the task", func() {
				Meta("struct:tag:encoding", "form")
			})
			Field(3, "taskId", String, "ID of the task")
			Required("file", "extension", "taskId")
		})
		Result(Empty)

		HTTP(func() {
			PATCH("/tasks/{taskId}/complete")
			MultipartRequest()
			Response(StatusOK)
		})
	})
})

var GetTextureByCoordinatesResponse = Type("GetTextureByCoordinatesResponse", func() {
	Description("Texture set payload")
	Attribute(TaskCodeFieldName, String, "")
	Attribute("baseMapUrl", String, "")
	Attribute("contourMapUrl", String, "")
	Attribute("sub_state", String, "")

})
