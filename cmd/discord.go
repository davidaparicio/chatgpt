/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// discordCmd represents the discord command
var discordCmd = &cobra.Command{
	Use:   "discord",
	Short: "Open OpenAI discord",
	Long:  `Open OpenAI discord`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Opening the OpenAI Discord...")
		cmd2 := exec.Command("open", "https://discord.com/invite/openai")
		err := cmd2.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(discordCmd)
}
