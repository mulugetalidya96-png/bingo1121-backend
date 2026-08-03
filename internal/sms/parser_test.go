package sms

import (
	"fmt"
	"testing"
)

func TestParseTelebirrAmharicSMS(t *testing.T) {
	// Your Telebirr Amharic SMS example
	smsText := `ውድMuse
ወደ Rodas Mulugeta(2519****4062) 50.00 ብር በ 03/08/2026 16:52:28  ልከዋል። የሂሳብ እንቅስቃሴ ቁጥርዎ DH31HOHKQR ነዉ። የአገልግሎት ክፍያው 0.87 ብር ፤ የአገልግሎት ክፍያው 15% VAT 0.13 ብር ነዉ።  አሁን ያለዎት ቀሪ ሂሳብ 175.08 ብር ነው። የክፍያ መረጃዎን ለማግኘት ማስፈንጠርያውን ይጫኑ፡ https://transactioninfo.ethiotelecom.et/receipt/DH31HOHKQR
በቴሌብር ስለተገለገሉ እናመሰግናለን 
ኢትዮ ቴሌኮም`

	fmt.Println("📱 Testing Telebirr Amharic SMS Parser")
	fmt.Println("=========================================")
	fmt.Printf("SMS Text: %s\n\n", smsText[:200]+"...")

	result := ParseTelebirrSMS(smsText)

	fmt.Println("📊 Parsed Results:")
	fmt.Printf("  ✅ IsValid: %v\n", result.IsValid)
	fmt.Printf("  💰 Amount: %.2f ETB\n", result.Amount)
	fmt.Printf("  🆔 Transaction ID: %s\n", result.TransactionID)
	fmt.Printf("  👤 Sender: %s\n", result.Sender)
	fmt.Printf("  👤 Receiver: %s\n", result.Receiver)
	fmt.Printf("  📱 Receiver Phone: %s\n", result.ReceiverPhone)
	fmt.Printf("  📅 Date: %s\n", result.Date)
	fmt.Printf("  💳 Balance: %.2f ETB\n", result.Balance)
	fmt.Printf("  💸 Fee: %.2f ETB\n", result.Fee)
	fmt.Printf("  💸 VAT: %.2f ETB\n", result.VAT)

	if !result.IsValid {
		t.Errorf("❌ Failed to parse Telebirr Amharic SMS")
	} else {
		fmt.Println("\n✅ Telebirr Amharic SMS parsed successfully!")
	}
}

func TestParseCBEBirrAmharicSMS(t *testing.T) {
	// Your CBE Birr Amharic SMS example
	smsText := `ውድTIBEBU
ወደ Frezer Wudneh(2519****3937) 20.00 ብር በ 01/08/2026 17:27:46  ልከዋል። የሂሳብ እንቅስቃሴ ቁጥርዎ DH18FW1Y4C ነዉ። የአገልግሎት ክፍያው 0.87 ብር ፤ የአገልግሎት ክፍያው 15% VAT 0.13 ብር ነዉ።  አሁን ያለዎት ቀሪ ሂሳብ 2.00 ብር ነው። የክፍያ መረጃዎን ለማግኘት ማስፈንጠርያውን ይጫኑ፡ https://transactioninfo.ethiotelecom.et/receipt/DH18FW1Y4C
በቴሌብር ስለተገለገሉ እናመሰግናለን 
ኢትዮ ቴሌኮም`

	fmt.Println("\n📱 Testing CBE Birr Amharic SMS Parser")
	fmt.Println("=========================================")
	fmt.Printf("SMS Text: %s\n\n", smsText[:200]+"...")

	result := ParseCBEBirrSMS(smsText)

	fmt.Println("📊 Parsed Results:")
	fmt.Printf("  ✅ IsValid: %v\n", result.IsValid)
	fmt.Printf("  💰 Amount: %.2f ETB\n", result.Amount)
	fmt.Printf("  🆔 Transaction ID: %s\n", result.TransactionID)
	fmt.Printf("  🆔 Reference Number: %s\n", result.ReferenceNumber)
	fmt.Printf("  👤 Sender: %s\n", result.Sender)
	fmt.Printf("  👤 Receiver: %s\n", result.Receiver)
	fmt.Printf("  📱 Receiver Phone: %s\n", result.ReceiverPhone)
	fmt.Printf("  📅 Date: %s\n", result.Date)
	fmt.Printf("  💳 Balance: %.2f ETB\n", result.Balance)
	fmt.Printf("  💸 Fee: %.2f ETB\n", result.Fee)
	fmt.Printf("  💸 VAT: %.2f ETB\n", result.VAT)
	fmt.Printf("  🔗 Receipt URL: %s\n", result.ReceiptURL)

	if !result.IsValid {
		t.Errorf("❌ Failed to parse CBE Birr Amharic SMS")
	} else {
		fmt.Println("\n✅ CBE Birr Amharic SMS parsed successfully!")
	}
}

