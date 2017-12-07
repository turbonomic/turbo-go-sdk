// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IdentityMetadata.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// EntityIdentityMetadata supplies meta information describing the properties used to
// identify an entity of a specific entity type.
type EntityIdentityMetadata struct {
	// The EntityType this metadata is for
	EntityType *EntityDTO_EntityType `protobuf:"varint,1,req,name=entityType,enum=proto.EntityDTO_EntityType" json:"entityType,omitempty"`
	// The version of the identifying properties for this entity type
	Version *int32 `protobuf:"varint,2,opt,name=version,def=0" json:"version,omitempty"`
	// The non-volatile identifying properties to be used
	// for this entity type. non-volatile identifying properties are the set of properties
	// necessary to identify an entity that will not change over the lifetime of the entity.
	// For example, "ID" will be a non-volatile identifying property for most entity types.
	NonVolatileProperties []*EntityIdentityMetadata_PropertyMetadata `protobuf:"bytes,3,rep,name=nonVolatileProperties" json:"nonVolatileProperties,omitempty"`
	// The volatile identifying properties to be used for this entity type.
	// Volatile identifying properties are the set of properties necessary to identify
	// an entity that may change over the lifetime of the entity. For example, for a VM,
	// the "PM_UUID" may be identifying, but moving the VM will cause the value of this property
	// to change.
	VolatileProperties []*EntityIdentityMetadata_PropertyMetadata `protobuf:"bytes,4,rep,name=volatileProperties" json:"volatileProperties,omitempty"`
	// The heuristic properties to be used for this entity type. Heuristic properties
	// are used to fuzzy match an entity's identity when an exact match using the
	// identifying non-volatile and volatile properties fails.
	HeuristicProperties []*EntityIdentityMetadata_PropertyMetadata `protobuf:"bytes,5,rep,name=heuristicProperties" json:"heuristicProperties,omitempty"`
	// The heuristic threshold is used by the identity service when matching heuristic properties
	// to determine what percentage of heuristic properties must match in order to consider
	// two objects to be the same. A heuristicThreshold of 50 would mean that at least 1/2 of
	// the heuristic properties must match for two entities to be considered to be the same.
	// This must be a value between 0 and 100.
	HeuristicThreshold *int32 `protobuf:"varint,6,opt,name=heuristicThreshold,def=75" json:"heuristicThreshold,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *EntityIdentityMetadata) Reset()                    { *m = EntityIdentityMetadata{} }
func (m *EntityIdentityMetadata) String() string            { return proto1.CompactTextString(m) }
func (*EntityIdentityMetadata) ProtoMessage()               {}
func (*EntityIdentityMetadata) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

const Default_EntityIdentityMetadata_Version int32 = 0
const Default_EntityIdentityMetadata_HeuristicThreshold int32 = 75

func (m *EntityIdentityMetadata) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *EntityIdentityMetadata) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Default_EntityIdentityMetadata_Version
}

func (m *EntityIdentityMetadata) GetNonVolatileProperties() []*EntityIdentityMetadata_PropertyMetadata {
	if m != nil {
		return m.NonVolatileProperties
	}
	return nil
}

func (m *EntityIdentityMetadata) GetVolatileProperties() []*EntityIdentityMetadata_PropertyMetadata {
	if m != nil {
		return m.VolatileProperties
	}
	return nil
}

func (m *EntityIdentityMetadata) GetHeuristicProperties() []*EntityIdentityMetadata_PropertyMetadata {
	if m != nil {
		return m.HeuristicProperties
	}
	return nil
}

func (m *EntityIdentityMetadata) GetHeuristicThreshold() int32 {
	if m != nil && m.HeuristicThreshold != nil {
		return *m.HeuristicThreshold
	}
	return Default_EntityIdentityMetadata_HeuristicThreshold
}

type EntityIdentityMetadata_PropertyMetadata struct {
	// The name of the property.
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EntityIdentityMetadata_PropertyMetadata) Reset() {
	*m = EntityIdentityMetadata_PropertyMetadata{}
}
func (m *EntityIdentityMetadata_PropertyMetadata) String() string { return proto1.CompactTextString(m) }
func (*EntityIdentityMetadata_PropertyMetadata) ProtoMessage()    {}
func (*EntityIdentityMetadata_PropertyMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0}
}

func (m *EntityIdentityMetadata_PropertyMetadata) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto1.RegisterType((*EntityIdentityMetadata)(nil), "proto.EntityIdentityMetadata")
	proto1.RegisterType((*EntityIdentityMetadata_PropertyMetadata)(nil), "proto.EntityIdentityMetadata.PropertyMetadata")
}

func init() { proto1.RegisterFile("IdentityMetadata.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x69, 0xbb, 0x2a, 0x3e, 0x41, 0xe5, 0x89, 0xa3, 0x6c, 0x07, 0xab, 0x07, 0xe9, 0x29,
	0x48, 0x41, 0x04, 0xbd, 0xa9, 0x3b, 0x78, 0x10, 0xa5, 0x14, 0x8f, 0x62, 0x6c, 0x23, 0x0d, 0x36,
	0x79, 0x25, 0xcd, 0x0a, 0xfb, 0xc7, 0x3d, 0xcb, 0x5a, 0x2d, 0x73, 0xeb, 0x69, 0xa7, 0x84, 0x2f,
	0xdf, 0xf7, 0xfd, 0x48, 0xf2, 0x60, 0xfc, 0x98, 0x0b, 0x6d, 0xa5, 0x5d, 0x3c, 0x09, 0xcb, 0x73,
	0x6e, 0x39, 0xab, 0x0c, 0x59, 0x42, 0xbf, 0x5d, 0x26, 0x87, 0xf7, 0xa4, 0x14, 0xe9, 0x87, 0xf4,
	0xb9, 0xd3, 0xcf, 0xbf, 0x3d, 0x18, 0xcf, 0xda, 0xc0, 0x7a, 0x10, 0x6f, 0x01, 0x3a, 0x25, 0x5d,
	0x54, 0x22, 0x70, 0x42, 0x37, 0x3a, 0x88, 0xa7, 0x5d, 0x8c, 0x75, 0x91, 0x65, 0xcd, 0xac, 0xb7,
	0x24, 0x2b, 0x76, 0x9c, 0xc2, 0x6e, 0x23, 0x4c, 0x2d, 0x49, 0x07, 0x6e, 0xe8, 0x44, 0xfe, 0x8d,
	0x73, 0x99, 0xfc, 0x29, 0x98, 0xc3, 0x89, 0x26, 0xfd, 0x4a, 0x25, 0xb7, 0xb2, 0x14, 0x2f, 0x86,
	0x2a, 0x61, 0xac, 0x14, 0x75, 0xe0, 0x85, 0x5e, 0xb4, 0x1f, 0xb3, 0x7f, 0x90, 0x8d, 0x0b, 0xfd,
	0xfa, 0x7b, 0x21, 0x19, 0x2e, 0xc3, 0x37, 0xc0, 0x66, 0x13, 0x31, 0xda, 0x0a, 0x31, 0xd0, 0x84,
	0xef, 0x70, 0x5c, 0x88, 0xb9, 0x91, 0xb5, 0x95, 0xd9, 0x0a, 0xc0, 0xdf, 0x0a, 0x30, 0x54, 0x85,
	0x31, 0x60, 0x2f, 0xa7, 0x85, 0x11, 0x75, 0x41, 0x65, 0x1e, 0xec, 0xb4, 0xef, 0xe9, 0x5e, 0x5f,
	0x25, 0x03, 0xa7, 0x93, 0x0b, 0x38, 0x5a, 0x2f, 0x47, 0x84, 0x91, 0xe6, 0x6a, 0xf9, 0x87, 0x4e,
	0xb4, 0x97, 0xb4, 0xfb, 0xbb, 0x33, 0x38, 0xcd, 0x48, 0xb1, 0x46, 0xd9, 0xb9, 0xf9, 0x20, 0x56,
	0x95, 0xdc, 0x7e, 0x92, 0x51, 0xac, 0xce, 0xbf, 0x58, 0xd6, 0x0e, 0xc9, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd9, 0x47, 0x95, 0x77, 0x4c, 0x02, 0x00, 0x00,
}
