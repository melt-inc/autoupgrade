package autoupgrade

import (
	"testing"
)

func Test_fullPath(t *testing.T) {
	expected := "github.com/melt-inc/autoupgrade/foo/bar@latest"
	actual := fullPath("github.com/melt-inc/autoupgrade", "foo/bar", "latest")
	if actual != expected {
		t.Errorf("\n--- '%s'\n+++ '%s'", expected, actual)
	}
}
