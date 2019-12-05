package initialize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/tendermint/go-amino"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"

	"github.com/commitHub/commitBlockchain/applications/hub"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
)

func ExportGenesisFile(
	genFile, chainID string, validators []types.GenesisValidator, appState json.RawMessage,
) error {

	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		Validators: validators,
		AppState:   appState,
	}

	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genDoc.SaveAs(genFile)
}

func ExportGenesisFileWithTime(
	genFile, chainID string, validators []types.GenesisValidator,
	appState json.RawMessage, genTime time.Time,
) error {

	genDoc := types.GenesisDoc{
		GenesisTime: genTime,
		ChainID:     chainID,
		Validators:  validators,
		AppState:    appState,
	}

	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genDoc.SaveAs(genFile)
}

func InitializeNodeValidatorFiles(
	config *cfg.Config) (nodeID string, valPubKey crypto.PubKey, err error,
) {

	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return nodeID, valPubKey, err
	}

	nodeID = string(nodeKey.ID())
	server.UpgradeOldPrivValFile(config)

	pvKeyFile := config.PrivValidatorKeyFile()
	if err := common.EnsureDir(filepath.Dir(pvKeyFile), 0777); err != nil {
		return nodeID, valPubKey, nil
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err := common.EnsureDir(filepath.Dir(pvStateFile), 0777); err != nil {
		return nodeID, valPubKey, nil
	}

	valPubKey = privval.LoadOrGenFilePV(pvKeyFile, pvStateFile).GetPubKey()

	return nodeID, valPubKey, nil
}

func LoadGenesisDoc(cdc *amino.Codec, genFile string) (genDoc types.GenesisDoc, err error) {
	genContents, err := ioutil.ReadFile(genFile)
	if err != nil {
		return genDoc, err
	}

	if err := cdc.UnmarshalJSON(genContents, &genDoc); err != nil {
		return genDoc, err
	}

	return genDoc, err
}

func initializeNodeValidatorFiles(config *cfg.Config) (nodeID string, valPubKey crypto.PubKey, err error) {

	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return nodeID, valPubKey, err
	}

	nodeID = string(nodeKey.ID())
	server.UpgradeOldPrivValFile(config)

	privateValidatorKeyFile := config.PrivValidatorKeyFile()
	if err := common.EnsureDir(filepath.Dir(privateValidatorKeyFile), 0777); err != nil {
		return nodeID, valPubKey, nil
	}

	privateValidatorStateFile := config.PrivValidatorStateFile()
	if err := common.EnsureDir(filepath.Dir(privateValidatorStateFile), 0777); err != nil {
		return nodeID, valPubKey, nil
	}

	valPubKey = privval.LoadOrGenFilePV(privateValidatorKeyFile, privateValidatorStateFile).GetPubKey()

	return nodeID, valPubKey, nil
}

func exportGenesisFile(genesisFIle, chainID string, validators []types.GenesisValidator, appState json.RawMessage) error {

	genesisDoc := types.GenesisDoc{
		ChainID:    chainID,
		Validators: validators,
		AppState:   appState,
	}

	if err := genesisDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genesisDoc.SaveAs(genesisFIle)
}

func initializeEmptyGenesis(cdc *codec.Codec, genesisFile string, overwrite bool) (appState json.RawMessage, err error) {

	if !overwrite && common.FileExists(genesisFile) {
		return nil, fmt.Errorf("genesis.json file already exists: %v", genesisFile)
	}

	return codec.MarshalJSONIndent(cdc, hub.application.NewDefaultGenesisState())
}
