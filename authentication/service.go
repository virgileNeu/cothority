// Package authentication centralizes authentication for services like
// skipchain, cisc, pop and others who need it. This first version simply
// holds a list of all authentication tokens that are allowed to do
// things with this conode. A later version will include policy-darcs to
// correctly authenticate using modern tools.
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
// darc can point to specific user darcs and then each user darc can
// update his access rights without having to contact the admin darc.
//
// Darcs have a 'Data' field that will be used to reflect the policy
// this darc is responsible. If the user wants to update an existing darc
// and replace it with another `Data`, then the system will refuse
// that darc.
//
// Policy definition
//
// When starting up, the authentication service sets up a new Darc with the
// public key of the conode and the policy "*", stored in the 'Data'
// field of the darc. This darc is the root darc and
// responsible for all services and all methods. Everybody having the private
// key of the conode is thus allowed to sign for any action of any service.
//
// For more fine-grained authentication, the client can ask to store a new
// darc with a 'Data' field set to the name of the service. This allows
// a client to give access to an external user only to a given service.
//
// For more fine-grained authentication, the 'Data' can hold the name
// of the service separated with a '.' from the method. For example to protect
// the `CreateSkipchain` method of the identity service, the client would have to
// create a darc with a 'Data' of 'Identity.CreateSkipchain'.
//
// How a service should use it
//
// Each service that wants to use authentication can include the following
// structure in its api-messages sent from the client:
//
//  type Auth struct{
//    Signature darc.Signature
//  }
//
// on the service-side, the skipchain-method has to call
//
//  authentication.Verify(s onet.Service, service, method string, auth Auth) error
package authentication

import (
	"errors"
	"log"

	"github.com/dedis/onchain-secrets/darc"
	"github.com/dedis/onet"
)

const ServiceName = "authentication"

// Service for the auhtentication holds all darcs that define who is allowed
// to do what. It interacts with the users through the api.
type Service struct {
	*onet.ServiceProcessor
}

// Auth can be included in an api-client-message to send an authentication to
// the service.
type Auth struct {
	Signature darc.Signature
}

// Verify checks the signature on the message given by "service.method"
// by going through all darcs and looking if one of the darc-policies matches the
// message and if the signature is correct.
func Verify(s onet.Context, service, method string, auth Auth) error {
	authService := s.Service(ServiceName).(*Service)
	log.Print(authService)
	return nil
}

// GetPolicy returns the latest version of the chosen policy with the ID
// given. If the ID == null, the latest version of the
// basic policy is returned.
func (s *Service) GetPolicy(req *GetPolicy) (*GetPolicyReply, error) {
	return nil, errors.New("not yet implemented")
}

//
// UpdatePolicy updates an existing policy. Following the policy-library,
// it needs to be signed by one of the admin users.
func (s *Service) UpdatePolicy(req *UpdatePolicy) (*UpdatePolicyReply, error) {
	return nil, errors.New("not yet implemented")
}

// UpdatePolicyPIN can be used in case the private key is not available,
// but if the user has access to the logs of the server. On the first
// call the PIN == "", and the server will print a 6-digit PIN in the log
// files. When he receives the policy and the correct PIN, the server will
// auto-sign the policy using his private key and add it to the policy-list.
func (s *Service) UpdatePolicyPIN(req *UpdatePolicyPIN) (*UpdatePolicyPINReply, error) {
	return nil, errors.New("not yet implemented")
}
