package mktextr

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
)

type ITextureStorage interface {
	StoreTexture(rawData []byte) (TextureRef, error)
	GetTextureById(textureId uuid.UUID) (TextureRef, error)
}

type LocalTextureStorage struct {
	settings LocalTextureStorageSettings
}

func (t LocalTextureStorage) getFilePath(textureId uuid.UUID) string {
	return filepath.Join(t.settings.rootFolder, textureId.String())
}
func (t LocalTextureStorage) GetTextureById(textureId uuid.UUID) (TextureRef, error) {
	var path = t.getFilePath(textureId)

	// Read binary data from file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Open error:", err)
		return ()
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
}

func (t LocalTextureStorage) StoreTexture(rawData []byte) (error) {
	var textureId = uuid.New()
	var textureRef = TextureRef{
		id:  textureId,
		uri: getFilePath(t.settings, textureId),
	}

	// Create (or overwrite) the file
	file, err := os.Create("output.bin")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return EmptyTextureRef(), err
	}

	defer func(file *os.File) {
		closeError := file.Close()
		if closeError != nil {
			fmt.Println("Error closing file:", closeError)
		}
	}(file)

	// Write binary data to the file
	_, err = file.Write(rawData)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return EmptyTextureRef(), err
	}

	return textureRef, nil
}

func NewLocalTextureStorage(settings LocalTextureStorageSettings) ITextureStorage {
	return &LocalTextureStorage{settings: settings}
}