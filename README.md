# Реализованный функционал

Был реализован базовый функционал бота:

- Выдача пользователю списка вопросов
- Проверка ответа на них
- Выдача таблицы лидеров
- Выдача положение пользователя в текущем рейтинге
- Засечение таймера на определенное время. В нашем случае 60 секунд


Хотелось бы добавить:

- Обработчик вопросов. Заметил при проверке работоспособности бота, что не все вопросы подходят для выдачи. Так как в
в некоторых вопросы предназначены для проведения блица, в ряде вопросов есть сноски как их следует читать.
- Также заметил, что не у всех пользователей в телеграмме есть username. И поэтому было бы неплохо добавить свой сервис регистрации
- Также хотелось бы добавить при правильном ответе на вопрос всплывающую реакцию ✅, а при неправильном ❌
- Еще было бы неплохо добавить режим игры как в реальном Что? Где? Когда?. То есть возможность пользователем создать 
команду из 6 человек. И провести раунд до 6 очков
- Проверка правильности орфографии слова. Было бы неплохо засчитывать правильный ответ если пользователь ошибся в нескольких буквах
Но ответ все равно верный. Например, слово "превет" сервис орфографии должен будет преобразовать в слово "привет"
- Сервис статистики ответа на вопрос. Проверка как пользователи ответили на каждый из вопросов.
Например пользователь закончил игру и хочет узнать сколько правильно ответило пользователей на вопросы, которые у него были