package test2
import "fmt"
import "test1"

func init () {
	fmt.Println("test2 init")
}

func Show2 () {
	fmt.Println("Show2")
	test1.Show1()
}	