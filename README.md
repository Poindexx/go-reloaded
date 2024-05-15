# Text Completion/Editing/Auto-Correction Tool

## Описание проекта

Этот проект представляет собой инструмент для автозаполнения, редактирования и автокоррекции текста. Инструмент был создан мной и использует функции из предыдущих проектов, а также новые функции, написанные специально для данной задачи.

## Введение

Программа написана на языке Go и соответствует лучшим практикам программирования. Она включает тестовые файлы для модульного тестирования, чтобы обеспечить корректность работы и высокое качество кода.

### Функциональность

Инструмент принимает два аргумента: имя входного файла с текстом для модификации и имя выходного файла, в который будет сохранен измененный текст. Программа выполняет следующие модификации:

1. **(hex)**: Замена шестнадцатеричных чисел на десятичные.
2. **(bin)**: Замена двоичных чисел на десятичные.
3. **(up)**: Преобразование предыдущего слова в верхний регистр.
4. **(low)**: Преобразование предыдущего слова в нижний регистр.
5. **(cap)**: Преобразование предыдущего слова с заглавной буквы.
6. **(low, <number>)**, **(up, <number>)**, **(cap, <number>)**: Преобразование указанного количества предыдущих слов.
7. Обработка знаков препинания: `.`, `,`, `!`, `?`, `:` и `;`.
8. Обработка групп знаков препинания, таких как `...` и `!?`.
9. Корректное размещение одинарных кавычек `' '`.
10. Замена `a` на `an` перед гласными и `h`.

## Использование

```sh
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

$ go run . sample.txt result.txt

$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.

$ cat sample.txt
There is no greater agony than bearing a untold story inside you.

$ go run . sample.txt result.txt

$ cat result.txt
There is no greater agony than bearing an untold story inside you.

$ cat sample.txt
Punctuation tests are ... kinda boring ,don't you think !?

$ go run . sample.txt result.txt

$ cat result.txt
Punctuation tests are... kinda boring, don't you think!?
