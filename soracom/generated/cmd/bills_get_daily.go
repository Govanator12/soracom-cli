// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// BillsGetDailyCmdYyyyMM holds value of 'yyyyMM' option
var BillsGetDailyCmdYyyyMM string

func init() {
	BillsGetDailyCmd.Flags().StringVar(&BillsGetDailyCmdYyyyMM, "yyyy-mm", "", TRAPI("year and month"))

	BillsGetDailyCmd.MarkFlagRequired("yyyy-mm")

	BillsCmd.AddCommand(BillsGetDailyCmd)
}

// BillsGetDailyCmd defines 'get-daily' subcommand
var BillsGetDailyCmd = &cobra.Command{
	Use:   "get-daily",
	Short: TRAPI("/bills/{yyyyMM}/daily:get:summary"),
	Long:  TRAPI(`/bills/{yyyyMM}/daily:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectBillsGetDailyCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)

	},
}

func collectBillsGetDailyCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsGetDailyCmd("/bills/{yyyyMM}/daily"),
		query:  buildQueryForBillsGetDailyCmd(),
	}, nil
}

func buildPathForBillsGetDailyCmd(path string) string {

	escapedYyyyMM := url.PathEscape(BillsGetDailyCmdYyyyMM)

	path = strReplace(path, "{"+"yyyyMM"+"}", escapedYyyyMM, -1)

	return path
}

func buildQueryForBillsGetDailyCmd() url.Values {
	result := url.Values{}

	return result
}
