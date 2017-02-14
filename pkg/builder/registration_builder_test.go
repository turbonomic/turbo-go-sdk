package builder

import (
	"github.com/stretchr/testify/assert"
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
