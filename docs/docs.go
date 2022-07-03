// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/get-full-link": {
            "get": {
                "description": "Get full link",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Get full link by short form",
                "operationId": "get-full-link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "full link",
                        "name": "link",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/post-link": {
            "post": {
                "description": "Post link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Get short link",
                "operationId": "post-link",
                "parameters": [
                    {
                        "description": "full link",
                        "name": "linkInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LinkInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "405": {
                        "description": "Method not allowed"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.LinkInfo": {
            "type": "object",
            "properties": {
                "full_link": {
                    "type": "string",
                    "example": "https://pkg.go.dev/"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Ozon fintech",
	Description:      "Zip linker service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
