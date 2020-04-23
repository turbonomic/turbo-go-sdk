package probe

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg"
)

func TestNewProbeBuilder(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	probeUICat := "Cloud"

	builder := NewProbeBuilder(probeType, probeCat, probeUICat)

	_, _ = builder.Create()

	assert.EqualValues(t, pkg.DEFAULT_FULL_DISCOVERY_IN_SECS,
		builder.probeConf.discoveryMetadata.GetFullRediscoveryIntervalSeconds())
	assert.EqualValues(t, pkg.DISCOVERY_NOT_SUPPORTED,
		builder.probeConf.discoveryMetadata.GetIncrementalRediscoveryIntervalSeconds())
	assert.EqualValues(t, pkg.DISCOVERY_NOT_SUPPORTED,
		builder.probeConf.discoveryMetadata.GetPerformanceRediscoveryIntervalSeconds())
}

func TestNewProbeBuilderWithDiscoveryMetadata(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	probeUICat := "Cloud"

	table := []struct {
		full        int32
		incremental int32
		performance int32
	}{
		{full: 0, incremental: 0, performance: 0},
		{full: -1, incremental: -1, performance: -1},
		{full: 60, incremental: 120, performance: 300},
		{full: 30, incremental: 20, performance: 30},
		{full: 30},
		{full: -1},
		{full: 1200},
		{incremental: 20, performance: 30},
		{incremental: 60, performance: 60},
	}
	for _, item := range table {
		builder := NewProbeBuilder(probeType, probeCat, probeUICat)
		builder.WithDiscoveryOptions(IncrementalRediscoveryIntervalSecondsOption(item.incremental),
			FullRediscoveryIntervalSecondsOption(item.full),
			PerformanceRediscoveryIntervalSecondsOption(item.performance))

		_, _ = builder.Create()

		dm := builder.probeConf.discoveryMetadata
		checkDiscoveryMetadata(t, item.full, dm, pkg.FULL_DISCOVERY)
		checkDiscoveryMetadata(t, item.incremental, dm, pkg.INCREMENTAL_DISCOVERY)
		checkDiscoveryMetadata(t, item.performance, dm, pkg.PERFORMANCE_DISCOVERY)
	}
}

func TestNewProbeBuilderWithoutRegistrationClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	probeUICat := "Cloud"

	probe, err := createProbe(probeType, probeCat, probeUICat, nil, "", nil)

	expectedErr := ErrorInvalidRegistrationClient()
	if !reflect.DeepEqual(expectedErr, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", expectedErr, err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}

func TestNewProbeBuilderWithoutDiscoveryClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	probeUICat := "Cloud"
	targetID := "T1"

	registrationClient := &TestProbeRegistrationClient{}

	probe, err := createProbe(probeType, probeCat, probeUICat, registrationClient, targetID, nil)

	expectedErr := ErrorUndefinedDiscoveryClient()
	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", expectedErr, err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}

func TestNewProbeBuilderWithRegistrationAndDiscoveryClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	probeUICat := "Cloud"
	targetId := "T1"

	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	builder := NewProbeBuilder(probeType, probeCat, probeUICat)
	builder.RegisteredBy(registrationClient)
	builder.DiscoversTarget(targetId, discoveryClient)
	probe, err := builder.Create()

	if !reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, err)
	}

	if !reflect.DeepEqual(registrationClient.GetSupplyChainDefinition(),
		probe.RegistrationClient.GetSupplyChainDefinition()) {
		t.Errorf("\nExpected %+v, \ngot      %+v",
			registrationClient, probe.RegistrationClient)
	}
	if !reflect.DeepEqual(registrationClient.GetAccountDefinition(),
		probe.RegistrationClient.GetAccountDefinition()) {
		t.Errorf("\nExpected %+v, \ngot      %+v",
			registrationClient, probe.RegistrationClient)
	}

	dc := probe.GetTurboDiscoveryClient()
	if !reflect.DeepEqual(discoveryClient, dc) {
		t.Errorf("\nExpected %+v, \ngot      %+v", discoveryClient, dc)
	}
}

func TestNewProbeBuilderWithActionClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	probeUICat := "Cloud"
	targetId := "T1"
	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	actionClient := &TestProbeActionClient{}
	builder := NewProbeBuilder(probeType, probeCat, probeUICat)

	if registrationClient != nil {
		builder.RegisteredBy(registrationClient)
	}

	if targetId != "" || discoveryClient != nil {
		builder.DiscoversTarget(targetId, discoveryClient)
	}
	builder.ExecutesActionsBy(actionClient)
	probe, err := builder.Create()

	if !reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, err)
	}
	if !reflect.DeepEqual(actionClient, probe.ActionClient) {
		t.Errorf("\nExpected %+v, \ngot      %+v", actionClient, probe.ActionClient)
	}
}

func TestNewProbeBuilderWithInvalidActionClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	probeUICat := "Cloud"
	targetId := "T1"
	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	var actionClient TurboActionExecutorClient //:= &TestProbeActionClient{}
	builder := NewProbeBuilder(probeType, probeCat, probeUICat)

	if registrationClient != nil {
		builder.RegisteredBy(registrationClient)
	}

	if targetId != "" || discoveryClient != nil {
		builder.DiscoversTarget(targetId, discoveryClient)
	}
	builder.ExecutesActionsBy(actionClient)
	probe, err := builder.Create()

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidActionClient(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}

func TestNewProbeBuilderInvalidTargetId(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	probeUICat := "Cloud"

	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	probe, err := createProbe(probeType, probeCat, probeUICat, registrationClient, "", discoveryClient)
	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidTargetIdentifier(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}

}

func TestNewProbeBuilderInvalidProbeType(t *testing.T) {
	probeType := ""
	probeCat := "Cloud"
	probeUICat := "Cloud"

	probe, err := createProbe(probeType, probeCat, probeUICat, nil, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidProbeType(), err)
	}
	var expected1 *TurboProbe
	if !reflect.DeepEqual(expected1, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}

	probeType = "Type1"
	probeCat = ""
	probeUICat = "Cloud"

	probe, err = createProbe(probeType, probeCat, probeUICat, nil, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidProbeCategory(), err)
	}
	var expected2 *TurboProbe
	if !reflect.DeepEqual(expected2, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}

	probeType = "Type1"
	probeCat = "Cloud"
	probeUICat = ""

	probe, err = createProbe(probeType, probeCat, probeUICat, nil, "", nil)

	if !reflect.DeepEqual(ErrorInvalidProbeUICategory(), err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidProbeUICategory(), err)
	}
	var expected3 *TurboProbe
	if !reflect.DeepEqual(expected3, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}

func createProbe(probeType, probeCat, probeUICat string,
	registrationClient TurboRegistrationClient,
	targetId string, discoveryClient TurboDiscoveryClient) (*TurboProbe, error) {

	builder := NewProbeBuilder(probeType, probeCat, probeUICat)

	if registrationClient != nil {
		builder.RegisteredBy(registrationClient)
	}

	if targetId != "" || discoveryClient != nil {
		builder.DiscoversTarget(targetId, discoveryClient)
	}

	return builder.Create()
}
