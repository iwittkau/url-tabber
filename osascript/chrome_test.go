package osascript_test

import (
	"testing"

	"github.com/iwittkau/url-tabber/osascript"
)

func Test_OpenNewChromeTab(t *testing.T) {
	err := osascript.OpenNewChromeTab("https://example.com")
	if err != nil {
		t.Error(err)
	}
}
