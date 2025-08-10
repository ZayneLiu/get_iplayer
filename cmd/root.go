package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ZayneLiu/get_iplayer/pkg/api"
	"github.com/ZayneLiu/get_iplayer/pkg/config"
	"github.com/ZayneLiu/get_iplayer/pkg/models"
	"github.com/spf13/cobra"
)

const (
	Version = "3.36-go" // Go version based on Perl v3.36
)

var (
	// Global flags
	help         bool
	basicHelp    bool
	longHelp     bool
	version      bool
	tvQuality    string
	radioQuality string
	get          []string
	pid          []string
	channel      string
	typeFlag     string
	long         bool
	subtitles    bool
)

var rootCmd = &cobra.Command{
	Use:   "get_iplayer",
	Short: "BBC iPlayer/BBC Sounds Indexing Tool and PVR (Go implementation)",
	Long: `get_iplayer: BBC iPlayer/BBC Sounds Indexing Tool and PVR

Downloads TV and radio programmes from BBC iPlayer/BBC Sounds
Allows multiple programmes to be downloaded using a single command
Indexing of most available iPlayer/Sounds catch-up programmes
Caching of programme index with automatic updating
Regex search on programme name and description
Filter search results by channel
Direct download via programme ID or URL
PVR capability

This is a Go implementation of the original Perl get_iplayer tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Printf("get_iplayer %s (Go implementation)\n", Version)
			return
		}

		if basicHelp {
			cmd.Help()
			return
		}

		if longHelp {
			showLongHelp()
			return
		}

		// If no arguments provided, show help
		if len(args) == 0 && len(get) == 0 && len(pid) == 0 {
			cmd.Help()
			return
		}

		// Main search/download logic
		if len(get) > 0 {
			handleDownload()
		} else if len(pid) > 0 {
			handlePidDownload()
		} else if len(args) > 0 {
			handleSearch(args[0])
		}
	},
}

func init() {
	// Basic flags
	rootCmd.Flags().BoolVar(&version, "version", false, "Show version information")
	rootCmd.Flags().BoolVar(&basicHelp, "basic-help", false, "Show basic help")
	rootCmd.Flags().BoolVar(&longHelp, "long-help", false, "Show detailed help")

	// Search and display flags
	rootCmd.Flags().StringVar(&typeFlag, "type", "tv", "Programme type: tv, radio, or tv,radio")
	rootCmd.Flags().StringVar(&channel, "channel", "", "Filter by channel (regex)")
	rootCmd.Flags().BoolVar(&long, "long", false, "Show long programme descriptions")

	// Download flags
	rootCmd.Flags().StringSliceVar(&get, "get", []string{}, "Download programme by index number(s)")
	rootCmd.Flags().StringSliceVar(&pid, "pid", []string{}, "Download programme by PID(s)")
	rootCmd.Flags().StringVar(&tvQuality, "tv-quality", "hd,sd,web,mobile", "TV quality preference")
	rootCmd.Flags().StringVar(&radioQuality, "radio-quality", "high,std,med,low", "Radio quality preference")
	rootCmd.Flags().BoolVar(&subtitles, "subtitles", false, "Download subtitles if available")
}

func Execute() error {
	return rootCmd.Execute()
}

func showLongHelp() {
	fmt.Print(`get_iplayer: BBC iPlayer/BBC Sounds Indexing Tool and PVR (Go implementation)

USAGE:
  get_iplayer [OPTIONS] [SEARCH_TERM]

EXAMPLES:
  get_iplayer ".*"                     # List all TV programmes
  get_iplayer --type=radio ".*"        # List all radio programmes  
  get_iplayer "doctor who"             # Search for programmes with "doctor who" in name
  get_iplayer --channel="BBC One" ".*" # List all BBC One programmes
  get_iplayer --get 123                # Download programme index 123
  get_iplayer --pid=b01sc0wf           # Download programme by PID
  get_iplayer --long ".*"              # List programmes with long descriptions

SEARCH OPTIONS:
  --type string        Programme type: tv, radio, or tv,radio (default "tv")
  --channel string     Filter by channel (regex)
  --long              Show long programme descriptions

DOWNLOAD OPTIONS:
  --get strings       Download programme by index number(s)  
  --pid strings       Download programme by PID(s)
  --tv-quality string TV quality preference (default "hd,sd,web,mobile")
  --radio-quality string Radio quality preference (default "high,std,med,low")
  --subtitles         Download subtitles if available

OTHER OPTIONS:
  --version           Show version information
  --basic-help        Show basic help
  --long-help         Show this detailed help
  -h, --help          Show help

For more information, visit: https://github.com/get-iplayer/get_iplayer/wiki
`)
}

func handleSearch(searchTerm string) {
	fmt.Printf("Searching for: %s (type: %s)\n", searchTerm, typeFlag)
	if channel != "" {
		fmt.Printf("Channel filter: %s\n", channel)
	}

	// Create API client and search
	client := api.NewClient()
	result, err := client.Search(searchTerm, typeFlag, channel)
	if err != nil {
		fmt.Printf("Error searching: %v\n", err)
		return
	}

	if len(result.Programmes) == 0 {
		fmt.Println("No programmes found matching your search criteria.")
		return
	}

	fmt.Printf("\nFound %d programme(s):\n\n", result.Total)

	for _, prog := range result.Programmes {
		if long {
			printLongProgramme(prog)
		} else {
			printShortProgramme(prog)
		}
	}

	fmt.Println("\nNote: This is a Go implementation with mock data for demonstration.")
	fmt.Println("Use the original Perl version for accessing real BBC iPlayer/Sounds data.")
}

func printShortProgramme(prog models.Programme) {
	fmt.Printf("%d: %s - %s, %s, %s\n",
		prog.Index, prog.Name, prog.Episode, prog.Channel, prog.PID)
}

func printLongProgramme(prog models.Programme) {
	fmt.Printf("%d: %s - %s\n", prog.Index, prog.Name, prog.Episode)
	fmt.Printf("    Channel: %s\n", prog.Channel)
	fmt.Printf("    PID: %s\n", prog.PID)
	fmt.Printf("    Type: %s\n", prog.Type)
	fmt.Printf("    Duration: %d minutes\n", prog.Duration/60)
	fmt.Printf("    Available: %s\n", prog.Available.Format("2006-01-02 15:04"))
	if len(prog.Categories) > 0 {
		fmt.Printf("    Categories: %s\n", strings.Join(prog.Categories, ", "))
	}
	if prog.Description != "" {
		fmt.Printf("    Description: %s\n", prog.Description)
	}
	fmt.Printf("    URL: %s\n", prog.URL)
	fmt.Println()
}

func handleDownload() {
	fmt.Printf("Downloading programmes: %v\n", get)
	fmt.Printf("TV Quality: %s, Radio Quality: %s\n", tvQuality, radioQuality)
	if subtitles {
		fmt.Println("Subtitles will be downloaded if available")
	}

	// Load configuration
	cfg, err := config.LoadConfig("")
	if err != nil {
		fmt.Printf("Warning: Could not load config: %v\n", err)
		cfg = config.DefaultConfig()
	}

	// Create API client for looking up programmes
	// client := api.NewClient()

	for _, indexStr := range get {
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			fmt.Printf("Invalid index: %s\n", indexStr)
			continue
		}

		fmt.Printf("\nPreparing to download programme index %d...\n", index)
		fmt.Printf("Output directory: %s\n", cfg.OutputDir)
		fmt.Printf("Cache directory: %s\n", cfg.CacheDir)
		fmt.Println("Note: Actual download functionality not yet implemented.")
	}

	fmt.Println("\nNote: This is a Go implementation - download functionality is under development.")
	fmt.Println("Please use the original Perl version for full functionality.")
}

func handlePidDownload() {
	fmt.Printf("Downloading programmes by PID: %v\n", pid)
	fmt.Printf("TV Quality: %s, Radio Quality: %s\n", tvQuality, radioQuality)
	if subtitles {
		fmt.Println("Subtitles will be downloaded if available")
	}

	// Load configuration
	cfg, err := config.LoadConfig("")
	if err != nil {
		fmt.Printf("Warning: Could not load config: %v\n", err)
		cfg = config.DefaultConfig()
	}

	// Create API client for looking up programmes
	client := api.NewClient()

	for _, pidStr := range pid {
		fmt.Printf("\nLooking up programme with PID: %s\n", pidStr)

		prog, err := client.GetProgrammeByPID(pidStr)
		if err != nil {
			fmt.Printf("Error looking up PID %s: %v\n", pidStr, err)
			continue
		}

		fmt.Printf("Found: %s - %s\n", prog.Name, prog.Episode)
		fmt.Printf("Channel: %s\n", prog.Channel)
		fmt.Printf("Output directory: %s\n", cfg.OutputDir)
		fmt.Printf("Cache directory: %s\n", cfg.CacheDir)
		fmt.Println("Note: Actual download functionality not yet implemented.")
	}

	fmt.Println("\nNote: This is a Go implementation - download functionality is under development.")
	fmt.Println("Please use the original Perl version for full functionality.")
}
