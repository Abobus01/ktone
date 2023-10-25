package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/nats-io/stan.go"
)

type Tank struct {
    Year              int  json:"year"
    Caliber           int  json:"caliber"
    Weight            int  json:"weight"
    SubcaliberAmmo    bool json:"subcaliber_ammo"
    CrewMembersNumber int  json:"crew_members_number"
}

func main() {
    // ... (прочий код module2)

    // Создаем новый танк
    newTank := Tank{
        Year:              2023,
        Caliber:           120,
        Weight:            55000,
        SubcaliberAmmo:    true,
        CrewMembersNumber: 4,
    }

    // Преобразуем танк в JSON
    tankJSON, err := json.Marshal(newTank)
    if err != nil {
        log.Fatal(err)
    }

    // Отправляем POST-запрос для добавления танка
    resp, err := http.Post("http://localhost:8080/add", "application/json", bytes.NewBuffer(tankJSON))
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)

    // Ждем немного перед отправкой GET-запроса для получения списка танков
    time.Sleep(time.Second)

    // Отправляем GET-запрос для получения списка танков
    resp, err = http.Get("http://localhost:8080/list")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    var tanks []Tank
    err = json.NewDecoder(resp.Body).Decode(&tanks)
    if err != nil {
        log.Fatal(err)
    }

    // Выводим список танков
    fmt.Println("List of Tanks:")
    for _, tank := range tanks {
        fmt.Printf("Year: %d, Caliber: %d mm, Weight: %d kg, Subcaliber Ammo: %t, Crew Members: %d\n",
            tank.Year, tank.Caliber, tank.Weight, tank.SubcaliberAmmo, tank.CrewMembersNumber)
    }

    // ... (прочий код module2)
}