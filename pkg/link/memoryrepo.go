package link

import (
	"errors"
	"sync"
)

type LinkMemory struct {
	memory map[string]string
	mx     sync.RWMutex
}

func NewLinkMemory() *LinkMemory {
	return &LinkMemory{memory: make(map[string]string)}
}

func (lm *LinkMemory) Add(original string) (string, error) {
	var shortURL string
	lm.mx.Lock()
	defer lm.mx.Unlock()
	if _, ok := lm.memory[shortURL]; !ok {
		lm.memory[shortURL] = original
	}
	return shortURL, nil
}
func (lm *LinkMemory) Get(short string) (string, error) {
	lm.mx.RLock()
	defer lm.mx.RUnlock()
	if originalURL, ok := lm.memory[short]; ok {
		return originalURL, nil
	}
	return "", errors.New("short link is not valid")
}
