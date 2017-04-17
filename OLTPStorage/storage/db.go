package storage

import "errors"

var (
	ErrNotFound = errors.New("not found!")
)

// Интерфейс нашего хранилища
type DB interface {
	Set(key string, val string) error

	Get(key string) ([]byte, error)

	Delete(key string) error
}
