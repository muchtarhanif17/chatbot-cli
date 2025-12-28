package prompt

import "fmt"

func Build(context, question string) string {
	return fmt.Sprintf(`
Kamu adalah chatbot internal perusahaan.
Jawab Hanya berdasarkan data berikut.

DATA:
%s

PERTANYAAN:
%s
`, context, question)
}
