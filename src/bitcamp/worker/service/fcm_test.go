package service

import "testing"

func TestBroadcastMessage(t *testing.T) {
	err := BroadcastMessage(3, "fuck you", "455-555-5555", true)
	if err != nil{
		t.Fatal(err)
		t.FailNow()
	}
}
