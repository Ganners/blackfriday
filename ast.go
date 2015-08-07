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
	"fmt"
)

// Stores the syntax for an element
type Element struct {
	Name     string
	Rendered *bytes.Buffer
}

type ElementStringified struct {
	Name     string
	Rendered string
}

type AST struct {
	flags int

	// The sub renderer
	renderer Renderer

	// We'll populate this slice of elements as we go
	Tree []*Element

	tempRendered *bytes.Buffer
}

func ASTRenderer(subRenderer Renderer) *AST {

	// Create new renderer
	return &AST{
		renderer: subRenderer,
		Tree:     make([]*Element, 0),
	}
}

func (ast *AST) GetTree() []*ElementStringified {

	var elements = make([]*ElementStringified, len(ast.Tree))

	for i := 0; i < len(ast.Tree); i++ {

		elements[i] = &ElementStringified{
			Name:     ast.Tree[i].Name,
			Rendered: ast.Tree[i].Rendered.String(),
		}
	}

	return elements
}

func (ast *AST) BlockCode(out *bytes.Buffer, text []byte, lang string) {

	var rendered bytes.Buffer
	ast.renderer.BlockCode(&rendered, text, lang)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "code",
		Rendered: &rendered,
	})
}

func (ast *AST) BlockQuote(out *bytes.Buffer, text []byte) {

	var rendered bytes.Buffer
	ast.renderer.BlockQuote(&rendered, text)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "blockquote",
		Rendered: &rendered,
	})
}

func (ast *AST) BlockHtml(out *bytes.Buffer, text []byte) {

	var rendered bytes.Buffer
	ast.renderer.BlockHtml(&rendered, text)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "html",
		Rendered: &rendered,
	})
	ast.tempRendered = &rendered
}

func (ast *AST) Header(out *bytes.Buffer, text func() bool, level int, id string) {

	var rendered bytes.Buffer
	ast.renderer.Header(&rendered, text, level, id)

	ast.Tree = append(ast.Tree, &Element{
		Name:     fmt.Sprintf("h%d", level),
		Rendered: &rendered,
	})
}

func (ast *AST) HRule(out *bytes.Buffer) {

	var rendered bytes.Buffer
	ast.renderer.HRule(&rendered)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "hrule",
		Rendered: &rendered,
	})
}

func (ast *AST) List(out *bytes.Buffer, text func() bool, flags int) {

	var rendered bytes.Buffer
	ast.renderer.List(&rendered, text, flags)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "list",
		Rendered: &rendered,
	})
}

func (ast *AST) ListItem(out *bytes.Buffer, text []byte, flags int) {

	var rendered bytes.Buffer
	ast.renderer.ListItem(&rendered, text, flags)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "listitem",
		Rendered: &rendered,
	})
}

func (ast *AST) Paragraph(out *bytes.Buffer, text func() bool) {

	var rendered bytes.Buffer
	ast.renderer.Paragraph(&rendered, text)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "paragraph",
		Rendered: &rendered,
	})
}

func (ast *AST) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {

	var rendered bytes.Buffer
	ast.renderer.Table(&rendered, header, body, columnData)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "table",
		Rendered: &rendered,
	})
}

func (ast *AST) TableRow(out *bytes.Buffer, text []byte) {

	var rendered bytes.Buffer
	ast.renderer.TableRow(&rendered, text)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "tablerow",
		Rendered: &rendered,
	})
}

func (ast *AST) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {

	var rendered bytes.Buffer
	ast.renderer.TableHeaderCell(&rendered, text, flags)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "th",
		Rendered: &rendered,
	})
}

func (ast *AST) TableCell(out *bytes.Buffer, text []byte, flags int) {

	var rendered bytes.Buffer
	ast.renderer.TableCell(&rendered, text, flags)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "td",
		Rendered: &rendered,
	})
}

func (ast *AST) Footnotes(out *bytes.Buffer, text func() bool) {

	var rendered bytes.Buffer
	ast.renderer.Footnotes(&rendered, text)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "footnotes",
		Rendered: &rendered,
	})
}

func (ast *AST) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {

	var rendered bytes.Buffer
	ast.renderer.FootnoteItem(&rendered, name, text, flags)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "footnoteitem",
		Rendered: &rendered,
	})
}

func (ast *AST) TitleBlock(out *bytes.Buffer, text []byte) {

	var rendered bytes.Buffer
	ast.renderer.TitleBlock(&rendered, text)

	ast.Tree = append(ast.Tree, &Element{
		Name:     "tite",
		Rendered: &rendered,
	})
}

// Span-level callbacks
func (ast *AST) AutoLink(out *bytes.Buffer, link []byte, kind int) {

	ast.renderer.AutoLink(
		&ast.Tree[len(ast.Tree)-1].Rendered, link, kind)
}

func (ast *AST) CodeSpan(out *bytes.Buffer, text []byte) {

	ast.renderer.CodeSpan(
		&ast.Tree[len(ast.Tree)-1].Rendered, text)
}

func (ast *AST) DoubleEmphasis(out *bytes.Buffer, text []byte) {

	ast.renderer.DoubleEmphasis(
		&ast.Tree[len(ast.Tree)-1].Rendered, text)
}

func (ast *AST) Emphasis(out *bytes.Buffer, text []byte) {

	ast.renderer.Emphasis(
		&ast.Tree[len(ast.Tree)-1].Rendered, text)
}

func (ast *AST) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {

	ast.renderer.Image(
		&ast.Tree[len(ast.Tree)-1].Rendered, link, title, alt)
}

func (ast *AST) LineBreak(out *bytes.Buffer) {

	ast.renderer.LineBreak(
		&ast.Tree[len(ast.Tree)-1].Rendered)
}

func (ast *AST) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {

	ast.renderer.Link(
		&ast.Tree[len(ast.Tree)-1].Rendered, link, title, content)
}

func (ast *AST) RawHtmlTag(out *bytes.Buffer, tag []byte) {

	ast.renderer.RawHtmlTag(
		&ast.Tree[len(ast.Tree)-1].Rendered, tag)
}

func (ast *AST) TripleEmphasis(out *bytes.Buffer, text []byte) {

	ast.renderer.TripleEmphasis(
		&ast.Tree[len(ast.Tree)-1].Rendered, text)
}

func (ast *AST) StrikeThrough(out *bytes.Buffer, text []byte) {

	ast.renderer.StrikeThrough(
		&ast.Tree[len(ast.Tree)-1].Rendered, text)
}

func (ast *AST) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {

	ast.renderer.FootnoteRef(
		&ast.Tree[len(ast.Tree)-1].Rendered, ref, id)
}

// Low-level callbacks
func (ast *AST) Entity(out *bytes.Buffer, entity []byte) {

	ast.renderer.Entity(
		&ast.Tree[len(ast.Tree)-1].Rendered, entity)
}

func (ast *AST) NormalText(out *bytes.Buffer, text []byte) {

	ast.renderer.NormalText(
		&ast.tempRendered, text)
}

// Header and footer
func (ast *AST) DocumentHeader(out *bytes.Buffer) {

}

func (ast *AST) DocumentFooter(out *bytes.Buffer) {

}

func (ast *AST) GetFlags() int {

	return 0
}
