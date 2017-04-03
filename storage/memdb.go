package storage

import (
	"sync"
	"time"
)

//РЕАЛИЗАЦИЯ ПОВЕДЕНИЯ ХРАНИЛИЩА

//Врямя жизни ключа
const BOXLIFETIME = 30

//Ячейка нашего хралища
type Box struct {
	val         []byte
	boxLifetime int64
	Updated     int64
}

//Хранилище
type MemoryDB struct {
	syn        sync.RWMutex
	d          map[string]*Box
	cleanTimer *time.Timer
}

var instance DB
var once sync.Once

//Singleton, та как хранилище всегда в единственном экземпляре
func CreateMemoryDB() DB {

	once.Do(func() {
		instance = &MemoryDB{d: make(map[string]*Box)}
	})
	return instance
}

//Устанавливает значение для заданного ключа
func (s *MemoryDB) Set(key string, val []byte) error {

	s.syn.Lock()
	defer s.syn.Unlock()
	s.d[key] = &Box{val, BOXLIFETIME, time.Now().Unix()}
	return nil
}

//Возвращает значение по заданному ключу
func (s *MemoryDB) Get(key string) ([]byte, error) {
	s.syn.RLock()
	defer s.syn.RUnlock()
	v, ok := s.d[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v.val, nil
}

//Удаляет ключ со значением, если значение для заданного ключа отсутствует, то возвращает ошибку
func (s *MemoryDB) Delete(key string) error {
	s.syn.Lock()
	defer s.syn.Unlock()
	_, ok := s.d[key]
	if !ok {
		return ErrNotFound
	}
	delete(s.d, key)

	return nil
}

//Удаляет ключ со значением, если значение не обновлялось больше BOXLIFETIME
func (s *MemoryDB) clean() {
	s.syn.Lock()
	defer s.syn.Unlock()
	now := time.Now().Unix()
	for key, boxRecord := range s.d {
		if boxRecord.Updated < (now - boxRecord.boxLifetime) {
			delete(s.d, key)
		}
	}
}

//Метод запускает очистку хранилища каждые timeout секунд
func (s *MemoryDB) StartCleaning(timeout int) {
	defer func() {
		s.cleanTimer = time.AfterFunc(time.Duration(timeout)*time.Second, func() { s.StartCleaning(timeout) })
	}()
	s.clean()
}

//Метод запускает очистку хранилища каждые timeout секунд
func (s *MemoryDB) StopCleaning() {
	if s.cleanTimer != nil {
		s.cleanTimer.Stop()
	}
}
