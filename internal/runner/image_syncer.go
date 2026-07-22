package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/yyysay/registry-sync/internal/task"
)

type ImageSyncer struct {
}

func NewImageSyncer() *ImageSyncer {
	return &ImageSyncer{}
}

func (i *ImageSyncer) Run(
	tasks []*task.Task,
) error {

	config := "images:\n"

	for _, t := range tasks {

		config += fmt.Sprintf(
			"  - source: %s\n    destination: %s\n",
			t.Source.Reference,
			t.Target.Reference,
		)
	}

	file, err := os.CreateTemp(
		"",
		"image-syncer-*.yaml",
	)

	if err != nil {
		return err
	}

	defer os.Remove(
		file.Name(),
	)

	if _, err := file.WriteString(config); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	cmd := exec.Command(
		"image-syncer",
		"--config",
		file.Name(),
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf(
			"image-syncer failed: %w\n%s",
			err,
			string(output),
		)
	}

	return nil
}
