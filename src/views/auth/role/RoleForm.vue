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

        <p slot="expandedRowRender" slot-scope="row">
          <a-row>
            <a-col :key="index" class='rule-list' span="12" v-for="(item, index) in rules">
              <a-row>
                <a-col span="4">{{ item.title }}：</a-col>
                <a-col span="20">
                  <template v-for="(action, i) in item.action">
                    <a-tag :key="i" v-if="row.rules.split(',').indexOf(action.id.toString()) !== -1">
                      {{ action.title }}
                    </a-tag>
                  </template>
                </a-col>
              </a-row>
            </a-col>
          </a-row>
        </p>

        <a-row class="form-row" :gutter="16">
          <a-col :lg="18" :md="12" :sm="24">
            <a-form-item label="拥有权限">
              <a-row :key="index" v-for="(item, index) in rules">
                <a-col :span="4">{{ item.name }}</a-col>
                <a-col :span="20">
                  <a-checkbox-group :options="item.actions" v-model="subData"/>
                </a-col>
              </a-row>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>

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
        hidden: 0,
        icon: ''
      },
      rules: [
        {
          id: 95,
          pid: 0,
          role: 'rule',
          title: '规则管理',
          status: 1,
          condition: '',
          name: 'admin/rbac/rules',
          action: [
            {
              role: 'add',
              title: '添加',
              name: 'admin/rbac/addRule',
              pid: 95,
              id: 96,
              value: 96,
              label: '添加',
              status: 1
            },
            {
              role: 'update',
              title: '修改',
              name: 'admin/rbac/updateRule',
              pid: 95,
              id: 97,
              value: 97,
              label: '修改',
              status: 1
            },
            {
              role: 'delete',
              title: '删除',
              name: 'admin/rbac/deleteRule',
              pid: 95,
              id: 98,
              value: 98,
              label: '删除',
              status: 1
            },
            {
              role: 'list',
              title: '查看',
              name: 'admin/rbac/rules',
              pid: 95,
              id: 129,
              value: 129,
              label: '查看',
              status: 1
            }
          ],
          permissionId: 'rule'
        },
        {
          id: 99,
          pid: 0,
          role: 'role',
          title: '角色管理',
          status: 1,
          condition: '',
          name: 'admin/rbac/groups',
          action: [
            {
              role: 'add',
              title: '添加',
              name: 'admin/rbac/addGroup',
              pid: 99,
              id: 100,
              value: 100,
              label: '添加',
              status: 1
            },
            {
              role: 'update',
              title: '修改',
              name: 'admin/rbac/updateGroup',
              pid: 99,
              id: 101,
              value: 101,
              label: '修改',
              status: 1
            },
            {
              role: 'delete',
              title: '删除',
              name: 'admin/rbac/deleteGroup',
              pid: 99,
              id: 102,
              value: 102,
              label: '删除',
              status: 1
            },
            {
              role: 'list',
              title: '查看',
              name: 'admin/rbac/groups',
              pid: 99,
              id: 130,
              value: 130,
              label: '查看',
              status: 1
            }
          ],
          permissionId: 'role'
        },
        {
          id: 103,
          pid: 0,
          role: 'account',
          title: '用户管理',
          status: 1,
          condition: '',
          name: 'admin/rbac/users',
          action: [
            {
              role: 'add',
              title: '添加',
              name: 'admin/rbac/addUser',
              pid: 103,
              id: 104,
              value: 104,
              label: '添加',
              status: 1
            },
            {
              role: 'update',
              title: '修改',
              name: 'admin/rbac/updateUser',
              pid: 103,
              id: 105,
              value: 105,
              label: '修改',
              status: 1
            },
            {
              role: 'delete',
              title: '删除',
              name: 'admin/rbac/deleteUser',
              pid: 103,
              id: 106,
              value: 106,
              label: '删除',
              status: 1
            },
            {
              role: 'list',
              title: '查看',
              name: 'admin/rbac/users',
              pid: 103,
              id: 131,
              value: 131,
              label: '查看',
              status: 1
            }
          ],
          permissionId: 'account'
        }
      ],
      // 子数据
      subData: []
    }
  },
  methods: {
    add (permissionList) {
      console.log(permissionList)
      // setTimeout(() => {
      //   this.form.setFieldsValue(this.entity)
      // }, 0)
      this.rules = permissionList
      this.visible = true
      this.subData = []
      this.entityId = ''
    },
    edit (permissionList, record) {
      this.visible = true
      this.$nextTick(() => {
        this.rules = permissionList
        this.loadEditInfo(record)
      })
    },
    handleSubmit () {
      const { form: { validateFields } } = this
      this.confirmLoading = true
      validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          const action = this.entityId === '' ? 'addRole' : 'updateRole'
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
      this.subData = []
    },
    async loadEditInfo (data) {
      const { form } = this
      const { result } = await getRole(Object.assign(data.record_id))
      const formData = pick(result.data, ['name', 'sequence', 'hidden', 'icon', 'record_id', 'actions'])
      this.entityId = formData.record_id
      console.log('formData', formData)
      form.setFieldsValue(formData)
      if (formData.actions === null) {
        this.subData = []
      } else {
        this.subData = formData.actions
      }
    }
  }
}
</script>
