<template>
    <div>
        <el-card style="height: 100%">
            <div>
                <div style="display: flex;justify-content: space-between">
                    <div>
                        <!--<el-upload
                                :show-file-list="true"
                                :before-upload="beforeUpload"
                                :on-success="onSuccess"
                                :on-error="onError"
                                :disabled="importDataDisabled"
                                :limit="1"
                                style="display: inline-flex;margin-right: 8px"
                                action="/excel/import">
                            <el-button :disabled="importDataDisabled" type="success" :icon="importDataBtnIcon">
                                {{importDataBtnText}}
                            </el-button>
                        </el-upload>-->
                        <uploader
                                :options="options"
                                :file-status-text="statusText"
                                :auto-start="false"
                                class="uploader-example"
                                @file-added="fileAdded"
                                @file-progress="onFileProgress"
                                @file-success="onFileSuccess"
                                @file-error="onFileError"
                        >
                            <uploader-unsupport></uploader-unsupport>
                            <uploader-drop>
                                <p>将文件拖放到此处进行上传或</p>
                                <uploader-btn style=" display: inline-block;line-height: 1;white-space: nowrap;cursor: pointer;text-align: center;
                                    box-sizing: border-box; margin: 0;transition: .1s;font-weight: 500; padding: 12px 20px;
                                    font-size: 12px; border-radius: 4px;    color: #FFF;background-color: #409EFF;border-color: #409EFF;">
                                    选择文件
                                </uploader-btn>
                            </uploader-drop>
                            <uploader-list/>
                        </uploader>
                    </div>
                    <div v-show="false">
                        <el-button type="success" @click="exportData" icon="el-icon-download">
                            导出数据
                        </el-button>
                    </div>
                </div>
            </div>
            <!--<div style="margin-top: 10px">

                <el-select v-show="resultString.trim().length !== 0"
                           v-model="radioValue"
                           @change="changeValue"
                           placeholder="请选择sheet表">
                    <el-option
                            v-for="(item,idx) in sheetList"
                            :key="item.sheet_index"
                            :label="item.sheet_name"
                            :value="item.sheet_index">
                    </el-option>
                </el-select>
            </div>-->
            <div style="margin-top: 10px">
                <!--<el-tabs type="border-card"
                         :value="getActiveName(sheetList)"
                         @tab-click="handleClick"
                         v-show="showForm">
                    <el-tab-pane style="width: 100%;height: 300px"
                                 :key="item.sheet_index"
                                 v-for="(item,idx) in sheetList"
                                 :name="item.sheet_index.toString()"
                                 :label="item.sheet_name">
                        <u-table
                                :data="tableData"
                                use-virtual
                                :row-height="rowHeight"
                                :height="height"
                                border
                                v-loading="loading"
                                element-loading-text="正在加载..."
                                element-loading-spinner="el-icon-loading"
                                element-loading-background="rgba(0, 0, 0, 0.8)"
                                style="width: 100%">
                            <u-table-column
                                    v-for="(item,index) in tableHeader"
                                    :key="index"
                                    align="center"
                                    :label="item"
                                    :resizable="item.resizable"
                            >-->
                                <!--<template slot-scope="scope">
                                    {{scope.row[index]}}
                                </template>
                            </u-table-column>
                        </u-table>

                    </el-tab-pane>
                </el-tabs>-->
            </div>
            <div style="margin-top: 20px">
                <el-form label-width="80px" ref="exportForm" v-show="showForm">
                    <el-form-item label="群成员">
                        <el-input style="width: 600px"
                                  type="textarea"
                                  placeholder="请将群成员粘贴到此处"
                                  v-model="textarea">
                        </el-input>
                    </el-form-item>
                    <el-form-item label="匹配的列">
                        <el-select v-model="columnValue" disabled placeholder="请选择">
                            <el-option
                                    v-for="(item,index) in tableHeader"
                                    :key="index"
                                    :value="index"
                                    :label="item">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="导出的列">
                        <el-select v-model="exportColumnValue" disabled placeholder="请选择">
                            <el-option
                                    v-for="(item,index) in tableHeader"
                                    :key="index"
                                    :value="index"
                                    :label="item">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </el-form>

                <el-button v-show="showForm" type="primary" icon="el-icon-s-check" @click="doExport">匹配未激活用户
                </el-button>
                <div v-show="this.resultString.length !== 0">
                    <el-input style="margin-top: 20px;width: 40%"
                              type="textarea"
                              autosize
                              readonly
                              placeholder="匹配结果将在此处显示"
                              v-model="resultString">
                    </el-input>
                    <span style="font-size: 12px;margin-left: 10px">该群共 {{ count }} 人未激活</span>
                </div>
                <div style="margin-top: 20px" v-show="this.resultString.length !== 0">
                    <el-button
                            size="mini"
                            type="primary"
                            v-clipboard:copy="resultString"
                            v-clipboard:success="copySuccess"
                            v-clipboard:error="copyError">复制
                    </el-button>
                </div>
            </div>
        </el-card>

    </div>
