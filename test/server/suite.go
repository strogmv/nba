package server

import (
	"context"

	. "github.com/onsi/ginkgo/v2"

	"nba-task-main/internal/app/aggregate"
)

type Suite struct {
	*aggregate.App
}

func NewSuite() *Suite {
	app, err := aggregate.NewApp()
	if err != nil {
		AbortSuite(err.Error())
	}

	return &Suite{
		App: app,
	}
}

func (s *Suite) Setup(_ context.Context) {
	go func() {
		defer GinkgoRecover()
		var err = <-s.App.Start()
		AbortSuite(err.Error())
	}()

}

func (s *Suite) TearDown(ctx context.Context) {
	if err := s.App.Shutdown(ctx); err != nil {
		AbortSuite(err.Error())
	}
}
