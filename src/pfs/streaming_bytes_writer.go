package pfs

import (
	"bufio"
	"io"
)

func writeToStreamingBytesServer(reader io.Reader, streamingBytesServer streamingBytesServer) error {
	_, err := bufio.NewReader(reader).WriteTo(newStreamingBytesWriter(streamingBytesServer))
	return err
}

type streamingBytesServer interface {
	Send(streamingBytes *StreamingBytes) error
}

type streamingBytesWriter struct {
	streamingBytesServer streamingBytesServer
}

func newStreamingBytesWriter(
	streamingBytesServer streamingBytesServer,
) *streamingBytesWriter {
	return &streamingBytesWriter{
		streamingBytesServer,
	}
}

func (s *streamingBytesWriter) Write(p []byte) (int, error) {
	if err := s.streamingBytesServer.Send(
		&StreamingBytes{
			Payload: p,
		},
	); err != nil {
		return 0, err
	}
	return len(p), nil
}
