package api

import (
	"basel2053/ps-board/db"
	"basel2053/ps-board/ps"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterRecordHandlers(ctx context.Context, db *db.Postgres) {
	getRecords := func(w http.ResponseWriter, r *http.Request) {
		records, err := db.FindRecords(ctx, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, records)
	}

	getRecord := func(w http.ResponseWriter, r *http.Request) {
		recordId := r.PathValue("id")
		record, err := db.FindRecordById(ctx, recordId)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			http.Error(w, "Unable to get record "+recordId, http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, *record)
	}

	createRecord := func(w http.ResponseWriter, r *http.Request) {
		var record ps.Record
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&record)
		if err != nil {
			http.Error(w, "unable to parse request body", http.StatusBadRequest)
			return
		}
		err = db.CreateRecord(ctx, record)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			http.Error(w, "Unable to create record ", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Record created successfully!")
	}

	removeRecord := func(w http.ResponseWriter, r *http.Request) {
		recordId := r.PathValue("id")
		err := db.RemoveRecord(ctx, recordId)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			http.Error(w, "Unable to delete record ", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Record (%s) deleted successfully!", recordId)
	}

	http.HandleFunc("GET /records", getRecords)
	http.HandleFunc("GET /records/{id}", getRecord)
	http.HandleFunc("POST /records", createRecord)
	http.HandleFunc("DELETE /records/{id}", removeRecord)
}
