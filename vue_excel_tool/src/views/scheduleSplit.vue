<template>
    <!--    <el-empty description="敬请期待"></el-empty>-->
    <el-card>
        <div style="width: 400px">
            <el-upload
                    class="upload-demo"
                    drag
                    action="/excel/schedule/upload"
                    :on-success="onSuccess"
            >
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
        methods: {
            onSuccess(response, file, fileList) {
                this.importDataBtnText = '导入数据';
                this.importDataBtnIcon = 'el-icon-upload2';
                this.importDataDisabled = false;
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
