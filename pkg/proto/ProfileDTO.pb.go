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

type EntityProfileDTO_VMProfileDTO_VMTier int32

const (
	// Basic VM tier: VMs can't participate in Azure Load Balancers
	EntityProfileDTO_VMProfileDTO_BASIC EntityProfileDTO_VMProfileDTO_VMTier = 1
	// Standard VM tier: VMs can be part of Azure Load Balancers backend pools
	EntityProfileDTO_VMProfileDTO_STANDARD EntityProfileDTO_VMProfileDTO_VMTier = 2
)

var EntityProfileDTO_VMProfileDTO_VMTier_name = map[int32]string{
	1: "BASIC",
	2: "STANDARD",
}
var EntityProfileDTO_VMProfileDTO_VMTier_value = map[string]int32{
	"BASIC":    1,
	"STANDARD": 2,
}

func (x EntityProfileDTO_VMProfileDTO_VMTier) Enum() *EntityProfileDTO_VMProfileDTO_VMTier {
	p := new(EntityProfileDTO_VMProfileDTO_VMTier)
	*p = x
	return p
}
func (x EntityProfileDTO_VMProfileDTO_VMTier) String() string {
	return proto.EnumName(EntityProfileDTO_VMProfileDTO_VMTier_name, int32(x))
}
func (x *EntityProfileDTO_VMProfileDTO_VMTier) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EntityProfileDTO_VMProfileDTO_VMTier_value, data, "EntityProfileDTO_VMProfileDTO_VMTier")
	if err != nil {
		return err
	}
	*x = EntityProfileDTO_VMProfileDTO_VMTier(value)
	return nil
}
func (EntityProfileDTO_VMProfileDTO_VMTier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 1, 0}
}

type EntityProfileDTO_VMProfileDTO_InstanceDiskType int32

const (
	EntityProfileDTO_VMProfileDTO_NONE EntityProfileDTO_VMProfileDTO_InstanceDiskType = 1
	// AWS HDD
	EntityProfileDTO_VMProfileDTO_HDD EntityProfileDTO_VMProfileDTO_InstanceDiskType = 2
	// AWS SSD
	EntityProfileDTO_VMProfileDTO_SSD EntityProfileDTO_VMProfileDTO_InstanceDiskType = 3
	// AWS  non-volatile memory express
	EntityProfileDTO_VMProfileDTO_NVME_SSD EntityProfileDTO_VMProfileDTO_InstanceDiskType = 4
	// Azure standard storage
	EntityProfileDTO_VMProfileDTO_AZURE_HDD EntityProfileDTO_VMProfileDTO_InstanceDiskType = 10
	// Azure premium storage
	EntityProfileDTO_VMProfileDTO_AZURE_SSD EntityProfileDTO_VMProfileDTO_InstanceDiskType = 11
)

var EntityProfileDTO_VMProfileDTO_InstanceDiskType_name = map[int32]string{
	1:  "NONE",
	2:  "HDD",
	3:  "SSD",
	4:  "NVME_SSD",
	10: "AZURE_HDD",
	11: "AZURE_SSD",
}
var EntityProfileDTO_VMProfileDTO_InstanceDiskType_value = map[string]int32{
	"NONE":      1,
	"HDD":       2,
	"SSD":       3,
	"NVME_SSD":  4,
	"AZURE_HDD": 10,
	"AZURE_SSD": 11,
}

func (x EntityProfileDTO_VMProfileDTO_InstanceDiskType) Enum() *EntityProfileDTO_VMProfileDTO_InstanceDiskType {
	p := new(EntityProfileDTO_VMProfileDTO_InstanceDiskType)
	*p = x
	return p
}
func (x EntityProfileDTO_VMProfileDTO_InstanceDiskType) String() string {
	return proto.EnumName(EntityProfileDTO_VMProfileDTO_InstanceDiskType_name, int32(x))
}
func (x *EntityProfileDTO_VMProfileDTO_InstanceDiskType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EntityProfileDTO_VMProfileDTO_InstanceDiskType_value, data, "EntityProfileDTO_VMProfileDTO_InstanceDiskType")
	if err != nil {
		return err
	}
	*x = EntityProfileDTO_VMProfileDTO_InstanceDiskType(value)
	return nil
}
func (EntityProfileDTO_VMProfileDTO_InstanceDiskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 1, 1}
}

// AWS VM tenancy types
type EntityProfileDTO_VMProfileDTO_TenancyType int32

const (
	EntityProfileDTO_VMProfileDTO_SHARED    EntityProfileDTO_VMProfileDTO_TenancyType = 1
	EntityProfileDTO_VMProfileDTO_DEDICATED EntityProfileDTO_VMProfileDTO_TenancyType = 2
	EntityProfileDTO_VMProfileDTO_HOST      EntityProfileDTO_VMProfileDTO_TenancyType = 3
)

