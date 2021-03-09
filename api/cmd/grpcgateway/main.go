package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	proto "github.com/khoerling/flux/api/lib/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)
var serveSwaggerJSON = serveFileHandler("lib/swagger/swagger.json", "application/json")
var serveSwaggerUI = serveFileHandler("lib/swagger/swagger-ui.html", "text/html")

// Maximum upload of 25 MB
const maxUploadSizeBytes = 1024 * 1024 * 25

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterFluxHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	mux.HandlePath("GET", "/swagger.json", serveSwaggerJSON)
	mux.HandlePath("GET", "/swagger", serveSwaggerUI)

	conn, err := grpc.Dial(*grpcServerEndpoint, grpc.WithInsecure(), grpc.WithMaxMsgSize(maxUploadSizeBytes))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewFluxClient(conn)

	// Upload translator for grpc (accept multipart on frontend)
	mux.HandlePath("POST", "/upload", uploadFileHandler(ctx, client))

	return http.ListenAndServe(apiPort(), allowCORS(mux))
}

// https://github.com/rephus/grpc-gateway-example/blob/master/main.go
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func apiPort() string {
	apiPort := os.Getenv("PORT")
	if apiPort == "" {
		apiPort = "8081"
	}
	log.Println("> listening on port " + apiPort)
	return ":" + apiPort
}

func serveFileHandler(path string, mimeType string) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Header().Add("content-type", mimeType)
		http.ServeFile(w, r, path)
	}
}

func uploadFileHandler(ctx context.Context, flux proto.FluxClient) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

		r.ParseMultipartForm(maxUploadSizeBytes)

		// Get handler for filename, size and headers
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		blob := make([]byte, maxUploadSizeBytes)
		n, err := file.Read(blob)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		ctx := metadata.NewIncomingContext(ctx, metadata.MD{
			"authorization": []string{handler.Header.Get("authorization")},
		})
		resp, err := flux.UploadFile(ctx, &proto.UploadFileRequest{
			Filename: handler.Filename,
			MimeType: handler.Header.Get("content-type"),
			Size:     int32(n),
			Body:     blob[:n],
		})
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		log.Println("resp: ", resp)

		w.Header().Add("content-type", "application/json")

		out, err := json.Marshal(map[string]string{
			"fileId": resp.FileId,
		})
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(out)
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