func TestIsTelebirrSMSCheck(t *testing.T) {
	amharicSMS := `ውድMuse ወደ Rodas Mulugeta(2519****4062) 50.00 ብር በ 03/08/2026 ልከዋል።`

	fmt.Println("\n📱 Testing Telebirr SMS Detection")
	fmt.Println("=========================================")

	isTelebirr := IsTelebirrSMS(amharicSMS)
	fmt.Printf("  Is Telebirr SMS: %v\n", isTelebirr)

	if !isTelebirr {
		t.Errorf("❌ Failed to detect Telebirr SMS")
	} else {
		fmt.Println("✅ Telebirr SMS detected successfully!")
	}
}

func TestIsCBEBirrSMSCheck(t *testing.T) {
	amharicSMS := `ውድTIBEBU ወደ Frezer Wudneh 20.00 ብር በ 01/08/2026 ልከዋል።`

	fmt.Println("\n📱 Testing CBE Birr SMS Detection")
	fmt.Println("=========================================")

	isCBEBirr := IsCBEBirrSMS(amharicSMS)
	fmt.Printf("  Is CBE Birr SMS: %v\n", isCBEBirr)

	if !isCBEBirr {
		t.Errorf("❌ Failed to detect CBE Birr SMS")
	} else {
		fmt.Println("✅ CBE Birr SMS detected successfully!")
	}
}

func TestGetTelebirrType(t *testing.T) {
	amharicSMS := `ውድMuse ወደ Rodas Mulugeta 50.00 ብር ልከዋል።`
	englishSMS := `Dear Muse, you transferred ETB 50.00 to Rodas Mulugeta.`

	fmt.Println("\n📱 Testing Telebirr Type Detection")
	fmt.Println("=========================================")

	type1 := GetTelebirrType(amharicSMS)
	type2 := GetTelebirrType(englishSMS)

	fmt.Printf("  Amharic SMS type: %s\n", type1)
	fmt.Printf("  English SMS type: %s\n", type2)

	if type1 != "amharic" {
		t.Errorf("❌ Failed to detect Amharic Telebirr type")
	}
	if type2 != "english" {
		t.Errorf("❌ Failed to detect English Telebirr type")
	}
	fmt.Println("✅ Telebirr type detection successful!")
}

func TestGetCBEBirrType(t *testing.T) {
	amharicSMS := `ውድTIBEBU ወደ Frezer Wudneh 20.00 ብር ልከዋል።`
	englishSMS := `Dear TIBEBU, you sent 20.00 Br to Frezer Wudneh.`

	fmt.Println("\n📱 Testing CBE Birr Type Detection")
	fmt.Println("=========================================")

	type1 := GetCBEBirrType(amharicSMS)
	type2 := GetCBEBirrType(englishSMS)

	fmt.Printf("  Amharic SMS type: %s\n", type1)
	fmt.Printf("  English SMS type: %s\n", type2)

	if type1 != "amharic" {
		t.Errorf("❌ Failed to detect Amharic CBE Birr type")
	}
	if type2 != "english" {
		t.Errorf("❌ Failed to detect English CBE Birr type")
	}
	fmt.Println("✅ CBE Birr type detection successful!")
}