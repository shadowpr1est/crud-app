# Simple React + FastAPI CRUD App

## Описание

Это простое fullstack-приложение для управления списком элементов (CRUD), состоящее из:
- **Frontend:** React (Create React App)
- **Backend:** FastAPI (Python)

## Быстрый старт

### 1. Клонируйте репозиторий
```bash
git clone https://github.com/shadowpr1est/crud-app.git
cd crud-app
```

### 2. Установите зависимости

#### Backend (FastAPI)
```bash
pip install fastapi 'uvicorn[standard]'
```

#### Frontend (React)
```bash
npm install
```

### 3. Запустите backend
```bash
uvicorn main:app --reload --host 0.0.0.0 --port 8000
```

### 4. Запустите frontend
В новом терминале:
```bash
npm start
```

### 5. Откройте приложение
- Frontend: [http://localhost:3000](http://localhost:3000)
- Backend docs: [http://localhost:8000/docs](http://localhost:8000/docs)

## Структура проекта
```
crud-app/
├── main.py         # FastAPI entrypoint
├── schemas.py      # Pydantic-схемы
├── storage.py      # CRUD-логика и хранение
├── src/            # React frontend
├── public/         # React static files
├── package.json    # React dependencies
├── ...
```

## Примечания
- Данные хранятся в памяти (при перезапуске backend всё сбрасывается).
- Для доступа с других устройств используйте свой локальный IP вместо localhost.
- Для продакшена рекомендуется добавить базу данных и авторизацию.

---

**Автор:** alisher 
