package models

import (
	"time"
)

type ColumnsGoadminMenu struct {
	ID, ParentID, Type, Order, Title, Header, Icon, Uri, CreatedAt, UpdatedAt string
}

type ColumnsGoadminOperationLog struct {
	ID, UserID, Path, Method, Ip, Input, CreatedAt, UpdatedAt string
}

type ColumnsGoadminPermission struct {
	ID, Name, Slug, HttpMethod, HttpPath, CreatedAt, UpdatedAt string
}

type ColumnsGoadminRoleMenu struct {
	RoleID, MenuID, CreatedAt, UpdatedAt string
}

type ColumnsGoadminRolePermission struct {
	RoleID, PermissionID, CreatedAt, UpdatedAt string
}

type ColumnsGoadminRoleUser struct {
	RoleID, UserID, CreatedAt, UpdatedAt string
}

type ColumnsGoadminRole struct {
	ID, Name, Slug, CreatedAt, UpdatedAt string
}

type ColumnsGoadminSession struct {
	ID, Sid, Values, CreatedAt, UpdatedAt string
}

type ColumnsGoadminUserPermission struct {
	UserID, PermissionID, CreatedAt, UpdatedAt string
}

type ColumnsGoadminUser struct {
	ID, Username, Password, Name, Avatar, RememberToken, CreatedAt, UpdatedAt string
}

type ColumnsSt struct {
	GoadminMenu           ColumnsGoadminMenu
	GoadminOperationLog   ColumnsGoadminOperationLog
	GoadminPermission     ColumnsGoadminPermission
	GoadminRoleMenu       ColumnsGoadminRoleMenu
	GoadminRolePermission ColumnsGoadminRolePermission
	GoadminRoleUser       ColumnsGoadminRoleUser
	GoadminRole           ColumnsGoadminRole
	GoadminSession        ColumnsGoadminSession
	GoadminUserPermission ColumnsGoadminUserPermission
	GoadminUser           ColumnsGoadminUser
}

