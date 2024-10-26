package main

import (
	"context"
	"fmt"
	"grpc-demo/invoicer"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type invoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (server *invoicerServer) Create(ctx context.Context, request *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  request.From,
		Docx: request.To,
	}, nil
}

func (s *invoicerServer) BiDirectionalTest(stream invoicer.Invoicer_BiDirectionalTestServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			return nil
		}

		if err != nil {
			log.Fatalln(err)
			return err
		}

		log.Printf("Received BiDirectionalTest request: %v", req)

		resp := &invoicer.CreateResponse{
			Pdf:  "Sample PDF content",
			Docx: "Sample DOCX content",
		}

		if err := stream.Send(resp); err != nil {
			log.Fatalln(err)
			return err
		}
	}
}

func (s *invoicerServer) ServerStreamTest(req *invoicer.CreateRequest, stream invoicer.Invoicer_ServerStreamTestServer) error {

	for i := 0; i < 10; i++ {
		stream.Send(&invoicer.CreateResponse{
			Pdf:  req.From,
			Docx: req.To,
		})
	}

	return nil
}

func (s *invoicerServer) ClientStreamTest(stream invoicer.Invoicer_ClientStreamTestServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &invoicer.CreateResponse{
				Pdf:  "Sample PDF content",
				Docx: "Sample DOCX content",
			}

			if err := stream.SendAndClose(resp); err != nil {
				log.Fatalln(err)
				return err
			}

			return nil
		}

		if req != nil {
			fmt.Println(req)
		}

	}
}

func main() {

	listener, httpError := net.Listen("tcp", ":5000")

	if httpError != nil {
		log.Fatalln("ERROR")
	}

	grpcServer := grpc.NewServer()

	invoicerService := &invoicerServer{}
	invoicer.RegisterInvoicerServer(grpcServer, invoicerService)
	grpcError := grpcServer.Serve(listener)

	if grpcError != nil {
		log.Fatalln("GRPC ERROR")
	}
}
