package main

type provider struct {
	registry *registry
}

func newProvider() *provider {
	return &provider{
		registry: makeRegistry(),
	}
}
