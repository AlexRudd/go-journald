package mock

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const journalSeed = 666
const machineId = "testtesttesttesttest"
const host = "test.mock"

type Entry map[string]string

func JournalGatewayd(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/browse", MockJournaldBrowse)
	mux.HandleFunc("/entries", MockJournaldEntries)
	mux.HandleFunc("/machine", MockJournaldMachine)
	mux.HandleFunc("/fields/", MockJournaldFields)

	mux.ServeHTTP(w, r)
}

func getJournalEntries(count, offset int) []Entry {
	r := rand.New(rand.NewSource(journalSeed))
	entries := []Entry{}
	clock := time.Date(2018, time.June, 19, 20, 36, 43, 0, time.Local).Unix() * int64(time.Microsecond)
	timeOffset := int64(0)
	for idx := 0; idx < offset+count; idx++ {
		entry := Entry{
			// User Journal Fields
			"MESSAGE":  fmt.Sprint("Test message #", r.Int()),
			"PRIORITY": fmt.Sprint(r.Intn(7)),
			// Custom User Field
			"CUSTOM_FIELD": fmt.Sprint(r.Intn(3)),
			// Trusted Fields
			"_PID":        fmt.Sprint(r.Int31n(50)),
			"_UID":        fmt.Sprint(r.Int31n(10)),
			"_GID":        fmt.Sprint(r.Int31n(10)),
			"_CMDLINE":    fmt.Sprint("/usr/bin/", r.Int31n(50)),
			"_MACHINE_ID": machineId,
			"_HOSTNAME":   host,
			"_TRANSPORT":  "journal",
			// Address Fields
			"__CURSOR":              fmt.Sprint(idx),
			"__REALTIME_TIMESTAMP":  fmt.Sprint(clock + timeOffset),
			"__MONOTONIC_TIMESTAMP": fmt.Sprint(timeOffset),
		}
		if idx >= offset {
			entries = append(entries, entry)
		}
		timeOffset += int64(r.Intn(100000))
	}
	return entries
}

func serialiseAs(acceptHeader string, v interface{}) (string, error) {
	switch acceptHeader {
	case "application/json":
		b, err := json.Marshal(v)
		return string(b), err
	default:
		return "", fmt.Errorf("Unrecognised Accept header: %s", acceptHeader)
	}
}

func MockJournaldBrowse(w http.ResponseWriter, r *http.Request) {
}

func MockJournaldEntries(w http.ResponseWriter, r *http.Request) {
}

func MockJournaldMachine(w http.ResponseWriter, r *http.Request) {
	machine := map[string]string{
		"machine_id":     machineId,
		"boot_id":        "3d3c9efaf556496a9b04259ee35df7f7",
		"hostname":       host,
		"os_pretty_name": "testy",
		"virtualization": "kvm",
	}

	resp, err := serialiseAs(r.Header.Get("Accept"), machine)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintln(w, resp)
}

func MockJournaldFields(w http.ResponseWriter, r *http.Request) {
}
