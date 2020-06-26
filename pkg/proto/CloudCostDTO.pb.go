// Code generated by protoc-gen-go.
// source: CloudCostDTO.proto
// DO NOT EDIT!

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Historical demand data type.
type DemandType int32

const (
	// Consumption based data to be used
	DemandType_CONSUMPTION DemandType = 1
	// Allocation based data to be used
	DemandType_ALLOCATION DemandType = 2
)

var DemandType_name = map[int32]string{
	1: "CONSUMPTION",
	2: "ALLOCATION",
}
var DemandType_value = map[string]int32{
	"CONSUMPTION": 1,
	"ALLOCATION":  2,
}

func (x DemandType) Enum() *DemandType {
	p := new(DemandType)
	*p = x
	return p
}
func (x DemandType) String() string {
	return proto.EnumName(DemandType_name, int32(x))
}
func (x *DemandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DemandType_value, data, "DemandType")
	if err != nil {
		return err
	}
	*x = DemandType(value)
	return nil
}
func (DemandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// The tenancy of an instance defines what hardware the instance is running on.
type Tenancy int32

const (
	// Instance runs on shared/default hardware.
	// This is typically the cheapest option.
	Tenancy_DEFAULT Tenancy = 1
	// Instance runs on single-tenant hardware.
	// That means your instance runs on a host that's separate from other customers,
	// but the host details are abstracted away, and you're not paying for the whole host.
	Tenancy_DEDICATED Tenancy = 2
	// Instance runs on a dedicated Host.
	// This means your instance runs on a specific host, and you are paying for the full host and
	// are responsible for managing it.
	Tenancy_HOST Tenancy = 3
)

var Tenancy_name = map[int32]string{
	1: "DEFAULT",
	2: "DEDICATED",
	3: "HOST",
}
var Tenancy_value = map[string]int32{
	"DEFAULT":   1,
	"DEDICATED": 2,
	"HOST":      3,
}

func (x Tenancy) Enum() *Tenancy {
	p := new(Tenancy)
	*p = x
	return p
}
func (x Tenancy) String() string {
	return proto.EnumName(Tenancy_name, int32(x))
}
func (x *Tenancy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Tenancy_value, data, "Tenancy")
	if err != nil {
		return err
	}
	*x = Tenancy(value)
	return nil
}
func (Tenancy) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// The supported operating systems. Keep in sync with com.vmturbo.mediation.hybrid.cloud.common.OsType.
// WINDOWS_SERVER and WINDOWS_SERVER_BURST are deprecated as they were only used to capture license
// costs for Azure templates. They were never directly associated in VM.
type OSType int32

const (
	OSType_UNKNOWN_OS OSType = 0
	// Unix OS.
	OSType_LINUX                     OSType = 2
	OSType_SUSE                      OSType = 3
	OSType_RHEL                      OSType = 4
	OSType_LINUX_WITH_SQL_ENTERPRISE OSType = 5
	OSType_LINUX_WITH_SQL_STANDARD   OSType = 6
	OSType_LINUX_WITH_SQL_WEB        OSType = 7
	// Windows OS.
	OSType_WINDOWS                     OSType = 20
	OSType_WINDOWS_WITH_SQL_STANDARD   OSType = 21
	OSType_WINDOWS_WITH_SQL_WEB        OSType = 22
	OSType_WINDOWS_WITH_SQL_ENTERPRISE OSType = 23
	OSType_WINDOWS_BYOL                OSType = 24
	OSType_WINDOWS_SERVER              OSType = 25
	OSType_WINDOWS_SERVER_BURST        OSType = 26
)

var OSType_name = map[int32]string{
	0:  "UNKNOWN_OS",
	2:  "LINUX",
	3:  "SUSE",
	4:  "RHEL",
	5:  "LINUX_WITH_SQL_ENTERPRISE",
	6:  "LINUX_WITH_SQL_STANDARD",
	7:  "LINUX_WITH_SQL_WEB",
	20: "WINDOWS",
	21: "WINDOWS_WITH_SQL_STANDARD",
	22: "WINDOWS_WITH_SQL_WEB",
	23: "WINDOWS_WITH_SQL_ENTERPRISE",
	24: "WINDOWS_BYOL",
	25: "WINDOWS_SERVER",
	26: "WINDOWS_SERVER_BURST",
}
var OSType_value = map[string]int32{
	"UNKNOWN_OS": 0,
	"LINUX":      2,
	"SUSE":       3,
	"RHEL":       4,
	"LINUX_WITH_SQL_ENTERPRISE":   5,
	"LINUX_WITH_SQL_STANDARD":     6,
	"LINUX_WITH_SQL_WEB":          7,
	"WINDOWS":                     20,
	"WINDOWS_WITH_SQL_STANDARD":   21,
	"WINDOWS_WITH_SQL_WEB":        22,
	"WINDOWS_WITH_SQL_ENTERPRISE": 23,
	"WINDOWS_BYOL":                24,
	"WINDOWS_SERVER":              25,
	"WINDOWS_SERVER_BURST":        26,
}

