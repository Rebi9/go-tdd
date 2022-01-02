package money

type PairAccesser interface {
	getFrom() string
	getTo() string
}

type Pair struct {
	from string
	to string
}

func (this *Pair) equals(accesser PairAccesser) bool {
	return this.getFrom() == accesser.getFrom() && this.getTo() == accesser.getTo()
}

func (this *Pair) getFrom() string {
	return this.from
}

func (this *Pair) getTo() string {
	return this.to
}

func (this *Pair) hashcode() int {
	return 0
}

func NewPair(from string, to string) Pair {
	return Pair{ from: from, to: to }
}