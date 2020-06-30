package apigee

import "testing"

func TestApigeeEdge_addOptions(t *testing.T) {
	type TestOptions struct {
		Name string `url:"name"`
	}

	url, err := addOptions("/foo/bar", TestOptions{Name: "foo bar"})
	if err != nil {
		t.Errorf("addOptions: got error %s", err.Error())
	}

	expected := "/foo/bar?name=foo+bar"
	if url != expected {
		t.Errorf("addOptions: did not format url correctly, expected %s got %s", expected, url)
	}
}
