/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/identityOrg/cerberus-api/backend/setup"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server of Cerberus IAM",
	Long: `The API server for Cerberus Identity and Access Management Software:

Cerberus is a Identity and Access Management Software, built on Golang.
Which is fast and reliable. To start the server execute

  cerberus serve
  cerberus serve --debug  //to enable debug logging

Thanks for using it.`,
	Run: setup.StartServer,
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("addr", "a", "localhost:9090", "Listen port for server")
}
