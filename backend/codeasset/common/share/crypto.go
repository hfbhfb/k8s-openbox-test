package share

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"regexp"
	"strings"
	"time"

	//"github.com/hyperledger/fabric/bccsp/utils"
	"golang.org/x/crypto/ripemd160"
)

var (
	b58Alphabet        = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	addressChecksumLen = 4
)

type ECDSASignature struct {
	R, S *big.Int
}

type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

// newKeyPair create privatekey and pulibckey by ecdsa
func NewKeyPair() *KeyPair {
	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil
	}
	publicKey := &private.PublicKey
	encPriv, encPub := encodeKey(private, publicKey)

	return &KeyPair{
		PublicKey:  encPub,
		PrivateKey: encPriv,
	}
}

func encodeKey(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	privateKeyBase64 := Base64Encode(pemEncoded)
	publicKeyBase64 := Base64Encode(pemEncodedPub)
	return privateKeyBase64, publicKeyBase64
}

func (kp *KeyPair) GetAddress() []byte {
	return GetAddress([]byte(kp.PublicKey))
}

func EncodePrivateKey(privateKey *ecdsa.PrivateKey) string {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	privateKeyBase64 := Base64Encode(pemEncoded)
	return privateKeyBase64
}

func DecodePrivateKey(pemEncoded string) (*ecdsa.PrivateKey, error) {
	dstPrivate, err := Base64Decode(pemEncoded)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(dstPrivate)
	if block == nil {
		return nil, errors.New("privatekey decode error")
	}
	x509Encoded := block.Bytes
	privateKey, err := x509.ParseECPrivateKey(x509Encoded)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func DecodePulibcKey(pemEncodedPub string) (*ecdsa.PublicKey, error) {
	dstPublic, err := Base64Decode(pemEncodedPub)
	if err != nil {
		return nil, err
	}

	blockPub, _ := pem.Decode(dstPublic)
	if blockPub == nil {
		return nil, errors.New("publickey decode error")
	}
	genericPublicKey, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := genericPublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("ParsePKIXPublicKey interface to ecdsa.PublicKey error")
	}
	return publicKey, nil
}

func SignatureHash(arr []byte) []byte {
	hash := sha256.Sum256(arr)
	return hash[:]
}

//func ECSign(k *ecdsa.PrivateKey, digest []byte) ([]byte, error) {
//	r, s, err := ecdsa.Sign(rand.Reader, k, digest)
//	if err != nil {
//		return nil, err
//	}
//	s, err = utils.ToLowS(&k.PublicKey, s)
//	if err != nil {
//		return nil, err
//	}
//	return utils.MarshalECDSASignature(r, s)
//}

func Verify(certFile, text string, sign string) {
	certBuff, err := ioutil.ReadFile(certFile)
	if err != nil {
		fmt.Printf("ERROR: failed to read keystore file: %s, error: %s\n", certFile, err)
		return
	}

	block, _ := pem.Decode(certBuff)
	if block == nil {
		fmt.Printf("ERROR: block of decoded private key is nil\n")
		return
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Printf("ERROR: failed get ECDSA private key, error: %v\n", err)
		return
	}

	arr := []byte(text)
	h := sha256.New()
	h.Write(arr)
	hashed := h.Sum(nil)

	signatureDec, _ := base64.StdEncoding.DecodeString(sign)
	sig := new(ECDSASignature)
	_, err = asn1.Unmarshal(signatureDec, sig)
	if err != nil {
		fmt.Printf("ERROR: failed unmashalling signature, error: %v", err)
		return
	}

	pub, _ := cert.PublicKey.(*ecdsa.PublicKey)
	if !ecdsa.Verify(pub, hashed[:], sig.R, sig.S) {
		fmt.Printf("ERROR: Failed to verify Signature: %v\n", err)
		return
	}
	fmt.Printf("Successed to verify Signature and nonce\n")
	return
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// GetAddress returns address
func GetAddress(publicKey []byte) []byte {
	pubKeyHash := hashPubKey((publicKey))
	versionedPayload := append([]byte(VERSION), pubKeyHash...)
	checksum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)
	return address
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	reverseBytes(result)
	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}
	return result
}

func hashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		return nil
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

//判断是邮箱还是手机- 返回 "email"/"phone"/""
func VerifyAccountType(account string) string {
	if strings.Contains(account, "@") {
		pattern := `^(\w-*\.*)+@(\w-?)+(\.\w{2,})+$`
		reg := regexp.MustCompile(pattern)
		if reg.MatchString(account) {
			return "email"
		}
	} else {
		regular := `^1[3-5789]\d{9}$`
		reg := regexp.MustCompile(regular)
		if reg.MatchString(account) {
			return "phone"
		}
	}
	return ""
}

func CheckUserPwd(pwd string) (bool, error) {
	pattern := `^[a-zA-Z0-9]{6,20}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(pwd) {
		err := errors.New("请输入6~20位数字或字母密码")
		return false, err
	} else {
		return true, nil
	}
}

func CheckIsAllNumber(content string) (bool, error) {
	pattern := `^[0-9]*$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(content) {
		return false, nil
	} else {
		return true, nil
	}
}

func JudgeInNameCompliance(name string) bool {
	if name == "" {
		return true
	}
	pattern := `^[^@#$%^&*{}:"\\<>?]{1,10}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(name)
}

func GetBeginOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

func GetEndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, month, day, 23, 59, 59, 999, location)
}

func GetBeginOfDayRFC1123(t time.Time) string {
	str := t.Format(time.RFC1123)
	//"Mon, 02 Jan 2006 15:04:05 MST"
	strArr1 := strings.Split(str, ",")
	str = strArr1[1]
	str = str[1:]
	strArr2 := strings.Split(str, " ")
	return fmt.Sprintf("%v/%v/%v:00:00:00", strArr2[0], strArr2[1], strArr2[2])
}
func GetEndOfDayRFC1123(t time.Time) string {
	str := t.Format(time.RFC1123)
	//"Mon, 02 Jan 2006 15:04:05 MST"
	strArr1 := strings.Split(str, ",")
	str = strArr1[1]
	str = str[1:]
	strArr2 := strings.Split(str, " ")
	return fmt.Sprintf("%v/%v/%v:23:59:59", strArr2[0], strArr2[1], strArr2[2])
}
