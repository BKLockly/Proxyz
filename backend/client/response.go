/*
 * @Author: Lockly
 * @Date: 2025-02-17 16:56:24
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-17 16:58:48
 */

package client

import "fmt"

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func (a *App) errorResponse(message ...interface{}) Response {
	return Response{
		Code:    400,
		Message: fmt.Sprint(message...),
		Data:    nil,
	}
}

func (a *App) successResponse(msg string, data interface{}) Response {
	return Response{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}
