<template>
    <div class="search">
        <router-link to="/layout"><img src="@/assets/校园论坛.png" style="width: 100px; height: 60px; margin-right: 10px;"
                alt="校园论坛"></router-link>
        <el-input v-model="input1" style="max-width: 600px" placeholder="请输入搜索内容" class="input-with-select"
            size="large">
            <template #append>
                <el-button type="primary" size="large" @click="handleSearch">搜索</el-button>
            </template>
        </el-input>
    </div>
</template>
<!-- style="width: 400px; height: 50px; left: 40%; transform: translateX(-50%); margin-top: 8px;"
            placeholder=";" -->
<script setup lang="ts">
import { ElInput } from 'element-plus';
import { ref } from 'vue'


const input1 = ref('')


const handleSearch = () => {
    console.log(input1.value)
}

</script>
<script lang="ts">
export default {
    name: "Menu",
}
</script>


<style scoped>
.search {
    height: 100px;
    border-style: solid;
    border-color: red;
    border-radius: 1px;
    display: flex;
    justify-content: center;
    /* 水平居中 */
    align-items: center;
    /* 垂直居中 */
    height: 100vh;
    /* 保证垂直居中在视口中 */

}

.input-with-select {
    width: 100%;
    height: 50px;
}


.carousel-container img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}
</style>