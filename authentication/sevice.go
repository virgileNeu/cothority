// Package authentication centralizes authentication for services like
// skipchain, cisc, pop and others who need it. This first version simply
// holds a list of all authentication tokens that are allowed to do
// things with this conode. A later version will include policy-darcs to
// correctly authenticate using modern tools.
// It is based on the policy
// library from Sandra Siby which is based on the darc-library from Linus.
//
// API for the clients
//
// The API of the services wanting to authenticate need to add a new field in
// their messages that is of type `policy.Signature` and is the first field
// available. In the protobuf definition this field is optional, as under some
// circumstances it might be nil.
//
// Usage of Darcs
//
// Distributed Access Rights Control is a data structure for handling access
// rights to given resources. A darc has a list of Owners who are allowed
// to propose a new version of the darc. It also has a list of Users who are
// allowed to sign as authentication for the execution of an action defined
// by the darc.
//
// Both Owners and Users can point to other darcs, so that a general admin
// darc can point to speicific user darcs and then each user darc can
// update his access rights without having to contact the admin darc.
//
// On the API side, the following methods are available:
//   // GetPolicy returns the latest version of the chosen policy with the ID
//   // given. If the ID == null, the basic policy is returned.
//   GetPolicy(ID darc.ID)(*darc.Policy, error)
//
//   // UpdatePolicy updates an existing policy. Following the policy-library,
//   // it needs to be signed by one of the admin users.
//   UpdatePolicy(newPolicy darc.Policy) error
//
//   // UpdatePolicyPIN can be used in case the private key is not available,
//   // but if the user has access to the logs of the server. On the first
//   // call the PIN == "", and the server will print a 6-digit PIN in the log
//   // files. When he receives the policy and the correct PIN, the server will
//   // auto-sign the policy using his private key and add it to the policy-list.
//   UpdatePolicyPIN(newPolicy policy, PIN string) error
//
//   // AddPolicy can be used to add a new policy to the system that will be
//   // referenced later with UpdatePolicy. For a new policy, it must be signed
//   // by a user of the root-policy.
//   AddPolicy(newPolicy policy, signature Signature)
package authentication
