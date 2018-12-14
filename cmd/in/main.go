package in

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cappyzawa/dummy-resource"
)

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
	Params  Params           `json:"params"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}

type Params struct {
	GetParam1 string `json:"get_param1"`
}

func main() {
	var request Request
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode: %s\n", err.Error())
		os.Exit(1)
		return
	}

	dest := os.Args[1]

	dayPath := fmt.Sprintf("%s/day", dest)
	yearPath := fmt.Sprintf("%s/year", dest)
	monthPath := fmt.Sprintf("%s/month", dest)

	dayFile, _ := os.Create(dayPath)
	yearFile, _ := os.Create(yearPath)
	monthFile, _ := os.Create(monthPath)

	defer dayFile.Close()
	defer yearFile.Close()
	defer monthFile.Close()

	t := time.Now()
	dayFile.WriteString(strconv.Itoa(t.Day()))
	yearFile.WriteString(strconv.Itoa(t.Year()))
	monthFile.WriteString(t.Month().String())

	response := Response{
		resource.Version{Date: request.Version.Date},
		[]resource.MetadataPair{
			{Name: "Year", Value: strconv.Itoa(t.Year())},
			{Name: "Month", Value: t.Month().String()},
			{Name: "Day", Value: strconv.Itoa(t.Day())},
		},
	}
	json.NewEncoder(os.Stdout).Encode(response)
}
