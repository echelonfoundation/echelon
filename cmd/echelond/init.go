package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/types"

	"github.com/cosmos/go-bip39"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
)

type printInfo struct {
	Moniker    string          `json:"moniker" yaml:"moniker"`
	ChainID    string          `json:"chain_id" yaml:"chain_id"`
	NodeID     string          `json:"node_id" yaml:"node_id"`
	GenTxsDir  string          `json:"gentxs_dir" yaml:"gentxs_dir"`
	AppMessage json.RawMessage `json:"app_message" yaml:"app_message"`
}

func newPrintInfo(moniker, chainID, nodeID, genTxsDir string, appMessage json.RawMessage) printInfo {
	return printInfo{
		Moniker:    moniker,
		ChainID:    chainID,
		NodeID:     nodeID,
		GenTxsDir:  genTxsDir,
		AppMessage: appMessage,
	}
}

func displayInfo(info printInfo) error {
	out, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(os.Stderr, "%s\n", string(sdk.MustSortJSON(out))); err != nil {
		return err
	}

	return nil
}

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func InitCmd(mbm module.BasicManager, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [moniker]",
		Short: "Initialize private validator, p2p, genesis, and application configuration files",
		Long:  `Initialize validators's and node's configuration files.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			// Echelon Default Configurations also found in cmd/echelond/root.go
			config.P2P.MaxNumInboundPeers = 240 // 8 to 1 ratio
			config.P2P.MaxNumOutboundPeers = 30

			// Set default seeds
			seeds := []string{
				"480db41faea6713405c93c505ff710a05d1fc801@94.250.203.190:26656",
				"41535ab44424500f44bb1b8d85fd941859991067@66.94.117.122:26656",
				"302ccf96853501c14060ffac2e1885bed6385f00@154.53.63.119:26656",
				"c06315472b5489b8d8b88622b86bc1e29b94002d@209.145.61.212:26656",
			}
			config.P2P.Seeds = strings.Join(seeds, ",")

			peers := []string{
				"00619d0710e367e00421223c03eda410bc1605f9@65.21.200.224:20656",
				"13b15c072e638982ba0c102e7e7c5add60256eae@66.94.125.156:20304",
				"cbf34b5ba7820b3bbf0a39e492fdc8665cfd8417@185.182.184.237:26656",
				"eac7a65f64247da11922300df7aef2eb446cde8b@24.117.162.116:26656",
				"a117b9bdd2b311097482be204a8b457317e21de5@209.145.62.69:26656",
				"4cb7a9a2ea5c238cb56242669553bbf5f7d2cad6@154.12.245.199:26696",
				"6b283c0dd751b015b4a9b6a60c514a7d1d169b8f@38.242.131.199:26656",
				"395dc53caf836f04474aa8069e8099b0629763a1@154.53.63.113:26656",
				"17fef6bc47f7fd69d3ca72d8da84fae785d59678@154.12.245.166:26656",
				"31d04668d9b22b281f5f148ef93d8da9288ecf5f@51.15.93.52:26656",
				"302ccf96853501c14060ffac2e1885bed6385f00@154.53.63.119:26656",
				"ba4b0793a0ff10675939cd6be2b28b4429a63efd@185.182.184.20:26656",
				"2bdc83cc8b257db83ff2960551d27095b0a05297@209.126.86.142:26656",
				"480db41faea6713405c93c505ff710a05d1fc801@94.250.203.190:26656",
				"c06315472b5489b8d8b88622b86bc1e29b94002d@209.145.61.212:26656",
				"d563d717ec4bcb7547ed1a67f07743b1673ace63@65.108.105.25:10756",
				"1fb4f150199532dc494893b8ec2b2dd3667100d3@185.169.252.163:26656",
			}
			config.P2P.PersistentPeers = strings.Join(peers, ",")

			config.Mempool.Size = 10000
			config.StateSync.TrustPeriod = 112 * time.Hour

			config.SetRoot(clientCtx.HomeDir)

			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("echelon_3000-%v", tmrand.Str(6))
			}

			// Get bip39 mnemonic
			var mnemonic string
			recover, _ := cmd.Flags().GetBool(genutilcli.FlagRecover)
			if recover {
				inBuf := bufio.NewReader(cmd.InOrStdin())
				value, err := input.GetString("Enter your bip39 mnemonic", inBuf)
				if err != nil {
					return err
				}

				mnemonic = value
				if !bip39.IsMnemonicValid(mnemonic) {
					return errors.New("invalid mnemonic")
				}
			}

			nodeID, _, err := genutil.InitializeNodeValidatorFilesFromMnemonic(config, mnemonic)
			if err != nil {
				return err
			}

			config.Moniker = args[0]

			genFile := config.GenesisFile()
			overwrite, _ := cmd.Flags().GetBool(genutilcli.FlagOverwrite)

			if !overwrite && tmos.FileExists(genFile) {
				return fmt.Errorf("genesis.json file already exists: %v", genFile)
			}

			appState, err := json.MarshalIndent(mbm.DefaultGenesis(cdc), "", " ")
			if err != nil {
				return errors.Wrap(err, "Failed to marshall default genesis state")
			}

			genDoc := &types.GenesisDoc{}
			if _, err := os.Stat(genFile); err != nil {
				if !os.IsNotExist(err) {
					return err
				}
			} else {
				genDoc, err = types.GenesisDocFromFile(genFile)
				if err != nil {
					return errors.Wrap(err, "Failed to read genesis doc from file")
				}
			}

			genDoc.ChainID = chainID
			genDoc.Validators = nil
			genDoc.AppState = appState

			if err := genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return errors.Wrap(err, "Failed to export gensis file")
			}

			toPrint := newPrintInfo(config.Moniker, chainID, nodeID, "", appState)

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
			return displayInfo(toPrint)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().BoolP(genutilcli.FlagOverwrite, "o", false, "overwrite the genesis.json file")
	cmd.Flags().Bool(genutilcli.FlagRecover, false, "provide seed phrase to recover existing key instead of creating")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")

	return cmd
}
