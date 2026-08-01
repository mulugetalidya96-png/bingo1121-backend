package sms

import (
	"strings"
	"testing"
)

func TestParseTelebirrSMS_Amharic(t *testing.T) {
	// Test the Amharic SMS sample you provided
	amharicSMS := `ውድTIBEBU ወደ Frezer Wudneh(2519****3937) 20.00 ብር በ 01/08/2026 17:27:46 ልከዋል። የሂሳብ እንቅስቃሴ ቁጥርዎ DH18FW1Y4C ነዉ። የአገልግሎት ክፍያው 0.87 ብር ፤ የአገልግሎት ክፍያው 15% VAT 0.13 ብር ነዉ። አሁን ያለዎት ቀሪ ሂሳብ 2.00 ብር ነው። የክፍያ መረጃዎን ለማግኘት ማስፈንጠርያውን ይጫኑ፡ https://transactioninfo.ethiotelecom.et/receipt/DH18FW1Y4C በቴሌብር ስለተገለገሉ እናመሰግናለን ኢትዮ ቴሌኮም`

	result := ParseTelebirrSMS(amharicSMS)

	if !result.IsValid {
		t.Error("Expected valid transaction, got invalid")
	}

	// Test Amount
	expectedAmount := 20.00
	if result.Amount != expectedAmount {
		t.Errorf("Expected amount %.2f, got %.2f", expectedAmount, result.Amount)
	}

	// Test Transaction ID
	expectedTxnID := "DH18FW1Y4C"
	if result.TransactionID != expectedTxnID {
		t.Errorf("Expected transaction ID %s, got %s", expectedTxnID, result.TransactionID)
	}

	// Test Receiver
	expectedReceiver := "Frezer Wudneh"
	if result.Receiver != expectedReceiver {
		t.Errorf("Expected receiver %s, got %s", expectedReceiver, result.Receiver)
	}

	// Test Sender
	expectedSender := "TIBEBU"
	if result.Sender != expectedSender {
		t.Errorf("Expected sender %s, got %s", expectedSender, result.Sender)
	}

	// Test Balance
	expectedBalance := 2.00
	if result.Balance != expectedBalance {
		t.Errorf("Expected balance %.2f, got %.2f", expectedBalance, result.Balance)
	}

	// Test Date
	expectedDate := "01/08/2026 17:27:46"
	if result.Date != expectedDate {
		t.Errorf("Expected date %s, got %s", expectedDate, result.Date)
	}
}

func TestParseTelebirrSMS_English(t *testing.T) {
	// Create an English Telebirr SMS sample
	englishSMS := `Dear John Doe, you have transferred ETB 150.00 to Jane Smith(2519****1234) on 01/08/2026 15:30:00. Your transaction number is TX123ABC456. Your balance is ETB 250.00. Service fee 1.50 ETB, VAT 0.23 ETB. Thank you for using Telebirr. Ethio Telecom`

	result := ParseTelebirrSMS(englishSMS)

	if !result.IsValid {
		t.Error("Expected valid transaction, got invalid")
	}

	// Test Amount
	expectedAmount := 150.00
	if result.Amount != expectedAmount {
		t.Errorf("Expected amount %.2f, got %.2f", expectedAmount, result.Amount)
	}

	// Test Transaction ID
	expectedTxnID := "TX123ABC456"
	if result.TransactionID != expectedTxnID {
		t.Errorf("Expected transaction ID %s, got %s", expectedTxnID, result.TransactionID)
	}

	// Test Receiver
	expectedReceiver := "Jane Smith"
	if result.Receiver != expectedReceiver {
		t.Errorf("Expected receiver %s, got %s", expectedReceiver, result.Receiver)
	}

	// Test Sender
	expectedSender := "John Doe"
	if result.Sender != expectedSender {
		t.Errorf("Expected sender %s, got %s", expectedSender, result.Sender)
	}

	// Test Balance
	expectedBalance := 250.00
	if result.Balance != expectedBalance {
		t.Errorf("Expected balance %.2f, got %.2f", expectedBalance, result.Balance)
	}

	// Test Date
	expectedDate := "01/08/2026 15:30:00"
	if result.Date != expectedDate {
		t.Errorf("Expected date %s, got %s", expectedDate, result.Date)
	}
}

