package contract

import "testing"

func TestHelloWorld(t *testing.T) {
	ClaimService.GetClaim("hhh")
	t.Log("hello world")
}
