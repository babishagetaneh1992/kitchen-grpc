package main



func main () {
	grpcServer := NewGrpcServer(":50051")
	grpcServer.Run()
}



 