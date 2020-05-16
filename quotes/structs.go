package quotes

type quotesDataSource struct {
	quotes []quote
}

type quote struct {
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
}
