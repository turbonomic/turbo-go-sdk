// Code generated by protoc-gen-go.
// source: ProfileDTO.proto
// DO NOT EDIT!

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This file lists all the objects related to Service Entity profiles
// created by user in environment or in VMTurbo
type EntityProfileDTO struct {
	// id of the entity profile.  This should allow access to the object
	// in the system it was discovered from and it should be unique in VMT also.
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// display name to be displayed to the user
	DisplayName *string `protobuf:"bytes,2,req,name=displayName" json:"displayName,omitempty"`
	// Type of entity this profile applies to
	EntityType *EntityDTO_EntityType `protobuf:"varint,3,req,name=entityType,enum=common_dto.EntityDTO_EntityType" json:"entityType,omitempty"`
	// The profile should contain multiple related commodity profiles for example
	// profile for MEM, CPU, VSTORAGE...
	CommodityProfile []*CommodityProfileDTO `protobuf:"bytes,4,rep,name=commodityProfile" json:"commodityProfile,omitempty"`
	// Model related to the profile
	Model *string `protobuf:"bytes,5,opt,name=model" json:"model,omitempty"`
	// Vendor related to the profile
	Vendor *string `protobuf:"bytes,6,opt,name=vendor" json:"vendor,omitempty"`
	// Description of the profile
	Description *string `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	// If this is a profile for VMs, vmProfileDTO must be specified
	// If this is a profile for PMs, pmProfileDTO must be specified
	// If this is a profile for DBs or DBInstances, dbProfileDTO must be specified
	//
	// Types that are valid to be assigned to EntityTypeSpecificData:
	//	*EntityProfileDTO_VmProfileDTO
	//	*EntityProfileDTO_PmProfileDTO
	//	*EntityProfileDTO_DbProfileDTO
	EntityTypeSpecificData isEntityProfileDTO_EntityTypeSpecificData `protobuf_oneof:"EntityTypeSpecificData"`
	// This flag indicates where existing entities can be matched against this profile
	EnableProvisionMatch *bool `protobuf:"varint,10,opt,name=enableProvisionMatch" json:"enableProvisionMatch,omitempty"`
	// This flag indicates whether a resize action may use this profile to resize to
	EnableResizeMatch *bool `protobuf:"varint,11,opt,name=enableResizeMatch" json:"enableResizeMatch,omitempty"`
	// Allow entity properties to be specified related to the entity profile dto.
	// Entity properties are a list of <string, string, string> namespace, key, value triplets
	EntityProperties []*EntityDTO_EntityProperty `protobuf:"bytes,12,rep,name=entityProperties" json:"entityProperties,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *EntityProfileDTO) Reset()                    { *m = EntityProfileDTO{} }
func (m *EntityProfileDTO) String() string            { return proto.CompactTextString(m) }
func (*EntityProfileDTO) ProtoMessage()               {}
func (*EntityProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type isEntityProfileDTO_EntityTypeSpecificData interface {
	isEntityProfileDTO_EntityTypeSpecificData()
}

type EntityProfileDTO_VmProfileDTO struct {
	VmProfileDTO *EntityProfileDTO_VMProfileDTO `protobuf:"bytes,8,opt,name=vmProfileDTO,oneof"`
}
type EntityProfileDTO_PmProfileDTO struct {
	PmProfileDTO *EntityProfileDTO_PMProfileDTO `protobuf:"bytes,9,opt,name=pmProfileDTO,oneof"`
}
type EntityProfileDTO_DbProfileDTO struct {
	DbProfileDTO *EntityProfileDTO_DBProfileDTO `protobuf:"bytes,13,opt,name=dbProfileDTO,oneof"`
}

func (*EntityProfileDTO_VmProfileDTO) isEntityProfileDTO_EntityTypeSpecificData() {}
func (*EntityProfileDTO_PmProfileDTO) isEntityProfileDTO_EntityTypeSpecificData() {}
func (*EntityProfileDTO_DbProfileDTO) isEntityProfileDTO_EntityTypeSpecificData() {}

func (m *EntityProfileDTO) GetEntityTypeSpecificData() isEntityProfileDTO_EntityTypeSpecificData {
	if m != nil {
		return m.EntityTypeSpecificData
	}
	return nil
}

func (m *EntityProfileDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *EntityProfileDTO) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *EntityProfileDTO) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *EntityProfileDTO) GetCommodityProfile() []*CommodityProfileDTO {
	if m != nil {
		return m.CommodityProfile
	}
	return nil
}

