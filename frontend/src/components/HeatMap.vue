<template>
  <div class="heatmap-container">
    <div class="controls" v-if="showControls">
      <button @click="generateRandomData" class="btn">随机生成数据</button>
      <button @click="clearData" class="btn">清空数据</button>
      <button @click="showStats" class="btn">显示统计</button>
    </div>

    <div class="months-container">
      <div
        v-for="(monthData, index) in monthsData"
        :key="index"
        class="month-section"
      >
        <div class="month-header">{{ monthData.year }}年{{ monthData.month + 1 }}月</div>

        <div class="heatmap-grid-wrapper">
          <div class="weekdays-header">
            <div
              v-for="(weekday, idx) in weekdays"
              :key="idx"
              class="weekday-label"
            >
              {{ weekday }}
            </div>
          </div>

          <div class="month-grid">
            <div
              v-for="(week, weekIndex) in monthData.weeks"
              :key="weekIndex"
              class="week-column"
            >
              <div
                v-for="(day, dayIndex) in week"
                :key="dayIndex"
                :class="getDayCellClass(day)"
                :data-date="day ? day.date : ''"
                :data-count="day ? day.count : 0"
                @mouseenter="day && showTooltip($event, day)"
                @mouseleave="hideTooltip"
              >
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="legend" v-if="showLegend">
      <span class="legend-label">打卡频率:</span>
      <div class="legend-colors">
        <div
          v-for="(item, index) in legendItems"
          :key="index"
          class="legend-item"
        >
          <div :class="['legend-color', `level-${item.level}`]"></div>
          <span>{{ item.label }}</span>
        </div>
      </div>
    </div>

    <div class="stats" v-if="showStatsInfo" id="stats">
      <div class="stats-grid">
        <div class="stat-item">
          <div class="stat-number">{{ stats.totalDays }}</div>
          <div class="stat-label">总打卡天数</div>
        </div>
        <div class="stat-item">
          <div class="stat-number">{{ stats.totalCount }}</div>
          <div class="stat-label">总打卡次数</div>
        </div>
        <div class="stat-item">
          <div class="stat-number">{{ stats.avgCount }}</div>
          <div class="stat-label">平均每日</div>
        </div>
        <div class="stat-item">
          <div class="stat-number">{{ stats.maxCount }}</div>
          <div class="stat-label">单日最高</div>
        </div>
        <div class="stat-item">
          <div class="stat-number">{{ stats.recent3MonthsDays }}</div>
          <div class="stat-label">近{{ months }}月天数</div>
        </div>
        <div class="stat-item">
          <div class="stat-number">{{ stats.recentAvg }}</div>
          <div class="stat-label">近{{ months }}月平均</div>
        </div>
      </div>
    </div>

    <div
      v-if="showTooltip"
      class="tooltip"
      :style="{ left: tooltipPosition.x + 'px', top: tooltipPosition.y + 'px' }"
    >
      {{ tooltipContent }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'HeatMap',
  props: {
    months: {
      type: Number,
      default: 3
    },
    showControls: {
      type: Boolean,
      default: true
    },
    showLegend: {
      type: Boolean,
      default: true
    },
    showStatsInfo: {
      type: Boolean,
      default: true
    },
    initialData: {
      type: Object,
      default: function() {
        return {}
      }
    }
  },
  data: function() {
    return {
      heatmapData: {},
      monthsData: [],
      weekdays: ['一', '二', '三', '四', '五', '六', '日'],
      legendItems: [
        { level: 0, label: '0次' },
        { level: 1, label: '1次' },
        { level: 2, label: '2次' },
        { level: 3, label: '3次' },
        { level: 4, label: '4+次' }
      ],
      showTooltip: false,
      tooltipContent: '',
      tooltipPosition: { x: 0, y: 0 },
      stats: {
        totalDays: 0,
        totalCount: 0,
        avgCount: 0,
        maxCount: 0,
        recent3MonthsDays: 0,
        recentAvg: 0
      }
    }
  },
  watch: {
    initialData: {
      handler: function(newData) {
        this.heatmapData = Object.assign({}, newData)
        this.generateMonthsHeatmap()
        this.calculateStats()
      },
      immediate: true
    }
  },
  mounted: function() {
    this.generateMonthsHeatmap()
    this.calculateStats()
  },
  methods: {
    // 生成最近几个月的热力图数据
    generateMonthsHeatmap: function() {
      var today = new Date()
      this.monthsData = []

      // 获取最近几个月（包括当前月）
      for (var i = this.months - 1; i >= 0; i--) {
        var monthDate = new Date(today.getFullYear(), today.getMonth() - i, 1)
        var monthData = this.createMonthData(monthDate)
        this.monthsData.push(monthData)
      }
    },

    // 创建单个月份的数据
    createMonthData: function(monthDate) {
      var year = monthDate.getFullYear()
      var month = monthDate.getMonth()
      var firstDay = new Date(year, month, 1)
      var lastDay = new Date(year, month + 1, 0)
      var daysInMonth = lastDay.getDate()

      // 计算第一个星期一是哪一天
      var startDayOfWeek = firstDay.getDay()
      if (startDayOfWeek === 0) startDayOfWeek = 6
      else startDayOfWeek -= 1

      // 按周生成数据
      var currentDay = 1
      var weeks = []

      // 第一周
      var firstWeek = []
      for (var i = 0; i < startDayOfWeek; i++) {
        firstWeek.push(null)
      }
      var daysInFirstWeek = 7 - startDayOfWeek
      for (var i = 0; i < daysInFirstWeek && currentDay <= daysInMonth; i++) {
        var date = new Date(year, month, currentDay)
        var dateStr = this.formatDate(date)
        firstWeek.push({
          date: dateStr,
          day: currentDay,
          count: this.heatmapData[dateStr] || 0
        })
        currentDay++
      }
      weeks.push(firstWeek)

      // 后续几周
      while (currentDay <= daysInMonth) {
        var week = []
        for (var i = 0; i < 7 && currentDay <= daysInMonth; i++) {
          var date = new Date(year, month, currentDay)
          var dateStr = this.formatDate(date)
          week.push({
            date: dateStr,
            day: currentDay,
            count: this.heatmapData[dateStr] || 0
          })
          currentDay++
        }
        // 补充null
        while (week.length < 7) {
          week.push(null)
        }
        weeks.push(week)
      }

      return {
        year: year,
        month: month,
        weeks: weeks
      }
    },

    // 获取日期格子的CSS类
    getDayCellClass: function(day) {
      if (!day) {
        return 'day-cell empty-cell'
      }

      var count = day.count
      var level = 0
      if (count === 1) level = 1
      else if (count === 2) level = 2
      else if (count === 3) level = 3
      else if (count >= 4) level = 4

      return 'day-cell level-' + level
    },

    // 切换打卡状态
    toggle打卡: function(date) {
      if (!this.heatmapData[date]) {
        this.heatmapData[data] = 1
      } else {
        this.heatmapData[date]++
      }

      this.generateMonthsHeatmap()
      this.calculateStats()
      this.$emit('update:heatmapData', Object.assign({}, this.heatmapData))
    },

    // 随机生成数据
    generateRandomData: function() {
      this.heatmapData = {}
      var today = new Date()

      for (var i = this.months - 1; i >= 0; i--) {
        var monthDate = new Date(today.getFullYear(), today.getMonth() - i, 1)
        var year = monthDate.getFullYear()
        var month = monthDate.getMonth()
        var daysInMonth = new Date(year, month + 1, 0).getDate()

        for (var day = 1; day <= daysInMonth; day++) {
          var date = new Date(year, month, day)
          var dateStr = this.formatDate(date)

          if (Math.random() < 0.7) {
            var randomCount = Math.floor(Math.random() * 6)
            if (randomCount > 0) {
              // this.$set(this.heatmapData, dateStr, randomCount)
              this.heatmapData[dateStr] = randomCount
            }
          }
        }
      }

      this.generateMonthsHeatmap()
      this.calculateStats()
      this.$emit('update:heatmapData', Object.assign({}, this.heatmapData))
    },

    // 清空数据
    clearData: function() {
      this.heatmapData = {}
      this.generateMonthsHeatmap()
      this.calculateStats()
      this.$emit('update:heatmapData', {})
    },

    // 显示统计信息
    showStats: function() {
      this.calculateStats()
    },

    // 计算统计数据
    calculateStats: function() {
      var totalDays = Object.keys(this.heatmapData).length
      var totalCount = 0
      var values = Object.values(this.heatmapData)
      for (var i = 0; i < values.length; i++) {
        totalCount += values[i]
      }

      var avgCount = totalDays > 0 ? (totalCount / totalDays).toFixed(1) : 0
      var maxCount = 0
      for (var i = 0; i < values.length; i++) {
        if (values[i] > maxCount) {
          maxCount = values[i]
        }
      }

      // 计算最近几个月的统计数据
      var today = new Date()
      var recentMonthsDays = 0
      var recentMonthsTotal = 0

      for (var i = this.months - 1; i >= 0; i--) {
        var monthDate = new Date(today.getFullYear(), today.getMonth() - i, 1)
        var year = monthDate.getFullYear()
        var month = monthDate.getMonth()
        var daysInMonth = new Date(year, month + 1, 0).getDate()

        for (var day = 1; day <= daysInMonth; day++) {
          var date = new Date(year, month, day)
          var dateStr = this.formatDate(date)
          if (this.heatmapData[dateStr]) {
            recentMonthsDays++
            recentMonthsTotal += this.heatmapData[dateStr]
          }
        }
      }

      var recentAvg = recentMonthsDays > 0 ? (recentMonthsTotal / recentMonthsDays).toFixed(1) : 0

      this.stats = {
        totalDays: totalDays,
        totalCount: totalCount,
        avgCount: avgCount,
        maxCount: maxCount,
        recent3MonthsDays: recentMonthsDays,
        recentAvg: recentAvg
      }
    },

    // 显示工具提示
    showTooltip: function(event, day) {
      if (!day) return

      this.tooltipContent = day.date + '<br>' + day.count + ' 次打卡'
      this.tooltipPosition = {
        x: event.pageX + 10,
        y: event.pageY - 20
      }
      this.showTooltip = true
    },

    // 隐藏工具提示
    hideTooltip: function() {
      this.showTooltip = false
    },

    // 格式化日期为 YYYY-MM-DD
    formatDate: function(date) {
      var year = date.getFullYear()
      var month = String(date.getMonth() + 1)
      if (month.length === 1) month = '0' + month
      var day = String(date.getDate())
      if (day.length === 1) day = '0' + day
      return year + '-' + month + '-' + day
    }
  }
}
</script>

