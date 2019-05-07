package ginadmin

import (
	"context"

	"gin-admin/internal/app/ginadmin/bll"
	"gin-admin/internal/app/ginadmin/config"
	"gin-admin/internal/app/ginadmin/schema"
	"gin-admin/pkg/util"
)

// InitData 初始化应用数据
func InitData(ctx context.Context, obj *Object) error {
	err := loadCasbinPolicyData(ctx, obj)
	if err != nil {
		return err
	}

	if config.GetGlobalConfig().AllowInitPermission {
		return initPermissionData(ctx, obj)
	}

	return nil
}

func loadCasbinPolicyData(ctx context.Context, obj *Object) error {
	err := obj.Bll.Role.LoadAllPolicy(ctx)
	if err != nil {
		return err
	}

	return obj.Bll.User.LoadAllPolicy(ctx)
}

// initPermissionData 初始化权力数据
func initPermissionData(ctx context.Context, obj *Object) error {
	// 检查是否存在权力数据，如果不存在则初始化
	exists, err := obj.Bll.Permission.CheckDataInit(ctx)
	if err != nil {
		return err
	} else if exists {
		return nil
	}

	// 初始化权力
	var data schema.PermissionTrees
	err = util.JSONUnmarshal([]byte(PermissionData), &data)
	if err != nil {
		return err
	}

	return createPermissions(ctx, obj, "", data)
}

func createPermissions(ctx context.Context, obj *Object, parentID string, list schema.PermissionTrees) error {
	return bll.ExecTrans(ctx, obj.Model.Trans, func(ctx context.Context) error {
		for _, item := range list {
			sitem := schema.Permission{
				Name:      item.Name,
				Sequence:  item.Sequence,
				Icon:      item.Icon,
				Router:    item.Router,
				Hidden:    item.Hidden,
				ParentID:  parentID,
				Actions:   item.Actions,
				Resources: item.Resources,
			}
			nsitem, err := obj.Bll.Permission.Create(ctx, sitem)
			if err != nil {
				return err
			}

			if item.Children != nil && len(*item.Children) > 0 {
				err := createPermissions(ctx, obj, nsitem.RecordID, *item.Children)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// 初始化权力数据
const PermissionData = `
[
  {
    "name": "首页",
    "icon": "dashboard",
    "router": "/dashboard",
    "sequence": 1900000
  },
  {
    "name": "系统管理",
    "icon": "setting",
    "sequence": 1100000,
    "children": [
      {
        "name": "权力管理",
        "icon": "solution",
        "router": "/system/Permission",
        "sequence": 1190000,
        "actions": [
          { "code": "add", "name": "新增" },
          { "code": "edit", "name": "编辑" },
          { "code": "del", "name": "删除" },
          { "code": "query", "name": "查询" }
        ],
        "resources": [
          {
            "code": "query",
            "name": "查询权力数据",
            "method": "GET",
            "path": "/api/v1/Permissions"
          },
          {
            "code": "get",
            "name": "精确查询权力数据",
            "method": "GET",
            "path": "/api/v1/Permissions/:id"
          },
          {
            "code": "create",
            "name": "创建权力数据",
            "method": "POST",
            "path": "/api/v1/Permissions"
          },
          {
            "code": "update",
            "name": "更新权力数据",
            "method": "PUT",
            "path": "/api/v1/Permissions/:id"
          },
          {
            "code": "delete",
            "name": "删除权力数据",
            "method": "DELETE",
            "path": "/api/v1/Permissions/:id"
          }
        ]
      },
      {
        "name": "角色管理",
        "icon": "audit",
        "router": "/system/role",
        "sequence": 1180000,
        "actions": [
          { "code": "add", "name": "新增" },
          { "code": "edit", "name": "编辑" },
          { "code": "del", "name": "删除" },
          { "code": "query", "name": "查询" }
        ],
        "resources": [
          {
            "code": "query",
            "name": "查询角色数据",
            "method": "GET",
            "path": "/api/v1/roles"
          },
          {
            "code": "get",
            "name": "精确查询角色数据",
            "method": "GET",
            "path": "/api/v1/roles/:id"
          },
          {
            "code": "create",
            "name": "创建角色数据",
            "method": "POST",
            "path": "/api/v1/roles"
          },
          {
            "code": "update",
            "name": "更新角色数据",
            "method": "PUT",
            "path": "/api/v1/roles/:id"
          },
          {
            "code": "delete",
            "name": "删除角色数据",
            "method": "DELETE",
            "path": "/api/v1/roles/:id"
          },
          {
            "code": "queryPermission",
            "name": "查询权力数据",
            "method": "GET",
            "path": "/api/v1/Permissions"
          }
        ]
      },
      {
        "name": "用户管理",
        "icon": "user",
        "router": "/system/user",
        "sequence": 1170000,
        "actions": [
          { "code": "add", "name": "新增" },
          { "code": "edit", "name": "编辑" },
          { "code": "del", "name": "删除" },
          { "code": "query", "name": "查询" }
        ],
        "resources": [
          {
            "code": "query",
            "name": "查询用户数据",
            "method": "GET",
            "path": "/api/v1/users"
          },
          {
            "code": "get",
            "name": "精确查询用户数据",
            "method": "GET",
            "path": "/api/v1/users/:id"
          },
          {
            "code": "create",
            "name": "创建用户数据",
            "method": "POST",
            "path": "/api/v1/users"
          },
          {
            "code": "update",
            "name": "更新用户数据",
            "method": "PUT",
            "path": "/api/v1/users/:id"
          },
          {
            "code": "delete",
            "name": "删除用户数据",
            "method": "DELETE",
            "path": "/api/v1/users/:id"
          },
          {
            "code": "disable",
            "name": "禁用用户数据",
            "method": "PATCH",
            "path": "/api/v1/users/:id/disable"
          },
          {
            "code": "enable",
            "name": "启用用户数据",
            "method": "PATCH",
            "path": "/api/v1/users/:id/enable"
          },
          {
            "code": "queryRole",
            "name": "查询角色数据",
            "method": "GET",
            "path": "/api/v1/roles"
          }
        ]
      }
    ]
  }
]
`
