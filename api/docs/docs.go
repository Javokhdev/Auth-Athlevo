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
        "/change-role": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update the role of a user with the specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Change user role",
                "parameters": [
                    {
                        "description": "Change role request",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.Role"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Role Changed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dashboard/access-count": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get the number of times a gym accessed the gym within the specified date range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dashboard"
                ],
                "summary": "Get the access count of a gym",
                "parameters": [
                    {
                        "type": "string",
                        "description": "gym ID",
                        "name": "gym_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "end_date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.AccessCountRes"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dashboard/booking-revenue": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get the total booking revenue within the specified date range for a gym",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dashboard"
                ],
                "summary": "Get total booking revenue for a gym",
                "parameters": [
                    {
                        "type": "string",
                        "description": "gym ID",
                        "name": "gym_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "end_date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.TotalRevenueRes"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dashboard/subscription/access-count": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get the number of times a user accessed the gym based on subscription ID within a date range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dashboard"
                ],
                "summary": "Get the access count by subscription ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Subscription ID",
                        "name": "subscription_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "end_date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.SubscriptionCountRes"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dashboard/subscription/booking-revenue": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get total booking revenue based on subscription ID within a date range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dashboard"
                ],
                "summary": "Get booking revenue by subscription ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Subscription ID",
                        "name": "subscription_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "end_date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.TotalRevenueBySubscriptionRes"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/forgot-password": {
            "post": {
                "description": "Request to reset user's password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Forgot password",
                "parameters": [
                    {
                        "description": "Email Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.GetByEmail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset email sent successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/refresh-token/{email}": {
            "get": {
                "description": "Request to reset user's access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRes"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Register User Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registered successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reset-password": {
            "put": {
                "description": "Reset user's password with a reset code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "Password Reset Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ResetPassReqBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a user with the specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/get-profiles": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve user profiles with optional filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user profiles",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by user ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by full name",
                        "name": "full_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by gym ID",
                        "name": "gym_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by phone number",
                        "name": "phone_number",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.UserRepeated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/passwords": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update the password of a user with the specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Change user password",
                "parameters": [
                    {
                        "description": "Updated password details",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ChangePasswordReqBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/profiles": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update the profile of a user with the specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Edit user profile",
                "parameters": [
                    {
                        "description": "Updated profile details",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.EditProfileReqBpdy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/setting": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve the settings of a user with the specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user settings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.Setting"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
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
                "description": "Update the settings of a user with the specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Edit user settings",
                "parameters": [
                    {
                        "description": "Updated setting details",
                        "name": "setting",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.Setting"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Setting updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/validate": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Validates a JWT token and returns the user ID and role.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Validate Token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.AccessCountRes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "auth.ChangePasswordReqBody": {
            "type": "object",
            "properties": {
                "current_password": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "auth.EditProfileReqBpdy": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gym_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.GetByEmail": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "auth.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.LoginRes": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.RegisterReq": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gym_id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.ResetPassReqBody": {
            "type": "object",
            "properties": {
                "email_code": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "auth.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "auth.Setting": {
            "type": "object",
            "properties": {
                "language": {
                    "type": "string"
                },
                "notification": {
                    "type": "string"
                },
                "privacy_level": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                }
            }
        },
        "auth.SubscriptionCountRes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "auth.TotalRevenueBySubscriptionRes": {
            "type": "object",
            "properties": {
                "totalRevenue": {
                    "type": "integer"
                }
            }
        },
        "auth.TotalRevenueRes": {
            "type": "object",
            "properties": {
                "totalRevenue": {
                    "type": "integer"
                }
            }
        },
        "auth.UserRepeated": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/auth.UserRes"
                    }
                }
            }
        },
        "auth.UserRes": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gym_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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
	Title:            "Authentication Service API",
	Description:      "API for Authentication Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
