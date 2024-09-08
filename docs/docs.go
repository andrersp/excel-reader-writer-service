// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/reader": {
            "post": {
                "description": "Reader excel file and return data",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Reader"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "this is a excel file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/writer": {
            "post": {
                "description": "Create excel file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/vnd.ms-excel"
                ],
                "tags": [
                    "Writer"
                ],
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "Payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/XlsxRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Column": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "width": {
                    "type": "number"
                }
            }
        },
        "Sheet": {
            "type": "object",
            "properties": {
                "columns": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Column"
                    }
                },
                "data": {
                    "description": "AdditionalInfo *[]AdditionalData ` + "`" + `json:\"additioanlData\"` + "`" + `",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/xlsx.Data"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "XlsxRequest": {
            "type": "object",
            "properties": {
                "sheets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Sheet"
                    }
                }
            }
        },
        "xlsx.Data": {
            "type": "object",
            "additionalProperties": true
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Excel Reader and Writer",
	Description:      "Service to read and writer xlsx.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
