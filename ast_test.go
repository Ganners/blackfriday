package blackfriday

import (
	"reflect"
	"strings"
	"testing"
)

func launchAstTest(input []byte) []*ElementStringified {

	renderer := ASTRenderer(HtmlRenderer(commonHtmlFlags, "", ""))
	MarkdownOptions(
		input,
		renderer,
		Options{Extensions: commonExtensions})

	return renderer.GetTree()
}

func TestHeaderParagraph(t *testing.T) {

	input := []byte(strings.Join([]string{
		"# Header 1",
		"Some paragraph of text which spans across",
		"multiple lines"}, "\n"))

	output := launchAstTest(input)

	fixture := []*ElementStringified{
		{
			Name:     "h1",
			Rendered: "<h1>Header 1</h1>\n",
		},
		{
			Name:     "paragraph",
			Rendered: "<p>Some paragraph of text which spans across\nmultiple lines</p>\n",
		},
	}

	if !reflect.DeepEqual(output, fixture) {
		t.Errorf("Output of header did not match fixture")
	}
}
