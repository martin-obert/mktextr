package domain

import (
	"encoding/base64"
	"strconv"
	"strings"
)

func TextureTaskId(worldId string, x int, y int) string {
	var v = strings.Join([]string{"render_image_task", worldId, strconv.Itoa(x), strconv.Itoa(y)}, ":")
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(v))
}
