<template>
  <a-modal
    title="新建规则"
    :width="640"
    :visible="visible"
    :confirmLoading="confirmLoading"
    @ok="handleSubmit"
    @cancel="handleCancel"
  >
    <a-spin :spinning="confirmLoading">
      <a-form :form="form">
        <a-row class="form-row" :gutter="16">
          <a-col :lg="12" :md="12" :sm="24">
            <a-form-item label="唯一标识号">
              <a-input v-decorator="['index_code', {rules: [{required: true, min: 2, message: '请输入至少两个字符的唯一标识号！'}]}]" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row class="form-row" :gutter="16">
          <a-col :lg="12" :md="12" :sm="24">
            <a-form-item label="名称">
              <a-input v-decorator="['name', {rules: [{required: true, min: 2, message: '请输入至少两个字符的规则名称！'}]}]" />
            </a-form-item>
          </a-col>
          <a-col :lg="12" :md="12" :sm="24">
            <a-form-item label="排序值">
              <a-input-number
                v-decorator="['sequence', {
                  initialValue: 1000000,
                  rules: [{ required: true, message: '排序值不能为空!' }]
                }]"
              />
            </a-form-item>
          </a-col>
        </a-row>

        <a-row class="form-row" :gutter="16">
          <a-col :lg="12" :md="12" :sm="24">
            <a-form-item label="显示状态">
              <a-select
                v-decorator="['status', {
                  initialValue: 0,
                  rules: [{ required: true, message: '请选择状态!' }]
                }]"
                placeholder="请选择"
              >
                <a-select-option :value="0">
                  显示
                </a-select-option>
                <a-select-option :value="1">
                  隐藏
                </a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row class="form-row" :gutter="16">
          <a-col :lg="18" :md="12" :sm="24">
            <a-form-item label="菜单动作管理">
              <a-button style="" type="dashed" icon="plus" @click="newSubData">新增动作</a-button>
              <a-divider type="vertical" />
              <a-button style="" type="primary" icon="plus" @click="batchAddition">批量添加</a-button>
              <a-table
                :columns="subColumns"
                :dataSource="subData"
                :pagination="false"
                :loading="subDataLoading"
              >
                <template v-for="(col, i) in ['value', 'label']" :slot="col" slot-scope="text, record">
                  <a-input
                    :key="col"
                    style="margin: -5px 0"
                    :value="text"
                    :placeholder="subColumns[i].title"
                    @change="e => handleChange(e.target.value, record.key, col)"
                  />
                </template>
                <template slot="action" slot-scope="text, record">
                  <a-divider type="vertical" />
                  <a-popconfirm title="是否要删除此行？" @confirm="subDataRemove(record.key)">
                    <a>删除</a>
                  </a-popconfirm>
                </template>
              </a-table>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import { getPermission } from '@/api/permission'
import pick from 'lodash.pick'

export default {
  data () {
    return {
      visible: false,
      confirmLoading: false,
      form: this.$form.createForm(this),
      entityId: '',
      entity: {
        name: '',
        sequence: 100000,
        hidden: 0,
        icon: ''
      },
      subDataLoading: false,
      subColumns: [
        {
          title: '动作',
          key: 'value',
          dataIndex: 'value',
          scopedSlots: { customRender: 'value' }
        },
        {
          title: '名称',
          key: 'label',
          dataIndex: 'label',
          scopedSlots: { customRender: 'label' }
        },
        {
          title: '操作',
          dataIndex: 'action',
          width: '150px',
          scopedSlots: { customRender: 'action' }
        }
      ],
      // 子数据
      subData: []
    }
  },
  methods: {
    add () {
      this.visible = true
    },
    clear () {
      this.form.setFieldsValue(this.entity)
      this.subData = []
      this.entityId = ''
    },
    edit (record) {
      this.visible = true
      this.$nextTick(() => {
        this.loadEditInfo(record)
      })
    },
    handleSubmit () {
      const { form: { validateFields } } = this
      this.confirmLoading = true
      validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          const action = this.entityId === '' ? 'addMenu' : 'updateMenu'
          values.record_id = this.entityId
          values.actions = this.subData
          this.$store.dispatch(action, values).then(res => {
            console.log(res)
            this.$notification['success']({
              message: '成功通知',
              description: this.entityId === '' ? '添加成功！' : '更新成功！'
            })
            this.visible = false
            this.confirmLoading = false
            this.clear()
            this.$emit('ok', values)
          })
            .finally(() => {
              this.confirmLoading = false
            })
        } else {
          this.confirmLoading = false
        }
      })
    },
    handleCancel () {
      this.clear()
      this.visible = false
    },
    loadEditInfo (data) {
      const { form } = this
      getPermission(Object.assign(data.record_id))
        .then(res => {
          const formData = pick(res.result.data, ['name', 'sequence', 'hidden', 'icon', 'record_id', 'actions'])
          this.entityId = formData.record_id
          console.log('formData', formData)
          form.setFieldsValue(formData)
          if (formData.actions === null) {
            this.subData = []
          } else {
            this.subData = formData.actions
          }
        })
    },
    newSubData () {
      const length = this.subData.length
      this.subData.push({
        key: length === 0 ? '1' : (parseInt(this.subData[length - 1].key) + 1).toString(),
        code: '',
        name: ''
      })
    },
    subDataRemove (key) {
      this.subData = this.subData.filter(item => item.key !== key)
    },
    // 子数据保存
    handleChange (value, key, column) {
      const newData = [...this.subData]
      const target = newData.filter(item => key === item.key)[0]
      if (target) {
        target[column] = value
        this.subData = newData
      }
    },
    batchAddition () {
      const data = [{
        key: '1',
        value: 'add',
        label: '添加'
      },
      {
        key: '2',
        value: 'edit',
        label: '修改'
      },
      {
        key: '3',
        value: 'delete',
        label: '删除'
      },
      {
        key: '4',
        value: 'list',
        label: '查看'
      },
      {
        key: '5',
        value: 'get',
        label: '详情'
      }]

      this.subData = data
    }
  }
}
</script>
