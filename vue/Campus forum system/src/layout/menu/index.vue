<template>
    <div class="search">
        <ElInput v-model="input1" style="width: 300px; height: 60px; "></ElInput>
    </div>
</template>

<script setup lang="ts">
import { ElInput } from 'element-plus';
import { ref } from 'vue'

const input1 = ref('')

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
.search{
    height: 100px;
    width: auto;
    border-color: red;
    border-radius: 1px;
    position: flex;
}
</style>