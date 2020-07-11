package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sunspa/sdk"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Long:  `Start main server`,
	Run: func(cmd *cobra.Command, args []string) {
		runApp()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runApp() {
	sv := sdk.NewService()
	defer sv.Close()

	if err := sv.Init(true, false); err != nil {
		log.Fatalln(err)
	}

	sv.Handle(Router)
	sv.Start()
}
