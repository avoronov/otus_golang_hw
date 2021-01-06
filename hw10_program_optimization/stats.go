package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

var (
	// ErrMalformedData is the error which thrown when input data malformed.
	ErrMalformedData = errors.New("malformed data")
	emailFieldPrefix = []byte(`"Email":"`)
)

// DomainStat is the struct describing domain stats.
type DomainStat map[string]int

// GetDomainStat cal stats for given domain.
func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	domainSuffix := []byte("." + domain)

	result := make(DomainStat)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		userDomain, err := parseEmailDomain(scanner.Bytes())
		if err != nil {
			return DomainStat{}, fmt.Errorf("can't parse email domain: %w", err)
		}

		if !bytes.HasSuffix(userDomain, domainSuffix) {
			continue
		}

		result[string(userDomain)]++
	}

	if scanner.Err() != nil {
		return DomainStat{}, fmt.Errorf("got error while process input data: %w", scanner.Err())
	}

	return result, nil
}

func parseEmailDomain(rec []byte) ([]byte, error) {
	// NB: string are valid JSON!

	if len(rec) == 0 || rec[0] != '{' || rec[len(rec)-1] != '}' {
		return nil, ErrMalformedData
	}

	rec = rec[1 : len(rec)-1] // trim { and } from the beginning and the end of the input

	startEmailRec := bytes.Index(rec, emailFieldPrefix) // assume here there are no another entires of emailFieldPrefix
	if startEmailRec == -1 {
		return nil, ErrMalformedData
	}

	rec = rec[startEmailRec+len(emailFieldPrefix):] // trim emailFieldPrefix

	endEmailRec := bytes.IndexByte(rec, ',') // assume here there are no another entires of ,
	if endEmailRec != -1 {
		rec = rec[:endEmailRec] // trim rest of input beyond the email key-value pair
	}

	if rec[len(rec)-1] != '"' {
		return nil, ErrMalformedData
	}
	rec = rec[:len(rec)-1] // trim " from the end of the input

	startDomain := bytes.IndexByte(rec, '@') // assume here there are no another entires of @
	if startDomain == -1 {
		return nil, ErrMalformedData
	}
	rec = rec[startDomain+1:] // trim @ from the beginning of the input

	if len(rec) == 0 {
		return nil, ErrMalformedData
	}

	return bytes.ToLower(rec), nil
}
