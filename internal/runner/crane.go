package runner

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/yyysay/registry-sync/internal/task"
)

type Crane struct {
	Platform string
	Workers  int
}

func NewCrane(
	platform string,
	workers int,
) *Crane {

	if workers <= 0 {
		workers = 1
	}

	return &Crane{
		Platform: platform,
		Workers:  workers,
	}
}

func (c *Crane) Run(
	tasks []*task.Task,
) error {

	total := len(tasks)

	jobs := make(chan *task.Task)

	errs := make(chan error, total)

	var wg sync.WaitGroup

	for i := 0; i < c.Workers; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for t := range jobs {

				fmt.Printf(
					"copying %s -> %s\n",
					t.Source.Reference,
					t.Target.Reference,
				)

				if err := c.copy(t); err != nil {
					errs <- err
					continue
				}

				fmt.Printf(
					"done %s\n",
					t.Target.Reference,
				)
			}

		}()

	}

	for _, t := range tasks {
		jobs <- t
	}

	close(jobs)

	wg.Wait()

	close(errs)

	for err := range errs {

		if err != nil {
			return err
		}

	}

	return nil
}

func (c *Crane) copy(
	t *task.Task,
) error {

	args := []string{
		"copy",
	}

	if c.Platform != "" {

		args = append(
			args,
			"--platform="+c.Platform,
		)

	}

	args = append(
		args,
		t.Source.Reference,
		t.Target.Reference,
	)

	cmd := exec.Command(
		"crane",
		args...,
	)

	output, err := cmd.CombinedOutput()

	if err != nil {

		return fmt.Errorf(
			"crane copy failed\nsource: %s\ntarget: %s\noutput: %s\nerror: %w",
			t.Source.Reference,
			t.Target.Reference,
			string(output),
			err,
		)

	}

	return nil
}
