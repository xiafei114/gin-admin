<template>
  <div class="page-header-index-wide">
    <a-card :bordered="false" :bodyStyle="{ padding: '16px 0', height: '100%' }" :style="{ height: '100%' }">
      <div class="card-list" ref="content">
        <a-list
          :grid="{gutter: 24, lg: 3, md: 2, sm: 1, xs: 1}"
          :dataSource="dataSource"
        >
          <a-list-item slot="renderItem" slot-scope="item">
            <template v-if="item === null">
              <a-upload name="fileData" :multiple="false" :showUploadList="false" action="http://localhost:8088/api/v1/medias/upload" :headers="headers" @change="handleChange">
                <a-button>
                  <a-icon type="upload" /> Click to Upload
                </a-button>
              </a-upload>
            </template>
            <template v-else>
              <a-card :hoverable="true">
                <img slot="cover" src="../../../assets/images/mifeng.png" :alt="item.title" />
                <template class="ant-card-actions" slot="actions">
                  <a>删除</a>
                </template>
              </a-card>
            </template>
          </a-list-item>
        </a-list>
      </div>
    </a-card>
  </div>
</template>

<script>
import { mixinDevice } from '@/utils/mixin.js'
import { STable } from '@/components'

const dataSource = []
dataSource.push(null)
for (let i = 0; i < 11; i++) {
  dataSource.push({
    title: 'Alipay',
    cover: 'https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png',
    content: '在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。'
  })
}

export default {
  components: {
    STable
  },
  mixins: [mixinDevice],
  data () {
    return {
      menuKey: new Date().getTime(),
      description: '段落示意：蚂蚁金服务设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态， 提供跨越设计与开发的体验解决方案。',
      linkList: [
        { icon: 'rocket', href: '#', title: '快速开始' },
        { icon: 'info-circle-o', href: '#', title: '产品简介' },
        { icon: 'file-text', href: '#', title: '产品文档' }
      ],
      dataSource,
      headers: {
        Authorization: 'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTkyNjc3MDIsImlhdCI6MTU1ODY2MjkwMiwibmJmIjoxNTU4NjYyOTAyLCJzdWIiOiI4Y2U0ZmQ2Ny1kMzJlLTRiZGItYjNkNS03ZmM1NjQxZTM0YjcifQ.-N4byXHzyn1bKRsX96IlNbiLiTvEPzuuctM6MIrm5DcLjwTGPBYtZaod7DpvI6EI4PiRExXh3sSWZv5uBu-utw'
      }
    }
  },
  created () {
  },
  methods: {
    handleChange (info) {
      if (info.file.status !== 'uploading') {
        console.log(info.file, info.fileList)
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} file uploaded successfully`)
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} file upload failed.`)
      }
    }
  },
  computed: {
    record_id: function () {
      return this.$route.query.record_id
    }
  }
}
</script>

<style lang="less" scoped>

  .card-list{
    flex: 1 1;
    padding: 8px 40px;

    .account-settings-info-title {
      color: rgba(0,0,0,.85);
      font-size: 20px;
      font-weight: 500;
      line-height: 28px;
      margin-bottom: 12px;
    }
    .account-settings-info-view {
      padding-top: 12px;
    }

    .card-avatar {
      width: 48px;
      height: 48px;
      border-radius: 48px;
    }

    .ant-card-actions {
      background: #f7f9fa;
      li {
        float: left;
        text-align: center;
        margin: 12px 0;
        color: rgba(0, 0, 0, 0.45);
        width: 50%;

        &:not(:last-child) {
          border-right: 1px solid #e8e8e8;
        }

        a {
          color: rgba(0, 0, 0, .45);
          line-height: 22px;
          display: inline-block;
          width: 100%;
          &:hover {
            color: #1890ff;
          }
        }
      }
    }

    .new-btn {
      background-color: #fff;
      border-radius: 2px;
      width: 100%;
      height: 180px;
    }

    .meta-content {
      position: relative;
      overflow: hidden;
      text-overflow: ellipsis;
      display: -webkit-box;
      height: 64px;
      -webkit-line-clamp: 3;
      -webkit-box-orient: vertical;
    }
  }

</style>
