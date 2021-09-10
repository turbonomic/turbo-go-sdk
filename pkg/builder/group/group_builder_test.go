package group

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"testing"
)

func TestGroupBuilderInvalid(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	selectionSpec := StringProperty().
		Name("Property1").
		Expression(proto.GroupDTO_SelectionSpec_EQUAL_TO).
		SetProperty("Hello")

	// valid entity type, no members
	_, err := StaticRegularGroup(id).OfType(eType).Build()
	fmt.Printf("******** StaticGroup Error %s\n", err)
	assert.NotNil(t, err)

	_, err = DynamicRegularGroup(id).OfType(eType).Build()
	fmt.Printf("******** DynamicGroup Error %s\n", err)
	assert.NotNil(t, err)

	// no entity type
	_, err = StaticRegularGroup(id).WithEntities([]string{"abc", "xyz"}).Build()
	fmt.Printf("******** StaticGroup Error %s\n", err)
	assert.NotNil(t, err)

	_, err = DynamicRegularGroup(id).MatchingEntities(SelectedBy(selectionSpec)).Build()
	fmt.Printf("******** DynamicGroup Error %s\n", err)
	assert.NotNil(t, err)

	// matching criterion for static group
	_, err = StaticRegularGroup(id).OfType(eType).MatchingEntities(SelectedBy(selectionSpec)).WithEntities([]string{"abc", "xyz"}).Build()
	fmt.Printf("####### StaticGroup Error %s\n", err)
	assert.NotNil(t, err)

	// member list for dynamic group
	_, err = DynamicRegularGroup(id).OfType(eType).WithEntities([]string{"abc", "xyz"}).MatchingEntities(SelectedBy(selectionSpec)).Build()
	fmt.Printf("####### DynamicGroup Error %s\n", err)
	assert.NotNil(t, err)
}

func TestGroupBuilderEntityType(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	// valid entity type
	groupBuilder := StaticRegularGroup(id).OfType(eType).WithEntities([]string{"abc", "xyz"})
	assert.NotNil(t, groupBuilder)

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, groupDTO.EntityType)
	assert.EqualValues(t, eType, groupDTO.GetEntityType())

	// unknown entity type
	var fakeType proto.EntityDTO_EntityType = 200
	groupBuilder = StaticRegularGroup(id).OfType(fakeType).WithEntities([]string{"abc", "xyz"})

	groupDTO, err = groupBuilder.Build()
	assert.NotNil(t, err)
	//
	// default entity type
	eType = 0
	groupBuilder3 := StaticRegularGroup(id).OfType(eType).WithEntities([]string{"abc", "xyz"})

	groupDTO3, err := groupBuilder3.Build()
	assert.Nil(t, err)
	assert.NotNil(t, groupDTO3.EntityType)
	assert.EqualValues(t, proto.EntityDTO_SWITCH, groupDTO3.GetEntityType())

	// overwrite existing entity type
	groupBuilder3 = groupBuilder3.OfType(30).WithEntities([]string{"abc", "xyz"})

	groupDTO3, err = groupBuilder3.Build()
	assert.NotNil(t, err)
}

func TestGroupBuilderDynamic(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	selectionSpec := StringProperty().
		Name("Property1").
		Expression(proto.GroupDTO_SelectionSpec_EQUAL_TO).
		SetProperty("Hello")

	selectionSpec2 := StringProperty().
		Name("Property2").
		Expression(proto.GroupDTO_SelectionSpec_EQUAL_TO).
		SetProperty("World")

	groupBuilder := DynamicRegularGroup(id).
		OfType(eType).
		MatchingEntities(SelectedBy(selectionSpec).and(selectionSpec2))

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)

	assert.EqualValues(t, eType, groupDTO.GetEntityType())
	assert.EqualValues(t, id, groupDTO.GetDisplayName())

	assert.EqualValues(t, 2, len(groupDTO.GetSelectionSpecList().SelectionSpec))
	assert.Nil(t, groupDTO.GetMemberList())
}

func TestGroupBuilderStatic(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	groupBuilder := StaticRegularGroup(id).
		OfType(eType).
		WithEntities([]string{"abc", "xyz"})

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)
	assert.EqualValues(t, eType, groupDTO.GetEntityType())
	assert.EqualValues(t, id, groupDTO.GetDisplayName())
	assert.EqualValues(t, 2, len(groupDTO.GetMemberList().GetMember()))
	assert.Nil(t, groupDTO.GetSelectionSpecList())
}

func TestGroupBuilderStaticNodePool(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_VIRTUAL_MACHINE

	groupBuilder := StaticNodePool(id).
		OfType(eType).
		WithEntities([]string{"abc", "xyz"})

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)
	assert.EqualValues(t, eType, groupDTO.GetEntityType())
	assert.EqualValues(t, id, groupDTO.GetDisplayName())
	assert.EqualValues(t, 2, len(groupDTO.GetMemberList().GetMember()))
	assert.EqualValues(t, proto.GroupDTO_NODE_POOL.Type(), groupDTO.GetGroupType().Type())
	assert.Nil(t, groupDTO.GetSelectionSpecList())
}

func TestGroupBuilderSetConsistentResize(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	groupBuilder := StaticRegularGroup(id).
		OfType(eType).
		WithEntities([]string{"abc", "xyz"})

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)
	consistentResize := groupDTO.GetIsConsistentResizing()
	assert.False(t, consistentResize)

	fmt.Printf("%++v\n", groupDTO)

	groupBuilder = StaticRegularGroup(id).
		OfType(eType).
		WithEntities([]string{"abc", "xyz"}).
		ResizeConsistently()

	groupDTO, err = groupBuilder.Build()
	assert.Nil(t, err)
	consistentResize = groupDTO.GetIsConsistentResizing()
	assert.True(t, consistentResize)

	fmt.Printf("%++v\n", groupDTO)
}

func TestGroupBuilderDisplayName(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	groupBuilder := StaticRegularGroup(id).
		OfType(eType).
		WithEntities([]string{"abc", "xyz"})

	groupDTO, _ := groupBuilder.Build()

	assert.Equal(t, id, groupDTO.GetDisplayName())

	displayName := "Test Group"
	groupBuilder.WithDisplayName(displayName)
	groupDTO, _ = groupBuilder.Build()

	assert.Equal(t, displayName, groupDTO.GetDisplayName())
}
