<script setup>
import INTJ from '../assets/intj-a.png'
import Avatar from '../assets/42tr.jpg'
</script>

<template>
    <a href="/resume" target="_blank">
        <div class="card">
            <div class="top">
                设计开发一条龙
            </div>
            <div class="middle">
                <div class="avatar-bg">
                    <img class="avatar" :src="Avatar">
                </div>
            </div>
            <div class="middle-text">42tr</div>
            <div class="bottom">
                <a href="https://www.16personalities.com/intj-personality" target="_blank">
                    <img :src="INTJ" style="width: 140px; height: 140px;">
                    <div style="font-size: 20px; position: absolute; bottom: 0; right: 0; width: 140px; text-align: center; color: #9096ee;">建筑师 INTJ-A</div>
                </a>
            </div>
        </div>
    </a>
</template>

<style scoped>
.card {
    position: relative;
    font-size: 35px;
    color: aliceblue;
    height: 380px;
    box-shadow: 0 8px 16px -4px rgba(0,0,0,0.314);
    overflow: hidden;
    opacity: 1;
}
.top {
    min-height: 120px;
    background: linear-gradient(-25deg, #9096ee, rgba(144,150,238,0.867), #9096ee, rgba(144,150,238,0.867));
    background-size: 400%;
    display: block;
    unicode-bidi: isolate;
    font-size: 14px;
    text-align: center;
    padding-top: 20px;
}
.middle {
    width: 114px;
    height: 114px;
    right: calc(50% - 57px);
    top: 75px;
    border-radius: 500px;
    position: absolute;
    transition: cubic-bezier(.69,.39,0,1.21) .3s;
    transform-origin: bottom;
    border: 7px solid #1b1c20;
    display: flex;
    justify-content: center;
}
.avatar-bg {
    overflow: hidden;
    border-radius: 500px;
    width: 114px;
    height: 114px;
    background: #1b1c20;
}
.avatar-bg:hover {
    transform: scale(1.1);
}
.middle-text {
    margin-top: 60px;
    text-align: center;
    font-weight: 700;
    color: aliceblue;
    font-size: 20px;
    line-height: 1;
    margin-bottom: 5px;
}
.bottom {
    position: absolute;
    bottom: 0;
    left: 0;
    height: 140px;
    width: 100%;
}
.avatar {
    object-fit: cover;
    width: 100%;
    height: 100%;
    cursor: pointer;
}
img {
    transition: transform 0.5s ease;
}
img:hover {
    transform: scale(1.1) rotate(10deg);
}
</style>
