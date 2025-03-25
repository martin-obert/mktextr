package mktextrapi

import (
	"io"
	"mime"
	"mime/multipart"
	"mktextr/gen/mktextr"
	"strings"
)

// MktextrCompleteTaskDecoderFunc implements the multipart decoder for service
// "mktextr" endpoint "completeTask". The decoder must populate the argument p
// after encoding.
func MktextrCompleteTaskDecoderFunc(mr *multipart.Reader, p **mktextr.CompleteTaskPayload) error {
	var res mktextr.CompleteTaskPayload
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		_, params, err := mime.ParseMediaType(p.Header.Get("Content-Disposition"))
		if err != nil {
			// can't process this entry, it probably isn't an image
			continue
		}

		disposition, _, err := mime.ParseMediaType(p.Header.Get("Content-Type"))
		// the disposition can be, for example 'image/jpeg' or 'video/mp4'
		// I want to support only image files!
		if err != nil || !strings.HasPrefix(disposition, "image/") {
			// can't process this entry, it probably isn't an image
			continue
		}

		if params["name"] == "file" {
			bytes, err := io.ReadAll(p)
			if err != nil {
				// can't process this entry, for some reason
				panic(err)
			}
			res.File = bytes
			res.Extension = disposition

			//imageUpload := images.ImageUpload{
			//	Type:  &disposition,
			//	Bytes: bytes,
			//	Name:  &filename,
			//}
			//res.Files = append(res.Files, &imageUpload)
		}
	}
	*p = &res
	return nil
}
