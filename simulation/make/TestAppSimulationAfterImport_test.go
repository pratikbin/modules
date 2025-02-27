// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"fmt"
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/schema/applications/base"
)

func TestAppSimulationAfterImport(t *testing.T) {
	config, db, _, logger, skip, closeFn, err := setupRun(t, "leveldb-app-sim", "Simulation")
	defer closeFn()

	if skip {
		t.Skip("skipping application simulation after import")
	}

	require.NoError(t, err, "simulation setup failed")

	prototype := base.NewSimulationApplication(applicationName, moduleBasicManager, wasm.EnableAllProposals, moduleAccountPermissions, tokenReceiveAllowedModules).(*base.SimulationApplication)
	simulationApplication := prototype.InitializeSimulationApplication(logger, db, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, prototype.GetDefaultNodeHome(), fauxMerkleModeOpt).(*base.SimulationApplication)
	require.Equal(t, "SimulationApplication", simulationApplication.Name())

	// Run randomized simulation
	stopEarly, simParams, simErr := simulation.SimulateFromSeed(
		t, os.Stdout, simulationApplication.GetBaseApp(), simapp.AppStateFn(simulationApplication.Codec(), simulationApplication.SimulationManager()),
		simapp.SimulationOperations(simulationApplication, simulationApplication.Codec(), config),
		simulationApplication.ModuleAccountAddresses(), config,
	)

	// export state and simParams before the simulation error is checked
	err = simapp.CheckExportSimulation(simulationApplication, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simapp.PrintStats(db)
	}

	if stopEarly {
		fmt.Printf("can't export or import a zero-validator genesis, exiting test...")
		return
	}

	fmt.Printf("exporting genesis...\n")

	appState, _, err := simulationApplication.ExportAppStateAndValidators(true, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, _, _, _, newCloseFn, err := setupRun(t, "leveldb-app-sim-2", "Simulation-2")
	defer newCloseFn()

	require.NoError(t, err, "simulation setup failed")

	newSimulationApplication := prototype.Initialize(logger, newDB, nil, true, simapp.FlagPeriodValue, map[int64]bool{}, prototype.GetDefaultNodeHome(), fauxMerkleModeOpt).(*base.SimulationApplication)
	require.Equal(t, "SimulationApplication", newSimulationApplication.Name())

	newSimulationApplication.InitChain(abci.RequestInitChain{
		AppStateBytes: appState,
	})

	_, _, err = simulation.SimulateFromSeed(
		t, os.Stdout, newSimulationApplication.GetBaseApp(), simapp.AppStateFn(simulationApplication.Codec(), simulationApplication.SimulationManager()),
		simapp.SimulationOperations(newSimulationApplication, newSimulationApplication.Codec(), config),
		newSimulationApplication.ModuleAccountAddresses(), config,
	)
	require.NoError(t, err)
}
