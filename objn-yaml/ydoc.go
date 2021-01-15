package objnyaml

import (
	"fmt"
	"github.com/golangee/architecture/objn"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
)

type YamlDoc struct {
	pos  objn.Pos
	root *yaml.Node
	text string
}

func (n *YamlDoc) Comment() string {
	return n.root.HeadComment
}

func (n *YamlDoc) Pos() objn.Pos {
	return n.pos
}

func (n *YamlDoc) String() string {
	return n.text
}

func (n *YamlDoc) Validate() error {
	if v, ok := n.Root().(validateable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (n *YamlDoc) Root() objn.Node {
	if len(n.root.Content) == 0 {
		return nil
	}

	return NewNode(n.pos.File, n.root.Content[0])
}

func NewYamlDoc(fname string) (*YamlDoc, error) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}

	node := &yaml.Node{}
	err = yaml.Unmarshal(buf, node)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal yaml file: %w", err)
	}

	if len(buf) > 0 {
		if node.Kind != yaml.DocumentNode {
			return nil, fmt.Errorf("invalid node kind: " + strconv.Itoa(int(node.Kind)))
		}
	}

	if len(node.Content) > 1 {
		return nil, fmt.Errorf("invalid content length")
	}

	n := &YamlDoc{
		pos: objn.Pos{
			File: fname,
			Line: node.Line,
			Col:  node.Column,
		},
		root: node,
		text: string(buf),
	}

	return n, nil
}

func NewNode(filename string, node *yaml.Node) objn.Node {
	switch node.Kind {
	case yaml.MappingNode:
		return NewYamlMap(filename, node)
	case yaml.SequenceNode:
		return NewYamlSeq(filename, node)
	case yaml.ScalarNode:
		return NewYamlLit(objn.Pos{
			File: filename,
			Line: node.Line,
			Col:  node.Column,
		}, node.Value, node.HeadComment)
	default:
		panic("yaml node type not supported: " + strconv.Itoa(int(node.Kind)))
	}
}