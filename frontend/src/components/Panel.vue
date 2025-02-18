<!--
 * @Author: Lockly
 * @Date: 2025-02-12 23:03:06
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 15:57:28
 -->
<script setup lang="ts">

import {onMounted, onUnmounted, ref} from "vue";
import Run from "@/components/Run.vue";
import Fetch from "@/components/Fetch.vue";
import Config from "@/components/Config.vue";
import About from "@/components/About.vue";
import {EventsOff, EventsOn} from "../../wailsjs/runtime";
import {useConfigStore} from "@/store/types";

interface Metrics {
  cpu_usage: string
  mem_usage_self: string
  mem_usage_total: string
}

const current = ref("N/A")
const configStore = useConfigStore()
const dataset = ref<Metrics>({
  cpu_usage: '0%',
  mem_usage_self: '0 MB',
  mem_usage_total: '0%'
})
const handleMetricsUpdate = (metrics: Metrics) => {
  dataset.value = {
    cpu_usage: metrics.cpu_usage || '0%',
    mem_usage_self: metrics.mem_usage_self || '0 MB',
    mem_usage_total: metrics.mem_usage_total || '0%'
  }
}

onMounted(() => {
  EventsOn('metrics_update', handleMetricsUpdate)
  EventsOn('status_update', (r: string) => {
    current.value = r
  })
})

onUnmounted(() => {
  EventsOff('metrics_update')
  EventsOff('status_update')
})
</script>


<template>
  <a-layout style="height: 95.6vh;">
    <a-layout-header class="header">
      <a-alert center :type='configStore.getStatus() == 1 ? "info" : (configStore.getStatus() == 2 ? "success" : "error")'>
        <p v-if="configStore.getStatus() == 0">尚未指定数据源，请导入txt文件或在线获取。</p>
        <p v-else-if="configStore.getStatus() == 1">已指定，等待可用性测试完成...</p>
        <p v-else-if="configStore.getStatus() == 2">数据源已准备就绪, 当前IP: {{current}}。</p>
        <p v-else>执行异常，请检查配置。</p>
      </a-alert>
    </a-layout-header>
    
    <a-layout-content class="content">
      <a-tabs>
        <a-tab-pane key="1">
          <template #title>
            <icon-poweroff/> 运行
          </template>
            <Run/>
        </a-tab-pane>
        <a-tab-pane key="2">
          <template #title>
            <icon-clock-circle/> 获取
          </template>
            <Fetch/>
        </a-tab-pane>
        <a-tab-pane key="3">
          <template #title>
            <icon-settings/> 配置
          </template>
            <Config/>
        </a-tab-pane>
        <a-tab-pane key="4">
          <template #title>
            <icon-info-circle/> 关于
          </template>
            <About/>
        </a-tab-pane>
      </a-tabs>

    </a-layout-content>
    
    <a-layout-footer class="footer">
      <a-row :gutter="[24, 12]">
        <a-col :span="8">CPU使用率: {{dataset.cpu_usage}}</a-col>
        <a-col :span="8">自身使用: {{dataset.mem_usage_self}}</a-col>
        <a-col :span="8">内存使用: {{dataset.mem_usage_total}}</a-col>
      </a-row>
    </a-layout-footer>
  </a-layout>
</template>

<style scoped>

.header {
  margin: 8px;
}
.content {
  margin: 8px;
}

.footer {
  background-color: rgba(22, 22, 24, 0.22);
  padding: 10px;
  margin: 8px;
  color: #ebebeb;
}
</style>