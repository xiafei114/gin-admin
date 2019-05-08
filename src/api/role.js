import { axios } from '@/utils/request'

const router = 'roles'

// 角色查询
export function getRoleList (parameter) {
  return axios({
    url: `/v1/${router}?q=page`,
    method: 'get',
    params: parameter
  })
}

// 获得单一角色
export function getRole (id) {
  return axios({
    url: `/v1/${router}/${id}`,
    method: 'get'
  })
}

// 添加角色
export function addRole (data) {
  return axios({
    url: `/v1/${router}`,
    method: 'post',
    data
  })
}
// 修改角色
export function updateRole (id, data) {
  return axios({
    url: `/v1/${router}/${id}`,
    method: 'put',
    data
  })
}

// 删除角色
export function deleteRole (params) {
  return axios({
    url: `/v1/${router}/${params.id}`,
    method: 'delete'
  })
}
