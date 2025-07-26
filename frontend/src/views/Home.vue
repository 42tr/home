<script setup>
import { onMounted, onUnmounted, ref } from 'vue'
import { NLayout, NLayoutHeader, NLayoutContent, NSpace, NRow, NCol } from 'naive-ui'
import Masonry from 'masonry-layout'
import SelfInfo from '../components/SelfInfo.vue'
import Leetcode from '../components/Leetcode.vue'
import Bookmark from '../components/Bookmark.vue'

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

const bookmarks = [
  {
    title: '博客',
    url: 'https://42tr.cn',
    tags: ['42tr'],
  },
  {
    title: 'GitHub',
    url: 'https://github.com/42tr',
    tags: ['42tr'],
  },
  {
    title: '任务管理',
    url: 'https://tasks.42tr.com',
    tags: ['42tr', 'task'],
  },
  {
    title: '简历管理',
    url: 'https://talents.42tr.com',
    tags: ['42tr', 'talent'],
  },
  {
    title: '禅道',
    url: 'http://222.190.139.186:8888',
    tags: ['bug'],
  },
  {
    title: 'Gitlab',
    url: 'http://222.190.139.186:9999',
    tags: ['code'],
  },
  {
    title: '内部知识分享平台',
    url: 'http://192.168.0.46:8002',
    tags: ['knowledge'],
  },
]
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
                    <Bookmark class="masonry-item"
                      v-for="(item, index) in bookmarks"
                      :key="index"
                      :bookmark="item"
                    />
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
    padding-top: 10px;
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
  transition: transform 0.5s ease;
}
.masonry-item:hover {
  transform: translateX(2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.3);
  /* border-color: #CC9900; */
}
</style>
