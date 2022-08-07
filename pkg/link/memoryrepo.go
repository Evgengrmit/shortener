package link

import (
	"errors"
	"ozonTask/shorter"
	"sync"
)

type LinkMemory struct {
	memory map[string]string
	mx     sync.RWMutex
}

func NewLinkMemory() *LinkMemory {
	return &LinkMemory{memory: make(map[string]string)}
}

func (lm *LinkMemory) Add(originalURL string) (string, error) {
	shortURL, err := shorter.GetShort(originalURL)
	if err != nil {
		return "", err
	}
	lm.mx.Lock()
	_, ok := lm.memory[shortURL]
	lm.mx.Unlock()
	if !ok {
		lm.mx.Lock()
		lm.memory[shortURL] = originalURL
		lm.mx.Unlock()
	}
	return shortURL, nil
}
func (lm *LinkMemory) Get(shortURL string) (string, error) {
	lm.mx.RLock()
	originalURL, ok := lm.memory[shortURL]
	lm.mx.RUnlock()
	if ok {
		return originalURL, nil
	}
	return "", errors.New("short link is not valid")
}
