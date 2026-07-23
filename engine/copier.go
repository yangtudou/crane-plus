package engine

import "context"

type Copier interface {
	Copy(
		ctx context.Context,
		source string,
		target string,
		platform []string,
	) error
}
