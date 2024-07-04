package test

// iota可理解为const语句块中的行索引
// 每次 const 出现时，都会让 iota 初始化为0.
const (
	a = iota
	b = iota
)

const (
	num  = 1
	user = "user"
	c    = iota
	name = "name"
	d    = iota
)
