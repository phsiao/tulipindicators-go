// Generate wrapped go functions from https://tulipindicators.org/list
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	IndexURL        = "https://tulipindicators.org/list"
	FunctionURLBase = "https://tulipindicators.org/"
)

func assertNoError(e error) {
	if e != nil {
		panic(e.Error())
	}
}

type Indicator struct {
	Identifier    string
	IndicatorName string
	Type          string
	Inputs        int
	Options       int
	Outputs       int
}

func index() []Indicator {
	u, err := url.Parse(IndexURL)
	assertNoError(err)

	resp, err := http.Get(u.String())
	assertNoError(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	assertNoError(err)

	indicators := []Indicator{}

	doc.Find("table.sortable").Each(func(i int, s *goquery.Selection) {
		// verify thead length
		columns := s.Find("thead tr th")
		if len(columns.Nodes) != 6 {
			log.Fatal(fmt.Sprintf("unexpected number of columns (expect 6 got %d", len(columns.Nodes)))
		}

		s.Find("tbody tr").Each(func(j int, row *goquery.Selection) {
			indicator := Indicator{}
			row.Find("td").Each(func(k int, element *goquery.Selection) {
				switch k {
				case 0:
					indicator.Identifier = element.Text()
				case 1:
					indicator.IndicatorName = element.Text()
				case 2:
					indicator.Type = element.Text()
				case 3:
					n, err := strconv.Atoi(element.Text())
					assertNoError(err)
					indicator.Inputs = n
				case 4:
					n, err := strconv.Atoi(element.Text())
					assertNoError(err)
					indicator.Options = n
				case 5:
					n, err := strconv.Atoi(element.Text())
					assertNoError(err)
					indicator.Outputs = n
				}
			})
			indicators = append(indicators, indicator)
		})
	})

	return indicators
}

func sanitizeTokens(tokens []string) []string {
	rval := tokens
	for i, t := range rval {
		rval[i] = strings.ReplaceAll(t, " ", "_")
		rval[i] = strings.ReplaceAll(rval[i], "%", "pct")
		switch t {
		case "var":
			rval[i] = "var_"
		case "real":
			if len(tokens) > 1 {
				rval[i] = fmt.Sprintf("real%d", i+1)
			}
		}
	}
	return rval
}

func getParameterNameMap(indicator string) (inputs, options, outputs []string, err error) {
	u, err := url.Parse(FunctionURLBase)
	assertNoError(err)
	u.Path = indicator

	resp, err := http.Get(u.String())
	assertNoError(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	assertNoError(err)

	doc.Find("code").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			return
		}
		blob := s.Text()
		lines := strings.Split(blob, "\n")
		var nInputs, nOptions, nOutputs int
		_, err = fmt.Sscanf(lines[2], "/* Input arrays: %d    Options: %d    Output arrays: %d */",
			&nInputs, &nOptions, &nOutputs)
		assertNoError(err)

		var tokens []string
		inputLine := lines[3]
		inputLine = inputLine[len("/* Inputs: ") : len(inputLine)-len(" */")]
		tokens = strings.Split(inputLine, ", ")
		inputs = sanitizeTokens(tokens)

		optionLine := lines[4]
		optionLine = optionLine[len("/* Options: ") : len(optionLine)-len(" */")]
		tokens = strings.Split(optionLine, ", ")
		options = sanitizeTokens(tokens)

		outputLine := lines[5]
		outputLine = outputLine[len("/* Outputs: ") : len(outputLine)-len(" */")]
		tokens = strings.Split(outputLine, ", ")
		outputs = sanitizeTokens(tokens)
	})

	return
}

