package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
)

type Tank struct {
    ID                int    json:"id"
    Year              int    json:"year"
    Caliber           int    json:"caliber"
    Weight            int    json:"weight"
    SubcaliberAmmo    bool   json:"subcaliber_ammo"
    CrewMembersNumber int    json:"crew_members_number"
}

func main() {
    // ... (прочий код module1)

    http.HandleFunc("/add", func(w http.ResponseWriter, r http.Request) {
        var tank Tank
        err := json.NewDecoder(r.Body).Decode(&tank)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        _, err = db.Exec("INSERT INTO tanks (year, caliber, weight, subcaliber_ammo, crew_members_number) VALUES (?, ?, ?, ?, ?)",
            tank.Year, tank.Caliber, tank.Weight, tank.SubcaliberAmmo, tank.CrewMembersNumber)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "Tank added successfully!")
    })

    http.HandleFunc("/list", func(w http.ResponseWriter, rhttp.Request) {
        rows, err := db.Query("SELECT id, year, caliber, weight, subcaliber_ammo, crew_members_number FROM tanks")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var tanks []Tank
        for rows.Next() {
            var tank Tank
            err := rows.Scan(&tank.ID, &tank.Year, &tank.Caliber, &tank.Weight, &tank.SubcaliberAmmo, &tank.CrewMembersNumber)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            tanks = append(tanks, tank)
        }

        json.NewEncoder(w).Encode(tanks)
    })

    // ... (прочий код module1)
}