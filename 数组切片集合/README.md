### 容器

#### 数组

> go 中的数组是值类型

```go
// 声明方式
var arr [3]int // [0 0 0]
arr1 := [3]int{1, 2, 3} // 必须赋值
arr2 := [...]int{1, 2, 3} // 没有指定数组长度

// 多维数组
var arr3 [4][5]int

// 遍历
for k, v := range arr {....}
```

#### 切片

> 切片可以理解为数组的 view

```go
// 切片创建方式
arr := [3]int{1, 2, 3, 4, 5}
slice1 := arr[:] // [1 2 3 4 5]

// 使用make创建
slice2 := make([]int, len, cap)

// 切片可以向后扩展，不可以向前扩展。向后扩展不可以超过底层数组的cap

// slice的基本操作
// append
slice3 := append(slice1, 6) // [1 2 3 4 5 6]

// copy(dest, src)
copy(slice2, slice3)

// go的slice中没有"+"，所以拼接或删除需要用append实现
```

#### map

> map 是无序的，以 key-value 的形式存在

```go
// 声明方式
// map[key-type]value-type{}
// make(map[key-type]value-type)
```
