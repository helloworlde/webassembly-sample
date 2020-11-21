# WebAssembly

## 使用 

- docker

```shell script
docker run --name webassembly-sample -p 8080:8080 docker.pkg.github.com/helloworlde/webassembly-sample/webassembly-sample:main
```

- 构建

```
git clone https://github.com/helloworlde/webassembly-sample.git

make all 
```

然后访问 [http://localhost:8080](http://localhost:8080) 即可看到结果

## 执行流程

项目结构有三部分
### 1. 前端页面

`static` 下面的 `index.html` 文件

- 初始化 WebAssembly

这里引入了 WebAssembly 的执行文件`wasm_exec.js`，同时指定了加载的 WebAssembly 文件 `wasm/index.wasm`

```html
<script src="wasm_exec.js"></script>
<script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("wasm/index.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
</script>
``` 

### 2. WebAssembly 文件

相应的 WebAssembly 逻辑在 `wasm` 目录下的 `index.go` 文件中，编译后的 `index.go` 文件会生成一个 `wasm` 格式的文件，该文件在 JS 中引入后可以被浏览器执行

在 main 函数中指定了两个全局的函数 `hello` 和 `changeContent`，分别用于响应 `index.html`中的两个按钮的事件

```go
func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("hello", js.FuncOf(hello))
	global.Set("changeContent", js.FuncOf(changeContent))

	<-done
}
```

需要注意的是 `done` 这个信号是必须的，如果没有，则会导致在执行的时候提示  `wasm_exec.js:539 Uncaught Error: Go program has already exited` 错误

以 `hello` 函数为例，`fmt.Println("Hello WebAssembly")` 会在 console 中输出日志，`js.Global().Call("alert", "Hello WebAssembly")` 会调用 JavaScript 的 `alert` 函数，触发弹窗

```go
func hello(this js.Value, args []js.Value) interface{} {
	fmt.Println("Hello WebAssembly")
	js.Global().Call("alert", "Hello WebAssembly")
	return nil
}
```

### 3. Server 服务

在 `main.go` 中实现，Server 服务本身和 WebAssembly 无关，只是为了提供一个可以访问 WebAssembly 项目的 Web 服务

```go
func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/wasm/", http.StripPrefix("/wasm/", http.FileServer(http.Dir("wasm"))))
	_ = http.ListenAndServe(":8080", nil)
}
```

在这个 Server 中，访问 `/` 路径，会指向`static`目录，最终指向 `index.html`文件；在加载 `index.html` 时会引用 `wasm/index.wasm` 文件，通过指定处理路径 `/wasm/`，会从 `wasm`这个文件夹下加载文件

----

- 如果使用 Goland 提示找不到 `syscall/js`，在设置 -> Go -> Build Tags & Vendoring 中选择 OS 为 js, Arch 为 wasm 即可