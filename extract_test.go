package extract
import(
    "testing"
    "os"
    "io"
    "bufio"
    "strings"
    "fmt"

    "github.com/huichen/sego"
)

func Test_Extract(t *testing.T)  {
    var segmenter sego.Segmenter
    segmenter.LoadDictionary("dict.txt")
    stopwords,_ := getStopwrods("stopwords.txt")
    extractor := New(segmenter,stopwords)

    fi,_ := os.Open("test.html")
    doc := extractor.Extract(fi)

    for _,kw := range doc.Keywords {
        fmt.Print(kw.Text() + " ")
        fmt.Print(kw.Weight())
        fmt.Println()
    }
}

func getStopwrods(f string) (map[string]int,error){
    stopwords := make(map[string]int)
    file, err := os.Open(f)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    rd := bufio.NewReader(file)
    for {
        word, err := rd.ReadString('\n')
        word = strings.TrimSpace(word)
        if io.EOF == err {
            break
        }
        if err != nil {
            return nil, err
        }
        stopwords[word] = 0
    }
    return stopwords, nil
}
