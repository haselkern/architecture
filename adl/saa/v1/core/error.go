package core

import (
	"errors"
	"fmt"
	"github.com/golangee/src/ast"
	"strconv"
	"strings"
)

type node struct {
	NodeSrc          string
	NodePos, NodeEnd Pos
}

func (n node) Src() string {
	return n.NodeSrc
}

func (n node) Pos() Pos {
	return n.NodePos
}

func (n node) End() Pos {
	return n.NodeEnd
}

type Node interface {
	// Src contains the original textual representation to which Pos and End refers.
	Src() string
	Pos() Pos
	End() Pos
}

func NewNodeFromAst(n ast.Node) Node {
	return node{
		NodePos: Pos(n.Pos()),
		NodeEnd: Pos(n.End()),
	}
}

type ErrDetail struct {
	Node    Node
	Message string
}

// PosError represents a very specific positional error with a lot of explaining noise. Use Explain.
type PosError struct {
	Node    Node
	Message string
	Details []ErrDetail
	Cause   error
	Hint    string
}

// NewPosError creates a new PosError with the given root cause and optional details.
func NewPosError(node Node, msg string, details ...ErrDetail) *PosError {
	return &PosError{
		Node:    node,
		Message: msg,
		Details: details,
	}
}

func (p *PosError) SetCause(err error) *PosError {
	p.Cause = err
	if err != nil {
		p.Details = append(p.Details, ErrDetail{p.Node, err.Error()})
	}
	return p
}

func (p *PosError) SetHint(str string) *PosError {
	p.Hint = str
	return p
}

func (p *PosError) Unwrap() error {
	return p.Cause
}

func (p *PosError) Error() string {
	if p.Cause == nil {
		return p.Message
	}

	return p.Message + ": " + p.Cause.Error()
}

// docLines returns associated source lines to the given node. It evaluate the magic attribute "src" from Obj
// which has the Stereotype Document.
func docLines(n Node) []string {
	return strings.Split(n.Src(), "\n")
}

// posLine returns the line from lines which fits to the given pos.
func posLine(lines []string, pos Pos) string {
	no := pos.Line - 1

	if no > len(lines) {
		no = len(lines) - 1
	}

	ltext := ""
	if no < len(lines) && no >= 0 {
		ltext = lines[no]
	}

	return ltext
}

// Explain returns a multi-line text suited to be printed into the console.
func (p PosError) Explain() string {
	// grab the required indent for the line numbers
	indent := 0
	for _, detail := range p.Details {
		l := len(strconv.Itoa(detail.Node.Pos().Line))
		if l > indent {
			indent = l
		}
	}

	sb := &strings.Builder{}
	sb.WriteString("error: ")
	sb.WriteString(p.Message)
	sb.WriteString("\n")
	for i := 0; i < indent; i++ {
		sb.WriteByte(' ')
	}
	sb.WriteString("--> ")
	if p.Node == nil {
		sb.WriteString("node is nil")
		return sb.String()
	}

	sb.WriteString(p.Node.Pos().String())
	sb.WriteString("\n")

	for i, detail := range p.Details {
		source := docLines(detail.Node)
		line := posLine(source, detail.Node.Pos())

		if detail.Node.Pos().File != p.Node.Pos().File {
			sb.WriteString(p.Node.Pos().String())
			sb.WriteString("\n")
		}

		sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"s |\n", ""))
		sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"d |", detail.Node.Pos().Line))
		sb.WriteString(line)
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"s |", ""))

		if detail.Node.End().Col-detail.Node.Pos().Col <= 1 {
			sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(detail.Node.Pos().Col-1)+"s", ""))
			sb.WriteString("^~~~ ")
		} else {
			sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(detail.Node.Pos().Col-1)+"s", ""))
			for i := 0; i < detail.Node.End().Col-detail.Node.Pos().Col; i++ {
				sb.WriteRune('^')
			}
			sb.WriteRune(' ')
		}

		sb.WriteString(detail.Message)
		sb.WriteString("\n")

		if i < len(p.Details)-1 {
			for i := 0; i < indent; i++ {
				sb.WriteByte(' ')
			}
			sb.WriteString("...")
			sb.WriteByte('\n')
		}
	}

	if p.Hint != "" {
		sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"s |\n", ""))
		sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"s = hint: %s\n", "", p.Hint))
	}

	return sb.String()
}

// Explain takes the given wrapped error chain and explains it, if it can.
func Explain(err error) string {
	var posErr *PosError
	if errors.As(err, &posErr) {
		return posErr.Explain()
	}

	return err.Error()
}