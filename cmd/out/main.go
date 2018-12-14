package out

import (
	"encoding/json"
	"fmt"
	"github.com/cappyzawa/dummy-resource"
	"os"
	"strconv"
	"time"
)

type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

type Params struct {
	PutParam1 string `json:"put_param1"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
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

	// 今回はresourceに対する操作がないためsrcは必要ない。
	//src := os.Args[1]

	t := time.Now()
	response := Response{
		resource.Version{Date: t.String()},
		[]resource.MetadataPair{
			{Name: "Year", Value: strconv.Itoa(t.Year())},
			{Name: "Month", Value: t.Month().String()},
			{Name: "Day", Value: strconv.Itoa(t.Day())},
		},
	}

	json.NewEncoder(os.Stdout).Encode(response)
}
