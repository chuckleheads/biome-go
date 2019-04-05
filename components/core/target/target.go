package target

import (
	"strings"

	"github.com/biome-sh/biome-go/components/core/os"
	"github.com/pkg/errors"
)

type PackageTarget struct {
	Platform     os.Platform
	Architecture os.Architecture
}

func (pt PackageTarget) FromString(target string) (PackageTarget, error) {
	parts := strings.Split(target, "-")
	if len(parts) < 2 || len(parts) > 4 {
		return PackageTarget{}, errors.Errorf("invalid package target '%s'", target)
	}

	for i, part := range parts {
		switch i {
		case 0:
			pt.Platform = os.PlatformFromString(part)
		case 1:
			pt.Architecture = os.ArchitectureFromString(part)
		}
	}

	return pt, nil
}
