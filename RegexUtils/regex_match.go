package RegexUtils

import (
	"regexp"
)

func match(text string, regex *regexp.Regexp) []string {
	parsed := regex.FindAllString(text, -1)
	return parsed
}

// Date finds all date strings
func Date(text string) []string {
	return match(text, DateRegex)
}

// Time finds all time strings
func Time(text string) []string {
	return match(text, TimeRegex)
}

// Phones finds all phone numbers
func Phones(text string) []string {
	return match(text, PhoneRegex)
}

// PhonesWithExts finds all phone numbers with ext
func PhonesWithExts(text string) []string {
	return match(text, PhonesWithExtsRegex)
}

// Links finds all link strings
func Links(text string) []string {
	return match(text, LinkRegex)
}

// Emails finds all email strings
func Emails(text string) []string {
	return match(text, EmailRegex)
}

// IPv4s finds all IPv4 addresses
func IPv4s(text string) []string {
	return match(text, IPv4Regex)
}

// IPv6s finds all IPv6 addresses
func IPv6s(text string) []string {
	return match(text, IPv6Regex)
}

// IPs finds all IP addresses (both IPv4 and IPv6)
func IPs(text string) []string {
	return match(text, IPRegex)
}

// NotKnownPorts finds all not-known port numbers
func NotKnownPorts(text string) []string {
	return match(text, NotKnownPortRegex)
}

// Prices finds all price strings
func Prices(text string) []string {
	return match(text, PriceRegex)
}

// HexColors finds all hex color values
func HexColors(text string) []string {
	return match(text, HexColorRegex)
}

// CreditCards finds all credit card numbers
func CreditCards(text string) []string {
	return match(text, CreditCardRegex)
}

// BtcAddresses finds all bitcoin addresses
func BtcAddresses(text string) []string {
	return match(text, BtcAddressRegex)
}

// StreetAddresses finds all street addresses
func StreetAddresses(text string) []string {
	return match(text, StreetAddressRegex)
}

// ZipCodes finds all zip codes
func ZipCodes(text string) []string {
	return match(text, ZipCodeRegex)
}

// PoBoxes finds all po-box strings
func PoBoxes(text string) []string {
	return match(text, PoBoxRegex)
}

// SSNs finds all SSN strings
func SSNs(text string) []string {
	return match(text, SSNRegex)
}

// MD5Hexes finds all MD5 hex strings
func MD5Hexes(text string) []string {
	return match(text, MD5HexRegex)
}

// SHA1Hexes finds all SHA1 hex strings
func SHA1Hexes(text string) []string {
	return match(text, SHA1HexRegex)
}

// SHA256Hexes finds all SHA256 hex strings
func SHA256Hexes(text string) []string {
	return match(text, SHA256HexRegex)
}

// GUIDs finds all GUID strings
func GUIDs(text string) []string {
	return match(text, GUIDRegex)
}

// ISBN13s finds all ISBN13 strings
func ISBN13s(text string) []string {
	return match(text, ISBN13Regex)
}

// ISBN10s finds all ISBN10 strings
func ISBN10s(text string) []string {
	return match(text, ISBN10Regex)
}

// VISACreditCards finds all VISA credit card numbers
func VISACreditCards(text string) []string {
	return match(text, VISACreditCardRegex)
}

// MCCreditCards finds all MasterCard credit card numbers
func MCCreditCards(text string) []string {
	return match(text, MCCreditCardRegex)
}

// MACAddresses finds all MAC addresses
func MACAddresses(text string) []string {
	return match(text, MACAddressRegex)
}

// IBANs finds all IBAN strings
func IBANs(text string) []string {
	return match(text, IBANRegex)
}

// GitRepos finds all git repository addresses which have protocol prefix
func GitRepos(text string) []string {
	return match(text, GitRepoRegex)
}

func Emojis(text string) []string {
	return match(text, EmojiRegex)
}
