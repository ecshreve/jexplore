package jexplore

import (
	"embed"
	"io/fs"
)

var (
	//go:embed templates/Routes.gotsx
	//go:embed templates/Diagram.gotsx
	_customRefineTemplates embed.FS

	//go:embed templates/Main.tmpl
	_customServerTemplates embed.FS
)

type CustomRefine struct{}

var DefaultCustomTemplate = CustomRefine{}

func (c CustomRefine) GetName() string {
	return "custom-refine"
}

func (c CustomRefine) GetFS() fs.FS {
	return _customRefineTemplates
}

func (c CustomRefine) GetTemplates() []string {
	return []string{
		"templates/Diagram.gotsx",
		"templates/Routes.gotsx",
	}
}

type CustomServer struct{}

var DefaultCustomServerTemplate = CustomServer{}

func (c CustomServer) GetName() string {
	return "custom-server"
}

func (c CustomServer) GetFS() fs.FS {
	return _customServerTemplates
}

func (c CustomServer) GetTemplates() []string {
	return []string{
		"templates/Main.tmpl",
	}
}
