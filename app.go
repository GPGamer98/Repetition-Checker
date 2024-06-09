package main

import (
	"context"
	"bufio"
	"os"
	"path/filepath"
	"encoding/csv"
	"sort"
	"strings"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"Cogito-Ergo-Vet/internal/update"
)

// App struct
type App struct {
	ctx context.Context
}

type Entry string

type CSVFile struct {
	Name     string // Name of the CSV file.
	Entries  []Entry // Unique entries found in the CSV file.
}

type EntryCount struct {
	Entry Entry
	Count int
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	runtime.LogInfo(a.ctx, "Checking Updates...")
	a.UpdateCheckUI()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SelectFile() (string, error) {
		file, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{Title: "Seleziona il file CSV", Filters: []runtime.FileFilter{{DisplayName: "File CSV (*.csv)", Pattern: "*.csv"}}})
		return file, err
}

func (a *App) AboutWindow() {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{Title: "Informazioni"})
}

func (a *App) AnalyzeCSVFiles(directory string) ([]EntryCount, error) {
	files, err := filepath.Glob(filepath.Join(directory, "*.csv"))
	if err!= nil {
		fmt.Printf("Error listing files: %v\n", err)
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	fmt.Println("Found CSV files:", files)

	countMap := make(map[string]int)
	for _, file := range files {
		f, err := os.Open(file)
		if err!= nil {
			fmt.Printf("Error opening file %s: %v\n", file, err)
			continue // Skip this file and move to the next one
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			normalizedLine := strings.TrimSpace(strings.ToLower(line)) // Normalize entry
			countMap[normalizedLine]++ // Increment count for each normalized entry
			fmt.Printf("Added entry '%s' to countMap\n", normalizedLine)
		}
		if err := scanner.Err(); err!= nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
		}
	}

	fmt.Println("Final countMap:", countMap)

	// Convert map to a slice of EntryCount for sorting
	var counts []EntryCount
	for entry, count := range countMap {
		counts = append(counts, EntryCount{Entry: Entry(entry), Count: count})
	}

	// Sort the slice by count (descending) for demonstration purposes
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count // Change to < for ascending order
	})

	fmt.Println("Sorted counts:", counts)

	return counts, nil
}

func (a *App) ExportDataToCSV(data []EntryCount, filename string) error {
	file, err := os.Create(filename)
	if err!= nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Comma = ','

	headers := []string{"Email", "N° volte"}
	err = writer.Write(headers)
	if err!= nil {
		return fmt.Errorf("failed to write headers: %w", err)
	}

	for _, entry := range data {
		fmt.Printf("Row: %s,%d\n", entry.Entry, entry.Count)
		row := []string{string(entry.Entry), fmt.Sprintf("%d", entry.Count)}
		err = writer.Write(row)
		if err!= nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}	

	fmt.Printf("Data exported to %s successfully.\n", filename)
	return nil
}

func (a *App) SelectExport() (string, error) {
	file, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{Title: "Seleziona il file CSV per esportare", DefaultFilename: "data.csv", Filters: []runtime.FileFilter{{DisplayName: "File CSV (*.csv)", Pattern: "*.csv"}}})
	return file, err
}

func (a *App) UpdateCheckUI() {
	shouldUpdate, latestVersion, currentVersion := update.CheckUpdate()
	if shouldUpdate {
		updateMessage := fmt.Sprintf("Nuova versione disponibile\nVorresti aggiornare da v%s a v%s?", currentVersion, latestVersion)
		buttons := []string{"Yes", "No"}
		dialogOpts := runtime.MessageDialogOptions{Title: "Aggiornamento Disponibile", Message: updateMessage, Type: runtime.QuestionDialog, Buttons: buttons, DefaultButton: "Sì", CancelButton: "No"}
		action, err := runtime.MessageDialog(a.ctx, dialogOpts)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// runtime.LogInfo(a.ctx, action)
		if action == "Yes" {
			// runtime.LogInfo(a.ctx, "Update clicked")
			updated := update.DoSelfUpdate()
			if updated {
				buttons = []string{"Ok"}
				dialogOpts = runtime.MessageDialogOptions{Title: "Aggiornamento Completato", Message: "Aggiornamento completato\nRiavvia l'app per applicare le modifiche", Type: runtime.InfoDialog, Buttons: buttons, DefaultButton: "Ok"}
				runtime.MessageDialog(a.ctx, dialogOpts)
			} else {
				buttons = []string{"Ok"}
				dialogOpts = runtime.MessageDialogOptions{Title: "Aggiornamento Fallito", Message: "Aggiornamento fallito\nAggiornare automaticamente da GitHub:\nhttps://github.com/GPGamer98/Repetition-Checker", Type: runtime.InfoDialog, Buttons: buttons, DefaultButton: "Ok"}
				runtime.MessageDialog(a.ctx, dialogOpts)
			}
		}
	}
}

func (a *App) GetCurrentVersion() string {
	return update.Version
}