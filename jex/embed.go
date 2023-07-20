package jexplore

import (
	"embed"
	"io/fs"
)

var (
	//go:embed templates/*
	_customTemplates embed.FS
)

type CustomTemplate struct{}

var DefaultCustomTemplate = CustomTemplate{}

func (c CustomTemplate) GetName() string {
	return "custom"
}

func (c CustomTemplate) GetFS() fs.FS {
	return _customTemplates
}

func (c CustomTemplate) GetTemplates() []string {
	return []string{
		"templates/Diagram.gotsx",
	}
}
