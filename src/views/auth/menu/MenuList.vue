<template>
  <div class="page-header-index-wide">
    <a-modal
      title="详情"
      :visible="visible"
      @ok="handleSubmit"
      :confirmLoading="confirmLoading"
      @cancel="handleCancel"
    >
      <a-form :form="form">
        <a-form-item label="唯一识别码">
          <a-input
            placeholder="前端路由菜单识别码"
            v-decorator="['role', {
              rules: [{ required: true, message: '请输入唯一识别码!' }]
            }]"
          />
        </a-form-item>

        <a-form-item label="权限名称">
          <a-input
            placeholder="权限名称"
            v-decorator="['title', {
              rules: [{ required: true, message: '请输入权限名称!' }]
            }]"
          />
        </a-form-item>

        <a-form-item label="权限规则">
          <a-input
            placeholder="模块/控制器/方法"
            v-decorator="['name', {
              rules: [{ required: true, message: '请输入权限规则!' }]
            }]"
          />
        </a-form-item>

        <a-form-item label="所属组别">
          <a-select
            v-decorator="['pid', {
              rules: [{ required: true, message: '请选择所属组别!' }]
            }]"
            placeholder="请选择所属组别"
          >
            <a-select-option :value="0">
              顶级分类
            </a-select-option>

            <a-select-option v-for="(item, index) in tree" :value="item.id" :key="index">
              {{ item.cname | html }}
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item label="状态">
          <a-select
            v-decorator="['status', {
              rules: [{ required: true, message: '请选择状态!' }]
            }]"
            placeholder="请选择"
          >
            <a-select-option :value="1">
              正常
            </a-select-option>
            <a-select-option :value="0">
              禁用
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item label="菜单动作管理">
          <a-button style="" type="dashed" icon="plus" @click="newMember">新增成员</a-button>
          <a-table
            :columns="columns"
            :dataSource="data"
            :pagination="false"
            :loading="memberLoading"
          >
            <template v-for="(col, i) in ['name', 'workId', 'department']" :slot="col" slot-scope="text, record">
              <a-input
                :key="col"
                v-if="record.editable"
                style="margin: -5px 0"
                :value="text"
                :placeholder="columns[i].title"
                @change="e => handleChange(e.target.value, record.key, col)"
              />
              <template v-else>{{ text }}</template>
            </template>
            <template slot="operation" slot-scope="text, record">
              <template v-if="record.editable">
            <span v-if="record.isNew">
              <a @click="saveRow(record)">添加</a>
              <a-divider type="vertical" />
              <a-popconfirm title="是否要删除此行？" @confirm="remove(record.key)">
                <a>删除</a>
              </a-popconfirm>
            </span>
                <span v-else>
              <a @click="saveRow(record)">保存</a>
              <a-divider type="vertical" />
              <a @click="cancel(record.key)">取消</a>
            </span>
              </template>
              <span v-else>
            <a @click="toggle(record.key)">编辑</a>
            <a-divider type="vertical" />
            <a-popconfirm title="是否要删除此行？" @confirm="remove(record.key)">
              <a>删除</a>
            </a-popconfirm>
          </span>
            </template>
          </a-table>

        </a-form-item>
      </a-form>
    </a-modal>

    <a-card>
      <a-row class="tools">
        <a-button v-action:add type="primary" ghost @click="openModal">添加</a-button>
      </a-row>

      <a-table
        :columns="columns"
        :rowKey="item => item.record_id"
        :dataSource="data"
        :pagination="pagination"
        :loading="loading"
      >

        <template slot="hidden" slot-scope="row">
          <template v-if="row.hidden === 0">正常</template>
          <template v-else>禁用</template>
        </template>

        <template slot="tools" slot-scope="row">
          <a-button
            v-action:add
            type="primary"
            ghost
            @click="openActionModal(row)"
            style="margin-right: 15px">编辑</a-button>
          <a-button
            v-action:delete
            type="danger"
            ghost
            @click="showDeleteConfirm(row.record_id)">删除</a-button>
        </template>
      </a-table>
    </a-card>
  </div>
</template>
<script>

import { STable } from '@/components'
import { mapActions } from 'vuex'
import { menuList } from '@/api/auth'

const columns = [{
  title: '菜单名称',
  dataIndex: 'name'
}, {
  title: '排序值',
  dataIndex: 'sequence'
}, {
  title: '隐藏状态',
  scopedSlots: { customRender: 'hidden' }
}, {
  title: '菜单图标',
  dataIndex: 'icon'
}, {
  title: '操作',
  scopedSlots: { customRender: 'tools' }
}]

export default {
  components: {
    STable
  },
  data () {
    return {
      description: '列表使用场景：后台管理中的权限管理以及角色管理，可用于基于 RBAC 设计的角色权限控制，颗粒度细到每一个操作类型。',
      columns: columns,
      data: [],
      pagination: {},
      loading: false,
      visible: false,
      confirmLoading: false,
      form: this.$form.createForm(this),
      info: {},
      tree: [],
      selected: 0
    }
  },
  filters: {
    html: (value) => {
      const arrEntities = { 'nbsp': '  ' }
      return value.replace(/&(nbsp);/ig, function (all, t) { return arrEntities[t] })
    }
  },
  mounted () {
    this.fetch()
  },
  methods: {
    ...mapActions([
      'fetchRule',
      'fetchTree',
      'deleteRule'
    ]),
    fetch () {
      this.loading = true
      menuList().then(res => {
        this.data = res.list
      }).finally(() => {
        this.loading = false
      })
    },
    openModal () {
      this.visible = true
    },
    openActionModal: function (row) {
      this.visible = true
      this.selected = row.id
      this.$nextTick(() => {
        this.form.setFieldsValue(
          {
            title: row.title,
            name: row.name,
            role: row.role,
            pid: row.pid,
            status: row.status
          }
        )
      })
    },
    handleSubmit () {
      this.form.validateFields(
        (err, values) => {
          if (!err) {
            this.confirmLoading = true
            const action = this.selected === 0 ? 'addRule' : 'updateRule'
            values.selectId = this.selected
            this.$store.dispatch(action, values).then(res => {
              this.$notification['success']({
                message: '成功通知',
                description: this.selected === 0 ? '添加成功！' : '更新成功！'
              })
              this.fetch()
              this.handleCancel()
            })
              .finally(() => {
                this.confirmLoading = false
              })
          }
        }
      )
    },
    handleCancel () {
      this.visible = false
      this.selected = 0
      this.form.resetFields()
    },
    showDeleteConfirm (id) {
      const {
        deleteRule
      } = this
      this.$confirm({
        title: '确定删除此规则吗?',
        content: '',
        okText: '确定',
        okType: 'danger',
        cancelText: '取消',
        onOk: () => {
          deleteRule({ id: id }).then(res => {
            this.$notification['success']({
              message: '成功通知',
              description: '删除成功！'
            })
            this.fetch()
          })
        }
      })
    }
  }
}
</script>
