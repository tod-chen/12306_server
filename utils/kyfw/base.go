package kyfw

import (
	"github.com/chuanshuo843/12306_server/utils"
)

var (
	request *utils.Request
)

type Base struct {

}

func init(){
	request = utils.InitRequest()
}

// func (base *Base) InitBase(){
// 	request.InitRequest()
// }