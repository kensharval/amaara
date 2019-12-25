package main

import (
	"bufio"
	"fmt"
	"github.com/qamarian-dtp/err"
	errLib "github.com/qamarian-lib/err"
	"github.com/qamarian-lib/newton"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main () {
	for {
		fmt.Print (`

==================================================================
==  ready to serve, please input; ready to serve, please input  ==
==================================================================

`)
		defer fmt.Print ("\n\n\n\n")

		input, errX := i.ReadString (byte ('\n'))
		fmt.Println ()
		if errX != nil {
			e := err.New ("Error: unable to read input.", nil, nil, errX)
			fmt.Println (errLib.Fup (e))
			continue
		}
		input = strings.TrimRight (input, "\n")

		if ! inputPattern.MatchString (input) {
			e := err.New ("Error: invalid input.", nil, nil)
			fmt.Println (errLib.Fup (e))
			continue
		}

		parts := strings.Split (input, " ")
		if len (parts) != 2 {
			e := err.New ("Error: bug detected; ref 0.", nil, nil)
			fmt.Println (errLib.Fup (e))
			continue
		}

		quantity, errY := strconv.Atoi (parts [1])
		if errY != nil {
			e := err.New ("Error: number of spelling may be too high.", nil,
				nil, errY)
			fmt.Println (errLib.Fup (e))
			continue
		}

		for x := 1; x <= quantity; x ++ {
			name, errZ := newton.Name_Rand (parts [0])
			if errZ != nil {
				e := err.New ("Generation of a name failed: unable to " +
					"obtain a name.", nil, nil, errZ)
				fmt.Println (errLib.Fup (e))
				continue
			}

			name.Polish ()

			spellings, errA := name.Spelling (8)
			if errA != nil {
				e := err.New ("Generate of a name failed: unable to " +
					"obtain spelling of the name sound.", nil, nil,
					errA)
				fmt.Println (errLib.Fup (e))
				continue
			}

			o := fmt.Sprintf ("%s:  %v", name.String (), spellings)
			fmt.Println (o)
		}
	}
}; var (
	i = bufio.NewReader (os.Stdout)
	inputPattern *regexp.Regexp
); func init () {
	var errX error
	inputPattern, errX = regexp.Compile (`^[cve]+ [1-9][\d]*$`)
	if errX != nil {
		e := err.New ("Startup failed: regular expression compilation failed.",
			nil, nil, errX)
		fmt.Println (e)
		os.Exit (1)
	}
}
