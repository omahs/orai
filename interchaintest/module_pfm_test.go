package interchaintest

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	interchaintestrelayer "github.com/strangelove-ventures/interchaintest/v4/relayer"
	"github.com/strangelove-ventures/interchaintest/v4/testreporter"
	"github.com/strangelove-ventures/interchaintest/v4/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	helpers "github.com/oraichain/orai/tests/interchaintest/helpers"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PacketMetadata struct {
	Forward *ForwardMetadata `json:"forward"`
}

type ForwardMetadata struct {
	Receiver       string        `json:"receiver"`
	Port           string        `json:"port"`
	Channel        string        `json:"channel"`
	Timeout        time.Duration `json:"timeout"`
	Retries        *uint8        `json:"retries,omitempty"`
	Next           *string       `json:"next,omitempty"`
	RefundSequence *uint64       `json:"refund_sequence,omitempty"`
}

// TestPacketForwardMiddlewareRouter ensures the PFM module is set up properly and works as expected.
func TestPacketForwardMiddlewareRouter(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		ctx                                        = context.Background()
		client, network                            = interchaintest.DockerSetup(t)
		rep                                        = testreporter.NewNopReporter()
		eRep                                       = rep.RelayerExecReporter(t)
		chainID_A, chainID_B, chainID_C, chainID_D = "chain-a", "chain-b", "chain-c", "chain-d"
		chainA, chainB, chainC, chainD             *cosmos.CosmosChain
	)

	// base config which all networks will use as defaults.
	baseCfg := ibc.ChainConfig{
		Type:                   "cosmos",
		Name:                   "orai",
		ChainID:                "", // change this for each
		Images:                 []ibc.DockerImage{OraiImage},
		Bin:                    "oraid",
		Bech32Prefix:           "orai",
		Denom:                  Denom,
		CoinType:               "118",
		GasPrices:              "0orai",
		GasAdjustment:          2.0,
		TrustingPeriod:         "112h",
		NoHostMount:            false,
		ConfigFileOverrides:    nil,
		EncodingConfig:         junoEncoding(),
		UsingNewGenesisCommand: false,
		ModifyGenesis:          cosmos.ModifyGenesis(defaultGenesisKV),
	}

	// Set specific chain ids for each so they are their own unique networks
	baseCfg.ChainID = chainID_A
	configA := baseCfg

	baseCfg.ChainID = chainID_B
	configB := baseCfg

	baseCfg.ChainID = chainID_C
	configC := baseCfg

	baseCfg.ChainID = chainID_D
	configD := baseCfg

	// Create chain factory with multiple Juno individual networks.
	numVals := 1
	numFullNodes := 0

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "orai",
			ChainConfig:   configA,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "orai",
			ChainConfig:   configB,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "orai",
			ChainConfig:   configC,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "orai",
			ChainConfig:   configD,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
	})

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	chainA, chainB, chainC, chainD = chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain), chains[2].(*cosmos.CosmosChain), chains[3].(*cosmos.CosmosChain)

	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		interchaintestrelayer.CustomDockerImage(IBCRelayerImage, IBCRelayerVersion, "100:1000"),
		interchaintestrelayer.StartupFlags("--processor", "events", "--block-history", "100"),
	).Build(t, client, network)

	const pathAB = "ab"
	const pathBC = "bc"
	const pathCD = "cd"

	ic := interchaintest.NewInterchain().
		AddChain(chainA).
		AddChain(chainB).
		AddChain(chainC).
		AddChain(chainD).
		AddRelayer(r, "relayer").
		AddLink(interchaintest.InterchainLink{
			Chain1:  chainA,
			Chain2:  chainB,
			Relayer: r,
			Path:    pathAB,
		}).
		AddLink(interchaintest.InterchainLink{
			Chain1:  chainB,
			Chain2:  chainC,
			Relayer: r,
			Path:    pathBC,
		}).
		AddLink(interchaintest.InterchainLink{
			Chain1:  chainC,
			Chain2:  chainD,
			Relayer: r,
			Path:    pathCD,
		})

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),

		SkipPathCreation: false,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})

	const userFunds = int64(10_000_000_000)
	users := make([]ibc.Wallet, 0, 10)

	users = append(users,
		helpers.GetAndFundTestUserWithMnemonic(t, ctx, t.Name(), "", userFunds, chainA),
		helpers.GetAndFundTestUserWithMnemonic(t, ctx, t.Name(), "", userFunds, chainB),
		helpers.GetAndFundTestUserWithMnemonic(t, ctx, t.Name(), "", userFunds, chainC),
		helpers.GetAndFundTestUserWithMnemonic(t, ctx, t.Name(), "", userFunds, chainD),
	)
	err = testutil.WaitForBlocks(ctx, 5, chainA, chainB, chainC, chainD)
	require.NoError(t, err)
	abChan, err := ibc.GetTransferChannel(ctx, r, eRep, chainID_A, chainID_B)
	require.NoError(t, err)

	baChan := abChan.Counterparty

	cbChan, err := ibc.GetTransferChannel(ctx, r, eRep, chainID_C, chainID_B)
	require.NoError(t, err)

	bcChan := cbChan.Counterparty

	dcChan, err := ibc.GetTransferChannel(ctx, r, eRep, chainID_D, chainID_C)
	require.NoError(t, err)

	cdChan := dcChan.Counterparty

	// Start the relayer on all paths
	err = r.StartRelayer(ctx, eRep, pathAB, pathBC, pathCD)
	require.NoError(t, err)

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				t.Logf("an error occurred while stopping the relayer: %s", err)
			}
		},
	)

	// Get original account balances
	userA, userB, userC, userD := users[0], users[1], users[2], users[3]

	const transferAmount int64 = 100000

	// Compose the prefixed denoms and ibc denom for asserting balances
	firstHopDenom := transfertypes.GetPrefixedDenom(baChan.PortID, baChan.ChannelID, chainA.Config().Denom)
	secondHopDenom := transfertypes.GetPrefixedDenom(cbChan.PortID, cbChan.ChannelID, firstHopDenom)
	thirdHopDenom := transfertypes.GetPrefixedDenom(dcChan.PortID, dcChan.ChannelID, secondHopDenom)

	firstHopDenomTrace := transfertypes.ParseDenomTrace(firstHopDenom)
	secondHopDenomTrace := transfertypes.ParseDenomTrace(secondHopDenom)
	thirdHopDenomTrace := transfertypes.ParseDenomTrace(thirdHopDenom)

	firstHopIBCDenom := firstHopDenomTrace.IBCDenom()
	secondHopIBCDenom := secondHopDenomTrace.IBCDenom()
	thirdHopIBCDenom := thirdHopDenomTrace.IBCDenom()

	firstHopEscrowAccount := sdk.MustBech32ifyAddressBytes(chainA.Config().Bech32Prefix, transfertypes.GetEscrowAddress(abChan.PortID, abChan.ChannelID))
	secondHopEscrowAccount := sdk.MustBech32ifyAddressBytes(chainB.Config().Bech32Prefix, transfertypes.GetEscrowAddress(bcChan.PortID, bcChan.ChannelID))
	thirdHopEscrowAccount := sdk.MustBech32ifyAddressBytes(chainC.Config().Bech32Prefix, transfertypes.GetEscrowAddress(cdChan.PortID, abChan.ChannelID))

	t.Run("multi-hop a->b->c->d", func(t *testing.T) {
		// Send packet from Chain A->Chain B->Chain C->Chain D

		transfer := ibc.WalletAmount{
			Address: userB.FormattedAddress(),
			Denom:   chainA.Config().Denom,
			Amount:  transferAmount,
		}

		secondHopMetadata := &PacketMetadata{
			Forward: &ForwardMetadata{
				Receiver: userD.FormattedAddress(),
				Channel:  cdChan.ChannelID,
				Port:     cdChan.PortID,
			},
		}
		nextBz, err := json.Marshal(secondHopMetadata)
		require.NoError(t, err)
		next := string(nextBz)

		firstHopMetadata := &PacketMetadata{
			Forward: &ForwardMetadata{
				Receiver: userC.FormattedAddress(),
				Channel:  bcChan.ChannelID,
				Port:     bcChan.PortID,
				Next:     &next,
			},
		}

		memo, err := json.Marshal(firstHopMetadata)
		require.NoError(t, err)

		chainAHeight, err := chainA.Height(ctx)
		require.NoError(t, err)

		transferTx, err := chainA.SendIBCTransfer(ctx, abChan.ChannelID, userA.KeyName(), transfer, ibc.TransferOptions{Memo: string(memo)})
		require.NoError(t, err)
		_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+30, transferTx.Packet)
		require.NoError(t, err)
		err = testutil.WaitForBlocks(ctx, 1, chainA)
		require.NoError(t, err)

		chainABalance, err := chainA.GetBalance(ctx, userA.FormattedAddress(), chainA.Config().Denom)
		require.NoError(t, err)

		chainBBalance, err := chainB.GetBalance(ctx, userB.FormattedAddress(), firstHopIBCDenom)
		require.NoError(t, err)

		chainCBalance, err := chainC.GetBalance(ctx, userC.FormattedAddress(), secondHopIBCDenom)
		require.NoError(t, err)

		chainDBalance, err := chainD.GetBalance(ctx, userD.FormattedAddress(), thirdHopIBCDenom)
		require.NoError(t, err)

		require.Equal(t, userFunds-transferAmount, chainABalance)
		require.Equal(t, int64(0), chainBBalance)
		require.Equal(t, int64(0), chainCBalance)
		require.Equal(t, transferAmount, chainDBalance)

		firstHopEscrowBalance, err := chainA.GetBalance(ctx, firstHopEscrowAccount, chainA.Config().Denom)
		require.NoError(t, err)

		secondHopEscrowBalance, err := chainB.GetBalance(ctx, secondHopEscrowAccount, firstHopIBCDenom)
		require.NoError(t, err)

		thirdHopEscrowBalance, err := chainC.GetBalance(ctx, thirdHopEscrowAccount, secondHopIBCDenom)
		require.NoError(t, err)

		require.Equal(t, transferAmount, firstHopEscrowBalance)
		require.Equal(t, transferAmount, secondHopEscrowBalance)
		require.Equal(t, transferAmount, thirdHopEscrowBalance)
	})
}
