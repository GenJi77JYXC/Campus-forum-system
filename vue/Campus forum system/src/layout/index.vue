<template>
    <div class="common-layout">
        <el-container class="header">
            <el-header style="text-align: right; font-size: 12px; height: 40px; border-radius: 2px">
                <a>登录</a>
                <a>注册</a>
            </el-header>
            <hr align=center width=100% color=#987cb9 SIZE=1>
            <Menu></Menu>
            <Header></Header>
        </el-container>
        <el-container class="bottom">
            <el-aside width="200px" class="aside">Aside</el-aside>
            <el-main class="main">Main</el-main>
        </el-container>
    </div>
</template>

<script setup lang="ts">
import Header from './header/index.vue'
import Menu from './menu/index.vue'

</script>

<style scoped>
.common-layout {

    background-color: palegreen;
    width: 100%;
    height: 100vh;
}

.header {
    display: flex;
    justify-content: center;
    height: 300px;
}

.header a {
    display: inline-block;
    margin-right: 22px;
    line-height: 34px;
}

.aside {
    height: calc(100vh - 300px);
    background-color: bisque;
}

.main {
    background-color: aqua;
}
.menu{
    height: 100px;
    width: auto;
    border-color: red;
    border-radius: 1px;
    position: flex;
}
</style>