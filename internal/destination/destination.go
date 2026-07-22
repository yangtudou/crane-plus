package destination

import (
	"strings"

	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
)

type Destination struct {
	Registry string
	mapper   *mapper.Mapper
}

func New(registry string, m *mapper.Mapper) *Destination {
	return &Destination{
		Registry: strings.TrimRight(registry, "/"),
		mapper:   m,
	}
}

func (d *Destination) Map(img *image.Image) *image.Image {
	mapped := d.mapper.Map(img)

	return &image.Image{
		Reference: buildReference(d.Registry, mapped),
		Registry:  d.Registry,
		Namespace: mapped.Namespace,
		Name:      mapped.Name,
		Tag:       mapped.Tag,
		Digest:    mapped.Digest,
	}
}

func buildReference(registry string, img *image.Image) string {
	ref := registry + "/" + img.Name

	if img.Tag != "" {
		ref += ":" + img.Tag
	}

	if img.Digest != "" {
		ref += "@" + img.Digest
	}

	return ref
}
