
# Useful samples to add new APIs

API는 쉽게 말해 개발자들을 위한 편리한 dev tool이라 할 수 있습니다. 
한편, 기능을 개발하는 것부터 API를 제공하기 까지는 시일이 오래 걸리고 다소 복잡하거나 어려운 작업이 될 것 입니다.

본 프로젝트에서는 Go 언어, Echo web framework, Swagger 활용하여 REST API 및 문서를 제공하고 있습니다. 
또한, API 개발을 조금이나마 구조화/단순화 하기 위해 Model-View-Controller (MVC) 패턴을 적용하였습니다(Model과 Controller만 적용).

이를 바탕으로, 새로운 API 추가를 위해 참고할만한 Sample 코드를 제공합니다. 
이를 활용하면 단순히 Copy, paste, modify 하는 것으로 새로운 API를 추가 할 수 있을 것이라 생각합니다(✔표시 참고).

이 글은 다음 내용을 간략히 설명하고 있습니다.
- REST API 개발을 위한 프로젝트 구조 또는 레이아웃
- REST API의 URL 및 Method 설정
- REST API의 Handler, request body, response body 설정
- REST API의 Model 설정
- REST API 문서 생성
- REST API 문서 확인

###  REST API 개발을 위한 프로젝트 구조 또는 레이아웃

