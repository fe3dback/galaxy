package vulkan

type (
	Config struct {
		debug bool
		vSync bool
	}

	Configure = func(*Config)
)

func NewConfig(opts ...Configure) Config {
	cfg := &Config{
		debug: false,
		vSync: false,
	}

	for _, configure := range opts {
		configure(cfg)
	}

	return *cfg
}

// WithDebug will print vulkan validation errors
// on stdout. Its require vulkan SDK to work
func WithDebug(enabled bool) Configure {
	return func(config *Config) {
		config.debug = enabled
	}
}

// WithVSync will use FIFO rendering
// true - vsync, good for mobile (small power consumption)
// false - low latency, high power consumption
func WithVSync(enabled bool) Configure {
	return func(config *Config) {
		config.vSync = enabled
	}
}
