package archive

type Metafile int

// TED: I don't know if all of this needs to be public

const (
	CFlags Metafile = iota
	Deps
	TDeps
	Exposes
	Ident
	LdRunPath
	LdFlags
	SvcUser
	Services
	ResolvedServices
	Manifest
	Path
	Target
	Type
	Config
)

func (file Metafile) String() string {
	// declare an array of strings
	// ... operator counts how many
	// items in the array (7)
	names := [...]string{
		"CFLAGS",
		"DEPS",
		"TDEPS",
		"EXPOSES",
		"IDENT",
		"LD_RUN_PATH",
		"LD_FLAGS",
		"SVC_USER",
		"SERVICES",
		"RESOLVED_SERVICES",
		"MANIFEST",
		"PATH",
		"TARGET",
		"TYPE",
		"default.toml",
	}
	if file < CFlags || file > Type {
		return "Unknown"
	}
	return names[file]
}
