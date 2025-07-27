<script setup>
defineProps({
  src: String,
});
</script>

<template>
    <div class="card">
        <img :src="src" style="width: 100%; height: 100%; object-fit: cover;">
    </div>
</template>

<style scoped>
.card {
    height: 530px;
}
</style>
