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
            <a-form-item label="状态">
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

        <a-form-item label="拥有权限">
          <a-row :gutter="16" v-for="(permission, index) in permissions" :key="index">
            <a-col :xl="4" :lg="24">
              {{ permission.name }}：
            </a-col>
            <a-col :xl="20" :lg="24">
              <a-checkbox
                v-if="permission.actions != null && permission.actions.length > 0"
                :indeterminate="permission.indeterminate"
                :checked="permission.checkedAll"
                @change="onChangeCheckAll($event, permission)">
                全选
              </a-checkbox>
              <a-checkbox-group :options="permission.actions" v-model="permission.selected" @change="onChangeCheck(permission)" />
            </a-col>
          </a-row>
        </a-form-item>

      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>

import { getPermissionList } from '@/api/permission'
import { getRole } from '@/api/role'
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
        status: 0
      },
      rules: [],
      // 子数据
      permissions: []
    }
  },
  methods: {
    clear () {
      this.form.setFieldsValue(this.entity)
      this.entityId = ''
    },
    add () {
      // console.log(JSON.stringify(permissionList))
      // // setTimeout(() => {
      // //   this.form.setFieldsValue(this.entity)
      // // }, 0)
      // this.loadPermissions(permissionList)
      this.$nextTick(() => {
        this.loadPermissions()
        this.visible = true
        this.clear()
      })
    },
    onChangeCheck (permission) {
      permission.indeterminate = !!permission.selected.length && (permission.selected.length < permission.actions.length)
      permission.checkedAll = permission.selected.length === permission.actions.length
    },
    onChangeCheckAll (e, permission) {
      console.log('permission:', permission)

      Object.assign(permission, {
        selected: e.target.checked ? permission.actions.map(obj => obj.value) : [],
        indeterminate: false,
        checkedAll: e.target.checked
      })
    },
    async loadPermissions () {
      const { result } = await getPermissionList()
      this.permissions = result.data.map(permission => {
        // const options = permission.actions
        permission.checkedAll = false
        permission.selected = []
        permission.indeterminate = false
        if (permission.actions === null) {
          // permission.actionsOptions = options.map(option => {
          //   return {
          //     label: option.label,
          //     value: option.value
          //   }
          // })
          permission.actions = []
        }
        return permission
      })

      // const { result } await getPermissionList().then(res => {
      //   const result = res.result.data
      //   this.permissions = result.map(permission => {
      //     // const options = permission.actions
      //     permission.checkedAll = false
      //     permission.selected = []
      //     permission.indeterminate = false
      //     if (permission.actions === null) {
      //       // permission.actionsOptions = options.map(option => {
      //       //   return {
      //       //     label: option.label,
      //       //     value: option.value
      //       //   }
      //       // })
      //       permission.actions = []
      //     }
      //     return permission
      //   })
      // })
    },
    edit (record) {
      this.$nextTick(() => {
        this.visible = true
        this.loadEditInfo(record)
      })
    },
    handleSubmit () {
      const { form: { validateFields } } = this
      this.confirmLoading = true
      validateFields((errors, values) => {
        if (!errors) {
          // console.log('values', values)
          const action = this.entityId === '' ? 'addRole' : 'updateRole'
          values.record_id = this.entityId

          // 拼装个新的，删除没用的
          const formData = []
          this.permissions.forEach(permission => {
            if (permission.selected.length > 0) {
              const data = pick(permission, ['record_id', 'selected', 'checkedAll'])
              data.permission_id = data.record_id
              data.actions = data.selected
              delete data['selected']
              delete data['record_id']
              formData.push(data)
            }
          })

          console.log(JSON.stringify(formData))

          values.permissions = formData
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
    async loadEditInfo (data) {
      await this.loadPermissions()
      const { form } = this
      const { result } = await getRole(Object.assign(data.record_id))
      const formData = pick(result.data, ['index_code', 'name', 'sequence', 'hidden', 'icon', 'record_id', 'permissions'])
      this.entityId = data.record_id
      // console.log('formData', formData)
      form.setFieldsValue(formData)

      if (this.permissions) {
        // 先处理要勾选的权限结构
        const permissionsAction = {}
        formData.permissions.forEach(permission => {
          permissionsAction[permission.permission_id] = permission.actions
        })

        // console.log('permissionsAction', permissionsAction)
        // console.log(JSON.stringify(this.permissions))
        // 把权限表遍历一遍，设定要勾选的权限 action
        this.permissions.forEach(permission => {
          const selected = permissionsAction[permission.record_id]

          // console.log('selected', selected)
          permission.selected = selected || []
          this.onChangeCheck(permission)
        })

        // console.log('this.permissions', this.permissions)
      }
    }
  }
}
</script>
