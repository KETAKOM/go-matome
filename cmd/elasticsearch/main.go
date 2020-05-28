package main

import (
	"log"
	"os"

	"context"

	elastic "github.com/olivere/elastic/v7"
)

func main() {
	url := os.Getenv("GOES")
	if url == "" {
		log.Println("invaild url")
		return
	}
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println("err", err)
		return
	}

	defer client.Stop()

	var res *elastic.GetResult
	res, err = client.Get().Index("user").Type("_doc").Id("2").Do(context.Background())
	if err != nil {
		log.Println("err", err)
		return
	}

	todo := Todo{}
	if err = todo.UnmarshalJSON(res.Source); err != nil {
		return
	}
	log.Println(todo.Title)
}

type Todo struct {
	Title string `json:"title,string"`
}

func (t *Todo) UnmarshalJSON(data []byte) error {
	t.Title = string(data)
	return nil
}
