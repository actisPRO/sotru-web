package controllers

import (
	"gopkg.in/boj/redistore.v1"
	"sotru-web/utils"
)

var (
	config utils.Config
	store  *redistore.RediStore
)

// Globally sets config for the controllers package
func UseConfig(conf utils.Config) {
	config = conf
}

// Globally sets session storage for the controllers package
func UseStore(newStore *redistore.RediStore) {
	store = newStore
}
