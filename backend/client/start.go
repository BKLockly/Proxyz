/*
 * @Author: Lockly
 * @Date: 2025-02-17 22:39:32
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 15:38:14
 */

package client

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) startListening() Response {
	if len(a.config.LiveProxyLists) == 0 {
		a.Error("缓存代理数为0，任务取消。")
		runtime.EventsEmit(a.ctx, "log_update", "[ERR] 缓存代理数为0。")
		runtime.EventsEmit(a.ctx, "log_update", "=========================== 任务取消 ============================")
		return a.errorResponse("缓存代理数为0，任务取消。")
	}

	listener, err := net.Listen("tcp", a.config.SocksAddress)
	if err != nil {
		a.Error("Error: %s\n", err.Error())
		runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[ERR] 监听失败 %s ", err.Error()))
		runtime.EventsEmit(a.ctx, "log_update", "=========================== 任务取消 ============================")
		return a.errorResponse(err.Error())
	}
	defer listener.Close()

	runtime.EventsEmit(a.ctx, "log_update", "=========================== 开始监听 ============================")
	runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[INF] 开始监听 socks5:\\%s -- 挂上代理以使用", a.config.SocksAddress))

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, a.config.CoroutineCount)

	for {
		conn, err := listener.Accept()
		if err != nil {
			runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[ERR] 接受连接失败 %s ", err.Error()))
			continue
		}

		semaphore <- struct{}{}
		wg.Add(1)
		go func(conn net.Conn) {
			defer wg.Done()
			defer func() { <-semaphore }()
			a.handleConnection(conn)
		}(conn)
	}
}
func (a *App) handleConnection(conn net.Conn) {
	defer conn.Close()

	if len(a.config.LiveProxyLists) == 0 {
		runtime.EventsEmit(a.ctx, "log_update", "[ERR] 没有可用代理")
		return
	}

	current := a.config.LiveProxyLists[rand.Intn(len(a.config.LiveProxyLists))]
	runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[INF] 当前使用代理 %s ", current))
	runtime.EventsEmit(a.ctx, "status_update", current)

	socks, err := net.DialTimeout("tcp", current, a.config.Timeout)
	if err != nil {
		runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[ERR] 连接代理失败 %s ", err.Error()))
		a.handleConnection(conn)
		return
	}
	defer socks.Close()

	var wg sync.WaitGroup
	ioCopy := func(dst io.Writer, src io.Reader) {
		defer wg.Done()
		_, err := io.Copy(dst, src)
		if err != nil {
			runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[ERR] 数据传输失败 %s ", err.Error()))
		}
	}

	wg.Add(2)
	go ioCopy(socks, conn)
	go ioCopy(conn, socks)
	wg.Wait()
}
