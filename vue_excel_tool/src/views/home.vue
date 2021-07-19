<template>
    <div>
        <el-card>
            <div>
                <div style="display: flex;justify-content: space-between">
                    <div>
                        <el-upload
                                :show-file-list="false"
                                :before-upload="beforeUpload"
                                :on-success="onSuccess"
                                :on-error="onError"
                                :disabled="importDataDisabled"
                                style="display: inline-flex;margin-right: 8px"
                                action="/excel/import">
                            <el-button :disabled="importDataDisabled" type="success" :icon="importDataBtnIcon">
                                {{importDataBtnText}}
                            </el-button>
                        </el-upload>
                        <el-button type="success" @click="exportData" icon="el-icon-download">
                            导出数据
                        </el-button>
                    </div>
                </div>
            </div>
            <div style="margin-top: 10px">

                <el-select
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
            </div>
            <div style="margin-top: 10px">
                <!--<el-tabs type="border-card">
                    <el-tab-pane :key="item.name"
                                 v-for="item in sheetNameList"
                                 :label="item">-->
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
                    >
                        <template slot-scope="scope">
                            {{scope.row[index]}}
                        </template>
                    </u-table-column>
                </u-table>

                <!--</el-tab-pane>
            </el-tabs>-->
            </div>
            <div style="margin-top: 20px">
                <el-form label-width="80px" ref="exportForm">
                    <el-form-item label="群成员">
                        <el-input style="width: 600px"
                                  type="textarea"
                                  placeholder="请将群成员粘贴到此处"
                                  v-model="textarea">
                        </el-input>
                    </el-form-item>
                    <el-form-item label="匹配的列">
                        <el-select v-model="columnValue" placeholder="请选择">
                            <el-option
                                    v-for="(item,index) in tableHeader"
                                    :key="index"
                                    :value="index"
                                    :label="item">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="导出的列">
                        <el-select v-model="exportColumnValue" placeholder="请选择">
                            <el-option
                                    v-for="(item,index) in tableHeader"
                                    :key="index"
                                    :value="index"
                                    :label="item">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </el-form>

                <el-button type="primary" icon="el-icon-s-check" @click="doExport">匹配未激活用户
                </el-button>
                <div>
                    <el-input v-show="this.resultString.length !== 0" style="margin-top: 20px;width: 30%"
                              type="textarea"
                              autosize
                              readonly
                              placeholder="匹配结果将在此处显示"
                              v-model="resultString">
                    </el-input>
                </div>
            </div>
        </el-card>

    </div>
</template>

<script>
    import {postRequest} from "../utils/api";

    export default {
        name: "Home",
        data() {
            return {
                resultString: "",
                textarea: "",
                height: 400,
                rowHeight: 55,
                radioValue: "",
                columnValue: "",
                exportColumnValue: "",
                importDataBtnText: '导入数据',
                importDataBtnIcon: 'el-icon-upload2',
                importDataDisabled: false,
                loading: false,
                total: 0,
                tableHeader: [],
                tableData: [],
                sheetNameList: [],
                sheetList: [],
                file: null,

            }
        },
        mounted() {
            this.initData();
        },
        methods: {
            doExport() {
                if (this.textarea.trim() === ""){
                    this.$message.error("群成员不能为空")
                }
                const formdata = new FormData();
                formdata.append("file", this.file);
                formdata.append("textarea", this.textarea);
                formdata.append("columnValue", this.columnValue);
                formdata.append("exportColumnValue", this.exportColumnValue);
                formdata.append("sheetIndex", this.radioValue);
                postRequest("/excel/inactive/user", formdata).then(res => {
                    console.log(res);
                    this.resultString = res.data
                })
            },

            changeValue(value) {
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
                // this.tableHeader = response.data.sheet.table_header
                // this.tableData = response.data.sheet.table_data
                // this.sheetNameList = response.data.sheet_name_list
                this.sheetList = response.data.sheet_list;
                this.file = file.raw;
                this.loading = false

            },
            beforeUpload() {
                this.importDataBtnText = '正在导入';
                this.importDataBtnIcon = 'el-icon-loading';
                this.importDataDisabled = true;
            },
            exportData() {
                window.open('/employee/basic/export', '_parent');
            },

            initData() {

            },
        },

        computed: {
            //计算属性

        },
    }
</script>

<style>
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