func TestParseTelebirrSMS_Amharic_AlternativeFormat(t *testing.T) {
	// Test another Amharic format variation
	amharicSMS := `ውድ MULUGETA ወደ Abebe Kebede(2519****5678) 45.50 ብር በ 02/08/2026 09:15:30 ልከዋል። የሂሳብ እንቅስቃሴ ቁጥርዎ AB12CD34EF ነው። አሁን ያለዎት ቀሪ ሂሳብ 100.00 ብር ነው። በቴሌብር ስለተገለገሉ እናመሰግናለን ኢትዮ ቴሌኮም`

	result := ParseTelebirrSMS(amharicSMS)

	if !result.IsValid {
		t.Error("Expected valid transaction, got invalid")
	}

	// Test Amount
	expectedAmount := 45.50
	if result.Amount != expectedAmount {
		t.Errorf("Expected amount %.2f, got %.2f", expectedAmount, result.Amount)
	}

	// Test Transaction ID
	expectedTxnID := "AB12CD34EF"
	if result.TransactionID != expectedTxnID {
		t.Errorf("Expected transaction ID %s, got %s", expectedTxnID, result.TransactionID)
	}

	// Test Receiver
	expectedReceiver := "Abebe Kebede"
	if result.Receiver != expectedReceiver {
		t.Errorf("Expected receiver %s, got %s", expectedReceiver, result.Receiver)
	}

	// Test Sender
	expectedSender := "MULUGETA"
	if result.Sender != expectedSender {
		t.Errorf("Expected sender %s, got %s", expectedSender, result.Sender)
	}

	// Test Balance
	expectedBalance := 100.00
	if result.Balance != expectedBalance {
		t.Errorf("Expected balance %.2f, got %.2f", expectedBalance, result.Balance)
	}
}

func TestParseTelebirrSMS_Invalid(t *testing.T) {
	// Test invalid SMS (not a Telebirr message)
	invalidSMS := `This is just a regular SMS message with no transaction info`

	result := ParseTelebirrSMS(invalidSMS)

	if result.IsValid {
		t.Error("Expected invalid transaction, got valid")
	}

	// Check that all fields are empty/zero
	if result.Amount != 0 {
		t.Errorf("Expected amount 0, got %.2f", result.Amount)
	}
	if result.TransactionID != "" {
		t.Errorf("Expected empty transaction ID, got %s", result.TransactionID)
	}
	if result.Balance != 0 {
		t.Errorf("Expected balance 0, got %.2f", result.Balance)
	}
}

func TestParseTelebirrSMS_MissingFields(t *testing.T) {
	// Test SMS with missing some fields
	sms := `ቁጥርዎ XYZ123 ነው። 50.00 ብር ልከዋል።`

	result := ParseTelebirrSMS(sms)

	// Should still be valid if we have transaction ID and amount
	if !result.IsValid {
		t.Error("Expected valid transaction even with missing fields")
	}

	if result.Amount != 50.00 {
		t.Errorf("Expected amount 50.00, got %.2f", result.Amount)
	}
	if result.TransactionID != "XYZ123" {
		t.Errorf("Expected transaction ID XYZ123, got %s", result.TransactionID)
	}
}

