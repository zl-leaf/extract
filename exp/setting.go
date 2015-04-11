package exp
import(
    "github.com/huichen/sego"
)

type Setting struct {
    Segmenter sego.Segmenter
    Stopwords map[string]int
}
