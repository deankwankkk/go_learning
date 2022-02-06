## 基础语法

### 1. 表达方式

```golang
// var关键字 变量名 变量类型 = 变量赋值
var name string = "kwan"

// var关键字 括号内表示是一组变量
var (name string = "kwan", age int = 18)

// 不使用var关键字声明变量，并且不为变量指定数据类型，go会自动推断变量类型
name = "kwan"

// 只限于函数内声明
name, age, sex := "kwan", 18, "man"

// 常量: 常量名不需要大写，大写在go中有另外含义
const name = "kwan"

// 枚举: 如果枚举值是递增1的关系，可以使用iota代替
const (
  java = iota  // 0
  javascript   // 1
  python       // 2
  cpp          // 3
)
```

### 2. 内置变量类型

- `bool, string`
- `(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr` u 代表无符号， uintptr: ptr 是指针的意思，int 类型字节数是依据系统的位数
- `byte, rune` rune 是字符型，32 位，等同于 char，由于国际化语言出现
- `float32, float64, complex64, complex128` complex 是复数

### 3. 语句

#### 3.1 条件语句

> if 和 switch 的用法，与其他语言相似，但是 switch 后可以没有表达式以及不用 break

#### 3.2 迭代语句

> go 中没有 while 语句，使用 for 的缺参语句进行代替

### 4. 函数

#### 4.1 函数表达式

> 函数声明，以 func 开头，传入参数以及返回参数的声明方式与变量声明一致

```golang
// 常规函数
func speak() {}

// 多参数传入
func girls(name ...string) {}

// 返回参数多个
func beauty() (string, string, string) {}

// 传入参数为函数或匿名函数
func competition(fn func(int int) int, a, b int) int {
  return fn(a, b)
}
```

### 5. 指针

> 这里简单介绍指针概念，go 语言中的指针，只存在值传递的情况

```golang
func main() {
  a, b := 3, 4
  swap(&a, &b)
  fmt.Println(a, b) // 4, 3
}

func swap(a,b *int) {
  *a, *b = *b, *a
}
```
