package exp
import(
    "io"
    "strings"
    "sort"

    "github.com/huichen/sego"

    "github.com/zl-leaf/studom"
    "github.com/zl-leaf/studom/dom"
)

type Document struct {
    StuDomTree *dom.StuDomTree
    Keywords Words
    Setting *Setting
}

/**
 * 获取标题
 */
func (doc *Document) Title() string {
    tree := doc.StuDomTree
    return tree.Title()
}

/**
 * 获取body中的内容
 */
func (doc *Document) Body() string {
    tree := doc.StuDomTree
    return tree.Body()
}

func (doc *Document) Load(fi io.Reader)  {
    stuDomTree,err := studom.Parse(fi)
    if err == nil {
        stuDomTree.CutStuDomTree()
        doc.StuDomTree = stuDomTree
    }
    initKeywords(doc)
    return
}

func (doc *Document) LoadHTML(html string)  {
    stuDomTree,err := studom.ParseString(html)
    if err == nil {
        stuDomTree.CutStuDomTree()
        doc.StuDomTree = stuDomTree
    }
    initKeywords(doc)
    return
}

func initKeywords(doc *Document) {
    if doc.StuDomTree != nil {
        segmenter := doc.Setting.Segmenter
        stopwords := doc.Setting.Stopwords
        content := doc.StuDomTree.AllText()
        doc.Keywords = getKeywords(segmenter, stopwords, content)
    }
}

func getKeywords(segmenter sego.Segmenter, stopwords map[string]int, content string) Words {
    totalwordnum := segmenter.Dictionary().TotalFrequency()
    maxWordLength := 0
    totalLine := 0

    segments := make([]sego.Segment, 0)
    if content != "" {
        segments = segmenter.Segment([]byte(content))
    }

    wordMap := make(map[string]*Word)
    row := 1
    index := 1

    for _,segment := range segments {
        token := segment.Token()
        w := strings.TrimSpace(token.Text())
        if w == "\n" {
            row++
            index = 1
        }
        if w != ""{
            if len(w) > maxWordLength {
                maxWordLength = len([]rune(w))
            }
            if _,exist := stopwords[w];!exist {
                keyword,exist := wordMap[w]
                if !exist {
                    keyword = &Word{text:w,frequency:token.Frequency(),pos:token.Pos()}
                    wordMap[w] = keyword
                }

                // 记录位置
                pos := &Position{Row:row,Index:index}
                keyword.positions = append(keyword.positions, pos)
            }
            index++
        }
    }
    totalLine = row

    tmpKeywords := make(Words, 0)
    for _,kw := range wordMap {
        // 词频
        kw.freqi = float32(len(kw.positions))/float32(len(kw.positions) + 1)
        // 词长
        kw.lengthi = float32(len([]rune(kw.text)))/float32(maxWordLength)
        // 词性
        if kw.pos == "n" || kw.pos == "i" {
            kw.posi = 0.8
        } else if kw.pos == "a" || kw.pos == "d" {
            kw.posi = 0.6
        } else {
            kw.posi = 0
        }
        // 位置
        w1 := 0
        w2 := 0
        w3 := 0
        for _,p := range kw.positions {
            if p.Row == 1 {
                w1++
            } else if p.Row == 2 || p.Row == 3 {
                w2++
            } else if p.Row >= totalLine-1 {
                w3++
            }
        }
        kw.addi = 10.0*(5.0*float32(w1) + 3.0*float32(w2) + 2.0*float32(w3))/float32(len(wordMap))
        // 网络词频
        kw.selecti = float32(kw.frequency)/float32(totalwordnum) * float32(len(kw.positions))
        tmpKeywords = append(tmpKeywords, kw)
    }

    keywords := filterKeywords(tmpKeywords)

    if len(wordMap) < 400 {
        e := float32(len(wordMap))*0.5
        return keywords[:int(e)]
    }
    return keywords[:20]
}

func filterKeywords(words Words) Words {
    keywords := make(Words, 0)
    for _,w := range words {
        keywords = append(keywords, w)
    }
    sort.Sort(keywords)
    return keywords
}
