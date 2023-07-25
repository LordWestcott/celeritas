//Special file to run some func before we run tests.

package render

import (
	"os"
	"testing"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./testdata/views"),
	jet.InDevelopmentMode(),
)

var testRenderer = Render{
	Renderer: "",
	RootPath: "",
	JetViews: views,
}

func TestMain(m *testing.M) {
	//Exit, but before you exit, run our tests.
	os.Exit(m.Run())
}
