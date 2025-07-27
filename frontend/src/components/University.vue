<script setup>
import UESTC from '../assets/uestc.svg';
</script>

<template>
    <div class="card">
        <img :src="UESTC" style="width: 80px;">
        <div class="text">
            <p>本科 2014 ~ 2018</p>
            <p>软件工程</p>
        </div>
    </div>
</template>

<style scoped>
.card {
    height: 100px;
    padding: 10px;
    display: flex;
    font-size: 14px;
    background:
        radial-gradient(white, transparent);
}
.text {
    margin-left: 20px;
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
}
p {
    margin: 0;
}
</style>
