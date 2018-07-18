// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/infracloudio/vloadgenerator/src"
	"github.com/infracloudio/vloadgenerator/src/types"
	"github.com/spf13/cobra"
)

var appConfig *types.AppConfig

// attackCmd represents the attack command
var attackCmd = &cobra.Command{
	Use:   "attack",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: attack,
}

func init() {
	rootCmd.AddCommand(attackCmd)
	appConfig = new(types.AppConfig)

	attackCmd.PersistentFlags().IntVarP(&appConfig.Rate, "rate", "n", 5, "Request Rate")
	attackCmd.MarkFlagRequired("rate")

	attackCmd.PersistentFlags().IntVarP(&appConfig.Duration, "duration", "d", 50, "Duration")
	attackCmd.MarkFlagRequired("duration")

	attackCmd.PersistentFlags().StringVarP(&appConfig.Name, "app", "a", "hsl", "")
	attackCmd.MarkFlagRequired("app")

	attackCmd.PersistentFlags().StringVarP(&appConfig.URL, "url", "u", "http://localhost:8081", "")
	attackCmd.MarkFlagRequired("url")
}

func attack(cmd *cobra.Command, args []string) {
	//src.GenerateLoadData(rate, duration, api)
	src.Attack(appConfig)
}
