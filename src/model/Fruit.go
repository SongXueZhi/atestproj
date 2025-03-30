package model

// Fruit 接口定义了水果的抽象行为
type Fruit interface {
	// Name 返回水果的名称
	Name() string
	// Color 返回水果的颜色
	Color() string
}

// Apple 结构体实现了 Fruit 接口
type Apple struct{}

func (a Apple) Name() string {
	return "Apple"
}

func (a Apple) Color() string {
	return "Red"
}

// Banana 结构体实现了 Fruit 接口
type Banana struct{}

func (b Banana) Name() string {
	return "Banana"
}

func (b Banana) Color() string {
	return "Yellow"
}
