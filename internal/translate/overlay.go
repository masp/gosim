package translate

import (
	_ "embed"
	"fmt"
	"go/version"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

//go:embed overlays/internal_sync_hashtriemap.go.txt
var internalSyncHashTrieMap []byte

// stdlibOverlays lists standard library files that are replaced before
// translation, keyed by their path relative to GOROOT/src.
var stdlibOverlays = map[string][]byte{
	"internal/sync/hashtriemap.go": internalSyncHashTrieMap,
}

var (
	stdlibOverlayOnce   sync.Once
	stdlibOverlayResult map[string][]byte
	stdlibOverlayErr    error
)

var goToolchainVersionOnce = sync.OnceValue(func() string {
	out, err := exec.Command("go", "env", "GOVERSION").Output()
	if err != nil {
		log.Fatalf("finding go version: %s", err)
	}
	// strip any suffix such as " X:nocoverageredesign" or "-devel ..."
	v := version.Lang(strings.TrimSpace(string(out)))
	if !version.IsValid(v) {
		log.Fatalf("unexpected go version %q", out)
	}
	return v
})

// goToolchainVersion returns the language version (eg. "go1.27") of the go
// command used to build translated code.
func goToolchainVersion() string {
	return goToolchainVersionOnce()
}

// stdlibOverlay returns a go/packages overlay that replaces the files in
// stdlibOverlays in the GOROOT used by the go command.
func stdlibOverlay(env []string) (map[string][]byte, error) {
	stdlibOverlayOnce.Do(func() {
		cmd := exec.Command("go", "env", "GOROOT")
		cmd.Env = env
		out, err := cmd.Output()
		if err != nil {
			stdlibOverlayErr = fmt.Errorf("finding GOROOT: %w", err)
			return
		}
		goroot := strings.TrimSpace(string(out))

		stdlibOverlayResult = make(map[string][]byte)
		for path, contents := range stdlibOverlays {
			stdlibOverlayResult[filepath.Join(goroot, "src", filepath.FromSlash(path))] = contents
		}
	})
	return stdlibOverlayResult, stdlibOverlayErr
}
