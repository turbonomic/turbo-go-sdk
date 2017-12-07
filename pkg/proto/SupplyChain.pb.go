// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SupplyChain.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TemplateDTO_TemplateType int32

const (
	TemplateDTO_BASE      TemplateDTO_TemplateType = 0
	TemplateDTO_EXTENSION TemplateDTO_TemplateType = 1
)

var TemplateDTO_TemplateType_name = map[int32]string{
	0: "BASE",
	1: "EXTENSION",
}
var TemplateDTO_TemplateType_value = map[string]int32{
	"BASE":      0,
	"EXTENSION": 1,
}

func (x TemplateDTO_TemplateType) Enum() *TemplateDTO_TemplateType {
	p := new(TemplateDTO_TemplateType)
	*p = x
	return p
}
func (x TemplateDTO_TemplateType) String() string {
	return proto1.EnumName(TemplateDTO_TemplateType_name, int32(x))
}
func (x *TemplateDTO_TemplateType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(TemplateDTO_TemplateType_value, data, "TemplateDTO_TemplateType")
	if err != nil {
		return err
	}
	*x = TemplateDTO_TemplateType(value)
	return nil
}
func (TemplateDTO_TemplateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

type Provider_ProviderType int32

const (
	// HOSTING is a To One relationship toward the provider, and it enforces containment.
	// This means that if the provider is removed, then every contained consumer will also be removed.
	Provider_HOSTING Provider_ProviderType = 0
	// LAYERED_OVER is a To Many relationship toward the provider, without containment.
	Provider_LAYERED_OVER Provider_ProviderType = 1
)

var Provider_ProviderType_name = map[int32]string{
	0: "HOSTING",
	1: "LAYERED_OVER",
}
var Provider_ProviderType_value = map[string]int32{
	"HOSTING":      0,
	"LAYERED_OVER": 1,
}

func (x Provider_ProviderType) Enum() *Provider_ProviderType {
	p := new(Provider_ProviderType)
	*p = x
	return p
}
func (x Provider_ProviderType) String() string {
	return proto1.EnumName(Provider_ProviderType_name, int32(x))
}
func (x *Provider_ProviderType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Provider_ProviderType_value, data, "Provider_ProviderType")
	if err != nil {
		return err
	}
	*x = Provider_ProviderType(value)
	return nil
}
func (Provider_ProviderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{2, 0} }

//
// The TemplateDTO message represents entity types (templates) that the probe expects to
// discover in the target. For the probe to load in Operations Manager, it must discover
// entity types that are valid members of the supply chain, and these entities must have
// valid buy/sell relationships. Specifying the set of templates for a probe serves to
// validate that the specific entities the probe discovers and sends to Operations Manager do
// indeed match the entity descriptions the probe is expected to discover.
//
// Specify entity type by setting an EntityType value to the templateClass field.
//
// An entity can maintain a list of commodities that it sells.
//
// An entity can maintain a map of commodities bought (TemplateCommodity objects). Each map key is
// an instance of Provider. For each provider, the map entry is a list of the commodities bought
// from that provider.
//
// The templateType can be either {@code Base} or
// Extension (see TemplateType).
//
// A Base template indicates the initial representation
// of an entity, which means this probe performs the primary discovery of the entity and places it in the market.
// Note that there can be more than one probe that discovers the same Base entity. The template has a
// templatePriority setting that resolves such a collision. The template with the highest priority value
// wins, and discoveries made for the lower-priority template are ignored.
//
// An extension template adds data to already discovered entities. This is a way to extend the
// commodities managed by a base template.
//
type TemplateDTO struct {
	// The type of entity that the template represents. See EntityType
	// for the available types.
	TemplateClass *EntityDTO_EntityType `protobuf:"varint,1,req,name=templateClass,enum=proto.EntityDTO_EntityType" json:"templateClass,omitempty"`
	// The template type (Base or Extension), used during the validation process.
	TemplateType *TemplateDTO_TemplateType `protobuf:"varint,2,req,name=templateType,enum=proto.TemplateDTO_TemplateType" json:"templateType,omitempty"`
	// The priority of a Base template. For equivalent Base templates, Operations Manager uses the highest-priority
	// template, and discards discovered data from lower-priority Base templates.
	TemplatePriority *int32 `protobuf:"varint,3,req,name=templatePriority" json:"templatePriority,omitempty"`
	// This entity's list of {@link TemplateCommodity} items that it provides.
	CommoditySold []*TemplateCommodity `protobuf:"bytes,5,rep,name=commoditySold" json:"commoditySold,omitempty"`
	// The commodities bought from the different providers.
	// This Map contains the commodities bought where:
	CommodityBought []*TemplateDTO_CommBoughtProviderProp `protobuf:"bytes,6,rep,name=commodityBought" json:"commodityBought,omitempty"`
	// A map that defines the entity types that will be providers or consumers for this template entity.
	// The entry key is an entity type, from the EntityType enumeration. There can only be
	// one instance of each entity type in this map. The entry value is an instance of
	// ExternalEntityLink. Each entity link describes an entity type in the supply chain,
	// and the commodities it buys from or sells to the template entity.
	ExternalLink []*TemplateDTO_ExternalEntityLinkProp `protobuf:"bytes,7,rep,name=externalLink" json:"externalLink,omitempty"`
	// Each set represents a case where the entity must buy one commodity of the set ( a logical or of the set)
	// Note, the entity may buy more than one of the commodities in the set.
	CommBoughtOrSet  []*TemplateDTO_CommBoughtProviderOrSet `protobuf:"bytes,8,rep,name=commBoughtOrSet" json:"commBoughtOrSet,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *TemplateDTO) Reset()                    { *m = TemplateDTO{} }
func (m *TemplateDTO) String() string            { return proto1.CompactTextString(m) }
func (*TemplateDTO) ProtoMessage()               {}
func (*TemplateDTO) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *TemplateDTO) GetTemplateClass() EntityDTO_EntityType {
	if m != nil && m.TemplateClass != nil {
		return *m.TemplateClass
	}
	return EntityDTO_SWITCH
}

func (m *TemplateDTO) GetTemplateType() TemplateDTO_TemplateType {
	if m != nil && m.TemplateType != nil {
		return *m.TemplateType
	}
	return TemplateDTO_BASE
}

func (m *TemplateDTO) GetTemplatePriority() int32 {
	if m != nil && m.TemplatePriority != nil {
		return *m.TemplatePriority
	}
	return 0
}

func (m *TemplateDTO) GetCommoditySold() []*TemplateCommodity {
	if m != nil {
		return m.CommoditySold
	}
	return nil
}

func (m *TemplateDTO) GetCommodityBought() []*TemplateDTO_CommBoughtProviderProp {
	if m != nil {
		return m.CommodityBought
	}
	return nil
}

func (m *TemplateDTO) GetExternalLink() []*TemplateDTO_ExternalEntityLinkProp {
	if m != nil {
		return m.ExternalLink
	}
	return nil
}

func (m *TemplateDTO) GetCommBoughtOrSet() []*TemplateDTO_CommBoughtProviderOrSet {
	if m != nil {
		return m.CommBoughtOrSet
	}
	return nil
}

// In some cases, an entity may buy one commodity or another, but it must buy one of the two
// This set represents the set of commodities where the entity must buy one of these.
// It could be that the set contains multiple commodities from the same provider - where only
// one of these will be bought.  Or it could be that there are multiple provider types and the
// entity must buy one.  However, for this set, the entity is only required to buy one of the
// commodities.
type TemplateDTO_CommBoughtProviderOrSet struct {
	CommBought       []*TemplateDTO_CommBoughtProviderProp `protobuf:"bytes,1,rep,name=commBought" json:"commBought,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *TemplateDTO_CommBoughtProviderOrSet) Reset()         { *m = TemplateDTO_CommBoughtProviderOrSet{} }
func (m *TemplateDTO_CommBoughtProviderOrSet) String() string { return proto1.CompactTextString(m) }
func (*TemplateDTO_CommBoughtProviderOrSet) ProtoMessage()    {}
func (*TemplateDTO_CommBoughtProviderOrSet) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 0}
}

