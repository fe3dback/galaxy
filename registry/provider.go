package registry

type Provider struct {
	Registry *Registry
}

type Flags struct {
	IsProfiling   bool
	ProfilingPort int
	FullScreen    bool
}

func NewProvider(flags Flags) *Provider {
	return &Provider{
		Registry: makeRegistry(flags),
	}
}
