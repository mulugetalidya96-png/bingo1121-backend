package sms

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// TransactionInfo represents parsed transaction details
type TransactionInfo struct {
	Amount          float64
	TransactionID   string
	Sender          string
	Receiver        string
	Date            string
	Balance         float64
	IsValid         bool
}

// ParseTelebirrSMS parses Telebirr SMS messages (supports both English and Amharic)
func ParseTelebirrSMS(smsText string) *TransactionInfo {
	result := &TransactionInfo{
		IsValid: false,
	}

	// Check if SMS is in Amharic
	isAmharic := containsAmharic(smsText)

	if isAmharic {
		parseAmharicTelebirrSMS(smsText, result)
	} else {
		parseEnglishTelebirrSMS(smsText, result)
	}

	// Check if we have enough info
	if result.TransactionID != "" && result.Amount > 0 {
		result.IsValid = true
	}

	return result
}

// parseEnglishTelebirrSMS parses English Telebirr SMS
func parseEnglishTelebirrSMS(smsText string, result *TransactionInfo) {
	// Extract transaction number
	txnRegex := regexp.MustCompile(`transaction number is ([A-Z0-9]+)`)
	if matches := txnRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		result.TransactionID = matches[1]
	}

	// Extract amount
	amountRegex := regexp.MustCompile(`transferred ETB ([0-9.]+) to`)
	if matches := amountRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		if amount, err := strconv.ParseFloat(matches[1], 64); err == nil {
			result.Amount = amount
		}
	}

	// Extract sender
	senderRegex := regexp.MustCompile(`Dear ([A-Za-z ]+)`)
	if matches := senderRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		result.Sender = strings.TrimSpace(matches[1])
	}

	// Extract receiver - more flexible pattern
	receiverRegex := regexp.MustCompile(`to ([A-Za-z ]+)\([0-9]+\*+[0-9]+\)`)
	if matches := receiverRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		result.Receiver = strings.TrimSpace(matches[1])
	}

	// Alternative receiver pattern with space before (
	if result.Receiver == "" {
		receiverRegex2 := regexp.MustCompile(`to ([A-Za-z ]+) \([0-9]+\*+[0-9]+\)`)
		if matches := receiverRegex2.FindStringSubmatch(smsText); len(matches) > 1 {
			result.Receiver = strings.TrimSpace(matches[1])
		}
	}

	// Extract balance
	balanceRegex := regexp.MustCompile(`balance is ETB ([0-9.]+)`)
	if matches := balanceRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		if balance, err := strconv.ParseFloat(matches[1], 64); err == nil {
			result.Balance = balance
		}
	}

	// Extract date
	dateRegex := regexp.MustCompile(`on (\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2})`)
	if matches := dateRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		result.Date = matches[1]
	}
}

