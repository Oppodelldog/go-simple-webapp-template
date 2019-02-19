//go:generate statics -i=assets/js -o=assets/javascript.go -pkg=assets -group=Javascript -ignore=.gitignore -prefix=assets

package assets

import (
	"os"

	"github.com/go-playground/statics/static"
)

// newStaticJavascript initializes a new *static.Files instance for use
func newStaticJavascript(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{
		Path:    "/js",
		Name:    "js",
		Size:    4096,
		Mode:    os.FileMode(2147484157),
		ModTime: 1550601214,
		IsDir:   true,
		Compressed: `
`,
		Files: []*static.DirFile{{
			Path:    "/js/index.js",
			Name:    "index.js",
			Size:    27,
			Mode:    os.FileMode(436),
			ModTime: 1550601214,
			IsDir:   false,
			Compressed: `
H4sIAAAAAAAA/0rOzyvOz0nVy8lP11DySM3JyVcoyUgtSlXStAYEAAD//wpItUkbAAAA
`,
			Files: []*static.DirFile{},
		},
		},
	})
}
