package controllers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type EntropiRequest struct {
	Text string `json:"text"`
}

type EntropiResponse struct {
	Entropi       float64         `json:"entropi"`
	Frekuensi     map[string]int  `json:"frekuensi"`
	TotalKarakter int             `json:"total_karakter"`
}

func HitungEntropiText(c *gin.Context) {
	var req EntropiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	text := strings.ToLower(req.Text)
	hitungDanBalas(c, text)
}

func HitungEntropiCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File CSV tidak ditemukan"})
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka file"})
		return
	}
	defer src.Close()

	reader := csv.NewReader(src)
	var allText string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membaca CSV"})
			return
		}
		allText += strings.Join(record, "")
	}
	hitungDanBalas(c, strings.ToLower(allText))
}

func HitungEntropiImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File gambar tidak ditemukan"})
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka gambar"})
		return
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format gambar tidak dikenali"})
		return
	}

	var data []byte
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			data = append(data, gray.Y)
		}
	}

	// Konversi ke string untuk entropi
	text := string(data)
	hitungDanBalas(c, text)
}

func hitungDanBalas(c *gin.Context, text string) {
	freq := make(map[rune]float64)
	freqInt := make(map[string]int)
	total := float64(len(text))

	for _, char := range text {
		freq[char]++
		freqInt[string(char)]++
	}

	var entropy float64
	for _, count := range freq {
		p := count / total
		entropy += -p * math.Log2(p)
	}

	c.JSON(http.StatusOK, EntropiResponse{
		Entropi:       entropy,
		Frekuensi:     freqInt,
		TotalKarakter: len(text),
	})
}

func GeneratePDF(c *gin.Context) {
	// Dummy data atau ambil dari cache
	resp := EntropiResponse{
		Entropi:       3.3219,
		TotalKarakter: 100,
		Frekuensi: map[string]int{
			"a": 10, "b": 15, "c": 5, "d": 20,
		},
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Laporan Entropi")

	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Total Karakter: %d", resp.TotalKarakter))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Entropi: %.4f", resp.Entropi))
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Frekuensi Karakter:")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	for k, v := range resp.Frekuensi {
		pdf.Cell(40, 10, fmt.Sprintf("%s : %d", k, v))
		pdf.Ln(7)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate PDF"})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=laporan_entropi.pdf")
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}
