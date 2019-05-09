import {
  addPermission, updatePermission
} from '@/api/permission'

import {
  addRole, updateRole
} from '@/api/role'

import {
  addUser, updateUser
} from '@/api/user'

const auth = {
  state: {
  },

  mutations: {
  },

  actions: {

    // 添加权力
    addMenu (state, data) {
      return new Promise((resolve, reject) => {
        addPermission(data).then(response => {
          const result = response.data
          resolve(result)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 更新权力
    updateMenu (state, data) {
      const id = data.record_id
      delete data.record_id
      return new Promise((resolve, reject) => {
        updatePermission(id, data).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 添加角色
    addRole (state, data) {
      return new Promise((resolve, reject) => {
        addRole(data).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 更新角色
    updateRole (state, data) {
      const id = data.record_id
      delete data.entityId
      return new Promise((resolve, reject) => {
        updateRole(id, data).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 添加管理员
    addUser (state, data) {
      return new Promise((resolve, reject) => {
        addUser(data).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 更新管理员
    updateUser (state, data) {
      const id = data.record_id
      delete data.entityId
      return new Promise((resolve, reject) => {
        updateUser(id, data).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    }

  }
}

export default auth
