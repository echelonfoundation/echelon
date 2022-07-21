package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/echelonfoundation/echelon/v3/app"
	"github.com/echelonfoundation/echelon/v3/testutil"
	claimtypes "github.com/echelonfoundation/echelon/v3/x/claims/types"
	"github.com/echelonfoundation/echelon/v3/x/recovery/types"
)

var _ = Describe("Recovery: Performing an IBC Transfer", Ordered, func() {
	coinEchelon := sdk.NewCoin("aechelon", sdk.NewInt(10000))
	coinOsmo := sdk.NewCoin("uosmo", sdk.NewInt(10))
	coinAtom := sdk.NewCoin("uatom", sdk.NewInt(10))

	var (
		sender, receiver       string
		senderAcc, receiverAcc sdk.AccAddress
		timeout                uint64
		claim                  claimtypes.ClaimsRecord
	)

	BeforeEach(func() {
		s.SetupTest()
	})

	Describe("from a non-authorized chain", func() {
		BeforeEach(func() {
			params := claimtypes.DefaultParams()
			params.AuthorizedChannels = []string{}
			s.EchelonChain.App.(*app.Echelon).ClaimsKeeper.SetParams(s.EchelonChain.GetContext(), params)

			sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
			receiver = s.EchelonChain.SenderAccount.GetAddress().String()
			senderAcc, _ = sdk.AccAddressFromBech32(sender)
			receiverAcc, _ = sdk.AccAddressFromBech32(receiver)
		})
		It("should transfer and not recover tokens", func() {
			s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

			nativeEchelon := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
			Expect(nativeEchelon).To(Equal(coinEchelon))
			ibcOsmo := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
			Expect(ibcOsmo).To(Equal(sdk.NewCoin(uosmoIbcdenom, coinOsmo.Amount)))
		})
	})

	Describe("from an authorized, non-EVM chain (e.g. Osmosis)", func() {

		Describe("to a different account on Echelon (sender != recipient)", func() {
			BeforeEach(func() {
				sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				receiver = s.EchelonChain.SenderAccount.GetAddress().String()
				senderAcc, _ = sdk.AccAddressFromBech32(sender)
				receiverAcc, _ = sdk.AccAddressFromBech32(receiver)
			})

			It("should transfer and not recover tokens", func() {
				s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

				nativeEchelon := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
				Expect(nativeEchelon).To(Equal(coinEchelon))
				ibcOsmo := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
				Expect(ibcOsmo).To(Equal(sdk.NewCoin(uosmoIbcdenom, coinOsmo.Amount)))
			})
		})

		Describe("to the sender's own eth_secp256k1 account on Echelon (sender == recipient)", func() {
			BeforeEach(func() {
				sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				receiver = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				senderAcc, _ = sdk.AccAddressFromBech32(sender)
				receiverAcc, _ = sdk.AccAddressFromBech32(receiver)
			})

			Context("with disabled recovery parameter", func() {
				BeforeEach(func() {
					params := types.DefaultParams()
					params.EnableRecovery = false
					s.EchelonChain.App.(*app.Echelon).RecoveryKeeper.SetParams(s.EchelonChain.GetContext(), params)
				})

				It("should not transfer or recover tokens", func() {
					s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)

					nativeEchelon := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
					Expect(nativeEchelon).To(Equal(coinEchelon))
					ibcOsmo := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
					Expect(ibcOsmo).To(Equal(sdk.NewCoin(uosmoIbcdenom, coinOsmo.Amount)))
				})
			})

			Context("with a sender's claims record", func() {
				Context("without completed actions", func() {
					BeforeEach(func() {
						amt := sdk.NewInt(int64(100))
						claim = claimtypes.NewClaimsRecord(amt)
						s.EchelonChain.App.(*app.Echelon).ClaimsKeeper.SetClaimsRecord(s.EchelonChain.GetContext(), senderAcc, claim)
					})

					It("should not transfer or recover tokens", func() {
						// Prevent further funds from getting stuck
						s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)

						nativeEchelon := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
						Expect(nativeEchelon).To(Equal(coinEchelon))
						ibcOsmo := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
					})
				})

				Context("with completed actions", func() {
					// Already has stuck funds
					BeforeEach(func() {
						amt := sdk.NewInt(int64(100))
						coins := sdk.NewCoins(sdk.NewCoin("aechelon", sdk.NewInt(int64(75))))
						claim = claimtypes.NewClaimsRecord(amt)
						claim.MarkClaimed(claimtypes.ActionIBCTransfer)
						s.EchelonChain.App.(*app.Echelon).ClaimsKeeper.SetClaimsRecord(s.EchelonChain.GetContext(), senderAcc, claim)

						// update the escrowed account balance to maintain the invariant
						err := testutil.FundModuleAccount(s.EchelonChain.App.(*app.Echelon).BankKeeper, s.EchelonChain.GetContext(), claimtypes.ModuleName, coins)
						s.Require().NoError(err)

						// aechelon & ibc tokens that originated from the sender's chain
						s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
						timeout = uint64(s.EchelonChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
					})

					It("should transfer tokens to the recipient and perform recovery", func() {
						// Escrow before relaying packets
						balanceEscrow := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), transfertypes.GetEscrowAddress("transfer", "channel-0"), "aechelon")
						Expect(balanceEscrow).To(Equal(coinEchelon))
						ibcOsmo := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())

						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisEchelon.RelayPacket(CreatePacket("10000", "aechelon", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisEchelon.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Check that the aechelon were recovered
						nativeEchelon := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
						Expect(nativeEchelon.IsZero()).To(BeTrue())
						ibcEchelon := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, aechelonIbcdenom)
						Expect(ibcEchelon).To(Equal(sdk.NewCoin(aechelonIbcdenom, coinEchelon.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo = s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))
					})

					It("should not claim/migrate/merge claims records", func() {
						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisEchelon.RelayPacket(CreatePacket("10000", "aechelon", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisEchelon.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						claimAfter, _ := s.EchelonChain.App.(*app.Echelon).ClaimsKeeper.GetClaimsRecord(s.EchelonChain.GetContext(), senderAcc)
						Expect(claim).To(Equal(claimAfter))
					})
				})
			})

			Context("without a sender's claims record", func() {
				When("recipient has no ibc vouchers that originated from other chains", func() {

					It("should transfer and recover tokens", func() {
						// aechelon & ibc tokens that originated from the sender's chain
						s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
						timeout = uint64(s.EchelonChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())

						// Escrow before relaying packets
						balanceEscrow := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), transfertypes.GetEscrowAddress("transfer", "channel-0"), "aechelon")
						Expect(balanceEscrow).To(Equal(coinEchelon))
						ibcOsmo := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())

						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisEchelon.RelayPacket(CreatePacket("10000", "aechelon", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisEchelon.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Check that the aechelon were recovered
						nativeEchelon := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
						Expect(nativeEchelon.IsZero()).To(BeTrue())
						ibcEchelon := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, aechelonIbcdenom)
						Expect(ibcEchelon).To(Equal(sdk.NewCoin(aechelonIbcdenom, coinEchelon.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo = s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))
					})
				})

				// Do not recover uatom sent from Cosmos when performing recovery through IBC transfer from Osmosis
				When("recipient has additional ibc vouchers that originated from other chains", func() {
					BeforeEach(func() {
						params := types.DefaultParams()
						params.EnableRecovery = false
						s.EchelonChain.App.(*app.Echelon).RecoveryKeeper.SetParams(s.EchelonChain.GetContext(), params)

						// Send uatom from Cosmos to Echelon
						s.SendAndReceiveMessage(s.pathCosmosEchelon, s.IBCCosmosChain, coinAtom.Denom, coinAtom.Amount.Int64(), s.IBCCosmosChain.SenderAccount.GetAddress().String(), receiver, 1)

						params.EnableRecovery = true
						s.EchelonChain.App.(*app.Echelon).RecoveryKeeper.SetParams(s.EchelonChain.GetContext(), params)

					})
					It("should not recover tokens that originated from other chains", func() {
						// Send uosmo from Osmosis to Echelon
						s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

						// Relay both packets that were sent in the ibc_callback
						timeout := uint64(s.EchelonChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err := s.pathOsmosisEchelon.RelayPacket(CreatePacket("10000", "aechelon", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisEchelon.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Aechelon was recovered from user address
						nativeEchelon := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
						Expect(nativeEchelon.IsZero()).To(BeTrue())
						ibcEchelon := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, aechelonIbcdenom)
						Expect(ibcEchelon).To(Equal(sdk.NewCoin(aechelonIbcdenom, coinEchelon.Amount)))

						// Check that the uosmo were retrieved
						ibcOsmo := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						// Check that the atoms were not retrieved
						ibcAtom := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, uatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(uatomIbcdenom, coinAtom.Amount)))

						// Repeat transaction from Osmosis to Echelon
						s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 2)

						timeout = uint64(s.EchelonChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err = s.pathOsmosisEchelon.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 3, timeout))
						s.Require().NoError(err)

						// No further tokens recovered
						nativeEchelon = s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
						Expect(nativeEchelon.IsZero()).To(BeTrue())
						ibcEchelon = s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, aechelonIbcdenom)
						Expect(ibcEchelon).To(Equal(sdk.NewCoin(aechelonIbcdenom, coinEchelon.Amount)))

						ibcOsmo = s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo = s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						ibcAtom = s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, uatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(uatomIbcdenom, coinAtom.Amount)))
					})
				})

				// Recover ibc/uatom that was sent from Osmosis back to Osmosis
				When("recipient has additional non-native ibc vouchers that originated from senders chains", func() {
					BeforeEach(func() {
						params := types.DefaultParams()
						params.EnableRecovery = false
						s.EchelonChain.App.(*app.Echelon).RecoveryKeeper.SetParams(s.EchelonChain.GetContext(), params)

						s.SendAndReceiveMessage(s.pathOsmosisCosmos, s.IBCCosmosChain, coinAtom.Denom, coinAtom.Amount.Int64(), s.IBCCosmosChain.SenderAccount.GetAddress().String(), receiver, 1)

						// Send IBC transaction of 10 ibc/uatom
						transferMsg := transfertypes.NewMsgTransfer(s.pathOsmosisEchelon.EndpointA.ChannelConfig.PortID, s.pathOsmosisEchelon.EndpointA.ChannelID, sdk.NewCoin(uatomIbcdenom, sdk.NewInt(10)), sender, receiver, timeoutHeight, 0)
						_, err := s.IBCOsmosisChain.SendMsgs(transferMsg)
						s.Require().NoError(err) // message committed
						transfer := transfertypes.NewFungibleTokenPacketData("transfer/channel-1/uatom", "10", sender, receiver)
						packet := channeltypes.NewPacket(transfer.GetBytes(), 1, s.pathOsmosisEchelon.EndpointA.ChannelConfig.PortID, s.pathOsmosisEchelon.EndpointA.ChannelID, s.pathOsmosisEchelon.EndpointB.ChannelConfig.PortID, s.pathOsmosisEchelon.EndpointB.ChannelID, timeoutHeight, 0)
						// Receive message on the echelon side, and send ack
						err = s.pathOsmosisEchelon.RelayPacket(packet)
						s.Require().NoError(err)

						// Check that the ibc/uatom are available
						osmoIBCAtom := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uatomOsmoIbcdenom)
						s.Require().Equal(osmoIBCAtom.Amount, coinAtom.Amount)

						params.EnableRecovery = true
						s.EchelonChain.App.(*app.Echelon).RecoveryKeeper.SetParams(s.EchelonChain.GetContext(), params)

					})
					It("should not recover tokens that originated from other chains", func() {
						s.SendAndReceiveMessage(s.pathOsmosisEchelon, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 2)

						// Relay packets that were sent in the ibc_callback
						timeout := uint64(s.EchelonChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err := s.pathOsmosisEchelon.RelayPacket(CreatePacket("10000", "aechelon", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisEchelon.RelayPacket(CreatePacket("10", "transfer/channel-0/transfer/channel-1/uatom", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisEchelon.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 3, timeout))
						s.Require().NoError(err)

						// Aechelon was recovered from user address
						nativeEchelon := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), senderAcc, "aechelon")
						Expect(nativeEchelon.IsZero()).To(BeTrue())
						ibcEchelon := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, aechelonIbcdenom)
						Expect(ibcEchelon).To(Equal(sdk.NewCoin(aechelonIbcdenom, coinEchelon.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						// Check that the ibc/uatom were retrieved
						osmoIBCAtom := s.EchelonChain.App.(*app.Echelon).BankKeeper.GetBalance(s.EchelonChain.GetContext(), receiverAcc, uatomOsmoIbcdenom)
						Expect(osmoIBCAtom.IsZero()).To(BeTrue())
						ibcAtom := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), senderAcc, uatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(uatomIbcdenom, sdk.NewInt(10))))
					})
				})
			})
		})
	})
})