<style scoped>
.heatmap-container {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  margin: 0;
  padding: 20px;
  background-color: #f6f8fa;
  color: #24292e;
  border-radius: 6px;
}

.controls {
  text-align: center;
  margin-bottom: 30px;
}

.btn {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  margin: 0 5px;
}

.btn:hover {
  background-color: #218838;
}

.months-container {
  display: flex;
  gap: 40px;
  justify-content: center;
  flex-wrap: wrap;
  background-color: white;
  padding: 30px;
  border-radius: 6px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12);
}

.month-section {
  min-width: 200px;
}

.month-header {
  text-align: center;
  font-weight: 600;
  margin-bottom: 15px;
  color: #24292e;
  font-size: 16px;
}

.heatmap-grid-wrapper {
  display: flex;
}

.weekdays-header {
  display: grid;
  grid-template-rows: repeat(7, 1fr);
  gap: 2px;
  margin-right: 5px;
}

.weekday-label {
  font-size: 12px;
  color: #586069;
  text-align: center;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.month-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(22px, auto));
  grid-auto-flow: column;
  gap: 2px;
}

.week-column {
  display: grid;
  grid-template-rows: repeat(7, 1fr);
  gap: 2px;
}

.day-cell {
  width: 20px;
  height: 20px;
  border-radius: 2px;
  background-color: #ebedf0;
  cursor: pointer;
  transition: all 0.1s ease;
  border: 1px solid transparent;
}

