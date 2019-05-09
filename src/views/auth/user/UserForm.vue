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
            <a-form-item label="用户名">
              <a-input v-decorator="['user_name', {rules: [{required: true, min: 2, message: '请输入至少两个字符的唯一标识号！'}]}]" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row class="form-row" :gutter="16">
          <a-col :lg="12" :md="12" :sm="24">
            <a-form-item label="真实姓名">
              <a-input v-decorator="['real_name', {rules: [{required: true, min: 2, message: '请输入至少两个字符的规则名称！'}]}]" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row class="form-row" :gutter="16">
          <a-col :lg="12" :md="12" :sm="24">
            <a-form-item label="密码">
              <a-input
                type="password"
                :placeholder="this.entityId === '' ? '请入登录密码' : '如需修改密码请输入新密码'"
                v-decorator="['password', {rules: [{required: this.entityId === '', message: '请输入登录密码!'}]}]" />
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
                  启用
                </a-select-option>
                <a-select-option :value="1">
                  禁用
                </a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>

        <a-form-item label="选择角色">
          <a-select
            mode="tags"
            v-decorator="['rules', {
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
      entity: {
        user_name: '',
        real_name: '',
        password: '',
        roles: []
      },
      roles: []
    }
  },
  methods: {
    add () {
      this.$nextTick(() => {
        this.loadRole()
        this.visible = true
      })
    },
    clear () {
      this.form.setFieldsValue(this.entity)
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
          const action = this.entityId === '' ? 'addUser' : 'updateUser'
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
    async loadEditInfo (data) {
      await this.loadRole()
      const { form } = this
      getUser(Object.assign(data.record_id))
        .then(res => {
          const formData = pick(res.result.data, ['user_name', 'real_name', 'status', 'roles', 'record_id'])
          this.entityId = formData.record_id
          console.log('formData', formData)
          form.setFieldsValue(formData)
        })
    },
    async loadRole () {
      const { result } = await getRoleList()
      this.roles = result.data
    }
  }
}
</script>
