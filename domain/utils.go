package domain

import (
	"encoding/base64"
	"strconv"
	"strings"
)

func TextureTaskId(worldId string, x int, y int, textureType TextureType) string {
	var v = strings.Join([]string{"render_image_task", worldId, strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(int(textureType))}, ":")
	return v
}
func EncodeTaskId(id string) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(id))
}

func DecodeTaskId(value string) (string, error) {
	b, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(value)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
