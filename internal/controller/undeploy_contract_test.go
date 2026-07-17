package controller

import (
	"os"
	"strings"
	"testing"
)

func TestUndeployContractKeepsCRDDeletionInUninstall(t *testing.T) {
	overlay, err := os.ReadFile("../../config/undeploy/kustomization.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(overlay), "../crd") {
		t.Fatal("undeploy overlay must not include CRDs; make uninstall owns CRD deletion")
	}
	for _, required := range []string{"../rbac", "../manager", "../default/metrics_service.yaml"} {
		if !strings.Contains(string(overlay), required) {
			t.Fatalf("undeploy overlay missing %q", required)
		}
	}
	makefile, err := os.ReadFile("../../Makefile")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(makefile), `"$(KUSTOMIZE)" build config/undeploy`) {
		t.Fatal("make undeploy does not use the CRD-free overlay")
	}
}
