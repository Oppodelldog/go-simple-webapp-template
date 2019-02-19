//go:generate statics -i=assets/icons -o=assets/icons.go -pkg=assets -group=Icons -ignore=.gitignore -prefix=assets

package assets

import (
	"os"

	"github.com/go-playground/statics/static"
)

// newStaticIcons initializes a new *static.Files instance for use
func newStaticIcons(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{
		Path:    "/icons",
		Name:    "icons",
		Size:    4096,
		Mode:    os.FileMode(2147484157),
		ModTime: 1550600088,
		IsDir:   true,
		Compressed: `
`,
		Files: []*static.DirFile{{
			Path:    "/icons/favicon.ico",
			Name:    "favicon.ico",
			Size:    4286,
			Mode:    os.FileMode(509),
			ModTime: 1550601531,
			IsDir:   false,
			Compressed: `
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
`,
			Files: []*static.DirFile{},
		},
		},
	})
}
