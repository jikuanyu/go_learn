package handler

import (
	"context"
	pb "helloworld/protovcom"
)

type Greeter struct {
}

func (g *Greeter) Hello(context context.Context, in *pb.Request, out *pb.Response) error {
	out.Msg = "hello ,你好" + in.GetName()
	return nil
}
