package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// 前端传递的参数 从map找到对应的实现
	_ = dbMaps["mysql"].Insert(nil)
	_ = dbMaps["pgsql"].Update(nil)

	// 默认实现
	_ = dbMaps["mysql"].Update(nil)
}

var dbMaps = map[string]IDatabase{
	"mysql": Mysql{DefaultImpl: &DefaultImpl{}},
	"pgsql": Pgsql{DefaultImpl: &DefaultImpl{}},
}

var (
	_ IDatabase = (*Mysql)(nil)
	_ IDatabase = (*Pgsql)(nil)
)

type Mysql struct {
	*DefaultImpl
}

func (m Mysql) Insert(ele any) error {
	fmt.Println("mysql insert")
	return nil
}

type Pgsql struct {
	*DefaultImpl
}

func (p Pgsql) Update(ele any) error {
	fmt.Println("pgsql update")
	return nil
}

type IDatabase interface {
	Insert(ele any) error
	Update(ele any) error
}

// 默认实现 IDatabase
type DefaultImpl struct{}

func (impl DefaultImpl) Insert(ele any) error {
	panic("未实现")
}

func (impl DefaultImpl) Update(ele any) error {
	panic("未实现")
}
