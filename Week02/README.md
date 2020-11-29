学习笔记

# 题目

> 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


## Get 到的点

```go
// Go 避免野协程 panic
func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover:", err)
			}
		}()
		x()
	}()
}
```