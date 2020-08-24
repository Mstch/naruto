package util

import "testing"

func TestInt32(t *testing.T) {
	if v := BytesToInt32(Int32ToBytes(-100)); v != -100 {
		print(v)
		t.Fail()
	}
	if v := BytesToInt32(Int32ToBytes(100)); v != 100 {
		print(v)

		t.Fail()
	}
	if v := BytesToInt32(Int32ToBytes(0)); v != 0 {
		print(v)

		t.Fail()
	}

}
