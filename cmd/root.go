package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Scrape-us",
	Short: "Poorly Implemented web scraping cli tool",
	Long:  "Web scraper implemented in GO. Though poorly implemented it still scrapes.",
}

func Execute() error {
	return rootCmd.Execute()
}
