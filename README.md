## Настройка и запуск.

1. Скачайте архив по ссылке.
2. Переместите содержимое в любую папку.
3. Откройте командую строку и перейдите в эту папку.

### Пример:
cd C:\Users\\"ваш_юзер"\Desktop\\"ваша_папка"

4. В командной строке введите "go mod tidy" без кавычек.
5. Запустите сервер командой "go run ." без кавычек. Сервер работает на порте 8080.


## Для добавления, получения, удаления задач используется curl запросы.
## Для этого:
1. Откройте новую командную строку.
2. Введите одну из команд:

   1) curl -X POST http://localhost:8080/task - Создать задачу.

   2) curl http://localhost:8080/task/id - Получение задачи.
      id который вы получите после ввода 1 команды вставлять вместе с кавычками.

   3) curl http://localhost:8080/tasks - Получение всех задач.

   4) curl -X DELETE http://localhost:8080/task/id - Удаление задачи.