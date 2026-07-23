package model

type Plan struct {
	Image   Image
	Mirrors []Mirror
	Targets []Target
}
