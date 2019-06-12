package transformers

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLinkToHTMLContent(t *testing.T) {
	content := linkToHTMLContent("https://www.dinamalar.com/news_detail.asp?id=2296515", nil)
	ioutil.WriteFile("./temp.html", []byte(content.(string)), os.ModePerm)

}
