package config

const (
	// DefaultUsage is the default description
	DefaultUsage = "[no usage description provided]"
)

// Base defines a base set of configuration
type Base struct {
	Shorthand string
	Usage     string
}

func (b *Base) GetShorthand() string {
	return b.Shorthand
}

func (b *Base) GetUsage() string {
	return b.Usage
}
