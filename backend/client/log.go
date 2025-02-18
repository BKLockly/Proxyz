/*
 * @Author: Lockly
 * @Date: 2025-02-17 17:15:36
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-17 17:16:30
 */

package client

import (
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) Debug(msg ...interface{}) {
	runtime.LogDebug(a.ctx, fmt.Sprint(msg...))
}

func (a *App) Info(msg ...interface{}) {
	runtime.LogInfo(a.ctx, fmt.Sprint(msg...))
}

func (a *App) Warn(msg ...interface{}) {
	runtime.LogWarning(a.ctx, fmt.Sprint(msg...))
}

func (a *App) Error(msg ...interface{}) {
	runtime.LogError(a.ctx, fmt.Sprint(msg...))
}

func (a *App) Fatal(msg ...interface{}) {
	runtime.LogFatal(a.ctx, fmt.Sprint(msg...))
}