func (m *EntityProfileDTO) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *EntityProfileDTO) GetVendor() string {
	if m != nil && m.Vendor != nil {
		return *m.Vendor
	}
	return ""
}

func (m *EntityProfileDTO) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *EntityProfileDTO) GetVmProfileDTO() *EntityProfileDTO_VMProfileDTO {
	if x, ok := m.GetEntityTypeSpecificData().(*EntityProfileDTO_VmProfileDTO); ok {
		return x.VmProfileDTO
	}
	return nil
}

func (m *EntityProfileDTO) GetPmProfileDTO() *EntityProfileDTO_PMProfileDTO {
	if x, ok := m.GetEntityTypeSpecificData().(*EntityProfileDTO_PmProfileDTO); ok {
		return x.PmProfileDTO
	}
	return nil
}

func (m *EntityProfileDTO) GetDbProfileDTO() *EntityProfileDTO_DBProfileDTO {
	if x, ok := m.GetEntityTypeSpecificData().(*EntityProfileDTO_DbProfileDTO); ok {
		return x.DbProfileDTO
	}
	return nil
}

func (m *EntityProfileDTO) GetEnableProvisionMatch() bool {
	if m != nil && m.EnableProvisionMatch != nil {
		return *m.EnableProvisionMatch
	}
	return false
}

func (m *EntityProfileDTO) GetEnableResizeMatch() bool {
	if m != nil && m.EnableResizeMatch != nil {
		return *m.EnableResizeMatch
	}
	return false
}

func (m *EntityProfileDTO) GetEntityProperties() []*EntityDTO_EntityProperty {
	if m != nil {
		return m.EntityProperties
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EntityProfileDTO) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EntityProfileDTO_OneofMarshaler, _EntityProfileDTO_OneofUnmarshaler, _EntityProfileDTO_OneofSizer, []interface{}{
		(*EntityProfileDTO_VmProfileDTO)(nil),
		(*EntityProfileDTO_PmProfileDTO)(nil),
		(*EntityProfileDTO_DbProfileDTO)(nil),
	}
}

func _EntityProfileDTO_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EntityProfileDTO)
	// EntityTypeSpecificData
	switch x := m.EntityTypeSpecificData.(type) {
	case *EntityProfileDTO_VmProfileDTO:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VmProfileDTO); err != nil {
			return err
		}
	case *EntityProfileDTO_PmProfileDTO:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PmProfileDTO); err != nil {
			return err
		}
	case *EntityProfileDTO_DbProfileDTO:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DbProfileDTO); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EntityProfileDTO.EntityTypeSpecificData has unexpected type %T", x)
	}
	return nil
}

