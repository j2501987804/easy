package main

// 司机张三，李四  汽车：奔驰，宝马
// 先把司机和汽车进行抽象
// 然后实现司机对应的Drive方法，汽车主要实现对应的run方法
// ------> 抽象层 <--------
// 抽象的汽车
type Carer interface {
	Run()
}

// 抽象司机
type Driver interface {
	Drive(car Carer)
}

// ------> 实现层 <--------
// 汽车品类
type Benz struct{}

func (benz *Benz) Run() {
	println("benz run")
}

type Bwm struct{}

func (b *Bwm) Run() {
	println("bwm run")
}

// ...跟多车型随意添加，不需要去改动原有的代码逻辑

// 司机
type ZS struct{}

func (zs *ZS) Drive(car Carer) {
	println("zs drive car")
	car.Run()
}

type LS struct{}

func (ls *LS) Drive(car Carer) {
	println("ls drive car")
	car.Run()
}

// ...跟多司机随意添加，不需要去改动原有的代码逻辑

// ------> 业务逻辑层 <--------
func main() {
	// 这样业务逻辑编写也很灵活
	var car Carer = new(Benz)  // 类型关心对应的抽象类，不关心具体实现,调用者可以选择任意汽车品牌
	var drive Driver = new(ZS) // 调用者可以调用任何司机
	drive.Drive(car)

	// 例如张三突然想开宝马
	car = new(Bwm)
	drive.Drive(car)
}
