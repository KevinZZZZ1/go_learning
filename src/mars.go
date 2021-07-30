package main

import "fmt"

func main() {

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// 切片的长度和容量
	// 切片的长度就是当前切片包含的元素数量
	// 切片的容量就是当前切片对应数组的长度

	dump("dwarfs", dwarfs)
	// 下面这两个的cap输出是不一样的，第一个是1，第二个是3，由此可以推断出，cap的计算是从切片起始位置开始后，对应的数组长度
	dump("dwarfs[4:5]", dwarfs[4:5])
	dump("dwarfs[2:4]", dwarfs[2:4])

	dwarfs2 := append(dwarfs, "Orcus")

	dwarfs3 := append(dwarfs, "Salacia", "Quaoar", "Sedna")

	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)

}

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
