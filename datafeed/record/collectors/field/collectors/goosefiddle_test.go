package collectors

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExtractContentGoose(t *testing.T) {
	article, err := ExtractContentGoose("http://www.puthiyathalaimurai.com/news/world/65279-water-usage-in-various-industries.html")
	if err != nil {
		t.Error(err.Error())
	}

	ioutil.WriteFile("./temp.txt", []byte(article.CleanedText), os.ModePerm)

}
