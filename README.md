# Boilerplate Clean Architecture

# Architecture

![Architecture](architecture.drawio.png)

```
.
├── adapter
│   ├── controller
│   │   ├── user_handler.go
│   │   └── user_handler_mapper.go
│   └── gateway
│       ├── user_repository_impl.go
│       └── user_repository_mapper.go
├── domain
│   └── entity
│       └── user.go
└── usecase
    ├── user_repository_port.go
    ├── user_usecase_impl.go
    ├── user_usecase_mapper.go
    └── user_usecase_port.go
```

## Adapter

Web API やデータベース など外部サービスとのやり取りを担当する。外部フォーマット ⇔ Use Case Data Structure（または Entity）の変換やエラーフォーマットの変換も行う。  
これにより、外部とのやり取り・仕様に関する詳細を Adapter レイヤーに隠蔽し、外部仕様変更に伴う上位レイヤーへの影響を最小限にする。  
本アーキテクチャでは、Controller が Web API とのやり取りを担当し、Gateway がデータベースとのやり取りを担当する。

### Controller

Web API とのやり取りを担当する。  
リクエストデータを Use Case Data Structure に変換の上、Use Case レイヤーに連携する。レスポンス時は、Use Case レイヤーの出力データを Web API フォーマットに変換して返却する。  
Use Case レイヤーとの連携は、Use Case Port を介して行う。

Web API フォーマットと Use Case Data Structure の変換は、Controller Mapper を参照する。

```go
// user_handler.go

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	req, err := ToDTO(r.Body)
	if err != nil {
		HttpError(w, err)
		return
	}

	result, err := h.usecase.AddUser(r.Context(), req)
	if err != nil {
		HttpError(w, err)
		return
	}

	h.HandleOK(w, FromDTO(result))
}
```

### Controller Mapper

Web API フォーマットと Use Case Data Structure の変換を行う。  
リクエスト時は、リクエストボディを Use Case Data Structure に変換。レスポンス時は、Use Case Data Structure を API レスポンスに変換する。

