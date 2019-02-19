//go:generate statics -i=assets/templates -o=assets/templates.go -pkg=assets -group=Templates -ignore=.gitignore -prefix=assets

package assets

import (
	"os"

	"github.com/go-playground/statics/static"
)

// newStaticTemplates initializes a new *static.Files instance for use
func newStaticTemplates(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{
		Path:    "/templates",
		Name:    "templates",
		Size:    4096,
		Mode:    os.FileMode(2147484157),
		ModTime: 1550603185,
		IsDir:   true,
		Compressed: `
`,
		Files: []*static.DirFile{{
			Path:    "/templates/index.html",
			Name:    "index.html",
			Size:    146,
			Mode:    os.FileMode(509),
			ModTime: 1550603185,
			IsDir:   false,
			Compressed: `
H4sIAAAAAAAA/1TOQa7CMAwE0H1O4Z8D/Kh747u4xMRITm2JAKqq3h1BxYLVLN5oNPhX/TzWENDRjRK+
A4yXdsqZ0rYN6WE8BLIK17zvCWev6y+ZN/+QTtQcbtceJvCUmSPgW8OiEyUMUn4IXO7LP5aghOXYw3I8
eAUAAP//w1og/ZIAAAA=
`,
			Files: []*static.DirFile{},
		},
			{
				Path:    "/templates/partials",
				Name:    "partials",
				Size:    4096,
				Mode:    os.FileMode(2147484157),
				ModTime: 1550604049,
				IsDir:   true,
				Compressed: `
`,
				Files: []*static.DirFile{{
					Path:    "/templates/partials/head.html",
					Name:    "head.html",
					Size:    606,
					Mode:    os.FileMode(509),
					ModTime: 1550601552,
					IsDir:   false,
					Compressed: `
H4sIAAAAAAAA/5SSwXKzMAyE7/9TeHT+HU9vPeC8izAiKDG2iwUhZXj3jiGdSdJeetOy4tvVjJeloZYD
KegIG1jXf0opVRVx3MZN9iSoXIdDJrEwSqvf4dXuRJKmj5EnC7MeUbvYJxSuPYFyMQgFscBkqTnR49/C
4ul4ijpznzzpK9WYkhbqkxeqzO6/pAXsycLEdE1xkIeAKzfS2YYmdqQ38V9xYGH0Ojv0ZN9+VN9hmEp6
H2veS+jSwmHC5wtulMH8gZAFZcy6xkFnuT2hao/u8gTzHC5qIG9h280dkYCSWyILQrMYlzOobqDWQpnN
/TT+ROEYDsX+Fccuhm8Q93giM+v92x1WRDYtTmU4sIuPnOwGTqLy4CyYc0ltaD6cMxwrs3v7cmX2h7Ms
FJp1Xb8CAAD//ymwLFFeAgAA
`,
					Files: []*static.DirFile{},
				},
					{
						Path:    "/templates/partials/logo.html",
						Name:    "logo.html",
						Size:    99,
						Mode:    os.FileMode(509),
						ModTime: 1550604049,
						IsDir:   false,
						Compressed: `
H4sIAAAAAAAA/yTKQQrEMAiF4X1PIR5gvEDbuxRNgpCJMEI24t2HJrvH9/4IKVVHAezWDDMPAIBTdAL3
x/3afi9el34b+I8vJGUbTvWZ7/goG9LuThKd9xFRhmT+AwAA//84sRGMYwAAAA==
`,
						Files: []*static.DirFile{},
					},
				},
			},
		},
	})
}
