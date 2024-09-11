// Dient der Pruefung von Sortier-Algorithmen, die auf Slices von
// operierenn. Enthaelt folgende Funktionen:
//
// Generate - Generiert Zufallswerte (Anzahl wird durch das aufrufende
//
//	Programm bestimmt)
//
// Check    - Prueft einen uebergebenen Slice, ob er a) sortier ist und b)
//
//	ob dieselben Zahlen enthaelt wie bei Generate erzeugt.
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

var (
	bck []float64
)

// Initialisiert alle verfuegbaren Plaetze in s mit Fliesskomma-Zufallszahlen
// aus dem Intervall [0,1). Der Slice wird ausserdem kopiert, diese Kopie
// mit Hilfe von sort sortiert und als globale Variable bck gespeichert.
func Generate(s []float64) {
	n := len(s)
	for i := range s {
		s[i] = rand.Float64()
	}
	bck = make([]float64, n)
	copy(bck, s)
	sort.Float64s(bck)
}

// Vergleicht den Slice s mit dem Slice in bck. Falls sie gleich sind, d.h.
// Element fuer Element die gleichen Zahlen enthalten, dann ist s mit grosser
// Wahrscheinlichkeit die Grundlage fuer bck gewesen, d.h. die Daten wurden
// effektiv eigenhaendig, resp. eigen-programmig sortiert.
func Check(s []float64) error {
	if len(s) != len(bck) {
		return errors.New(fmt.Sprintf("Zu pruefender Slice ist '%d' Felder gross, der generierte aber '%d'", len(s), len(bck)))
	}
	for i := range s {
		if s[i] != bck[i] {
			errors.New(fmt.Sprintf("Stimmt an der Stelle '%d' nicht: habe %.3f, erwarte %.3f", i, s[i], bck[i]))
		}
	}
	return nil
}
