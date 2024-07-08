package handlers

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"nba-task-main/test/server"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

var s *server.Suite

func setupSuite(ctx context.Context) {
	s = server.NewSuite()
	s.Setup(ctx)
	DeferCleanup(s.Shutdown)
}
