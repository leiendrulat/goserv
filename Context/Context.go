package Contexter

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	//only 3rd parties

	. "github.com/logrusorgru/aurora"

)

var err error

var Start = time.Now()
var Durations = time.Now().Sub(Start)

//getting context
type Contexter struct {
	M      string
	U      *url.URL
	P      string
	B      io.ReadCloser
	Gb     func() (io.ReadCloser, error)
	Host   string
	Form   url.Values
	Cancel <-chan struct{}
	R      *http.Response
	H      http.Header
	D      time.Duration
	I      string
}

//used to shorten use of Contexter
var CC Contexter

//initializing context
func AddContext(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Start := time.Now()
		Duration := time.Now().Sub(Start)

		CC = Contexter{
			M:      r.WithContext(ctx).Method,
			U:      r.WithContext(ctx).URL,
			P:      r.WithContext(ctx).Proto,
			B:      r.WithContext(ctx).Body,
			Host:   r.WithContext(ctx).Host,
			Form:   r.WithContext(ctx).Form,
			Cancel: r.WithContext(ctx).Cancel,
			R:      r.WithContext(ctx).Response,
			D:      Duration,
			H:      r.WithContext(ctx).Header,
			//I:      Clients.ReadUserIP(r),
		}

		fmt.Println(Blue("/ʕ◔ϖ◔ʔ/````````````````````````````````````````````"))
		fmt.Printf("Method:%s\n - Status:%s\n - URL:%s - Body:%v\n - Host:%s\n - Form:%v\n - Cancel:%d\n - Response:%d\n - Dur:%02d-00:00\n - Cache-Control:%s - Accept:%s\n - IP:%s\n",
			Cyan(CC.M),
			Brown(r.Header.Get("X-Forwarded-Port")),
			Red(CC.U),
			Blue(CC.B),
			Yellow(CC.Host),
			BgRed(CC.Form),
			BgGreen(CC.Cancel),
			BgBrown(CC.R),
			BgMagenta(CC.D),
			Red(CC.H.Get("Cache-Control")),
			Blue(CC.H.Get("Accept")),
			Yellow(CC.I),
		)
		//this is to spit out the ctx.header in pieces cause its exhaustive
		// for k, v := range CC.H {
		// 	fmt.Println("\n", k, v)
		// }
		cookie, _ := r.Cookie("username")

		if cookie != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {

			if err != nil {
				// Error occurred while parsing request body
				w.WriteHeader(http.StatusBadRequest)

				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
