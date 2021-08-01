package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name)

	if err != nil {
		return err
	}

	// defer表示延迟执行，会在函数返回之前执行
	// defer可以作用于任何函数或方法
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully")
	return err

}

func main() {
	// go允许函数和方法返回多个值，而且按照惯例，返回的最后一个值表示错误
	// 在调用函数之后应该立即检查是否发生错误，而且在错误发生的时候函数返回的其他值就不可信了

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

}
