package list

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"

	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"

	"github.com/antihax/optional"
)

type options struct {
	outputFormat string
	page         int
	limit        int

	Config     config.IConfig
	Connection func() (connection.IConnection, error)
}

// NewListCommand creates a new command for listing kafkas.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		page:       0,
		limit:      100,
		Config:     f.Config,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all Kafka instances",
		Long:  "List all Kafka instances",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "json" && opts.outputFormat != "yaml" && opts.outputFormat != "yml" && opts.outputFormat != "table" {
				return fmt.Errorf("Invalid output format '%v'", opts.outputFormat)
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "table", "Format to display the Kafka instances. Choose from: \"json\", \"yaml\", \"yml\", \"table\"")
	cmd.Flags().IntVarP(&opts.page, "page", "", 0, "Page that should be returned from server")
	cmd.Flags().IntVarP(&opts.limit, "limit", "", 100, "Limit of items that should be returned from server")
	return cmd
}

func runList(opts *options) error {
	connection, err := opts.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	options := managedservices.ListKafkasOpts{
		Page: optional.NewString(strconv.Itoa(opts.page)),
		Size: optional.NewString(strconv.Itoa(opts.limit))}
	response, _, err := client.DefaultApi.ListKafkas(context.Background(), &options)

	if err != nil {
		return fmt.Errorf("Error retrieving Kafka instances: %w", err)
	}

	if response.Size == 0 {
		fmt.Fprintln(os.Stderr, "No Kafka instances found.")
		return nil
	}

	jsonResponse, _ := json.Marshal(response)

	var kafkaList kafka.ClusterList

	outputFormat := opts.outputFormat

	if err = json.Unmarshal(jsonResponse, &kafkaList); err != nil {
		fmt.Fprintf(os.Stderr, "Could not unmarshal Kakfa items into table, defaulting to JSON: %v\n", err)
		outputFormat = "json"
	}

	switch outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
		fmt.Print(string(data))
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		fmt.Print(string(data))
	default:
		printer := tableprinter.New(os.Stdout)
		printer.Print(kafkaList.Items)
		fmt.Fprint(os.Stderr, "\n")
	}

	return nil
}
