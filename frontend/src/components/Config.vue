<!--
 * @Author: Lockly
 * @Date: 2025-02-16 16:58:40
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 14:15:08
 -->
<script setup lang="ts">

import {computed, onMounted, ref} from "vue";
import {useConfigStore} from "@/store/types";

const configState = useConfigStore();
const formData = ref({
  Code: 0, Error: "",
  FilePath: computed({
    get: () => configState.getFilePath(),
    set: (value) => configState.setFilePath(value)
  }),
    CoroutineCount: computed({
      get: () => configState.getCoroutineCount(),
      set: (value) => configState.setCoroutineCount(value)
    }),
    LiveProxies: configState.getLiveProxies(),
    AllProxies: configState.getAllProxies(),
    Timeout: computed({
      get: () => configState.getTimeout(),
      set: (value) => configState.setTimeout(value)
    }),
    SocksAddress: configState.getSocksAddress(),
    Status: 0
})

const rules = {
  CoroutineCount: [
    {
      required: true, 
      message: '请输入协程数量',
      trigger: 'blur'
    },
  ],
  Timeout: [
    {
      required: true,
      message: '请指定超时时长',
      trigger: 'blur'
    },
  ]

}

onMounted(() => {
  configState.getProfile()
})
</script>

<template>
  <a-form ref="formRef" :model="formData" :layout="'vertical'">
    <a-form-item label="文件路径" name="FilePath">
      <a-input :model-value="formData.FilePath" readonly></a-input>
    </a-form-item>
    
    <a-row :gutter="[24, 12]">
      <a-col :span="12">
        <a-form-item label="协程数量" name="CoroutineCount">
          <a-input-number v-model="formData.CoroutineCount" :mode="'button'"></a-input-number>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="超时时长" name="Timeout">
          <a-input-number v-model="formData.Timeout" :mode="'button'"></a-input-number>
        </a-form-item>
      </a-col>
    </a-row>    
    
    <a-row :gutter="[24, 12]">
      <a-col :span="12">
        <a-form-item label="可用代理数量" name="LiveProxies">
          <a-input-number v-model="formData.LiveProxies" read-only>
            <template #suffix>
              条
            </template>
          </a-input-number>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="全部代理数量" name="AllProxies">
          <a-input-number v-model="formData.AllProxies" read-only></a-input-number>
        </a-form-item>
      </a-col>
    </a-row>    
    
    <a-row :gutter="[24, 12]">
      <a-col :span="12">
      </a-col>
      <a-col :span="12">
      </a-col>
    </a-row>
    
    <a-form-item label="Socks地址" name="SocksAddress">
      <a-input v-model="formData.SocksAddress"></a-input>
    </a-form-item>

  </a-form>
</template>

<style scoped>

</style>