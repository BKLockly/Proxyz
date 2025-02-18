/*
 * @Author: Lockly
 * @Date: 2025-02-17 17:40:35
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 18:10:44
 */

package client

import (
	"fmt"
	"strings"
	"sync"

	"github.com/imroc/req/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) CheckDatasets() Response {
	a.Debug("workers: ", a.config.CoroutineCount)
	availableProxies := make(chan string, a.config.CoroutineCount)
	var wg sync.WaitGroup

	checkedProxies := 0
	var mu sync.Mutex

	runtime.EventsEmit(a.ctx, "start_task", a.config.GetProfile())
	for _, ip := range a.config.LiveProxyLists {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			if a.checkProxy(ip) {
				availableProxies <- ip
			}

			mu.Lock()
			checkedProxies++
			mu.Unlock()

			progress := float64(checkedProxies) / float64(a.config.AllProxies)
			progressStr := fmt.Sprintf("%.2f", progress)
			runtime.EventsEmit(a.ctx, "task_progress", progressStr)
		}(ip)
	}

	go func() {
		wg.Wait()
		close(availableProxies)
	}()

	var availableProxiesList []string
	for proxy := range availableProxies {
		availableProxiesList = append(availableProxiesList, proxy)
	}

	a.config.SetLiveProxies(availableProxiesList)
	msg := fmt.Sprintf("共有 %d 条有效数据", a.config.LiveProxies)
	a.Debug("msg: ", msg)
	runtime.EventsEmit(a.ctx, "is_ready", a.config.LiveProxies)
	runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[INF] %s ", msg))

	return a.startListening()
}

func (a *App) checkProxy(proxyIP string) bool {
	client := req.C()
	client.SetProxyURL(fmt.Sprintf("socks5://%s", proxyIP))
	client.SetTimeout(a.config.Timeout)
	resp, err := client.R().Get("http://myip.ipip.net")
	if err != nil {
		a.Error("不可用： %s, 错误: %v\n", proxyIP, err)
		runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[ERR] %s <-- : --> %v", proxyIP, err))
		return false
	}

	if strings.Contains(resp.String(), "当前 IP") {
		a.Error("可用： %s\n", proxyIP, "resp: ", resp.String())
		runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[INF] 有效值 %s ", resp.String()))
		return true
	}

	a.Error("不可用： %s\n", proxyIP)
	runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[WAR] 不稳定 %s -- %v", proxyIP, err))
	return false
}
