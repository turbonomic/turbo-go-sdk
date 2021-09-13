package group

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type GroupType string

const (
	REGULAR   GroupType = "REGULAR"
	NODE_POOL GroupType = "NODE_POOL"
)

// Builder for creating a GroupDTO
type AbstractBuilder struct {
	groupId          string
	displayName      string
	owner            string
	entityTypePtr    *proto.EntityDTO_EntityType
	memberList       []string
	matching         *Matching
	consistentResize bool
	//groupDTO *proto.GroupDTO
	ec        *builder.ErrorCollector
	groupType GroupType
	isStatic  bool
}

// Create a new instance of AbstractBuilder.
// Specify the group id and if the group is static or dynamic.
func newAbstractBuilder(id string, groupType GroupType, isStatic bool) *AbstractBuilder {
	groupBuilder := &AbstractBuilder{
		groupType:        groupType,
		groupId:          id,
		ec:               new(builder.ErrorCollector),
		consistentResize: false,
		isStatic:         isStatic,
	}
	return groupBuilder
}

// Create a new instance of builder for creating Static groups.
// Static group contains a fixed list of entity id's
func StaticGroup(id string, groupType GroupType) *AbstractBuilder {
	groupBuilder := newAbstractBuilder(id, groupType, true)
	return groupBuilder
}

// Create a new instance of builder for creating REGULAR Static groups.
// Static group contains a fixed list of entity id's
func StaticRegularGroup(id string) *AbstractBuilder {
	groupBuilder := newAbstractBuilder(id, REGULAR, true)
	return groupBuilder
}

// Create a new instance of builder for creating NODE_POOL Static groups.
// Static group contains a fixed list of entity id's
func StaticNodePool(id string) *AbstractBuilder {
	groupBuilder := newAbstractBuilder(id, NODE_POOL, true)
	return groupBuilder
}

// Create a new instance of builder for creating Dynamic groups.
// Dynamic group contains selection criteria using entity properties to select entities.
func DynamicGroup(id string, groupType GroupType) *AbstractBuilder {
	groupBuilder := newAbstractBuilder(id, groupType, false)
	return groupBuilder
}

// Create a new instance of builder for creating REGULAR Dynamic groups.
// Dynamic group contains selection criteria using entity properties to select entities.
func DynamicRegularGroup(id string) *AbstractBuilder {
	groupBuilder := newAbstractBuilder(id, REGULAR, false)
	return groupBuilder
}

// Create a new instance of builder for creating NODE_POOL Dynamic groups.
// Dynamic group contains selection criteria using entity properties to select entities.
func DynamicNodePoolGroup(id string) *AbstractBuilder {
	groupBuilder := newAbstractBuilder(id, NODE_POOL, false)
	return groupBuilder
}

// Return the Protobuf GroupDTO object. There is no constraint object with this group.
// Return error if errors were collected during the building of the group properties.
func (groupBuilder *AbstractBuilder) Build() (*proto.GroupDTO, error) {

	groupId := &proto.GroupDTO_GroupName{
		GroupName: groupBuilder.groupId,
	}
	groupDTO := &proto.GroupDTO{
		DisplayName: &groupBuilder.groupId,
		Info:        groupId,
	}

	if groupBuilder.displayName != "" {
		groupDTO.DisplayName = &groupBuilder.displayName
	}

	err := groupBuilder.setupEntityType(groupDTO)
	if err != nil {
		groupBuilder.ec.Collect(err)
	}

	if groupBuilder.isStatic {
		err := groupBuilder.setUpStaticMembers(groupDTO)
		if err != nil {
			groupBuilder.ec.Collect(err)
		}
	} else {
		err := groupBuilder.setUpDynamicGroup(groupDTO)
		if err != nil {
			groupBuilder.ec.Collect(err)
		}
	}

	if groupBuilder.groupType == REGULAR {
		regular := proto.GroupDTO_REGULAR
		groupDTO.GroupType = &regular
	} else if groupBuilder.groupType == NODE_POOL {
		resource := proto.GroupDTO_NODE_POOL
		groupDTO.GroupType = &resource
	}

	if groupBuilder.owner != "" {
		groupDTO.Owner = &groupBuilder.owner
	}

	groupDTO.IsConsistentResizing = &groupBuilder.consistentResize

	if groupBuilder.ec.Count() > 0 {
		glog.Errorf("%s : %s", groupBuilder.groupId, groupBuilder.ec.Error())
		return nil, fmt.Errorf("%s: %s", groupBuilder.groupId, groupBuilder.ec.Error())
	}

	return groupDTO, nil
}

