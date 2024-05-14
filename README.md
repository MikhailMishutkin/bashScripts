# bashScripts

Проект находиться в ветке my.

Запускается make-файлом в корневой папке проекта.

Основная моя идея состоит в том, что bash-комманды это набор строк для интерпретатора, выполняемых одна за другой, а следовательно, нет смысла разбивать отдельные команды на составляющие. Встроенный пакет os/exec считывает команду как строку без дополнительных манипуляций. Поэтому, основная структура данных, которую я выбрал для использования в проекте это массив/слайс строк.

Основной транспорт между слоями - структуры, на вход и выход сериализуется в JSON. Тестовые запросы в формате JSON направлял несложные: простейший вывод и/или поиск строк в файлах проекта, пинг известных ресурсов. Осталось неясным зачем нужен результат выполнения запроса в базе, и я решил, что его надо выводить при запросе списка всех сохранённых скриптов, хотя это делает вывод списка некрасивым и трудно читаемым.

При проектировании базы данных я сделал уникальным значение имени скрипта, во избежание дублирования одинаковых имён. Конфликт обработал тектом ошибки в слое репозитория, не стал выводить на уровень сервиса, это на мой взгляд возможное сокращение для тестовой работы.

Накат миграций сделал через make-файл, хотя и знаком с миграциями внутри кода, однако такое решение у меня не всегда ранее корректно работало, поэтому склонился к надёжному варианту.

Логика названий в слоях строится по принципу - заходит сущность "скрипт", в сервисе разбивается на составные: имя, комманды, команда, результат, репозиторий выполняет простейшие операции записи/чтения, сложные запросы не видел смысла использовать, как и транзакции.

На представленную работу ушло 4 неполных дня, тестами, к сожалению, покрыть не успел.