func (m *TemplateDTO_CommBoughtProviderOrSet) GetCommBought() []*TemplateDTO_CommBoughtProviderProp {
	if m != nil {
		return m.CommBought
	}
	return nil
}

type TemplateDTO_CommBoughtProviderProp struct {
	// Provider entity type created by the probe
	Key *Provider `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	// The list of commodities bought from the provider specified as key.
	Value []*TemplateCommodity `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
	// Specifies if the provider is optional or not.
	IsOptional       *bool  `protobuf:"varint,3,opt,name=isOptional,def=0" json:"isOptional,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TemplateDTO_CommBoughtProviderProp) Reset()         { *m = TemplateDTO_CommBoughtProviderProp{} }
func (m *TemplateDTO_CommBoughtProviderProp) String() string { return proto1.CompactTextString(m) }
func (*TemplateDTO_CommBoughtProviderProp) ProtoMessage()    {}
func (*TemplateDTO_CommBoughtProviderProp) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 1}
}

const Default_TemplateDTO_CommBoughtProviderProp_IsOptional bool = false

func (m *TemplateDTO_CommBoughtProviderProp) GetKey() *Provider {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *TemplateDTO_CommBoughtProviderProp) GetValue() []*TemplateCommodity {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TemplateDTO_CommBoughtProviderProp) GetIsOptional() bool {
	if m != nil && m.IsOptional != nil {
		return *m.IsOptional
	}
	return Default_TemplateDTO_CommBoughtProviderProp_IsOptional
}

type TemplateDTO_ExternalEntityLinkProp struct {
	Key              *EntityDTO_EntityType `protobuf:"varint,1,req,name=key,enum=proto.EntityDTO_EntityType" json:"key,omitempty"`
	Value            *ExternalEntityLink   `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *TemplateDTO_ExternalEntityLinkProp) Reset()         { *m = TemplateDTO_ExternalEntityLinkProp{} }
func (m *TemplateDTO_ExternalEntityLinkProp) String() string { return proto1.CompactTextString(m) }
func (*TemplateDTO_ExternalEntityLinkProp) ProtoMessage()    {}
func (*TemplateDTO_ExternalEntityLinkProp) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 2}
}

func (m *TemplateDTO_ExternalEntityLinkProp) GetKey() EntityDTO_EntityType {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return EntityDTO_SWITCH
}

func (m *TemplateDTO_ExternalEntityLinkProp) GetValue() *ExternalEntityLink {
	if m != nil {
		return m.Value
	}
	return nil
}

