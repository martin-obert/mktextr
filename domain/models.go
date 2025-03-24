package domain

type Task struct {
	Id string
}

type RenderImageTask struct {
	Task
	WorldId string
	X       int
	Y       int
}

type TextureRef struct {
	Id      string
	Uri     string
	WorldId string
	X       int
	Y       int
}
