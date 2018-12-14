package resource

type Source struct {
	Config1 string `json:"config1"`
	Config2 string `json:"config2"`
}

type Version struct {
	Date string `json:"date"`
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
