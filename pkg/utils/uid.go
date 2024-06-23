package utils

import "github.com/segmentio/ksuid"

func UniqueID() string {
	return ksuid.New().String()
}
