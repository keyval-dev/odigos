package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/odigos-io/odigos/common/consts"
	"github.com/odigos-io/odigos/destinations"
	"github.com/odigos-io/odigos/k8sutils/pkg/env"

	"github.com/gin-contrib/cors"

	v1alpha1 "github.com/odigos-io/odigos/api/odigos/v1alpha1"
	"github.com/odigos-io/odigos/frontend/endpoints/actions"
	"github.com/odigos-io/odigos/frontend/endpoints/sse"
	"github.com/odigos-io/odigos/frontend/kube"
	"github.com/odigos-io/odigos/frontend/version"

	"github.com/odigos-io/odigos/frontend/endpoints"

	"github.com/gin-gonic/gin"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

const (
	defaultPort = 3000
)

type Flags struct {
	Version    bool
	Address    string
	Port       int
	Debug      bool
	KubeConfig string
	Namespace  string
}

// type SSEMessage struct {
// 	Time string `json:"time"`
// 	Type string `json:"type"`
// }

//go:embed all:webapp/out/*
var uiFS embed.FS

func parseFlags() Flags {
	defaultKubeConfig := env.GetDefaultKubeConfigPath()

	var flags Flags
	flag.BoolVar(&flags.Version, "version", false, "Print Odigos UI version.")
	flag.StringVar(&flags.Address, "address", "localhost", "Address to listen on")
	flag.IntVar(&flags.Port, "port", defaultPort, "Port to listen on")
	flag.BoolVar(&flags.Debug, "debug", false, "Enable debug mode")
	flag.StringVar(&flags.KubeConfig, "kubeconfig", defaultKubeConfig, "Path to kubeconfig file")
	flag.StringVar(&flags.Namespace, "namespace", consts.DefaultNamespace, "Kubernetes namespace where Odigos is installed")
	flag.Parse()
	return flags
}

func initKubernetesClient(flags *Flags) error {
	client, err := kube.CreateClient(flags.KubeConfig)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes client: %w", err)
	}

	kube.SetDefaultClient(client)
	return nil
}

