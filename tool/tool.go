package tool

import (
	"strings"

	"github.com/CCDirectLink/CCUpdaterCLI/internal/game"
	"github.com/CCDirectLink/CCUpdaterCLI/pkg"
)

//Available tools.
func Available(game game.Game) ([]pkg.Package, error) {
	return []pkg.Package{
		ccloader{game},
	}, nil
}

//Installed tools.
func Installed(game game.Game) ([]pkg.Package, error) {
	avail, err := Available(game)
	result := make([]pkg.Package, 0, len(avail))

	for _, tool := range avail {
		if tool.Installed() {
			result = append(result, tool)
		}
	}

	return result, err
}

//Get a tool by exact name.
func Get(game game.Game, name string) (pkg.Package, error) {
	switch strings.ToLower(name) {
	case "ccloader":
		return ccloader{game}, nil
	default:
		return nil, pkg.NewError(pkg.ModeUnknown, nil, pkg.ErrNotFound)
	}
}
