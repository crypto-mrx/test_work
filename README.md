# todo

## Getting started


## Usage 
Есть две таблицы:  
- **users** (пользователи)  
- **todos** (список задач, привязанный к пользователю)  

Ниже описаны **Postman-запросы** для работы с `/api/v1/todos`.

---

## **1️⃣ Создать задачу (`Create Todo`)**  
📌 **Метод:** `POST http://localhost:8080/api/v1/todos`  
📌 **Body (JSON):**
```json
{
  "title": "Купить продукты",
  "description": "Купить молоко, хлеб и яйца",
  "user_id": "550e8400-e29b-41d4-a716-446655440000"
}
```
📌 **Ответ (201 Created):**  
```json
{
  "id": "1a2b3c4d-5678-90ab-cdef-1234567890ab",
  "title": "Купить продукты",
  "description": "Купить молоко, хлеб и яйца",
  "completed": false,
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "created_at": "2024-03-01T12:00:00Z",
  "updated_at": "2024-03-01T12:00:00Z"
}
```
📌 **Ошибки:**  
- `400 Bad Request` → если переданы некорректные данные  
- `500 Internal Server Error` → если ошибка на сервере  

---

## **2️⃣ Получить все задачи (`Get All Todos`)**  
📌 **Метод:** `GET http://localhost:8080/api/v1/todos`  
📌 **Ответ (200 OK):**  
```json
[
  {
    "id": "1a2b3c4d-5678-90ab-cdef-1234567890ab",
    "title": "Купить продукты",
    "description": "Купить молоко, хлеб и яйца",
    "completed": false,
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "created_at": "2024-03-01T12:00:00Z",
    "updated_at": "2024-03-01T12:00:00Z"
  },
  {
    "id": "2b3c4d5e-6789-01ab-cdef-2345678901bc",
    "title": "Выучить Golang",
    "description": "Прочитать документацию и написать тестовый проект",
    "completed": false,
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "created_at": "2024-03-01T13:00:00Z",
    "updated_at": "2024-03-01T13:00:00Z"
  }
]
```

---

## **3️⃣ Получить задачу по ID (`Get Todo By ID`)**  
📌 **Метод:** `GET http://localhost:8080/api/v1/todos/{id}`  
📌 **Пример запроса:**  
```
GET http://localhost:8080/api/v1/todos/1a2b3c4d-5678-90ab-cdef-1234567890ab
```
📌 **Ответ (200 OK):**  
```json
{
  "id": "1a2b3c4d-5678-90ab-cdef-1234567890ab",
  "title": "Купить продукты",
  "description": "Купить молоко, хлеб и яйца",
  "completed": false,
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "created_at": "2024-03-01T12:00:00Z",
  "updated_at": "2024-03-01T12:00:00Z"
}
```
📌 **Ошибки:**  
- `400 Bad Request` → если ID некорректный  
- `404 Not Found` → если задачи с таким ID нет  

---

## **4️⃣ Обновить задачу (`Update Todo`)**  
📌 **Метод:** `PUT http://localhost:8080/api/v1/todos/{id}`  
📌 **Пример запроса:**  
```
PUT http://localhost:8080/api/v1/todos/1a2b3c4d-5678-90ab-cdef-1234567890ab
```
📌 **Body (JSON):**
```json
{
  "title": "Купить продукты на неделю",
  "description": "Добавить овощи и мясо",
  "completed": true
}
```
📌 **Ответ (200 OK):**  
```json
{
  "id": "1a2b3c4d-5678-90ab-cdef-1234567890ab",
  "title": "Купить продукты на неделю",
  "description": "Добавить овощи и мясо",
  "completed": true,
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "created_at": "2024-03-01T12:00:00Z",
  "updated_at": "2024-03-01T14:00:00Z"
}
```
📌 **Ошибки:**  
- `400 Bad Request` → если ID неверный  
- `404 Not Found` → если задача не найдена  

---

## **5️⃣ Удалить задачу (`Delete Todo`)**  
📌 **Метод:** `DELETE http://localhost:8080/api/v1/todos/{id}`  
📌 **Пример запроса:**  
```
DELETE http://localhost:8080/api/v1/todos/1a2b3c4d-5678-90ab-cdef-1234567890ab
```
📌 **Ответ (204 No Content):**  
```
(пустой ответ)
```
📌 **Ошибки:**  
- `400 Bad Request` → если ID неверный  
- `404 Not Found` → если задачи нет  

---

## **Как импортировать в Postman?**
1. **Откройте Postman**  
2. **Создайте новую коллекцию** (`Todo API`)  
3. **Добавьте запросы**:
   - `POST /api/v1/todos` → Создание задачи  
   - `GET /api/v1/todos` → Получение всех задач  
   - `GET /api/v1/todos/{id}` → Получение одной задачи  
   - `PUT /api/v1/todos/{id}` → Обновление задачи  
   - `DELETE /api/v1/todos/{id}` → Удаление задачи  

