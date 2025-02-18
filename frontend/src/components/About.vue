<!--
 * @Author: Lockly
 * @Date: 2025-02-16 17:25:26
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 15:12:25
 -->
<script setup lang="ts">

import {BrowserOpenURL} from "../../wailsjs/runtime";

const url = `https://github.com/BKLockly/proxyz`
const comments = [
  { content: '在使用本工具之前，请务必了解并遵守以下事项：', sub: [
      { content: '仅将本工具用于合法授权的测试环境或您拥有明确权限的系统。未经授权的使用可能违反法律。' },
      { content: '本工具按“原样”提供，不提供任何形式的明示或暗示的担保，包括但不限于适用性、可靠性或安全性。' },
      { content: '开发者不对因使用本工具导致的任何数据丢失、系统损坏、法律责任或其他后果负责。' },
      { content: '使用者需确保其行为符合所在国家或地区的法律法规。开发者不对任何非法使用行为负责。' },
    ] 
  },
  { content: '如果您使用本工具，即表示您已阅读、理解并同意本免责声明的所有条款。' },
]
</script>

<template>
  <a-row :gutter="6">
    <a-col :span="6">
      <a-tag checkable color="cyan" :default-checked="true">
        <template #icon>
          <icon-user/>
        </template>
        开发者: &nbsp;<span>Lockly</span>
      </a-tag>
    </a-col>
    <a-col :span="7">
      <a-tag checkable color="blue" :default-checked="true">
        <template #icon>
          <icon-info-circle/>
        </template>
        版本: &nbsp;<span>v1.0.0</span>
      </a-tag>
    </a-col>
    <a-col :span="11" >
      <a-tooltip content="点击跳转项目主页" position="right">
        <a-tag color="arcoblue" @click="BrowserOpenURL(url)">
          <template #icon>
            <icon-github/>
          </template>
          项目地址: &nbsp;<span>{{ url }}</span>
        </a-tag>
      </a-tooltip>
    </a-col>
  </a-row>



  
  <a-comment
    v-for="(comment, index) in comments"
    :key="index"
    author="Lockly"
    :content="comment.content"
    style="margin-top: 20px"
  >
    <template #avatar>
      <img src="https://avatars.githubusercontent.com/u/113906689?v=4" alt="avatar">
    </template>
    <template v-if="comment.sub">
      <a-comment
        v-for="(subComment, subIndex) in comment.sub"
        :key="subIndex"
        author="Lockly"
        :content="subComment.content"
      >
        <template #avatar>
          <img src="https://avatars.githubusercontent.com/u/113906689?v=4" alt="avatar">
        </template>
      </a-comment>
    </template>
  </a-comment>
</template>

<style scoped>

</style>