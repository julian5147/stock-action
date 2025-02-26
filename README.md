# Stock Action 📈

Una aplicación web moderna para el seguimiento y análisis de acciones bursátiles.

## 🚀 Tecnologías Utilizadas

### Frontend
- React con TypeScript
- Vite como bundler
- Tailwind CSS para estilos
- Gestión de paquetes con pnpm

### Backend
- Go
- CockroachDB como base de datos
- Docker para contenerización
- Air para hot-reload en desarrollo

## 📋 Prerrequisitos

- Node.js (v22)
- pnpm
- Go (v1.24)
- Docker y Docker Compose
- CockroachDB

## 🛠️ Instalación

### Frontend

1. Navega al directorio frontend:
```bash
cd frontend
```

2. Instala las dependencias:
```bash
pnpm install
```

3. Copia el archivo de variables de entorno:
```bash
cp .env.example .env
```

4. Inicia el servidor de desarrollo:
```bash
pnpm dev
```

### Backend

1. Navega al directorio backend:
```bash
cd backend
```

2. Instala las dependencias de Go:
```bash
go mod download
```

3. Copia el archivo de variables de entorno:
```bash
cp .env.example .env
```

4. Inicia los servicios con Docker:
```bash
docker-compose up -d
```

5. Inicia el servidor de desarrollo con Docker:
```bash
docker-compose up backend --watch
```

## 🌟 Características

- Interfaz de usuario moderna y responsive
- Seguimiento en tiempo real de acciones
- Análisis de datos históricos
- API RESTful
- Base de datos distribuida con CockroachDB
- Contenerización completa con Docker

## 👥 Autores

- Julián Salgado - [Tu GitHub](https://github.com/julian5147)