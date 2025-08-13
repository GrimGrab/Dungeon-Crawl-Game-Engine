package cmd

import (
	"SoB/internal/server"
	"log"

	"github.com/spf13/cobra"
)

var port string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the WebSocket engine server",
	Long: `Start the WebSocket engine server that echoes "hello world" to all connected clients.
	
The server will start on the specified port (default: 8080) and provide:
- WebSocket endpoint at /ws that echoes "hello world"
- Web interface at / for testing the WebSocket connection

Example usage:
  gameserver serve
  gameserver serve --port 3000`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Starting WebSocket server on port %s...", port)

		if err := server.StartServer(port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Add port flag
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
}
