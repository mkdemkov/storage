package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mkdemkov/storage/internal/file"
)

type Storage struct {
	f map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		f: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}

	s.f[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.f[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileID)
	}

	return foundFile, nil
}
