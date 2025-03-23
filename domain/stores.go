package domain

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
)

type ITextureStorage interface {
	StoreTexture(textureId uuid.UUID, rawData []byte) error
	GetTextureById(textureId uuid.UUID) ([]byte, error)
	DeleteTextureById(textureId uuid.UUID) error
}

type LocalTextureStorage struct {
	rootFolder string
}

func (t LocalTextureStorage) DeleteTextureById(textureId uuid.UUID) error {
	var path = t.getFilePath(textureId)
	return os.Remove(path)
}

func (t LocalTextureStorage) getFilePath(textureId uuid.UUID) string {
	return filepath.Join(t.rootFolder, fmt.Sprintf("%s.png", textureId.String()))
}

func (t LocalTextureStorage) GetTextureById(textureId uuid.UUID) ([]byte, error) {
	var path = t.getFilePath(textureId)

	// Read binary data from file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Open error:", err)
		return nil, err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {

		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Read error:", err)
		return nil, err
	}

	return data, nil
}

func (t LocalTextureStorage) StoreTexture(textureId uuid.UUID, rawData []byte) error {
	// Create (or overwrite) the file
	var path = t.getFilePath(textureId)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
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
		return err
	}

	return nil
}

func NewLocalTextureStorage(rootFolder string) ITextureStorage {
	return &LocalTextureStorage{rootFolder: rootFolder}
}
