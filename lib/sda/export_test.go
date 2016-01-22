package sda

import (
	"github.com/dedis/cothority/lib/network"
)

// Export some private functions of Host for testing

func (h *Host) AddPendingTreeMarshal(tm *TreeMarshal) {
	h.addPendingTreeMarshal(tm)
}

func (h *Host) CheckPendingTreeMarshal(el *EntityList) {
	h.checkPendingTreeMarshal(el)
}

func (h *Host) SendSDAData(id *network.Entity, msg *SDAData) error {
	return h.sendSDAData(id, msg)
}

func (h *Host) Receive() network.NetworkMessage {
	return h.receive()
}

func (h *Host) ProtocolInstantiate(tok *Token, tn *TreeNode) (ProtocolInstance, error) {
	return h.overlay.protocolInstantiate(tok, tn)
}
