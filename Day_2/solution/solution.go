package main
import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Funkcja do sprawdzania warunków dla tablicy różnic
func meetsConditions(differences []int) bool {
	if len(differences) == 0 {
		return false
	}

	// Sprawdź, czy wszystkie wartości są dodatnie lub ujemne
	allPositive := true
	allNegative := true
	for _, diff := range differences {
		if diff > 0 {
			allNegative = false
		} else if diff < 0 {
			allPositive = false
		}
	}

	// Jeśli nie są wszystkie dodatnie lub wszystkie ujemne, zwróć false
	if !(allPositive || allNegative) {
		return false
	}

	// Sprawdź, czy wszystkie wartości bezwzględne są > 1
	for _, diff := range differences {
		if (abs(diff) < 1 || abs(diff) > 3) {
			return false
		}
	}

	// Spełnia warunki
	return true
}

func solution() {
	// Wczytaj plik
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Błąd wczytywania pliku:", err)
		return
	}

	// Zamień dane na string
	content := string(data)

	// Podziel tekst na linie
	lines := strings.Split(content, "\n")

	// Przechowuj wszystkie linie jako tablicę tablic
	allIntElements := [][]int{} // Inicjalizacja pustej tablicy tablic

	for _, line := range lines {
		// Ignoruj puste linie
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Podziel linię na elementy
		elements := strings.Split(line, " ")

		// Przechowuj liczby całkowite
		intElements := []int{}

		for _, element := range elements {
			// Konwertuj każdy element na int
			num, err := strconv.Atoi(strings.TrimSpace(element))
			if err != nil {
				fmt.Println("Błąd konwersji:", err)
				continue
			}
			intElements = append(intElements, num)
		}

		// Dodaj intElements do allIntElements
		allIntElements = append(allIntElements, intElements)
	}

	// Oblicz różnice dla każdej linii
	// Oblicz różnice dla każdej linii
	result := 0
	for _, intElements := range allIntElements {
		// Oblicz różnice dla pierwotnej tablicy
		differences := calculateDifferences(intElements)

		// Sprawdź pierwotne warunki
		if meetsConditions(differences) {
			result++
			continue
		}

		// Jeśli warunki nie są spełnione, sprawdź wszystkie możliwe podzbiory
		isValid := false
		for i := 0; i < len(intElements); i++ {
			// Tworzenie nowego podzbioru bez elementu na pozycji `i`
			subset := make([]int, 0, len(intElements)-1) // Tworzymy nową tablicę
			subset = append(subset, intElements[:i]...)
			subset = append(subset, intElements[i+1:]...)

			// Obliczamy różnice dla tego podzbioru
			newDifferences := calculateDifferences(subset)

			// Debug print to verify subsets

			// Sprawdzenie warunków dla podzbioru
			if meetsConditions(newDifferences) {
				isValid = true
				break
			}
		}

		// Jeśli którykolwiek podzbiór spełnia warunki, zwiększ wynik
		if isValid {
			result++
		}
	}
	

	// Wyświetl wynik
	fmt.Println("Wynik:", result)
}

func calculateDifferences(elements []int) []int {
	differences := []int{}
	for i := 1; i < len(elements); i++ {
		differences = append(differences, elements[i]-elements[i-1])
	}
	return differences
}

func main() {
	solution()
}