func (x OSType) Enum() *OSType {
	p := new(OSType)
	*p = x
	return p
}
func (x OSType) String() string {
	return proto.EnumName(OSType_name, int32(x))
}
func (x *OSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OSType_value, data, "OSType")
	if err != nil {
		return err
	}
	*x = OSType(value)
	return nil
}
func (OSType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// The engine for a database tier.
// This is an enum to save on space - and also because
// the list of supported engines across cloud providers is pretty small.
type DatabaseEngine int32

const (
	DatabaseEngine_UNKNOWN          DatabaseEngine = 0
	DatabaseEngine_MYSQL            DatabaseEngine = 1
	DatabaseEngine_MARIADB          DatabaseEngine = 2
	DatabaseEngine_POSTGRESQL       DatabaseEngine = 3
	DatabaseEngine_ORACLE           DatabaseEngine = 4
	DatabaseEngine_SQLSERVER        DatabaseEngine = 5
	DatabaseEngine_AURORA           DatabaseEngine = 6
	DatabaseEngine_AURORAMYSQL      DatabaseEngine = 7
	DatabaseEngine_AURORAPOSTGRESQL DatabaseEngine = 8
	DatabaseEngine_MONGO            DatabaseEngine = 9
)

var DatabaseEngine_name = map[int32]string{
	0: "UNKNOWN",
	1: "MYSQL",
	2: "MARIADB",
	3: "POSTGRESQL",
	4: "ORACLE",
	5: "SQLSERVER",
	6: "AURORA",
	7: "AURORAMYSQL",
	8: "AURORAPOSTGRESQL",
	9: "MONGO",
}
var DatabaseEngine_value = map[string]int32{
	"UNKNOWN":          0,
	"MYSQL":            1,
	"MARIADB":          2,
	"POSTGRESQL":       3,
	"ORACLE":           4,
	"SQLSERVER":        5,
	"AURORA":           6,
	"AURORAMYSQL":      7,
	"AURORAPOSTGRESQL": 8,
	"MONGO":            9,
}

func (x DatabaseEngine) Enum() *DatabaseEngine {
	p := new(DatabaseEngine)
	*p = x
	return p
}
func (x DatabaseEngine) String() string {
	return proto.EnumName(DatabaseEngine_name, int32(x))
}
func (x *DatabaseEngine) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DatabaseEngine_value, data, "DatabaseEngine")
	if err != nil {
		return err
	}
	*x = DatabaseEngine(value)
	return nil
}
func (DatabaseEngine) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// LicenseModel describes all supported license models in cloud projects.
type LicenseModel int32

const (
	LicenseModel_BRING_YOUR_OWN_LICENSE LicenseModel = 0
	LicenseModel_LICENSE_INCLUDED       LicenseModel = 1
	LicenseModel_NO_LICENSE_REQUIRED    LicenseModel = 2
)

var LicenseModel_name = map[int32]string{
	0: "BRING_YOUR_OWN_LICENSE",
	1: "LICENSE_INCLUDED",
	2: "NO_LICENSE_REQUIRED",
}
var LicenseModel_value = map[string]int32{
	"BRING_YOUR_OWN_LICENSE": 0,
	"LICENSE_INCLUDED":       1,
	"NO_LICENSE_REQUIRED":    2,
}

