<template>
    <section class="main">
    <div class="container">
      <div class="main-body no-bg">
        <div class="widget signup">
          <div class="widget-header">注册</div>
          <div class="widget-content">
            <div class="field">
              <label class="label">昵称</label>
              <div class="control has-icons-left">
                <input
                  v-model="username"
                  class="input is-success"
                  type="text"
                  placeholder="请输入昵称"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-username" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">邮箱</label>
              <div class="control has-icons-left">
                <input
                  v-model="email"
                  class="input is-success"
                  type="text"
                  placeholder="请输入邮箱"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-email" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">密码</label>
              <div class="control has-icons-left">
                <input
                  v-model="password"
                  class="input"
                  type="password"
                  placeholder="请输入密码"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">确认密码</label>
              <div class="control has-icons-left">
                <input
                  v-model="rePassword"
                  class="input"
                  type="password"
                  placeholder="请再次输入密码"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>

            <div class="field">
              <div class="control">
                <button class="button is-success" @click="signup">注册</button>
                <github-login :ref-url="ref" />
                <qq-login :ref-url="ref" />
              </div>
            </div>

            <div class="field">
              <router-link to="/user/signin">已有账号，前往登录&gt;&gt;</router-link>
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
const email = ref('');
const password = ref('');
const rePassword = ref('');

const signup = () => {

}
</script>

<style scoped>
.signup {
  max-width: 480px;
  margin: auto;
  padding: 0 20px;
}
</style>