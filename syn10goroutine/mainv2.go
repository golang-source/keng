package main

import (
	"sync"
	"time"
	"fmt"
)
//主要是管理 goroutine 同步执行
//类似于 WaitGroup的功能
type Context struct{
	Count int
	sync.Mutex
	sig chan bool
	fnList []func(*Context)
}

func createContext()  *Context{
	var ctx =new(Context)
	ctx.Count=0
	ctx.sig = make(chan bool)
	return ctx
}

func (this *Context) newGoutine(fn func(*Context)) {
	  this.Lock()
	  defer this.Unlock()
	  this.Count++
	  this.fnList = append(this.fnList,fn)
}

func (this * Context) Commit() {

	for _,v:= range  this.fnList{
		//f:=v
		go v(this)
	}
}

func (this * Context) Exit() {
	  this.Lock()
	  this.Count--
	  if this.Count<=0{
	  	this.sig <-true
	  }
	  this.Unlock()
}
func (this *Context)destructor(){
	close(this.sig)
}
func (this *Context)Wait()  {
	  <-this.sig
	  this.destructor()
}

func main() {

	var ctx =createContext()

	ctx.newGoutine(func(ctx *Context) {
		for i:=0;i<10;i++{
			time.Sleep(time.Second)
			fmt.Println("area 1: index ",i)
		}

		ctx.Exit()
	})
	ctx.newGoutine(func(ctx *Context) {

		for i:=10;i<20;i++{
			time.Sleep(time.Second)
			fmt.Println("area 2: index ",i)
		}
		ctx.Exit()
	})

	ctx.newGoutine(func(ctx *Context) {
		for i:=20;i<30;i++{
			time.Sleep(time.Second)
			fmt.Println("area 3: index ",i)
		}

		ctx.Exit()
	})
    ctx.Commit()

	ctx.Wait()


}