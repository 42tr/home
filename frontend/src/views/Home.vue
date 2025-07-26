<script setup>
import { onMounted, onUnmounted, ref } from 'vue'
import { NLayout, NLayoutHeader, NLayoutContent, NSpace, NRow, NCol } from 'naive-ui'
import Masonry from 'masonry-layout'
import SelfInfo from '../components/SelfInfo.vue'
import Leetcode from '../components/Leetcode.vue'

const container = ref(null)
onMounted(() => {
  // 加载背景光效
  const script = document.createElement('script')
  script.src = '/universe.min.js'
  script.onload = () => {
    console.log('universe.js loaded')
  }
  document.head.appendChild(script)
  onUnmounted(() => {
    document.head.removeChild(script)
  })
  // 瀑布流展示卡片
  var masonryInstance = new Masonry(container.value, {
    itemSelector: '.masonry-item',
    columnWidth: '.masonry-item',
    percentPosition: true,
    gutter: 16
  })
  onUnmounted(() => {
    if (masonryInstance) {
      masonryInstance.destroy()
    }
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
                <div ref="container" class="content masonry-wrapper">
                    <SelfInfo class="masonry-item"/>
                    <Leetcode class="masonry-item"/>
                </div>
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
    width: 100vw;
}
.n-layout-content {
    background-color: transparent !important;
    height: calc(100vh - 64px);
    width: 100vw;
    padding-left: 30px;
    padding-right: 30px;
}
.content {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 8px;
}
.bg {
    z-index: -2;
    background-image: url('/wallhaven-l871yl.png');
    visibility: visible;
    border: none;
    height: 100vh;
    pointer-events: none;
    position: fixed;
    top: 0;
    width: 100vw;
    background-position: center center;
    background-repeat: no-repeat no-repeat;
    background-size: cover;
}
.gradient {
    background: radial-gradient(transparent 50%, black), linear-gradient(transparent, black);
    opacity: 0.2;
    position: fixed;
    top: 0;
    height: 100vh;
    width: 100vw;
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
    width: 100vw;
    height: 100vh;
    pointer-events: none;
    z-index: 1;
}
.masonry-wrapper {
  display: flex;
  flex-wrap: wrap;
  margin: -8px;
}
.masonry-item {
  margin: 8px;
}
</style>
