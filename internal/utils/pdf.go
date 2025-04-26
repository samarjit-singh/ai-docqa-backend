package utils

import (
	"bytes"
	"fmt"

	"github.com/ledongthuc/pdf"
)

func ExtractPDFContent(fileBytes []byte) (string, error) {
	reader := bytes.NewReader(fileBytes)

	r, err := pdf.NewReader(reader, int64(len(fileBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to create PDF reader: %w", err)
	}

	var buf bytes.Buffer
	numPages := r.NumPage()

	for i := 1; i <= numPages; i++ {
		p := r.Page(i)
		if p.V.IsNull() {
			continue
		}

		text, err := p.GetPlainText(nil)
		if err != nil {
			return "", fmt.Errorf("failed to extract text from page %d: %w", i, err)
		}

		buf.WriteString(text)
		buf.WriteString("\n")
	}

	return buf.String(), nil
}
