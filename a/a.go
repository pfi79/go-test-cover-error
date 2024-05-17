package a

// Argument defines the argument for endorsement
type Argument interface {
	Dependency
	// Arg returns the bytes of the argument
	Arg() []byte
}

// Dependency marks a dependency passed to the Init() method
type Dependency interface{}

// Plugin endorses a proposal response
type Plugin interface {
	Eval(int) (int, error)

	Init(dependencies ...Dependency) error
}

// PluginFactory creates a new instance of a Plugin
type PluginFactory interface {
	New() Plugin
}
