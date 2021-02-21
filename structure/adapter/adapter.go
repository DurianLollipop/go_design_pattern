// 结构型模式:适配器模式

// 适配器模式用于转换一种接口适配另一种接口。
// 实际使用中Adapter一般为接口，并且使用工厂函数生成实例。
// 在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest的实例方法，
// 又因为Go语言中非入侵式接口特征，其实Adapter接口也适配Adaptee
package adapter

// Target是适配的目标接口
type Target interface {
	Request() string
}

// Adaptee是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// AdapteeImpl是被适配的目标类
type adapteeImpl struct{}

// SpecificRequest是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// NewAdapter是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
