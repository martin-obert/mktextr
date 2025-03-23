package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("mktextr", func() {
	Description("Texture store")

	Method("getTextureById", func() {
		Payload(String, "Texture unique identifier")
		Result(TextureReferencePayload, "Texture reference")

		HTTP(func() {
			GET("/textures/{id}")
		})
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
