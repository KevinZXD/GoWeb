package applications

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)
var  num int = 1
// 路由handle提出来了而已
// 匿名函数方式 不重要
func Hello(c echo.Context) error {
	num = add(num)
	var  contactStr ="感恩相遇，未来可期，你是第"+strconv.Itoa(num)+"位获取我的联系方式：0001 0000 0110 0111 1000 1001 0010 0101 0000 0011"
	return c.String(http.StatusOK, contactStr)
}
