/*
 * Copyright (c) 2014 by Farsight Security, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package framestream

import (
	"io"
	"time"
)

type EncoderOptions struct {
	ContentType   []byte
	Bidirectional bool
	Timeout       time.Duration
}

type Encoder struct {
	*Writer
}

func NewEncoder(w io.Writer, opt *EncoderOptions) (enc *Encoder, err error) {
	if opt == nil {
		opt = &EncoderOptions{}
	}
	wopt := &WriterOptions{
		Bidirectional: opt.Bidirectional,
		Timeout:       opt.Timeout,
	}
	if opt.ContentType != nil {
		wopt.ContentTypes = append(wopt.ContentTypes, opt.ContentType)
	}
	writer, err := NewWriter(w, wopt)
	if err != nil {
		return nil, err
	}
	return &Encoder{Writer: writer}, nil
}
