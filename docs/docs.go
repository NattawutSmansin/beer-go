// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/": {
            "get": {
                "description": "ตัวอย่างการใช้งาน API เบื้องต้น",
                "tags": [
                    "Welcome to My Beer"
                ],
                "summary": "แสดงข้อความต้อนรับ",
                "responses": {}
            }
        },
        "/beer": {
            "get": {
                "description": "ตัวอย่างการใช้งาน API เบื้องต้น",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beer"
                ],
                "summary": "เรียกข้อมูล Beer ทั้งหมด",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Beer Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page (optional, default: 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit (optional, default: 10)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "ตัวอย่างการใช้งาน API เบื้องต้น",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beer"
                ],
                "summary": "เพิ่มข้อมูล Beer",
                "parameters": [
                    {
                        "description": "Beer Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBeer"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/beer/{id}": {
            "get": {
                "description": "ตัวอย่างการใช้งาน API เบื้องต้น",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beer"
                ],
                "summary": "เรียกข้อมูล Beer ราย item",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "ตัวอย่างการใช้งาน API เบื้องต้น",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beer"
                ],
                "summary": "อัปเดตข้อมูล Beer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Beer Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateBeer"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "ตัวอย่างการใช้งาน API เบื้องต้น",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beer"
                ],
                "summary": "ลบข้อมูล Beer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/upload": {
            "post": {
                "description": "ตัวอย่างการใช้งาน API เบื้องต้น",
                "tags": [
                    "Upload"
                ],
                "summary": "อัปโหลดไฟล์",
                "parameters": [
                    {
                        "type": "file",
                        "description": "ไฟล์ที่ต้องการอัปโหลด",
                        "name": "file",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.CreateBeer": {
            "type": "object",
            "required": [
                "category_id",
                "description",
                "image_files",
                "name"
            ],
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "image_files": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.UpdateBeer": {
            "type": "object",
            "required": [
                "category_id",
                "description",
                "image_files",
                "name"
            ],
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_files": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Beer API Title",
	Description:      "This is a sample Beer API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
