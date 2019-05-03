package fs

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

const CacheKeyPath = "hab/cache/keys"
const CacheArtifactPath = "hab/cache/artifacts"

// CacheKeyPath returns the defaultCacheKeyPath given the home directory
func GetCacheKeyPath() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(home, CacheKeyPath) // Preserve legacy behavior
}
