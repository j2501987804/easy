package main

// bad
// 穿衣服的方式
type Clothes struct{}

func (c *Clothes) OnWork() {
	println("工作的装扮")
}

func (c *Clothes) OnShop() {
	println(" 购物的装扮")
}

//
// func main() {
// 	c := &Clothes{}
//
// 	c.OnWork()
// 	c.OnShop()
// }

// good
type workClothes struct{}

func (wc *workClothes) style() {
	println("工作的装扮")
}

type shopClothes struct{}

func (sc *shopClothes) style() {
	println("购物的装扮")
}

func main() {
	wc := &workClothes{}
	wc.style()

	sc := &shopClothes{}
	sc.style()
}
