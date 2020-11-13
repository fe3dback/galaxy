package registry

type Provider struct {
	Registry *Registry
}

type Flags struct {
	IsProfiling   bool
	ProfilingPort int
	FullScreen    bool
	Seed          int64
}

func NewProvider(flags Flags) *Provider {
	return &Provider{
		Registry: makeRegistry(flags),
	}
}