func _EntityProfileDTO_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EntityProfileDTO)
	switch tag {
	case 8: // EntityTypeSpecificData.vmProfileDTO
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EntityProfileDTO_VMProfileDTO)
		err := b.DecodeMessage(msg)
		m.EntityTypeSpecificData = &EntityProfileDTO_VmProfileDTO{msg}
		return true, err
	case 9: // EntityTypeSpecificData.pmProfileDTO
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EntityProfileDTO_PMProfileDTO)
		err := b.DecodeMessage(msg)
		m.EntityTypeSpecificData = &EntityProfileDTO_PmProfileDTO{msg}
		return true, err
	case 13: // EntityTypeSpecificData.dbProfileDTO
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EntityProfileDTO_DBProfileDTO)
		err := b.DecodeMessage(msg)
		m.EntityTypeSpecificData = &EntityProfileDTO_DbProfileDTO{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EntityProfileDTO_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EntityProfileDTO)
	// EntityTypeSpecificData
	switch x := m.EntityTypeSpecificData.(type) {
	case *EntityProfileDTO_VmProfileDTO:
		s := proto.Size(x.VmProfileDTO)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EntityProfileDTO_PmProfileDTO:
		s := proto.Size(x.PmProfileDTO)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EntityProfileDTO_DbProfileDTO:
		s := proto.Size(x.DbProfileDTO)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Specific data related to a vm profile
type EntityProfileDTO_VMProfileDTO struct {
	// At least one of numVCPUs and vCPUSpeed should be specified.
	// One of the two can be derived from the other using the capacity
	// from the commodityDTO
	// number of VCPUs
	NumVCPUs *int32 `protobuf:"varint,1,opt,name=numVCPUs" json:"numVCPUs,omitempty"`
	// VCPU speed
	VCPUSpeed *float32 `protobuf:"fixed32,2,opt,name=vCPUSpeed" json:"vCPUSpeed,omitempty"`
	// Number of storage entities that this VM will use storage from
	NumStorageConsumed *int32 `protobuf:"varint,3,opt,name=numStorageConsumed" json:"numStorageConsumed,omitempty"`
	// Disk type related to the VM
	DiskType         *string `protobuf:"bytes,4,opt,name=diskType" json:"diskType,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EntityProfileDTO_VMProfileDTO) Reset()         { *m = EntityProfileDTO_VMProfileDTO{} }
func (m *EntityProfileDTO_VMProfileDTO) String() string { return proto.CompactTextString(m) }
func (*EntityProfileDTO_VMProfileDTO) ProtoMessage()    {}
func (*EntityProfileDTO_VMProfileDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 0}
}

func (m *EntityProfileDTO_VMProfileDTO) GetNumVCPUs() int32 {
	if m != nil && m.NumVCPUs != nil {
		return *m.NumVCPUs
	}
	return 0
}

func (m *EntityProfileDTO_VMProfileDTO) GetVCPUSpeed() float32 {
	if m != nil && m.VCPUSpeed != nil {
		return *m.VCPUSpeed
	}
	return 0
}

func (m *EntityProfileDTO_VMProfileDTO) GetNumStorageConsumed() int32 {
	if m != nil && m.NumStorageConsumed != nil {
		return *m.NumStorageConsumed
	}
	return 0
}

func (m *EntityProfileDTO_VMProfileDTO) GetDiskType() string {
	if m != nil && m.DiskType != nil {
		return *m.DiskType
	}
	return ""
}

// Specific data related to a pm profile
type EntityProfileDTO_PMProfileDTO struct {
	// At least one of numCores and cpuCoreSpeed should be specified
	// The other can be derived from the cpu capacity in
	// the commodity dto.
	NumCores         *int32   `protobuf:"varint,1,opt,name=numCores" json:"numCores,omitempty"`
	CpuCoreSpeed     *float32 `protobuf:"fixed32,2,opt,name=cpuCoreSpeed" json:"cpuCoreSpeed,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EntityProfileDTO_PMProfileDTO) Reset()         { *m = EntityProfileDTO_PMProfileDTO{} }
func (m *EntityProfileDTO_PMProfileDTO) String() string { return proto.CompactTextString(m) }
func (*EntityProfileDTO_PMProfileDTO) ProtoMessage()    {}
func (*EntityProfileDTO_PMProfileDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 1}
}

func (m *EntityProfileDTO_PMProfileDTO) GetNumCores() int32 {
	if m != nil && m.NumCores != nil {
		return *m.NumCores
	}
	return 0
}

func (m *EntityProfileDTO_PMProfileDTO) GetCpuCoreSpeed() float32 {
	if m != nil && m.CpuCoreSpeed != nil {
		return *m.CpuCoreSpeed
	}
	return 0
}

// Specific data related to a db profile or db instance profile
// Only used by vendors: AWS and Azure
type EntityProfileDTO_DBProfileDTO struct {
	// region
	Region *string `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	// database code, only used by AWS
	DbCode *int32 `protobuf:"varint,2,opt,name=dbCode" json:"dbCode,omitempty"`
	// database edition
	DbEdition *string `protobuf:"bytes,3,opt,name=dbEdition" json:"dbEdition,omitempty"`
	// database engine
	DbEngine *string `protobuf:"bytes,4,opt,name=dbEngine" json:"dbEngine,omitempty"`
	// deployment option
	DeploymentOption *string `protobuf:"bytes,5,opt,name=deploymentOption" json:"deploymentOption,omitempty"`
	// number of VCPUs
	NumVCPUs         *int32 `protobuf:"varint,6,opt,name=numVCPUs" json:"numVCPUs,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EntityProfileDTO_DBProfileDTO) Reset()         { *m = EntityProfileDTO_DBProfileDTO{} }
func (m *EntityProfileDTO_DBProfileDTO) String() string { return proto.CompactTextString(m) }
func (*EntityProfileDTO_DBProfileDTO) ProtoMessage()    {}
func (*EntityProfileDTO_DBProfileDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 2}
}

func (m *EntityProfileDTO_DBProfileDTO) GetRegion() string {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return ""
}

func (m *EntityProfileDTO_DBProfileDTO) GetDbCode() int32 {
	if m != nil && m.DbCode != nil {
		return *m.DbCode
	}
	return 0
}

func (m *EntityProfileDTO_DBProfileDTO) GetDbEdition() string {
	if m != nil && m.DbEdition != nil {
		return *m.DbEdition
	}
	return ""
}

func (m *EntityProfileDTO_DBProfileDTO) GetDbEngine() string {
	if m != nil && m.DbEngine != nil {
		return *m.DbEngine
	}
	return ""
}

func (m *EntityProfileDTO_DBProfileDTO) GetDeploymentOption() string {
	if m != nil && m.DeploymentOption != nil {
		return *m.DeploymentOption
	}
	return ""
}

func (m *EntityProfileDTO_DBProfileDTO) GetNumVCPUs() int32 {
	if m != nil && m.NumVCPUs != nil {
		return *m.NumVCPUs
	}
	return 0
}

// Data related to a commodity profile used in an entity profile
// Note typically only a subset of these fields may be specified in a profile for
// each commmodity.
type CommodityProfileDTO struct {
	// The type of commodity such as MEM, CPU, STORAGE
	CommodityType *CommodityDTO_CommodityType `protobuf:"varint,1,req,name=commodityType,enum=common_dto.CommodityDTO_CommodityType" json:"commodityType,omitempty"`
	// The capacity of the commodity
	Capacity *float32 `protobuf:"fixed32,2,opt,name=capacity" json:"capacity,omitempty"`
	// Consumed factor may be used to calculate consumed based on capacity
	ConsumedFactor *float32 `protobuf:"fixed32,3,opt,name=consumedFactor" json:"consumedFactor,omitempty"`
	// Consumed value to be used in the profile
	Consumed *float32 `protobuf:"fixed32,4,opt,name=consumed" json:"consumed,omitempty"`
	// A reservation related to this commodity
	Reservation *float32 `protobuf:"fixed32,5,opt,name=reservation" json:"reservation,omitempty"`
	// Overhead related to this commodity - for example overheadMem
	Overhead         *float32 `protobuf:"fixed32,6,opt,name=overhead" json:"overhead,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CommodityProfileDTO) Reset()                    { *m = CommodityProfileDTO{} }
func (m *CommodityProfileDTO) String() string            { return proto.CompactTextString(m) }
func (*CommodityProfileDTO) ProtoMessage()               {}
func (*CommodityProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *CommodityProfileDTO) GetCommodityType() CommodityDTO_CommodityType {
	if m != nil && m.CommodityType != nil {
		return *m.CommodityType
	}
	return CommodityDTO_CLUSTER
}

func (m *CommodityProfileDTO) GetCapacity() float32 {
	if m != nil && m.Capacity != nil {
		return *m.Capacity
	}
	return 0
}

func (m *CommodityProfileDTO) GetConsumedFactor() float32 {
	if m != nil && m.ConsumedFactor != nil {
		return *m.ConsumedFactor
	}
	return 0
}

func (m *CommodityProfileDTO) GetConsumed() float32 {
	if m != nil && m.Consumed != nil {
		return *m.Consumed
	}
	return 0
}

func (m *CommodityProfileDTO) GetReservation() float32 {
	if m != nil && m.Reservation != nil {
		return *m.Reservation
	}
	return 0
}

func (m *CommodityProfileDTO) GetOverhead() float32 {
	if m != nil && m.Overhead != nil {
		return *m.Overhead
	}
	return 0
}

// This represents a deployment profile (service catalog item) which is related
// to a service entity profile (template)
// This DTO ties image information with scope and a profile to allow for
// the deployment of an entity related to a profile
type DeploymentProfileDTO struct {
	// id related to this.  This may be an id from the system where this was discovered
	// it must be unique in VMTurbo
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// Display name for the deployment profile
	ProfileName *string `protobuf:"bytes,2,req,name=profileName" json:"profileName,omitempty"`
	// Context data needed to actually invoke deploy - such as URIs
	ContextData []*ContextData `protobuf:"bytes,3,rep,name=contextData" json:"contextData,omitempty"`
	// related service entity profiles (templates)
	RelatedEntityProfileId []string `protobuf:"bytes,4,rep,name=relatedEntityProfileId" json:"relatedEntityProfileId,omitempty"`
	// scopes in which this can be used for example cluster, network
	RelatedScopeId []string `protobuf:"bytes,5,rep,name=relatedScopeId" json:"relatedScopeId,omitempty"`
	// accessible scopes in which this can be used for example clusters
	// this id allows for a set of clusters, where the relatedScopeId would typically
	// only allow for 1 cluster or data center.  This is treated as an OR of scopes
	AccessibleScopeId []string `protobuf:"bytes,6,rep,name=accessibleScopeId" json:"accessibleScopeId,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *DeploymentProfileDTO) Reset()                    { *m = DeploymentProfileDTO{} }
func (m *DeploymentProfileDTO) String() string            { return proto.CompactTextString(m) }
func (*DeploymentProfileDTO) ProtoMessage()               {}
func (*DeploymentProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *DeploymentProfileDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *DeploymentProfileDTO) GetProfileName() string {
	if m != nil && m.ProfileName != nil {
		return *m.ProfileName
	}
	return ""
}

func (m *DeploymentProfileDTO) GetContextData() []*ContextData {
	if m != nil {
		return m.ContextData
	}
	return nil
}

func (m *DeploymentProfileDTO) GetRelatedEntityProfileId() []string {
	if m != nil {
		return m.RelatedEntityProfileId
	}
	return nil
}

func (m *DeploymentProfileDTO) GetRelatedScopeId() []string {
	if m != nil {
		return m.RelatedScopeId
	}
	return nil
}

func (m *DeploymentProfileDTO) GetAccessibleScopeId() []string {
	if m != nil {
		return m.AccessibleScopeId
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityProfileDTO)(nil), "common_dto.EntityProfileDTO")
	proto.RegisterType((*EntityProfileDTO_VMProfileDTO)(nil), "common_dto.EntityProfileDTO.VMProfileDTO")
	proto.RegisterType((*EntityProfileDTO_PMProfileDTO)(nil), "common_dto.EntityProfileDTO.PMProfileDTO")
	proto.RegisterType((*EntityProfileDTO_DBProfileDTO)(nil), "common_dto.EntityProfileDTO.DBProfileDTO")
	proto.RegisterType((*CommodityProfileDTO)(nil), "common_dto.CommodityProfileDTO")
	proto.RegisterType((*DeploymentProfileDTO)(nil), "common_dto.DeploymentProfileDTO")
}

func init() { proto.RegisterFile("ProfileDTO.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x25, 0xe9, 0x5a, 0xda, 0xdb, 0x6e, 0x74, 0xd9, 0x34, 0xbc, 0x0a, 0xb1, 0x68, 0x42, 0x28,
	0x48, 0xd0, 0x87, 0x09, 0x21, 0xf1, 0x00, 0x12, 0x4d, 0x8b, 0xe0, 0x61, 0xac, 0xa2, 0xdb, 0x5e,
	0x91, 0x6b, 0xdf, 0x6d, 0x16, 0x49, 0x6c, 0x39, 0x6e, 0x45, 0x79, 0xe1, 0xb3, 0xe0, 0x33, 0xf8,
	0x24, 0x64, 0xb7, 0x4d, 0x32, 0x36, 0xd0, 0xde, 0x9a, 0x7b, 0x8e, 0xcf, 0xbd, 0xf7, 0xf8, 0xd4,
	0xd0, 0x1d, 0x6b, 0x79, 0x21, 0x12, 0x1c, 0x9e, 0x9e, 0xf4, 0x95, 0x96, 0x46, 0x06, 0xc0, 0x64,
	0x9a, 0xca, 0xec, 0x0b, 0x37, 0xb2, 0xf7, 0x20, 0x76, 0xbf, 0x0b, 0xf0, 0xf0, 0x77, 0x03, 0xba,
	0xa3, 0xcc, 0x08, 0xb3, 0x28, 0xcf, 0x05, 0x00, 0xbe, 0xe0, 0xc4, 0x0b, 0xfd, 0xa8, 0x15, 0xec,
	0x40, 0x9b, 0x8b, 0x5c, 0x25, 0x74, 0xf1, 0x89, 0xa6, 0x48, 0x7c, 0x57, 0x7c, 0x09, 0x80, 0xee,
	0xd0, 0xe9, 0x42, 0x21, 0xa9, 0x85, 0x7e, 0xb4, 0x75, 0x14, 0xf6, 0xcb, 0x3e, 0xfd, 0xa5, 0xa4,
	0x6d, 0x33, 0x2a, 0x78, 0xc1, 0x6b, 0xe8, 0x3a, 0x0a, 0x2f, 0xbb, 0x91, 0x8d, 0xb0, 0x16, 0xb5,
	0x8f, 0x0e, 0xaa, 0x67, 0xe3, 0xbf, 0x38, 0x76, 0xa2, 0x4d, 0xa8, 0xa7, 0x92, 0x63, 0x42, 0xea,
	0xa1, 0x17, 0xb5, 0x82, 0x2d, 0x68, 0xcc, 0x31, 0xe3, 0x52, 0x93, 0x86, 0xfb, 0xb6, 0x43, 0x62,
	0xce, 0xb4, 0x50, 0x46, 0xc8, 0x8c, 0xdc, 0x77, 0xc5, 0x77, 0xd0, 0x99, 0xa7, 0xa5, 0x06, 0x69,
	0x86, 0x5e, 0xd4, 0x3e, 0x7a, 0x76, 0x73, 0xcc, 0x8a, 0x63, 0xe7, 0xc7, 0xe5, 0xc7, 0x87, 0x7b,
	0x56, 0x42, 0x55, 0x25, 0x5a, 0x77, 0x90, 0x18, 0xdf, 0x90, 0xe0, 0xd3, 0x8a, 0xc4, 0xe6, 0x1d,
	0x24, 0x86, 0x83, 0x6b, 0x12, 0x8f, 0x60, 0x17, 0x33, 0x3a, 0x4d, 0x70, 0xac, 0xe5, 0x5c, 0xe4,
	0x42, 0x66, 0xc7, 0xd4, 0xb0, 0x2b, 0x02, 0xa1, 0x17, 0x35, 0x83, 0x7d, 0xd8, 0x5e, 0xa2, 0x9f,
	0x31, 0x17, 0xdf, 0x71, 0x09, 0xb5, 0x1d, 0xf4, 0x16, 0xba, 0xb8, 0xd6, 0x56, 0xa8, 0x8d, 0xc0,
	0x9c, 0x74, 0x9c, 0xe1, 0x4f, 0xfe, 0x77, 0x59, 0x2b, 0xf6, 0xa2, 0x47, 0xa1, 0x53, 0x35, 0x24,
	0xe8, 0x42, 0x33, 0x9b, 0xa5, 0xe7, 0xf1, 0xf8, 0x2c, 0x27, 0x5e, 0xe8, 0x45, 0xf5, 0x60, 0x1b,
	0x5a, 0xf3, 0x78, 0x7c, 0x36, 0x51, 0x88, 0x9c, 0xf8, 0xa1, 0x17, 0xf9, 0x41, 0x0f, 0x82, 0x6c,
	0x96, 0x4e, 0x8c, 0xd4, 0xf4, 0x12, 0x63, 0x99, 0xe5, 0xb3, 0x14, 0x39, 0xa9, 0x39, 0x7a, 0x17,
	0x9a, 0x5c, 0xe4, 0x5f, 0x5d, 0x6a, 0x36, 0xec, 0x25, 0xf5, 0x5e, 0x41, 0x67, 0x7c, 0xb3, 0x45,
	0x2c, 0x35, 0xae, 0x5b, 0xec, 0x42, 0x87, 0xa9, 0x99, 0xad, 0x54, 0xba, 0xf4, 0x7e, 0x40, 0xa7,
	0xea, 0x92, 0x4d, 0x84, 0xc6, 0x4b, 0x7b, 0xf9, 0xde, 0x3a, 0x21, 0x7c, 0x1a, 0x4b, 0x8e, 0x8e,
	0xef, 0x06, 0xe5, 0xd3, 0x11, 0x17, 0x2e, 0x1f, 0x35, 0x47, 0xb1, 0xc3, 0x4c, 0x47, 0xd9, 0xa5,
	0xc8, 0x56, 0xc3, 0x04, 0x04, 0xba, 0x1c, 0x55, 0x22, 0x17, 0x29, 0x66, 0xe6, 0x64, 0x99, 0xa5,
	0xfa, 0x9a, 0x5b, 0x6c, 0x6e, 0x23, 0x57, 0x1f, 0x10, 0xd8, 0x2b, 0xa3, 0x3d, 0x51, 0xc8, 0xc4,
	0x85, 0x60, 0x43, 0x6a, 0xe8, 0xe1, 0x4f, 0x0f, 0x76, 0x6e, 0xcb, 0xf0, 0x1b, 0xd8, 0x2c, 0xe2,
	0xef, 0x1c, 0xf0, 0xdc, 0xff, 0xe6, 0xe9, 0xad, 0xd9, 0xb7, 0xb7, 0x11, 0x57, 0xd9, 0x76, 0x04,
	0x46, 0x15, 0x65, 0xc2, 0x2c, 0x56, 0x4e, 0xef, 0xc1, 0x16, 0x5b, 0xf9, 0xfb, 0x9e, 0x32, 0x23,
	0xb5, 0x5b, 0xcc, 0x77, 0xcc, 0xb5, 0xef, 0x1b, 0xae, 0xb2, 0x03, 0x6d, 0x8d, 0x39, 0xea, 0x39,
	0x2d, 0x76, 0x72, 0x34, 0x39, 0x47, 0x7d, 0x85, 0x94, 0xbb, 0x9d, 0xfc, 0xc3, 0x5f, 0x1e, 0xec,
	0x0e, 0x0b, 0x03, 0xfe, 0xfd, 0x20, 0xa8, 0x25, 0x52, 0x79, 0x10, 0x9e, 0x43, 0x9b, 0xc9, 0xcc,
	0xe0, 0x37, 0x63, 0x2d, 0x20, 0x35, 0x17, 0xb2, 0x87, 0xd7, 0x37, 0x2b, 0xe0, 0xe0, 0x31, 0xec,
	0x69, 0x4c, 0xa8, 0x41, 0x7e, 0x2d, 0xfa, 0x1f, 0xb9, 0x7b, 0x0e, 0x5a, 0x76, 0xb1, 0x15, 0x3e,
	0x61, 0x52, 0xd9, 0x7a, 0xdd, 0xd5, 0xf7, 0x61, 0x9b, 0x32, 0x86, 0x79, 0x2e, 0xa6, 0x09, 0xae,
	0xa1, 0x86, 0x85, 0x06, 0x2f, 0xe0, 0x80, 0xc9, 0xb4, 0x3f, 0x4f, 0xcd, 0x4c, 0x4f, 0x65, 0x5f,
	0x25, 0xd4, 0x5c, 0x48, 0x9d, 0xae, 0x26, 0xe8, 0x73, 0x23, 0x07, 0x50, 0x2e, 0xf4, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0x73, 0xc5, 0xf9, 0x26, 0x05, 0x00, 0x00,
}
