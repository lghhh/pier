package rulemgr

import (
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/pier/internal/peermgr"
	peerproto "github.com/meshplus/pier/internal/peermgr/proto"
)

func (rm *RuleMgr) handleRule(net network.Stream, msg *peerproto.Message) {
	data := msg.Payload.Data
	rule := &Rule{}
	if err := json.Unmarshal(data, rule); err != nil {
		logger.Error(err)
		return
	}
	ok := true
	err := rm.Ledger.SetCode(types.String2Address(rule.Address), rule.Code)
	if err != nil {
		ok = false
		logger.Error(err)
	}

	res := RuleResponse{
		Ok: ok,
	}

	ackData, err := json.Marshal(res)
	if err != nil {
		logger.Error(err)
		return
	}

	ackMsg := peermgr.Message(peerproto.Message_ACK, true, ackData)

	err = rm.PeerManager.AsyncSendWithStream(net, ackMsg)
	if err != nil {
		logger.Error(err)
	}
}
