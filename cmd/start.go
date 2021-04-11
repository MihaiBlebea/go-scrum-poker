package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/MihaiBlebea/go-scrum-poker/conn"
	"github.com/MihaiBlebea/go-scrum-poker/poker"
	"github.com/MihaiBlebea/go-scrum-poker/server"
	"github.com/MihaiBlebea/go-scrum-poker/server/handler"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application server.",
	Long:  "Start the application server.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		conn, err := conn.ConnectSQL()
		if err != nil {
			return err
		}

		poker := poker.New(conn)

		handler := handler.New(poker, l)

		svr := server.New(handler, l)

		svr.Server()

		return nil
	},
}
