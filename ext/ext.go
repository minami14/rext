package ext

import (
	"errors"
	"mime"
	"net/http"
	"os"
)

var (
	ErrExtensionNotFound = errors.New("extension not found")
)

func ExtensionFromFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	if _, err := f.Read(buf); err != nil {
		return nil, err
	}

	return ExtensionFromBinary(buf)
}

func ExtensionFromBinary(data []byte) ([]string, error) {
	ct := http.DetectContentType(data)
	if ct == "application/octet-stream" {
		return nil, ErrExtensionNotFound
	}

	ext, err := mime.ExtensionsByType(ct)
	if err != nil {
		return nil, err
	}

	if len(ext) == 0 {
		return nil, ErrExtensionNotFound
	}

	return ext, nil
}
