// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/family/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "家族模块"
                ],
                "summary": "创建家族",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FamilyAPICreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"dm_error\":0,\"error_msg\":\"\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/family/join": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "家族模块"
                ],
                "summary": "加入家族",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FamilyAPIJoinReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"dm_error\":0,\"error_msg\":\"\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/family/quit": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "家族模块"
                ],
                "summary": "退出家族",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FamilyAPIQuitReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"dm_error\":0,\"error_msg\":\"\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/family_graph/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "家族模块"
                ],
                "summary": "创建族谱图节点",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FamilyGraphAPICreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"dm_error\":0,\"error_msg\":\"\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/family_graph/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "家族模块"
                ],
                "summary": "删除族谱图节点",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FamilyGraphAPIDelReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"dm_error\":0,\"error_msg\":\"\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/family_graph/detail": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "家族模块"
                ],
                "summary": "族谱图节点详情",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FamilyGraphAPIDetailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FamilyGraphAPIDetailResp"
                        }
                    }
                }
            }
        },
        "/api/family_graph/graph": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "家族模块"
                ],
                "summary": "族谱图",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FamilyGraphAPIGraphReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FamilyGraphAPIGraphResp"
                        }
                    }
                }
            }
        },
        "/api/family_graph/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "家族模块"
                ],
                "summary": "更新族谱图节点",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FamilyGraphAPIUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"dm_error\":0,\"error_msg\":\"\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.FamilyAPICreateReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "portrait": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "model.FamilyAPIJoinReq": {
            "type": "object",
            "properties": {
                "family_id": {
                    "type": "integer"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "model.FamilyAPIQuitReq": {
            "type": "object",
            "properties": {
                "family_id": {
                    "type": "integer"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "model.FamilyGraphAPICreateReq": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "integer"
                },
                "current_node": {
                    "type": "integer"
                },
                "death_time": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "family_id": {
                    "type": "integer"
                },
                "father_node": {
                    "type": "integer"
                },
                "gender": {
                    "type": "integer"
                },
                "hometown": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "option": {
                    "description": "1-添加跟节点；2-添加父节点；3-添加孩子节点；4添加配偶节点",
                    "type": "integer"
                },
                "portrait": {
                    "type": "string"
                }
            }
        },
        "model.FamilyGraphAPIDelReq": {
            "type": "object",
            "properties": {
                "node": {
                    "type": "integer"
                }
            }
        },
        "model.FamilyGraphAPIDetailReq": {
            "type": "object",
            "properties": {
                "node": {
                    "type": "integer"
                }
            }
        },
        "model.FamilyGraphAPIDetailResp": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "integer"
                },
                "death_time": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "hometown": {
                    "type": "string"
                },
                "index_num": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "node": {
                    "type": "integer"
                },
                "portrait": {
                    "type": "string"
                }
            }
        },
        "model.FamilyGraphAPIGraphReq": {
            "type": "object",
            "properties": {
                "family_id": {
                    "type": "integer"
                }
            }
        },
        "model.FamilyGraphAPIGraphResp": {
            "type": "object",
            "properties": {
                "family_id": {
                    "type": "integer"
                },
                "graph": {
                    "$ref": "#/definitions/model.FamilyGraphTree"
                }
            }
        },
        "model.FamilyGraphAPIUpdateReq": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "integer"
                },
                "death_time": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "hometown": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "node": {
                    "type": "integer"
                },
                "portrait": {
                    "type": "string"
                }
            }
        },
        "model.FamilyGraphNode": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "integer"
                },
                "death_time": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "hometown": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "node": {
                    "type": "integer"
                },
                "portrait": {
                    "type": "string"
                }
            }
        },
        "model.FamilyGraphTree": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "integer"
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.FamilyGraphTree"
                    }
                },
                "death_time": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "hometown": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "node": {
                    "type": "integer"
                },
                "portrait": {
                    "type": "string"
                },
                "wives": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.FamilyGraphNode"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "39.108.101.229:80",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "家谱服务",
	Description: "家谱服务 API 文档.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}