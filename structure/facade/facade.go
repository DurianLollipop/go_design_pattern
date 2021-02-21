// 结构型模式:外观模式
// API为facade模块的外观接口，大部分代码使用此接口简化对facade类的访问。

// facade模块同时暴露了a和b两个Module的NewXXX和interface，其他代码如果
// 需要使用细节功能时可以直接调用。
package facade

import "fmt"

func NewApi() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API is facade interface of facade package
type API interface {
	Test() string
}

// facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (api *apiImpl) Test() string {
	aRet := api.a.TestA()
	bRet := api.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImp{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleImp struct{}

func (*bModuleImp) TestB() string {
	return "B module running"
}
