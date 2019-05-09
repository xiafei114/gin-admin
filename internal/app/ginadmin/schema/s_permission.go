package schema

import (
	"time"
)

// Permission 权力对象
type Permission struct {
	RecordID  string              `json:"record_id" swaggo:"false,记录ID"`
	IndexCode string              `json:"index_code" binding:"required" swaggo:"true,唯一标识码"`
	Name      string              `json:"name" binding:"required" swaggo:"true,权力名称"`
	Sequence  int                 `json:"sequence" swaggo:"false,排序值"`
	Icon      string              `json:"icon" swaggo:"false,权力图标"`
	Status    int                 `json:"status" swaggo:"false,隐藏权力(0:不隐藏 1:隐藏)"`
	Creator   string              `json:"creator" swaggo:"false,创建者"`
	CreatedAt *time.Time          `json:"created_at" swaggo:"false,创建时间"`
	UpdatedAt *time.Time          `json:"updated_at" swaggo:"false,更新时间"`
	Actions   PermissionActions   `json:"actions" swaggo:"false,动作列表"`
	Resources PermissionResources `json:"resources" swaggo:"false,资源列表"`
}

// PermissionAction 权力动作对象
type PermissionAction struct {
	ID   string `json:"key" swaggo:"true,序号"`
	Code string `json:"value" swaggo:"true,动作编号"`
	Name string `json:"label" swaggo:"true,动作名称"`
}

// PermissionResource 权力资源对象
type PermissionResource struct {
	Code   string `json:"code" swaggo:"true,资源编号"`
	Name   string `json:"name" swaggo:"true,资源名称"`
	Method string `json:"method" swaggo:"true,请求方式"`
	Path   string `json:"path" swaggo:"true,请求路径"`
}

// PermissionQueryParam 查询条件
type PermissionQueryParam struct {
	RecordIDs        []string // 记录ID列表
	LikeName         string   // 权力名称(模糊查询)
	ParentID         *string  // 父级内码
	PrefixParentPath string   // 父级路径(前缀模糊查询)
	Status           *int     // 隐藏权力
}

// PermissionQueryOptions 查询可选参数项
type PermissionQueryOptions struct {
	PageParam        *PaginationParam // 分页参数
	IncludeActions   bool             // 包含动作列表
	IncludeResources bool             // 包含资源列表
}

// PermissionQueryResult 查询结果
type PermissionQueryResult struct {
	Data       Permissions
	PageResult *PaginationResult
}

// Permissions 权力列表
type Permissions []*Permission

// ToMap 转换为键值映射
func (a Permissions) ToMap() map[string]*Permission {
	m := make(map[string]*Permission)
	for _, item := range a {
		m[item.RecordID] = item
	}
	return m
}

// // ToTrees 转换为权力列表
// func (a Permissions) ToTrees() PermissionTrees {
// 	list := make(PermissionTrees, len(a))
// 	for i, item := range a {
// 		list[i] = &PermissionTree{
// 			RecordID:  item.RecordID,
// 			Name:      item.Name,
// 			Sequence:  item.Sequence,
// 			Icon:      item.Icon,
// 			Status:    item.Status,
// 			Actions:   item.Actions,
// 			Resources: item.Resources,
// 		}
// 	}
// 	return list
// }

func (a Permissions) fillLeafNodeID(tree *[]*PermissionTree, leafNodeIDs *[]string) {
	for _, node := range *tree {
		if node.Children == nil || len(*node.Children) == 0 {
			*leafNodeIDs = append(*leafNodeIDs, node.RecordID)
			continue
		}
		a.fillLeafNodeID(node.Children, leafNodeIDs)
	}
}

// // ToLeafRecordIDs 转换为叶子节点记录ID列表
// func (a Permissions) ToLeafRecordIDs() []string {
// 	var leafNodeIDs []string
// 	tree := a.ToTrees().ToTree()
// 	a.fillLeafNodeID(&tree, &leafNodeIDs)
// 	return leafNodeIDs
// }

// PermissionResources 权力资源列表
type PermissionResources []*PermissionResource

// ForEach 遍历资源数据
func (a PermissionResources) ForEach(fn func(*PermissionResource, int)) PermissionResources {
	for i, item := range a {
		fn(item, i)
	}
	return a
}

// ToMap 转换为键值映射
func (a PermissionResources) ToMap() map[string]*PermissionResource {
	m := make(map[string]*PermissionResource)
	for _, item := range a {
		m[item.Code] = item
	}
	return m
}

// PermissionActions 权力动作列表
type PermissionActions []*PermissionAction

// PermissionTree 权力树
type PermissionTree struct {
	RecordID   string              `json:"record_id" swaggo:"false,记录ID"`
	Name       string              `json:"name" binding:"required" swaggo:"true,权力名称"`
	Sequence   int                 `json:"sequence" swaggo:"false,排序值"`
	Icon       string              `json:"icon" swaggo:"false,权力图标"`
	Router     string              `json:"router" swaggo:"false,访问路由"`
	Status     int                 `json:"status" swaggo:"false,隐藏权力(0:不隐藏 1:隐藏)"`
	ParentID   string              `json:"parent_id" swaggo:"false,父级ID"`
	ParentPath string              `json:"parent_path" swaggo:"false,父级路径"`
	Resources  PermissionResources `json:"resources" swaggo:"false,资源列表"`
	Actions    PermissionActions   `json:"actions" swaggo:"false,动作列表"`
	Children   *[]*PermissionTree  `json:"children,omitempty" swaggo:"false,子级树"`
}

// PermissionTrees 权力树列表
type PermissionTrees []*PermissionTree

// ForEach 遍历数据项
func (a PermissionTrees) ForEach(fn func(*PermissionTree, int)) PermissionTrees {
	for i, item := range a {
		fn(item, i)
	}
	return a
}

// ToTree 转换为树形结构
func (a PermissionTrees) ToTree() []*PermissionTree {
	mi := make(map[string]*PermissionTree)
	for _, item := range a {
		mi[item.RecordID] = item
	}

	var list []*PermissionTree
	for _, item := range a {
		if item.ParentID == "" {
			list = append(list, item)
			continue
		}
		if pitem, ok := mi[item.ParentID]; ok {
			if pitem.Children == nil {
				var children []*PermissionTree
				children = append(children, item)
				pitem.Children = &children
				continue
			}
			*pitem.Children = append(*pitem.Children, item)
		}
	}
	return list
}
