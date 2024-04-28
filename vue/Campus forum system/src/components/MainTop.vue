<template>
    <div class="main-top">
        <div class="forum">
            <div class="class_title">
                <div>
                    热门论坛
                </div>
            </div>
            <div class="class_list">
                <ul class="class_ul">
                    <li class="class_li">
                        <a>item1
                            <div class="class_info">
                                <div class="forum_name">name</div>
                                <div class="forum_member">member</div>
                                <div class="forum_comment">comment</div>
                            </div>
                        </a>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

</script>
<script lang="ts">
export default {
    name: "MainTop"
}
</script>

<style scoped>
.main-top {
    margin-bottom: 15px;
}

.forum {
    padding: 16px 16px 0;
    width: 780px;
    background: #fff;
    border: 1px solid rgba(230, 230, 230, 1);
    box-sizing: border-box;
    font-family: STHeiti, "Microsoft Yahei", "Microsoft YaHei", Arial, sans-serif;
}
.forum .class_title {
    font-size: 16px;
    color: #000;
    font-weight: 700;
    margin-bottom: 16px;
    display: -ms-flexbox;
    display: flex;
    -ms-flex-align: center;
    align-items: center;
    -ms-flex-pack: justify;
    justify-content: space-between;
}
</style>