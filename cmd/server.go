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
		server, err := api.New()
		if err != nil {
			log.Fatal("cannot create server: ", err)
		}
		// Close closes all connections in the pool and rejects future Acquire calls.
		// Blocks until all connections are returned to pool and closed.
		defer server.DbPool.Close()

		err = server.Start(":" + viper.GetString("http_server.port"))
		if err != nil {
			log.Fatal("cannot start server:", err)
		}

	},
}
