// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/duc-cnzj/mars/v4/internal/ent/schema\",\"Package\":\"github.com/duc-cnzj/mars/v4/internal/ent\",\"Schemas\":[{\"name\":\"AccessToken\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"token\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":100,\"unique\":true,\"validators\":2,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"usage\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":50,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"expired_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"last_used_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"user_info\",\"type\":{\"Type\":3,\"Ident\":\"schematype.UserInfo\",\"PkgPath\":\"github.com/duc-cnzj/mars/v4/internal/ent/schema/schematype\",\"PkgName\":\"schematype\",\"Nillable\":false,\"RType\":{\"Name\":\"UserInfo\",\"Ident\":\"schematype.UserInfo\",\"Kind\":25,\"PkgPath\":\"github.com/duc-cnzj/mars/v4/internal/ent/schema/schematype\",\"Methods\":{\"GetID\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"IsAdmin\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"Json\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"optional\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"email\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"CacheLock\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"key\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"owner\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"expired_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}}]},{\"name\":\"Changelog\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"git_project\",\"type\":\"GitProject\",\"field\":\"git_project_id\",\"ref_name\":\"changelogs\",\"unique\":true,\"inverse\":true},{\"name\":\"project\",\"type\":\"Project\",\"field\":\"project_id\",\"ref_name\":\"changelogs\",\"unique\":true,\"inverse\":true}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"version\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1,\"default_kind\":2,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"username\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":100,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"修改人\"},{\"name\":\"manifest\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"config\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"用户提交的配置\"},{\"name\":\"config_type\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"optional\":true,\"validators\":1,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_branch\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"validators\":2,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_commit\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"validators\":2,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"docker_image\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"env_values\",\"type\":{\"Type\":3,\"Ident\":\"[]*types.KeyValue\",\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/types\",\"PkgName\":\"types\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]*types.KeyValue\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"可用的环境变量值\"},{\"name\":\"extra_values\",\"type\":{\"Type\":3,\"Ident\":\"[]*websocket.ExtraValue\",\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/websocket\",\"PkgName\":\"websocket\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]*websocket.ExtraValue\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":9,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"用户表单传入的额外值\"},{\"name\":\"final_extra_values\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":10,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"用户表单传入的额外值 + 系统默认的额外值\"},{\"name\":\"git_commit_web_url\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"optional\":true,\"validators\":1,\"position\":{\"Index\":11,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_commit_title\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"optional\":true,\"validators\":1,\"position\":{\"Index\":12,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_commit_author\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"optional\":true,\"validators\":1,\"position\":{\"Index\":13,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_commit_date\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":14,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"config_changed\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":15,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"project_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":16,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_project_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":17,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"project_id\",\"config_changed\",\"deleted_at\",\"version\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"DBCache\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"key\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"value\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"longtext\"}},{\"name\":\"expired_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}}],\"annotations\":{\"EntSQL\":{\"table\":\"db_cache\"}}},{\"name\":\"Event\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"file\",\"type\":\"File\",\"field\":\"file_id\",\"ref_name\":\"events\",\"unique\":true,\"inverse\":true}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"action\",\"type\":{\"Type\":11,\"Ident\":\"types.EventActionType\",\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/types\",\"PkgName\":\"types\",\"Nillable\":false,\"RType\":{\"Name\":\"EventActionType\",\"Ident\":\"types.EventActionType\",\"Kind\":5,\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/types\",\"Methods\":{\"Descriptor\":{\"In\":[],\"Out\":[{\"Name\":\"EnumDescriptor\",\"Ident\":\"protoreflect.EnumDescriptor\",\"Kind\":20,\"PkgPath\":\"google.golang.org/protobuf/reflect/protoreflect\",\"Methods\":null}]},\"Enum\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"*types.EventActionType\",\"Kind\":22,\"PkgPath\":\"\",\"Methods\":null}]},\"EnumDescriptor\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"\",\"Ident\":\"[]int\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Number\":{\"In\":[],\"Out\":[{\"Name\":\"EnumNumber\",\"Ident\":\"protoreflect.EnumNumber\",\"Kind\":5,\"PkgPath\":\"google.golang.org/protobuf/reflect/protoreflect\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Type\":{\"In\":[],\"Out\":[{\"Name\":\"EnumType\",\"Ident\":\"protoreflect.EnumType\",\"Kind\":20,\"PkgPath\":\"google.golang.org/protobuf/reflect/protoreflect\",\"Methods\":null}]}}}},\"default\":true,\"default_value\":0,\"default_kind\":5,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"username\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"用户名称\"},{\"name\":\"message\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"old\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"longtext\"}},{\"name\":\"new\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"longtext\"}},{\"name\":\"duration\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"file_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"action\"]},{\"fields\":[\"username\",\"created_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"File\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"events\",\"type\":\"Event\"}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"upload_type\",\"type\":{\"Type\":7,\"Ident\":\"schematype.UploadType\",\"PkgPath\":\"github.com/duc-cnzj/mars/v4/internal/ent/schema/schematype\",\"PkgName\":\"schematype\",\"Nillable\":false,\"RType\":{\"Name\":\"UploadType\",\"Ident\":\"schematype.UploadType\",\"Kind\":24,\"PkgPath\":\"github.com/duc-cnzj/mars/v4/internal/ent/schema/schematype\",\"Methods\":{}}},\"size\":100,\"default\":true,\"default_value\":\"local\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"path\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"文件全路径\"},{\"name\":\"size\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":11,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"int\"},\"comment\":\"文件大小\"},{\"name\":\"username\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"用户名称\"},{\"name\":\"namespace\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":100,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"pod\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":100,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"container\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":100,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"container_path\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"容器中的文件路径\"}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"GitProject\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"changelogs\",\"type\":\"Changelog\"}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"default_branch\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_project_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"enabled\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"global_enabled\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"global_config\",\"type\":{\"Type\":3,\"Ident\":\"*mars.Config\",\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/mars\",\"PkgName\":\"mars\",\"Nillable\":true,\"RType\":{\"Name\":\"Config\",\"Ident\":\"mars.Config\",\"Kind\":22,\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/mars\",\"Methods\":{\"Descriptor\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"\",\"Ident\":\"[]int\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"GetBranches\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"GetConfigField\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetConfigFile\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetConfigFileType\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetConfigFileValues\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetDisplayName\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetElements\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]*mars.Element\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"GetIsSimpleEnv\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"GetLocalChartPath\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetValuesYaml\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"ProtoMessage\":{\"In\":[],\"Out\":[]},\"ProtoReflect\":{\"In\":[],\"Out\":[{\"Name\":\"Message\",\"Ident\":\"protoreflect.Message\",\"Kind\":20,\"PkgPath\":\"google.golang.org/protobuf/reflect/protoreflect\",\"Methods\":null}]},\"Reset\":{\"In\":[],\"Out\":[]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Validate\":{\"In\":[],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"ValidateAll\":{\"In\":[],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"optional\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"全局配置\"}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"Namespace\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"projects\",\"type\":\"Project\"}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":100,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"项目空间名\"},{\"name\":\"image_pull_secrets\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"default\":true,\"default_value\":[],\"default_kind\":23,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"image pull secrets\"}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"Project\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"changelogs\",\"type\":\"Changelog\"},{\"name\":\"namespace\",\"type\":\"Namespace\",\"field\":\"namespace_id\",\"ref_name\":\"projects\",\"unique\":true,\"inverse\":true}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":100,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"项目名\"},{\"name\":\"git_project_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_branch\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"git 分支\"},{\"name\":\"git_commit\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"validators\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"git commit\"},{\"name\":\"config\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"override_values\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"longtext\"}},{\"name\":\"docker_image\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"docker 镜像\"},{\"name\":\"pod_selectors\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"pod 选择器\"},{\"name\":\"atomic\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"deploy_status\",\"type\":{\"Type\":11,\"Ident\":\"types.Deploy\",\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/types\",\"PkgName\":\"types\",\"Nillable\":false,\"RType\":{\"Name\":\"Deploy\",\"Ident\":\"types.Deploy\",\"Kind\":5,\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/types\",\"Methods\":{\"Descriptor\":{\"In\":[],\"Out\":[{\"Name\":\"EnumDescriptor\",\"Ident\":\"protoreflect.EnumDescriptor\",\"Kind\":20,\"PkgPath\":\"google.golang.org/protobuf/reflect/protoreflect\",\"Methods\":null}]},\"Enum\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"*types.Deploy\",\"Kind\":22,\"PkgPath\":\"\",\"Methods\":null}]},\"EnumDescriptor\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"\",\"Ident\":\"[]int\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Number\":{\"In\":[],\"Out\":[{\"Name\":\"EnumNumber\",\"Ident\":\"protoreflect.EnumNumber\",\"Kind\":5,\"PkgPath\":\"google.golang.org/protobuf/reflect/protoreflect\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Type\":{\"In\":[],\"Out\":[{\"Name\":\"EnumType\",\"Ident\":\"protoreflect.EnumType\",\"Kind\":20,\"PkgPath\":\"google.golang.org/protobuf/reflect/protoreflect\",\"Methods\":null}]}}}},\"default\":true,\"default_value\":0,\"default_kind\":5,\"position\":{\"Index\":9,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"部署状态\"},{\"name\":\"env_values\",\"type\":{\"Type\":3,\"Ident\":\"[]*types.KeyValue\",\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/types\",\"PkgName\":\"types\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]*types.KeyValue\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":10,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"环境变量值\"},{\"name\":\"extra_values\",\"type\":{\"Type\":3,\"Ident\":\"[]*websocket.ExtraValue\",\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/websocket\",\"PkgName\":\"websocket\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]*websocket.ExtraValue\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":11,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"额外值\"},{\"name\":\"final_extra_values\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":12,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"用户表单传入的额外值 + 系统默认的额外值\"},{\"name\":\"version\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1,\"default_kind\":2,\"position\":{\"Index\":13,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"版本\"},{\"name\":\"config_type\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"optional\":true,\"validators\":1,\"position\":{\"Index\":14,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"manifest\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":15,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"manifest\"},{\"name\":\"git_commit_web_url\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":16,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_commit_title\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":17,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_commit_author\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":18,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_commit_date\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":19,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"namespace_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":20,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"git_project_id\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"Repo\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"schema_type\":{\"mysql\":\"datetime\"}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"validators\":2,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"charset\":\"utf8mb4\",\"collation\":\"utf8mb4_0900_ai_ci\"}},\"comment\":\"默认使用的名称: helm create {name}\"},{\"name\":\"default_branch\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"optional\":true,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"git_project_name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"关联的 git 项目 name\"},{\"name\":\"git_project_id\",\"type\":{\"Type\":11,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"关联的 git 项目 id\"},{\"name\":\"enabled\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"need_git_repo\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"mars_config\",\"type\":{\"Type\":3,\"Ident\":\"*mars.Config\",\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/mars\",\"PkgName\":\"mars\",\"Nillable\":true,\"RType\":{\"Name\":\"Config\",\"Ident\":\"mars.Config\",\"Kind\":22,\"PkgPath\":\"github.com/duc-cnzj/mars/api/v4/mars\",\"Methods\":{\"Descriptor\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"\",\"Ident\":\"[]int\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"GetBranches\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"GetConfigField\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetConfigFile\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetConfigFileType\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetConfigFileValues\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetDisplayName\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetElements\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]*mars.Element\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"GetIsSimpleEnv\":{\"In\":[],\"Out\":[{\"Name\":\"bool\",\"Ident\":\"bool\",\"Kind\":1,\"PkgPath\":\"\",\"Methods\":null}]},\"GetLocalChartPath\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"GetValuesYaml\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"ProtoMessage\":{\"In\":[],\"Out\":[]},\"ProtoReflect\":{\"In\":[],\"Out\":[{\"Name\":\"Message\",\"Ident\":\"protoreflect.Message\",\"Kind\":20,\"PkgPath\":\"google.golang.org/protobuf/reflect/protoreflect\",\"Methods\":null}]},\"Reset\":{\"In\":[],\"Out\":[]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Validate\":{\"In\":[],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"ValidateAll\":{\"In\":[],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"optional\":true,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"mars 配置\"}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]}],\"Features\":[\"intercept\",\"schema/snapshot\",\"sql/upsert\"]}"
