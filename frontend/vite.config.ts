/*
 * @Author: Lockly
 * @Date: 2025-02-12 22:54:13
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 14:13:50
 */
import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path';


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: [
      {
        find: '@',
        replacement: resolve(__dirname, './src')
      }
    ]
  }

})