var EntityProfileDTO_VMProfileDTO_TenancyType_name = map[int32]string{
	1: "SHARED",
	2: "DEDICATED",
	3: "HOST",
}
var EntityProfileDTO_VMProfileDTO_TenancyType_value = map[string]int32{
	"SHARED":    1,
	"DEDICATED": 2,
	"HOST":      3,
}

func (x EntityProfileDTO_VMProfileDTO_TenancyType) Enum() *EntityProfileDTO_VMProfileDTO_TenancyType {
	p := new(EntityProfileDTO_VMProfileDTO_TenancyType)
	*p = x
	return p
}
func (x EntityProfileDTO_VMProfileDTO_TenancyType) String() string {
	return proto.EnumName(EntityProfileDTO_VMProfileDTO_TenancyType_name, int32(x))
}
func (x *EntityProfileDTO_VMProfileDTO_TenancyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EntityProfileDTO_VMProfileDTO_TenancyType_value, data, "EntityProfileDTO_VMProfileDTO_TenancyType")
	if err != nil {
		return err
	}
	*x = EntityProfileDTO_VMProfileDTO_TenancyType(value)
	return nil
}
func (EntityProfileDTO_VMProfileDTO_TenancyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 1, 2}
}

// This file lists all the objects related to Service Entity profiles
// created by user in environment or in VMTurbo
type EntityProfileDTO struct {
	// id of the entity profile.  This should allow access to the object
	// in the system it was discovered from and it should be unique in VMT also.
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// display name to be displayed to the user
	DisplayName *string `protobuf:"bytes,2,opt,name=displayName" json:"displayName,omitempty"`
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
	// Whether the profile is being created/updated or removed
	UpdateType       *UpdateType `protobuf:"varint,14,opt,name=updateType,enum=common_dto.UpdateType,def=0" json:"updateType,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *EntityProfileDTO) Reset()                    { *m = EntityProfileDTO{} }
func (m *EntityProfileDTO) String() string            { return proto.CompactTextString(m) }
func (*EntityProfileDTO) ProtoMessage()               {}
func (*EntityProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

const Default_EntityProfileDTO_UpdateType UpdateType = UpdateType_UPDATED

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

func (m *EntityProfileDTO) GetUpdateType() UpdateType {
	if m != nil && m.UpdateType != nil {
		return *m.UpdateType
	}
	return Default_EntityProfileDTO_UpdateType
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

// define a multimap between region name and available license name in each region.
type EntityProfileDTO_LicenseMapEntry struct {
	Region           *string  `protobuf:"bytes,1,req,name=region" json:"region,omitempty"`
	LicenseName      []string `protobuf:"bytes,2,rep,name=licenseName" json:"licenseName,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EntityProfileDTO_LicenseMapEntry) Reset()         { *m = EntityProfileDTO_LicenseMapEntry{} }
func (m *EntityProfileDTO_LicenseMapEntry) String() string { return proto.CompactTextString(m) }
func (*EntityProfileDTO_LicenseMapEntry) ProtoMessage()    {}
func (*EntityProfileDTO_LicenseMapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 0}
}

func (m *EntityProfileDTO_LicenseMapEntry) GetRegion() string {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return ""
}