func (x LicenseModel) Enum() *LicenseModel {
	p := new(LicenseModel)
	*p = x
	return p
}
func (x LicenseModel) String() string {
	return proto.EnumName(LicenseModel_name, int32(x))
}
func (x *LicenseModel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LicenseModel_value, data, "LicenseModel")
	if err != nil {
		return err
	}
	*x = LicenseModel(value)
	return nil
}
func (LicenseModel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// DeploymentType describes all supported deployment types by cloud probes.
type DeploymentType int32

const (
	DeploymentType_SINGLE_AZ DeploymentType = 0
	DeploymentType_MULTI_AZ  DeploymentType = 1
)

var DeploymentType_name = map[int32]string{
	0: "SINGLE_AZ",
	1: "MULTI_AZ",
}
var DeploymentType_value = map[string]int32{
	"SINGLE_AZ": 0,
	"MULTI_AZ":  1,
}

func (x DeploymentType) Enum() *DeploymentType {
	p := new(DeploymentType)
	*p = x
	return p
}
func (x DeploymentType) String() string {
	return proto.EnumName(DeploymentType_name, int32(x))
}
func (x *DeploymentType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeploymentType_value, data, "DeploymentType")
	if err != nil {
		return err
	}
	*x = DeploymentType(value)
	return nil
}
func (DeploymentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// The edition of a database engine.
// The edition enum is closely related to the DatabaseEngine enum, and in the future it may be
// worth it to have a separate "database identifier" message that forbids illegal
// engine-edition combinations. For now, there are only two database engines with editions,
// so this seems manageable.
type DatabaseEdition int32

const (
	// Possible list of edition values.
	DatabaseEdition_NONE        DatabaseEdition = 0
	DatabaseEdition_ENTERPRISE  DatabaseEdition = 1
	DatabaseEdition_STANDARD    DatabaseEdition = 2
	DatabaseEdition_STANDARDONE DatabaseEdition = 3
	DatabaseEdition_STANDARDTWO DatabaseEdition = 4
	DatabaseEdition_WEB         DatabaseEdition = 12
	DatabaseEdition_EXPRESS     DatabaseEdition = 13
)

var DatabaseEdition_name = map[int32]string{
	0:  "NONE",
	1:  "ENTERPRISE",
	2:  "STANDARD",
	3:  "STANDARDONE",
	4:  "STANDARDTWO",
	12: "WEB",
	13: "EXPRESS",
}
var DatabaseEdition_value = map[string]int32{
	"NONE":        0,
	"ENTERPRISE":  1,
	"STANDARD":    2,
	"STANDARDONE": 3,
	"STANDARDTWO": 4,
	"WEB":         12,
	"EXPRESS":     13,
}

func (x DatabaseEdition) Enum() *DatabaseEdition {
	p := new(DatabaseEdition)
	*p = x
	return p
}
func (x DatabaseEdition) String() string {
	return proto.EnumName(DatabaseEdition_name, int32(x))
}
func (x *DatabaseEdition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DatabaseEdition_value, data, "DatabaseEdition")
	if err != nil {
		return err
	}
	*x = DatabaseEdition(value)
	return nil
}
func (DatabaseEdition) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type ReservedInstanceType_OfferingClass int32

const (
	// Most of the attributes of a standard reserved instance are "fixed" at the time it's
	// bought.
	ReservedInstanceType_STANDARD ReservedInstanceType_OfferingClass = 1
	// A convertible reserved instance can be exchanged for a different
	// instance type. Not all service providers offer convertible reserved
	// instances.
	ReservedInstanceType_CONVERTIBLE ReservedInstanceType_OfferingClass = 2
)

var ReservedInstanceType_OfferingClass_name = map[int32]string{
	1: "STANDARD",
	2: "CONVERTIBLE",
}
var ReservedInstanceType_OfferingClass_value = map[string]int32{
	"STANDARD":    1,
	"CONVERTIBLE": 2,
}

func (x ReservedInstanceType_OfferingClass) Enum() *ReservedInstanceType_OfferingClass {
	p := new(ReservedInstanceType_OfferingClass)
	*p = x
	return p
}
func (x ReservedInstanceType_OfferingClass) String() string {
	return proto.EnumName(ReservedInstanceType_OfferingClass_name, int32(x))
}
func (x *ReservedInstanceType_OfferingClass) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReservedInstanceType_OfferingClass_value, data, "ReservedInstanceType_OfferingClass")
	if err != nil {
		return err
	}
	*x = ReservedInstanceType_OfferingClass(value)
	return nil
}
func (ReservedInstanceType_OfferingClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 0}
}

type ReservedInstanceType_PaymentOption int32

