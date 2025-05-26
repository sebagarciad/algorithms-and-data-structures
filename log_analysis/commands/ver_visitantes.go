package comandos

import (
	"fmt"
)

const (
	_VISITANTES = "Visitantes:"
)

// ================== VER VISITANTES ==================

// VerVisitantes imprime los visitantes que se encuentran en el rango [desde, hasta]
func (analyzer *dataAnalyzer) VerVisitantes(desde, hasta string) error {
	fmt.Println(_VISITANTES)
	for iterador := analyzer.ips.IteratorRange(&desde, &hasta); iterador.HasNext(); iterador.Next() {
		ip, _ := iterador.Current()
		fmt.Printf("\t%s\n", ip)
	}
	return nil
}
