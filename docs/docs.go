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
        "/profiles/delete/{email}": {
            "delete": {
                "description": "Delete a user profile by email",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a profile",
                "operationId": "delete-profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address of the profile to be deleted",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        },
        "/profiles/user": {
            "post": {
                "description": "Post a user profile",
                "produces": [
                    "application/json"
                ],
                "summary": "Post a profile",
                "operationId": "post-profile",
                "parameters": [
                    {
                        "description": "User data to be posted",
                        "name": "userData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PostProfile"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        },
        "/profiles/user/{email}": {
            "get": {
                "description": "Get a user profile by email",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a profile",
                "operationId": "get-profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address of the profile to be retrieved",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Edit a user profile by email",
                "produces": [
                    "application/json"
                ],
                "summary": "Edit a profile",
                "operationId": "edit-profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address of the profile to be edited",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data to be edited",
                        "name": "userData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EditProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.EditProfile": {
            "type": "object",
            "required": [
                "profile_image",
                "user_name"
            ],
            "properties": {
                "profile_image": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 30
                }
            }
        },
        "models.PostProfile": {
            "type": "object",
            "required": [
                "email",
                "profile_image",
                "user_name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "profile_image": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 30
                }
            }
        },
        "responses.UserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4100",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Profiles API",
	Description:      "This is a simple CRUD API for profiles",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
