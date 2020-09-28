package note

/*
空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。
可以存储任意类型的数值
*/

/*
通过switch来判断变量得类型
switch value := element.(type) {
	case int:
		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	case string:
		fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	case Person:
		fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
	default:
		fmt.Println("list[%d] is of a different type", index)
}
*/
