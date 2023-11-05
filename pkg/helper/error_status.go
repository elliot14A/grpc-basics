package helper

import "google.golang.org/grpc/status"

func ErrorStatusConverter(err error) (string, string) {
	s := status.Convert(err)
	return s.Code().String(), s.Message()
}