const (
	// The user must pay the entire price of this instance upfront. There is no recurring
	// cost.
	// (e.g. $10000.00 upfront for the year)
	ReservedInstanceType_ALL_UPFRONT ReservedInstanceType_PaymentOption = 1
	// The user must pay some part of the instance price upfront, and the rest over time.
	// (e.g. $1000.00 upfront, and $0.5 per instance-hour afterwards).
	ReservedInstanceType_PARTIAL_UPFRONT ReservedInstanceType_PaymentOption = 2
	// The entire price of the instance is recurring
	// (e.g. $0.7 per instance-hour)
	ReservedInstanceType_NO_UPFRONT ReservedInstanceType_PaymentOption = 3
)

var ReservedInstanceType_PaymentOption_name = map[int32]string{
	1: "ALL_UPFRONT",
	2: "PARTIAL_UPFRONT",
	3: "NO_UPFRONT",
}
var ReservedInstanceType_PaymentOption_value = map[string]int32{
	"ALL_UPFRONT":     1,
	"PARTIAL_UPFRONT": 2,
	"NO_UPFRONT":      3,
}

func (x ReservedInstanceType_PaymentOption) Enum() *ReservedInstanceType_PaymentOption {
	p := new(ReservedInstanceType_PaymentOption)
	*p = x
	return p
}
func (x ReservedInstanceType_PaymentOption) String() string {
	return proto.EnumName(ReservedInstanceType_PaymentOption_name, int32(x))
}
func (x *ReservedInstanceType_PaymentOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReservedInstanceType_PaymentOption_value, data, "ReservedInstanceType_PaymentOption")
	if err != nil {
		return err
	}
	*x = ReservedInstanceType_PaymentOption(value)
	return nil
}
func (ReservedInstanceType_PaymentOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 1}
}

