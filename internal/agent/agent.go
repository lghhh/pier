package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
	rpcx "github.com/meshplus/go-bitxhub-client"
	"github.com/meshplus/pier/internal/repo"
)

// agent is responsible for interacting with bitxhub
var _ Agent = (*BxhAgent)(nil)

// BxhAgent represents the necessary data for interacting with bitxhub
type BxhAgent struct {
	client     rpcx.Client
	from       types.Address
	addr       string
	validators string
}

// New create an instance of BxhAgent given the client of bitxhub and
// the appchain id and some configuration of bitxhub
func New(client rpcx.Client, pierID types.Address, bitxhub repo.Relay) (*BxhAgent, error) {
	return &BxhAgent{
		client:     client,
		from:       pierID,
		addr:       bitxhub.Addr,
		validators: strings.Join(bitxhub.Validators, ","),
	}, nil
}

func (agent *BxhAgent) Stop() error {
	return agent.client.Stop()
}

// Appchain implements Agent
func (agent *BxhAgent) Appchain() (*rpcx.Appchain, error) {
	receipt, err := agent.client.InvokeBVMContract(rpcx.AppchainMgrContractAddr, "Appchain")
	if err != nil {
		return nil, err
	}

	appchain := &rpcx.Appchain{}
	if receipt.Status == pb.Receipt_FAILED {
		return nil, fmt.Errorf("receipt: %s", receipt.Ret)
	}

	if err := json.Unmarshal(receipt.Ret, appchain); err != nil {
		return nil, fmt.Errorf("unmarshal appchain from bitxhub: %w", err)
	}

	return appchain, nil
}

func (agent *BxhAgent) GetInterchainMeta() (*rpcx.Interchain, error) {
	receipt, err := agent.client.InvokeBVMContract(rpcx.InterchainContractAddr, "Interchain")
	if err != nil {
		return nil, err
	}

	if !receipt.IsSuccess() {
		return nil, fmt.Errorf("receipt: %s", receipt.Ret)
	}

	ret := &rpcx.Interchain{}
	if err := json.Unmarshal(receipt.Ret, ret); err != nil {
		return nil, fmt.Errorf("unmarshal interchain meta from bitxhub: %w", err)
	}

	return ret, nil
}

func (agent *BxhAgent) SyncBlockHeader(ctx context.Context, headerCh chan *pb.BlockHeader) error {
	ch, err := agent.client.Subscribe(ctx, pb.SubscriptionRequest_BLOCK_HEADER, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case h, ok := <-ch:
				if !ok {
					close(headerCh)
					return
				}
				headerCh <- h.(*pb.BlockHeader)
			}
		}
	}()

	return nil
}

func (agent *BxhAgent) GetBlockHeader(ctx context.Context, begin, end uint64, ch chan *pb.BlockHeader) error {
	if err := agent.client.GetBlockHeader(ctx, begin, end, ch); err != nil {
		return err
	}

	return nil
}

func (agent *BxhAgent) SyncInterchainTxWrapper(ctx context.Context, txCh chan *pb.InterchainTxWrapper) error {
	ch, err := agent.client.Subscribe(ctx, pb.SubscriptionRequest_INTERCHAIN_TX_WRAPPER, agent.from.Bytes())
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case h, ok := <-ch:
				if !ok {
					close(txCh)
					return
				}
				txCh <- h.(*pb.InterchainTxWrapper)
			}
		}
	}()

	return nil
}

// GetInterchainTxWrapper implements Agent
func (agent *BxhAgent) GetInterchainTxWrapper(ctx context.Context, begin, end uint64, ch chan *pb.InterchainTxWrapper) error {
	return agent.client.GetInterchainTxWrapper(ctx, agent.from.String(), begin, end, ch)
}

// SendTransaction implements Agent
func (agent *BxhAgent) SendTransaction(tx *pb.Transaction) (*pb.Receipt, error) {
	return agent.client.SendTransactionWithReceipt(tx)
}

// SendIBTP implements Agent
func (agent *BxhAgent) SendIBTP(ibtp *pb.IBTP) (*pb.Receipt, error) {
	b, err := ibtp.Marshal()
	if err != nil {
		return nil, err
	}
	return agent.client.InvokeContract(pb.TransactionData_BVM, rpcx.InterchainContractAddr,
		"HandleIBTP", rpcx.Bytes(b))
}

// GetReceipt implements Agent
func (agent *BxhAgent) GetReceipt(hash string) (*pb.Receipt, error) {
	return agent.client.GetReceipt(hash)
}

// GetIBTPByID implements Agent
func (agent *BxhAgent) GetIBTPByID(id string) (*pb.IBTP, error) {
	receipt, err := agent.client.InvokeContract(pb.TransactionData_BVM, rpcx.InterchainContractAddr,
		"GetIBTPByID", rpcx.String(id))
	if err != nil {
		return nil, err
	}

	hash := types.Bytes2Hash(receipt.Ret)

	response, err := agent.client.GetTransaction(hash.Hex())
	if err != nil {
		return nil, err
	}

	return response.Tx.GetIBTP()
}

// GetChainMeta implements Agent
func (agent *BxhAgent) GetChainMeta() (*pb.ChainMeta, error) {
	return agent.client.GetChainMeta()
}

func (agent *BxhAgent) GetAssetExchangeSigns(id string) ([]byte, error) {
	resp, err := agent.client.GetAssetExchangeSigns(id)
	if err != nil {
		return nil, err
	}

	if resp == nil || resp.Sign == nil {
		return nil, fmt.Errorf("get empty signatures for asset exchange id %s", id)
	}

	var signs []byte
	for _, sign := range resp.Sign {
		signs = append(signs, sign...)
	}

	return signs, nil
}
