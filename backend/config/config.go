/*
 * @Author: Lockly
 * @Date: 2025-02-16 18:44:58
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 18:10:49
 */

package config

import "time"

type ProfileItem struct {
	CoroutineCount int
	LiveProxies    int
	AllProxies     int
	LiveProxyLists []string
	Timeout        time.Duration
	SocksAddress   string
	FilePath       string

	Status int

	Code  int
	Error string
}

func NewProfile() *ProfileItem {
	return &ProfileItem{
		Code:           200,
		CoroutineCount: 200,
		LiveProxies:    0,
		AllProxies:     0,
		Timeout:        5 * time.Second,
		SocksAddress:   "127.0.0.1:1080",
	}
}

func (p *ProfileItem) GetProfile() ProfileItem {
	return *p
}

func (p *ProfileItem) GetCoroutineCount() int {
	return p.CoroutineCount
}

func (p *ProfileItem) GetLiveProxies() int {
	return p.LiveProxies
}

func (p *ProfileItem) SetAllProxies(datasets []string) {
	p.AllProxies = len(datasets)
	p.LiveProxyLists = datasets
}

func (p *ProfileItem) SetLiveProxies(datasets []string) {
	p.LiveProxyLists = datasets
	p.LiveProxies = len(datasets)
}

func (p *ProfileItem) GetTimeout() time.Duration {
	return p.Timeout
}

func (p *ProfileItem) GetSocksAddress() string {
	return p.SocksAddress
}
