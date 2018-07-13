package news

import (
	"fmt"
	"testing"
)

func TestGetArticles(t *testing.T) {
	api := New("fd2ae5f6a806473fb60bf56c2f2403c8")
	res, err := api.GetArticles("finance")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Result: ", res.Status)
	fmt.Println("TotalResults: ", res.TotalResults)
	fmt.Println("Articles: ", len(res.Articles))
	for i := 0; i < len(res.Articles); i++ {
		fmt.Println(res.Articles[i].Title)
	}
}

func TestGetHeadlines(t *testing.T) {
	api := New("fd2ae5f6a806473fb60bf56c2f2403c8")
	res, err := api.GetHeadlines("business")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Result: ", res.Status)
	fmt.Println("TotalResults: ", res.TotalResults)
	fmt.Println("Articles: ", len(res.Articles))
	for i := 0; i < len(res.Articles); i++ {
		fmt.Println(res.Articles[i].Title)
	}
}
