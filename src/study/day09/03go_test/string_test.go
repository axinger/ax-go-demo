package string_test

import (
	"axinger.xyz/ax_go_demo/study/day09/do_string"
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {

}

func TestAdd(t *testing.T) {

	s1 := do_string.Split()

	want := "abc"
	fmt.Println("TestAdd:", reflect.DeepEqual(s1, want))

}

// 性能基准测试
func BenchmarkName(b *testing.B) {
	//b.ResetTimer() //重置计算时间
	for i := 0; i < b.N; i++ {
		do_string.Split()
	}
}
