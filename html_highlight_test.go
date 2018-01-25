package static

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestSyntaxHighlight(t *testing.T) {
	r := ioutil.NopCloser(strings.NewReader(`
<pre><code>
package main

func main() {
	fmt.Println("foo")
}
</code></pre>

<pre><code class="language-yaml">
foo: bar
list:
- 1
- 2
</code></pre>

<pre><code>
this is not even a lang
</code></pre>
	`))

	r = SyntaxHighlight(r)
	got, _ := ioutil.ReadAll(r)
	expect, _ := ioutil.ReadFile("testdata/code.html")
	if string(got) != string(expect) {
		t.Errorf("expected %s but got %s", string(expect), string(got))
		// ioutil.WriteFile("testdata/code.html", got, 0644)
	}
}