func TestIsTelebirrSMS(t *testing.T) {
	testCases := []struct {
		name     string
		sms      string
		expected bool
	}{
		{
			name:     "Amharic SMS",
			sms:      `ውድTIBEBU ወደ Frezer Wudneh(2519****3937) 20.00 ብር ልከዋል። ቁጥርዎ DH18FW1Y4C ነው። በቴሌብር ስለተገለገሉ እናመሰግናለን ኢትዮ ቴሌኮም`,
			expected: true,
		},
		{
			name:     "English SMS",
			sms:      `Dear John, you have transferred ETB 100.00 to Jane. Transaction number is TX123. Thank you for using Telebirr.`,
			expected: true,
		},
		{
			name:     "Regular SMS",
			sms:      `Hello, how are you? This is a regular message.`,
			expected: false,
		},
		{
			name:     "SMS with Amharic keyword",
			sms:      `አሁን ያለዎት ቀሪ ሂሳብ 100.00 ብር ነው።`,
			expected: true,
		},
		{
			name:     "SMS with Ethio Telecom mention",
			sms:      `እንኳን ደስ አለዎት ኢትዮ ቴሌኮም ለደንበኞች አዲስ አገልግሎት አስተዋውቋል`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsTelebirrSMS(tc.sms)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestContainsAmharic(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		expected bool
	}{
		{
			name:     "Pure Amharic text",
			text:     `ውድ ደንበኛ ሰላም ነው`,
			expected: true,
		},
		{
			name:     "Mixed Amharic and English",
			text:     `ውድ TIBEBU ወደ Frezer Wudneh`,
			expected: true,
		},
		{
			name:     "Pure English text",
			text:     `Dear customer, hello`,
			expected: false,
		},
		{
			name:     "Numbers and special characters",
			text:     `123 456 7890`,
			expected: false,
		},
		{
			name:     "Amharic numbers",
			text:     `፩ ፪ ፫`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := containsAmharic(tc.text)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFormatTransactionSummary(t *testing.T) {
	info := &TransactionInfo{
		Amount:        20.00,
		TransactionID: "DH18FW1Y4C",
		Sender:        "TIBEBU",
		Receiver:      "Frezer Wudneh",
		Date:          "01/08/2026 17:27:46",
		Balance:       2.00,
		IsValid:       true,
	}

	summary := FormatTransactionSummary(info)

	// Check that the summary contains expected information
	expectedParts := []string{
		"✅ Transaction Details:",
		"💰 Amount: 20.00 ETB",
		"🆔 Transaction ID: DH18FW1Y4C",
		"👤 From: TIBEBU",
		"👤 To: Frezer Wudneh",
		"📅 Date: 01/08/2026 17:27:46",
		"💰 Balance: 2.00 ETB",
	}

	for _, part := range expectedParts {
		if !strings.Contains(summary, part) {
			t.Errorf("Expected summary to contain '%s', got:\n%s", part, summary)
		}
	}

	// Test invalid transaction
	invalidInfo := &TransactionInfo{IsValid: false}
	invalidSummary := FormatTransactionSummary(invalidInfo)
	expected := "Invalid transaction"
	if invalidSummary != expected {
		t.Errorf("Expected '%s', got '%s'", expected, invalidSummary)
	}
}

// Benchmark tests
func BenchmarkParseTelebirrSMS_Amharic(b *testing.B) {
	sms := `ውድTIBEBU ወደ Frezer Wudneh(2519****3937) 20.00 ብር በ 01/08/2026 17:27:46 ልከዋል። የሂሳብ እንቅስቃሴ ቁጥርዎ DH18FW1Y4C ነዉ። አሁን ያለዎት ቀሪ ሂሳብ 2.00 ብር ነው።`

	for i := 0; i < b.N; i++ {
		ParseTelebirrSMS(sms)
	}
}

func BenchmarkParseTelebirrSMS_English(b *testing.B) {
	sms := `Dear John Doe, you have transferred ETB 150.00 to Jane Smith(2519****1234) on 01/08/2026 15:30:00. Your transaction number is TX123ABC456. Your balance is ETB 250.00.`

	for i := 0; i < b.N; i++ {
		ParseTelebirrSMS(sms)
	}
}

func BenchmarkIsTelebirrSMS(b *testing.B) {
	sms := `ውድTIBEBU ወደ Frezer Wudneh(2519****3937) 20.00 ብር ልከዋል።`

	for i := 0; i < b.N; i++ {
		IsTelebirrSMS(sms)
	}
}