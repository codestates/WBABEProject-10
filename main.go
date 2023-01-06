package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lecture/WBABEProject-10/logger"
	rt "lecture/WBABEProject-10/router"

	"lecture/WBABEProject-10/model"

	ctl "lecture/WBABEProject-10/controller"

	"golang.org/x/sync/errgroup"

	conf "lecture/WBABEProject-10/config"
)

var (
	g errgroup.Group
)

func main() {
	var configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")
	/* [코드리뷰]
	 * 시스템과 관련된 config를 main 함수 첫번째에서 잘 가져와 주셨습니다.
	 * command 라인에서 config를 별도로 지정하여 환경에 따라 다른 시스템을 다르게 실행할 수 있게된 좋은 코드입니다.
	 */
	flag.Parse()
	cf := conf.NewConfig(*configFlag)

	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	logger.Debug("ready server....")

	if mod, err := model.NewModel(); err != nil {
		panic(err)
	} else if controller, err := ctl.NewCTL(mod); err != nil {
		panic(err)
	} else if rt, err := rt.NewRouter(controller); err != nil {
		panic(err)
	} else {
		mapi := &http.Server{
			Addr:           ":8080",
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		g.Go(func() error {
			return mapi.ListenAndServe()
		})
		stopSig := make(chan os.Signal) //chan 선언
		// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
		signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM)
		<-stopSig //메세지 등록

		// 해당 context 타임아웃 설정, 5초후 server stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			fmt.Println("timeout 5 seconds.")
		}
		fmt.Println("Server stop")
	}

	g.Wait()

}
