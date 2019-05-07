import { axios } from '@/utils/request'

const router = 'menus'

// 菜单查询
export function menuList (parameter) {
  return axios({
    url: `/v1/${router}?q=page`,
    method: 'get',
    params: parameter
  })
}

// 获得单一菜单
export function getMenu (id) {
  return axios({
    url: `/v1/${router}/${id}`,
    method: 'get'
  })
}

// 添加菜单
export function addMenu (data) {
  return axios({
    url: `/v1/${router}`,
    method: 'post',
    data
  })
}
// 修改菜单
export function updateMenu (id, data) {
  return axios({
    url: `/v1/${router}/${id}`,
    method: 'put',
    data
  })
}

// 删除菜单
export function deleteMenu (params) {
  return axios({
    url: `/v1/${router}/${params.id}`,
    method: 'delete'
  })
}