</template>

<script>
    import {getRequest, postRequest} from "../utils/api";
    import SparkMD5 from 'spark-md5'

    export default {
        name: "matchInactive",
        data() {
            return {
                isUploaded: false,
                md5: '',
                resultString: "",
                textarea: "",
                height: 300,
                rowHeight: 55,
                columnValue: "用户状态",
                exportColumnValue: "微信昵称",
                importDataBtnText: '导入数据',
                importDataBtnIcon: 'el-icon-upload2',
                importDataDisabled: false,
                loading: false,
                count: 0,
                tableHeader: [],
                tableData: [],
                sheetNameList: [],
                notUploadedChunks: [],
                sheetList: [{
                    sheet_index: 0,
                    sheet_name: ""
                }],
                file: null,
                showForm: false,
                sheetIndex: 0
            }
        },
        methods: {
            // 上传单个文件
            fileAdded(file) {
                this.computeMD5(file) // 生成MD5
            },
            // 计算MD5值
            computeMD5(file) {
                const that = this;
                this.isUploaded = false; // 这个文件是否已经上传成功过
                this.notUploadedChunks = []; // 未成功的chunkNumber
                const fileReader = new FileReader();
                let md5 = '';
                file.pause();
                fileReader.readAsArrayBuffer(file.file);
                fileReader.onload = async function (e) {
                    if (file.size !== e.target.result.byteLength) {
                        this.$message.error(
                            'Browser reported success but could not read the file until the end.'
                        );
                        return false
                    }
                    md5 = SparkMD5.ArrayBuffer.hash(e.target.result, false);
                    file.uniqueIdentifier = md5;
                    if (md5 !== '') {
                        const res = await this.getRequest("/excel/simple/check", {md5: md5});
                        console.log(res);
                        if (res.code === 0) {
                            if (res.data.isDone) {
                                // 上传成功过
                                this.isUploaded = true;
                                that.$message({
                                    message: '该文件已经上传成功过了，秒传成功。',
                                    type: 'success'
                                });
                                file.cancel()
                            } else {
                                this.isUploaded = false;
                                this.notUploadedChunks = res.data.chunks;
                                if (this.notUploadedChunks.length) {
                                    file.resume()
                                }
                            }
                        }
                    }
                };
                fileReader.onerror = function () {
                    this.$.message.error(
                        'generator md5 时FileReader异步读取文件出错了，FileReader onerror was triggered, maybe the browser aborted due to high memory usage.'
                    );
                    return false
                }
            },
            // 上传进度
            onFileProgress() {
            },
            // 上传成功
            async onFileSuccess(rootFile, file) {
                const response = await this.getRequest("/excel/simple/merge", {
                    md5: file.uniqueIdentifier,
                    fileName: file.name
                });
                console.log(response);
                this.tableHeader = response.data.sheet.table_header;
                // this.tableData = response.data.sheet.table_data;
                this.sheetNameList = response.data.sheet_name_list;
                this.sheetList = response.data.sheet_list;
                this.file = file.file;
                this.loading = false;
                this.showForm = true;
            },
            onFileError(rootFile, file, response) {
                this.$message({
                    message: response,
                    type: 'error'
                })
            },

            copySuccess(e) {
                this.$message({
                    type: "success",
                    message: "复制成功",
                    duration: 1500,
                    showClose: true,
                });
            },
            copyError(e) {
                this.$message({
                    message: 'Copy error',
                    type: 'error',
                    duration: 1500
                });
            },
            async doExport() {
                if (this.textarea.trim() === "") {
                    this.$message.error("群成员不能为空");
                    return false
                }
                const formdata = new FormData();
                formdata.append("filename", this.file.name);
                formdata.append("textarea", this.textarea);
                formdata.append("columnValue", this.getColumnIndex(this.columnValue));
                formdata.append("exportColumnValue", this.getColumnIndex(this.exportColumnValue));
                formdata.append("sheetIndex", this.sheetIndex);
                this.postRequest("/excel/inactive/user", formdata).then(res=>{
                    if (res.code === 1000) {
                        console.log(res);
                        if (res.data.result.trim().length === 0) {
                            this.$message.success("该群全部成员都已激活！")
                        } else {
                            this.$message.success("@文本生成成功");
                            this.resultString = res.data.result;
                            this.count = res.data.count;
                        }
                    }
                })

            },

            getColumnIndex(columnName) {
                for (let i = 0; i < this.tableHeader.length; i++) {
                    if (this.tableHeader[i] === columnName) {
                        return i
                    }
                }
            }
            ,
            getActiveName(sheetList) {
                // 默认选择第一项
                return sheetList[0].sheet_index.toString();
            }
            ,
            handleClick(tab, event) {
                this.changeValue(tab.name)
            }
            ,
            changeValue(value) {
                this.sheetIndex = value;
                const formdata = new FormData();
                formdata.append("file", this.file);
                let params = {
                    index: value,
                };
                formdata.append("index", params.index);
                postRequest("/excel/table", formdata)
                    .then(res => {
                        console.log(res);
                        this.tableHeader = res.data.sheet.table_header;
                        this.tableData = res.data.sheet.table_data;
                    })
            },
            onError(err, file, fileList) {
                this.importDataBtnText = '导入数据';
                this.importDataBtnIcon = 'el-icon-upload2';
                this.importDataDisabled = false;
            },
            onSuccess(response, file, fileList) {
                this.importDataBtnText = '导入数据';
                this.importDataBtnIcon = 'el-icon-upload2';
                this.importDataDisabled = false;
                console.log(response);
                this.tableHeader = response.data.sheet.table_header;
                this.tableData = response.data.sheet.table_data;
                this.sheetNameList = response.data.sheet_name_list;
                this.sheetList = response.data.sheet_list;
                this.file = file.raw;
                this.loading = false;
                this.showForm = true;
            },
            beforeUpload() {
                this.importDataBtnText = '正在导入';
                this.importDataBtnIcon = 'el-icon-loading';
                this.importDataDisabled = true;
            },
            exportData() {
                window.open('/employee/basic/export', '_parent');
            },
        },

        computed: {
            statusText() {
                return {
                    success: '成功',
                    error: '错误',
                    uploading: '上传中',
                    paused: '暂停中',
                    waiting: '等待中'
                }
            }
            ,
            options() {
                return {
                    target: '/excel/simple/upload',
                    testChunks: false,
                    simultaneousUploads: 1,
                    chunkSize: 2 * 1024 * 1024,
                    checkChunkUploadedByResponse(chunk) {
                        if (this.isUploaded) {
                            return true // return true 会忽略当前文件，不会再发送给后台
                        } else {
                            // 根据已经上传过的切片来进行忽略
                            return (
                                this.notUploadedChunks &&
                                this.notUploadedChunks.some(
                                    item => item.chunkNumber === chunk.offset + 1
                                )
                            )
                        }
                    }
                }
            }
        },
    }
</script>

<style>
    .uploader-example {
        width: 880px;
        padding: 15px;
        margin: 15px 15px 20px;
        font-size: 12px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
    }

    .uploader-example .uploader-btn {
        margin-right: 4px;
    }

    .uploader-example .uploader-list {
        margin-top: 15px;
        max-height: 440px;
        overflow: auto;
        overflow-x: hidden;
    }

    /* 可以设置不同的进入和离开动画 */
    /* 设置持续时间和动画函数 */
    .slide-fade-enter-active {
        transition: all .8s ease;
    }

    .slide-fade-leave-active {
        transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
    }

    .slide-fade-enter, .slide-fade-leave-to
        /* .slide-fade-leave-active for below version 2.1.8 */
    {
        transform: translateX(10px);
        opacity: 0;
    }

    /*.el-table--scrollable-x .el-table__body-wrapper {*/
    /*    overflow: scroll !important;*/
    /*    height: 29rem !important;*/
    /*}*/
</style>
