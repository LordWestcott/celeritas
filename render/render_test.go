package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "no error rendering non-existent template (go template)"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering non-existent template (jet template)"},
	{"invalid_render_engine", "foo", "home", true, "no error when using incorrect render engine"},
	{"invalid_render_engine", "", "home", true, "no error when using no render engine"},
}

// Test (*Render).Page
func TestRender_Page(t *testing.T) {

	for _, e := range pageData {
		//Make a new request
		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}

		//Make a response-writer recorder.
		w := httptest.NewRecorder()

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	if err = testRenderer.Page(w, r, "home", nil, nil); err != nil {
		t.Error("Error rendering GoTemplates page", err)
	}

	td := &TemplateData{
		ServerName: "Test",
	}

	if err = testRenderer.Page(w, r, "home", nil, td); err != nil {
		t.Error("Error rendering GoTemplates page", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata"

	if err = testRenderer.Page(w, r, "home", nil, nil); err != nil {
		t.Error("Error rendering GoTemplates page", err)
	}

	//With Template Data
	td := &TemplateData{
		ServerName: "Test",
	}

	if err = testRenderer.Page(w, r, "home", nil, td); err != nil {
		t.Error("Error rendering GoTemplates page", err)
	}
}
