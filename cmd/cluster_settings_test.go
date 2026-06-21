package cmd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/altinity/acmctl/pkg/api"
)

const settingsFixture = `{"data":[
  {"id":"67045","name":"config.d/sql-play.xml","value":"<clickhouse/>","isFile":true,"id_cluster":"9536"},
  {"id":"65152","name":"display_name","value":"antalya","isFile":false,"id_cluster":"9536"}
]}`

// serve wires apiClient at an httptest server for the duration of a test.
func serve(t *testing.T, h http.HandlerFunc) {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	apiClient = api.NewClient(srv.URL, "tok")
}

func TestLooksLikeFile(t *testing.T) {
	for name, want := range map[string]bool{
		"config.d/sql-play.xml": true,
		"users.d/readonly.xml":  true,
		"foo.sql":               true,
		"bar.yaml":              true,
		"https_port":            false,
		"display_name":          false,
	} {
		if got := looksLikeFile(name); got != want {
			t.Errorf("looksLikeFile(%q)=%v, want %v", name, got, want)
		}
	}
}

func TestFindSetting(t *testing.T) {
	serve(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/cluster/9536/settings" {
			io.WriteString(w, settingsFixture)
			return
		}
		t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
	})

	if s, err := findSetting("9536", "config.d/sql-play.xml"); err != nil || s == nil || s.ID != "67045" || !s.IsFile {
		t.Fatalf("by name: %+v err=%v", s, err)
	}
	if s, _ := findSetting("9536", "65152"); s == nil || s.Name != "display_name" {
		t.Fatalf("by id: %+v", s)
	}
	if s, _ := findSetting("9536", "missing"); s != nil {
		t.Fatalf("absent should be nil, got %+v", s)
	}
}

func TestSetUpdatesExistingByID(t *testing.T) {
	var gotPath, gotBody string
	serve(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet:
			io.WriteString(w, settingsFixture)
		case r.Method == http.MethodPost:
			b, _ := io.ReadAll(r.Body)
			gotPath, gotBody = r.URL.Path, string(b)
			io.WriteString(w, `{"data":{"ok":true}}`)
		default:
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
	})

	c := clusterSettingsSetCmd
	_ = c.Flags().Set("value", "NEWVAL")
	if err := c.RunE(c, []string{"9536", "config.d/sql-play.xml"}); err != nil {
		t.Fatal(err)
	}
	if gotPath != "/cluster-setting/67045" {
		t.Errorf("update should POST by id, got path %q", gotPath)
	}
	if !strings.Contains(gotBody, "NEWVAL") {
		t.Errorf("body missing new value: %s", gotBody)
	}
}

func TestSetCreatesNewWithInferredIsFile(t *testing.T) {
	var gotPath, gotBody string
	serve(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet:
			io.WriteString(w, settingsFixture)
		case r.Method == http.MethodPost:
			b, _ := io.ReadAll(r.Body)
			gotPath, gotBody = r.URL.Path, string(b)
			io.WriteString(w, `{"data":{"ok":true}}`)
		default:
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
	})

	c := clusterSettingsSetCmd
	_ = c.Flags().Set("value", "<new/>")
	if err := c.RunE(c, []string{"9536", "config.d/brand-new.xml"}); err != nil {
		t.Fatal(err)
	}
	if gotPath != "/cluster/9536/settings" {
		t.Errorf("create should POST to /cluster/<id>/settings, got %q", gotPath)
	}
	if !strings.Contains(gotBody, `"isFile":true`) {
		t.Errorf("isFile should be inferred true for a config.d/*.xml name: %s", gotBody)
	}
}

func TestRmResolvesNameToID(t *testing.T) {
	var delPath string
	serve(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			io.WriteString(w, settingsFixture)
		case http.MethodDelete:
			delPath = r.URL.Path
		default:
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
	})

	c := clusterSettingsRmCmd
	if err := c.RunE(c, []string{"9536", "config.d/sql-play.xml"}); err != nil {
		t.Fatal(err)
	}
	if delPath != "/cluster-setting/67045" {
		t.Errorf("rm by name should DELETE id 67045, got %q", delPath)
	}
}
