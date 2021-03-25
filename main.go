//+build !test

package main

import (
	"github.com/webmalc/vishleva-backend/common/config"
	"github.com/webmalc/vishleva-backend/common/logger"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	log.Info("init")
}