type TemplateCommodity struct {
	CommodityType *CommodityDTO_CommodityType `protobuf:"varint,1,req,name=commodityType,enum=proto.CommodityDTO_CommodityType" json:"commodityType,omitempty"`
	Key           *string                     `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	// Type of the commodity, that charges this one. This must be on of the commodities from
	// the entity (template) is expected to buy. So, this is a link between bought and sold
	// commodity of the same entity
	ChargedBy        []CommodityDTO_CommodityType `protobuf:"varint,3,rep,name=chargedBy,enum=proto.CommodityDTO_CommodityType" json:"chargedBy,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *TemplateCommodity) Reset()                    { *m = TemplateCommodity{} }
func (m *TemplateCommodity) String() string            { return proto1.CompactTextString(m) }
func (*TemplateCommodity) ProtoMessage()               {}
func (*TemplateCommodity) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *TemplateCommodity) GetCommodityType() CommodityDTO_CommodityType {
	if m != nil && m.CommodityType != nil {
		return *m.CommodityType
	}
	return CommodityDTO_CLUSTER
}

func (m *TemplateCommodity) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *TemplateCommodity) GetChargedBy() []CommodityDTO_CommodityType {
	if m != nil {
		return m.ChargedBy
	}
	return nil
}

// The Provider class creates a template entity that sells commodities to a
// consumer template.
//
// Each Provider instance has a templateClass to define the entity type, which is expressed
// as a member of the EntityType enumeration.
//
// A provider can have one of two types of relationship with the consumer entity -
// HOSTING or LAYERED_OVER (see ProviderType):
//
// HOSTING is a One Provider/Many Consumers relationship, where the provider contains the consumer.
// This means that if the provider is removed, then every consumer it contains will also be removed.
// For example, a PhysicalMachine contains many VirtualMachines. If you remove the PhysicalMachine
// entity, then its contained VMs will also be removed. You should move VMs off of a host before removing it.
//
// LAYERED_OVER is a Many/Many relationship, with no concept of containment. For example, many VMs
// can share more than one datastore. For LayeredOver relationships, you must specify max and min limits
// to determine how many providers can be layered over the given type of consumer. These values are set in the
// cardinalityMax and cardinalityMin members of this class.
type Provider struct {
	// The type of entity that the provider represents. See {@link Entity}
	// for the available types.
	TemplateClass *EntityDTO_EntityType `protobuf:"varint,1,req,name=templateClass,enum=proto.EntityDTO_EntityType" json:"templateClass,omitempty"`
	// ProviderType specifies the type of relationship between the provider and the consumer
	ProviderType *Provider_ProviderType `protobuf:"varint,2,req,name=providerType,enum=proto.Provider_ProviderType" json:"providerType,omitempty"`
	// For LAYERED_OVER providers, the maximum number of providers allowed for the consumer.
	CardinalityMax *int32 `protobuf:"varint,3,req,name=cardinalityMax" json:"cardinalityMax,omitempty"`
	// For LAYERED_OVER providers, the minimum number of providers allowed for the consumer.
	CardinalityMin   *int32 `protobuf:"varint,4,req,name=cardinalityMin" json:"cardinalityMin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Provider) Reset()                    { *m = Provider{} }
func (m *Provider) String() string            { return proto1.CompactTextString(m) }
func (*Provider) ProtoMessage()               {}
func (*Provider) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *Provider) GetTemplateClass() EntityDTO_EntityType {
	if m != nil && m.TemplateClass != nil {
		return *m.TemplateClass
	}
	return EntityDTO_SWITCH
}

func (m *Provider) GetProviderType() Provider_ProviderType {
	if m != nil && m.ProviderType != nil {
		return *m.ProviderType
	}
	return Provider_HOSTING
}

func (m *Provider) GetCardinalityMax() int32 {
	if m != nil && m.CardinalityMax != nil {
		return *m.CardinalityMax
	}
	return 0
}

func (m *Provider) GetCardinalityMin() int32 {
	if m != nil && m.CardinalityMin != nil {
		return *m.CardinalityMin
	}
	return 0
}

// ExternalEntityLink is a subclass of {@link EntityLink} that
// describes the buy/sell relationship between an entity discovered by the probe, and
// an external entity.
//
// An external entity is one that exists in the
// Operations Manager topology, but has not been discovered by the probe.
// Operations Manager uses this link to stitch discovered entities into the
// existing topology that's managed by the Operations Manager market. This external
// entity can be a provider or a consumer. The ExternalEntityLink object
// contains a full description of the relationship between the external entity and
// the node entity.
// This description includes the entity types for the buyer and seller, the ProviderType
// (the relationship type for the provider, either HOSTING or LAYERED_OVER}),
// and the list of commodities bought from the provider.
//
// To enable stitching, the external link includes a map of {@code probeEntityDef} items
// and a list of ServerEntityPropertyDef items. These work together to identify which
// external entity to stitch together with the probe's discovered entity. The {@code probeEntityDef}
// items store data to identify the appropriate external entity. For example, a discovered application
// can store the IP address of the hosting VM.
//
// The ServerEntityPropertyDef items
// tell Operations Manager how to find identifying information in the external entities.
// For example, the discovered application stores IP address of the hosting VM. Operations Manager
// will use the ServerEntityPropertyDef to test the current VMs for a matching IP address.
type ExternalEntityLink struct {
	// Consumer entity in the link
	BuyerRef *EntityDTO_EntityType `protobuf:"varint,1,req,name=buyerRef,enum=proto.EntityDTO_EntityType" json:"buyerRef,omitempty"`
	// Provider entity in the link
	SellerRef *EntityDTO_EntityType `protobuf:"varint,2,req,name=sellerRef,enum=proto.EntityDTO_EntityType" json:"sellerRef,omitempty"`
	// Provider relationship type
	Relationship *Provider_ProviderType `protobuf:"varint,3,req,name=relationship,enum=proto.Provider_ProviderType" json:"relationship,omitempty"`
	// The list of commodities the consumer entity buys from the provider entity.
	CommodityDefs []*ExternalEntityLink_CommodityDef `protobuf:"bytes,4,rep,name=commodityDefs" json:"commodityDefs,omitempty"`
	// Commodity key
	Key *string `protobuf:"bytes,5,opt,name=key" json:"key,omitempty"`
	// If one of the entity is to be found outside the probe
	HasExternalEntity *bool `protobuf:"varint,6,opt,name=hasExternalEntity" json:"hasExternalEntity,omitempty"`
	// Map of the name and description of the property belonging to the entity instances
	// discovered by the probe.
	ProbeEntityPropertyDef []*ExternalEntityLink_EntityPropertyDef `protobuf:"bytes,7,rep,name=probeEntityPropertyDef" json:"probeEntityPropertyDef,omitempty"`
	// The meta data representing the property definition of the external entity.
	// The value of the property is used for matching the entity instances.
	ExternalEntityPropertyDefs []*ExternalEntityLink_ServerEntityPropDef `protobuf:"bytes,8,rep,name=externalEntityPropertyDefs" json:"externalEntityPropertyDefs,omitempty"`
	// if the provider can replace a placeholder entity created outside of the probe,
	// give a list of EntityTypes it can replace.  For example, a LogicalPool can replace
	// a DiskArray or LogicalPool created by another probe.  The replaced entity must be
	// marked REPLACEABLE by the probe that creates it.
	ReplacesEntity   []EntityDTO_EntityType `protobuf:"varint,9,rep,name=replacesEntity,enum=proto.EntityDTO_EntityType" json:"replacesEntity,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ExternalEntityLink) Reset()                    { *m = ExternalEntityLink{} }
func (m *ExternalEntityLink) String() string            { return proto1.CompactTextString(m) }
func (*ExternalEntityLink) ProtoMessage()               {}
func (*ExternalEntityLink) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ExternalEntityLink) GetBuyerRef() EntityDTO_EntityType {
	if m != nil && m.BuyerRef != nil {
		return *m.BuyerRef
	}
	return EntityDTO_SWITCH
}

