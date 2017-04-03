package storage

import "errors"

var (
	ErrNotFound = errors.New("not found!")
)

// Интерфейс нашего хранилища
type DB interface {
	Set(key string, val []byte) error

	Get(key string) ([]byte, error)

	Delete(key string) error

	StartCleaning(timeout int)

	StopCleaning()
}
