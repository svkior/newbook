package main

import (
	//"log"
	"net/http"
)

type Middleware []http.Handler

func (m *Middleware) Add(handler http.Handler) {
	*m = append(*m, handler)
}

type MiddlewareResponseWriter struct {
	http.ResponseWriter
	written bool
}

func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWriter {
	return &MiddlewareResponseWriter{ResponseWriter: w}
}

func (w *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(bytes)
}

func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw := NewMiddlewareResponseWriter(w)

	for _, handler := range m {
		//log.Printf("Handle %p, IN mw : ", handler, mw.written)
		handler.ServeHTTP(mw, r)
		//log.Printf("Handle %p, OUT mw : ", handler, mw.written)
		if mw.written {
			return
		}
	}
	http.NotFound(w, r)
}
