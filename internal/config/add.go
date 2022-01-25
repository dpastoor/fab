package config

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func appendStringNode(node *yaml.Node, value string, quote bool) {
	var newNode yaml.Node
	newNode.Kind = yaml.ScalarNode
	newNode.Value = value
	newNode.Tag = "!!str"
	if quote {
		newNode.Style = yaml.DoubleQuotedStyle
	} else {
		newNode.Style = yaml.LiteralStyle
	}
	node.Content = append(node.Content, &newNode)
}

func addStringToArrayNode(root yaml.Node, identifier string, s string, quote bool) (yaml.Node, error) {
	collectionNode := iterateNode(&root, identifier)
	appendStringNode(collectionNode, s, quote)
	return root, nil
}

func AddPathsToCollections(root yaml.Node, paths []string, quote bool, validate bool) (yaml.Node, error) {
	if validate {
		var missingPaths []string
		for _, c := range paths {
			if !exists(c) {
				missingPaths = append(missingPaths, c)
			}
		}
		if len(missingPaths) > 0 {
			return root, fmt.Errorf("no collections at paths:\n- %s", strings.Join(missingPaths, "\n- "))
		}
	}
	for _, c := range paths {
		t, err := addStringToArrayNode(root, "collections", c, quote)
		if err != nil {
			return t, err
		}
	}
	return yaml.Node{}, nil
}
