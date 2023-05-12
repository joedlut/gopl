package memo5

import (
	"io/ioutil"
	"net/http"
)

type result struct {
	value interface{}
	err   error
}

type request struct {
	key string
	//发送
	response chan<- result
}

type Memo struct {
	requests chan request
	//f  Func
	//mu sync.Mutex
	// *entry
	//cache map[string]*entry
}

type entry struct {
	res result
	//res准备好后关闭该通道
	ready chan struct{}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	//调用函数 f
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)

	//创建一个响应通道，放到request里面
	memo.requests <- request{key, response}

	//阻塞，直到response准备好数据
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

// 用作监控的goroutine
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		//第一次请求key
		if e == nil {
			//e := &entry{ready: make(chan struct{})}  fix
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			//调用f(key)
			//fmt.Println("begin to call e.call")
			//fmt.Println(e)
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	//执行fkey
	//fmt.Println("call f ")
	e.res.value, e.res.err = f(key)
	//广播数据已经准备好
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	//等待数据准备完毕
	<-e.ready
	//向客户端发送结果
	response <- e.res
}
