<script setup>
import { onMounted, ref } from 'vue'
import http from '../api/http'

var data = ref({})
onMounted(() => {
  http.get('/api/leetcode').then((response) => {
    console.log(response)
    data.value = response
  })
})
</script>

<template>
  <div class="card">

    <div class="card-body">

      <div class="stat-item">
        <span class="label">LeetCode 竞赛分数：</span>
        <span class="value">{{ data.rating }}</span>
      </div>

      <div class="stat-item">
        <span class="label">全球排名：</span>
        <span class="value">{{ data.globalRanking }} / {{ data.globalTotalParticipants }}</span>
      </div>

      <div class="stat-item">
        <span class="label">全国排名：</span>
        <span class="value">{{ data.localRanking }} / {{ data.localTotalParticipants }}</span>
      </div>

      <div class="stat-item">
        <span class="label">已解答：</span>
        <span class="value">{{ data.questionSolved }} / {{ data.questionTotal }}</span>
      </div>

    </div>
  </div>
</template>

<style scoped>
.card {
  box-sizing: border-box;
  background-color: #000000;
  color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.1);
  padding: 24px;
  text-align: center;
  height: 254px;
}

.leetcode-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
}

.card-header {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 20px;
}

.card-body {
  display: flex;
  flex-direction: column;
  gap: 12px;
  text-align: left;
}

.stat-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #f0f0f0;
}

.stat-item:last-child {
  border-bottom: none;
}

.label {
  font-size: 0.95rem;
  font-weight: 500;
}

.value {
  font-size: 1rem;
  font-weight: 600;
}
</style>
