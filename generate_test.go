package cryptorandomstring

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	if got, err := WithLength(10).WithKind("base64").Generate(); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("got: %v\n", got)
	}
}
