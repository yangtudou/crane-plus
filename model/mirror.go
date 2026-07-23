package model

type MirrorType string

const (
	MirrorExact   MirrorType = ""
	MirrorFlatten MirrorType = "flatten"
)

type Mirror struct {
	Name string
	URL  string
	Type MirrorType
	Auth *Auth
}
