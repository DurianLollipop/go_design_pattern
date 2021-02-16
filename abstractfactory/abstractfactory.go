// 抽象工厂模式

// 抽象工厂模式用于生成产品族的工厂, 所生成的是对象是有关联的.
// 如果抽象工厂退化成生成的对象无关联则成为工厂函数模式.
// 比如本例子中使用RDB和XML储存订单信息, 抽象工厂分别能生成相关的主订单信息和订单详情信息.
// 如果业务逻辑中需要替换使用的时候只需要改动工厂函数相关的类就能替换使用不同的存储方式了.
package abstractfactory

import "fmt"

// 订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// 订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// 关系型数据库OrderMainDAO实现
type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main save")
}

// 关系型数据库OrderDetailDAO实现
type RDBDetailDAO struct{}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

// RDB 抽象工厂实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XML储存OrderMainDAO实现
type XMLMainDAO struct{}

func (*XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

// XML储存OrderDetailDAO实现
type XMLDetailDAO struct{}

func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

// XML 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
