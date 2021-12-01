package validation

import "testing"

var checkURLCases = []struct {
	name      string
	url       string
	expectErr bool
}{
	{
		name:      "valid http://",
		url:       "http://url.io",
		expectErr: false,
	},
	{
		name:      "valid https://",
		url:       "https://url.io",
		expectErr: false,
	},
	{
		name:      "invalid protocol",
		url:       "htp://url.io",
		expectErr: true,
	},
	{
		name:      "disallowed protocol",
		url:       "irc://url.io",
		expectErr: true,
	},
	{
		name:      "empty url",
		url:       "",
		expectErr: true,
	},
}

func checkURLTest(t *testing.T) {
	for _, tt := range checkURLCases {
		out, err := checkURL(tt.url)
		if err != nil && !tt.expectErr {
			t.Errorf("%s :: %s", tt.name, err.Error())
		}

		if out != tt.url && !tt.expectErr {
			t.Errorf("URL mangled. Got %s - expected %s", out, tt.url)
		}

		if out != "" && err != nil && tt.expectErr {
			t.Errorf("Didn't fail when expected")
		}
	}
}
