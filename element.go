package mdfunc

import (
	"fmt"
	"io"
	"strings"
)

// Element

type Element struct {
	// Raw is basically a bypass, allowing for element creation from raw MD
	Raw    string
	Render string
}

func (el Element) Output() (string, bool) {
	// Raw trumps everything and is just returned as is
	if el.Raw != "" {
		return fmt.Sprintf("%s\n", el.Raw), true
	}

	if el.Render == "" {
		return "", false
	}

	return el.Render, true
}

func (el Element) String() string {
	s, _ := el.Output()
	return s
}

func (el Element) Bytes() []byte {
	renderedElement, _ := el.Output()
	return []byte(renderedElement)
}

func (el Element) WriteDoc(w io.Writer) error {
	_, err := w.Write(el.Bytes())
	return err
}

func (el Element) RenderIf(doRender bool) Element {
	if doRender {
		return el
	}

	return Element{}
}

func (el Element) RenderIfNotBlank() Element {
	if el.Raw != "" || el.Render != "" {
		return el
	}

	return Element{}
}

// Elements

type Elements []Element

func (els Elements) Output() string {
	var renderedEls []string

	for _, el := range els {
		if renderedElement, doRender := el.Output(); doRender {
			renderedEls = append(
				renderedEls,
				renderedElement,
			)
		}
	}

	return strings.TrimSpace(strings.Join(renderedEls, ""))
}

func (els Elements) String() string {
	return els.Output()
}

func (els Elements) Bytes() []byte {
	renderedElement := els.Output()
	return []byte(renderedElement)
}

func (els Elements) WriteDoc(w io.Writer) error {
	_, err := w.Write(els.Bytes())
	return err
}

// Tags

func Doc(els ...Element) Elements {
	return els
}

func H1(s string) Element {
	return Element{
		Render: fmt.Sprintf("# %s\n", s),
	}
}

func H2(s string) Element {
	return Element{
		Render: fmt.Sprintf("## %s\n", s),
	}
}

func H3(s string) Element {
	return Element{
		Render: fmt.Sprintf("### %s\n", s),
	}
}

func H4(s string) Element {
	return Element{
		Render: fmt.Sprintf("#### %s\n", s),
	}
}

func H5(s string) Element {
	return Element{
		Render: fmt.Sprintf("##### %s\n", s),
	}
}

func H6(s string) Element {
	return Element{
		Render: fmt.Sprintf("###### %s\n", s),
	}
}

func Bold(s string) Element {
	return Element{
		Render: fmt.Sprintf("**%s** ", s),
	}
}

func Italic(s string) Element {
	return Element{
		Render: fmt.Sprintf("*%s* ", s),
	}
}

func Text(s string) Element {
	return Element{
		Render: fmt.Sprintf("%s ", s),
	}
}

func P(els ...Element) Element {
	return Element{
		Render: fmt.Sprintf("\n%s\n", Elements(els).Output()),
	}
}

func Line(els ...Element) Element {
	return Element{
		Render: fmt.Sprintf("%s\n", Elements(els).Output()),
	}
}

func BlockQuote(s string) Element {
	return Element{
		Render: fmt.Sprintf("> %s\n", s),
	}
}

func Link(alt, dest string) Element {
	return Element{
		Render: fmt.Sprintf("[%s](%s) ", alt, dest),
	}
}

func Image(alt, dest string) Element {
	return Element{
		Render: fmt.Sprintf("![%s](%s) ", alt, dest),
	}
}

func Ul(els ...Element) Element {
	var out []string

	for _, el := range els {
		if output, ok := el.Output(); ok {
			li := fmt.Sprintf("- %s", output)
			out = append(out, strings.TrimSpace(li))
		}
	}

	render := strings.Join(out, "\n")

	return Element{
		Render: fmt.Sprintf("\n%s\n", render),
	}
}

func Ol(els ...Element) Element {
	var out []string

	for i, el := range els {
		if output, ok := el.Output(); ok {
			li := fmt.Sprintf("%v %s", i+1, output)
			out = append(out, strings.TrimSpace(li))
		}
	}

	render := strings.Join(out, "\n")

	return Element{
		Render: fmt.Sprintf("\n%s\n", render),
	}
}
