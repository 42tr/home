<script setup>
// import HelloWorld from './components/HelloWorld.vue'
import { NLayout, NLayoutHeader, NLayoutContent, NSpace } from 'naive-ui'
import SelfInfo from './components/SelfInfo.vue'
</script>

<template>
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
</style>
