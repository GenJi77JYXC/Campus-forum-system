<template>
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" :ellipsis="false" @select="handleSelect">
        <el-menu-item index="0">
            <img style="width: 50px" src="@\assets\logo.svg" alt="Element logo" />
        </el-menu-item>
        <div class="center">
            <el-input v-model="input" style="width: 400px" placeholder="Please input" clearable />
        <el-button type="primary">进入贴吧</el-button>
        </div>
    </el-menu>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const input = ref('')

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}

</script>
<script lang="ts">
export default {
    name: "Menu",
}
</script>


<style scoped>
</style>