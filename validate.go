package main

import (
	"fmt"
	"strings"
	"time"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateAF(numero int64, fornecedor, descricao, dataInicio, dataFim string, status bool) []ValidationError {
	var errors []ValidationError

	if numero <= 0 {
		errors = append(errors, ValidationError{
			Field:   "numero",
			Message: "NÚMERO DA AF DEVE SER MAIOR QUE ZERO",
		})
	}

	if strings.TrimSpace(fornecedor) == "" {
		errors = append(errors, ValidationError{
			Field:   "fornecedor",
			Message: "FORNECEDOR NÃO PODE ESTAR VAZIO",
		})
	}

	if strings.TrimSpace(descricao) == "" {
		errors = append(errors, ValidationError{
			Field:   "descricao",
			Message: "DESCRIÇÃO NÃO PODE ESTAR VAZIA",
		})
	}

	inicio, err := time.Parse(time.DateOnly, dataInicio)
	if err != nil {
		errors = append(errors, ValidationError{
			Field:   "data_inicio",
			Message: "DATA DE INÍCIO INVÁLIDA - USE O FORMATO AAAA-MM-DD",
		})
	}

	fim, err := time.Parse(time.DateOnly, dataFim)
	if err != nil {
		errors = append(errors, ValidationError{
			Field:   "data_fim",
			Message: "DATA DE FIM INVÁLIDA - USE O FORMATO AAAA-MM-DD",
		})
	} else if !inicio.IsZero() && !fim.IsZero() && fim.Before(inicio) {
		errors = append(errors, ValidationError{
			Field:   "data_fim",
			Message: "DATA DE FIM NÃO PODE SER ANTERIOR À DATA DE INÍCIO",
		})
	}

	return errors
}

func validateItem(descricao string, bandaMaxima, bandaInstalada int64, dataInstalacao string, quantidade int64, status bool) []ValidationError {
	var errors []ValidationError

	if strings.TrimSpace(descricao) == "" {
		errors = append(errors, ValidationError{
			Field:   "descricao",
			Message: "DESCRIÇÃO NÃO PODE ESTAR VAZIA",
		})
	}

	if bandaMaxima <= 0 {
		errors = append(errors, ValidationError{
			Field:   "banda_maxima",
			Message: "BANDA MÁXIMA DEVE SER MAIOR QUE ZERO",
		})
	}

	if bandaInstalada <= 0 {
		errors = append(errors, ValidationError{
			Field:   "banda_instalada",
			Message: "BANDA INSTALADA DEVE SER MAIOR QUE ZERO",
		})
	}

	if quantidade <= 0 {
		errors = append(errors, ValidationError{
			Field:   "quantidade",
			Message: "QUANTIDADE DEVE SER MAIOR QUE ZERO",
		})
	}

	if bandaInstalada > bandaMaxima {
		errors = append(errors, ValidationError{
			Field:   "banda_instalada",
			Message: "BANDA INSTALADA NÃO PODE EXCEDER A BANDA MÁXIMA",
		})
	}

	if dataInstalacao == "" {
		errors = append(errors, ValidationError{
			Field:   "data_instalacao",
			Message: "DATA DE INSTALAÇÃO É OBRIGATÓRIA",
		})
	} else {
		_, err := time.Parse(time.DateOnly, dataInstalacao)
		if err != nil {
			errors = append(errors, ValidationError{
				Field:   "data_instalacao",
				Message: "DATA DE INSTALAÇÃO INVÁLIDA - USE O FORMATO AAAA-MM-DD",
			})
		}
	}

	return errors
}

func validateValor(valor int64, dataInicio, dataFim string, itemID int64) []ValidationError {
	var errors []ValidationError

	if valor <= 0 {
		errors = append(errors, ValidationError{
			Field:   "valor",
			Message: "VALOR DEVE SER MAIOR QUE ZERO",
		})
	}

	if dataInicio == "" {
		errors = append(errors, ValidationError{
			Field:   "data_inicio",
			Message: "DATA DE INÍCIO É OBRIGATÓRIA",
		})
	} else {
		inicio, err := time.Parse("2006-01-02", dataInicio)
		if err != nil {
			errors = append(errors, ValidationError{
				Field:   "data_inicio",
				Message: "DATA DE INÍCIO INVÁLIDA - USE O FORMATO AAAA-MM-DD",
			})
		} else if inicio.After(time.Now()) {
			errors = append(errors, ValidationError{
				Field:   "data_inicio",
				Message: "DATA DE INÍCIO NÃO PODE SER FUTURA",
			})
		}
	}

	if dataFim != "" {
		fim, err := time.Parse("2006-01-02", dataFim)
		if err != nil {
			errors = append(errors, ValidationError{
				Field:   "data_fim",
				Message: "DATA DE FIM INVÁLIDA - USE O FORMATO AAAA-MM-DD",
			})
		} else {
			inicio, _ := time.Parse("2006-01-02", dataInicio)
			if !inicio.IsZero() && fim.Before(inicio) {
				errors = append(errors, ValidationError{
					Field:   "data_fim",
					Message: "DATA DE FIM NÃO PODE SER ANTERIOR À DATA DE INÍCIO",
				})
			}
		}
	}

	return errors
}
