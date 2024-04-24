<template>
    <div>
        <el-card>
    <template #header>
      <div class="card-header">
        <span>贴吧精选专题</span>
      </div>
    </template>
    <p>高校精选专题</p>
  </el-card>
    </div>
    <el-divider content-position="left">贴吧分类</el-divider>
    <div>
        <el-card style="max-width: 480px">
    <template #header>
      <div class="card-header">
        <span>Card name</span>
      </div>
    </template>
    <p v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</p>
  </el-card>
    </div>
    <div>
        <el-card style="max-width: 480px">
    <template #header>
      <div class="card-header">
        <span>贴吧精选专题</span>
      </div>
    </template>
    <p v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</p>
  </el-card>
    </div>
</template>

<script setup lang="ts">

</script>

<script lang="ts">
export default {
    name: "HomeAside",
}
</script>

<style scoped>

</style>