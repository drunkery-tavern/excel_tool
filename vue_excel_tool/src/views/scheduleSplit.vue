<template>
    <!--    <el-empty description="敬请期待"></el-empty>-->
    <el-card>
        <div style="width: 400px">
            <el-upload
                    class="upload-demo"
                    drag
                    :httpRequest="onUpload" action="string">
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                <div class="el-upload__tip" slot="tip">大文件上传较慢请耐心等待...</div>
                <div class="el-upload__tip" slot="tip">表格数据越多，解析越慢，解析完成将自动下载</div>
            </el-upload>
        </div>
    </el-card>
</template>

<script>
    export default {
        name: "scheduleSplit",
        data() {
            return {
                downloadDisabled: false,
            }
        },
        methods: {
            onUpload(config) {
                console.log(config);
                //action="/excel/schedule/upload"
                let paramFormData = new FormData();
                paramFormData.append("file", config.file);
                //axios 发出请求
                this.axios({
                    url: "/excel/schedule/upload",
                    method: 'post',
                    data: paramFormData,
                    headers: {'Content-Type': 'multipart/form-data'},
                    responseType: 'arraybuffer',
                }).then(res => {
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
                })
            },
        }
    }
</script>

<style>
    .el-upload-dragger {
        background-color: #fff;
        border: 1px dashed #d96421;
        border-radius: 6px;
        box-sizing: border-box;
        width: 360px;
        height: 150px;
        text-align: center;
        position: relative;
        overflow: hidden;
    }

    .el-upload--picture-card, .el-upload-dragger {
        -webkit-box-sizing: border-box;
        cursor: pointer;
    }

</style>