.day-cell:hover {
  border-color: rgba(0,0,0,0.2);
  transform: scale(1.1);
}

.empty-cell {
  background-color: transparent !important;
  cursor: default;
}

.empty-cell:hover {
  transform: none;
  border: 1px solid transparent;
}

.level-0 { background-color: #ebedf0; }
.level-1 { background-color: #9be9a8; }
.level-2 { background-color: #40c463; }
.level-3 { background-color: #30a14e; }
.level-4 { background-color: #216e39; }

.legend {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 30px;
  font-size: 12px;
  color: #586069;
}

.legend-label {
  margin-right: 10px;
}

.legend-colors {
  display: flex;
  align-items: center;
  gap: 3px;
}

.legend-item {
  display: flex;
  align-items: center;
  margin: 0 5px;
}

.legend-color {
  width: 14px;
  height: 14px;
  border-radius: 2px;
  margin-right: 3px;
}

.stats {
  text-align: center;
  margin: 25px 0;
  padding: 20px;
  background-color: #f6f8fa;
  border-radius: 6px;
  font-size: 14px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 20px;
  max-width: 600px;
  margin: 0 auto;
}

.stat-item {
  padding: 10px;
}

.stat-number {
  font-size: 24px;
  font-weight: bold;
  color: #28a745;
}

.stat-label {
  font-size: 12px;
  color: #586069;
  margin-top: 5px;
}

.tooltip {
  position: fixed;
  background-color: rgba(0,0,0,0.8);
  color: white;
  padding: 8px 12px;
  border-radius: 4px;
  font-size: 12px;
  pointer-events: none;
  z-index: 1000;
  white-space: nowrap;
}
</style>
