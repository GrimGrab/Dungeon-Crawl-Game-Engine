package cmd

import (
	"SoB/internal/controller"
	"SoB/internal/engine"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var (
	port  string
	debug bool
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the game server",
	Long: `Start the WebSocket game server with the configured controller.
	
The server will start on the specified port (default: 8080) and provide:
- WebSocket endpoint at /ws for game connections
- Structured logging with debug mode support
- Pluggable controller architecture

Example usage:
  gameserver serve
  gameserver serve --port 3000
  gameserver serve --port 3000 --debug`,
	Run: func(cmd *cobra.Command, args []string) {
		// Setup logger
		logger := setupLogger(debug)

		logger.Info("starting game server",
			slog.String("port", port),
			slog.Bool("debug", debug))

		engine := engine.New()

		// Create game controller
		gameController := controller.NewGameController(
			logger.With(slog.String("component", "game")),
		)

		// Create server
		server := controller.NewGameServer(
			gameController,
			engine,
			logger.With(slog.String("component", "server")),
		)

		// Start server
		if err := server.Start(":" + port); err != nil {
			logger.Error("server failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	},
}

func setupLogger(debug bool) *slog.Logger {
	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if debug {
		opts.Level = slog.LevelDebug
		// Text handler for development (human-readable)
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		// JSON handler for production (machine-parseable)
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Add flags
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	serveCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
}
