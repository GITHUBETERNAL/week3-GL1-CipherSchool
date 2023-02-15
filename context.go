package main

import "context"
func main() {
  ctc := context.Background()
  ctx, err := WithTimeout(ctx,time.Second*2)
  defer cancel()
  go doSomething(ctx)
  select {
    case v :=  <-ctx8.Done(): fmt.Println("timeline exceeded of 2 sec")
  }
  time.Sleep(time.Second *3)
}
func doSomething(ctx context.Context, chn chan<- chan int) {
  for {
    select {
      case <- ctx.Done():
      fmt.Println("timeout")
      return
      default:
      fmt.Println("Do something bakwaas")
    }
  }
}
// 2unc main() {
//   ctx := context.Background()
//   //seed some data in ctx
//   seedContext(ctx)
//   readCtx(cx)
// }

func readCtx(ctx cntext.Context() {
  value := ctx.Value("one")
  fmt.Println(value)
}
             
func seedContext(ctx context.Context) {
  context.WithValue(ctx, "key", "value")
  return ctx
}
  //store and pass the information across the different layer of the application
  // contesxt variable and pass it through the main -> router -> hanfler -> dbfunction
  // the ability to cancellation the job in between of execution
  
  // consume restful api anf if I as not done with it in within 2 mins I need to cancel that job
  
  
  
