package main

import (
	"context"
	"fmt"
	"errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func serverApp(ctx context.Context) error {
	mux:=http.NewServeMux()
	srv:=&http.Server{Addr: "0.0.0.0:8000",Handler: mux}
	mux.HandleFunc("/",func(resp http.ResponseWriter,req *http.Request) {
		time.Sleep(time.Microsecond)
		fmt.Println(resp,"hello new conn")
	})
	go func(){
		timeoutCtx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
		select {
		case<-ctx.Done():
			srv.Shutdown(timeoutCtx)
			fmt.Println("App Shutdown")
		}
	}()
	fmt.Println("App Start")
	return srv.ListenAndServe()

}

func serverDebug(ctx context.Context) error {
	srv:=&http.Server{Addr: "127.0.0.1:8081",Handler: http.DefaultServeMux}
	go func(){
		timeoutCtx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
		select {
		case<-ctx.Done():
			srv.Shutdown(timeoutCtx)
			fmt.Println("App Shutdown")
		}
	}()
	fmt.Println("App Debug")
	return srv.ListenAndServe()
}

func serverSignal(ctx context.Context) error {
	signalCh:=make(chan os.Signal)
	signal.Notify(signalCh,os.Interrupt,syscall.SIGINT,syscall.SIGHUP,syscall.SIGTERM,syscall.SIGQUIT,syscall.SIGKILL)
	select{
	case <-ctx.Done():
		return nil
	case <-signalCh:
		return errors.New("err:stop signal")
	}
}

func main(){
	group,ctx:=errgroup.WithContext(context.Background())
	//start app
	group.Go(func()error{
		return serverApp(ctx)
	})

	group.Go(func() error{
		return serverDebug(ctx)
	})

	group.Go(func() error {
		return serverSignal(ctx)
	})
	if err:=group.Wait();err!=nil {
		fmt.Printf("server error:%v\n",err)
	}
}