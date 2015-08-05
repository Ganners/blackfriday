//
// A simple AST renderer, built to solve a basic problem and so does not have
// full AST output.
//
// Current implementation does not branch into span element, instead, it will
// just flatten them into HTML and append to the parent. It isn't a feature
// I require just yet.
//

package blackfriday

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Stores the syntax for an element
type Element struct {
	Name     string
	Rendered string
}

type AST struct {
	flags int

	// The sub renderer
	renderer Renderer

	// We'll populate this slice of elements as we go
	Tree []*Element
}

func ASTRenderer(subRenderer Renderer) *AST {

	// Create new renderer
	return &AST{
		renderer: subRenderer,
		Tree:     make([]*Element, 0),
	}
}

func (ast *AST) BlockCode(out *bytes.Buffer, text []byte, lang string) {

	var rendered *bytes.Buffer
	ast.renderer.BlockCode(out, text, lang)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "code",
		Rendered: rendered.String(),
	})

	binary.Write(out, binary.BigEndian, ast.Tree)
}

func (ast *AST) BlockQuote(out *bytes.Buffer, text []byte) {

}

func (ast *AST) BlockHtml(out *bytes.Buffer, text []byte) {

}

func (ast *AST) Header(out *bytes.Buffer, text func() bool, level int, id string) {

	var rendered bytes.Buffer
	ast.renderer.Header(&rendered, text, level, id)

	ast.Tree = append(ast.Tree, &Element{
		Name:     fmt.Sprintf("h%d", level),
		Rendered: rendered.String(),
	})
}

func (ast *AST) HRule(out *bytes.Buffer) {

}

func (ast *AST) List(out *bytes.Buffer, text func() bool, flags int) {

}

func (ast *AST) ListItem(out *bytes.Buffer, text []byte, flags int) {

}

func (ast *AST) Paragraph(out *bytes.Buffer, text func() bool) {

}

func (ast *AST) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {

}

func (ast *AST) TableRow(out *bytes.Buffer, text []byte) {

}

func (ast *AST) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {

}

func (ast *AST) TableCell(out *bytes.Buffer, text []byte, flags int) {

}

func (ast *AST) Footnotes(out *bytes.Buffer, text func() bool) {

}

func (ast *AST) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {

}

func (ast *AST) TitleBlock(out *bytes.Buffer, text []byte) {

}

// Span-level callbacks
func (ast *AST) AutoLink(out *bytes.Buffer, link []byte, kind int) {

}

func (ast *AST) CodeSpan(out *bytes.Buffer, text []byte) {

}

func (ast *AST) DoubleEmphasis(out *bytes.Buffer, text []byte) {

}

func (ast *AST) Emphasis(out *bytes.Buffer, text []byte) {

}

func (ast *AST) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {

}

func (ast *AST) LineBreak(out *bytes.Buffer) {

}

func (ast *AST) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {

}

func (ast *AST) RawHtmlTag(out *bytes.Buffer, tag []byte) {

}

func (ast *AST) TripleEmphasis(out *bytes.Buffer, text []byte) {

}

func (ast *AST) StrikeThrough(out *bytes.Buffer, text []byte) {

}

func (ast *AST) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {

}

// Low-level callbacks
func (ast *AST) Entity(out *bytes.Buffer, entity []byte) {

}

func (ast *AST) NormalText(out *bytes.Buffer, text []byte) {

}

// Header and footer
func (ast *AST) DocumentHeader(out *bytes.Buffer) {

}

func (ast *AST) DocumentFooter(out *bytes.Buffer) {

}

func (ast *AST) GetFlags() int {

	return 0
}
