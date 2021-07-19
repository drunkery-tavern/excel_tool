<template>
    <div>
        <el-card>
            <div>
                <div style="display: flex;justify-content: space-between">
                    <!--<div>
                        <el-input placeholder="请输入员工名进行搜索，可以直接回车搜索..." prefix-icon="el-icon-search"
                                  clearable
                                  @clear="initEmps"
                                  style="width: 350px;margin-right: 10px" v-model="keyword"
                                  @keydown.enter.native="initEmps" :disabled="showAdvanceSearchView"></el-input>
                        <el-button icon="el-icon-search" type="primary" @click="initEmps" :disabled="showAdvanceSearchView">
                            搜索
                        </el-button>
                        <el-button type="primary" @click="showAdvanceSearchView = !showAdvanceSearchView">
                            <i :class="showAdvanceSearchView?'fa fa-angle-double-up':'fa fa-angle-double-down'"
                               aria-hidden="true"></i>
                            高级搜索
                        </el-button>
                    </div>-->
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
                <!--<transition name="slide-fade">
                    <div v-show="showAdvanceSearchView"
                         style="border: 1px solid #409eff;border-radius: 5px;box-sizing: border-box;padding: 5px;margin: 10px 0px;">
                        <el-row>
                            <el-col :span="5">
                                政治面貌:
                                <el-select v-model="searchValue.politicId" placeholder="政治面貌" size="mini"
                                           style="width: 130px;">
                                    <el-option
                                            v-for="item in politicsstatus"
                                            :key="item.id"
                                            :label="item.name"
                                            :value="item.id">
                                    </el-option>
                                </el-select>
                            </el-col>
                            <el-col :span="4">
                                民族:
                                <el-select v-model="searchValue.nationId" placeholder="民族" size="mini"
                                           style="width: 130px;">
                                    <el-option
                                            v-for="item in nations"
                                            :key="item.id"
                                            :label="item.name"
                                            :value="item.id">
                                    </el-option>
                                </el-select>
                            </el-col>
                            <el-col :span="4">
                                职位:
                                <el-select v-model="searchValue.posId" placeholder="职位" size="mini" style="width: 130px;">
                                    <el-option
                                            v-for="item in positions"
                                            :key="item.id"
                                            :label="item.name"
                                            :value="item.id">
                                    </el-option>
                                </el-select>
                            </el-col>
                            <el-col :span="4">
                                职称:
                                <el-select v-model="searchValue.jobLevelId" placeholder="职称" size="mini"
                                           style="width: 130px;">
                                    <el-option
                                            v-for="item in joblevels"
                                            :key="item.id"
                                            :label="item.name"
                                            :value="item.id">
                                    </el-option>
                                </el-select>
                            </el-col>
                            <el-col :span="7">
                                聘用形式:
                                <el-radio-group v-model="searchValue.engageForm">
                                    <el-radio label="劳动合同">劳动合同</el-radio>
                                    <el-radio label="劳务合同">劳务合同</el-radio>
                                </el-radio-group>
                            </el-col>
                        </el-row>
                        <el-row style="margin-top: 10px">
                            <el-col :span="5">
                                所属部门:
                                <el-popover
                                        placement="right"
                                        title="请选择部门"
                                        width="200"
                                        trigger="manual"
                                        v-model="popVisible2">
                                    <el-tree default-expand-all :data="allDeps" :props="defaultProps"
                                             @node-click="searvhViewHandleNodeClick"></el-tree>
                                    <div slot="reference"
                                         style="width: 130px;display: inline-flex;font-size: 13px;border: 1px solid #dedede;height: 26px;border-radius: 5px;cursor: pointer;align-items: center;padding-left: 8px;box-sizing: border-box;margin-left: 3px"
                                         @click="showDepView2">{{inputDepName}}
                                    </div>
                                </el-popover>
                            </el-col>
                            <el-col :span="10">
                                入职日期:
                                <el-date-picker
                                        v-model="searchValue.beginDateScope"
                                        type="daterange"
                                        size="mini"
                                        unlink-panels
                                        value-format="yyyy-MM-dd"
                                        range-separator="至"
                                        start-placeholder="开始日期"
                                        end-placeholder="结束日期">
                                </el-date-picker>
                            </el-col>
                            <el-col :span="5" :offset="4">
                                <el-button size="mini">取消</el-button>
                                <el-button size="mini" icon="el-icon-search" type="primary" @click="initEmps('advanced')">
                                    搜索
                                </el-button>
                            </el-col>
                        </el-row>
                    </div>
                </transition>-->
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
                height: 450,
                rowHeight: 55,
                radioValue: "",
                columnValue: "",
                exportColumnValue: "",
                searchValue: {
                    politicId: null,
                    nationId: null,
                    jobLevelId: null,
                    posId: null,
                    engageForm: null,
                    departmentId: null,
                    beginDateScope: null
                },
                title: '',
                importDataBtnText: '导入数据',
                importDataBtnIcon: 'el-icon-upload2',
                importDataDisabled: false,
                showAdvanceSearchView: false,
                allDeps: [],
                emps: [],
                loading: false,
                popVisible: false,
                popVisible2: false,
                dialogVisible: false,
                total: 0,
                page: 1,
                keyword: '',
                size: 10,

                defaultProps: {
                    children: 'children',
                    label: 'name'
                },
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
            searvhViewHandleNodeClick(data) {
                this.inputDepName = data.name;
                this.searchValue.departmentId = data.id;
                this.popVisible2 = !this.popVisible2
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