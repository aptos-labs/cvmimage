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
	if experimental, ok := daemon["experimental"].(bool); ok && experimental {
		t.Fatal("shipping daemon.json must not enable experimental")
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
		t.Fatal("shipping runsc must not set runtimeArgs; debug images overlay daemon.json")
	}
	if args, ok := runsc["args"].([]any); ok && len(args) > 0 {
		t.Fatal("shipping runsc must not set args; debug images overlay daemon.json")
	}
}

func TestDebugDockerDaemonJSONEnablesCheckpointPolicy(t *testing.T) {
	raw, err := os.ReadFile("../../../image/debug-rootfs/etc/docker/daemon.json")
	if err != nil {
		t.Fatal(err)
	}
	var daemon map[string]any
	if err := json.Unmarshal(raw, &daemon); err != nil {
		t.Fatal(err)
	}
	experimental, ok := daemon["experimental"].(bool)
	if !ok || !experimental {
		t.Fatal("debug daemon.json must set experimental: true for docker checkpoint")
	}
	runtimes, ok := daemon["runtimes"].(map[string]any)
	if !ok {
		t.Fatal("debug daemon.json is missing runtimes")
	}
	runsc, ok := runtimes["runsc"].(map[string]any)
	if !ok {
		t.Fatal("debug daemon.json is missing the runsc runtime")
	}
	args, ok := runsc["runtimeArgs"].([]any)
	if !ok {
		t.Fatal("debug runsc must set runtimeArgs")
	}
	want := []string{"--platform=systrap", "--overlay2=none"}
	if len(args) != len(want) {
		t.Fatalf("runtimeArgs = %v, want %v", args, want)
	}
	for i, arg := range want {
		if args[i] != arg {
			t.Fatalf("runtimeArgs[%d] = %v, want %v", i, args[i], arg)
		}
	}
}
