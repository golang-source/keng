package main

import "sync"

type UserMoney struct{
	moneys map[string]int
	sync.Mutex
}

func (this *UserMoney) ChangeMoney(name string,money int) {
	this.Lock()
	defer this.Unlock()
	this.moneys[name]=money
}

func (this *UserMoney) GetMoney(name string) int{
	this.Lock()
	defer this.Unlock()
	if  v,exist:=this.moneys[name];exist{
		return v
	}
	return -1
}

func main() {

}