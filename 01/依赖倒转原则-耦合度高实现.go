package main

// 司机张三，李四  汽车：奔驰，宝马

// 奔驰汽车
type Benz struct{}

func (b *Benz) run() {
	println("Benz run")
}

// 宝马汽车
type Bwm struct{}

func (b *Bwm) run() {
	println("Bwm run")
}

// 张三
type ZS struct{}

func (zs *ZS) Drive() {
	println("zs Drive")
	b := &Benz{}
	b.run()
}

// 李四
type LS struct{}

func (ls *LS) Drive() {
	println("ls Drive")
	b := &Bwm{} // 需要变动为Benz,但是所有调用该方法的都会变,如果未来再加入其他车型，不方便扩展
	b.run()
}

//  如果未来李四想开奔驰，张三想开宝马需要改方法逻辑,这样会影响到所有调用该方法的
// 如果不改原方法，新加个方法实现，如果后续再加入司机和车型，会变的很混乱
