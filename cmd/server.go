package cmd

import (
	"basket-api/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the basket REST api",
	Run: func(cmd *cobra.Command, args []string) {
		server := api.New()

		err := server.Start(":" + viper.GetString("http_server.port"))
		if err != nil {
			log.Fatal("cannot start server:", err)
		}

	},
}