func (m *ExternalEntityLink) GetSellerRef() EntityDTO_EntityType {
	if m != nil && m.SellerRef != nil {
		return *m.SellerRef
	}
	return EntityDTO_SWITCH
}

func (m *ExternalEntityLink) GetRelationship() Provider_ProviderType {
	if m != nil && m.Relationship != nil {
		return *m.Relationship
	}
	return Provider_HOSTING
}

func (m *ExternalEntityLink) GetCommodityDefs() []*ExternalEntityLink_CommodityDef {
	if m != nil {
		return m.CommodityDefs
	}
	return nil
}

func (m *ExternalEntityLink) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ExternalEntityLink) GetHasExternalEntity() bool {
	if m != nil && m.HasExternalEntity != nil {
		return *m.HasExternalEntity
	}
	return false
}

func (m *ExternalEntityLink) GetProbeEntityPropertyDef() []*ExternalEntityLink_EntityPropertyDef {
	if m != nil {
		return m.ProbeEntityPropertyDef
	}
	return nil
}

func (m *ExternalEntityLink) GetExternalEntityPropertyDefs() []*ExternalEntityLink_ServerEntityPropDef {
	if m != nil {
		return m.ExternalEntityPropertyDefs
	}
	return nil
}

func (m *ExternalEntityLink) GetReplacesEntity() []EntityDTO_EntityType {
	if m != nil {
		return m.ReplacesEntity
	}
	return nil
}

