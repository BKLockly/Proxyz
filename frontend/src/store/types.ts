/*
 * @Author: Lockly
 * @Date: 2025-02-18 10:10:54
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 12:21:43
 */
import {defineStore} from "pinia";
import {GetProfile} from "../../wailsjs/go/client/App";

export interface ProfileItem {
    Code: number;
    Error: string;
    
    Status: number;

    FilePath: string;
    CoroutineCount: number;
    LiveProxies: number;
    AllProxies: number;
    Timeout: number;
    SocksAddress: string;
}

export const useConfigStore = defineStore('config', {
    state: (): ProfileItem => ({
        Code: 0,
        Error: '',
        
        Status: 0,

        FilePath: '路径为空',
        CoroutineCount: 0,
        LiveProxies: 0,
        AllProxies: 0,
        Timeout: 0,
        SocksAddress: 'NULL',
    }),
    actions: {
        getProfile() {
            GetProfile().then((profile) => {
                this.FilePath = profile.FilePath
                this.CoroutineCount = profile.CoroutineCount
                this.LiveProxies = profile.LiveProxies
                this.AllProxies = profile.AllProxies
                this.Timeout = profile.Timeout / 1000000000
                this.SocksAddress = profile.SocksAddress
            })
            return this
        },
        getSocksAddress() {
            return `socks5://${this.SocksAddress}`
        },
        getTimeout() {
            return this.Timeout
        },
        setTimeout(timeout: number) {
            this.Timeout = timeout
        },
        getLiveProxies() {
            return this.LiveProxies
        },
        getAllProxies() {
            return this.AllProxies
        },
        getCoroutineCount() {
            return this.CoroutineCount
        },
        setCoroutineCount(count: number) {
            this.CoroutineCount = count
        },
        getFilePath() {
            return this.FilePath
        },
        setFilePath(path: string) {
            this.FilePath = path
        },
        getStatus() {
            return this.Status
        },
        setStatus(status: number) {
            this.Status = status
        },
    }
})