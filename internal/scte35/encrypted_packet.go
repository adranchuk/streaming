// Copyright 2024 The Untangled Authors. Use of this source code is
// governed by the ISC license available in the LICENSE.ISC file.

package scte35

const (
	EncryptionAlgorithmNone uint32 = iota
	EncryptionAlgorithmDESECB
	EncryptionAlgorithmDESCBC
	EncryptionAlgorithmTripleDES
)

// EncryptedPacket represents payload encryption information.
type EncryptedPacket struct {
	// Specifies which cipher is used.
	// See SCTE 35 section 11.3.
	EncryptionAlgorithm uint32
	CWIndex             uint32
}

// cipher is a 6-bit field specifying the algorithm used to encrypt
// payloads as defined in SCTE 35 section 11.3.
type Cipher uint8

const (
	CipherNone Cipher = iota
	DES_ECB           // SCTE 35 section 11.3.1
	DES_CBC           // SCTE 35 section 11.3.2
	TripleDES         // SCTE 35 section 11.3.3
	reserved
	// Values 32 through 63 are available for "User private"
	// algorithms. See SCTE 35 section 11.3.4.
)

const maxCipher = 63

func (c Cipher) String() string {
	switch c {
	case CipherNone:
		return "none"
	case DES_ECB:
		return "DES – ECB mode"
	case DES_CBC:
		return "DES – CBC mode"
	case TripleDES:
		return "Triple DES EDE3 – ECB mode"
	}
	if c >= reserved && c <= 31 {
		return "reserved"
	} else if c <= maxCipher {
		return "user private"
	}
	return "invalid"
}
