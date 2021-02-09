package controllers

import "sotru-web/utils"

var config utils.Config

// Globally sets config for the controllers package
func UseConfig(conf utils.Config) {
	config = conf
}
