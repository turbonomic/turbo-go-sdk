package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlowBuildSuccess(t *testing.T) {
	assert := assert.New(t)
	builder := NewFlowDTOBuilder()
	builder.Source("10.1.1.1").
		Destination("10.1.2.1", 8080).
		Protocol(TCP).
		FlowAmount(1024).
		Latency(100).Received(512).Transmitted(512)
	dto, err := builder.Create()
	assert.True(err == nil)
	assert.True(dto != nil)
	assert.EqualValues("10.1.1.1", *dto.SourceEntityIdentityData.IpAddress)
	assert.EqualValues(0, *dto.SourceEntityIdentityData.Port)
	assert.EqualValues("10.1.2.1", *dto.DestEntityIdentityData.IpAddress)
	assert.EqualValues(8080, *dto.DestEntityIdentityData.Port)
	assert.EqualValues(TCP, *dto.Protocol)
	assert.EqualValues(1024, *dto.FlowAmount)
	assert.EqualValues(100, *dto.Latency)
	assert.EqualValues(512, *dto.ReceivedAmount)
	assert.EqualValues(512, *dto.TransmittedAmount)
}

func TestFlowBuildFailure(t *testing.T) {
	assert := assert.New(t)
	builder := NewFlowDTOBuilder()
	builder.Source("10.1.1.1").
		Destination("10.1.2.1", 8080).
		Protocol(1000).
		FlowAmount(1024).
		Latency(100).Received(512).Transmitted(512)
	_, err := builder.Create()
	assert.True(err != nil)
}
