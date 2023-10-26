package storage

import "github.com/mkdemkov/storage/internal/file"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) file.File {
	return file.File{
		Name: filename,
		Data: blob,
	}
}
