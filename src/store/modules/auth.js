import {
  fetchAccount, addAccount, updateAccount, deleteAccount
} from '@/api/auth'

import {
  addPermission, updatePermission
} from '@/api/permission'

import {
  addRole, updateRole
} from '@/api/role'

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
      const id = data.selectId
      delete data.selectId
      return new Promise((resolve, reject) => {
        updateRole(id, data).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 获取管理员
    fetchAccount (state) {
      return new Promise((resolve, reject) => {
        fetchAccount().then(response => {
          const result = response.data
          resolve(result)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 添加管理员
    addAccount (state, data) {
      return new Promise((resolve, reject) => {
        addAccount(data).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 更新管理员
    updateAccount (state, data) {
      const id = data.selectId
      delete data.selectId
      return new Promise((resolve, reject) => {
        updateAccount(id, data).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 删除管理员
    deleteAccount (state, params) {
      return new Promise((resolve, reject) => {
        deleteAccount(params).then(_ => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default auth
