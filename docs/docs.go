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
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categorys": {
            "get": {
                "description": "Get all Category from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get all Category",
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Category from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update Category",
                "parameters": [
                    {
                        "description": "Update Category",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/category.CategoryUpdate"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Insert a new Category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Insert a new Category",
                "parameters": [
                    {
                        "description": "Array Category to Inset",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/category.Category"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/categorys/:id": {
            "delete": {
                "description": "Delete category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id Category",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get all Product from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all Product",
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Product from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update Product",
                "parameters": [
                    {
                        "description": "Update Product",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/product.ProductUpdate"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Insert a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Insert a new product",
                "parameters": [
                    {
                        "description": "Array Product to insert",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/product.ProductInsert"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/products/:id": {
            "delete": {
                "description": "Delete product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id Product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/products/hide/:id": {
            "delete": {
                "description": "Spft Delete product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Soft Delete product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id Product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/products/name": {
            "get": {
                "description": "Get Product from the database Filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Product Filter Name data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pname",
                        "name": "pname",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/helper.SuccessResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "category.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "category.CategoryUpdate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "helper.SuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 2000
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "product.ProductInsert": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "product.ProductUpdate": {
            "type": "object",
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
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "update_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Intern_shopping",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
