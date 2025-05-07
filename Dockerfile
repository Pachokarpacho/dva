FROM ubuntu:latest

# Установим необходимые зависимости
RUN apt-get update && apt-get install -y dpkg

# Скопируем .deb пакет из артефакта
COPY myprogram.deb /tmp/

# Установим пакет
RUN dpkg -i /tmp/myprogram.deb

# Указываем точку входа для контейнера
ENTRYPOINT ["/usr/bun/lab"]
