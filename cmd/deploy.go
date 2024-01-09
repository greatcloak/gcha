package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/greatcloak/gcha/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command.
// This deploys an app.
var deployCmd = &cobra.Command{
	Use: "deploy [app] [enviroment]",

	Aliases: []string{"d"},
	Short:   "Deploys an app to gcha.",
	Long:    `Deploy an app to gcha servers.`,
	Args:    cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		app := args[0]
		env := args[1]

		err := deployApp(cmd.Context(), app, env)
		if err != nil {
			log.WithError(err).Error("Error deploying app")
			return
		}
	},
}

var authToken string

func init() {
	deployCmd.PersistentFlags().StringVar(&authToken, "token", "", "auth token used to authenticate access to your gcha account")

	rootCmd.AddCommand(deployCmd)
}

type deployReq struct {
	Header api.CommandHeader

	App         string
	Environment string
}

func deployApp(ctx context.Context, app, environment string) error {
	appEnvLogger := log.WithFields(log.Fields{
		"app":         app,
		"environment": environment,
	})
	appEnvLogger.Info("Deploying application...")

	reqData := &deployReq{
		Header: api.CommandHeader{
			AuthToken: authToken,
		},

		App:         app,
		Environment: environment,
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(reqData)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, api.BaseAPIEndpoint+"deploy", buf)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.WithField("status", resp.Status).WithField("body", string(out)).Info("Response")

	if resp.StatusCode == http.StatusOK {
		appEnvLogger.Info("Successfully deployed application!")
	}

	return nil
}
