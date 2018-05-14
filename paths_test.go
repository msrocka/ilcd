package ilcd

import "testing"

func TestFindUUID(t *testing.T) {
	if FindUUID("no/uuid") != "" {
		t.Fatal("no/uuid has no UUID")
	}
	uuid := FindUUID("ILCD/unitgroups/93a60a57-a3c8-11da-a746-0800200c9a66.xml")
	if uuid != "93a60a57-a3c8-11da-a746-0800200c9a66" {
		t.Fatal("Did not extracted UUID")
	}
	uuid = FindUUID("ILCD/unitgroups/93A60A57-A3C8-11DA-A746-0800200C9A66.xml")
	if uuid != "93A60A57-A3C8-11DA-A746-0800200C9A66" {
		t.Fatal("Did not extracted UUID")
	}
}
