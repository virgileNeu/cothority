package authentication

import (
	"github.com/dedis/onchain-secrets/darc"
	"github.com/dedis/onet/network"
)

func init() {
	network.RegisterMessages(
		GetPolicy{}, GetPolicyReply{},
		UpdatePolicy{}, UpdatePolicyReply{},
		UpdatePolicyPIN{}, UpdatePolicyPINReply{},
		AddPolicy{}, AddPolicyReply{},
	)
}

// PROTOSTART
//
// option java_package = "ch.epfl.dedis.proto";
// option java_outer_classname = "AuthProto";

// ***
// These are the messages used in the API-calls
// ***

// GetPolicy can be called from a client to get the latest version of a given
// policy. The system will only return an exact match.
type GetPolicy struct {
	Rule string
}

// GetPolicyReply contains the latest known version of the darc in the
// authentication service. If the Rule has not been found, then the closes
// matching rule will be returned.
type GetPolicyReply struct {
	Latest *darc.Darc
}

// UpdatePolicy proposes a new darc for a given policy. If the proposed darc
// is new (version == 0), then it must be signed by the latest version of the
// root-darc.
// If an updated darc is given, it must have the same 'Data' field as the previous
// darc, else it will be rejected by the system.
type UpdatePolicy struct {
	NewDarc *darc.Darc
}

// UpdatePolicyReply doesn't return anything and is empty.
type UpdatePolicyReply struct{}

// UpdatePolicyPIN takes a new policy and a PIN. If the PIN is empty, then the
// authentication service will print a PIN to the server logs, then the client
// must read this out, and resend the data using that PIN.
// A darc submitted with this method must be a new darc (version = 0), but it
// can overwrite an existing darc!
type UpdatePolicyPIN struct {
	NewDarc *darc.Darc
}
type UpdatePolicyPINReply struct{}
type AddPolicy struct{}
type AddPolicyReply struct{}
