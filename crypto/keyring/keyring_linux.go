//go:build linux
// +build linux

package keyring

import (
	"fmt"
	"io"

	"github.com/99designs/keyring"
	"github.com/cosmos/cosmos-sdk/codec"
)

// Linux-only backend options.
const BackendKeyctl = "keyctl"

func newKeyctlBackendConfig(appName, _ string, _ io.Reader) keyring.Config {
	return keyring.Config{
		AllowedBackends: []keyring.BackendType{keyring.KeyCtlBackend},
		ServiceName:     appName,
		KeyCtlScope:     "user",
		KeyCtlPerm:      0x3f3f0000,
	}
}

// New creates a new instance of a keyring.
// Keyring options can be applied when generating the new instance.
// Available backends are "os", "file", "kwallet", "memory", "pass", "test", "keyctl".
func New(
	appName, backend, rootDir string, userInput io.Reader, cdc codec.Codec, opts ...Option,
) (Keyring, error) {

	if backend != BackendKeyctl {
		return newKeyringGeneric(appName, backend, rootDir, userInput, cdc, opts...)
	}

	db, err := keyring.Open(newKeyctlBackendConfig(appName, "", userInput))
	if err != nil {
		return nil, fmt.Errorf("couldn't open keyring for %q: %w", appName, err)
	}

	return newKeystore(db, cdc, backend, opts...), nil
}