Go 애플리케이션 프로젝트의 기본 레이아웃을 소개하는 [Standard Go Project Layout](https://github.com/golang-standards/project-layout)를 바탕으로 인터넷의 다양한 자료들을 참고하여 프로젝트 구조/레이아웃을 구성하였습니다.

- Go 언어의 Package Oriented Design/Development 방식을 적용하여 관련된 기능과 타입을 그룹화하여 모듈성과 재사용성을 높임
	- 유사한 기능들을 한 File에 정의하고, 여러 File들을(관련된 여러 기능) 하나의 Package로 묶어 관리함
	- 예를 들어, `controller` 패키지에서 마이그레이션 관련 핸들러들은 `migration.go` 파일에, 추천 관련 핸들러들은 `recommendation.go` 파일에 정의됨

프로젝트의 구조/레이아웃 중 **REST API를 제공하는데 직접적으로 관련된 디렉토리 및 파일은 다음과 같습니다.** 
MVC 패턴을 접해봤으면 쉽게 이해할 수 있으리라 생각합니다.

**참고 - 프로젝트 경로는 `${HOME}/cm-beetle`로 가정합니다.**

```
.
├── Makefile
└── pkg
    ├── Makefile
    └── api
        └── rest
            ├── controller
            │   └── sample.go
            ├── docs
            │   ├── docs.go
            │   ├── swagger.json
            │   └── swagger.yaml
            ├── model
            │   └── sample.go
            ├── route
            │   └── sample.go
            └── server
                └── server.go 
```


### REST API의 URL 및 Method 설정

REST API URL 및 Method를 만들기 위해 다음 두 가지 Sample을 참고하시기 바랍니다.

예를 들어, `GET /beetle/sample/users`(i.e., Method: GET, URL: /beetle/sample/users) API 등을 생성할 때 참고하시기 바랍니다.

#### URL(i.e., route) group

관련된 API들을 묶어서 관리하기 위해 Group을 사용하였습니다.

`server/server.go`
```go
	// Beetle API group which has /beetle as prefix
	groupBase := e.Group("/beetle")

	// Sample API group (for developers to add new API)
	groupSample := groupBase.Group("/sample")
	route.RegisterSampleRoutes(groupSample)
```

✔Sample과 관련된 **소스 코드를 복사**하고 용도에 맞게 수정하시기 바랍니다.

#### URL, Method, and handler

위에서 설정한 `/beetle/sample` group에 세부 URL 및 Method를 설정할 수 있습니다.

예를 들어, API `GET /beetle/sample/users`가 호출될 시, 이를 처리할 핸들러`controller.GetUsers`가 할당되어 있습니다. 

`route/sample.go`
```go
package route

import (
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/labstack/echo/v4"
)

// /beetle/sample/*
func RegisterSampleRoutes(g *echo.Group) {
	g.GET("/users", controller.GetUsers)
	g.GET("/users/:id", controller.GetUser)
	g.POST("/users", controller.CreateUser)
	g.PUT("/users/:id", controller.UpdateUser)
	g.DELETE("/users/:id", controller.DeleteUser)
}
```

✔ `route/sample.go` **파일을 복사**하고 용도에 맞게 수정하시기 바랍니다.

### REST API의 Handler, request body, response body 설정

API handler 및 Request/response body를 만들기 위해 `controller/sample.go`의 Sample 소스 코드를 참고하시기 바랍니다.

먼저, **API documentation 생성 자동화를 위해 각 Handler 마다 Swagger declarative comments format (annotaion)이 작성 되어 있습니다.** 
자세한 내용을 [swaggo/swag](https://github.com/swaggo/swag)을 참고하시기 바랍니다.

기본적으로, **1 Handler, 1 Request body, 1 Response body인 틀을 구성하고자 했으며**,  
GET method와 같이 Request body를 필요로 하지 않는 경우에는 Reqeust body를 정의하지 않았습니다. 

위 1-1-1 구성을 유지하면서, Request/response body와 model이 연관되도록 Go 언어의 struct embedding을 적용하고 있습니다. 
자세한 내용은 아래 POST 예시를 참고하시기 바랍니다.
- Struct embedding: 객체 프로그래밍에서 서브클래스와 유사한 개념으로 하나의 구조체를 다른 구조체 내에 포함시킬 수 있습니다.

#### GET method

REST API에서 GET method의 경우 Request body를 필요로 하지 않아 Request body를 생략하였습니다. 
`GET /beetle/sample/users`의 Response body는 User들로 구성된 리스트입니다.

`GET /beetle/sample/users` 요청을 다루는 GetUsers handler가 정의되어 있고, 
상단에 API 문서 생성을 위한 Annotation 이 주석으로 추가되어 있습니다.

```go
// [Note]
// No RequestBody required for "GET /users"

type GetUsersResponse struct {
	Users []model.MyUser `json:"users"`
}

// GetUsers godoc
// @Summary Get a list of users
// @Description Get information of all users.
// @Tags [Sample] Users
// @Accept  json
// @Produce  json
// @Success 200 {object} GetUsersResponse "(sample) This is a sample description for success response in Swagger UI"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users [get]
func GetUsers(c echo.Context) error {

	// In this example, hardcoded data is returned
	users := []model.MyUser{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Anne Jacqueline Hathaway", Email: "Anne@example.com"},
		{ID: 3, Name: "Robert John Downey Jr.", Email: "Robert@example.com"},
	}
	return c.JSON(http.StatusOK, users)
}
```

#### POST method

REST API에서 POST method의 경우 Request body와 Request body를 필요로 합니다. 
1 Handler, 1 Request body, 1 Response body 구성을 유지하고 있습니다. 

Request body `CreateUserRequest`와 Response body `CreateUserResponse` 모두에 `model.MyUser` 구조체를 임베딩하여 상호 연관되도록 만들었습니다.

예를 들어, 아래 예제 코드의 두 구조체는 이름만 다르고 멤버 데이터는 같습니다.

```go
// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type CreateUserRequest struct {
	model.MyUser
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type CreateUserResponse struct {
	model.MyUser
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the given information.
// @Tags [Sample] Users
// @Accept  json
// @Produce  json
// @Param User body CreateUserRequest true "User information"
// @Success 201 {object} GetUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 400 {object} object "Invalid Request"
// @Router /sample/users [post]
func CreateUser(c echo.Context) error {
	u := new(model.MyUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	// Implement user creation logic (this is a simple example)
	u.ID = 100 // Unique ID generation logic needed in actual implementation

	return c.JSON(http.StatusCreated, u)
}

```

✔ `controller/sample.go` **파일을 복사**하고 용도에 맞게 수정하시기 바랍니다.

### REST API의 Model 설정

API handler, Request/response body에서 필요로 하는 Model을 만들기 위해 `model/sample.go`의 Sample 소스 코드를 참고하시기 바랍니다.

```go
type MyUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
```

✔ `model/sample.go` **파일을 복사**하고 용도에 맞게 수정하시기 바랍니다.


### REST API 문서 생성

미리 작성해둔 Swagger declarative comments format (annotaion)을 바탕으로 API documentation을 생성/업데이트 합니다.

swag가 설치되어 있지 않다면 아래 명령어를 통해 설치합니다.

(Optional) If you got an error because of missing `swag`, install `swag`:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

아래 명령을 수행하여 API documentation을 생성합니다.
`/pkg/api/rest/docs` 에 문서가 생성된다.

Update Swagger API document
```bash
cd ${HOME}/cm-beetle
make swag
```

### REST API 문서 확인

REST API 문서를 확인하기 위해 아래 명령들을 순차적으로 실행합니다.
요약하면 서버 빌드, 실행, Swagger UI에 접속하는 과정 입니다.

Setup environment variables
```bash
cd ${HOME}/cm-beetle
source conf/setup.env
```

Build server
```bash
cd ${HOME}/cm-beetle
make
```

Run server 
```bash
cd ${HOME}/cm-beetle
make run
```

Health-check CM-Beetle
```bash
curl http://localhost:8056/beetle/health

# Output if it's running successfully
# {"message":"CM-Beetle API server is running"}
```

Access to Swagger UI
- (Default link) http://localhost:8056/beetle/swagger/index.html
- 서버 실행 시, Swagger UI 링크, Default ID 및 Password를 확인할 수 있습니다.
