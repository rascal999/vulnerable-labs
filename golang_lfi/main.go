package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

const baseDir = "/home/user/git/vulnerable-labs/golang_lfi/files"

func handleGetFile0(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Query().Get("file")
    if file == "" {
        http.Error(w, "missing 'file' parameter in request", http.StatusBadRequest)
        return
    }

    f, err := os.Open(file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer f.Close()

    _, err = io.Copy(w, f)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleGetFile1(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Query().Get("file")
    if file == "" {
        http.Error(w, "missing 'file' parameter in request", http.StatusBadRequest)
        return
    }

    if strings.HasPrefix(file, "/") {
        http.Error(w, "File path cannot start with /", http.StatusInternalServerError)
        return
    }

    f, err := os.Open(file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer f.Close()

    _, err = io.Copy(w, f)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleGetFile2(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Query().Get("file")
    if file == "" {
        http.Error(w, "missing 'file' parameter in request", http.StatusBadRequest)
        return
    }

    if strings.HasPrefix(file, "/") {
        http.Error(w, "File path cannot start with /", http.StatusInternalServerError)
        return
    }

    //file = filepath.Clean(file)
    file = strings.Replace(file, "../", "", -1)

    f, err := os.Open(file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer f.Close()

    _, err = io.Copy(w, f)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleGetFile3(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Query().Get("file")
    if file == "" {
        http.Error(w, "missing 'file' parameter in request", http.StatusBadRequest)
        return
    }

    absFile := filepath.Join(baseDir, file)
    absBase, err := filepath.Abs(baseDir)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if !strings.HasPrefix(absFile, absBase) {
        http.Error(w, "accessing files outside of base directory is not allowed", http.StatusForbidden)
        return
    }

    f, err := os.Open(absFile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer f.Close()

    _, err = io.Copy(w, f)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleGetFile4(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Query().Get("file")
    if file == "" {
        http.Error(w, "missing 'file' parameter in request", http.StatusBadRequest)
        return
    }

    ext := filepath.Ext(file)
    if ext != ".pdf" {
        http.Error(w, "only '.pdf' files are allowed", http.StatusBadRequest)
        return
    }

    absFile := filepath.Join(baseDir, file)
    absBase, err := filepath.Abs(baseDir)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if !strings.HasPrefix(absFile, absBase) {
        http.Error(w, "accessing files outside of base directory is not allowed", http.StatusForbidden)
        return
    }

    f, err := os.Open(absFile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer f.Close()

    _, err = io.Copy(w, f)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleGetFile5(w http.ResponseWriter, r *http.Request) {
    file := r.URL.Query().Get("file")
    if file == "" {
        http.Error(w, "missing 'file' parameter in request", http.StatusBadRequest)
        return
    }

    ext := filepath.Ext(file)
    if ext != ".pdf" {
        http.Error(w, "only '.pdf' files are allowed", http.StatusBadRequest)
        return
    }

    absFile := filepath.Join(baseDir, file)
    absBase, err := filepath.Abs(baseDir)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if !strings.HasPrefix(absFile, absBase) {
        http.Error(w, "accessing files outside of base directory is not allowed", http.StatusForbidden)
        return
    }

    f, err := os.Open(absFile)
    if err != nil {
        http.Error(w, "File not found", http.StatusInternalServerError)
        return
    }
    defer f.Close()

    _, err = io.Copy(w, f)
    if err != nil {
        http.Error(w, "Error", http.StatusInternalServerError)
    }
}

func main() {
    // Solution: Y3VybCAiaHR0cDovL2xvY2FsaG9zdDo4MDgwLzA/ZmlsZT0vZXRjL3Bhc3N3ZCI=
    http.HandleFunc("/0", handleGetFile0)
    // Solution: Y3VybCAiaHR0cDovL2xvY2FsaG9zdDo4MDgwLzE/ZmlsZT0uLi8uLi8uLi8uLi8uLi9ldGMvcGFzc3dkIg==
    http.HandleFunc("/1", handleGetFile1)
    // Solution: Y3VybCAiaHR0cDovL2xvY2FsaG9zdDo4MDgwLzI/ZmlsZT0uLi4vLi8uLi4vLi8uLi4vLi8uLi4vLi8uLi4vLi9ldGMvcGFzc3dkIgo=
    http.HandleFunc("/2", handleGetFile2)
    // Solution: Y3VybCAiaHR0cDovL2xvY2FsaG9zdDo4MDgwLzM/ZmlsZT1iYWQudHh0Ig==
    http.HandleFunc("/3", handleGetFile3)
    // Solution: Y3VybCAiaHR0cDovL2xvY2FsaG9zdDo4MDgwLzQ/ZmlsZT0vZXRjL3Bhc3N3ZC5wZGYi
    http.HandleFunc("/4", handleGetFile4)
    http.HandleFunc("/5", handleGetFile5)

    fmt.Println("Starting server on http://localhost:8000")
    http.ListenAndServe(":8080", nil)
}
