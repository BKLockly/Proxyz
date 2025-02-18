<!--
 * @Author: Lockly
 * @Date: 2025-02-16 16:57:57
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 15:59:47
 -->
<template>
  <a-row :gutter="12">
    <a-col :span="18">
      <a-input
          :readonly="true"
          :model-value="configState.getFilePath()"
      >
        <template #suffix>
          <a-tooltip content="可以从此处获取, 点击即可打开">
            <icon-question-circle @click="BrowserOpenURL('http://proxycompass.com/cn/free-proxies/asia/china/')"/>
          </a-tooltip>
        </template>
      </a-input>
    </a-col>
    <a-col :span="6">
      <a-button
          :disabled="disabled"
          v-show="configState.getStatus() != 2"
          @click="openFile"
          size="medium"
          type='outline' long
      >导入
      </a-button>
      <a-button
          v-show="configState.getStatus() == 2"
          type="outline"
          size="medium"
          @click="configState.setStatus(0)"
      >停止
      </a-button>
    </a-col>
  </a-row>

  <br>

  <a-row :gutter="12">
    <a-col :span="5">
      <a-card hoverable :bordered="false" class="card progress">
        <a-progress
            :percent="percent"
            :type="'circle'"
            :status='configState.getStatus() === 0 || configState.getStatus() === 3 ? "danger" : (configState.getStatus() === 2 ? "success" : "normal")'
            size="medium"
            style="margin-bottom: 10px;"
        />
        <a-tag :color='configState.getStatus() === 0 || configState.getStatus() === 3 ? "red" : (configState.getStatus() === 2 ? "green" : "blue")'>
          <span v-if="configState.getStatus() === 0">尚未测试</span>
          <span v-else-if="configState.getStatus() === 1">正在测试</span>
          <span v-else-if="configState.getStatus() === 2">测试完成</span>
          <span v-else>任务取消</span>
        </a-tag>
      </a-card>
    </a-col>

    <a-col :span="19">
      <a-card hoverable :bordered="false" class="card">
        <a-descriptions
            :column="3"
            v-model:data="details"
            title='当前配置'
            :align="{ label: 'right' }"
        />
      </a-card>
    </a-col>
  </a-row>

  <br>

  <div class="log-viewer" ref="logContainer">
    <div v-for="(log, index) in logs" :key="index">
      <span :style="logStyle(log)">{{ log }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, ref } from "vue";
import { ChooseFile } from "../../wailsjs/go/client/App";
import { Notification } from '@arco-design/web-vue';
import {BrowserOpenURL, EventsOff, EventsOn} from "../../wailsjs/runtime";
import {ProfileItem, useConfigStore} from "@/store/types";

const logs = ref<string[]>([]);
const logContainer = ref<HTMLElement | null>(null);

const percent = ref(0.00);
const disabled = ref(false);

const started = ref(false);

const configState = useConfigStore();
const data = ref<ProfileItem>();
const details = ref([
  {
    label: '监听绑定',
    value: configState.getSocksAddress(),
  },
  {
    label: '超时时间',
    value: `${configState.getTimeout()} s`,
  },
  {
    label: '协程数',
    value: `${configState.getCoroutineCount()}`,
  },
  {
    label: '代理总数',
    value: `${configState.getAllProxies()} 条`,
  },
  {
    label: '有效代理',
    value: `${configState.getLiveProxies()} 条`,
  },
]);

function openFile() {
  disabled.value = true;
  ChooseFile().then((res) => {
    data.value = res as unknown as ProfileItem;
    if (data.value.Code !== 200) {
      configState.setStatus(3)
      Notification.error({
        title: '任务失败',
        content: data.value.Error,
      });
      return;
    }
  }).catch((err) => {
    configState.setStatus(3)
    Notification.error({
      title: '导入失败',
      content: err,
    });
  }).finally(() => {
    disabled.value = false;
  });
}

const handleLogEmits = (log: string) => {
  logs.value.push(log);
  if (logContainer.value) {
    logContainer.value.scrollTop = logContainer.value.scrollHeight;
  }
};

const logStyle = (log: string) => {
  if (log.includes('[ERR]')) {
    return { color: '#b35351' };
  } else if (log.includes('[INF]')) {
    return { color: '#29b445' };
  } else if (log.includes('[WAR]')) {
    return { color: '#b16f34' };
  }
  return {};
};

onMounted(() => {
  configState.getProfile()
  EventsOn('log_update', handleLogEmits);
  EventsOn('task_progress', (p: number) => {
    percent.value = p;
  });
  EventsOn('start_task', (profile: ProfileItem) => {
    started.value = true;
    configState.setFilePath(profile.FilePath);
    configState.setStatus(1)
    details.value = [
      {
        label: '监听绑定',
        value: profile.SocksAddress,
      },
      {
        label: '协程数',
        value: `${profile.CoroutineCount}`,
      },
      {
        label: '超时时间',
        value: `${profile.Timeout / 1000000000}s`,
      },
      {
        label: '代理总数',
        value: `${profile.AllProxies}条`,
      },
    ]
  });
  EventsOn('is_ready', (callback: string) => {
    started.value = false;
    configState.setStatus(2)
    details.value.push({
      label: '有效代理',
      value: `${callback} 条`,
    })
    Notification.success({
      title: "任务完成",
      content: `共有 ${callback} 条有效数据`,
      duration: 3000,
      closable: true,
    });
  });
});

onUnmounted(() => {
  EventsOff('log_update');
  EventsOff('task_progress');
  EventsOff('start_task');
  EventsOff('is_ready');
});
</script>

<style scoped>
.log-viewer {
  height: 40vh;
  border-radius: 8px;
  overflow-y: auto;
  background-color: rgba(30, 33, 33, 0.3);
  color: #ffffff;
  scrollbar-width: thin;
  scrollbar-color: #6b7280 #1e1e1e;
}

.log-viewer::-webkit-scrollbar {
  width: 8px;
}

.log-viewer::-webkit-scrollbar-track {
  background: #1e1e1e;
}

.log-viewer::-webkit-scrollbar-thumb {
  background: #6b7280;
  border-radius: 4px;
}

.log-viewer::-webkit-scrollbar-thumb:hover {
  background: #4b5563;
}

.log-viewer > div {
  margin-left: 10px;
  padding: 2px;
}

.card {
  background-color: rgba(30, 33, 33, 0.3);
}

.progress {
  display: flex; 
  flex-direction: column; 
  align-items: center;
  padding-left: 15px;
}
</style>
