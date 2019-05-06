import { axios } from '@/utils/request'

// 菜单查询
export function menuList (parameter) {
  return axios({
    url: '/v1/menus?q=page',
    method: 'get',
    params: parameter
  })
}

// 获得单一菜单
export function getMenu (id) {
  return axios({
    url: '/v1/menus/' + id,
    method: 'get'
  })
}

// 添加菜单
export function addMenu (data) {
  return axios({
    url: '/v1/menus',
    method: 'post',
    data
  })
}
// 修改菜单
export function updateMenu (id, data) {
  return axios({
    url: '/v1/menus/' + id,
    method: 'put',
    data
  })
}

// 删除菜单
export function deleteMenu (params) {
  return axios({
    url: '/v1/menus/' + params.id,
    method: 'delete'
  })
}

export function fetchRule () {
  return axios({
    url: '/auth/rule',
    method: 'get'
  })
}

export function fetchTree () {
  return axios({
    url: '/auth/tree',
    method: 'get'
  })
}



export function fetchRole () {
  return axios({
    url: '/auth/role',
    method: 'get'
  })
}

export function addRole (data) {
  return axios({
    url: '/auth/role',
    method: 'post',
    data
  })
}

export function updateRole (id, data) {
  return axios({
    url: '/auth/role/' + id,
    method: 'put',
    data
  })
}

export function deleteRole (params) {
  return axios({
    url: '/auth/role/' + params.id,
    method: 'delete'
  })
}

export function fetchAccount () {
  return axios({
    url: '/auth/user',
    method: 'get'
  })
}

export function addAccount (data) {
  return axios({
    url: '/auth/user',
    method: 'post',
    data
  })
}

export function updateAccount (id, data) {
  return axios({
    url: '/auth/user/' + id,
    method: 'put',
    data
  })
}

export function deleteAccount (params) {
  return axios({
    url: '/auth/user/' + params.id,
    method: 'delete'
  })
}
