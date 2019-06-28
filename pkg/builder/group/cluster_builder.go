package group

import "github.com/turbonomic/turbo-go-sdk/pkg/proto"

type ClusterBuilder struct {
	*AbstractConstraintGroupBuilder
}

// Cluster is the builder for a group with Cluster constraint
func Cluster(id string) *ClusterBuilder {
	return &ClusterBuilder{
		&AbstractConstraintGroupBuilder{
			StaticGroup(id),
			newConstraintBuilder(proto.GroupDTO_CLUSTER, id),
		},
	}
}

func (c *ClusterBuilder) OfType(eType proto.EntityDTO_EntityType) *ClusterBuilder {
	c.AbstractBuilder.OfType(eType)
	return c
}

func (c *ClusterBuilder) WithEntities(entities []string) *ClusterBuilder {
	c.AbstractBuilder.WithEntities(entities)
	return c
}

func (c *ClusterBuilder) WithDisplayName(displayName string) *ClusterBuilder {
	c.AbstractBuilder.WithDisplayName(displayName)
	c.ConstraintInfoBuilder.WithDisplayName(displayName)
	return c
}

func (c *ClusterBuilder) Build() (*proto.GroupDTO, error) {
	return c.AbstractConstraintGroupBuilder.Build()
}
