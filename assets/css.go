//go:generate statics -i=assets/css -o=assets/css.go -pkg=assets -group=Styles -ignore=.gitignore -prefix=assets

package assets

import (
	"os"

	"github.com/go-playground/statics/static"
)

// newStaticStyles initializes a new *static.Files instance for use
func newStaticStyles(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{
		Path:    "/css",
		Name:    "css",
		Size:    4096,
		Mode:    os.FileMode(2147484157),
		ModTime: 1550604031,
		IsDir:   true,
		Compressed: `
`,
		Files: []*static.DirFile{{
			Path:    "/css/initialization.css",
			Name:    "initialization.css",
			Size:    339,
			Mode:    os.FileMode(509),
			ModTime: 1550604031,
			IsDir:   false,
			Compressed: `
H4sIAAAAAAAA/2yOS27DMAxE9zoF14FUOP0tlNNIER0RVUhBplG7Re5eRE3bAPWOxMzDvKznAp8GACDK
4ib6ID55iNISNhdlOZiLMTsLOx9xlIbXK4yKbYMizthIOxIlrbfKObQTsYehB/PPXg0pdWw49L/QpG7S
taDTtaIHFsaOPByFFVkt5L2F/GghP1nIzxbyi4X8+n/n73UFR/WwH+pym98s38tcjCGus34X3TvGN1IX
asXQAh9/za6pzFqI72XNVwAAAP//cI7YTlMBAAA=
`,
			Files: []*static.DirFile{},
		},
		},
	})
}
