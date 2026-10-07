package report

import (
	"encoding/json"
	"fmt"

	"github.com/Amezco-Group-LLC/sekd/internal/analysis"
)

func RenderJSON(r *analysis.Report) error {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
