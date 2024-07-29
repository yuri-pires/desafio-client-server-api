package arquivos

import (
	"bufio"
	"fmt"
	"os"
)

func SalvarCotacao(bid string) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	texto := fmt.Sprintf("Dolár: %s\n", bid)

	if _, err = writer.WriteString(texto); err != nil {
		return err
	}
	writer.Flush()

	return nil
}
