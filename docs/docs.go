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
        "/songs": {
            "put": {
                "description": "Update song info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Update song",
                "parameters": [
                    {
                        "description": "Data for update song info",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateSongDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/song.Song"
                        }
                    },
                    "400": {
                        "description": "bad request data"
                    }
                }
            },
            "post": {
                "description": "Find info by song name and group then record it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Create song",
                "parameters": [
                    {
                        "description": "Data for create song",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateSongDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/song.Song"
                        }
                    },
                    "400": {
                        "description": "bad request data"
                    }
                }
            },
            "delete": {
                "description": "Delete song by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Delete song",
                "parameters": [
                    {
                        "description": "Data with song id",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.IdDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/song.Song"
                        }
                    },
                    "400": {
                        "description": "bad request data"
                    }
                }
            }
        },
        "/songs/find": {
            "post": {
                "description": "Find by info and paginate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Find song",
                "parameters": [
                    {
                        "description": "Data for find song",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.FindSongDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/song.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request data"
                    }
                }
            }
        },
        "/songs/text": {
            "post": {
                "description": "Get song text and paginate by verses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Get text",
                "parameters": [
                    {
                        "description": "Data for get song text",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.GetTextDTO"
                        }
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
                    },
                    "400": {
                        "description": "bad request data"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CreateSongDTO": {
            "type": "object",
            "required": [
                "group",
                "song"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "Muse"
                },
                "song": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "Supermassive Black Hole"
                }
            }
        },
        "controller.FindSongDTO": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "Muse"
                },
                "limit": {
                    "type": "string",
                    "example": "1"
                },
                "link": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
                },
                "offset": {
                    "type": "string",
                    "example": "1"
                },
                "releaseDateBegin": {
                    "type": "string",
                    "example": "1717188495"
                },
                "releaseDateEnd": {
                    "type": "string",
                    "example": "1727188495"
                },
                "song": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "Supermassive Black Hole"
                },
                "text": {
                    "type": "string",
                    "maxLength": 10000,
                    "minLength": 3,
                    "example": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
                }
            }
        },
        "controller.GetTextDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string",
                    "example": "147367f5-93ef-432d-8a97-b06f716f9fad"
                },
                "limit": {
                    "type": "string",
                    "example": "1"
                },
                "offset": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "controller.IdDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string",
                    "example": "147367f5-93ef-432d-8a97-b06f716f9fad"
                }
            }
        },
        "controller.UpdateSongDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "Muse"
                },
                "id": {
                    "type": "string",
                    "example": "147367f5-93ef-432d-8a97-b06f716f9fad"
                },
                "link": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
                },
                "releaseDate": {
                    "type": "string",
                    "example": "16.07.2006"
                },
                "song": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "Supermassive Black Hole"
                },
                "text": {
                    "type": "string",
                    "maxLength": 10000,
                    "minLength": 3,
                    "example": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
                }
            }
        },
        "song.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "example": "Muse"
                },
                "id": {
                    "type": "string",
                    "example": "147367f5-93ef-432d-8a97-b06f716f9fad"
                },
                "link": {
                    "type": "string",
                    "example": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
                },
                "name": {
                    "type": "string",
                    "example": "Supermassive Black Hole"
                },
                "releaseDate": {
                    "type": "string",
                    "example": "2006-07-16T00:00:00+03:00"
                },
                "text": {
                    "type": "string",
                    "example": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Songs Library API",
	Description:      "API for library of songs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}