package probe

import (
	"testing"
	"reflect"
)


func TestNewProbeBuilderWithoutRegistrationClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"

	probe, err := createProbe(probeType, probeCat, nil, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidRegistrationClient(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}


func TestNewProbeBuilderWithoutDiscoveryClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	targetID := "T1"

	registrationClient :=  &TestProbeRegistrationClient{}

	probe, err := createProbe(probeType, probeCat, registrationClient, targetID, nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorUndefinedDiscoveryClient(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}


func TestNewProbeBuilderWithRegistrationAndDiscoveryClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	targetId := "T1"

	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	builder := NewProbeBuilder(probeType, probeCat)
	builder.RegisteredBy(registrationClient)
	builder.DiscoversTarget(targetId, discoveryClient)
	probe, err := builder.Create()

	if !reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, err)
	}

	if !reflect.DeepEqual(registrationClient, probe.RegistrationClient) {
		t.Errorf("\nExpected %+v, \ngot      %+v", registrationClient, probe.RegistrationClient)
	}

	dc := probe.getDiscoveryClient(targetId)
	if !reflect.DeepEqual(discoveryClient, dc) {
		t.Errorf("\nExpected %+v, \ngot      %+v", discoveryClient, dc)
	}
}

func TestNewProbeBuilderWithActionClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	targetId := "T1"
	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	actionClient := &TestProbeActionClient{}
	builder := NewProbeBuilder(probeType, probeCat)

	if (registrationClient != nil) {
		builder.RegisteredBy(registrationClient)
	}

	if (targetId != "" || discoveryClient != nil) {
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
	targetId := "T1"
	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	var actionClient TurboActionExecutorClient	//:= &TestProbeActionClient{}
	builder := NewProbeBuilder(probeType, probeCat)

	if (registrationClient != nil) {
		builder.RegisteredBy(registrationClient)
	}

	if (targetId != "" || discoveryClient != nil) {
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

	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	probe, err := createProbe(probeType, probeCat, registrationClient, "", discoveryClient)
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

	probe, err := createProbe(probeType, probeCat, nil, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidProbeType(), err)
	}
	var expected1 *TurboProbe
	if !reflect.DeepEqual(expected1, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}

	probeType = "Type1"
	probeCat = ""

	probe, err = createProbe(probeType, probeCat, nil, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidProbeCategory(), err)
	}
	var expected2 *TurboProbe
	if !reflect.DeepEqual(expected2, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}


func createProbe(probeType, probeCat string,
			registrationClient TurboRegistrationClient,
			targetId string, discoveryClient TurboDiscoveryClient) (*TurboProbe, error) {

	builder := NewProbeBuilder(probeType, probeCat)

	if (registrationClient != nil) {
		builder.RegisteredBy(registrationClient)
	}

	if (targetId != "" || discoveryClient != nil) {
		builder.DiscoversTarget(targetId, discoveryClient)
	}

	return builder.Create()
}


