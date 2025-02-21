// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/cobra"

	"github.com/AssetMantle/modules/schema/helpers"
)

type cliCommand struct {
	use         string
	short       string
	long        string
	cliFlagList []helpers.CLIFlag
}

var _ helpers.CLICommand = (*cliCommand)(nil)

func (cliCommand cliCommand) registerFlags(command *cobra.Command) {
	for _, cliFlag := range cliCommand.cliFlagList {
		cliFlag.Register(command)
	}
}

func (cliCommand cliCommand) ReadInt64(cliFlag helpers.CLIFlag) int64 {
	switch cliFlag.GetValue().(type) {
	case int64:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(int64)
			}
		}
	default:
		panic(fmt.Errorf("flag %v not an int64 flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue()))
	}
	panic(fmt.Errorf("uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue()))
}

func (cliCommand cliCommand) ReadInt(cliFlag helpers.CLIFlag) int {
	switch cliFlag.GetValue().(type) {
	case int:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(int)
			}
		}
	default:
		panic(fmt.Errorf("flag %v not an int flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue()))
	}
	panic(fmt.Errorf("uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue()))
}

func (cliCommand cliCommand) ReadBool(cliFlag helpers.CLIFlag) bool {
	switch cliFlag.GetValue().(type) {
	case bool:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(bool)
			}
		}
	default:
		panic(fmt.Errorf("flag %v not an bool flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue()))
	}
	panic(fmt.Errorf("uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue()))
}

func (cliCommand cliCommand) ReadString(cliFlag helpers.CLIFlag) string {
	switch cliFlag.GetValue().(type) {
	case string:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(string)
			}
		}
	default:
		panic(fmt.Errorf("falg %v not an string flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue()))
	}
	panic(fmt.Errorf("uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue()))
}

func (cliCommand cliCommand) ReadBaseReq(context client.Context) rest.BaseReq {
	return rest.BaseReq{
		From:     context.GetFromAddress().String(),
		ChainID:  context.ChainID,
		Simulate: context.Simulate,
	}
}
func (cliCommand cliCommand) CreateCommand(runE func(command *cobra.Command, args []string) error) *cobra.Command {
	command := &cobra.Command{
		Use:   cliCommand.use,
		Short: cliCommand.short,
		Long:  cliCommand.long,
		RunE:  runE,
	}
	cliCommand.registerFlags(command)
	flags.AddTxFlagsToCmd(command)

	return command
}

func NewCLICommand(use string, short string, long string, cliFlagList []helpers.CLIFlag) helpers.CLICommand {
	return cliCommand{
		use:         use,
		short:       short,
		long:        long,
		cliFlagList: cliFlagList,
	}
}
