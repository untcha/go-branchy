package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/untcha/go-branchy/internal/appmeta"
)

func TestVersionCommandPrintsAppMetadata(t *testing.T) {
	oldVersion := appmeta.Version
	oldCommit := appmeta.Commit
	oldBuildDate := appmeta.BuildDate
	t.Cleanup(func() {
		appmeta.Version = oldVersion
		appmeta.Commit = oldCommit
		appmeta.BuildDate = oldBuildDate
	})

	appmeta.Version = "1.2.3"
	appmeta.Commit = "abc1234"
	appmeta.BuildDate = "2026-05-06T08:30:00Z"

	cmd := newVersionCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}

	got := out.String()
	for _, want := range []string{
		"Version: 1.2.3",
		"Commit: abc1234",
		"Build date: 2026-05-06T08:30:00Z",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("output %q does not contain %q", got, want)
		}
	}
}
