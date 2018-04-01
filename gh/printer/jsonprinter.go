// Copyright © 2018 Mikael Berthe <mikael@lilotux.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package printer

import (
	"encoding/json"
	"fmt"

	"github.com/McKael/ghreleasechecker/gh"
)

// JSONPrinter is a JSON printer (the default one)
type JSONPrinter struct {
}

// NewPrinterJSON returns a JSON printer
func NewPrinterJSON(options Options) (*JSONPrinter, error) {
	p := &JSONPrinter{}
	return p, nil
}

// PrintReleases displays a list of releases in JSON format
func (p *JSONPrinter) PrintReleases(rr []gh.Release) error {
	if len(rr) == 0 {
		return nil
	}
	bb, err := json.Marshal(rr)
	if err != nil {
		return err
	}
	fmt.Println(string(bb))
	return nil
}
