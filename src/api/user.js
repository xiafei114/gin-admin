import { axios } from '@/utils/request'

const router = 'users'

// 角色查询
export function getUserList (parameter) {
  return axios({
    url: `/v1/${router}?q=page`,
    method: 'get',
    params: parameter
  })
}

// 获得单一角色
export function getUser (id) {
  return axios({
    url: `/v1/${router}/${id}`,
    method: 'get'
  })
}

// 添加角色
export function addUser (data) {
  return axios({
    url: `/v1/${router}`,
    method: 'post',
    data
  })
}
// 修改角色
export function updateUser (id, data) {
  return axios({
    url: `/v1/${router}/${id}`,
    method: 'put',
    data
  })
}

// 删除角色
export function deleteUser (params) {
  return axios({
    url: `/v1/${router}/${params.id}`,
    method: 'delete'
  })
}
