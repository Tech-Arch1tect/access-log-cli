package CLIConfiguration

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var Conf Toml
var configFileLocation string
var logLocation string
var output string

type Toml struct {
	SourceIPRegex   string
	MethodRegex     string
	URIRegex        string
	ProtocolRegex   string
	StatusCodeRegex string
	DateRegex       string
	LogFileLocation string
	TimeRegex       string
	Output          string
	DateFormat      string
}

func LoadConfig() {
	readFlags()
	var conf Toml
	if _, err := toml.DecodeFile(configFileLocation, &conf); err != nil {
		log.Fatal("Error reading configuration file: " + err.Error())
	}
	Conf = conf
	Conf.LogFileLocation = logLocation
	Conf.Output = output
}

func readFlags() {
	flag.StringVar(&configFileLocation, "c", "access-log-cli.toml", "Optional, specify config location.")
	flag.StringVar(&logLocation, "l", "access.log", "Required, location of log")
	flag.StringVar(&output, "o", "report.html", "Required, location of output report")
	flag.Parse()
}