type ExternalEntityLink_CommodityDef struct {
	Type             *CommodityDTO_CommodityType `protobuf:"varint,1,req,name=type,enum=proto.CommodityDTO_CommodityType" json:"type,omitempty"`
	HasKey           *bool                       `protobuf:"varint,2,opt,name=hasKey,def=0" json:"hasKey,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *ExternalEntityLink_CommodityDef) Reset()         { *m = ExternalEntityLink_CommodityDef{} }
func (m *ExternalEntityLink_CommodityDef) String() string { return proto1.CompactTextString(m) }
func (*ExternalEntityLink_CommodityDef) ProtoMessage()    {}
func (*ExternalEntityLink_CommodityDef) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{3, 0}
}

const Default_ExternalEntityLink_CommodityDef_HasKey bool = false

func (m *ExternalEntityLink_CommodityDef) GetType() CommodityDTO_CommodityType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CommodityDTO_CLUSTER
}

func (m *ExternalEntityLink_CommodityDef) GetHasKey() bool {
	if m != nil && m.HasKey != nil {
		return *m.HasKey
	}
	return Default_ExternalEntityLink_CommodityDef_HasKey
}

// The ServerEntityPropDef class provides metadata properties for entities
// in the Operations Manager topology that have not been discovered by this probe.
// Operations Manager uses these property values to stitch external entities to the
// entities discovered by the probe.
// An external entity is one that exists in the Operations Manager topology, but has
// not been discovered by the probe.
//
// The link definition identifies:
// * The entity type for this external entity
// * An attribute of the entity to use to identify it (for example a physical machine's IP address)
// * A flag to set whether to fetch the attribute from an entity that is related in the
//   Operations Manager topology (for example, use the IP address of a VM's host physical machine)
// * Optionally, a handler that can traverse the topology to find the identifying value
//
// This class includes a set of constants for properties that apply to some of the supported
// entity types. Use these constants to create external links with the most common entity
// types in the Operations Manager topology. You can also use this class to create
// custom external entity link definitions.
//
// EXAMPLE: Connecting a DiskArray to Storage. To connect storage objects to disk arrays that the probe discovers,
// the entity link can use either the LUN ID, WWN, or export path properties. This class includes
// the STORAGE_LUNID, STORAGE_WWN, and STORAGE_REMOTE_HOST constants.
// You can use one of these constants as the ExternalEntityLinkDef in the
// ExternalEntityLink that you create for the discovered disk array.
//
// EXAMPLE: Connecting an Application to a VM. To connect an application the probe discovers to a VM,
// you typically use the VM's IP address. You could also use the VM's unique ID. This class includes
// the VM_IP constant for VM IP addresses, and the VM_UUID constant for the VM unique ID.
// You can use one of these constants as the ExternalEntityLinkDef in the
// ExternalEntityLink that you create for the discovered VM.
type ExternalEntityLink_ServerEntityPropDef struct {
	Entity           *EntityDTO_EntityType               `protobuf:"varint,1,req,name=entity,enum=proto.EntityDTO_EntityType" json:"entity,omitempty"`
	Attribute        *string                             `protobuf:"bytes,2,req,name=attribute" json:"attribute,omitempty"`
	UseTopoExt       *bool                               `protobuf:"varint,3,opt,name=useTopoExt" json:"useTopoExt,omitempty"`
	PropertyHandler  *ExternalEntityLink_PropertyHandler `protobuf:"bytes,4,opt,name=propertyHandler" json:"propertyHandler,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *ExternalEntityLink_ServerEntityPropDef) Reset() {
	*m = ExternalEntityLink_ServerEntityPropDef{}
}
func (m *ExternalEntityLink_ServerEntityPropDef) String() string { return proto1.CompactTextString(m) }
func (*ExternalEntityLink_ServerEntityPropDef) ProtoMessage()    {}
func (*ExternalEntityLink_ServerEntityPropDef) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{3, 1}
}

func (m *ExternalEntityLink_ServerEntityPropDef) GetEntity() EntityDTO_EntityType {
	if m != nil && m.Entity != nil {
		return *m.Entity
	}
	return EntityDTO_SWITCH
}

func (m *ExternalEntityLink_ServerEntityPropDef) GetAttribute() string {
	if m != nil && m.Attribute != nil {
		return *m.Attribute
	}
	return ""
}

func (m *ExternalEntityLink_ServerEntityPropDef) GetUseTopoExt() bool {
	if m != nil && m.UseTopoExt != nil {
		return *m.UseTopoExt
	}
	return false
}

func (m *ExternalEntityLink_ServerEntityPropDef) GetPropertyHandler() *ExternalEntityLink_PropertyHandler {
	if m != nil {
		return m.PropertyHandler
	}
	return nil
}

// Holds a property for the probe's discovered entity that Operations Manager can use to stitch the discovered entity
// into the Operations Manager topology. Each property contains a property name and a description.
//
// The property name specifies which property of the discovered entity you want to match. The discovered
// entity's DTO contains the list of properties and values for that entity. This link must include a property that matches a
// named property in the DTO. Note that the SDK includes builders for different types of entities.
// These builders add properties to the entity DTO, giving them names from the {@link SupplyChainConstants} enumeration.
// However, you can use arbitrary names for these properties, so long as the named property is declared in the
// entity DTO.
//
// The properties you create here match the property names in the target DTO.
// For example, the {link ApplicationBuilder} adds an IP address as a property named {@code SupplyChainConstants.IP_ADDRESS}.
// To match the application IP address in this link, add a property to the link with the same name. By doing that,
// the stitching process can access the value that is set in the discovered entity's DTO.
//
// The property description is an arbitrary string to describe the purpose of this property. This is useful
// when you print out the link via a {@code toString()} method.
type ExternalEntityLink_EntityPropertyDef struct {
	// An entity property name
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// An arbitrary description
	Description      *string `protobuf:"bytes,2,req,name=description" json:"description,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExternalEntityLink_EntityPropertyDef) Reset()         { *m = ExternalEntityLink_EntityPropertyDef{} }
func (m *ExternalEntityLink_EntityPropertyDef) String() string { return proto1.CompactTextString(m) }
func (*ExternalEntityLink_EntityPropertyDef) ProtoMessage()    {}
func (*ExternalEntityLink_EntityPropertyDef) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{3, 2}
}

func (m *ExternalEntityLink_EntityPropertyDef) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ExternalEntityLink_EntityPropertyDef) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

// The PropertyHandler class manages handler methods that
// Operations Manager can use to traverse the topology to get value(s) of the specified attribute.
// The class assembles a linked list of handlers that can be used to inspect multiple
// layers of the topology to get properties from different entity types.
// Forms LinkedList structure to maintain multiple layers of property names and entity types.
// Example:
// PropertyHandler ipHandler {
//                nextHandler: null
//                methodName: "getAddress"
//                entity: Entity.IP
//                directlyApply: false
//                }</pre></code>
// ipHandler can be used to retrieve IP address string values from IP object.
type ExternalEntityLink_PropertyHandler struct {
	MethodName *string               `protobuf:"bytes,1,req,name=methodName" json:"methodName,omitempty"`
	EntityType *EntityDTO_EntityType `protobuf:"varint,2,opt,name=entityType,enum=proto.EntityDTO_EntityType" json:"entityType,omitempty"`
	// it notifies if the method can be directly applied to what returned from the previous layer
	// For example, if it's for IP from VM.getUsesEndPoints(), then directlyApply should be false.
	// Since what returned from VM.getUsesEndPoints() is a list of IPs. So should go to each
	// instance and apply that method.
	DirectlyApply    *bool                               `protobuf:"varint,3,opt,name=directlyApply" json:"directlyApply,omitempty"`
	NextHandler      *ExternalEntityLink_PropertyHandler `protobuf:"bytes,4,opt,name=next_handler,json=nextHandler" json:"next_handler,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *ExternalEntityLink_PropertyHandler) Reset()         { *m = ExternalEntityLink_PropertyHandler{} }
func (m *ExternalEntityLink_PropertyHandler) String() string { return proto1.CompactTextString(m) }
func (*ExternalEntityLink_PropertyHandler) ProtoMessage()    {}
func (*ExternalEntityLink_PropertyHandler) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{3, 3}
}

