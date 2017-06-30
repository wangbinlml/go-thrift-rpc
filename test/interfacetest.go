package main

import (
	"fmt"
)

type DOInvoke interface {
	Invoke(service string) (string,error);
}
type DisInvoke struct {

}
func (dis *DisInvoke) Invoke(str string) (string,error){
	fmt.Println("DisInvoke")
	return str,nil
}
type DisInvoke2 struct {

}
func (dis *DisInvoke2) Invoke(str string) (string,error){
	fmt.Println("DisInvoke2")
	return str,nil
}
func main()  {
	var a DOInvoke
	var dis = new(DisInvoke2)
	a=dis

	str,err := a.Invoke("abvc")
	fmt.Println(str)
	fmt.Println(err)
}
