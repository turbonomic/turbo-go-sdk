package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
	"testing"
)

//
//// Tests that the argument passed to SetContainerInfo is assigned to this.clientMessage.ContainerInfo
//func TestSetContainerInfo(t *testing.T) {
//	assert := assert.New(t)
//	clientMsg := new(proto.MediationClientMessage)
//	cmbuilder := &ClientMessageBuilder{
//		clientMessage: clientMsg,
//	}
//	cInfo := new(proto.ContainerInfo)
//	cmb := cmbuilder.SetContainerInfo(cInfo)
//	if assert.NotEqual((*ContainerInfo)(nil), cmb.clientMessage.ContainerInfo) {
//		assert.Equal(cInfo, cmb.clientMessage.ContainerInfo)
//		assert.Equal(*cInfo, *cmb.clientMessage.ContainerInfo)
//	}
//}

// Tests that the method NewAccountDefEntryBuilder creates an AccountDefEntry struct with member
// variables set equal to the pointers to the arguments passed to it
func TestNewAccountDefEntryBuilder(t *testing.T) {
	assert := assert.New(t)
	name := rand.String(6)
	displayName := rand.String(7)
	description := rand.String(8)
	verificationRegex := rand.String(9)
	entryType := false //AccountDefEntry_OPTIONAL
	isSecret := true
	acctDefEntryBuilder := NewAccountDefEntryBuilder(name, displayName, description, verificationRegex, entryType, isSecret)
	acctDef := acctDefEntryBuilder.accountDefEntry
	customAcctDef := acctDef.GetCustomDefinition()
	if assert.NotEqual((*proto.AccountDefEntry)(nil), acctDef) {
		assert.Equal(name, *customAcctDef.Name)
		assert.Equal(displayName, *customAcctDef.DisplayName)
		assert.Equal(description, *customAcctDef.Description)
		assert.Equal(verificationRegex, *customAcctDef.VerificationRegex)
		assert.Equal(entryType, *acctDef.Mandatory)
		assert.Equal(isSecret, *customAcctDef.IsSecret)
	}
}

// Tests that NewProbeInfoBuilder creates a ProbeInfo struct with member variables set to the
// arguments passed to it and creates a ProbeInfoBuilder struct with its probeInfo variable set
// to the new ProbeInfo struct containing the passed arguments
func TestNewProbeInfoBuilder(t *testing.T) {
	assert := assert.New(t)
	probeType := "ProbeType1"    //rand.String(6)
	probeCat := "ProbeCategory1" //rand.String(7)
	var supplyCS []*proto.TemplateDTO
	var acctDef []*proto.AccountDefEntry
	probeInfoBldr := NewProbeInfoBuilder(probeType, probeCat, supplyCS, acctDef)
	assert.Equal(probeType, *probeInfoBldr.probeInfo.ProbeType)
}

// Tests that the method Create() returns a pointer to the ProbeInfo struct in the object that Create
// called on
func TestProbeInfo_Create(t *testing.T) {
	assert := assert.New(t)
	probeinfo := new(proto.ProbeInfo)
	builder := &ProbeInfoBuilder{
		probeInfo: probeinfo,
	}
	pi := builder.Create()
	assert.Equal(probeinfo, pi)
	assert.Equal(*probeinfo, *pi)
}

func TestProbeInfoBuilder(t *testing.T) {

	table := []struct {
		probeType     string
		probeCategory string
		full          int32
		incremental   int32
		performance   int32
	}{
		{"Type1", "Category1", 200, 10, 10},
		{"Type2", "Category2", 200, 10, 10},
		{"Type3", "Category3", -1, -1, -1},
		{probeType: "Type4", probeCategory: "Category4"},
	}

	for _, item := range table {
		builder := NewBasicProbeInfoBuilder(item.probeType, item.probeCategory)
		builder.WithFullDiscoveryInterval(item.full)
		builder.WithIncrementalDiscoveryInterval(item.incremental)
		builder.WithPerformanceDiscoveryInterval(item.performance)
		probeInfo := builder.Create()
		assert.Equal(t, item.probeType, probeInfo.GetProbeType())
		assert.Equal(t, item.probeCategory, probeInfo.GetProbeCategory())
		if item.full < pkg.DEFAULT_MIN_DISCOVERY_IN_SECS {
			assert.EqualValues(t, pkg.DEFAULT_MIN_DISCOVERY_IN_SECS, probeInfo.GetFullRediscoveryIntervalSeconds())
		} else {
			assert.EqualValues(t, item.full, probeInfo.GetFullRediscoveryIntervalSeconds())
		}
		if item.incremental <= pkg.DISCOVERY_NOT_SUPPORTED {
			assert.EqualValues(t, 0, probeInfo.GetIncrementalRediscoveryIntervalSeconds())
		} else {
			assert.EqualValues(t, item.incremental, probeInfo.GetIncrementalRediscoveryIntervalSeconds())
		}
		if item.performance <= pkg.DISCOVERY_NOT_SUPPORTED {
			assert.EqualValues(t, 0, probeInfo.GetPerformanceRediscoveryIntervalSeconds())
		} else {
			assert.EqualValues(t, item.performance, probeInfo.GetPerformanceRediscoveryIntervalSeconds())
		}
		assert.Nil(t, probeInfo.GetActionPolicy())
		assert.Nil(t, probeInfo.GetEntityMetadata())
	}
}

func TestActionPolicyBuilder(t *testing.T) {
	var supported, notSupported proto.ActionPolicyDTO_ActionCapability
	supported = proto.ActionPolicyDTO_SUPPORTED
	notSupported = proto.ActionPolicyDTO_NOT_SUPPORTED

	expectedMap := make(map[proto.EntityDTO_EntityType]map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability)
	expectedMap[proto.EntityDTO_VIRTUAL_MACHINE] = map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability{
		proto.ActionItemDTO_MOVE:      supported,
		proto.ActionItemDTO_RESIZE:    supported,
		proto.ActionItemDTO_PROVISION: notSupported,
	}
	expectedMap[proto.EntityDTO_CONTAINER_POD] = map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability{
		proto.ActionItemDTO_MOVE:   supported,
		proto.ActionItemDTO_RESIZE: notSupported,
	}
	expectedMap[proto.EntityDTO_CONTAINER] = map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability{
		proto.ActionItemDTO_MOVE:   notSupported,
		proto.ActionItemDTO_RESIZE: supported,
	}

	builder := NewActionPolicyBuilder()
	for eType, itemMap := range expectedMap {
		for aType, item := range itemMap {
			builder.WithEntityActions(eType, aType, item)
		}
	}

	actionPolicies := builder.Create()
	for _, actionPolicy := range actionPolicies {
		policies := actionPolicy.PolicyElement
		expectedPolicies, exists := expectedMap[*actionPolicy.EntityType]
		assert.True(t, exists)
		for _, policyElement := range policies {
			assert.EqualValues(t, expectedPolicies[*policyElement.ActionType], *policyElement.ActionCapability)
		}
	}
}
