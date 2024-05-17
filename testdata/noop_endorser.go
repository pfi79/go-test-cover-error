package main

import (
	"github.com/pfi79/go-test-cover-error/a"
)

type NoOpEndorser struct {
}

func (*NoOpEndorser) Eval(int) (int, error) {
	return 0, nil
}

func (*NoOpEndorser) Init(dependencies ...a.Dependency) error {
	return nil
}

type NoOpEndorserFactory struct {
}

func (*NoOpEndorserFactory) New() a.Plugin {
	return &NoOpEndorser{}
}

// NewPluginFactory is the function ran by the plugin infrastructure to create an endorsement plugin factory.
func NewPluginFactory() a.PluginFactory {
	return &NoOpEndorserFactory{}
}