また、エラーが発生した際はアプリケーションエラーから HTTP エラーへの変換も行う。（詳細は [Error Strategy](#error-strategy) を参照）

```go
// user_handler_mapper.go

// リクエストボディを Use Case Data Structure に変換
func ToDTO(
	body io.ReadCloser,
) (*usecase.User, *pkgErr.ApplicationError) {
	var dto usecase.User
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, pkgErr.NewApplicationError(err.Error(), pkgErr.LevelWarn, pkgErr.CodeBadRequest)
	}
	return &usecase.User{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Age:       dto.Age,
	}, nil
}

// Use Case Data Structure を API レスポンスに変換
func FromDTO(
	dto *usecase.User,
) *User {
	return &User{
		Id:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Age:       int(dto.Age),
	}
}
```

### Gateway

データベースとのやり取りを担当する。Repository Port（インターフェース）の実装。  
Use Case レイヤーから連携された Entity を ORM モデルに変換の上、データベースとのやり取りを行う。反対に、Use Case レイヤーへの返却時は、ORM モデルを Entity に変換して返却する。

ORM モデルと Entity の変換は、Gateway Mapper を参照する。

```go
// user_repository_impl.go

func (u *UserRepositoryImpl) Save(ctx context.Context, entity *entity.User) (*entity.User, *pkgErr.ApplicationError) {
	tx := ctx.Value(TX_KEY).(*bun.Tx)

	user := FromEntity(entity)
	if _, err := tx.NewInsert().Model(user).Exec(ctx); err != nil {
		return nil, RepositoryError(err)
	}
	return user.ToEntity(), nil
}
```

### Gateway Mapper

ORM モデルと Entity の変換を行う。

また、データベースエラーが発生した際はアプリケーションエラーへの変換も行う。（詳細は [Error Strategy](#error-strategy) を参照）

```go
// user_repository_mapper.go

// ORM モデル
type User struct {
	ID        string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	FirstName string    `bun:"first_name,notnull"`
	LastName  string    `bun:"last_name,notnull"`
	Age       int32     `bun:"age,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}


// ORM モデルを Entity に変換
func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}


// Entity を ORM モデルに変換
func FromEntity(
	entity *entity.User,
) *User {
	return &User{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Age:       entity.Age,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
```

## Use Case

ビジネスロジックを担当する。Use Case Port（インターフェース）の実装。  
Use Case Data Structure から Entity への変換を行い、Repository に渡す。Repository との連携は、Repository Port を介して行う。

```go
// user_usecase_port.go

type UserUsecase interface {
	AddUser(ctx context.Context, dto *User) (*User, *pkgErr.ApplicationError)
}
```

```go
// user_usecase_impl.go

func (u *UserUsecaseImpl) AddUser(
	ctx context.Context,
	dto *User,
) (*User, *pkgErr.ApplicationError) {
	entity, err := u.userRepository.Save(ctx, dto.ToEntity())
	if err != nil {
		return nil, err
	}
	return FromEntity(entity), nil
}
```

### Use Case Mapper

Use Case Data Structure と Entity の変換を行う。

```go
// user_usecase_mapper.go

// Use Case Data Structure
type User struct {
	ID        string
	FirstName string
	LastName  string
	Age       int32
}

// Use Case Data Structure を Entity に変換
func (u *User) ToEntity() *entity.User {
	return &entity.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}
}

// Entity を Use Case Data Structure に変換
func FromEntity(
	entity *entity.User,
) *User {
	return &User{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Age:       entity.Age,
	}
}
```

依存方向を下位レイヤー → 上位レイヤーに限定させるため、Repository Port は Use Case 側に定義。

```go
// user_repository_port.go

type UserRepository interface {
	Save(ctx context.Context, e *entity.User) (*entity.User, *pkgErr.ApplicationError)
}
```

## Domain

ビジネスルールを表す。どのレイヤーにも依存しない最重要ビジネスデータ。

```go
// entity/user.go

// Entity
type User struct {
	ID        string
	FirstName string
	LastName  string
	Age       int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
```

# Error Strategy

アプリケーション独自のカスタムエラーを定義。エラーメッセージ、エラーレベル、エラーコードの 3 要素で構成され、エラーレベルとエラーコードによってエラーの重要度を表現し、ハンドリングする。  
外部エラー（データベースエラーなど）をカスタムエラーに変換し、上位レイヤーへのエラー詳細の流入を防ぐ。また、API レスポンス時は、カスタムエラーをプロトコルエラー（HTTP エラーなど）に変換し、返却する。

```go
// Custom Error

type ErrorLevel int8

const (
	_ ErrorLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

type ErrorCode int

const (
	_ ErrorCode = iota
	CodeBadRequest
	CodeNotFound
	CodeDuplicate
	CodeInternalServerError
)

type ApplicationError struct {
	message string
	level   ErrorLevel
	code    ErrorCode
}

func (e *ApplicationError) Error() string {
	return e.message
}

func (e *ApplicationError) Level() ErrorLevel {
	return e.level
}

func (e *ApplicationError) Code() ErrorCode {
	return e.code
}

func NewApplicationError(message string, level ErrorLevel, code ErrorCode) *ApplicationError {
	return &ApplicationError{
		message: message,
		level:   level,
		code:    code,
	}
}

```

## Database Error Mapping

```go
// user_repository_mapper.go

func RepositoryError(err error) *pkgErr.ApplicationError {
	switch err {
	case sql.ErrNoRows:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelWarn, pkgErr.CodeNotFound)
	default:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelError, pkgErr.CodeInternalServerError)
	}
}
```

### Usage

```go
// user_repository_impl.go

func (u *UserRepositoryImpl) Save(ctx context.Context, entity *entity.User) (*entity.User, *pkgErr.ApplicationError) {
	tx := ctx.Value(TX_KEY).(*bun.Tx)

	user := FromEntity(entity)
	if _, err := tx.NewInsert().Model(user).Exec(ctx); err != nil {
		return nil, RepositoryError(err) // 変換
	}
	return user.ToEntity(), nil
}
```

## HTTP Error Mapping

```go
// user_handler_mapper.go

func NotFoundError(w http.ResponseWriter, err *pkgErr.ApplicationError) {
	setHeaderContentType(w)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Error{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	})
}

func InternalServerError(w http.ResponseWriter, err *pkgErr.ApplicationError) {
	setHeaderContentType(w)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Error{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}

func HttpError(w http.ResponseWriter, err *pkgErr.ApplicationError) {
	switch err.Code() {
	case pkgErr.CodeNotFound:
		NotFoundError(w, err)
	default:
		InternalServerError(w, err)
	}
}
```

### Usage

```go
// user_handler.go

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	req, err := ToDTO(r.Body)
	if err != nil {
		HttpError(w, err) // 変換
		return
	}

	result, err := h.usecase.AddUser(r.Context(), req)
	if err != nil {
		HttpError(w, err) // 変換
		return
	}

	h.HandleOK(w, FromDTO(result))
}
```
