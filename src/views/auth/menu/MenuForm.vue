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
            <a-form-item label="名称">
              <a-input v-decorator="['name', {rules: [{required: true, min: 2, message: '请输入至少两个字符的规则描述！'}]}]" />
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
                v-decorator="['hidden', {
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
          <a-col :lg="12" :md="12" :sm="24">
            <a-form-item label="图标">
              <a-input v-decorator="['icon', {rules: [{required: true, min: 2, message: '请输入图标！'}]}]" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>

import { getMenu } from '@/api/auth'
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
      }
    }
  },
  methods: {
    add () {
      this.visible = true
      this.form.setFieldsValue(this.entity)
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
          this.$store.dispatch(action, values).then(res => {
            console.log(res)
            this.$notification['success']({
              message: '成功通知',
              description: this.entityId === '' ? '添加成功！' : '更新成功！'
            })
            this.visible = false
            this.confirmLoading = false
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
      this.visible = false
    },
    loadEditInfo (data) {
      const { form } = this
      getMenu(Object.assign(data.record_id))
        .then(res => {
          const formData = pick(res.result.data, ['name', 'sequence', 'hidden', 'icon', 'record_id'])
          this.entityId = formData.record_id
          console.log('formData', formData)
          form.setFieldsValue(formData)
        })
    }
  }
}
</script>
