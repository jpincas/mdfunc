package mdfunc

import (
	"testing"
)

func TestOutput(t *testing.T) {
	doc := Doc(
		H1("Title"),
		H2("Subtitle"),
		Text("Just some text."),
		P(
			Text("This is a paragraph with some"),
			Italic("italic"),
			Text("and some"),
			Bold("bold"),
			Text("text."),
		),
		P(
			Text("Here is a list of links and images:"),
		),
		Ul(
			Text("This is just some text"),
			Link("Alt", "www.link.com"),
			Image("Alt", "www.link.com/image.jpeg"),
		),
		Ol(
			Text("Jon"),
			Text("Jessi"),
			Text("Allan"),
		),
		P(
			Text("this shouldn't render"),
		).RenderIf(false),
	)

	output := doc.Output()

	expected := `# Title
## Subtitle
Just some text.
This is a paragraph with some *italic* and some **bold** text.

Here is a list of links and images:

- This is just some text
- [Alt](www.link.com)
- ![Alt](www.link.com/image.jpeg)

1 Jon
2 Jessi
3 Allan`

	if output != expected {
		t.Errorf("Output was not as expected.\n\nOutput: \n\n%s\n\n Expected: \n\n%s\n", output, expected)
	}
}
