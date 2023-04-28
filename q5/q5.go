package q5

import "strings"

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	/*var final string
	var resultado string
	vogais := []string{"a", "e", "i", "o", "u"}
	novo := strings.ToLower(s)
	for i := 0; i < len(vogais); i++ {
		final = strings.ReplaceAll(novo, vogais[i], "")
		resultado = ""
	}

	for i := 0; i < len(final); i++ {
		resultado += "." + string(final[i])
	}
	return resultado*/
	novo := strings.ToLower(s)
	a := strings.ReplaceAll(novo, "a", "")
	e := strings.ReplaceAll(a, "e", "")
	i := strings.ReplaceAll(e, "i", "")
	o := strings.ReplaceAll(i, "o", "")
	u := strings.ReplaceAll(o, "u", "")

	resultado := ""

	for i := 0; i < len(u); i++ {
		resultado += "." + string(u[i])
	}
	return resultado
}
