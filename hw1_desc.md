# Задание 1
- Простейшая авторизация пользователя.
- Возможность создания пользователя, где указывается следующая информация:
1. Имя
2. Фамилия
3. Возраст
4. Пол
5. Интересы
6. Город
- Страницы с анкетой.
## Нефункциональные требования:
- Любой язык программирования
- В качестве базы данных использовать MySQL (при остром желании PostgreSQL/MariaDB, но стоит иметь ввиду, что по этим базам может не быть возможности задать вопрос преподавателю)
- Не использовать ORM
- Программа должна представлять из себя монолитное приложение.
- Не рекомендуется использовать следующие технологии:
1. Репликация
2. Шардинг
3. Индексы
4. Кэширование
Для удобства разработки и проверки задания можно воспользоваться этой спецификацией и реализовать в ней методы:
- /login
- /user/register
- /user/get/{id}
Фронт опционален.
Сделать инструкцию по локальному запуску приложения, приложить Postman-коллекцию.
ДЗ принимается в виде исходного кода на github и Postman-коллекции.


Критерии оценки:
Оценка происходит по принципу зачет/незачет.
Требования:
Есть возможность авторизации, регистрации, получение анкет по ID.
Отсутствуют SQL-инъекции.
Пароль хранится безопасно.
