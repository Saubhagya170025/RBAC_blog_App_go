server/
├── main.go
├── go.mod
├── go.sum
├── models/
│   ├── user.go
│   ├── role.go
│   ├── permission.go
│   ├── category.go
│   └── blog.go
├── database/
│   ├── connection.go
│   ├── migration.go
│   └── repository/
│       ├── user_repository.go
│       ├── role_repository.go
│       ├── permission_repository.go
│       ├── category_repository.go
│       └── blog_repository.go
├── handlers/
│   ├── auth_handler.go          # NEW - Login, Refresh, Logout
│   ├── user_handler.go
│   ├── role_handler.go
│   ├── permission_handler.go
│   ├── blog_handler.go
│   └── category_handler.go
├── routes/
│   └── routes.go
├── middleware/
│   └── auth_middleware.go       # UPDATED - JWT verification
├── utils/                        # NEW
│   ├── jwt.go                   # JWT generation and validation
│   └── password.go              # Password hashing utilities
├── config/                       # NEW
│   └── config.go                # Configuration (JWT secrets, etc.)
└── migrations/
    ├── 001_create_roles_table.sql
    ├── 002_create_permissions_table.sql
    ├── 003_create_users_table.sql
    ├── 004_create_categories_table.sql
    └── 005_create_blogs_table.sql