// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://fazilnbr.github.io/mypeosolal.web.portfolio/",
            "email": "fazilkp2000@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "authentication login",
                "operationId": "authentication login",
                "parameters": [
                    {
                        "description": "auth login",
                        "name": "Login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_domain.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication",
                    "User Authentication"
                ],
                "summary": "SignUp for users",
                "operationId": "SignUp authentication",
                "parameters": [
                    {
                        "description": "Worker Login",
                        "name": "WorkerLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_domain.User"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "password": {
                                            "type": "string"
                                        },
                                        "username": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    }
                }
            }
        },
        "/auth/token-refresh": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Refresh Token"
                ],
                "summary": "Refresh The Access Token",
                "operationId": "Refresh access token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Retrieve a product by ID",
                "operationId": "retriveproduct",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product Id : ",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update an existing product",
                "operationId": "updateproduct",
                "parameters": [
                    {
                        "description": "Product Detials",
                        "name": "productdetials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_domain.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create a new product",
                "operationId": "createproduct",
                "parameters": [
                    {
                        "description": "Product Detials",
                        "name": "productdetials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_domain.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete a product by ID",
                "operationId": "deleteproduct",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product Id : ",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    }
                }
            }
        },
        "/product/all": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "List all available products",
                "operationId": "listproduct",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_domain.Product": {
            "type": "object",
            "required": [
                "description",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "minLength": 15
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_domain.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "github_com_fazilnbr_GoCart-grpc-API-Gateway_pkg_utils_response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go + Gin Workey API",
	Description:      "This is a sample server Job Portal server. You can visit the GitHub repository at https://github.com/fazilnbr/Job_Portal_Project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
