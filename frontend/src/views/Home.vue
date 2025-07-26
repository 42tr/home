<script setup>
// import HelloWorld from './components/HelloWorld.vue'
import { NLayout, NLayoutHeader, NLayoutContent, NSpace } from 'naive-ui'
import SelfInfo from '../components/SelfInfo.vue'

import { onMounted, onUnmounted } from 'vue'

onMounted(() => {
  // 动态创建 script 标签
  const script = document.createElement('script')
  script.src = '/universe.min.js'
  script.onload = () => {
    // 脚本加载完成后的回调
    console.log('universe.js loaded')
  }
  document.head.appendChild(script)

  // 组件卸载时移除脚本（可选）
  onUnmounted(() => {
    document.head.removeChild(script)
  })
})
</script>

<template>
    <div class="bg"></div>
    <div class="gradient"></div>
    <n-space vertical size="large">
        <n-layout>
            <n-layout-header>
                <div>
                    <p class="logo">42tr</p>
                </div>
            </n-layout-header>
            <n-layout-content>
                <SelfInfo />
            </n-layout-content>
        </n-layout>
    </n-space>

    <canvas id="universe"></canvas>
</template>

<style scoped>
.logo {
    font-size: 24px;
    color: white;
    margin-left: 30px;
}
.n-layout {
    background-color: transparent !important;
}
.n-layout-header {
    background-color: transparent !important;
    color: white;
    height: 64px;
}
.n-layout-content {
    background-color: transparent !important;
    height: calc(100vh - 64px);
}
.bg {
    z-index: -2;
    background-image: url('/wallhaven-l871yl.png');
    visibility: visible;
    border: none;
    height: 100%;
    pointer-events: none;
    position: fixed;
    top: 0;
    width: 100%;
    background-position: center center;
    background-repeat: no-repeat no-repeat;
    background-size: cover;
}
.gradient {
    background: radial-gradient(transparent 50%, black), linear-gradient(transparent, black);
    opacity: 0.2;
    position: fixed;
    top: 0;
    height: 100%;
    width: 100%;
    z-index: -1;
}
#universe {
    display: block;
    position: fixed;
    margin: 0;
    padding: 0;
    border: 0;
    outline: 0;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    pointer-events: none;
    z-index: 1;
}
</style>
