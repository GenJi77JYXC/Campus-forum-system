<template>
    <div class="common-layout">
        <el-container class="header">
            <el-header style="text-align: right; font-size: 12px; height: 40px; border-radius: 2px">
                <a>登录</a>
                <a>注册</a>
            </el-header>
            <!-- <hr align=center width=100% color=#987cb9 SIZE=1> -->
            <!-- <el-divider /> -->
            <Menu></Menu>
        </el-container>
        <el-container>
            <el-aside width="200px" class="aside"><HomeAside></HomeAside></el-aside>
            <el-main class="main">
                <MainTop></MainTop>
                <MainContent></MainContent>
            </el-main>
        </el-container>

    </div>
</template>

<script setup lang="ts">
import Carousel from '@/components/Carousel.vue'
import HomeAside from '@/components/HomeAside.vue';
import Menu from './menu/index.vue'
import MainTop from '@/components/MainTop.vue';
import MainContent from '@/components/MainContent.vue';
</script>

<script lang="ts">
export default {
    name : "Layout",
}
</script>

<style scoped>
.el-carousel__item h3 {
    display: flex;
    color: #475669;
    opacity: 0.75;
    line-height: 300px;
    margin: 0;
}

.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
    background-color: #d3dce6;
}

.common-layout {
    background-color: palegreen;
    width: 100%;
    height: 100vh;
}


.header {
    display: flex;
    justify-content: center;
    height: 200px;
    border-style: solid;
    border-color: black;
    border-radius: 1px;
}

.header a {
    display: inline-block;
    margin-right: 22px;
    line-height: 34px;
}

.carousel {
    border-style: solid;
    border-color: black;
    border-radius: 1px;
}

.aside {
    background-color: bisque;
}

.main {

    background-color: aqua;
}
</style>