func (groupBuilder *AbstractBuilder) WithDisplayName(displayName string) *AbstractBuilder {
	if displayName == "" {
		return groupBuilder
	}
	// Setup entity type
	groupBuilder.displayName = displayName

	return groupBuilder
}

func (groupBuilder *AbstractBuilder) WithOwner(owner string) *AbstractBuilder {
	if owner == "" {
		return groupBuilder
	}
	// Setup entity type
	groupBuilder.owner = owner

	return groupBuilder
}

// Set the entity type for the members of the group.
// All the entities in a group belong to the same entity type.
func (groupBuilder *AbstractBuilder) OfType(eType proto.EntityDTO_EntityType) *AbstractBuilder {

	// Check entity type
	if groupBuilder.entityTypePtr != nil && *groupBuilder.entityTypePtr != eType {
		groupBuilder.ec.Collect(fmt.Errorf("cannot add members, input entityType %v is not consistent with existing entityType %v",
			eType, *groupBuilder.entityTypePtr))
		return groupBuilder
	}

	// Setup entity type
	groupBuilder.entityTypePtr = &eType

	return groupBuilder
}

func (groupBuilder *AbstractBuilder) setupEntityType(groupDTO *proto.GroupDTO) error {
	if groupBuilder.entityTypePtr == nil {
		return fmt.Errorf("entity type is not set")
	}
	// Validate entity type
	entityType := *groupBuilder.entityTypePtr
	_, valid := proto.EntityDTO_EntityType_name[int32(entityType)]

	if !valid {
		return fmt.Errorf("invalid entity type %v", entityType)
	}

	// Setup entity type
	groupDTO.EntityType = &entityType
	return nil
}

// Set the members for a static group. Input is a list of UUIDs for the entities that belong to the group.
func (groupBuilder *AbstractBuilder) WithEntities(entities []string) *AbstractBuilder {

	// Assert that the group is a static group
	if !groupBuilder.isStatic {
		groupBuilder.ec.Collect(fmt.Errorf("cannot set member uuid list for dynamic group"))
		return groupBuilder
	}
	groupBuilder.memberList = entities

	return groupBuilder
}

func (groupBuilder *AbstractBuilder) setUpStaticMembers(groupDTO *proto.GroupDTO) error {

	if len(groupBuilder.memberList) == 0 {
		return fmt.Errorf("empty member list")
	}

	// Set the Group DTO member field
	memberList := &proto.GroupDTO_MemberList{
		MemberList: &proto.GroupDTO_MembersList{
			Member: groupBuilder.memberList,
		},
	}
	groupDTO.Members = memberList
	return nil
}

// Set the members matching criteria for a dynamic group.
func (groupBuilder *AbstractBuilder) MatchingEntities(matching *Matching) *AbstractBuilder {

	// Assert that the group is a dynamci group
	if groupBuilder.isStatic {
		groupBuilder.ec.Collect(fmt.Errorf("cannot set matching criteria for static group"))
		return groupBuilder
	}
	groupBuilder.matching = matching

	return groupBuilder
}

func (groupBuilder *AbstractBuilder) setUpDynamicGroup(groupDTO *proto.GroupDTO) error {

	if groupBuilder.matching == nil {
		return fmt.Errorf("null matching criteria for member selection")
	}

	var selectionSpecList []*proto.GroupDTO_SelectionSpec

	// Build the selection spec list from the matching criteria
	selectionSpecBuilderList := groupBuilder.matching.selectionSpecBuilderList
	for _, specBuilder := range selectionSpecBuilderList {
		selectionSpec := specBuilder.Build()
		selectionSpecList = append(selectionSpecList, selectionSpec)
	}

	// Set the Group DTO member field
	selectionSpecList_ := &proto.GroupDTO_SelectionSpecList_{
		SelectionSpecList: &proto.GroupDTO_SelectionSpecList{
			SelectionSpec: selectionSpecList,
		},
	}
	groupDTO.Members = selectionSpecList_
	return nil
}

func (groupBuilder *AbstractBuilder) ResizeConsistently() *AbstractBuilder {
	groupBuilder.consistentResize = true
	return groupBuilder
}
