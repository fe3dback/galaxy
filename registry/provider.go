package registry

type Provider struct {
	Registry *Registry
}

type Flags struct {
	IsProfiling   bool
	ProfilingPort int
}

func NewProvider(flags Flags) *Provider {
	return &Provider{
		Registry: makeRegistry(flags),
	}
}