func (m *EntityProfileDTO_LicenseMapEntry) GetLicenseName() []string {
	if m != nil {
		return m.LicenseName
	}
	return nil
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
	DiskType *string `protobuf:"bytes,4,opt,name=diskType" json:"diskType,omitempty"`
	// An identifier for matching profiles that belong to the same family.
	Family *string `protobuf:"bytes,5,opt,name=family" json:"family,omitempty"`
	// A quantitative way to compare different instance types in a family.
	NumberOfCoupons *int32 `protobuf:"varint,6,opt,name=numberOfCoupons" json:"numberOfCoupons,omitempty"`
	// Specifies the dedicated storage configuration state for the VM profile
	DedicatedStorageNetworkState *EntityDTO_ComputeTierData_DedicatedStorageNetworkState `protobuf:"varint,7,opt,name=dedicatedStorageNetworkState,enum=common_dto.EntityDTO_ComputeTierData_DedicatedStorageNetworkState" json:"dedicatedStorageNetworkState,omitempty"`
	License                      []*EntityProfileDTO_LicenseMapEntry                     `protobuf:"bytes,8,rep,name=license" json:"license,omitempty"`
	// UUID of the entity from which this entity is cloned.
	// Required for EBS enabled profiles on AWS.
	ClonedUuid *string `protobuf:"bytes,9,opt,name=clonedUuid" json:"clonedUuid,omitempty"`
	// The VM tier. Currently required for Azure.
	VmTier           *EntityProfileDTO_VMProfileDTO_VMTier           `protobuf:"varint,10,opt,name=vmTier,enum=common_dto.EntityProfileDTO_VMProfileDTO_VMTier" json:"vmTier,omitempty"`
	InstanceDiskType *EntityProfileDTO_VMProfileDTO_InstanceDiskType `protobuf:"varint,11,opt,name=instanceDiskType,enum=common_dto.EntityProfileDTO_VMProfileDTO_InstanceDiskType" json:"instanceDiskType,omitempty"`
	InstanceDiskSize *int32                                          `protobuf:"varint,12,opt,name=instanceDiskSize" json:"instanceDiskSize,omitempty"`
	// Show if it supports AWS encrypted volumes
	SupportEncryptedVolume *bool `protobuf:"varint,13,opt,name=supportEncryptedVolume" json:"supportEncryptedVolume,omitempty"`
	// Template providers are created on the platform side, so on the probe side we can only
	// find out a key for template family commodity and then use it for creating of commodity.
	TemplateFamilyCommodityKey *string `protobuf:"bytes,14,opt,name=templateFamilyCommodityKey" json:"templateFamilyCommodityKey,omitempty"`
	NumCloudComputeUnit        *int32  `protobuf:"varint,15,opt,name=numCloudComputeUnit" json:"numCloudComputeUnit,omitempty"`
	// instanceSizeFamily is the group that comprises of profiles of varying sizes that fit a
	// particular use case. Instance size flexible Reserved Instances can be applied to entities
	// using any profile within an instanceSizeFamily.
	InstanceSizeFamily *string `protobuf:"bytes,16,opt,name=instanceSizeFamily" json:"instanceSizeFamily,omitempty"`
	// An optional list of tenancy types supported by the VM profile
	Tenancy []EntityProfileDTO_VMProfileDTO_TenancyType `protobuf:"varint,17,rep,name=tenancy,enum=common_dto.EntityProfileDTO_VMProfileDTO_TenancyType" json:"tenancy,omitempty"`
	// quotaFamily is the group that comprises of profiles that count towards the usage limit for
	// a defined category imposed on a Business Account.
	QuotaFamily      *string `protobuf:"bytes,18,opt,name=quotaFamily" json:"quotaFamily,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EntityProfileDTO_VMProfileDTO) Reset()         { *m = EntityProfileDTO_VMProfileDTO{} }
func (m *EntityProfileDTO_VMProfileDTO) String() string { return proto.CompactTextString(m) }
func (*EntityProfileDTO_VMProfileDTO) ProtoMessage()    {}
func (*EntityProfileDTO_VMProfileDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 1}
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

func (m *EntityProfileDTO_VMProfileDTO) GetFamily() string {
	if m != nil && m.Family != nil {
		return *m.Family
	}
	return ""
}

func (m *EntityProfileDTO_VMProfileDTO) GetNumberOfCoupons() int32 {
	if m != nil && m.NumberOfCoupons != nil {
		return *m.NumberOfCoupons
	}
	return 0
}

func (m *EntityProfileDTO_VMProfileDTO) GetDedicatedStorageNetworkState() EntityDTO_ComputeTierData_DedicatedStorageNetworkState {
	if m != nil && m.DedicatedStorageNetworkState != nil {
		return *m.DedicatedStorageNetworkState
	}
	return EntityDTO_ComputeTierData_NOT_SUPPORTED
}

func (m *EntityProfileDTO_VMProfileDTO) GetLicense() []*EntityProfileDTO_LicenseMapEntry {
	if m != nil {
		return m.License
	}
	return nil
}

func (m *EntityProfileDTO_VMProfileDTO) GetClonedUuid() string {
	if m != nil && m.ClonedUuid != nil {
		return *m.ClonedUuid
	}
	return ""
}

func (m *EntityProfileDTO_VMProfileDTO) GetVmTier() EntityProfileDTO_VMProfileDTO_VMTier {
	if m != nil && m.VmTier != nil {
		return *m.VmTier
	}
	return EntityProfileDTO_VMProfileDTO_BASIC
}

func (m *EntityProfileDTO_VMProfileDTO) GetInstanceDiskType() EntityProfileDTO_VMProfileDTO_InstanceDiskType {
	if m != nil && m.InstanceDiskType != nil {
		return *m.InstanceDiskType
	}
	return EntityProfileDTO_VMProfileDTO_NONE
}

func (m *EntityProfileDTO_VMProfileDTO) GetInstanceDiskSize() int32 {
	if m != nil && m.InstanceDiskSize != nil {
		return *m.InstanceDiskSize
	}
	return 0
}

func (m *EntityProfileDTO_VMProfileDTO) GetSupportEncryptedVolume() bool {
	if m != nil && m.SupportEncryptedVolume != nil {
		return *m.SupportEncryptedVolume
	}
	return false
}

func (m *EntityProfileDTO_VMProfileDTO) GetTemplateFamilyCommodityKey() string {
	if m != nil && m.TemplateFamilyCommodityKey != nil {
		return *m.TemplateFamilyCommodityKey
	}
	return ""
}

func (m *EntityProfileDTO_VMProfileDTO) GetNumCloudComputeUnit() int32 {
	if m != nil && m.NumCloudComputeUnit != nil {
		return *m.NumCloudComputeUnit
	}
	return 0
}

func (m *EntityProfileDTO_VMProfileDTO) GetInstanceSizeFamily() string {
	if m != nil && m.InstanceSizeFamily != nil {
		return *m.InstanceSizeFamily
	}
	return ""
}

func (m *EntityProfileDTO_VMProfileDTO) GetTenancy() []EntityProfileDTO_VMProfileDTO_TenancyType {
	if m != nil {
		return m.Tenancy
	}
	return nil
}

func (m *EntityProfileDTO_VMProfileDTO) GetQuotaFamily() string {
	if m != nil && m.QuotaFamily != nil {
		return *m.QuotaFamily
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
	return fileDescriptor9, []int{0, 2}
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
	// An identifier for matching profiles that belong to the same family.
	Family *string `protobuf:"bytes,1,opt,name=family" json:"family,omitempty"`
	// number of VCPUs
	NumVCPUs *int32 `protobuf:"varint,2,opt,name=numVCPUs" json:"numVCPUs,omitempty"`
	// all database editions supported by profile.  Edition example: "Enterprise" or "web".
	DbEdition []string `protobuf:"bytes,3,rep,name=dbEdition" json:"dbEdition,omitempty"`
	// all database engines supported by profile.  Engine example: "MariaDB" or "MySql".
	DbEngine []string `protobuf:"bytes,4,rep,name=dbEngine" json:"dbEngine,omitempty"`
	// all the license (operating system or DB engine) supported by profile
	License          []*EntityProfileDTO_LicenseMapEntry `protobuf:"bytes,5,rep,name=license" json:"license,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *EntityProfileDTO_DBProfileDTO) Reset()         { *m = EntityProfileDTO_DBProfileDTO{} }
func (m *EntityProfileDTO_DBProfileDTO) String() string { return proto.CompactTextString(m) }
func (*EntityProfileDTO_DBProfileDTO) ProtoMessage()    {}
func (*EntityProfileDTO_DBProfileDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0, 3}
}

