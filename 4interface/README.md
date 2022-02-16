### interface

#### usage

```go
// 声明接口
package mock

type Retriever struct{}

func (r Retriever) Get(url string) string {}

// 接口调用
type Retriever interface {
	Get(url string) string
}

func main() {
	r := mock.Retriever{}
	fmt.Println(r)
}
```

#### 接口变量
1. 接口变量包含接口实现者的类型以及实现者的指针
2. 接口变量自带指针
3. 接口变量同样采用值传递，几乎不需要使用接口的指针
4. 指针接收者实现只能以指针方式使用；值接收者两者都可以

#### 查看接口变量
1. interface{}
2. Type Assertion
3. Type Switch // 通过xxx.(type) 获取对应的类型

#### 接口组合
```go
type MulInterface interface {
	AInterface
	BInterface
}
```