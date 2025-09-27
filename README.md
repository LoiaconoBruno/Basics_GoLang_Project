# GoAuthAPI 🚀

API RESTful desarrollada en **Go (Golang)** para gestión de usuarios con autenticación segura. Permite crear usuarios, iniciar sesión, eliminar usuarios y manejar tokens JWT.

---

## 🔹 Tecnologías

- **Go 1.21+**
- **Gin** - Framework HTTP
- **GORM** - ORM para PostgreSQL/MySQL/SQLite
- **JWT** - JSON Web Tokens para autenticación
- **bcrypt** - Para el hashing de contraseñas
- **PostgreSQL** (o SQLite/MySQL según configuración)

---

## 🔹 Características

- Registro de usuarios con correo electrónico y contraseña
- Login con JWT para autenticación segura
- Middleware para rutas protegidas
- Obtener lista de usuarios
- Eliminar usuarios
- Manejo de errores estandarizado
- Passwords hasheadas con **bcrypt**
- API documentada con JSON responses claros

---

## 🔹 Instalación

1. Clonar el repositorio:
```bash
git clone https://github.com/tuusuario/go-auth-api.git
cd go-auth-api
