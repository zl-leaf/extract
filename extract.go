package extract

import(
    "io"

    "github.com/huichen/sego"

    "github.com/zl-leaf/extract/exp"
)

type Extractor struct {
    setting *exp.Setting
}

func New(segmenter sego.Segmenter, stopwords map[string]int) *Extractor {
    setting := &exp.Setting{Segmenter:segmenter, Stopwords:stopwords}
    extractor := &Extractor{setting:setting}
    return extractor
}

func (extractor *Extractor) Extract(fi io.Reader) (doc *exp.Document) {
    doc = &exp.Document{Setting:extractor.setting}
    doc.Load(fi)

    return
}

func (extractor *Extractor) ExtractString(html string) (doc *exp.Document) {
    doc = &exp.Document{Setting:extractor.setting}
    doc.LoadHTML(html)

    return
}
