<template>
    <div id="app">
        <vue-particles
                class="login-background"
                color="#97D0F2"
                :particleOpacity="0.7"
                :particlesNumber="50"
                shapeType="circle"
                :particleSize="4"
                linesColor="#97D0F2"
                :linesWidth="1"
                :lineLinked="true"
                :lineOpacity="0.4"
                :linesDistance="150"
                :moveSpeed="3"
                :hoverEffect="true"
                hoverMode="grab"
                :clickEffect="true"
                clickMode="push">
        </vue-particles>
        <div class="login-form">
            <el-form :rules="rules" ref="loginForm" :model="loginForm" class="loginContainer">
                <h3 class="loginTitle">系统登录</h3>
                <el-form-item prop="username">
                    <el-input type="text" v-model="loginForm.username" auto-complete="off"
                              placeholder="请输入用户名"></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input type="password" v-model="loginForm.password" auto-complete="off"
                              placeholder="请输入用户密码"></el-input>
                </el-form-item>
                <el-checkbox v-model="checked">记住密码</el-checkbox>
                <el-button type="primary" style="width: 100%;margin-top: 15px" @click="submitLogin">登录</el-button>
            </el-form>
        </div>
    </div>
</template>

<script>
    import {generaMenu} from "../utils/menu";

    export default {
        name: "login",
        data() {
            return {
                loginForm: {
                    username: 'admin',
                    password: 'admin123456'
                },

                checked: true,

                rules: {
                    username:[{required:true,message:'请输入用户名',trigger:'blur'}],
                    password:[{required:true,message:'请输入密码',trigger:'blur'}]

                }
            }
        },

        methods: {
            submitLogin() {
                this.$refs.loginForm.validate((valid) => {
                    if (valid) {
                        this.$store.commit("login", this.loginForm);
                        generaMenu();
                        this.$message.success("登录成功");
                        this.$router.push({path:"/index"})
                    } else {
                        this.$message.error('请输入必输字段');
                        return false;
                    }
                });
            }
        }
    }
</script>

<style>
    html,body{
        margin:0;
        padding:0;
    }
    .login-background {
        background: linear-gradient(-180deg, #dcf2e6 0%, #304156 100%);
        /*background-color: #304156;*/
        width: 100%;
        height: 100%; /**宽高100%是为了图片铺满屏幕 */
        z-index: -1;
        position: absolute;
    }
    .login-form{
        z-index: 1;
        margin: 180px 0 0 calc(calc(100vw - 410px) / 2);
        position: absolute;
    }

    .loginContainer {
        border-radius: 15px;
        background-clip: padding-box;
        margin: 180px auto;
        width: 350px;
        padding: 15px 35px 15px 35px;
        background: white;
        border: 1px solid #eaeaea;
        box-shadow: 0 0 25px #cac6c6;
    }

    .loginTitle {
        margin: 15px auto 20px auto;
        text-align: center;
        font-family: "萍方0",sans-serif;
    }

</style>