func generateIndicators(indicators []Indicator) error {
	for _, indicator := range indicators {
		f, err := os.Create(fmt.Sprintf("indicators/%s.go", indicator.Identifier))
		assertNoError(err)
		defer f.Close()

		inputs, options, outputs, err := getParameterNameMap(indicator.Identifier)
		assertNoError(err)

		/*
			inputs := []string{}
			for idx := 0; idx < indicator.Inputs; idx++ {
				inputs = append(inputs, fmt.Sprintf("input%d", idx+1))
			}
			options := []string{}
			for idx := 0; idx < indicator.Options; idx++ {
				options = append(options, fmt.Sprintf("option%d", idx+1))
			}
			outputs := []string{}
			for idx := 0; idx < indicator.Outputs; idx++ {
				outputs = append(outputs, fmt.Sprintf("output%d", idx+1))
			}
		*/

		fmt.Fprintf(f, "package indicators\n\n")

		fmt.Fprintf(f, "//#include \"../tulipindicators/indicators/%s.c\"\n", indicator.Identifier)
		fmt.Fprintf(f, "import \"C\"\n")
		fmt.Fprintf(f, "import \"fmt\"\n")
		fmt.Fprintf(f, "\n")

		fmt.Fprintf(f, "// %s function wraps `%s' function that provides \"%s\"\n//\n",
			strings.ToUpper(indicator.Identifier),
			indicator.Identifier,
			indicator.IndicatorName)
		fmt.Fprintf(f, "// Reference: https://tulipindicators.org/%s\n", indicator.Identifier)
		if indicator.Options > 0 {
			fmt.Fprintf(f, "func %s(%s []float64, %s int) (%s []float64, err error) {\n",
				strings.ToUpper(indicator.Identifier),
				strings.Join(inputs, ", "),
				strings.Join(options, ", "),
				strings.Join(outputs, ", "),
			)
		} else {
			fmt.Fprintf(f, "func %s(%s []float64) (%s []float64, err error) {\n",
				strings.ToUpper(indicator.Identifier),
				strings.Join(inputs, ", "),
				strings.Join(outputs, ", "),
			)
		}

		fmt.Fprintf(f, "\tinput_length := len(%s)\n", inputs[0])

		if indicator.Options > 0 {
			floatOptions := []string{}
			for _, o := range options {
				floatOptions = append(floatOptions, fmt.Sprintf("float64(%s)", o))
			}
			fmt.Fprintf(f, "\toptions := []float64{%s}\n", strings.Join(floatOptions, ", "))
			fmt.Fprintf(f, "\toption_input := (*C.double)(&options[0])\n")
			fmt.Fprintf(f, "\tstart, err := C.ti_%s_start(option_input)\n", indicator.Identifier)
			fmt.Fprintf(f, "\tif err != nil {\n")
			fmt.Fprintf(f, "\t\treturn\n")
			fmt.Fprintf(f, "\t}\n")
		} else {
			fmt.Fprintf(f, "\toptions := []float64{0}\n")
			fmt.Fprintf(f, "\tstart := 0\n")
		}

		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "\tall_input_data := newIndicatorData(input_length, %d)\n", indicator.Inputs)
		fmt.Fprintf(f, "\tall_input_data.Set([][]float64{%s})\n", strings.Join(inputs, ","))
		fmt.Fprintf(f, "\tdefer all_input_data.Destroy()\n")
		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "\toutput_length := input_length - int(start)\n")
		fmt.Fprintf(f, "\tall_output_data := newIndicatorData(output_length, %d)\n", indicator.Outputs)
		fmt.Fprintf(f, "\tdefer all_output_data.Destroy()\n")

		fmt.Fprintf(f, "\tret, err := C.ti_%s(\n"+
			"\t\t(C.int)(input_length),\n"+
			"\t\t(**C.double)(all_input_data.buffer),\n"+
			"\t\t(*C.double)(&options[0]),\n"+
			"\t\t(**C.double)(all_output_data.buffer),\n"+
			"\t\t)\n", indicator.Identifier)
		fmt.Fprintf(f, "\n")

		fmt.Fprintf(f, "\tif err != nil {\n")
		fmt.Fprintf(f, "\t\treturn\n")
		fmt.Fprintf(f, "\t}\n")
		fmt.Fprintf(f, "\tif ret != C.TI_OKAY {\n")
		fmt.Fprintf(f, "\t\terr = fmt.Errorf(\"ret = %%d\", ret)\n")
		fmt.Fprintf(f, "\t\treturn\n")
		fmt.Fprintf(f, "\t}\n")

		// unpack results
		fmt.Fprintf(f, "\toutputs := all_output_data.Get()\n")
		for idx, o := range outputs {
			fmt.Fprintf(f, "\t%s = outputs[%d]\n", o, idx)
		}

		fmt.Fprintf(f, "\treturn\n")
		fmt.Fprintf(f, "}\n")
	}
	return nil
}

func main() {
	indicators := index()
	generateIndicators(indicators)
}
