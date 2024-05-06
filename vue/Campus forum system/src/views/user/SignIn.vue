<template>
    <section class="main">
        <div class="container">
            <div class="main-body no-bg">
                <div class="widget signin">
                    <div class="widget-header">登录</div>
                    <div class="widget-content">
                        <div class="feild">
                            <label class="label">用户名/邮箱</label>
                            <div class="control has-icons-left">
                                <span class="icon is-small is-left">
                                    <el-icon :size="20"><User /></el-icon>
                                </span>
                                <input type="text" v-model="username" placeholder="请输入用户名/邮箱" class="input is-success"
                                    @keyup.enter="submitLogin" />
                                
                            </div>
                        </div>

                        <div class="feild">
                            <label class="label">密码</label>
                            <div class="control has-icons-left">
                                <span class="icon is-small is-left">
                                    <el-icon><Lock /></el-icon>
                                </span>
                                <input type="password" v-model="password" placeholder="请输入密码" class="input is-success"
                                    @keyup.enter="submitLogin" />
                                
                            </div>
                        </div>

                        <div class="feild login-button">
                            <button class="button is-success" @click="submitLogin">登录</button>
                            <router-link to="/user/signup" class="button is-text">没有账号？点击这里去注册&gt;&gt;</router-link>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const username = ref('');
const password = ref('');

const submitLogin = () => {

}


</script>

<style scoped lang="scss">
.signin {
    max-width: 480px;
    margin: auto;
    padding: 0 20px;

    .login-captcha-input {
        width: 100%;
        margin-right: 20px;

        .input {
            width: 100% !important;
        }
    }

    .login-captcha-img {
        img {
            height: 40px;
        }
    }

    .login-button {
        .button {
            width: 100%;
        }
    }

    .third-party-line {
        border-bottom: 1px solid #dedede;
        margin-bottom: 24px;

        .third-party-title {
            margin-bottom: -12px;
            text-align: center;

            span {
                background-color: #fff;
                padding: 0 10px;
                font-size: 13px;
            }
        }
    }

    .third-parties {
        text-align: center;
        margin-bottom: 10px;
    }
}
</style>