func startHTTPServer(flags *Flags) (*gin.Engine, error) {
	var r *gin.Engine
	if flags.Debug {
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	}

	// Enable CORS
	r.Use(cors.Default())

	// Serve React app
	dist, err := fs.Sub(uiFS, "webapp/out")
	if err != nil {
		return nil, fmt.Errorf("error reading webapp/out directory: %s", err)
	}

	// Serve React app if page not found serve index.html
	r.NoRoute(gin.WrapH(httpFileServerWith404(http.FS(dist))))

	// Serve API
	apis := r.Group("/api")
	{
		apis.GET("/namespaces", func(c *gin.Context) { endpoints.GetNamespaces(c, flags.Namespace) })
		apis.POST("/namespaces", endpoints.PersistNamespaces)

		apis.GET("/sources", func(c *gin.Context) { endpoints.GetSources(c, flags.Namespace) })
		apis.GET("/sources/namespace/:namespace/kind/:kind/name/:name", endpoints.GetSource)
		apis.DELETE("/sources/namespace/:namespace/kind/:kind/name/:name", endpoints.DeleteSource)
		apis.PATCH("/sources/namespace/:namespace/kind/:kind/name/:name", endpoints.PatchSource)

		apis.GET("/applications/:namespace", endpoints.GetApplicationsInNamespace)
		apis.GET("/config", endpoints.GetConfig)
		apis.GET("/destination-types", endpoints.GetDestinationTypes)
		apis.GET("/destination-types/:type", endpoints.GetDestinationTypeDetails)
		apis.GET("/destinations", func(c *gin.Context) { endpoints.GetDestinations(c, flags.Namespace) })
		apis.GET("/destinations/:id", func(c *gin.Context) { endpoints.GetDestinationById(c, flags.Namespace) })
		apis.POST("/destinations", func(c *gin.Context) { endpoints.CreateNewDestination(c, flags.Namespace) })
		apis.PUT("/destinations/:id", func(c *gin.Context) { endpoints.UpdateExistingDestination(c, flags.Namespace) })
		apis.DELETE("/destinations/:id", func(c *gin.Context) { endpoints.DeleteDestination(c, flags.Namespace) })

		apis.GET("/actions", func(c *gin.Context) { actions.GetActions(c, flags.Namespace) })

		// AddClusterInfo
		apis.GET("/actions/types/AddClusterInfo/:id", func(c *gin.Context) { actions.GetAddClusterInfo(c, flags.Namespace, c.Param("id")) })
		apis.POST("/actions/types/AddClusterInfo", func(c *gin.Context) { actions.CreateAddClusterInfo(c, flags.Namespace) })
		apis.PUT("/actions/types/AddClusterInfo/:id", func(c *gin.Context) { actions.UpdateAddClusterInfo(c, flags.Namespace, c.Param("id")) })
		apis.DELETE("/actions/types/AddClusterInfo/:id", func(c *gin.Context) { actions.DeleteAddClusterInfo(c, flags.Namespace, c.Param("id")) })

		// DeleteAttribute
		apis.GET("/actions/types/DeleteAttribute/:id", func(c *gin.Context) { actions.GetDeleteAttribute(c, flags.Namespace, c.Param("id")) })
		apis.POST("/actions/types/DeleteAttribute", func(c *gin.Context) { actions.CreateDeleteAttribute(c, flags.Namespace) })
		apis.PUT("/actions/types/DeleteAttribute/:id", func(c *gin.Context) { actions.UpdateDeleteAttribute(c, flags.Namespace, c.Param("id")) })
		apis.DELETE("/actions/types/DeleteAttribute/:id", func(c *gin.Context) { actions.DeleteDeleteAttribute(c, flags.Namespace, c.Param("id")) })

		// RenameAttribute
		apis.GET("/actions/types/RenameAttribute/:id", func(c *gin.Context) { actions.GetRenameAttribute(c, flags.Namespace, c.Param("id")) })
		apis.POST("/actions/types/RenameAttribute", func(c *gin.Context) { actions.CreateRenameAttribute(c, flags.Namespace) })
		apis.PUT("/actions/types/RenameAttribute/:id", func(c *gin.Context) { actions.UpdateRenameAttribute(c, flags.Namespace, c.Param("id")) })
		apis.DELETE("/actions/types/RenameAttribute/:id", func(c *gin.Context) { actions.DeleteRenameAttribute(c, flags.Namespace, c.Param("id")) })
	}

	return r, nil
}

