// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package commands

//
import (
	"net/url"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/util/print"
	"github.com/nanobox-io/nanobox/util/server"
)

//
var execCmd = &cobra.Command{
	Hidden: true,
	
	Use:    "exec",
	Short:  "Runs a command from inside your app on the nanobox",
	Long:   ``,

	PreRun:  boot,
	Run:     execute,
	PostRun: halt,
}

// execute
func execute(ccmd *cobra.Command, args []string) {

	// PreRun: boot

	//
	if len(args) == 0 {
		args = append(args, print.Prompt("Please specify a command you wish to exec: "))
	}

	//
	v := url.Values{}
	v.Add("cmd", strings.Join(args[0:], " "))

	// if a container is found that matches args[0] then set that as a qparam, and
	// set the cmd equal to the remaining args
	if server.IsContainerExec(args) {
		v.Add("container", args[0])
		v.Set("cmd", strings.Join(args[1:], " "))
	}

	//
	server.Exec("command", v.Encode())

	// PostRun: halt
}