// Identifies the type of "reservation" for the instance, and the
// payment conditions.
type ReservedInstanceType struct {
	// The type of offering.
	OfferingClass *ReservedInstanceType_OfferingClass `protobuf:"varint,1,opt,name=offering_class,json=offeringClass,enum=common_dto.ReservedInstanceType_OfferingClass" json:"offering_class,omitempty"`
	// The payment option for this reserved instance.
	PaymentOption *ReservedInstanceType_PaymentOption `protobuf:"varint,2,opt,name=payment_option,json=paymentOption,enum=common_dto.ReservedInstanceType_PaymentOption" json:"payment_option,omitempty"`
	// The term, in years.
	TermYears        *uint32 `protobuf:"varint,3,opt,name=term_years,json=termYears" json:"term_years,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReservedInstanceType) Reset()                    { *m = ReservedInstanceType{} }
func (m *ReservedInstanceType) String() string            { return proto.CompactTextString(m) }
func (*ReservedInstanceType) ProtoMessage()               {}
func (*ReservedInstanceType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ReservedInstanceType) GetOfferingClass() ReservedInstanceType_OfferingClass {
	if m != nil && m.OfferingClass != nil {
		return *m.OfferingClass
	}
	return ReservedInstanceType_STANDARD
}

func (m *ReservedInstanceType) GetPaymentOption() ReservedInstanceType_PaymentOption {
	if m != nil && m.PaymentOption != nil {
		return *m.PaymentOption
	}
	return ReservedInstanceType_ALL_UPFRONT
}

func (m *ReservedInstanceType) GetTermYears() uint32 {
	if m != nil && m.TermYears != nil {
		return *m.TermYears
	}
	return 0
}

// An amount of money, expressed in some currency.
type CurrencyAmount struct {
	// The currency in which the amount is expressed.
	// This is the ISO 4217 numeric code.
	// The default (840) is the USD currency code.
	//
	// We use the ISO 4217 standard so that in the future it would be easier to integrate
	// with JSR 354: Money and Currency API.
	Currency *int32 `protobuf:"varint,1,opt,name=currency,def=840" json:"currency,omitempty"`
	// The value, in the currency.
	// This should be non-negative, with 0 representing "free".
	Amount           *float64 `protobuf:"fixed64,2,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CurrencyAmount) Reset()                    { *m = CurrencyAmount{} }
func (m *CurrencyAmount) String() string            { return proto.CompactTextString(m) }
func (*CurrencyAmount) ProtoMessage()               {}
func (*CurrencyAmount) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

const Default_CurrencyAmount_Currency int32 = 840

func (m *CurrencyAmount) GetCurrency() int32 {
	if m != nil && m.Currency != nil {
		return *m.Currency
	}
	return Default_CurrencyAmount_Currency
}

func (m *CurrencyAmount) GetAmount() float64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

// The ReservedInstanceSpec describes a solution offered by the cloud provider that allows the customer
// to buy in advance a number of compute instances, for a discounted price. Usually those solutions
// have long terms like 1 or 3 years.
type ReservedInstanceSpec struct {
	// The type of the reserved instance.
	Type *ReservedInstanceType `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// The tenancy of the reserved instance.
	Tenancy *Tenancy `protobuf:"varint,2,opt,name=tenancy,enum=common_dto.Tenancy" json:"tenancy,omitempty"`
	// The operating system of the reserved instance.
	Os *OSType `protobuf:"varint,3,opt,name=os,enum=common_dto.OSType" json:"os,omitempty"`
	// The entity profile of the reserved instance is using, such as t2.large.
	Tier *EntityDTO `protobuf:"bytes,4,opt,name=tier" json:"tier,omitempty"`
	// The region of the reserved instance.
	Region       *EntityDTO `protobuf:"bytes,5,opt,name=region" json:"region,omitempty"`
	SizeFlexible *bool      `protobuf:"varint,6,opt,name=size_flexible,json=sizeFlexible" json:"size_flexible,omitempty"`
	// Whether the reserved instance can be applied to any operating system. If true,
	// the os attribute of the RI spec will be ignored.
	PlatformFlexible *bool  `protobuf:"varint,7,opt,name=platform_flexible,json=platformFlexible" json:"platform_flexible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReservedInstanceSpec) Reset()                    { *m = ReservedInstanceSpec{} }
func (m *ReservedInstanceSpec) String() string            { return proto.CompactTextString(m) }
func (*ReservedInstanceSpec) ProtoMessage()               {}
func (*ReservedInstanceSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ReservedInstanceSpec) GetType() *ReservedInstanceType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ReservedInstanceSpec) GetTenancy() Tenancy {
	if m != nil && m.Tenancy != nil {
		return *m.Tenancy
	}
	return Tenancy_DEFAULT
}

func (m *ReservedInstanceSpec) GetOs() OSType {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return OSType_UNKNOWN_OS
}

func (m *ReservedInstanceSpec) GetTier() *EntityDTO {
	if m != nil {
		return m.Tier
	}
	return nil
}

func (m *ReservedInstanceSpec) GetRegion() *EntityDTO {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *ReservedInstanceSpec) GetSizeFlexible() bool {
	if m != nil && m.SizeFlexible != nil {
		return *m.SizeFlexible
	}
	return false
}

func (m *ReservedInstanceSpec) GetPlatformFlexible() bool {
	if m != nil && m.PlatformFlexible != nil {
		return *m.PlatformFlexible
	}
	return false
}

func init() {
	proto.RegisterType((*ReservedInstanceType)(nil), "common_dto.ReservedInstanceType")
	proto.RegisterType((*CurrencyAmount)(nil), "common_dto.CurrencyAmount")
	proto.RegisterType((*ReservedInstanceSpec)(nil), "common_dto.ReservedInstanceSpec")
	proto.RegisterEnum("common_dto.DemandType", DemandType_name, DemandType_value)
	proto.RegisterEnum("common_dto.Tenancy", Tenancy_name, Tenancy_value)
	proto.RegisterEnum("common_dto.OSType", OSType_name, OSType_value)
	proto.RegisterEnum("common_dto.DatabaseEngine", DatabaseEngine_name, DatabaseEngine_value)
	proto.RegisterEnum("common_dto.LicenseModel", LicenseModel_name, LicenseModel_value)
	proto.RegisterEnum("common_dto.DeploymentType", DeploymentType_name, DeploymentType_value)
	proto.RegisterEnum("common_dto.DatabaseEdition", DatabaseEdition_name, DatabaseEdition_value)
	proto.RegisterEnum("common_dto.ReservedInstanceType_OfferingClass", ReservedInstanceType_OfferingClass_name, ReservedInstanceType_OfferingClass_value)
	proto.RegisterEnum("common_dto.ReservedInstanceType_PaymentOption", ReservedInstanceType_PaymentOption_name, ReservedInstanceType_PaymentOption_value)
}

func init() { proto.RegisterFile("CloudCostDTO.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 921 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x52, 0xe3, 0x36,
	0x14, 0xc5, 0x4e, 0x48, 0xc2, 0x85, 0x18, 0x55, 0xb0, 0x90, 0x65, 0xbb, 0x03, 0x4d, 0x5f, 0x68,
	0x3a, 0xa4, 0x1d, 0x66, 0x1f, 0x3a, 0x7d, 0x73, 0x6c, 0x01, 0x9e, 0x3a, 0x56, 0x90, 0xed, 0xcd,
	0xa6, 0x2f, 0x1e, 0x93, 0x08, 0x26, 0xd3, 0xc4, 0x4a, 0x6d, 0xb3, 0xd3, 0xf4, 0x73, 0xfa, 0x19,
	0xfd, 0x85, 0xfe, 0x41, 0xbf, 0xa6, 0x23, 0xd9, 0x09, 0x61, 0x97, 0x69, 0xfb, 0xe6, 0x7b, 0xee,
	0xb9, 0x47, 0xd7, 0xe7, 0x4a, 0x17, 0xb0, 0x35, 0x13, 0x8f, 0x13, 0x4b, 0x64, 0xb9, 0x1d, 0xd0,
	0xee, 0x22, 0x15, 0xb9, 0xc0, 0x30, 0x16, 0xf3, 0xb9, 0x48, 0xa2, 0x49, 0x2e, 0x4e, 0xf6, 0x2d,
	0xf5, 0xbd, 0x4e, 0xb6, 0xff, 0xd6, 0xe1, 0x90, 0xf1, 0x8c, 0xa7, 0x1f, 0xf9, 0xc4, 0x49, 0xb2,
	0x3c, 0x4e, 0xc6, 0x3c, 0x58, 0x2e, 0x38, 0x0e, 0xc1, 0x10, 0xf7, 0xf7, 0x3c, 0x9d, 0x26, 0x0f,
	0xd1, 0x78, 0x16, 0x67, 0x59, 0x4b, 0x3b, 0xd3, 0xce, 0x8d, 0xcb, 0x6e, 0xf7, 0x49, 0xae, 0xfb,
	0x52, 0x65, 0x97, 0x96, 0x65, 0x96, 0xac, 0x62, 0x4d, 0xb1, 0x19, 0x4a, 0xd9, 0x45, 0xbc, 0x9c,
	0xf3, 0x24, 0x8f, 0xc4, 0x22, 0x9f, 0x8a, 0xa4, 0xa5, 0xff, 0x4f, 0xd9, 0x41, 0x51, 0x46, 0x55,
	0x15, 0x6b, 0x2e, 0x36, 0x43, 0xfc, 0x16, 0x20, 0xe7, 0xe9, 0x3c, 0x5a, 0xf2, 0x38, 0xcd, 0x5a,
	0x95, 0x33, 0xed, 0xbc, 0xc9, 0x76, 0x24, 0x32, 0x92, 0x40, 0xbb, 0x0b, 0xcd, 0x67, 0x5d, 0xe1,
	0x3d, 0x68, 0xf8, 0x81, 0xe9, 0xd9, 0x26, 0xb3, 0x91, 0x86, 0xf7, 0x61, 0xd7, 0xa2, 0xde, 0x7b,
	0xc2, 0x02, 0xa7, 0xe7, 0x12, 0xa4, 0xb7, 0x09, 0x34, 0x9f, 0x1d, 0x27, 0x19, 0xa6, 0xeb, 0x46,
	0xe1, 0xe0, 0x8a, 0x51, 0x2f, 0x40, 0x1a, 0x3e, 0x80, 0xfd, 0x81, 0xc9, 0x02, 0xc7, 0x7c, 0x02,
	0x75, 0x6c, 0x00, 0x78, 0x74, 0x1d, 0x57, 0xda, 0x0e, 0x18, 0xd6, 0x63, 0x9a, 0xf2, 0x64, 0xbc,
	0x34, 0xe7, 0xe2, 0x31, 0xc9, 0xf1, 0x29, 0x34, 0xc6, 0x25, 0xa2, 0xfc, 0xdc, 0xfe, 0xb1, 0xf2,
	0xc3, 0xbb, 0xef, 0xd9, 0x1a, 0xc4, 0x47, 0x50, 0x8b, 0x15, 0x55, 0xf9, 0xa2, 0xb1, 0x32, 0x6a,
	0xff, 0xf5, 0xc2, 0x9c, 0xfc, 0x05, 0x1f, 0xe3, 0x77, 0x50, 0xcd, 0x97, 0x0b, 0xae, 0xd4, 0x76,
	0x2f, 0xcf, 0xfe, 0xcb, 0x46, 0xa6, 0xd8, 0xf8, 0x02, 0xea, 0x39, 0x4f, 0x62, 0xd9, 0x46, 0xe1,
	0xff, 0xc1, 0x66, 0x61, 0x50, 0xa4, 0xd8, 0x8a, 0x83, 0xdb, 0xa0, 0x8b, 0xc2, 0x56, 0xe3, 0x12,
	0x6f, 0x32, 0xa9, 0xaf, 0x44, 0x75, 0x91, 0xe1, 0x6f, 0xa0, 0x9a, 0x4f, 0x79, 0xda, 0xaa, 0xaa,
	0x46, 0x5e, 0x6d, 0xb2, 0x48, 0x92, 0x4f, 0xf3, 0xa5, 0x1d, 0x50, 0xa6, 0x28, 0xf8, 0x02, 0x6a,
	0x29, 0x7f, 0x90, 0xc3, 0xdf, 0xfe, 0x37, 0x72, 0x49, 0xc2, 0x5f, 0x43, 0x33, 0x9b, 0xfe, 0xce,
	0xa3, 0xfb, 0x19, 0xff, 0x6d, 0x7a, 0x37, 0xe3, 0xad, 0xda, 0x99, 0x76, 0xde, 0x60, 0x7b, 0x12,
	0xbc, 0x2a, 0x31, 0xfc, 0x2d, 0x7c, 0xb1, 0x98, 0xc5, 0xf9, 0xbd, 0x48, 0xe7, 0x4f, 0xc4, 0xba,
	0x22, 0xa2, 0x55, 0x62, 0x45, 0xee, 0x5c, 0x00, 0xd8, 0x7c, 0x1e, 0x27, 0x13, 0x75, 0xd5, 0x8b,
	0xf1, 0xfb, 0x61, 0x7f, 0x10, 0x38, 0xd4, 0x43, 0x9a, 0x9c, 0xa3, 0xe9, 0xba, 0xd4, 0x32, 0x55,
	0xac, 0x77, 0xbe, 0x83, 0x7a, 0x69, 0x09, 0xde, 0x85, 0xba, 0x4d, 0xae, 0xcc, 0xd0, 0x95, 0x97,
	0xa0, 0x09, 0x3b, 0x36, 0xb1, 0x1d, 0xcb, 0x0c, 0x88, 0x8d, 0x74, 0xdc, 0x80, 0xea, 0x0d, 0xf5,
	0x03, 0x54, 0xe9, 0xfc, 0xa9, 0x43, 0xad, 0xb0, 0x46, 0x6a, 0x85, 0xde, 0x4f, 0x1e, 0x1d, 0x7a,
	0x11, 0xf5, 0xd1, 0x16, 0xde, 0x81, 0x6d, 0xd7, 0xf1, 0xc2, 0x0f, 0x05, 0xdf, 0x0f, 0x7d, 0x82,
	0x2a, 0xf2, 0x8b, 0xdd, 0x10, 0x17, 0x55, 0xf1, 0x5b, 0x78, 0xad, 0xd2, 0xd1, 0xd0, 0x09, 0x6e,
	0x22, 0xff, 0xd6, 0x8d, 0x88, 0x17, 0x10, 0x36, 0x60, 0x8e, 0x4f, 0xd0, 0x36, 0x7e, 0x03, 0xc7,
	0x9f, 0xa4, 0xd7, 0xd7, 0xb8, 0x86, 0x8f, 0x00, 0x7f, 0x92, 0x1c, 0x92, 0x1e, 0xaa, 0xcb, 0x9e,
	0x87, 0x8e, 0x67, 0xd3, 0xa1, 0x8f, 0x0e, 0xe5, 0x01, 0x65, 0xf0, 0x82, 0xc6, 0x2b, 0xdc, 0x82,
	0xc3, 0xcf, 0xd2, 0x52, 0xe5, 0x08, 0x9f, 0xc2, 0x9b, 0xcf, 0x32, 0x1b, 0xbd, 0x1d, 0x63, 0x04,
	0x7b, 0x2b, 0x42, 0x6f, 0x44, 0x5d, 0xd4, 0xc2, 0x47, 0x60, 0xac, 0x10, 0x9f, 0xb0, 0xf7, 0x84,
	0xa1, 0xd7, 0x27, 0x7a, 0x43, 0xc3, 0x5f, 0x3e, 0x1d, 0x52, 0xe0, 0x51, 0x2f, 0x64, 0x7e, 0x80,
	0x4e, 0x64, 0xb6, 0xf3, 0x87, 0x06, 0x86, 0x1d, 0xe7, 0xf1, 0x5d, 0x9c, 0x71, 0x92, 0x3c, 0x4c,
	0x13, 0x2e, 0xff, 0xa0, 0x34, 0xb1, 0x70, 0xb0, 0x3f, 0xf2, 0x6f, 0x5d, 0xa4, 0x49, 0xbc, 0x6f,
	0x32, 0xc7, 0xb4, 0x7b, 0xc5, 0xeb, 0x1b, 0x50, 0x3f, 0xb8, 0x66, 0x44, 0x26, 0x2b, 0x18, 0xa0,
	0x46, 0x99, 0x69, 0xb9, 0x04, 0x55, 0xe5, 0xa4, 0xfc, 0x5b, 0xb7, 0x6c, 0x62, 0x1b, 0x1b, 0x50,
	0x33, 0x43, 0x46, 0x99, 0x89, 0x6a, 0xaa, 0x21, 0xf9, 0xbc, 0x55, 0x5c, 0x08, 0xd7, 0xf1, 0x21,
	0xa0, 0x02, 0xd8, 0x50, 0x6c, 0xa8, 0x93, 0xa9, 0x77, 0x4d, 0xd1, 0x4e, 0x67, 0x04, 0x7b, 0xee,
	0x74, 0xcc, 0x93, 0x8c, 0xf7, 0xc5, 0x84, 0xcf, 0xf0, 0x09, 0x1c, 0xf5, 0x98, 0xe3, 0x5d, 0x47,
	0x23, 0x1a, 0xb2, 0x48, 0x4e, 0xdb, 0x75, 0x2c, 0xe2, 0xf9, 0x04, 0x6d, 0x49, 0xb1, 0x32, 0x88,
	0x1c, 0xcf, 0x72, 0x43, 0x9b, 0xc8, 0xa5, 0x73, 0x0c, 0x07, 0x1e, 0x5d, 0xb1, 0x22, 0x46, 0x6e,
	0x43, 0x87, 0xc9, 0x6b, 0xd4, 0xb9, 0x00, 0xc3, 0xe6, 0x8b, 0x99, 0x50, 0xfb, 0x47, 0xdd, 0x21,
	0xd9, 0xbd, 0xe3, 0x5d, 0xbb, 0x24, 0x32, 0x7f, 0x46, 0x5b, 0x72, 0x79, 0xf5, 0x43, 0x37, 0x70,
	0x64, 0xa4, 0x75, 0x7e, 0x85, 0xfd, 0xb5, 0x5b, 0x93, 0xa9, 0xda, 0x56, 0x0d, 0xa8, 0x7a, 0xd4,
	0x93, 0x47, 0x1b, 0x00, 0x1b, 0x33, 0xd2, 0x9e, 0xed, 0x3d, 0x5d, 0xfe, 0xf6, 0x2a, 0x92, 0xf4,
	0xca, 0x26, 0x10, 0x0c, 0x29, 0xaa, 0xe2, 0x3a, 0x54, 0xe4, 0xf4, 0xf7, 0xa4, 0xd3, 0xe4, 0xc3,
	0x80, 0x11, 0xdf, 0x47, 0xcd, 0xde, 0x57, 0x70, 0x3a, 0x16, 0xf3, 0xee, 0xc7, 0x79, 0xfe, 0x98,
	0xde, 0x89, 0xee, 0xea, 0x79, 0x75, 0xb3, 0xc9, 0x2f, 0xe5, 0x4b, 0xfe, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0x99, 0xa8, 0x0d, 0x89, 0x06, 0x00, 0x00,
}
