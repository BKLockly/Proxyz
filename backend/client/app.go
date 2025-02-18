/*
 * @Author: Lockly
 * @Date: 2025-02-12 22:54:13
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 14:57:36
 */

package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"proxyz/backend/config"
	"proxyz/backend/info"
	"proxyz/backend/request"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx      context.Context
	stopChan chan struct{}

	rpm     *request.ProxyManager
	metrics *info.ProxyMetrics
	config  *config.ProfileItem
}

func NewApp() *App {
	stopChan := make(chan struct{})
	return &App{
		stopChan: stopChan,
		config:   config.NewProfile(),
		rpm: request.NewProxyManager([]request.ProxyFetcher{
			&request.Free89,
			&request.FreeHappy,
			&request.FreeQiYun,
		}),
		metrics: info.NewProxyMetrics(stopChan),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	go a.metrics.StartMonitoring(1 * time.Second)
	go a.metricsPusher()
}

func (a *App) Shutdown(ctx context.Context) {
	a.metrics.StopMonitoring()
	close(a.stopChan)
}

func (a *App) GetProfile() config.ProfileItem {
	return a.config.GetProfile()
}

func (a *App) FetchProxies() Response {
	proxies, err := a.rpm.FetchAll()
	if err != nil {
		return a.errorResponse(err)
	}

	table, err := a.rpm.RenderTable()
	if err != nil {
		return a.errorResponse(err)
	}

	a.config.SetAllProxies(proxies)
	return Response{
		Code:    200,
		Message: "抓取成功",
		Data:    string(table),
	}
}

func (a *App) ChooseFile() config.ProfileItem {
	a.config.Code = 200
	a.config.FilePath, _ = runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:           "请选择配置文件",
		ShowHiddenFiles: true,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "配置文件",
				Pattern:     "*.txt",
			},
		},
	})

	if a.config.FilePath == "" {
		a.config.Code = 400
		a.config.Error = "未选择配置文件"
		return a.config.GetProfile()
	}

	var lists []string
	f, errOpen := os.Open(a.config.FilePath)
	if errOpen != nil {
		a.config.Code = 400
		a.config.Error = errOpen.Error()
		return a.config.GetProfile()
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			a.config.Code = 400
			a.config.Error = err.Error()
			return a.config.GetProfile()
		}
		if err == io.EOF {
			break
		}

		lists = append(lists, line)
	}

	a.config.SetAllProxies(lists)
	runtime.EventsEmit(a.ctx, "log_update", fmt.Sprintf("[INF] 配置文件 %s 读取成功", a.config.FilePath))
	rsp := a.CheckDatasets()
	if rsp.Code != 200 {
		a.Error("测试有误： %s\n", rsp.Message)
		a.config.Code = 400
		a.config.Error = rsp.Message
		return a.config.GetProfile()
	}

	return a.config.GetProfile()
}

func (a *App) UseFetchedDatasets() Response {
	runtime.EventsEmit(a.ctx, "log_update", "[INF] 使用抓取的代理")
	return a.CheckDatasets()
}

func (a *App) metricsPusher() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		runtime.EventsEmit(a.ctx, "metrics_update", a.metrics.GetMetrics())
	}
}
