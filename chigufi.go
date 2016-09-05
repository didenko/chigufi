/*
Copyright 2016 Vlad Didenko

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License in the included LICENSE file or at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	COL_IUCR_ID      = 0
	COL_IUCR_FACRIME = 4
	COL_IUCR_VW      = 5
	COL_IUCR_MURDER  = 6

	COL_CRIME_DATE = 2
	COL_CRIME_IUCR = 4

	SHOW_BATCH = 100000
)

const (
	crime int = iota
	violation
	murder
)

var (
	crimes_name string
	iucr_name   string
)

type belongs struct {
	key   string
	categ int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	flag.StringVar(&iucr_name, "iucr", "", "Name of the incident codes CSV file")
	flag.StringVar(&crimes_name, "crimes", "", "Name of the crime records CSV file")
}

func load_iucrs(filename string) map[belongs]bool {
	iucr_csv, err := os.Open(filename)
	check(err)
	defer iucr_csv.Close()

	riucr := csv.NewReader(iucr_csv)
	iucrs := make(map[belongs]bool)

	_, _ = riucr.Read() // Skip the headers line

	for {
		iucr_rec, err := riucr.Read()
		if err == io.EOF {
			break
		}

		iucr := fmt.Sprintf("%04s", iucr_rec[COL_IUCR_ID])

		iucrs[belongs{iucr, crime}] = iucr_rec[COL_IUCR_FACRIME] != ""
		iucrs[belongs{iucr, violation}] = iucr_rec[COL_IUCR_VW] != ""
		iucrs[belongs{iucr, murder}] = iucr_rec[COL_IUCR_MURDER] != ""
	}

	return iucrs
}

func load_crimes(filename string, iucrs map[belongs]bool) (map[belongs]int, map[string]struct{}) {
	crimes_csv, err := os.Open(filename)
	check(err)
	defer crimes_csv.Close()

	rcsv := csv.NewReader(crimes_csv)

	stats := make(map[belongs]int)
	months_set := make(map[string]struct{})

	// Skip the headers line
	_, _ = rcsv.Read()

	line := 0
	next_report := SHOW_BATCH
	for {
		rec, err := rcsv.Read()
		if err == io.EOF {
			break
		}
		check(err)

		line++
		if line >= next_report {
			fmt.Fprintf(os.Stderr, "%d\r", line)
			next_report += SHOW_BATCH
		}

		date, err := time.Parse("01/02/2006 03:04:05 PM", rec[COL_CRIME_DATE])
		month := date.Format("2006-01")
		months_set[month] = struct{}{}

		if err != nil {
			fmt.Fprintln(os.Stderr)
			fmt.Fprintln(os.Stderr, "Line ", line, "error parsing the date:", rec[COL_CRIME_DATE])
			continue
		}

		if iucrs[belongs{rec[COL_CRIME_IUCR], crime}] {
			stats[belongs{month, crime}]++
		}

		if iucrs[belongs{rec[COL_CRIME_IUCR], violation}] {
			stats[belongs{month, violation}]++
		}

		if iucrs[belongs{rec[COL_CRIME_IUCR], murder}] {
			stats[belongs{month, murder}]++
		}
	}

	fmt.Fprintln(os.Stderr, line, "crime records processed")

	return stats, months_set
}

func sort_months(months_set map[string]struct{}) []string {
	months := make([]string, len(months_set))

	for month, _ := range months_set {
		months = append(months, month)
	}
	sort.Strings(months)

	return months
}

func print_stats(months []string, stats map[belongs]int) {
	out_csv := csv.NewWriter(os.Stdout)

	if err := out_csv.Write([]string{"Month", "FA_Crimes", "Violations", "Murders"}); err != nil {
		fmt.Fprintln(os.Stderr, "error writing record to csv:", err)
	}

	for _, month := range months {
		if month == "" {
			continue
		}

		crimes := strconv.Itoa(stats[belongs{month, crime}])
		violations := strconv.Itoa(stats[belongs{month, violation}])
		murders := strconv.Itoa(stats[belongs{month, murder}])

		if err := out_csv.Write([]string{month, crimes, violations, murders}); err != nil {
			fmt.Fprintln(os.Stderr, "error writing record to csv:", err)
		}
	}

	out_csv.Flush()
	if err := out_csv.Error(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}
}

func main() {
	flag.Parse()

	iucrs := load_iucrs(iucr_name)

	stats, months_set := load_crimes(crimes_name, iucrs)

	months := sort_months(months_set)

	print_stats(months, stats)
}
