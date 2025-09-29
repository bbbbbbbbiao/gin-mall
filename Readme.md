## git仓库代理指令

### （一）、设置代理

```shell
git config --global http.proxy http://127.0.0.1:7890
git config --global https.proxy https://127.0.0.1:7890
```

### （二）、取消代理

```shell
git config --global --unset http.proxy
git config --global --unset https.proxy
```



## 需要完善的点

### （一）、完善事务管理

![image-20250928143821727](C:\Users\baibiao\AppData\Roaming\Typora\typora-user-images\image-20250928143821727.png)

```go
// 建议创建一个事务管理的工具类： 

package utils

import (
    "gin-mall/global"
    "gorm.io/gorm"
)

// Transaction 执行数据库事务
func Transaction(fn func(tx *gorm.DB) error) error {
    tx := global.App.DB.Begin()
    if tx.Error != nil {
        return tx.Error
    }
    
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()
    
    if err := fn(tx); err != nil {
        tx.Rollback()
        return err
    }
    
    return tx.Commit().Error
}
```
### （二）、完善响应的错误码

### （三）、完善请求体及其校验

### （四）、完善商品浏览量点击接口

### (五)、完善请求体，响应体的命名（DTO, VO, PO, DO, Entity）



