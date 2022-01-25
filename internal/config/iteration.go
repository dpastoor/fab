package config

import "gopkg.in/yaml.v3"

// from: https://stackoverflow.com/questions/55674853/modify-existing-yaml-file-and-add-new-data-and-comments
// iterateNode will recursive look for the node following the identifier Node,
// as go-yaml has a node for the key and the value itself
// we want to manipulate the value Node
func iterateNode(node *yaml.Node, identifier string) *yaml.Node {
	returnNode := false
	for _, n := range node.Content {
		if n.Value == identifier {
			returnNode = true
			continue
		}
		if returnNode {
			return n
		}
		if len(n.Content) > 0 {
			ac_node := iterateNode(n, identifier)
			if ac_node != nil {
				return ac_node
			}
		}
	}
	return nil
}
