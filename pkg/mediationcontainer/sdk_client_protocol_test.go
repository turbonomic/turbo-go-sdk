package mediationcontainer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateSdkClientProtocolHandler(t *testing.T) {
	communicationBindingChannel := "foo"
	sdkClientProtocol := CreateSdkClientProtocolHandler(nil, "1.0", communicationBindingChannel, nil)
	assert.Equal(t, time.Duration(300000000000), sdkClientProtocol.registrationResponseTimeout)
	assert.Equal(t, false, sdkClientProtocol.restartOnRegistrationTimeout)
}

func TestCreateSdkClientProtocolHandlerInvalidTimeout(t *testing.T) {
	communicationBindingChannel := "foo"
	sdkProtocolConfig := &SdkProtocolConfig{
		RegistrationTimeoutSec:       10,
		RestartOnRegistrationTimeout: true,
	}
	sdkClientProtocol := CreateSdkClientProtocolHandler(nil, "1.0", communicationBindingChannel, sdkProtocolConfig)
	assert.Equal(t, time.Duration(300000000000), sdkClientProtocol.registrationResponseTimeout)
	assert.Equal(t, true, sdkClientProtocol.restartOnRegistrationTimeout)
}

func TestTimeoutRead(t *testing.T) {
	du := time.Second * 3
	ch := make(chan *ParsedMessage)

	_, err := timeOutRead("test", du, ch)
	if err == nil {
		t.Error("Read should time out with error.")
	} else {
		fmt.Println(err.Error())
	}
}

func TestTimeoutRead2(t *testing.T) {
	du := time.Second * 5
	ch := make(chan *ParsedMessage)
	msg := &ParsedMessage{}

	go func() {
		ch <- msg
	}()

	msg2, err := timeOutRead("test", du, ch)
	if err != nil {
		t.Errorf("Read should success without error: %v", err)
	}

	if msg != msg2 {
		fmt.Println("received msg is different.")
	}
}

// TestCommunicationBindingChannel tests the population of the communication binding channel and ensures the produced
// container info has it.
func TestCommunicationBindingChannel(t *testing.T) {
	for _, communicationBindingChannel := range []string{"foo", ""} {
		sdkClientProtocol := CreateSdkClientProtocolHandler(nil, "1.0", communicationBindingChannel, nil)
		containInfo, err := sdkClientProtocol.MakeContainerInfo()
		if err != nil {
			t.Fatalf("Error making container info: %v", err)
		}
		if *containInfo.CommunicationBindingChannel != communicationBindingChannel {
			t.Fatalf("Binding channel in the container info %v is not the same as the original %v", containInfo,
				communicationBindingChannel)
		}
	}
}
