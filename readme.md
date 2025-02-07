## Go Secure
Gosec kütüphanesi golang dilinde yazılmış güvenlik kütüphanesidir. Bu kütüphane ile kodlarınızda bulunan olası güvenlik açıklarını tespit edebilirsiniz.

### Kurulum

Gosec kütüphanesini kurmak için aşağıdaki komutu çalıştırmanız yeterlidir.

```bash
go install github.com/securego/gosec/v2/cmd/gosec@latest
```
Sonrasında path değişkenine gosec kütüphanesinin path'ini ekleyin.

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Kullanım

Gosec kütüphanesini kullanmak için aşağıdaki komutu çalıştırmanız yeterlidir.

```bash
gosec ./...
```

### Örnek

Örnek bir kod parçası aşağıdaki gibidir.

```go
package main

import (
	"log"
	"os"
)

func main() {
	if err:=InSecureCode(); err!=nil{
		log.Println("Error: ", err)
		return
	}
}

func InSecureCode() error{
	err := os.Mkdir("./example", 0777)
	if err != nil {
		return err
	}
	return nil
}

```

Bu kod parçasını gosec kütüphanesi ile kontrol ettiğimizde aşağıdaki gibi bir çıktı alırız.

```bash
go-secure/main.go:16] - G301 (CWE-276): Expect directory permissions to be 0750 or less (Confidence: HIGH, Severity: MEDIUM)
    15: func InSecureCode() error{
  > 16:         err := os.Mkdir("./example", 0777)
    17:         if err != nil {

Autofix: 

Summary:
  Gosec  : dev
  Files  : 1
  Lines  : 21
  Nosec  : 0
  Issues : 1
```




Teşekkürler 💫