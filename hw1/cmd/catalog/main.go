package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/LightAlykard/GO-BackEnd-2/hw1/internal/api/handler"
	"github.com/LightAlykard/GO-BackEnd-2/hw1/internal/api/server"
	"github.com/LightAlykard/GO-BackEnd-2/hw1/internal/app/models/community"
	"github.com/LightAlykard/GO-BackEnd-2/hw1/internal/app/models/user"
	"github.com/LightAlykard/GO-BackEnd-2/hw1/internal/app/starter"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	var ust user.UserStore
	var css community.CommunityStore

	a := starter.NewApp(ust, css)
	us := user.NewUsers(ust)
	cs := community.NewCommunities(css)

	h := handler.NewRouter(us, cs)
	srv := server.NewServer(":8000", h)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
