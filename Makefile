# Makefile

.PHONY: all build test clean deb

# Сборка программы
build:
    @echo "Сборка программы..."
    go build -o lab ./cmd/lab

# Запуск тестов
test:
    @echo "Запуск тестов..."
    go test -v ./...

install:
    @echo "zav"
    go mod download

# Очистка проекта
clean:
    @echo "Очистка временных файлов..."
    rm -f lab

# Сборка .deb пакета
deb:
    @echo "Сборка .deb пакета..."
    mkdir -p myprogram/usr/bin
    mkdir -p myprogram/app
    cp lab myprogram/usr/bin/lab
    cp data/input.txt myprogram/app/input.txt  # Копируем input.txt
    mkdir -p myprogram/DEBIAN
    echo "Package: lab" > myprogram/DEBIAN/control
    echo "Version: 1.0" >> myprogram/DEBIAN/control
    echo "Architecture: amd64" >> myprogram/DEBIAN/control
    echo "Maintainer: Your Name <balashovpashabk@mail.ru>" >> myprogram/DEBIAN/control
    echo "Description: Программа для опознавания простого числа" >> myprogram/DEBIAN/control
    dpkg-deb --build myprogram
    rm -rf myprogram
