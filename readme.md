# go 
- 没有继承  匿名字段
- 指针
- 函数/方法绑定
- code fmt  :    ```go fmt main.go```
- switch  不需要break 每一个case都是独立的
- 没有while for代替

```
// 等同于while(true)
for{

}
```

- 数组切片

- 协程 channel
```
协程比线程轻量  不由os控制  
channel为各个协程服务 通信 同步（channel 理解为一个FIFO队列 协程一往队列中enqueue一个msg  协程二从队列中取一个msg）  在main中读取不符合个数的变量会阻塞main线程 
```


