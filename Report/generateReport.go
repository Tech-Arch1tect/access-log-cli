package Report

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/Tech-Arch1tect/access-log-cli/CLIConfiguration"
	"github.com/Tech-Arch1tect/access-log-stats/stats"
)

func GenerateReport(stats stats.Stat) {
	t, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(CLIConfiguration.Conf.Output)
	if err != nil {
		log.Fatalln(err)
	}
	json, err := json.Marshal(stats)
	if err != nil {
		log.Fatalln(err)
	}
	data := getReportData(string(json))
	// Template file
	err = t.ExecuteTemplate(f, "report.tmpl", data)
	if err != nil {
		log.Fatalln(err)
	}

}

func getReportData(json string) reportData {
	data := reportData{}
	data.Json = json
	mapData := make(map[string]map[string]string)
	mapData = getMapData("SourceIPs", "10", "bar", mapData)
	mapData = getMapData("Methods", "0", "bar", mapData)
	mapData = getMapData("Protocols", "10", "bar", mapData)
	mapData = getMapData("URI", "10", "bar", mapData)
	mapData = getMapData("StatusCodes", "0", "bar", mapData)
	mapData = getMapData("Date", "0", "line", mapData)
	mapData = getMapData("ByHour", "0", "line", mapData)
	data.Graphs = mapData
	data.Logname = filepath.Base(CLIConfiguration.Conf.LogFileLocation)
	currentTime := time.Now()
	data.GeneratedOn = currentTime.Format("02/01/2006 15:04:05")
	return data
}

func getMapData(name, limit string, graphtype string, mapData map[string]map[string]string) map[string]map[string]string {
	mapData[name] = make(map[string]string)
	mapData[name]["Name"] = name
	mapData[name]["Limit"] = limit
	mapData[name]["GraphType"] = graphtype
	return mapData
}
