package crypto

const (
	PublicKeySuffix = "pub"
	// The suffix on the end of a public sig file
	SecretSigKeySuffix = "sig.key"
	// The suffix on the end of a secret box file
	SecretBoxKeySuffix = "box.key"
	// The suffix on the end of a secret symmetric key file
	SecretSymKeySuffix = "sym.key"
	// The hashing function we're using during sign/verify
	// See also: https://download.libsodium.org/doc/hashing/generic_hashing.html
	SigHashType = "BLAKE2b"
	// This environment variable allows you to override the fs::CACHE_KEY_PATH
	// at runtime. This is useful for testing.
	CacheKeyPathEnvVar        = "HAB_CACHE_KEY_PATH"
	HartFormatVersion         = "HART-1"
	BoxFormatVersion          = "BOX-1"
	AnonymousBoxFormatVersion = "ANONYMOUS-BOX-1"
	// Create secret key files with these permissions
	PublicKeyPermissions = 0400
	SecretKeyPermissions = 0400

	PublicSigKeyVersion  string = "SIG-PUB-1"
	PrivateSigKeyVersion string = "SIG-SEC-1"
	PublicBoxKeyVersion  string = "BOX-PUB-1"
	PrivateBoxKeyVersion string = "BOX-SEC-1"
	PrivateSymKeyVersion string = "SYM-SEC-1"
)
