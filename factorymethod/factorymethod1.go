// 工厂方法模式
// 工厂方法模式使用子类的方式延迟生成对象到子类中实现。
// Go中不存在继承，所以使用匿名组合来实现
package factorymethod

//被封装的实际接口
type Operator1 interface {
	SetA1(int)
	SetB1(int)
	Result1() int
}

type OperatorFactory1 interface {
	Create1() Operator1
}

type OperatorBase1 struct {
	a, b int
}

func (o *OperatorBase1) SetA1(a1 int) {
	o.a = a1
}

func (o *OperatorBase1) SetB1(b1 int) {
	o.b = b1
}

type PlusOperatorFactory1 struct{}

func (PlusOperatorFactory1) Create1() Operator1 {
	return &PlusOperator1{
		&OperatorBase1{},
	}
}

type PlusOperator1 struct {
	*OperatorBase1
}

func (o PlusOperator1) Result1() int {
	return o.a + o.b
}

type MinusOperatorFactory1 struct{}

func (MinusOperatorFactory1) Create1() Operator1 {
	return &MinusOperator1{
		&OperatorBase1{},
	}
}

type MinusOperator1 struct {
	*OperatorBase1
}

func (o MinusOperator1) Result1() int {
	return o.a - o.b
}
