package main
 
func main() {
    var k int = 2
    if k == 1 {  //같은 라인
		println("One")
	}
	if k == 1 {
		println("One")
	} else if k == 2 {  //같은 라인
		println("Two")
	} else {   //같은 라인
		println("Other")
	}
    var i int = 1
    var max int = 3
	if val := i * 2; val < max {
		println(val)
	}
}