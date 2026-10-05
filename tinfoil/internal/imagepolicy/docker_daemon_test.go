package imagepolicy_test

import (
	"encoding/json"
	"os"
	"testing"
)

func TestShippingDockerDaemonJSONRegistersRunscByPath(t *testing.T) {
	raw, err := os.ReadFile("../../../image/rootfs/etc/docker/daemon.json")
	if err != nil {
		t.Fatal(err)
	}
	var daemon map[string]any
	if err := json.Unmarshal(raw, &daemon); err != nil {
		t.Fatal(err)
	}
	runtimes, ok := daemon["runtimes"].(map[string]any)
	if !ok {
		t.Fatal("daemon.json is missing runtimes")
	}
	if _, ok := runtimes["nvidia"]; !ok {
		t.Fatal("daemon.json is missing the nvidia runtime")
	}
	runsc, ok := runtimes["runsc"].(map[string]any)
	if !ok {
		t.Fatal("daemon.json is missing the runsc runtime")
	}
	if path, _ := runsc["path"].(string); path != "/usr/bin/runsc" {
		t.Fatalf("runsc path = %v, want /usr/bin/runsc", runsc["path"])
	}
	if _, ok := runsc["runtimeArgs"]; ok {
		t.Fatal("runsc must not set runtimeArgs; dockerd would wrap it onto the noexec ramdisk")
	}
	if args, ok := runsc["args"].([]any); ok && len(args) > 0 {
		t.Fatal("runsc must not set args; dockerd would wrap it onto the noexec ramdisk")
	}
	if _, err := os.Stat("../../../image/debug-rootfs/etc/docker/daemon.json"); err == nil {
		t.Fatal("debug overlay must not fork daemon.json; runsc belongs in the shipping file")
	}
}
