package destination

import (
	"path"

	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
)

type Destination struct {
	Registry string
	Mapper   *mapper.Mapper
}

func New(registry string, m *mapper.Mapper) *Destination {
	return &Destination{
		Registry: registry,
		Mapper:   m,
	}
}

func (d *Destination) Map(src *image.Image) *image.Image {
	dst := d.Mapper.Map(src)

	dst.Registry = d.Registry
	dst.Reference = buildReference(dst)

	return dst
}

func buildReference(img *image.Image) string {
	ref := img.Registry + "/"

	ref += path.Clean(img.Name)

	if img.Tag != "" {
		ref += ":" + img.Tag
	}

	if img.Digest != "" {
		ref += "@" + img.Digest
	}

	return ref
}