var Columns = ColumnsSt{
	GoadminMenu: ColumnsGoadminMenu{
		ID:        "id",
		ParentID:  "parent_id",
		Type:      "type",
		Order:     "order",
		Title:     "title",
		Header:    "header",
		Icon:      "icon",
		Uri:       "uri",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	GoadminOperationLog: ColumnsGoadminOperationLog{
		ID:        "id",
		UserID:    "user_id",
		Path:      "path",
		Method:    "method",
		Ip:        "ip",
		Input:     "input",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	GoadminPermission: ColumnsGoadminPermission{
		ID:         "id",
		Name:       "name",
		Slug:       "slug",
		HttpMethod: "http_method",
		HttpPath:   "http_path",
		CreatedAt:  "created_at",
		UpdatedAt:  "updated_at",
	},
	GoadminRoleMenu: ColumnsGoadminRoleMenu{
		RoleID:    "role_id",
		MenuID:    "menu_id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	GoadminRolePermission: ColumnsGoadminRolePermission{
		RoleID:       "role_id",
		PermissionID: "permission_id",
		CreatedAt:    "created_at",
		UpdatedAt:    "updated_at",
	},
	GoadminRoleUser: ColumnsGoadminRoleUser{
		RoleID:    "role_id",
		UserID:    "user_id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	GoadminRole: ColumnsGoadminRole{
		ID:        "id",
		Name:      "name",
		Slug:      "slug",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	GoadminSession: ColumnsGoadminSession{
		ID:        "id",
		Sid:       "sid",
		Values:    "values",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	GoadminUserPermission: ColumnsGoadminUserPermission{
		UserID:       "user_id",
		PermissionID: "permission_id",
		CreatedAt:    "created_at",
		UpdatedAt:    "updated_at",
	},
	GoadminUser: ColumnsGoadminUser{
		ID:            "id",
		Username:      "username",
		Password:      "password",
		Name:          "name",
		Avatar:        "avatar",
		RememberToken: "remember_token",
		CreatedAt:     "created_at",
		UpdatedAt:     "updated_at",
	},
}

type TableGoadminMenu struct {
	Name, Alias string
}

type TableGoadminOperationLog struct {
	Name, Alias string
}

type TableGoadminPermission struct {
	Name, Alias string
}

type TableGoadminRoleMenu struct {
	Name, Alias string
}

type TableGoadminRolePermission struct {
	Name, Alias string
}

type TableGoadminRoleUser struct {
	Name, Alias string
}

type TableGoadminRole struct {
	Name, Alias string
}

type TableGoadminSession struct {
	Name, Alias string
}

type TableGoadminUserPermission struct {
	Name, Alias string
}

type TableGoadminUser struct {
	Name, Alias string
}

type TablesSt struct {
	GoadminMenu           TableGoadminMenu
	GoadminOperationLog   TableGoadminOperationLog
	GoadminPermission     TableGoadminPermission
	GoadminRoleMenu       TableGoadminRoleMenu
	GoadminRolePermission TableGoadminRolePermission
	GoadminRoleUser       TableGoadminRoleUser
	GoadminRole           TableGoadminRole
	GoadminSession        TableGoadminSession
	GoadminUserPermission TableGoadminUserPermission
	GoadminUser           TableGoadminUser
}

var Tables = TablesSt{
	GoadminMenu: TableGoadminMenu{
		Name:  "goadmin_menu",
		Alias: "t",
	},
	GoadminOperationLog: TableGoadminOperationLog{
		Name:  "goadmin_operation_log",
		Alias: "t",
	},
	GoadminPermission: TableGoadminPermission{
		Name:  "goadmin_permissions",
		Alias: "t",
	},
	GoadminRoleMenu: TableGoadminRoleMenu{
		Name:  "goadmin_role_menu",
		Alias: "t",
	},
	GoadminRolePermission: TableGoadminRolePermission{
		Name:  "goadmin_role_permissions",
		Alias: "t",
	},
	GoadminRoleUser: TableGoadminRoleUser{
		Name:  "goadmin_role_users",
		Alias: "t",
	},
	GoadminRole: TableGoadminRole{
		Name:  "goadmin_roles",
		Alias: "t",
	},
	GoadminSession: TableGoadminSession{
		Name:  "goadmin_session",
		Alias: "t",
	},
	GoadminUserPermission: TableGoadminUserPermission{
		Name:  "goadmin_user_permissions",
		Alias: "t",
	},
	GoadminUser: TableGoadminUser{
		Name:  "goadmin_users",
		Alias: "t",
	},
}

type GoadminMenu struct {
	tableName struct{} `sql:"goadmin_menu,alias:t" pg:",discard_unknown_columns"`

	ID        int        `sql:"id,pk"`
	ParentID  int        `sql:"parent_id,notnull"`
	Type      *int       `sql:"type"`
	Order     int        `sql:"order,notnull"`
	Title     string     `sql:"title,notnull"`
	Header    *string    `sql:"header"`
	Icon      string     `sql:"icon,notnull"`
	Uri       string     `sql:"uri,notnull"`
	CreatedAt *time.Time `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type GoadminOperationLog struct {
	tableName struct{} `sql:"goadmin_operation_log,alias:t" pg:",discard_unknown_columns"`

	ID        int        `sql:"id,pk"`
	UserID    int        `sql:"user_id,notnull"`
	Path      string     `sql:"path,notnull"`
	Method    string     `sql:"method,notnull"`
	Ip        string     `sql:"ip,notnull"`
	Input     string     `sql:"input,notnull"`
	CreatedAt *time.Time `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type GoadminPermission struct {
	tableName struct{} `sql:"goadmin_permissions,alias:t" pg:",discard_unknown_columns"`

	ID         int        `sql:"id,pk"`
	Name       string     `sql:"name,notnull"`
	Slug       string     `sql:"slug,notnull"`
	HttpMethod *string    `sql:"http_method"`
	HttpPath   string     `sql:"http_path,notnull"`
	CreatedAt  *time.Time `sql:"created_at"`
	UpdatedAt  *time.Time `sql:"updated_at"`
}

type GoadminRoleMenu struct {
	tableName struct{} `sql:"goadmin_role_menu,alias:t" pg:",discard_unknown_columns"`

	RoleID    int        `sql:"role_id,notnull"`
	MenuID    int        `sql:"menu_id,notnull"`
	CreatedAt *time.Time `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type GoadminRolePermission struct {
	tableName struct{} `sql:"goadmin_role_permissions,alias:t" pg:",discard_unknown_columns"`

	RoleID       int        `sql:"role_id,notnull"`
	PermissionID int        `sql:"permission_id,notnull"`
	CreatedAt    *time.Time `sql:"created_at"`
	UpdatedAt    *time.Time `sql:"updated_at"`
}

type GoadminRoleUser struct {
	tableName struct{} `sql:"goadmin_role_users,alias:t" pg:",discard_unknown_columns"`

	RoleID    int        `sql:"role_id,notnull"`
	UserID    int        `sql:"user_id,notnull"`
	CreatedAt *time.Time `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type GoadminRole struct {
	tableName struct{} `sql:"goadmin_roles,alias:t" pg:",discard_unknown_columns"`

	ID        int        `sql:"id,pk"`
	Name      string     `sql:"name,notnull"`
	Slug      string     `sql:"slug,notnull"`
	CreatedAt *time.Time `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type GoadminSession struct {
	tableName struct{} `sql:"goadmin_session,alias:t" pg:",discard_unknown_columns"`

	ID        int        `sql:"id,pk"`
	Sid       string     `sql:"sid,notnull"`
	Values    string     `sql:"values,notnull"`
	CreatedAt *time.Time `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type GoadminUserPermission struct {
	tableName struct{} `sql:"goadmin_user_permissions,alias:t" pg:",discard_unknown_columns"`

	UserID       int        `sql:"user_id,notnull"`
	PermissionID int        `sql:"permission_id,notnull"`
	CreatedAt    *time.Time `sql:"created_at"`
	UpdatedAt    *time.Time `sql:"updated_at"`
}

type GoadminUser struct {
	tableName struct{} `sql:"goadmin_users,alias:t" pg:",discard_unknown_columns"`

	ID            int        `sql:"id,pk"`
	Username      string     `sql:"username,notnull"`
	Password      string     `sql:"password,notnull"`
	Name          string     `sql:"name,notnull"`
	Avatar        *string    `sql:"avatar"`
	RememberToken *string    `sql:"remember_token"`
	CreatedAt     *time.Time `sql:"created_at"`
	UpdatedAt     *time.Time `sql:"updated_at"`
}