func (m *ExternalEntityLink_PropertyHandler) GetMethodName() string {
	if m != nil && m.MethodName != nil {
		return *m.MethodName
	}
	return ""
}

func (m *ExternalEntityLink_PropertyHandler) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *ExternalEntityLink_PropertyHandler) GetDirectlyApply() bool {
	if m != nil && m.DirectlyApply != nil {
		return *m.DirectlyApply
	}
	return false
}

func (m *ExternalEntityLink_PropertyHandler) GetNextHandler() *ExternalEntityLink_PropertyHandler {
	if m != nil {
		return m.NextHandler
	}
	return nil
}

func init() {
	proto1.RegisterType((*TemplateDTO)(nil), "proto.TemplateDTO")
	proto1.RegisterType((*TemplateDTO_CommBoughtProviderOrSet)(nil), "proto.TemplateDTO.CommBoughtProviderOrSet")
	proto1.RegisterType((*TemplateDTO_CommBoughtProviderProp)(nil), "proto.TemplateDTO.CommBoughtProviderProp")
	proto1.RegisterType((*TemplateDTO_ExternalEntityLinkProp)(nil), "proto.TemplateDTO.ExternalEntityLinkProp")
	proto1.RegisterType((*TemplateCommodity)(nil), "proto.TemplateCommodity")
	proto1.RegisterType((*Provider)(nil), "proto.Provider")
	proto1.RegisterType((*ExternalEntityLink)(nil), "proto.ExternalEntityLink")
	proto1.RegisterType((*ExternalEntityLink_CommodityDef)(nil), "proto.ExternalEntityLink.CommodityDef")
	proto1.RegisterType((*ExternalEntityLink_ServerEntityPropDef)(nil), "proto.ExternalEntityLink.ServerEntityPropDef")
	proto1.RegisterType((*ExternalEntityLink_EntityPropertyDef)(nil), "proto.ExternalEntityLink.EntityPropertyDef")
	proto1.RegisterType((*ExternalEntityLink_PropertyHandler)(nil), "proto.ExternalEntityLink.PropertyHandler")
	proto1.RegisterEnum("proto.TemplateDTO_TemplateType", TemplateDTO_TemplateType_name, TemplateDTO_TemplateType_value)
	proto1.RegisterEnum("proto.Provider_ProviderType", Provider_ProviderType_name, Provider_ProviderType_value)
}

