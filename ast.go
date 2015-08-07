//
// A simple AST renderer, built to solve a basic problem and so does not have
// full AST output.
//
// Current implementation does not branch into span element, instead, it will
// just flatten them into HTML and append to the parent. It isn't a feature
// I require just yet.
//
// My use is I want to be able to enumerate my top level elements in order to
// build something which can populate different aspects of a web page.
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

// The type of element we output, has been kindly converted to as tring on the
// rendered part for us
type ElementStringified struct {
	Name     string
	Rendered string
}

// AST implements the Renderer interface
type AST struct {
	flags int

	// The sub renderer
	renderer Renderer

	// Populate the tree with AST (not exactly a tree yet, just a list)
	tree []*Element

	// Need to store a pointer to what we are writing to in the tree for easy
	// access to have something to put span elements in to
	tempRendered *bytes.Buffer
}

// Construct a new renderer, this accepts a sub renderer which will populate as
// it goes through.
func ASTRenderer(subRenderer Renderer) *AST {

	// Create new renderer
	return &AST{
		renderer: subRenderer,
		tree:     make([]*Element, 0),
	}
}

func (ast *AST) GetTree() []*ElementStringified {

	var elements = make([]*ElementStringified, len(ast.tree))

	for i := 0; i < len(ast.tree); i++ {

		elements[i] = &ElementStringified{
			Name:     ast.tree[i].Name,
			Rendered: ast.tree[i].Rendered.String(),
		}
	}

	return elements
}

func (ast *AST) BlockCode(out *bytes.Buffer, text []byte, lang string) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.BlockCode(rendered, text, lang)

	ast.tree = append(ast.tree, &Element{
		Name:     "code",
		Rendered: rendered,
	})
}

func (ast *AST) BlockQuote(out *bytes.Buffer, text []byte) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.BlockQuote(rendered, text)

	ast.tree = append(ast.tree, &Element{
		Name:     "blockquote",
		Rendered: rendered,
	})
}

func (ast *AST) BlockHtml(out *bytes.Buffer, text []byte) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.BlockHtml(rendered, text)

	ast.tree = append(ast.tree, &Element{
		Name:     "html",
		Rendered: rendered,
	})
	ast.tempRendered = rendered
}

func (ast *AST) Header(out *bytes.Buffer, text func() bool, level int, id string) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.Header(rendered, text, level, id)

	ast.tree = append(ast.tree, &Element{
		Name:     fmt.Sprintf("h%d", level),
		Rendered: rendered,
	})
}

func (ast *AST) HRule(out *bytes.Buffer) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.HRule(rendered)

	ast.tree = append(ast.tree, &Element{
		Name:     "hrule",
		Rendered: rendered,
	})
}

func (ast *AST) List(out *bytes.Buffer, text func() bool, flags int) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.List(rendered, text, flags)

	ast.tree = append(ast.tree, &Element{
		Name:     "list",
		Rendered: rendered,
	})
}

func (ast *AST) ListItem(out *bytes.Buffer, text []byte, flags int) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.ListItem(rendered, text, flags)

	ast.tree = append(ast.tree, &Element{
		Name:     "listitem",
		Rendered: rendered,
	})
}

func (ast *AST) Paragraph(out *bytes.Buffer, text func() bool) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.Paragraph(rendered, text)

	ast.tree = append(ast.tree, &Element{
		Name:     "paragraph",
		Rendered: rendered,
	})
}

func (ast *AST) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.Table(rendered, header, body, columnData)

	ast.tree = append(ast.tree, &Element{
		Name:     "table",
		Rendered: rendered,
	})
}

func (ast *AST) TableRow(out *bytes.Buffer, text []byte) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.TableRow(rendered, text)

	ast.tree = append(ast.tree, &Element{
		Name:     "tablerow",
		Rendered: rendered,
	})
}

func (ast *AST) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.TableHeaderCell(rendered, text, flags)

	ast.tree = append(ast.tree, &Element{
		Name:     "th",
		Rendered: rendered,
	})
}

func (ast *AST) TableCell(out *bytes.Buffer, text []byte, flags int) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.TableCell(rendered, text, flags)

	ast.tree = append(ast.tree, &Element{
		Name:     "td",
		Rendered: rendered,
	})
}

func (ast *AST) Footnotes(out *bytes.Buffer, text func() bool) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.Footnotes(rendered, text)

	ast.tree = append(ast.tree, &Element{
		Name:     "footnotes",
		Rendered: rendered,
	})
}

func (ast *AST) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.FootnoteItem(rendered, name, text, flags)

	ast.tree = append(ast.tree, &Element{
		Name:     "footnoteitem",
		Rendered: rendered,
	})
}

func (ast *AST) TitleBlock(out *bytes.Buffer, text []byte) {

	rendered := new(bytes.Buffer)
	ast.tempRendered = rendered
	ast.renderer.TitleBlock(rendered, text)

	ast.tree = append(ast.tree, &Element{
		Name:     "tite",
		Rendered: rendered,
	})
}

// Span-level callbacks
func (ast *AST) AutoLink(out *bytes.Buffer, link []byte, kind int) {

	ast.renderer.AutoLink(
		ast.tempRendered, link, kind)
}

func (ast *AST) CodeSpan(out *bytes.Buffer, text []byte) {

	ast.renderer.CodeSpan(
		ast.tempRendered, text)
}

func (ast *AST) DoubleEmphasis(out *bytes.Buffer, text []byte) {

	ast.renderer.DoubleEmphasis(
		ast.tempRendered, text)
}

func (ast *AST) Emphasis(out *bytes.Buffer, text []byte) {

	ast.renderer.Emphasis(
		ast.tempRendered, text)
}

func (ast *AST) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {

	ast.renderer.Image(
		ast.tempRendered, link, title, alt)
}

func (ast *AST) LineBreak(out *bytes.Buffer) {

	ast.renderer.LineBreak(
		ast.tempRendered)
}

func (ast *AST) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {

	ast.renderer.Link(
		ast.tempRendered, link, title, content)
}

func (ast *AST) RawHtmlTag(out *bytes.Buffer, tag []byte) {

	ast.renderer.RawHtmlTag(
		ast.tempRendered, tag)
}

func (ast *AST) TripleEmphasis(out *bytes.Buffer, text []byte) {

	ast.renderer.TripleEmphasis(
		ast.tempRendered, text)
}

func (ast *AST) StrikeThrough(out *bytes.Buffer, text []byte) {

	ast.renderer.StrikeThrough(
		ast.tempRendered, text)
}

func (ast *AST) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {

	ast.renderer.FootnoteRef(
		ast.tempRendered, ref, id)
}

// Low-level callbacks
func (ast *AST) Entity(out *bytes.Buffer, entity []byte) {

	ast.renderer.Entity(
		ast.tempRendered, entity)
}

func (ast *AST) NormalText(out *bytes.Buffer, text []byte) {

	ast.renderer.NormalText(
		ast.tempRendered, text)
}

// Header and footer
func (ast *AST) DocumentHeader(out *bytes.Buffer) {

}

func (ast *AST) DocumentFooter(out *bytes.Buffer) {

}

func (ast *AST) GetFlags() int {

	return 0
}