func (m *EntityProfileDTO_DBProfileDTO) GetFamily() string {
	if m != nil && m.Family != nil {
		return *m.Family
	}
	return ""
}

func (m *EntityProfileDTO_DBProfileDTO) GetNumVCPUs() int32 {
	if m != nil && m.NumVCPUs != nil {
		return *m.NumVCPUs
	}
	return 0
}

func (m *EntityProfileDTO_DBProfileDTO) GetDbEdition() []string {
	if m != nil {
		return m.DbEdition
	}
	return nil
}

func (m *EntityProfileDTO_DBProfileDTO) GetDbEngine() []string {
	if m != nil {
		return m.DbEngine
	}
	return nil
}

func (m *EntityProfileDTO_DBProfileDTO) GetLicense() []*EntityProfileDTO_LicenseMapEntry {
	if m != nil {
		return m.License
	}
	return nil
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
func (*CommodityProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

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
	ProfileName *string `protobuf:"bytes,2,opt,name=profileName" json:"profileName,omitempty"`
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
	// Whether the profile is being created/updated or removed
	UpdateType       *UpdateType `protobuf:"varint,7,opt,name=updateType,enum=common_dto.UpdateType,def=0" json:"updateType,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *DeploymentProfileDTO) Reset()                    { *m = DeploymentProfileDTO{} }
func (m *DeploymentProfileDTO) String() string            { return proto.CompactTextString(m) }
func (*DeploymentProfileDTO) ProtoMessage()               {}
func (*DeploymentProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

const Default_DeploymentProfileDTO_UpdateType UpdateType = UpdateType_UPDATED

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

func (m *DeploymentProfileDTO) GetUpdateType() UpdateType {
	if m != nil && m.UpdateType != nil {
		return *m.UpdateType
	}
	return Default_DeploymentProfileDTO_UpdateType
}

func init() {
	proto.RegisterType((*EntityProfileDTO)(nil), "common_dto.EntityProfileDTO")
	proto.RegisterType((*EntityProfileDTO_LicenseMapEntry)(nil), "common_dto.EntityProfileDTO.LicenseMapEntry")
	proto.RegisterType((*EntityProfileDTO_VMProfileDTO)(nil), "common_dto.EntityProfileDTO.VMProfileDTO")
	proto.RegisterType((*EntityProfileDTO_PMProfileDTO)(nil), "common_dto.EntityProfileDTO.PMProfileDTO")
	proto.RegisterType((*EntityProfileDTO_DBProfileDTO)(nil), "common_dto.EntityProfileDTO.DBProfileDTO")
	proto.RegisterType((*CommodityProfileDTO)(nil), "common_dto.CommodityProfileDTO")
	proto.RegisterType((*DeploymentProfileDTO)(nil), "common_dto.DeploymentProfileDTO")
	proto.RegisterEnum("common_dto.EntityProfileDTO_VMProfileDTO_VMTier", EntityProfileDTO_VMProfileDTO_VMTier_name, EntityProfileDTO_VMProfileDTO_VMTier_value)
	proto.RegisterEnum("common_dto.EntityProfileDTO_VMProfileDTO_InstanceDiskType", EntityProfileDTO_VMProfileDTO_InstanceDiskType_name, EntityProfileDTO_VMProfileDTO_InstanceDiskType_value)
	proto.RegisterEnum("common_dto.EntityProfileDTO_VMProfileDTO_TenancyType", EntityProfileDTO_VMProfileDTO_TenancyType_name, EntityProfileDTO_VMProfileDTO_TenancyType_value)
}

func init() { proto.RegisterFile("ProfileDTO.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 1221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdf, 0x8e, 0xda, 0x46,
	0x17, 0xff, 0x6c, 0x16, 0x16, 0x0e, 0x84, 0x75, 0x26, 0xd1, 0x7e, 0x16, 0x8a, 0xbe, 0xf0, 0xa1,
	0x2a, 0xa2, 0x55, 0x8a, 0x22, 0xa4, 0x56, 0x6a, 0x2e, 0xda, 0x2e, 0x98, 0x68, 0x57, 0xc9, 0x02,
	0x1a, 0x60, 0x2f, 0x72, 0x13, 0x19, 0xfb, 0x6c, 0x32, 0x0a, 0x9e, 0x71, 0xed, 0x31, 0x2d, 0xb9,
	0xeb, 0x4d, 0xdf, 0xa5, 0x0f, 0xd0, 0x37, 0xe8, 0x13, 0xf4, 0x89, 0xaa, 0x19, 0x1b, 0x30, 0x7f,
	0xb2, 0xda, 0xaa, 0x77, 0x3e, 0xff, 0x7e, 0x67, 0x7c, 0xe6, 0x9c, 0xdf, 0x1c, 0xb0, 0xc6, 0x91,
	0xb8, 0x65, 0x0b, 0x74, 0xa6, 0xa3, 0x4e, 0x18, 0x09, 0x29, 0x08, 0x78, 0x22, 0x08, 0x04, 0x7f,
	0xe7, 0x4b, 0xd1, 0x38, 0xeb, 0xeb, 0xef, 0x8d, 0xb1, 0xf5, 0x07, 0x01, 0x6b, 0xc0, 0x25, 0x93,
	0xab, 0x6d, 0x1c, 0xa9, 0x83, 0xc9, 0x7c, 0xdb, 0x68, 0x9a, 0xed, 0x0a, 0x35, 0x99, 0x4f, 0x9a,
	0x50, 0xf5, 0x59, 0x1c, 0x2e, 0xdc, 0xd5, 0xd0, 0x0d, 0xd0, 0x36, 0x9b, 0x46, 0xbb, 0x42, 0xf3,
	0x2a, 0xf2, 0x23, 0x00, 0x6a, 0x94, 0xe9, 0x2a, 0x44, 0xbb, 0xd0, 0x34, 0xdb, 0xf5, 0x6e, 0xb3,
	0xb3, 0x4d, 0xdc, 0x49, 0x73, 0xa8, 0xbc, 0x83, 0x8d, 0x1f, 0xcd, 0xc5, 0x90, 0xd7, 0x60, 0x69,
	0x77, 0x7f, 0x7b, 0x14, 0xfb, 0xa4, 0x59, 0x68, 0x57, 0xbb, 0x4f, 0xf3, 0x38, 0xfd, 0x3d, 0x1f,
	0x67, 0x3a, 0xa2, 0x07, 0x81, 0xe4, 0x31, 0x14, 0x03, 0xe1, 0xe3, 0xc2, 0x2e, 0xea, 0xa3, 0xa6,
	0x02, 0x39, 0x87, 0xd2, 0x12, 0xb9, 0x2f, 0x22, 0xbb, 0xa4, 0xd5, 0x99, 0xa4, 0x7f, 0x0f, 0x63,
	0x2f, 0x62, 0xa1, 0x64, 0x82, 0xdb, 0xa7, 0xd9, 0xef, 0x6d, 0x55, 0x64, 0x04, 0xb5, 0x65, 0xb0,
	0xcd, 0x68, 0x97, 0x9b, 0x46, 0xbb, 0xda, 0xfd, 0xf2, 0xf0, 0x07, 0x73, 0xc5, 0xbf, 0xb9, 0xde,
	0x0a, 0x97, 0xff, 0xa1, 0x3b, 0x00, 0x0a, 0x30, 0xcc, 0x03, 0x56, 0xee, 0x01, 0x38, 0xde, 0x03,
	0x0c, 0xf7, 0x00, 0xfd, 0x79, 0x0e, 0xf0, 0xc1, 0x3d, 0x00, 0x9d, 0xde, 0x2e, 0x60, 0x1e, 0x80,
	0x74, 0xe1, 0x31, 0x72, 0x77, 0xbe, 0xc0, 0x71, 0x24, 0x96, 0x2c, 0x66, 0x82, 0x5f, 0xbb, 0xd2,
	0xfb, 0x60, 0x43, 0xd3, 0x68, 0x97, 0xe9, 0x51, 0x1b, 0x79, 0x0e, 0x0f, 0x53, 0x3d, 0xc5, 0x98,
	0x7d, 0xc2, 0x34, 0xa0, 0xaa, 0x03, 0x0e, 0x0d, 0x64, 0x0c, 0x16, 0xae, 0x8f, 0x14, 0x62, 0x24,
	0x19, 0xc6, 0x76, 0x4d, 0xdf, 0xf8, 0x17, 0x77, 0x75, 0x4e, 0xe6, 0xbd, 0xa2, 0x07, 0xd1, 0xe4,
	0x07, 0x80, 0x24, 0xf4, 0x5d, 0x89, 0xba, 0x0b, 0xeb, 0x4d, 0xa3, 0x5d, 0xef, 0x9e, 0xe7, 0xb1,
	0x66, 0x1b, 0xeb, 0xcb, 0xd3, 0xd9, 0xd8, 0xb9, 0x98, 0x0e, 0x1c, 0x9a, 0x0b, 0x69, 0xbc, 0x86,
	0xb3, 0x37, 0xcc, 0x43, 0x1e, 0xe3, 0xb5, 0x1b, 0x0e, 0xb8, 0x8c, 0x56, 0xaa, 0x69, 0x22, 0x7c,
	0xaf, 0xfa, 0x22, 0x9d, 0x87, 0x4c, 0x52, 0x4d, 0xb3, 0x48, 0x5d, 0xb3, 0x99, 0x28, 0xa8, 0xa6,
	0xc9, 0xa9, 0x1a, 0xbf, 0x57, 0xa0, 0x96, 0x6f, 0x02, 0xd2, 0x80, 0x32, 0x4f, 0x82, 0x9b, 0xfe,
	0x78, 0x16, 0xdb, 0x46, 0xd3, 0x68, 0x17, 0xe9, 0x46, 0x26, 0x4f, 0xa0, 0xb2, 0xec, 0x8f, 0x67,
	0x93, 0x10, 0xd1, 0xd7, 0x03, 0x66, 0xd2, 0xad, 0x82, 0x74, 0x80, 0xf0, 0x24, 0x98, 0x48, 0x11,
	0xb9, 0xef, 0xb1, 0x2f, 0x78, 0x9c, 0x04, 0xe8, 0xdb, 0x05, 0x8d, 0x71, 0xc4, 0xa2, 0x32, 0xf9,
	0x2c, 0xfe, 0xa8, 0xcb, 0x70, 0xa2, 0xdb, 0x79, 0x23, 0x93, 0x06, 0x94, 0x6e, 0xdd, 0x80, 0x2d,
	0x56, 0xe9, 0x70, 0xf4, 0x4c, 0xdb, 0xa0, 0x99, 0x86, 0xb4, 0xe1, 0x8c, 0x27, 0xc1, 0x1c, 0xa3,
	0xd1, 0x6d, 0x5f, 0x24, 0xa1, 0xe0, 0xb1, 0x1e, 0x95, 0x22, 0xdd, 0x57, 0x93, 0xdf, 0x0c, 0x78,
	0xe2, 0xa3, 0xcf, 0x3c, 0x57, 0xa2, 0x9f, 0xa5, 0x1f, 0xa2, 0xfc, 0x59, 0x44, 0x1f, 0x27, 0xd2,
	0x95, 0xa8, 0xa7, 0xa8, 0xde, 0xed, 0x1d, 0xbf, 0xc9, 0xbe, 0x08, 0xc2, 0x44, 0xe2, 0x94, 0x61,
	0xe4, 0xb8, 0xd2, 0xed, 0x38, 0x77, 0x20, 0xd1, 0x3b, 0xf3, 0x90, 0x57, 0x70, 0x9a, 0x15, 0xdd,
	0x2e, 0xeb, 0xe6, 0x79, 0x7e, 0x67, 0xcf, 0xef, 0x5d, 0x2f, 0x5d, 0x07, 0x93, 0xff, 0x01, 0x78,
	0x0b, 0xc1, 0xd1, 0x9f, 0x25, 0xcc, 0xd7, 0xf3, 0x58, 0xa1, 0x39, 0x0d, 0xb9, 0x84, 0xd2, 0x32,
	0x50, 0x07, 0xd6, 0x13, 0x50, 0xef, 0xbe, 0xb8, 0xf7, 0xf0, 0x77, 0x6e, 0xae, 0x55, 0x1c, 0xcd,
	0xe2, 0xc9, 0x2d, 0x58, 0x8c, 0xc7, 0xd2, 0xe5, 0x1e, 0x3a, 0xeb, 0x4b, 0xaa, 0x6a, 0xcc, 0x97,
	0xf7, 0xc7, 0xbc, 0xda, 0x43, 0xa0, 0x07, 0x98, 0xe4, 0xab, 0xdd, 0x3c, 0x13, 0xf6, 0x09, 0xed,
	0x9a, 0xbe, 0xcd, 0x03, 0x3d, 0xf9, 0x16, 0xce, 0xe3, 0x24, 0x0c, 0x45, 0x24, 0x07, 0xdc, 0x8b,
	0x56, 0xa1, 0x44, 0xff, 0x46, 0x2c, 0x92, 0x00, 0x35, 0x91, 0x94, 0xe9, 0x67, 0xac, 0xe4, 0x7b,
	0x68, 0x48, 0x0c, 0xc2, 0x85, 0x2b, 0xf1, 0x95, 0x6e, 0xa1, 0x0d, 0x3f, 0xbf, 0xc6, 0x95, 0x9e,
	0xc0, 0x0a, 0xbd, 0xc3, 0x83, 0xbc, 0x80, 0x47, 0x3c, 0x09, 0xfa, 0x0b, 0x91, 0xf8, 0x59, 0x4f,
	0xcc, 0x38, 0x93, 0xf6, 0x99, 0x3e, 0xe6, 0x31, 0x93, 0x1a, 0x85, 0xf5, 0xe9, 0xd5, 0xc9, 0x53,
	0x4c, 0xdb, 0xd2, 0x99, 0x8e, 0x58, 0xc8, 0x08, 0x4e, 0x25, 0x72, 0x97, 0x7b, 0x2b, 0xfb, 0x61,
	0xb3, 0xd0, 0xae, 0x77, 0xbf, 0xb9, 0x7f, 0x91, 0xa7, 0x69, 0xa0, 0xae, 0xef, 0x1a, 0x45, 0x0d,
	0xfe, 0x4f, 0x89, 0x90, 0x6e, 0x96, 0x99, 0xa4, 0xaf, 0x45, 0x4e, 0xd5, 0xfa, 0x3f, 0x94, 0xd2,
	0x2b, 0x27, 0x15, 0x28, 0xf6, 0x2e, 0x26, 0x57, 0x7d, 0xcb, 0x20, 0x35, 0x28, 0x4f, 0xa6, 0x17,
	0x43, 0xe7, 0x82, 0x3a, 0x96, 0xd9, 0x7a, 0x0b, 0xd6, 0xfe, 0x0d, 0x92, 0x32, 0x9c, 0x0c, 0x47,
	0xc3, 0x81, 0x65, 0x90, 0x53, 0x28, 0x5c, 0x3a, 0x8e, 0x65, 0xaa, 0x8f, 0xc9, 0xc4, 0xb1, 0x0a,
	0x2a, 0x7a, 0x78, 0x73, 0x3d, 0x78, 0xa7, 0xa4, 0x13, 0xf2, 0x00, 0x2a, 0x17, 0x6f, 0x67, 0x74,
	0xf0, 0x4e, 0x79, 0xc1, 0x56, 0x54, 0xd6, 0x6a, 0xab, 0x0b, 0xd5, 0xdc, 0xc1, 0x09, 0x40, 0x69,
	0x72, 0x79, 0x41, 0x07, 0x8e, 0x65, 0x28, 0x4f, 0x67, 0xe0, 0x5c, 0xf5, 0x15, 0xf1, 0x59, 0xa6,
	0xca, 0x78, 0x39, 0x9a, 0x4c, 0xad, 0x42, 0x63, 0x08, 0xb5, 0xf1, 0x21, 0x55, 0xf5, 0x45, 0x84,
	0x79, 0xaa, 0xd2, 0x32, 0x69, 0x41, 0xcd, 0x0b, 0x13, 0xf5, 0x9d, 0x67, 0xab, 0x1d, 0x5d, 0xe3,
	0x4f, 0x03, 0x6a, 0xf9, 0xe7, 0x45, 0xd1, 0x68, 0xc6, 0x3a, 0x46, 0xfa, 0xf6, 0x66, 0x8c, 0x93,
	0xe7, 0x44, 0xf3, 0x90, 0x13, 0xfd, 0xf9, 0xc0, 0x67, 0xfa, 0x55, 0x2e, 0x68, 0x82, 0xdd, 0x2a,
	0x34, 0xc7, 0xcd, 0x07, 0xfc, 0x3d, 0xe3, 0xe9, 0xa2, 0xa0, 0x38, 0x2e, 0x93, 0xf3, 0xa4, 0x50,
	0xfc, 0x17, 0xa4, 0xd0, 0xb3, 0xe1, 0x7c, 0xbb, 0xae, 0x4c, 0x42, 0xf4, 0xd8, 0x2d, 0xf3, 0x14,
	0x63, 0xb5, 0x7e, 0x35, 0xe1, 0xd1, 0x91, 0x5d, 0x84, 0xbc, 0x81, 0x07, 0x9b, 0x6d, 0x44, 0x4f,
	0xb6, 0xa1, 0x77, 0xa1, 0x67, 0x47, 0x77, 0x98, 0x8c, 0x0a, 0xb7, 0xde, 0x74, 0x37, 0x58, 0xfd,
	0xa3, 0xe7, 0x86, 0xae, 0xc7, 0xe4, 0x2a, 0x2b, 0xf3, 0x46, 0x26, 0xcf, 0xa0, 0xee, 0x65, 0x7c,
	0xff, 0xca, 0xf5, 0xa4, 0x88, 0xf4, 0x7b, 0x60, 0xd2, 0x3d, 0xad, 0xc6, 0x58, 0xbf, 0x18, 0x27,
	0x19, 0xc6, 0xfa, 0x9d, 0x68, 0x42, 0x35, 0xc2, 0x18, 0xa3, 0xa5, 0xab, 0x6b, 0x5c, 0xd4, 0xe6,
	0xbc, 0x4a, 0x45, 0x8b, 0x25, 0x46, 0x1f, 0xd0, 0xf5, 0xf5, 0x53, 0x60, 0xd2, 0x8d, 0xdc, 0xfa,
	0xcb, 0x84, 0xc7, 0x0e, 0x86, 0x0b, 0xb1, 0x0a, 0x90, 0xcb, 0xbb, 0xf7, 0xc7, 0x30, 0xb5, 0xe6,
	0xf7, 0xc7, 0x9c, 0x8a, 0x7c, 0x07, 0x55, 0x4f, 0x70, 0x89, 0xbf, 0x48, 0x55, 0x5d, 0x7d, 0xd9,
	0xd5, 0xee, 0x7f, 0x77, 0x8b, 0xb6, 0x31, 0xd3, 0xbc, 0xaf, 0xa2, 0xae, 0x08, 0x15, 0xbd, 0xf8,
	0x3b, 0xf7, 0x7a, 0xe5, 0x67, 0x5d, 0xf1, 0x19, 0xab, 0xaa, 0x5f, 0x66, 0x99, 0x78, 0x22, 0x54,
	0xfe, 0x45, 0xed, 0xbf, 0xa7, 0x55, 0x4b, 0x8d, 0xeb, 0x79, 0x18, 0xc7, 0x6c, 0xbe, 0xc0, 0xb5,
	0x6b, 0x49, 0xbb, 0x1e, 0x1a, 0xf6, 0x56, 0x90, 0xd3, 0x7f, 0xbc, 0x82, 0xf4, 0xbe, 0x86, 0xa7,
	0x9e, 0x08, 0x3a, 0xcb, 0x40, 0x26, 0xd1, 0x5c, 0x74, 0x14, 0x71, 0xde, 0x8a, 0x28, 0xc8, 0x20,
	0x3a, 0xbe, 0x14, 0x3d, 0xd8, 0x96, 0xfa, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x99, 0x41,
	0x69, 0xef, 0x0b, 0x00, 0x00,
}
