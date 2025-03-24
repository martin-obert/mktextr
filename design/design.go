package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("mktextr", func() {
	Description("Texture store")

	Method("getTextureById", func() {
		Payload(func() {
			Field(1, "id", String, "Texture ID")
			Required("id")
		})
		Result(Empty)

		HTTP(func() {
			GET("/textures/{id}")
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
				Tag("status_code", "ok")
			})
			Response(StatusAccepted, func() {
				Tag("status_code", "accepted")
			})
			Response(StatusBadRequest)
		})
		Payload(func() {
			Field(1, "x", Int, "Texture X")
			Field(1, "y", Int, "Texture y")
			Field(1, "worldId", String, "WorldId")
			Required("x", "y", "worldId")
		})
		Result(GetTextureByCoordinatesResponse, func() {
			View("default")
		})
	})

	Method("completeTask", func() {
		Payload(func() {
			Field(1, "file", Bytes, "The file to upload", func() {
				Meta("struct:tag:encoding", "form")
			})
			Field(2, "filename", String, "Name of the file", func() {
				Meta("struct:tag:encoding", "form")
			})
			Field(3, "taskId", String, "ID of the task")
			Required("file", "filename", "taskId")
		})
		Result(Empty)

		HTTP(func() {
			PATCH("/tasks/{taskId}/complete")
			MultipartRequest()
			Response(StatusOK)
		})
	})
})

var GetTextureByCoordinatesResponse = ResultType("application/get-result", func() {
	Description("Texture set payload")
	Field(1, "status_code", String)

	Attributes(func() {
		Attribute("status_code", String)
		Attribute("taskId", String, "")
		Attribute("baseMapUrl", String, "")
		Attribute("contourMapUrl", String, "")
	})

	View("default", func() {
		Attribute("taskId", String, "")
		Attribute("baseMapUrl", String, "")
		Attribute("contourMapUrl", String, "")
	})
})
