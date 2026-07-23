package copier

import (
	"context"
	"fmt"
	"strings"

	"registry-sync/engine"

	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type CraneCopier struct {
	copyFunc func(
		string,
		string,
		...crane.Option,
	) error
}

func New() *CraneCopier {

	return &CraneCopier{
		copyFunc: crane.Copy,
	}
}

func (c *CraneCopier) Copy(
	ctx context.Context,
	source string,
	target string,
	platform []string,
) error {

	fmt.Println("CRANE COPY")
	fmt.Println("==========")

	fmt.Println("SOURCE:")
	fmt.Println(" ", source)

	fmt.Println()

	fmt.Println("TARGET:")
	fmt.Println(" ", target)

	fmt.Println()

	opts := []crane.Option{
		crane.WithContext(ctx),
	}

	if len(platform) > 0 {

		fmt.Println("PLATFORM:")

		for _, p := range platform {
			fmt.Println(" ", p)
		}

		fmt.Println()

		parts := strings.Split(
			platform[0],
			"/",
		)

		if len(parts) == 2 {

			opts = append(
				opts,
				crane.WithPlatform(
					&v1.Platform{
						OS:           parts[0],
						Architecture: parts[1],
					},
				),
			)
		}
	}

	fmt.Println("ENGINE:")
	fmt.Println(" crane.Copy")

	fmt.Println()

	err := c.copyFunc(
		source,
		target,
		opts...,
	)

	if err != nil {
		fmt.Println("STATUS:")
		fmt.Println(" FAILED")

		fmt.Println("ERROR:")
		fmt.Println(" ", err)

		fmt.Println()
		return err
	}

	fmt.Println("STATUS:")
	fmt.Println(" SUCCESS")
	fmt.Println()

	return nil
}

var _ engine.Copier = (*CraneCopier)(nil)
