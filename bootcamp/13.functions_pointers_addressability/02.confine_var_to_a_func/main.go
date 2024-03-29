package main

import ("fmt" 
		"strconv"
	)

func main (){
	local := 10
	show(local)
	incrWrong(local)
	fmt.Printf("main local n = %d\n", local)

	local = incr(local)
	show(local)

	local = incrBy(local, 5)
	show(local)

	_, err := incrByStr(local, "TWO")
	if err != nil {
		fmt.Printf("err = %s\n", err)
	}
	local, _ = incrByStr(local,"2")
	show(local)

	show(incrBy(local,2))
	show(local)

	local = sanitize(incrByStr(local,"2"))
	show(local)

	local = limit(incrBy(local,5),500)
	show(local)
}


func show(n int){
	fmt.Printf("show  n = %d\n", n)
}

func incrWrong(n int){
	// n := main's local variable's value
	n++
}

func incr(n int) int {
	n++
	return n
}

func incrBy(n, factor int) int {
	return n*factor
}

func incrByStr(n int, factor string) (int,error){
	m, err := strconv.Atoi(factor)
	n = incrBy(n,m)
	return n, err
}

func sanitize(n int, err error) int{
	if err != nil{
		return 0
	}
	return n
}

func limit(n, lim int) (m int){
	// var m int
	m = n 
	if m >= lim {
		m = lim
	}
	return m
}