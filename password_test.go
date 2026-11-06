package password

import "testing"

func TestHashAndCompare(t *testing.T) {
	hash, err := Hash("12345678")
	if err != nil {
		t.Fatal(err)
	}
	if hash == "12345678" {
		t.Fatal("o hash não pode ser a senha em texto puro")
	}
	if !Compare(hash, "12345678") {
		t.Error("senha correta deveria validar")
	}
	if Compare(hash, "outra-senha") {
		t.Error("senha errada não deveria validar")
	}
	if Compare("hash-invalido", "12345678") {
		t.Error("hash inválido não deveria validar")
	}
	CompareDummy("qualquer")
}
