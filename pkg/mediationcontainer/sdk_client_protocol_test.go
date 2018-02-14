package mediationcontainer

import (
	"fmt"
	"testing"
	"time"
)

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
