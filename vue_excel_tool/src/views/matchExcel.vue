<template>
    <!--    <el-empty description="敬请期待"></el-empty>-->
    <el-card v-loading="loading"
             element-loading-text="数据拼命解析中,解析完成将自动下载"
             element-loading-spinner="el-icon-loading"
             element-loading-background="rgba(0, 0, 0, 0.8)">
        <div style="width: 400px">
            <el-upload class="upload-demo"
                       ref="upload"
                       action="string"
                       :on-change="onChange"
                       :on-remove="handleRemove"
                       :before-remove="beforeRemove"
                       multiple
                       :on-exceed="handleExceed"
                       accept=".xlsx"
                       :limit="2"
                       :file-list="fileList"
                       :auto-upload="false"
                       show-file-list>
                <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
                <div style="margin-top: 20px" slot="tip" class="el-upload__tip">点击按钮选择文件进行上传，大文件上传较慢请耐心等待</div>
            </el-upload>
            <div v-show="progressFlag" class="head-img" style="margin-top: 15px">
                <el-progress :text-inside="true" :stroke-width="18" :percentage="progressPercent"
                             status="success"></el-progress>
            </div>
        </div>
        <div style="margin-top: 20px">
            <el-radio-group v-model="model" size="small">
                <el-radio label="1" border>BD模式</el-radio>
                <el-radio label="2" border>作品模式</el-radio>
            </el-radio-group>
        </div>
        <el-button style="margin-top: 20px;" size="small" type="success" @click="submitUpload">上传服务器</el-button>
        <div><el-divider content-position="left"></el-divider></div>
        <div style="margin-top: 15px;font-size: 14px;border: 1px dashed #c4d9d0;border-radius: 6px;">
            <span class="el-icon-warning-outline" style="margin-left: 20px;margin-top: 15px;font-size: 17px;font-weight: bolder"> 使用须知</span>
            <div style="margin-left: 20px;margin-top: 10px">
                <div>
                    <span>1、将两张表格同时上传;</span>
                </div>
                <div style="margin-top: 10px">
                    <span style="color: #d96421;font-weight: bolder">2、关于表格的上传顺序,引用的(需要修改的)的表格放在第一位上传，被引用(作为参照的)的表格放在第二位上传;</span>
                </div>
                <div style="margin-top: 10px">
                    <span style="color: #d96421;font-weight: bolder">3、关于模式的选择,BD模式为使用渠道匹配大区,作品模式为使用用户ID匹配作品链接和作品得分;</span>
                </div>
                <div style="margin-top: 10px">
                    <span style="color: #d96421;font-weight: bolder">4、作品模式引用的表格为左侧菜单栏班期拆分功能导出的表；</span>
                </div>
                <div style="margin-bottom: 20px">
                    <span>  </span>
                </div>
            </div>
        </div>

    </el-card>
</template>

<script>

    export default {
        name: "matchExcel",
        data() {
            return {
                loading: false,
                model: '1',
                fileList: [],//储存多文件
                progressFlag: false,//进度条初始值隐藏
                progressPercent: 0,//进度条初始值
            };
        },
        methods: {
            handleRemove(file, fileList) {
                console.log(file, fileList);
            },
            handleExceed(files, fileList) {
                this.$message.warning(`当前限制选择 2 个文件进行上传`);
            },
            beforeRemove(file, fileList) {
                return this.$confirm(`确定移除 ${file.name}？`);
            },
            onChange(file, fileList) {
                this.fileList = fileList;
            },
            //上传函数
            submitUpload(file) {
                this.loading = true;
                //重新命名 方便setTimeout函数 --因为setTimeout函数在vue内部中无效
                const that = this;
                that.$refs.upload.submit();
                //判断上传文件数量
                this.filenameList = [];
                this.fileList.forEach(item => {
                    this.filenameList.push(item.name)
                });
                if (this.fileList.length === 0) {
                    that.$message({
                        message: '请选择导入的文件',
                        type: 'warning',
                        duration: '2000'
                    });
                } else {
                    //创建FormData();主要用于发送表单数据
                    let paramFormData = new FormData();
                    console.log(this.model);
                    //遍历 fileList
                    that.fileList.forEach(file => {
                        paramFormData.append("files", file.raw);
                        paramFormData.append("fileNames", file.name);
                    });
                    paramFormData.append("model", this.model);
                    //修改progressFlag值
                    that.progressFlag = true;
                    //控制台打印文件信息
                    console.log(paramFormData.getAll("files"));
                    //axios 发出请求
                    that.axios({
                        url: "/excel/multi/upload",
                        method: 'post',
                        data: paramFormData,
                        headers: {'Content-Type': 'multipart/form-data'},
                        responseType: 'arraybuffer',
                        onUploadProgress: progressEvent => {
                            // progressEvent.loaded:已上传文件大小
                            // progressEvent.total:被上传文件的总大小
                            //进度条
                            that.progressPercent = ((progressEvent.loaded / progressEvent.total) * 100) | 0;
                        }
                    }).then(res => {
                        this.loading = false;
                        if (that.progressPercent === 100) {
                            that.$message({
                                message: '导入成功！',
                                type: 'success',
                                duration: '2000'
                            });
                            that.progressFlag = false;
                            that.progressPercent = 0;
                            that.$refs.upload.clearFiles();
                            const url = window.URL.createObjectURL(new Blob([res.data], {
                                type: "application/vnd.ms-excel"
                            }));
                            //获取heads中的filename文件名
                            let temp = res.headers["content-disposition"].split(";")[1].split("filename=")[1];
                            const iconv = require('iconv-lite');
                            let fileName = iconv.decode(temp, 'utf8');
                            const a = document.createElement('a');
                            a.style.display = 'none';
                            a.href = url;
                            a.download = fileName;
                            document.body.appendChild(a);
                            a.click();
                            document.body.removeChild(a)
                        }
                        this.filenameList = [];
                    })
                }
            },
        }
    }
</script>

<style scoped>

</style>
