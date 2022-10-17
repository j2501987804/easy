package main

// 银行
type Bank struct {
	Name string
}

func (b *Bank) Save() {
	println("存款业务。。。")
}

func (b *Bank) Pay() {
	println("支付业务。。。")
}

func (b *Bank) Darw() {
	println("取钱业务。。。")
}

// 例如后边增加业务需要更改Bank的内容，可能会影响其他业务使用
