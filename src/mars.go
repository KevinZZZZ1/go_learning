package main

import "fmt"

// type report struct {
// 	sol       int
// 	high, low float64
// 	lat, long float64
// }

type sol int

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

// 这个struct就是组合上面三个类型和struct
type report struct {
	sol  sol
	temp temperature
	loc  location
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {

	// go可以使用struct实现组合，使用不同struct完成不同功能
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}

	rep := report{
		sol:  15,
		temp: t,
		loc:  bradbury,
	}

	fmt.Printf("%+v", rep)
	fmt.Println(rep.temp.average())

	// go也提供了嵌入(embedding)特性，可以实现方法的转发(forwarding)
	// 这里的嵌入是指：在声明struct的时候只给定类型，不给定字段名

	repEm := reportEmbedding{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}

	// 这里我们就能在repEm上直接调用本来在temperature上的方法，和fmt.Println(repEm.temperature.average())是一样的
	fmt.Println(repEm.average())

}

// 声明struct嵌入
type reportEmbedding struct {
	sol sol
	temperature
	location
}
