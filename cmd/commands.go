package cmd

import (
	"fmt"
	"log"

	"scrape-us/scraper"

	"github.com/spf13/cobra"
)

var tag string
var ofJson, ofTxt, ofCsv bool

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(scrapeCmd)

	scrapeCmd.Flags().StringVarP(&tag, "tag", "t", "", "Tag to be scraped")
	scrapeCmd.MarkFlagRequired("tag")

	scrapeCmd.Flags().BoolVar(&ofJson, "json", false, "Output in JSON")
	scrapeCmd.Flags().BoolVar(&ofTxt, "txt", false, "Output in Text")
	scrapeCmd.Flags().BoolVar(&ofCsv, "csv", false, "Output in Csv")
	scrapeCmd.MarkFlagsOneRequired("json", "txt", "csv")
	scrapeCmd.MarkFlagsMutuallyExclusive("json", "txt", "csv")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Scrape-us",
	Long:  "All software has versions. Scrape-us has one too",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scrape-us cli web scraper v0.0.1 -- HEAD")
	},
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "I put the scrape in scrape-us",
	Long:  "This is the command used for scraping web pages",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: URL to scrape required")
		}
		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			log.Fatal(err)
		}
		url := args[0]
		fmt.Printf("%s : %s \n", tag, url)
		var ext string
		switch {
		case ofJson:
			ext = ".json"
		case ofTxt:
			ext = ".txt"
		case ofCsv:
			ext = ".csv"
		}
		scraper.Scrape(url, tag, ext)

	},
}
