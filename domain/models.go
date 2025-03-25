package domain

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"mktextr/design"
)

type TextureType int

const (
	TextureTypeBaseMap TextureType = iota + 1
	TextureTypeContourMap
)

type TextureState int

const (
	TextureMissing TextureState = iota + 1
	TextureReady
)

type TextureSetState int

const (
	TextureSetStateIncomplete TextureSetState = iota + 1
	TextureSetStateComplete
)

type Task struct {
	Id string
}

type RenderImageTask struct {
	Task
	Address     TextureAddress
	TextureType TextureType
}

func NewRenderBaseMapImageTask(address TextureAddress) RenderImageTask {
	return RenderImageTask{
		Task: Task{
			Id: TextureTaskId(address.WorldId, address.X, address.Y, TextureTypeBaseMap),
		},
		Address:     address,
		TextureType: TextureTypeBaseMap,
	}
}
func NewRenderContourMapImageTask(address TextureAddress) RenderImageTask {
	return RenderImageTask{
		Task: Task{
			Id: TextureTaskId(address.WorldId, address.X, address.Y, TextureTypeContourMap),
		},
		Address:     address,
		TextureType: TextureTypeContourMap,
	}
}

type TextureAddress struct {
	WorldId string
	X       int
	Y       int
}

func (a *TextureAddress) String() string {
	return fmt.Sprintf("%s:%d:%d", a.WorldId, a.X, a.Y)
}

type TextureSet struct {
	Id         string
	Address    TextureAddress
	TextureSet map[TextureType]TextureRef
}

func (ts *TextureSet) State() TextureSetState {
	for _, ref := range ts.TextureSet {
		if ref.State != TextureReady {
			return TextureSetStateIncomplete
		}
	}

	return TextureSetStateComplete
}

func (s TextureSetState) ToDesign() string {
	switch s {
	case TextureSetStateComplete:
		return design.TextureSetStateReady
	case TextureSetStateIncomplete:
		return design.TextureSetStateProcessing
	default:
		return "Unknown"
	}
}

func NewTextureSet(address TextureAddress) TextureSet {
	return TextureSet{
		Id:      bson.NewObjectID().Hex(),
		Address: address,
		TextureSet: map[TextureType]TextureRef{
			TextureTypeBaseMap:    NewMissingTextureRef(),
			TextureTypeContourMap: NewMissingTextureRef(),
		},
	}
}

type TextureRef struct {
	Uri    string
	FileId string
	State  TextureState
}

func NewMissingTextureRef() TextureRef {
	return TextureRef{
		Uri:    "",
		FileId: "",
		State:  TextureMissing,
	}
}
