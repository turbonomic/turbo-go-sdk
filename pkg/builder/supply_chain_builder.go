package builder

import (
	"fmt"

	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
)

type SupplyChainBuilder struct {
	SupplyChainNodes map[*proto.EntityDTO_EntityType]*SupplyChainNodeBuilder
	currentNode      *SupplyChainNodeBuilder

	err error
}

func NewSupplyChainBuilder() *SupplyChainBuilder {
	return &SupplyChainBuilder{}
}

// TODO: supply chain here is a set(map) or a list
// if a map, key is the node builder, value is a boolean
func (scb *SupplyChainBuilder) Create() ([]*proto.TemplateDTO, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	allNodeBuilders := make(map[*SupplyChainNodeBuilder]bool)
	for _, nodeBuilder := range scb.SupplyChainNodes {
		allNodeBuilders[nodeBuilder] = true
	}

	// create nodes from all node builders and put it into result
	var allNodes []*proto.TemplateDTO
	for nodeBuilder, _ := range allNodeBuilders {
		templateDTO, err := nodeBuilder.Create()
		if err != nil {
			return nil, fmt.Errorf("Error in creating supplychain node: %v", err)
		}
		allNodes = append(allNodes, templateDTO)
	}

	return allNodes, nil
}

// To build a supply chain, must specify the top nodebuilder first.
// Here it initialize SupplyChainNodes and currentNode
func (scb *SupplyChainBuilder) Top(topNode *SupplyChainNodeBuilder) *SupplyChainBuilder {
	supplyChianNodesBuilderMap := make(map[*proto.EntityDTO_EntityType]*SupplyChainNodeBuilder)
	scb.SupplyChainNodes = supplyChianNodesBuilderMap
	topNodeEntityType, err := topNode.getType()
	if err != nil {
		scb.err = err
		return scb
	}
	scb.SupplyChainNodes[topNodeEntityType] = topNode

	scb.currentNode = topNode

	return scb
}

// Add an entity node to supply chain
func (scb *SupplyChainBuilder) Entity(node *SupplyChainNodeBuilder) *SupplyChainBuilder {
	if hasTop := scb.hasTopNode(); !hasTop {
		scb.err = fmt.Errorf("Must set top supply chain node first.")
		return scb
	}
	nodeEntityType, err := node.getType()
	if err != nil {
		scb.err = err
		return scb
	}
	scb.SupplyChainNodes[nodeEntityType] = node

	scb.currentNode = node

	return scb
}

// check if the top node has been initialized
func (scb *SupplyChainBuilder) hasTopNode() bool {
	if scb.SupplyChainNodes == nil {
		return false
	}
	return true
}

// Adds an external entity link to the current node.
// An external entity is on ethat exists in teh Operations Manager supply chain, but has not been
// discovered by the probe. Operations Manager uses this link by the Operations Manager market. This
// external entity can be a provider or a consumer.
func (scb *SupplyChainBuilder) ConnectsTo(extEntityLink *proto.ExternalEntityLink) *SupplyChainBuilder {
	if scb.err != nil {
		return scb
	}
	err := scb.requireCurrentNode()
	if err != nil {
		scb.err = err
		return scb
	}

	scb.currentNode = scb.currentNode.Link(extEntityLink)

	return scb
}

// check if currentNode is set.
func (scb *SupplyChainBuilder) requireCurrentNode() error {
	if scb.currentNode == nil {
		return fmt.Errorf("Illegal state, currentNode is nil")
	}
	return nil
}