func httpFileServerWith404(fs http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(r.URL.Path)
		if err != nil {
			// Serve index.html
			r.URL.Path = "/"
		}
		http.FileServer(fs).ServeHTTP(w, r)
	})
}

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags := parseFlags()

	if flags.Version {
		fmt.Printf("version.Info{Version:'%s', GitCommit:'%s', BuildDate:'%s'}\n", version.OdigosVersion, version.OdigosCommit, version.OdigosDate)
		return
	}

	// Load destinations data
	err := destinations.Load()
	if err != nil {
		log.Fatalf("Error loading destinations data: %s", err)
	}

	// Connect to Kubernetes
	err = initKubernetesClient(&flags)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %s", err)
	}

	// Start server
	r, err := startHTTPServer(&flags)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	instrumentedApplicationWatcher, err := kube.DefaultClient.OdigosClient.InstrumentedApplications("").Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error creating watcher: %v", err)
	}

	// Handle the watch events for InstrumentedApplications
	go func() {
		ch := instrumentedApplicationWatcher.ResultChan()
		for event := range ch {
			switch event.Type {
			case watch.Added:
				fmt.Printf("New pod added: %s\n", event.Object.(*v1alpha1.InstrumentedApplication).Name)
				data := "InstrumentedApplication added successfully"
				sse.SendMessageToClient(sse.SSEMessage{Time: "2021-09-01 15:04:05", Type: "message", Data: data})
			case watch.Modified:
				fmt.Printf("Pod modified: %s\n", event.Object.(*v1alpha1.InstrumentedApplication).Name)
				conditions := event.Object.(*v1alpha1.InstrumentedApplication).Status.Conditions
				if len(conditions) == 0 {
					continue
				}

				// Send the message to the client
				for _, condition := range conditions {
					data := &condition.Message
					sse.SendMessageToClient(sse.SSEMessage{Time: "2021-09-01 15:04:05", Type: "message", Data: *data})
				}

			case watch.Deleted:
				fmt.Printf("Pod deleted: %s\n", event.Object.(*v1alpha1.InstrumentedApplication).Name)
			case watch.Error:
				fmt.Printf("Error watching pod: %v\n", event.Object)
			}
		}
	}()

	destinationWatcher, err := kube.DefaultClient.OdigosClient.Destinations(flags.Namespace).Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error creating watcher: %v", err)
	}

	// Handle the watch events for Destination
	go func() {
		ch := destinationWatcher.ResultChan()
		for event := range ch {
			switch event.Type {
			case watch.Added:
				fmt.Printf("New pod added: %s\n", event.Object.(*v1alpha1.Destination).Name)
				data := "Destination added successfully"
				sse.SendMessageToClient(sse.SSEMessage{Time: "2021-09-01 15:04:05", Type: "message", Data: data})

			case watch.Modified:
				fmt.Printf("Pod modified: %s\n", event.Object.(*v1alpha1.Destination).Name)
				conditions := event.Object.(*v1alpha1.Destination).Status.Conditions
				if len(conditions) == 0 {
					continue
				}

				// Send the message to the client
				for _, condition := range conditions {
					data := &condition.Message
					sse.SendMessageToClient(sse.SSEMessage{Time: "2021-09-01 15:04:05", Type: "message", Data: *data})
				}

			case watch.Deleted:
				fmt.Printf("Pod deleted: %s\n", event.Object.(*v1alpha1.Destination).Name)
			case watch.Error:
				fmt.Printf("Error watching pod: %v\n", event.Object)
			}
		}
	}()

	r.GET("/events", sse.HandleSSEConnections)

	// // Add SSE endpoint
	// r.GET("/events", func(c *gin.Context) {
	// 	c.Header("Content-Type", "text/event-stream")
	// 	c.Header("Cache-Control", "no-cache")
	// 	c.Header("Connection", "keep-alive")
	// 	c.Header("Transfer-Encoding", "chunked")

	// 	// Create a channel to send SSE messages
	// 	messageChan := make(chan SSEMessage)

	// 	// Start a goroutine to send SSE messages
	// 	go func() {
	// 		defer close(messageChan)
	// 		for {
	// 			// Create SSE message
	// 			message := SSEMessage{Time: time.Now().Format("2006-01-02 15:04:05"), Type: "message"}

	// 			// Marshal message to JSON
	// 			jsonData, err := json.Marshal(message)
	// 			if err != nil {
	// 				log.Printf("Error marshaling JSON: %s", err)
	// 				continue
	// 			}

	// 			// Send JSON data as SSE message
	// 			fmt.Fprintf(c.Writer, "data: %s\n\n", string(jsonData))
	// 			c.Writer.Flush()

	// 			// Wait for 2 seconds before sending the next message
	// 			time.Sleep(2 * time.Second)
	// 		}
	// 	}()

	// 	// Continuously send SSE messages to the client
	// 	for range c.Writer.CloseNotify() {
	// 		// Client connection closed, stop sending messages
	// 		return
	// 	}
	// })

	log.Println("Starting Odigos UI...")
	log.Printf("Odigos UI is available at: http://%s:%d", flags.Address, flags.Port)
	err = r.Run(fmt.Sprintf("%s:%d", flags.Address, flags.Port))
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
