package GoEnv

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	// テスト用の.envファイルを作成
	fileContent := `KEY1="value1"
KEY2='value2'
KEY3=` + "`value3`" + `
KEY4=value4
KEY5="value with spaces"
KEY6='value with "quotes" and ''single quotes'''
`
	file, err := os.CreateTemp("", ".env")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	if _, err := file.WriteString(fileContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	// LoadEnv関数をテスト
	env, err := LoadEnv(file.Name())
	if err != nil {
		t.Fatalf("LoadEnv returned an error: %v", err)
	}

	// 結果を検証
	expected := map[string]string{
		"KEY1": "value1",
		"KEY2": "value2",
		"KEY3": "value3",
		"KEY4": "value4",
		"KEY5": "value with spaces",
		"KEY6": `value with "quotes" and 'single quotes'`,
	}

	for key, expectedValue := range expected {
		if value, exists := env[key]; !exists || value != expectedValue {
			t.Errorf("Expected %s=%s, but got %s=%s", key, expectedValue, key, value)
		}
	}

	if _, exists := env["INVALID_LINE"]; exists {
		t.Errorf("Expected INVALID_LINE to be ignored, but it was found in the env map")
	}
}
