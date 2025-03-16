package utils

import "github.com/segmentio/ksuid"

type RandomUtils struct {
}

func NewRandomUtils() *RandomUtils {
	return &RandomUtils{}
}

func (rut *RandomUtils) UniqueID() string {
	return ksuid.New().String()
}
