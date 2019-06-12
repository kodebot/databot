package collectors

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExtractContent(t *testing.T) {
	content, err := ExtractContent("http://www.puthiyathalaimurai.com/news/world/65279-water-usage-in-various-industries.html")
	if err != nil {
		t.Error(err.Error())
	}

	ioutil.WriteFile("./temp.txt", []byte(content), os.ModePerm)

	t.Log(content)
}
