// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/rate": {
            "get": {
                "description": "Запит має повертати поточний курс BTC до UAH використовуючи будь-який third party сервіс з публічним АРІ",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rate"
                ],
                "summary": "Отримати поточний курс BTC до UAH",
                "operationId": "rate",
                "responses": {
                    "200": {
                        "description": "Повертається актуальний курс BTC до UAH",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Invalid status value"
                    }
                }
            }
        },
        "/sendEmails": {
            "post": {
                "description": "Запит має отримувати актуальний курс BTC до UAH за допомогою third-party сервісу та відправляти його на всі електронні адреси, які були підписані раніше.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "Відправити e-mail з поточним курсом на всі підписані електронні пошти.",
                "operationId": "sendEmails",
                "responses": {
                    "200": {
                        "description": "E-mailʼи відправлено"
                    }
                }
            }
        },
        "/subscribe": {
            "post": {
                "description": "Запит має перевірити, чи немає данної електронної адреси в поточній базі даних (файловій) і, в разі її відсутності, записувати її. Пізніше, за допомогою іншого запиту ми будемо відправляти лист на ті електронні адреси, які будуть в цій базі.",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "Підписати емейл на отримання поточного курсу",
                "operationId": "subscribe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Електронна адреса, яку потрібно підписати",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "E-mail додано"
                    },
                    "400": {
                        "description": "BadRequest"
                    },
                    "409": {
                        "description": "Повертати, якщо e-mail вже є в базі даних (файловій)"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "GSES3 BTC application",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
