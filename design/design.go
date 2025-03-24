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
		Result(TextureReferencePayload, "Texture reference")

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
		})
		Payload(func() {
			Field(1, "x", Int, "Texture X")
			Field(1, "y", Int, "Texture y")
			Field(1, "worldId", String, "WorldId")
			Required("x", "y", "worldId")
		})
		Result(TextureReferencePayload, "Texture reference")
	})

	Method("completeTask", func() {
		Payload(TaskCompletionPayload, "Complete task")
		Result(Empty)

		HTTP(func() {
			GET("/tasks/{taskId}/complete")
		})
	})
})

var TextureReferencePayload = Type("TextureReferencePayload", func() {
	Description("The texture unique reference")

	Attribute("id", String, "Unique identifier")
})

var TaskCompletionPayload = Type("TaskCompletionPayload", func() {
	Description("The texture payload")
	Attribute("taskId", String, "Unique identifier")
	Attribute("texture", Bytes, "The texture")
	Required("taskId", "texture")
})
