package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s = "www"
	a := fmt.Sprintf("'%s' is good", s)
	fmt.Println(a)
	var s2 = "DATE_FORMAT(admiss_date,'%Y%m%d') between"
	var buffer bytes.Buffer
	buffer.WriteString(s2)
	var s3 = "'20191201'"
	buffer.WriteString(s3)
	var s4 = "and '20191231'"
	buffer.WriteString(s4)
	fmt.Println(buffer.String())
}
