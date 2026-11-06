// Package password encapsula o hash e a verificação de senhas com bcrypt.
package password

import "golang.org/x/crypto/bcrypt"

// dummyHash equaliza o tempo de resposta quando não há hash real para comparar
var dummyHash = func() []byte {
	h, err := bcrypt.GenerateFromPassword([]byte("dummy-password-for-timing"), bcrypt.DefaultCost)
	if err != nil {
		panic("password: falha ao gerar o hash de equalização: " + err.Error())
	}
	return h
}()

func Hash(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hash), err
}

// Compare informa se a senha em texto puro corresponde ao hash
func Compare(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}

// CompareDummy gasta o mesmo custo de uma comparação real, evitando enumeração de usuários por tempo de resposta
func CompareDummy(plain string) {
	_ = bcrypt.CompareHashAndPassword(dummyHash, []byte(plain))
}
