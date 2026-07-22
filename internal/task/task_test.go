package task

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
)

func TestGenerate(t *testing.T) {
	images := []*image.Image{}

	img1, err := image.Parse("cloudflare/cloudflared:latest")
	if err != nil {
		t.Fatal(err)
	}

	img2, err := image.Parse("ghcr.io/sagernet/sing-box:latest")
	if err != nil {
		t.Fatal(err)
	}

	images = append(images, img1, img2)

	dst := destination.New(
		"registry.cn-hangzhou.aliyuncs.com/myspace",
		mapper.New(mapper.Basename),
	)

	tasks := Generate(images, dst)

	if len(tasks) != 2 {
		t.Fatalf("tasks = %d, want 2", len(tasks))
	}

	if tasks[0].Target.Reference != "registry.cn-hangzhou.aliyuncs.com/myspace/cloudflared:latest" {
		t.Fatalf("target[0] = %q", tasks[0].Target.Reference)
	}

	if tasks[1].Target.Reference != "registry.cn-hangzhou.aliyuncs.com/myspace/sing-box:latest" {
		t.Fatalf("target[1] = %q", tasks[1].Target.Reference)
	}
}
