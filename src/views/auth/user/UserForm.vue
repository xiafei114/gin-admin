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
        <a-form-item
          label="用户名"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }">
          <a-input v-decorator="['user_name', {rules: [{required: true, min: 2, message: '请输入至少两个字符的唯一标识号！'}]}]" />
        </a-form-item>
        <a-form-item
          label="真实姓名"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }">
          <a-input v-decorator="['real_name', {rules: [{required: true, min: 2, message: '请输入至少两个字符的规则名称！'}]}]" />
        </a-form-item>
        <a-form-item
          label="密码"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }">
          <a-input
            type="password"
            :placeholder="this.entityId === '' ? '请入登录密码' : '如需修改密码请输入新密码'"
            v-decorator="['password', {rules: [{required: this.entityId === '', message: '请输入登录密码!'}]}]" />
        </a-form-item>
        <a-form-item
          label="显示状态"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }">
          <a-select
            v-decorator="['status', {
              initialValue: 1,
              rules: [{ required: true, message: '请选择状态!' }]
            }]"
            placeholder="请选择"
          >
            <a-select-option :value="1">
              启用
            </a-select-option>
            <a-select-option :value="2">
              禁用
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item
          label="选择角色"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }">
          <a-select
            mode="tags"
            v-decorator="['roles', {
              rules: [{ required: true, message: '请选择所属组别!' }]
            }]"
            placeholder="请选择角色"
          >
            <a-select-option v-for="(role, index) in roles" :value="role.record_id" :key="index">
              {{ role.name }}
            </a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import { getUser } from '@/api/user'
import { getRoleList } from '@/api/role'
import pick from 'lodash.pick'

export default {
  data () {
    return {
      visible: false,
      confirmLoading: false,
      form: this.$form.createForm(this),
      entityId: '',
      roles: [],
      entity: {
        user_name: '',
        real_name: '',
        password: '',
        roles: []
      }
    }
  },
  methods: {
    add () {
      this.$nextTick(() => {
        this.loadRole()
      })
    },
    clear () {
      this.form.setFieldsValue(this.entity)
      this.entityId = ''
    },
    edit (record) {
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
          const action = this.entityId === '' ? 'addUser' : 'updateUser'
          values.record_id = this.entityId
          const roles = []

          values.roles.forEach(id => {
            const role = {}
            role.role_id = id
            roles.push(role)
          })

          values.roles = roles

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
    async loadEditInfo (data) {
      await this.loadRole()
      this.visible = false
      const { form } = this
      getUser(Object.assign(data.record_id))
        .then(res => {
          const formData = pick(res.result.data, ['user_name', 'real_name', 'status', 'roles', 'record_id'])
          this.entityId = formData.record_id
          console.log('formData', formData)
          const roles = []

          formData.roles.forEach(role => roles.push(role.role_id))
          formData.roles = roles
          form.setFieldsValue(formData)
          this.visible = true
        })
    },
    async loadRole () {
      const { result } = await getRoleList()
      this.roles = result.data
      this.visible = true
    }
  }
}
</script>
