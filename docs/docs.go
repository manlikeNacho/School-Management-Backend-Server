// Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "Iheanacho Emmanuel",
            "url": "github.com/manlikeNacho",
            "email": "eiheanacho52@gmail.com"
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
        "/course/:id/student": {
            "put": {
                "description": "create student account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Create Student Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StudentResponse"
                        }
                    }
                }
            }
        },
        "/list/course": {
            "get": {
                "description": "List courses total courses taken by all students",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "List courses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Student ID",
                        "name": "course_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CourseResponse"
                        }
                    }
                }
            }
        },
        "/list/student": {
            "get": {
                "description": "List total students in the db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "List Students",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.StudentResponse"
                            }
                        }
                    }
                }
            }
        },
        "/list/student/:course_name": {
            "get": {
                "description": "create student account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Create Student Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Student ID",
                        "name": "course_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/student": {
            "post": {
                "description": "create student account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Create Student Account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StudentResponse"
                        }
                    }
                }
            }
        },
        "/student/:id": {
            "get": {
                "description": "get student by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Get Student",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StudentResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update student Name/email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Update Student account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StudentResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete student account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Delete Student Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StudentResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CourseResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.StudentResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Alt-School Adminstration",
	Description:      "This is a sample server School admin server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
