import { axios } from '@/utils/request'

const router = 'permission'

// 菜单查询
export function getPermissionPageList (parameter) {
  return axios({
    url: `/v1/${router}?q=page`,
    method: 'get',
    params: parameter
  })
}

export function getPermissionList (parameter) {
  return axios({
    url: `/v1/${router}?q=list`,
    method: 'get',
    params: parameter
  })
}

// 获得单一菜单
export function getPermission (id) {
  return axios({
    url: `/v1/${router}/${id}`,
    method: 'get'
  })
}

// 添加菜单
export function addPermission (data) {
  return axios({
    url: `/v1/${router}`,
    method: 'post',
    data
  })
}
// 修改菜单
export function updatePermission (id, data) {
  return axios({
    url: `/v1/${router}/${id}`,
    method: 'put',
    data
  })
}

// 删除菜单
export function deletePermission (params) {
  return axios({
    url: `/v1/${router}/${params.id}`,
    method: 'delete'
  })
}
