<template>
  <a
    :href="bookmark.url"
    target="_blank"
    @click="$emit('click', bookmark)"
  >
    <div class="card">
      <!-- 图标 + 名称 -->
      <div class="bookmark-header">
        <h3 class="bookmark-title">{{ bookmark.title }}</h3>
        <!-- 标签 -->
        <div v-if="bookmark.tags && bookmark.tags.length" class="bookmark-tags">
          <span
            v-for="tag in bookmark.tags"
            :key="tag"
            class="tag"
            :style="{ backgroundColor: getTagColor(tag) }"
          >
            {{ tag }}
          </span>
        </div>
      </div>

      <!-- 网址 -->
      <p class="bookmark-url">{{ formatUrl(bookmark.url) }}</p>



      <!-- 收藏时间 -->
      <!-- <p class="bookmark-date">收藏于 {{ formatDate(bookmark.date) }}</p> -->
    </div>

    <!-- 悬浮图标：外部链接 -->
    <div class="external-icon">↗</div>
  </a>
</template>

<script>
export default {
  name: 'BookmarkCard',
  props: {
    bookmark: {
      type: Object,
      required: true,
      validator: (val) =>
        val.title && val.url,
    },
  },
  methods: {
    // 格式化网址（去掉 https:// 和 www）
    formatUrl(url) {
      return url.replace(/^https?:\/\//, '').replace(/^www\./, '').split('/')[0];
    },

    // 格式化日期
    formatDate(date) {
      const d = new Date(date);
      return d.toLocaleDateString('zh-CN', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
      });
    },

    // 为标签生成颜色（使用暗黄系）
    getTagColor(tag) {
      const colors = {
        bug: '#CC9900',     // 暗黄
        knowledge: '#CC9900',
        code: '#66BB6A',     // 绿
        backend: '#42A5F5',      // 蓝
        tool: '#FFB74D',         // 橙黄
        docs: '#AB47BC',         // 紫
        task: '#FF5722',         // 红
        talent: '#FF6F00',       // 橙
      };
      return colors[tag.toLowerCase()] || '#777777';
    },
  },
};
</script>

<style scoped>
.card {
  transition: all 0.2s ease;
  position: relative;
  text-decoration: none;
  padding: 16px;
  /* background: #1e1e1e; */
  background: radial-gradient(transparent 50%, black), linear-gradient(transparent, black);
  display: flex;
  flex-direction: column;
  gap: 6px;
  height: 100px;
}

/* 标题行 */
.bookmark-header {
  display: flex;
  align-items: center;
  gap: 10px;
}

.favicon {
  width: 16px;
  height: 16px;
  background-size: contain;
  background-position: center;
  background-repeat: no-repeat;
}

.bookmark-title {
  font-size: 1.1rem;
  font-weight: 600;
  color: #ffffff;
  margin: 0;
  flex: 1;
}

/* 网址 */
.bookmark-url {
  font-size: 0.9rem;
  color: #b0b0b0;
  margin: 0;
  font-family: monospace, sans-serif;
}

/* 标签 */
.bookmark-tags {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
  margin: 4px 0;
}

.tag {
  font-size: 0.75rem;
  padding: 2px 6px;
  border-radius: 6px;
  color: white;
  background-color: #808000;
  opacity: 0.9;
}

/* 日期 */
.bookmark-date {
  font-size: 0.8rem;
  color: #888;
  margin: 0;
}

/* 外部链接图标 */
.external-icon {
  position: absolute;
  top: 16px;
  right: 16px;
  font-size: 0.9rem;
  color: #CC9900;
  opacity: 0;
  transition: opacity 0.2s ease;
}

a:hover .external-icon {
  opacity: 1;
}

/* 响应式 */
@media (max-width: 480px) {
  .bookmark-card {
    padding: 14px;
  }
  .bookmark-title {
    font-size: 1rem;
  }
  .bookmark-url {
    font-size: 0.85rem;
  }
}
</style>
