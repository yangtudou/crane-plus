package source

type Mode int

const (
	Default Mode = iota
)

type Source struct {
	Registry string
	Mode     Mode
}

func New(registry string, mode Mode) *Source {
	return &Source{
		Registry: registry,
		Mode:     mode,
	}
}
