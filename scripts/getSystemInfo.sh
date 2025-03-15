#!/bin/bash

# Функция для получения информации о дисковом пространстве
get_disk_info() {
    df -h --output=source,size,avail,target | awk '
    NR==1 {next}
    {
        isRoot = ($4 == "/") ? "true" : "false"; # Проверяем, является ли раздел корневым
        printf "{\"filesystem\": \"%s\", \"totalSize\": \"%s\", \"available\": \"%s\", \"isRoot\": %s}", $1, $2, $3, isRoot
        if (NR < FNR) printf "," # Добавляем запятую между объектами, кроме последнего
        printf "\n"
    }' | awk 'BEGIN {print "["} {print} END {print "]"}'
}

# Функция для получения информации об ОЗУ
get_ram_info() {
    awk '
    /MemTotal/ {total=$2}
    /MemAvailable/ {available=$2}
    END {
        used=total-available;
        printf "{\"totalRam\": \"%.2fG\", \"usedRam\": \"%.2fG\"}", total/1024/1024, used/1024/1024
    }' /proc/meminfo
}

# Функция для получения загрузки CPU
get_cpu_load() {
    # Функция для чтения данных из /proc/stat
    read_cpu_stats() {
        awk '/^cpu / {
            # Используем массив для хранения значений
            split($0, cpu_fields, " ");
            idle=cpu_fields[5];
            total=0;
            for (i=2; i<=NF; i++) total+=cpu_fields[i];
            printf "%d %d", idle, total
        }' /proc/stat
    }

    # Первое измерение
    read -r idle1 total1 <<< $(read_cpu_stats)

    # Ждем 1 секунду
    sleep 1

    # Второе измерение
    read -r idle2 total2 <<< $(read_cpu_stats)

    # Вычисляем разницу
    idle_diff=$((idle2 - idle1))
    total_diff=$((total2 - total1))

    # Вычисляем загрузку CPU
    if [[ $total_diff -ne 0 ]]; then
        cpu_load=$(awk -v idle="$idle_diff" -v total="$total_diff" 'BEGIN {printf "%.2f", 100 * (1 - idle / total)}')
    else
        cpu_load="0.00" # Если total_diff == 0, загрузка CPU равна 0%
    fi

    # Возвращаем результат в JSON
    echo "{\"cpuLoad\": \"$cpu_load%\"}"
}

# Формируем итоговый JSON
echo "{ \"diskInfo\": $(get_disk_info), \"ramInfo\": $(get_ram_info), \"cpuInfo\": $(get_cpu_load) }"