// parseAmharicTelebirrSMS parses Amharic Telebirr SMS
func parseAmharicTelebirrSMS(smsText string, result *TransactionInfo) {
	// Extract transaction number - Amharic pattern
	// "የሂሳብ እንቅስቃሴ ቁጥርዎ DH18FW1Y4C ነው"
	txnRegex := regexp.MustCompile(`ቁጥርዎ ([A-Z0-9]+) ነው`)
	if matches := txnRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		result.TransactionID = matches[1]
	}
	
	// Also try alternative pattern with "ነው።" at the end
	if result.TransactionID == "" {
		txnRegex2 := regexp.MustCompile(`ቁጥርዎ ([A-Z0-9]+) ነው።`)
		if matches := txnRegex2.FindStringSubmatch(smsText); len(matches) > 1 {
			result.TransactionID = matches[1]
		}
	}

	// Also try pattern with "ነዉ" (common typo)
	if result.TransactionID == "" {
		txnRegex3 := regexp.MustCompile(`ቁጥርዎ ([A-Z0-9]+) ነዉ`)
		if matches := txnRegex3.FindStringSubmatch(smsText); len(matches) > 1 {
			result.TransactionID = matches[1]
		}
	}

	// Extract amount - Amharic pattern
	// "ወደ Frezer Wudneh(2519****3937) 20.00 ብር ልከዋል"
	amountRegex := regexp.MustCompile(`\) ([0-9.]+) ብር ልከዋል`)
	if matches := amountRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		if amount, err := strconv.ParseFloat(matches[1], 64); err == nil {
			result.Amount = amount
		}
	}
	
	// Alternative amount pattern (if the above doesn't match)
	if result.Amount == 0 {
		amountRegex2 := regexp.MustCompile(`([0-9.]+) ብር በ`)
		if matches := amountRegex2.FindStringSubmatch(smsText); len(matches) > 1 {
			if amount, err := strconv.ParseFloat(matches[1], 64); err == nil {
				result.Amount = amount
			}
		}
	}

	// Alternative amount pattern with space
	if result.Amount == 0 {
		amountRegex3 := regexp.MustCompile(`\) ([0-9.]+) ብር`)
		if matches := amountRegex3.FindStringSubmatch(smsText); len(matches) > 1 {
			if amount, err := strconv.ParseFloat(matches[1], 64); err == nil {
				result.Amount = amount
			}
		}
	}

	// Extract receiver - Amharic pattern
	// "ወደ Frezer Wudneh(2519****3937)"
	receiverRegex := regexp.MustCompile(`ወደ ([^(]+)\([0-9]+\*+[0-9]+\)`)
	if matches := receiverRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		result.Receiver = strings.TrimSpace(matches[1])
	}
	
	// Alternative receiver pattern
	if result.Receiver == "" {
		receiverRegex2 := regexp.MustCompile(`ወደ ([^(]+)\(`)
		if matches := receiverRegex2.FindStringSubmatch(smsText); len(matches) > 1 {
			result.Receiver = strings.TrimSpace(matches[1])
		}
	}

	// Extract sender - Amharic pattern
	// "ውድ TIBEBU" or "ውድTIBEBU" (no space)
	senderRegex := regexp.MustCompile(`ውድ([A-Za-z ]+)`)
	if matches := senderRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		result.Sender = strings.TrimSpace(matches[1])
	}
	
	// Alternative sender pattern with space
	if result.Sender == "" {
		senderRegex2 := regexp.MustCompile(`ውድ ([A-Za-z ]+)`)
		if matches := senderRegex2.FindStringSubmatch(smsText); len(matches) > 1 {
			result.Sender = strings.TrimSpace(matches[1])
		}
	}

	// Extract balance - Amharic pattern
	// "አሁን ያለዎት ቀሪ ሂሳብ 2.00 ብር ነው"
	balanceRegex := regexp.MustCompile(`ቀሪ ሂሳብ ([0-9.]+) ብር`)
	if matches := balanceRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		if balance, err := strconv.ParseFloat(matches[1], 64); err == nil {
			result.Balance = balance
		}
	}
	
	// Alternative balance pattern (with "ነው")
	if result.Balance == 0 {
		balanceRegex2 := regexp.MustCompile(`ሂሳብ ([0-9.]+) ብር ነው`)
		if matches := balanceRegex2.FindStringSubmatch(smsText); len(matches) > 1 {
			if balance, err := strconv.ParseFloat(matches[1], 64); err == nil {
				result.Balance = balance
			}
		}
	}
	
	// Alternative balance pattern (with "ነው።")
	if result.Balance == 0 {
		balanceRegex3 := regexp.MustCompile(`ሂሳብ ([0-9.]+) ብር ነው።`)
		if matches := balanceRegex3.FindStringSubmatch(smsText); len(matches) > 1 {
			if balance, err := strconv.ParseFloat(matches[1], 64); err == nil {
				result.Balance = balance
			}
		}
	}

	// Extract date - Amharic pattern
	// "በ 01/08/2026 17:27:46 ልከዋል"
	dateRegex := regexp.MustCompile(`በ (\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2})`)
	if matches := dateRegex.FindStringSubmatch(smsText); len(matches) > 1 {
		result.Date = matches[1]
	}
}

// containsAmharic checks if the text contains Amharic characters
func containsAmharic(text string) bool {
	// Iterate through each rune and check if it's in the Amharic Unicode range
	for _, r := range text {
		// Amharic Unicode range: U+1200 to U+137F
		if r >= 0x1200 && r <= 0x137F {
			return true
		}
	}
	return false
}

// IsTelebirrSMS checks if the SMS is from Telebirr (supports both English and Amharic)
func IsTelebirrSMS(smsText string) bool {
	// English keywords
	enKeywords := []string{
		"telebirr",
		"Ethio telecom",
		"transferred ETB",
		"transaction number is",
	}
	
	// Amharic keywords
	amKeywords := []string{
		"ቴሌብር",
		"ኢትዮ ቴሌኮም",
		"ብር ልከዋል",
		"የሂሳብ እንቅስቃሴ ቁጥር",
		"ቀሪ ሂሳብ",
	}
	
	lowerText := strings.ToLower(smsText)
	
	// Check English keywords
	for _, keyword := range enKeywords {
		if strings.Contains(lowerText, strings.ToLower(keyword)) {
			return true
		}
	}
	
	// Check Amharic keywords (no need to lower case for Amharic)
	for _, keyword := range amKeywords {
		if strings.Contains(smsText, keyword) {
			return true
		}
	}
	
	return false
}

// FormatTransactionSummary returns a human-readable summary of the transaction
func FormatTransactionSummary(info *TransactionInfo) string {
	if !info.IsValid {
		return "Invalid transaction"
	}
	
	var sb strings.Builder
	sb.WriteString("✅ Transaction Details:\n")
	sb.WriteString(fmt.Sprintf("💰 Amount: %.2f ETB\n", info.Amount))
	sb.WriteString(fmt.Sprintf("🆔 Transaction ID: %s\n", info.TransactionID))
	
	if info.Sender != "" {
		sb.WriteString(fmt.Sprintf("👤 From: %s\n", info.Sender))
	}
	if info.Receiver != "" {
		sb.WriteString(fmt.Sprintf("👤 To: %s\n", info.Receiver))
	}
	if info.Date != "" {
		sb.WriteString(fmt.Sprintf("📅 Date: %s\n", info.Date))
	}
	if info.Balance > 0 {
		sb.WriteString(fmt.Sprintf("💰 Balance: %.2f ETB\n", info.Balance))
	}
	
	return sb.String()
}