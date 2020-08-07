package main

import (
	"github.com/cqbqdd11519/apiserver-test/pkg/server"
	"github.com/operator-framework/operator-sdk/pkg/log/zap"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

func main() {
	logf.SetLogger(zap.Logger())

	s := server.New()
	s.Start()
}
