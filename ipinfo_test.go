package ipinfo

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	info, err := Query("")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", info)
}

func TestQuery(t *testing.T) {
	info, err := Query("8.8.8.8")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", info)
}
