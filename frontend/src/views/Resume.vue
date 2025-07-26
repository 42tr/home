<script setup>
import { NIcon, NThing, NButton, NTimeline, NTimelineItem, NSpace, NTag, NDivider } from 'naive-ui'
import { Email, Phone, Home, Certificate } from '@vicons/carbon'
import Avatar from '../assets/42tr.jpg'
import UESTC from '../assets/uestc.svg'
import asiainfo from '../assets/asiainfo.ico'
import Rich from '../assets/rich.png'
import Huawei from '../assets/huawei.png'
import Hontor from '../assets/hontor.png'

var open = (url) => {
  window.open(url, '_blank')
}

var mail = () => {
  window.location.href = 'mailto:1055894396@qq.com'
}
</script>

<template>
    <div class="body">
        <!-- header start -->
        <div style="position: absolute; right: 30px; top: 20px; z-index: 10;">
            <img :src="Avatar" style="height: 100px;">
        </div>
        <div style="text-align: center;">
            <h1>刘阳</h1>
            <n-divider style="margin-top: 0; margin-bottom: 0; width: 70%; left: 15%;" />
            <n-thing>
                <n-button size="small" :bordered="false">
                    <template #icon>
                        <n-icon><Phone /></n-icon>
                    </template>
                    18936861530
                </n-button>
                <n-button size="small" :bordered="false" @click="mail()">
                    <template #icon>
                        <n-icon><Email /></n-icon>
                    </template>
                    1055894396@qq.com
                </n-button>
                <n-button size="small" :bordered="false" @click="open('https://42tr.cn')">
                    <template #icon>
                        <n-icon><Home /></n-icon>
                    </template>
                    42tr.cn
                </n-button>
                <n-button size="small" :bordered="false">
                    <template #icon>
                        <n-icon><Certificate /></n-icon>
                    </template>
                    软考-软件设计师
                </n-button>
            </n-thing>
        </div>
        <!-- header end -->
        <!-- body start -->
        <h2>教育</h2>
        <n-timeline>
            <n-timeline-item title="软件工程-本科 | 电子科技大学" content="学习的主要课程：软件工程、计算机组成原理、计算机网络、嵌入式系统、数据结构与算法、机器视觉算法与应用、综合课程设计（Android & 人脸识别）等。获得英语四级证书和人民一等奖学金。">
                <template #icon>
                    <n-icon>
                        <img :src="UESTC" style="height: 16px; width: 16px;">
                    </n-icon>
                </template>
                <template #footer>
                    2014.09 - 2018.06
                </template>
            </n-timeline-item>
        </n-timeline>
        <h2>经历</h2>
        <n-timeline>
            <n-timeline-item type="success" title="中科南京人工智能创新研究院">
                <template #icon><n-icon><img :src="Hontor" style="height: 16px; width: 16px;"></n-icon></template>
                <p><b>智能办公助手 2025.03 ~</b></p>
                <p>基于 Qwen、Deepseek 等大语言模型，搭配 CosyVoice、SenseVoice、Mineru 等模型，实现公文撰写、PPT 生成、脑图生成、智能会议等功能。</p>
            </n-timeline-item>
            <n-timeline-item type="success" title="华为">
                <template #icon><n-icon><img :src="Huawei" style="height: 16px; width: 16px;"></n-icon></template>
                <p><b>维护辅助工具 2024.01 ~ 2024.11</b></p>
                <p>使用 SpringBoot + MyBatis + Vue2 + ElementUI + Echarts + Langchain 框架开发，用于根据用户提问，从历史数据中查找相似案例，并使用大模型通过 RAG 方式回答用户问题，并接入内部聊天机器人。作为整体设计参与开发。数据库使用 Mysql，分词搜索引擎使用 ElasticSearch，语义搜索使用内部引擎。工作成果：设计整体架构，搭建前后端框架，优化 ES 的搜索效果，并提出多路召回重排序方法，问题搜索准确率由 50% 提升至 90%，获“数据通信研发部长奖-维护效率及能力提升奖”，并成功申请专利。应用使用 Docker 方式部署。</p>
                <p><b>AR 接入路由器 2021.09 ~ 2023.12</b></p>
                <p>使用 C 语言参与 SSH/TELNET/FTP/NETCONF 等协议、支撑与 NCE 交互、补丁机制以及系统整体相关开发与维护工作。作为 CIE：使用 Jenkins 维护并优化门禁，提高 MR 合入效率 10% 以上；设计并搭建构建镜像，负责工具链迭代、镜像升级；优化版本 CI，提高构建发布效率 20% 以上；开发个人构建脚本，提高开发效率 20% 以上。使用 Python + Bootstrap 维护补丁网站，支撑其他开发快速构建补丁以及解决补丁编译问题。使用 Golang 的 Gin 框架开发查重网站，将补丁查重后台完全自动化，每年减少开发投入 2 人月。使用 Rust 开发自动化补丁代码检查工具，解决了人工查找隐式函数转换导致的寻址异常问题。作为 Redis 的 PMC 成员，参与产品代码定制化归一与升级工作，获“redis 定制归一攻关激励团队奖”。作为开发，协助国外客户处理设备相关问题。</p>
                <p></p>
            </n-timeline-item>
            <n-timeline-item title="北京睿企">
                <template #icon><n-icon><img :src="Rich" style="height: 16px; width: 16px;"></n-icon></template>
                <p><b>NLP 应用 2020.05 ~ 2021.07</b></p>
                <p>使用 SpringBoot + MyBatis + React + Antd + Echarts 框架开发，用于公安行业的大数据处理。作为开发与设计人员，负责开发大数据标准化处理模块，维护标准化处理流程，设计与实现质态监管、热词雷达、自动报告等各个应用的后端。对接其他公司，使用内部的 ETL 工具对数据进行处理，偶尔帮忙改改前端 bug 以及写一些简单的管理页面。数据存储主要使用 TiDB 和 Minio，容器管理使用 Kubernetes，搜索使用 ElasticSearch。</p>
            </n-timeline-item>
            <n-timeline-item title="亚信科技">
                <template #icon><n-icon><img :src="asiainfo" style="height: 16px; width: 16px;"></n-icon></template>
                <!-- <p><b>营业云平台 2020.03 ~ 2020.05</b></p>
                <p>营业员使用的 windows 桌面应用项目改造，前端主要使用 Electron + Aurelia + NWJS，该技术选型目的在于支持营业员使用较多的 XP 系统。完成的工作包括搭建测试环境、本地调试应用、使用 InnoSetup 对工程进行打包。项目中使用 emqtt(emqx) 实现了消息推送。改造内容包括前端优化消息展现，后端跨中心时通过使用消息队列 kafka 的方式同步员工信息、发送消息时使用 DFA 算法检测消息中是否包含敏感字符等。</p>
                <p><b>电信 CRM3.0 2018.09 ~ 2019.12</b></p>
                <p>电信运营商服务升级项目，前端使用 Vita-Plus，后端使用 SpringMVC/SpringBoot + MyBatis。主要负责核心部分营业受理开发与吉林和江苏本地化部分开发内容，包括前端读卡器接入与单点登录对接、Nginx反向代理等问题。</p>
                <p><b>EOP 平台 2018.03 ~ 2018.08</b></p>
                <p>CRM3.0 项目，能力开放平台，用于对接入的接口重新封装，使其获得 EOP 提供的流控、路由等能力，主要工作内容为接口封装与压测、流控、路由、协助问题排查、Jenkins 测试发布以及 Docker 环境问题排查等。</p> -->
                <p><b>电信 CRM3.0 2018.03 ~ 2020.05</b></p>
                <p>作为全栈工程师，负责前后端开发、数据库设计、部署等任务，后端使用 SpringBoot + Mybatis，常用中间件包括 Redis、Kafka、Mysql、Nginx 等，前端使用 Vita-Plus、Electron。</p>
            </n-timeline-item>
        </n-timeline>
        <h2>个人评价</h2>
        <div style="text-indent:28px">
            全栈工程师，熟悉 linux、Windows、Mac 操作系统，熟练 Java、Go、JavaScript 开发应用，熟悉 C、Rust 系统开发，熟练使用 Github Actions/Jenkins 等工具实现 CI/CD。算法能力较强，平时喜欢研究技术，业余时间会做一些基于树莓派的应用开发。
        </div>
        <!-- body end -->
    </div>
</template>

<style scoped>
.body {
    padding: 20px;
}
</style>
