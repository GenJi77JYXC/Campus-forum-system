<template>
  <el-card body-style="font-size: 14px; padding: 0; ">
    <h4 class="card-title">贴吧精选专题</h4>
    <div class="item_card">高校精选专题</div>
  </el-card>
  <el-divider content-position="left" class="divider">贴吧分类</el-divider>
  <div>
    <el-card style="max-width: 480px" body-style="font-size: 14px; padding: 0; margin: 10px;">
      <div class="card-header" style="font-size: 16px; ">
        <h4 class="card-title">title</h4>
      </div>
      <div class="card-body">
        <a>item1</a>
        <a>item2</a>
        <a>item3</a>
        <a>item4</a>
      </div>
    </el-card>
  </div>
  <div>
    <el-card style="max-width: 480px" body-style="font-size: 14px; padding: 0; margin: 10px;">
      <div class="card-header" style="font-size: 16px;">
        <h4 class="card-title">title</h4>
      </div>
      <div class="card-body">
        <a>item1</a>
        <a>item2</a>
        <a>item3</a>
        <a>item4</a>
      </div>
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
.card-title {
  color: #444;
  font-size: 16px;
  padding: 0;
}

.item_card {
  display: inline-block;
  padding: 4px 4px 2px 2px;
  color: #999;
}

.divider {
  margin: 10px 0;
  height: 1px;
  background-color: black;
}

.card-body a {
  display: inline-block;
  padding: 4px 4px 2px 2px;
  color: #999;
}
</style>