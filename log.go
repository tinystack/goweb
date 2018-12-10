// Copyright 2018 tinystack Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package goweb

import (
    "io"
)

type (
    Logger interface {
        Output(level int, s string) error
        SetOutput(w io.Writer)
        GetOutput() io.Writer
        Debug(s string, v ...interface{})
        Info(s string, v ...interface{})
        Notice(s string, v ...interface{})
        Warning(s string, v ...interface{})
        Error(s string, v ...interface{})
        Panic(s string, v ...interface{})
    }
)

