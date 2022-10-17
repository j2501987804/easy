package main

// 抽象的银行业务员
type Banker interface {
	DoBusiness() // 抽象处理业务接口
}

// 存款业务员
type SaveBank struct{}

func (savebank *SaveBank) DoBusiness() {
	println("存款业务")
}

// 支付业务
type PayBank struct{}

func (paybank *PayBank) DoBusiness() {
	println("支付业务")
}

// 执行业务，主要传对应的业务员
func DoBusiness(b Banker) {
	b.DoBusiness()
}
// 其他业务增加，继续增加结构体实现方法即可，这样就可以互相隔离


func main() {
	DoBusiness(&SaveBank{})
	DoBusiness(&PayBank{})
}
