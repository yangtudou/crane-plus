package image

type Image struct {
	Reference string

	Registry  string
	Namespace string
	Name      string
	Tag       string
	Digest    string
}

func New(
	registry string,
	namespace string,
	name string,
	tag string,
) *Image {
	fullName := name

	if namespace != "" {
		fullName = namespace + "/" + name
	}

	return &Image{
		Registry:  NormalizeRegistry(registry),
		Namespace: namespace,
		Name:      fullName,
		Tag:       tag,
		Reference: buildReference(
			registry,
			namespace,
			name,
			tag,
		),
	}
}

func (i *Image) Repository() string {
	return i.Name
}

func (i *Image) FullName() string {
	name := i.Name

	if i.Tag != "" {
		name += ":" + i.Tag
	}

	if i.Digest != "" {
		name += "@" + i.Digest
	}

	return name
}

func buildReference(
	registry string,
	namespace string,
	name string,
	tag string,
) string {
	ref := ""

	if registry != "" {
		ref += registry + "/"
	}

	if namespace != "" {
		ref += namespace + "/"
	}

	ref += name

	if tag != "" {
		ref += ":" + tag
	}

	return ref
}