func init() { proto1.RegisterFile("SupplyChain.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 953 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdb, 0x6e, 0x23, 0x45,
	0x10, 0xdd, 0xf1, 0x25, 0xb1, 0xcb, 0x4e, 0xe2, 0x14, 0x52, 0x18, 0xcc, 0xc2, 0x7a, 0x2d, 0x58,
	0xcc, 0x42, 0x8c, 0x14, 0x84, 0x10, 0x20, 0x01, 0xb1, 0x63, 0xed, 0x46, 0x64, 0xe3, 0x68, 0x6c,
	0x21, 0x78, 0x5a, 0x75, 0x3c, 0xed, 0xf5, 0x68, 0xc7, 0xd3, 0xa3, 0x9e, 0x76, 0x64, 0xff, 0x04,
	0x0f, 0xbc, 0xf1, 0x15, 0x7c, 0x11, 0x2f, 0xbc, 0xf0, 0x1b, 0xa8, 0x2f, 0xb6, 0x7b, 0x7c, 0xc9,
	0x05, 0xf1, 0xe4, 0xee, 0xea, 0x73, 0x4e, 0x57, 0x55, 0x57, 0xd5, 0x18, 0x0e, 0x7b, 0x93, 0x38,
	0x0e, 0x67, 0xed, 0x11, 0x09, 0xa2, 0x66, 0xcc, 0x99, 0x60, 0x98, 0x57, 0x3f, 0xd5, 0x83, 0x36,
	0x1b, 0x8f, 0x59, 0x74, 0xd6, 0xef, 0x6a, 0x7b, 0xfd, 0x8f, 0x5d, 0x28, 0xf5, 0xe9, 0x38, 0x0e,
	0x89, 0xa0, 0x67, 0xfd, 0x2e, 0x9e, 0xc2, 0x9e, 0x30, 0xdb, 0x76, 0x48, 0x92, 0xc4, 0x75, 0x6a,
	0x99, 0xc6, 0xfe, 0xc9, 0xfb, 0x1a, 0xde, 0xec, 0x44, 0x22, 0x10, 0x33, 0x49, 0xd7, 0xab, 0xfe,
	0x2c, 0xa6, 0x5e, 0x9a, 0x81, 0x6d, 0x28, 0xcf, 0x0d, 0xf2, 0xd8, 0xcd, 0x28, 0x85, 0x27, 0x46,
	0xc1, 0xba, 0x6c, 0xb1, 0x56, 0x2a, 0x29, 0x12, 0x3e, 0x87, 0xca, 0x7c, 0x7f, 0xc5, 0x03, 0xc6,
	0x03, 0x31, 0x73, 0xb3, 0xb5, 0x4c, 0x23, 0xef, 0xad, 0xd9, 0xf1, 0x7b, 0xd8, 0x1b, 0xc8, 0xb0,
	0xfc, 0x40, 0xcc, 0x7a, 0x2c, 0xf4, 0xdd, 0x7c, 0x2d, 0xdb, 0x28, 0x9d, 0xb8, 0x2b, 0x37, 0xb6,
	0xe7, 0x18, 0x2f, 0x0d, 0xc7, 0x1e, 0x1c, 0x2c, 0x0c, 0x2d, 0x36, 0x79, 0x33, 0x12, 0xee, 0x8e,
	0x52, 0xf8, 0x74, 0x83, 0xcf, 0x52, 0x45, 0x83, 0xae, 0x38, 0xbb, 0x09, 0x7c, 0xca, 0xaf, 0x38,
	0x8b, 0xbd, 0x55, 0x05, 0x7c, 0x05, 0x65, 0x3a, 0x15, 0x94, 0x47, 0x24, 0xbc, 0x08, 0xa2, 0xb7,
	0xee, 0xee, 0x56, 0xc5, 0x8e, 0x81, 0xe9, 0x8c, 0x4a, 0xb0, 0x52, 0x4c, 0xd1, 0xb1, 0xaf, 0x7d,
	0xd4, 0xe2, 0x5d, 0xde, 0xa3, 0xc2, 0x2d, 0x28, 0xc5, 0xe7, 0xf7, 0xf2, 0x51, 0x31, 0xbc, 0x55,
	0x89, 0xaa, 0x0f, 0xef, 0x6e, 0xc1, 0xe2, 0x39, 0xc0, 0x12, 0xed, 0x3a, 0x0f, 0xcd, 0x87, 0x45,
	0xae, 0xfe, 0xee, 0xc0, 0xd1, 0x66, 0x18, 0x3e, 0x85, 0xec, 0x5b, 0x3a, 0x53, 0x45, 0x56, 0x3a,
	0x39, 0x30, 0xf2, 0x73, 0x84, 0x27, 0xcf, 0xb0, 0x09, 0xf9, 0x1b, 0x12, 0x4e, 0x64, 0x1d, 0xdd,
	0xfe, 0xaa, 0x1a, 0x86, 0x1f, 0x03, 0x04, 0x49, 0x37, 0x16, 0x01, 0x8b, 0x48, 0xe8, 0x66, 0x6b,
	0x4e, 0xa3, 0xf0, 0x6d, 0x7e, 0x48, 0xc2, 0x84, 0x7a, 0xd6, 0x41, 0x75, 0x0a, 0x47, 0x9b, 0x13,
	0x8f, 0xc7, 0x4b, 0x9f, 0xee, 0x28, 0x7c, 0xe5, 0xdf, 0x17, 0x4b, 0xff, 0x64, 0x10, 0xef, 0xcd,
	0x09, 0x6b, 0xe2, 0xc6, 0xc1, 0xfa, 0x27, 0x50, 0xb6, 0x0b, 0x1f, 0x0b, 0x90, 0x6b, 0x9d, 0xf6,
	0x3a, 0x95, 0x47, 0xb8, 0x07, 0xc5, 0xce, 0x2f, 0xfd, 0xce, 0x65, 0xef, 0xbc, 0x7b, 0x59, 0x71,
	0xea, 0x7f, 0x3a, 0x70, 0xb8, 0x16, 0x26, 0xbe, 0xb0, 0xaa, 0x5d, 0xf5, 0x97, 0x76, 0xf4, 0xa9,
	0xb9, 0x77, 0x01, 0x9c, 0x3f, 0xce, 0x02, 0xe8, 0xa5, 0x79, 0x58, 0xd1, 0x71, 0x66, 0x6a, 0x4e,
	0xa3, 0xa8, 0x43, 0xf9, 0x01, 0x8a, 0x83, 0x11, 0xe1, 0x6f, 0xa8, 0xdf, 0x92, 0xdd, 0x96, 0xbd,
	0x9f, 0xec, 0x92, 0x53, 0xff, 0x2d, 0x03, 0x85, 0xf9, 0xeb, 0xfd, 0x1f, 0xa3, 0xe4, 0x47, 0x28,
	0xc7, 0x46, 0xce, 0x1a, 0x25, 0x8f, 0x57, 0xea, 0x64, 0xb1, 0xd0, 0x73, 0xc4, 0x66, 0xe0, 0x33,
	0xd8, 0x1f, 0x10, 0xee, 0x07, 0x11, 0x09, 0x03, 0x31, 0x7b, 0x45, 0xa6, 0x66, 0x8a, 0xac, 0x58,
	0x57, 0x71, 0x41, 0xe4, 0xe6, 0xd6, 0x71, 0x41, 0x54, 0x3f, 0x86, 0xb2, 0x7d, 0x1b, 0x96, 0x60,
	0xf7, 0x65, 0xb7, 0xd7, 0x3f, 0xbf, 0x7c, 0x51, 0x79, 0x84, 0x15, 0x28, 0x5f, 0x9c, 0xfe, 0xda,
	0xf1, 0x3a, 0x67, 0xaf, 0xbb, 0x3f, 0x77, 0xbc, 0x8a, 0x53, 0xff, 0xa7, 0x08, 0xb8, 0x5e, 0x09,
	0xf8, 0x35, 0x14, 0xae, 0x27, 0x33, 0xca, 0x3d, 0x3a, 0xbc, 0x4f, 0x56, 0x16, 0x60, 0xfc, 0x06,
	0x8a, 0x09, 0x0d, 0x43, 0xcd, 0xcc, 0xdc, 0xcd, 0x5c, 0xa2, 0x65, 0x2e, 0x39, 0x0d, 0x89, 0x2c,
	0xff, 0x64, 0x14, 0xc4, 0x2a, 0x0f, 0x77, 0xe6, 0xd2, 0x66, 0xe0, 0x85, 0x55, 0x79, 0x67, 0x74,
	0x98, 0xb8, 0x39, 0xd5, 0x91, 0xcf, 0xb6, 0x56, 0xbc, 0x55, 0x35, 0x74, 0xe8, 0xa5, 0xc9, 0xf3,
	0xf2, 0xcb, 0x2f, 0xcb, 0xef, 0x73, 0x38, 0x1c, 0x91, 0x24, 0x2d, 0xe3, 0xee, 0xc8, 0x06, 0xf6,
	0xd6, 0x0f, 0x70, 0x00, 0x47, 0x31, 0x67, 0xd7, 0x54, 0x6f, 0x65, 0xe7, 0x52, 0xae, 0xa4, 0xcd,
	0xa8, 0xfd, 0x6c, 0xbb, 0x5b, 0x6b, 0x14, 0x6f, 0x8b, 0x14, 0x8e, 0xa1, 0x4a, 0x53, 0x7c, 0xeb,
	0x30, 0x31, 0x13, 0xf8, 0x78, 0xfb, 0x45, 0x3d, 0xca, 0x6f, 0x28, 0x5f, 0x32, 0xe5, 0x55, 0xb7,
	0x08, 0x62, 0x1b, 0xf6, 0x39, 0x8d, 0x43, 0x32, 0xa0, 0x89, 0x09, 0xbf, 0xa8, 0xba, 0xf0, 0xd6,
	0x37, 0x5e, 0xa1, 0x54, 0x7d, 0x28, 0xdb, 0x79, 0xc7, 0xaf, 0x20, 0x27, 0x1e, 0x34, 0x27, 0x14,
	0x1c, 0x3f, 0x80, 0x9d, 0x11, 0x49, 0x7e, 0x32, 0x13, 0x62, 0x31, 0x43, 0x8d, 0xb1, 0xfa, 0x97,
	0x03, 0xef, 0x6c, 0x08, 0x0f, 0xbf, 0x84, 0x1d, 0xaa, 0x5d, 0xbf, 0x47, 0x61, 0x1b, 0x28, 0x3e,
	0x86, 0x22, 0x11, 0x82, 0x07, 0xd7, 0x13, 0xa1, 0x9b, 0xbc, 0xe8, 0x2d, 0x0d, 0xf8, 0x21, 0xc0,
	0x24, 0xa1, 0x7d, 0x16, 0xb3, 0xce, 0x54, 0xe8, 0x89, 0xee, 0x59, 0x16, 0xf9, 0xfd, 0x8e, 0x4d,
	0x16, 0x5f, 0x92, 0xc8, 0x0f, 0x29, 0x77, 0x73, 0x35, 0xc7, 0xfa, 0x5e, 0x6d, 0x78, 0x99, 0xab,
	0x34, 0xc1, 0x5b, 0x55, 0xa8, 0x9e, 0xc3, 0xe1, 0x7a, 0x39, 0x20, 0xe4, 0x22, 0x32, 0xd6, 0xa9,
	0x2c, 0x7a, 0x6a, 0x8d, 0x35, 0x28, 0xf9, 0x34, 0x19, 0xf0, 0x40, 0x7d, 0x59, 0x8c, 0xf7, 0xb6,
	0xa9, 0xfa, 0xb7, 0x03, 0x07, 0x2b, 0xf7, 0xc9, 0x98, 0xc6, 0x54, 0x8c, 0x98, 0x7f, 0xb9, 0xd4,
	0xb3, 0x2c, 0xf8, 0x1d, 0x00, 0x5d, 0xe4, 0x49, 0xbd, 0xc0, 0x1d, 0xa9, 0xb4, 0xe0, 0xf8, 0x11,
	0xec, 0xf9, 0x01, 0xa7, 0x03, 0x11, 0xce, 0x4e, 0xe5, 0x1f, 0x41, 0x93, 0xb3, 0xb4, 0x11, 0x2f,
	0xa0, 0x1c, 0xd1, 0xa9, 0x78, 0x3d, 0xfa, 0xaf, 0x39, 0x2b, 0x49, 0xba, 0xd9, 0xb4, 0x9a, 0xf0,
	0x64, 0xc0, 0xc6, 0xcd, 0x9b, 0xb1, 0x98, 0xf0, 0x6b, 0xd6, 0x94, 0x43, 0x7c, 0xc8, 0xf8, 0xb8,
	0xa9, 0x9a, 0x3e, 0x6a, 0xfa, 0x82, 0xb5, 0x4a, 0xd6, 0xdf, 0xd2, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x11, 0x7d, 0x8e, 0xf7, 0xa4, 0x0a, 0x00, 0x00,
}
