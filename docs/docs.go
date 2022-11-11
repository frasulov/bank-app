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
        "/accounts": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get accounts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page ID",
                        "name": "page_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/account.AccountOutput"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Create a new account.",
                "parameters": [
                    {
                        "description": "account",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.CreateAccountInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/account.AccountOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get a account by ID.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/account.AccountOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        },
        "/transfer": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Transfer money between 2 accounts",
                "parameters": [
                    {
                        "description": "transfer",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.TransferInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.TransferTxResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "account.AccountOutput": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "account.CreateAccountInput": {
            "type": "object",
            "required": [
                "balance",
                "currency"
            ],
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "currency": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "account.TransferInput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "currency": {
                    "type": "string"
                },
                "from_account_id": {
                    "type": "integer"
                },
                "to_account_id": {
                    "type": "integer"
                }
            }
        },
        "db.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "db.Entry": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "amount": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "db.Transfer": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "from_account_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "to_account_id": {
                    "type": "integer"
                }
            }
        },
        "db.TransferTxResult": {
            "type": "object",
            "properties": {
                "from_account": {
                    "$ref": "#/definitions/db.Account"
                },
                "from_entry": {
                    "$ref": "#/definitions/db.Entry"
                },
                "to_account": {
                    "$ref": "#/definitions/db.Account"
                },
                "to_entry": {
                    "$ref": "#/definitions/db.Entry"
                },
                "transfer": {
                    "$ref": "#/definitions/db.Transfer"
                }
            }
        },
        "errors.Response": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8001",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "Sample App API",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
