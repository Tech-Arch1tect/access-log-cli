package main

import (
	"github.com/Tech-Arch1tect/access-log-cli/CLIConfiguration"
	"github.com/Tech-Arch1tect/access-log-cli/Report"
	"github.com/Tech-Arch1tect/access-log-stats/configuration"
	"github.com/Tech-Arch1tect/access-log-stats/stats"
)

func main() {
	// get CLI specific config
	CLIConfiguration.LoadConfig()
	// Set regex values from our config
	configuration.Config.LogFile = CLIConfiguration.Conf.LogFileLocation
	configuration.Config.SourceIPRegex = CLIConfiguration.Conf.SourceIPRegex
	configuration.Config.MethodRegex = CLIConfiguration.Conf.MethodRegex
	configuration.Config.URIRegex = CLIConfiguration.Conf.URIRegex
	configuration.Config.ProtocolRegex = CLIConfiguration.Conf.ProtocolRegex
	configuration.Config.StatusCodeRegex = CLIConfiguration.Conf.StatusCodeRegex
	configuration.Config.DateRegex = CLIConfiguration.Conf.DateRegex
	configuration.Config.TimeRegex = CLIConfiguration.Conf.TimeRegex
	configuration.Config.DateFormat = CLIConfiguration.Conf.DateFormat

	//  get stats
	stats := stats.GetStats()
	Report.GenerateReport(stats)
}
