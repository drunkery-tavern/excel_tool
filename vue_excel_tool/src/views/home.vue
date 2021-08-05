<template>
    <div>
        <el-carousel :interval="4000" type="card" height="260px">
            <el-carousel-item datasrc="" v-for="item in images" :key="item">
                <img :src="item" width="100%" height="100%">
            </el-carousel-item>
        </el-carousel>
        <el-card>
            <el-row>
                <el-col :span="6">
                    <el-image style="margin-top: 20px" src="http://qiniu.drunkery.cn/logo.png"></el-image>
                </el-col>
                <el-col :span="18">
                    <el-row align="center">
                        <el-tag effect="dark" type="danger">excel</el-tag>
                        <el-tag style="margin-left: 20px" effect="dark" type="success">工具</el-tag>
                        <el-tag style="margin-left: 20px" effect="dark">开源</el-tag>
                    </el-row>
                    <el-row style="margin-top: 30px">
                        <div style="height: 160px;">
                            <el-steps direction="vertical">
                                <el-step icon="el-icon-s-finance" status="finish" title="匹配用户"
                                         description="文件采用分片上传的方式，可一键匹配企业微信群中未激活课程的用户，该工具需配合KeyboardMan一起使用"></el-step>
                                <el-step icon="el-icon-s-order" status="finish" title="表格合并"
                                         description="Excel的VLOOKUP函数功能自动化实现"></el-step>
                                <el-step icon="el-icon-s-management" status="finish" title="班期拆分"
                                         description="将班期总表激活的用户按班期拆分成若干sheet"></el-step>
                            </el-steps>
                        </div>
                    </el-row>
                </el-col>
            </el-row>
        </el-card>
            <el-row>
                <el-col :span="8">
                    <el-card style="margin-top: 15px;height: 300px">
                        <div style="width: 500px;height: 230px;display: grid;margin-top: 15px;margin-left: 15px">
                            <span style="font-weight: bold;font-size: 20px;text-align: center">技术栈</span>
                            <div style="padding-top:5px;" class="progress-item">
                                <span>Vue</span>
                                <el-progress :text-inside="true" :stroke-width="18" :percentage="65"/>
                            </div>
                            <div class="progress-item">
                                <span>JavaScript</span>
                                <el-progress :text-inside="true" :stroke-width="18" :percentage="25" status="exception"/>
                            </div>
                            <div class="progress-item">
                                <span>Go</span>
                                <el-progress :text-inside="true" :stroke-width="18" :percentage="90" status="success"/>
                            </div>
                            <div class="progress-item">
                                <span>Redis</span>
                                <el-progress :text-inside="true" :stroke-width="18" :percentage="10" status="warning"/>
                            </div>
                        </div>
                    </el-card>
                </el-col>
                <el-col :span="16">
                    <el-card style="margin-top: 15px;margin-left: 15px;height: 300px">
                        <!-- 表格展示 -->
                        <el-table
                                border
                                :data="fileList"
                                v-loading="loading"
                                height="225px"
                                stripe
                                size="mini"
                        >
                            <!-- 表格列 -->
                            <el-table-column prop="filename" label="文件名" align="center"/>
                            <el-table-column prop="file_size" label="大小" align="center" width="150">
                                <template slot-scope="scope">
                                    <el-tag v-if="(scope.row.file_size / 1024 / 1024) < 1" effect="dark">
                                        {{ scope.row.file_size / 1024 | numFilter }} kb
                                    </el-tag>
                                    <el-tag v-else effect="dark" type="danger">
                                        {{ scope.row.file_size / 1024 /1024 | numFilter }} Mb
                                    </el-tag>
                                </template>
                            </el-table-column>
                            <el-table-column
                                    prop="last_update_time"
                                    label="最后修改时间"
                                    align="center">
                                <template slot-scope="scope">
                                    <i class="el-icon-time" style="margin-right:5px"/>
                                    {{ scope.row.last_update_time}}
                                </template>
                            </el-table-column>
                        </el-table>
                        <h3 style="margin-left:20px;float: left;">已上传的文件列表<span style="font-size: 14px;margin-left: 20px;color: #d96421"> 多人同时操作时，请勿使用同名文件</span></h3>
                        <!-- 分页 -->
                        <el-pagination
                                class="pagination-container"
                                background
                                @size-change="sizeChange"
                                @current-change="currentChange"
                                :current-page="current"
                                :page-size="size"
                                :total="count"
                                :page-sizes="[5, 10]"
                                layout="total, sizes, prev, pager, next, jumper"
                        />
                    </el-card>
                </el-col>
            </el-row>
    </div>
</template>

<script>
    export default {
        name: "home",
        created() {
            this.listFiles();
        },
        data() {
            return {
                fileList: [],
                loading: true,
                current: 1,
                size: 10,
                count: 0,
                images: [
                    "http://qiniu.drunkery.cn/home.jpg",
                    "http://qiniu.drunkery.cn/home1.jpg",
                    "http://qiniu.drunkery.cn/wallhaven-6kdgpq.jpg",
                    "http://qiniu.drunkery.cn/wallhaven-rddgwm.jpg",
                ]
            }
        },
        methods: {
            sizeChange(size) {
                this.size = size;
                this.listFiles();
            },
            currentChange(current) {
                this.current = current;
                this.listFiles();
            },
            async listFiles() {
                const res = await this.getRequest("/excel/system/files", {
                    current: this.current,
                    size: this.size,
                });
                console.log(res.data);
                this.loading = false;
                this.fileList = res.data.data.fileList;
                this.count = res.data.data.count
            }
        },
        filters: {
            numFilter(value) {
                // 截取当前数据到小数点后两位
                return parseFloat(value).toFixed(2)
            }
        }
    }
</script>

<style>
    .progress-item {
        margin-bottom: 15px;
        font-size: 14px;
    }

    .el-carousel__item h3 {
        color: #475669;
        font-size: 14px;
        opacity: 0.75;
        line-height: 200px;
        margin: 0;
    }

    .pagination-container {
        float: right;
        margin-top: 1.25rem;
        margin-bottom: 1.25rem;
    }
</style>
