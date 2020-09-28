package note

/*
结构体声明方式:
1.按照顺序提供初始化值

P := person{"Tom", 25}

2.通过field:value的方式初始化，这样可以任意顺序

P := person{age:24, name:"Tom"}

3.当然也可以通过new函数分配一个指针，此处P的类型为*person

P := new(person)
*/

/*
method的概念，method是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在func后面增加了一个receiver(也就是method所依从的主体)。

A method is a function with an implicit first argument, called a receiver.

method的语法如下：
func (r ReceiverType) funcName(parameters) (results)
*/
type Color byte
type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}
