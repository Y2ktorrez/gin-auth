# Gin Auth

Este proyecto es una API de autenticación construida con el framework Gin en Go, utilizando GORM para la gestión de base de datos.

## Tecnologías utilizadas
- **Go** con el framework **Gin** para el desarrollo de la API.
- **GORM** como ORM para la interacción con la base de datos.
- **Docker** (opcional) para despliegue y ejecución en contenedores.
- **JWT** para autenticación y autorización.

## Estructura del Proyecto

```
gin-auth/
│── cmd/           # Contiene el punto de entrada principal de la aplicación.
│── config/        # Archivos de configuración, como variables de entorno y ajustes generales.
│── data/          # Puede incluir archivos de base de datos o datos predefinidos.
│── middleware/    # Middlewares de Gin, como autenticación, logging, etc.
│── src/           # Código fuente principal de la aplicación.
│── tmp/           # Archivos temporales generados durante la ejecución.
│── util/          # Utilidades y funciones auxiliares para la aplicación.
│── .air.toml      # Configuración de Air, una herramienta para hot-reload en Go.
│── .env           # Variables de entorno del proyecto.
│── .gitignore     # Archivo para ignorar archivos no necesarios en Git.
│── go.mod         # Módulo de Go con dependencias del proyecto.
│── go.sum         # Verificación de versiones de dependencias de Go.
```

## Instalación y Ejecución

1. **Clonar el repositorio:**
   ```sh
   git clone <repo-url>
   cd gin-auth
   ```

2. **Instalar dependencias:**
   ```sh
   go mod tidy
   ```

3. **Configurar variables de entorno:**
   - Copiar el archivo `.env.example` a `.env` y modificarlo según la configuración deseada.

4. **Ejecutar la aplicación:**
   ```sh
   go run cmd/main.go
   ```

5. **(Opcional) Ejecutar con Air (hot-reload):**
   ```sh
   air
   ```

## Endpoints Principales

- `POST /register` - Registro de usuario.
- `POST /login` - Inicio de sesión.
- `GET /profile` - Información del usuario autenticado.


