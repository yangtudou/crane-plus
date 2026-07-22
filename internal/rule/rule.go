package rule

import (
	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/source"
)

type Rule struct {
	Name        string
	Source      *source.Source
	Destination *destination.Destination
}

func New(
	name string,
	src *source.Source,
	dst *destination.Destination,
) *Rule {

	return &Rule{
		Name:        name,
		Source:      src,
		Destination: dst,
	}
}

// Match 判断规则是否匹配镜像
//
// Source 存在:
//
//	按 registry 匹配
//
// Source 不存在:
//
//	属于 mixed 规则，不参与主动匹配
func (r *Rule) Match(
	img *image.Image,
) bool {

	if r.Source == nil {
		return false
	}

	return r.Source.Match(img)
}

// IsMixed 判断是否为 mixed 默认规则
//
// mixed:
//
//	没有 source
//
// 用于 engine fallback
func (r *Rule) IsMixed() bool {

	return r.Source == nil
}
