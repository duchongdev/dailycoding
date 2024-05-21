package test

import (
	"encoding/json"
	"testing"
	"time"
)

/*
golang通过内嵌（embedding）模拟继承
Go不支持传统意义上的继承（inheritance）,因此无法直接使用继承来实现代码重用或多态。
但是，Go提供了一种称为“内嵌”（embedding）的特性，它允许在结构体中嵌入另一个结构体，从而实现类似继承的效果。
内嵌结构体可以像其他字段一样直接访问，并且可以【调用嵌入结构体的方法】，这使得内嵌结构体的接口可以被外部结构体使用。
这种方式在Go中被广泛用于模拟继承的功能。
*/
func TestEmbedding(t *testing.T) {
	nowTime := time.Now()
	// t1 := T1{nowTime, 1}
	// t1str, _ := json.Marshal(t1)
	// //t1 = 2024-05-21 17:41:02.178285714 +0800 CST m=+0.001328448
	// t.Logf("t1 =  %s", t1str)

	// t2 := T2{T: nowTime, N: 1}
	// t2str, _ := json.Marshal(t2)
	// // {T:2024-05-21 17:41:02.178285714 +0800 CST m=+0.001328448 N:1}
	// t.Logf("t2 = %s", t2str)

	// t3 := T3{nowTime, 1}
	// t3str, _ := json.Marshal(t3)
	// //t1 = 2024-05-21 17:41:02.178285714 +0800 CST m=+0.001328448
	// t.Logf("t3 = %s", t3str)

	t4 := T4{nowTime, 1}
	t4str, _ := json.Marshal(t4)
	t.Logf("t4 = %s", t4str)

}

// 因为内嵌，T1 的方法集包括了 time.Time 的方法集，所以，T1 自动实现了 Marshaler 接口
type T1 struct {
	//time.Time，没有导出任何字段的结构体类型，因此它肯定实现了 Marshaler 接口。
	time.Time
	N int
}

type T2 struct {
	T time.Time
	N int
}

type T3 struct {
	time.Time
	int
}

type T4 struct {
	time.Time
	int
}

// 自定义MarshalJSON方法允许你完全控制序列化过程，你可以添加任何逻辑，比如忽略某些字段、修改字段名称、处理复杂的数据结构等。
func (t *T4) MarshalJSON() ([]byte, error) {
	temp := struct {
		Time string `json:"time"`
		Num  int    `json:"num"`
	}{Time: "111", Num: t.int}
	return json.Marshal(&temp)
}
