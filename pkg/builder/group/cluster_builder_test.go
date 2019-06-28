package group

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"testing"
)

func TestClusterHasWrongEntityType(t *testing.T) {
	clusterBuilder := Cluster("cluster1")
	_, err := clusterBuilder.
		OfType(proto.EntityDTO_CONTAINER_POD).
		WithEntities([]string{"vm1", "vm2"}).
		Build()
	fmt.Printf("Cluster error %v\n", err)
	assert.NotNil(t, err)
}

func TestClusterHasNoMemberList(t *testing.T) {
	clusterBuilder := Cluster("cluster1")
	_, err := clusterBuilder.
		OfType(proto.EntityDTO_VIRTUAL_MACHINE).
		Build()
	fmt.Printf("Cluster error %v\n", err)
	assert.NotNil(t, err)
}

func TestClusterHasClusterConstraintInfo(t *testing.T) {
	clusterBuilder := Cluster("cluster1")
	clusterDTO, err := clusterBuilder.
		OfType(proto.EntityDTO_VIRTUAL_MACHINE).
		WithEntities([]string{"vm1", "vm2"}).
		Build()
	assert.Nil(t, err)
	assert.Equal(t, clusterDTO.GetConstraintInfo().GetConstraintType(), proto.GroupDTO_CLUSTER)
}
