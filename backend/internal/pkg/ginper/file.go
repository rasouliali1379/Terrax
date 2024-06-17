package ginper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

var (
	ErrInvalidFile          = errors.New("invalid file")
	ErrNoFileInRequest      = errors.New("there's no file in the request")
	ErrInvalidFileExtension = errors.New("invalid file extension")
)

type File struct {
	Content   io.Reader
	Extension string
	Name      string
}

// ParseFile parses file from form. You can pass your preferred extensions for validation or don't send any extension to
// accept all extensions.
func ParseFile(c *gin.Context, key string, extensions ...string) (*File, error) {

	formFile, err := c.FormFile(key)
	if err != nil {
		if errors.Is(err, http.ErrNotMultipart) {
			return nil, ErrNoFileInRequest
		}
		return nil, err
	}

	extExist := func(ext string) bool {
		if len(extensions) == 0 {
			return true
		}

		for i := range extensions {
			if extensions[i] == ext {
				return true
			}
		}
		return false
	}

	name := formFile.Filename
	splitted := strings.Split(name, ".")
	ext := splitted[len(splitted)-1]

	if len(splitted) < 2 || !extExist(ext) {
		return nil, ErrInvalidFileExtension
	}

	reader, err := formFile.Open()
	if err != nil {
		return nil, ErrInvalidFile
	}

	return &File{Content: reader, Extension: ext, Name: name}, nil
}
