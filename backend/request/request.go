/*
 * @Author: Lockly
 * @Date: 2025-02-17 00:34:00
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 14:23:05
 */

package request

import (
	"encoding/json"
	"strconv"
)

type ProxyInfo struct {
	Key     string `json:"key"`
	Address string `json:"address"`
	Source  string `json:"source"`
}

type ProxyFetcher interface {
	Fetch() ([]string, error)
	Name() string
}

type ProxyManager struct {
	fetchers   []ProxyFetcher
	allProxies []string
	proxies    []ProxyInfo
}

func NewProxyManager(fetchers []ProxyFetcher) *ProxyManager {
	return &ProxyManager{
		fetchers: fetchers,
	}
}

func (pm *ProxyManager) RenderTable() ([]byte, error) {
	marshal, err := json.Marshal(pm.proxies)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (pm *ProxyManager) FetchAll() ([]string, error) {
	i := 1
	for _, fetcher := range pm.fetchers {
		proxies, err := fetcher.Fetch()
		if err != nil {
			continue
		}

		for _, proxy := range proxies {
			pi := &ProxyInfo{
				Key:     strconv.Itoa(i),
				Address: proxy,
				Source:  fetcher.Name(),
			}
			pm.proxies = append(pm.proxies, *pi)
			i++
		}

		pm.allProxies = append(pm.allProxies, proxies...)
	}

	return pm.allProxies, nil
}
