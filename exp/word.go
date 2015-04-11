package exp

type Word struct {
    text string
    frequency int
    pos string
    positions []*Position
    freqi float32
    lengthi float32
    posi float32
    addi float32
    selecti float32
}

type Words []*Word

type Position struct {
    Row int
    Index int
}

func (w *Word) Text() string {
    return w.text
}

func (w *Word) Weight() float32 {
    return 1.5*w.freqi + 1.1*w.posi + 1.0*w.addi + 1.0*w.selecti + 0.8*w.lengthi
}

func (words Words) Len() int {return len(words)}
func (words Words) Less(i,j int) bool {
    if words[i].Weight() > words[j].Weight() {
        return true
    } else if words[i].Weight() == words[j].Weight() && len(words[i].positions) > len(words[j].positions) {
        return true
    }
    return false
}
func (words Words) Swap(i,j int) {
    words[i],words[j] = words[j],words[i]
}
