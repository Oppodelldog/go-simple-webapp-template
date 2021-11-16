package assets

import (
	"embed"
	_ "embed"
)

//go:embed templates/*
var Templates embed.FS

//go:embed css/*
var Styles embed.FS

//go:embed icons/*
var Icons embed.FS

//go:embed js/*
var Javascript embed.FS