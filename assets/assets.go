package assets

import (
	"github.com/Oppodelldog/go-simple-webapp-template/config"
	"github.com/go-playground/statics/static"
)

var Templates *static.Files
var Styles *static.Files
var Icons *static.Files
var Javascript *static.Files

func init() {
	var err error

	cfg := &static.Config{
		UseStaticFiles: config.UseStaticFiles,
		FallbackToDisk: false,
		AbsPkgPath:     config.AbsoluteAssetsPath,
	}

	Templates, err = newStaticTemplates(cfg)
	if err != nil {
		panic(err)
	}

	Styles, err = newStaticStyles(cfg)
	if err != nil {
		panic(err)
	}

	Icons, err = newStaticIcons(cfg)
	if err != nil {
		panic(err)
	}

	Javascript, err = newStaticJavascript(cfg)
	if err != nil {
		panic(err)
	}
}
