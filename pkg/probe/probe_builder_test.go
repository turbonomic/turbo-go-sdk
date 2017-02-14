package probe

import (
	"testing"
	"reflect"
)


func TestNewProbeBuilderWithoutRegistrationClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"

	probe, err := createProbe(probeType, probeCat, nil, "", nil)	//builder.Create()

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

	registrationClient :=  &TestProbeRegistrationClient{}

	probe, err := createProbe(probeType, probeCat, registrationClient, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorUndefineddDiscoveryClient(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}


func TestNewProbeBuilderInvalidTargetId(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	targetId := ""

	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}

	probe, err := createProbe(probeType, probeCat, registrationClient, targetId, discoveryClient)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidDiscoveryClient(targetId), err)
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


