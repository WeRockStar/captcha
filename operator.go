package captcha

type operator struct {
	integerOperator int
}

func (op operator) get() string {
	p := map[int]string{1: "+", 2: "-", 3: "*"}
	return p[op.integerOperator]
}
