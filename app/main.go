package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	bootstrap "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	router "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	tls "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	patch "github.com/geraldo-labs/merge-struct"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
)

var opt = options{
	Domain:        "*",
	FullChain:     "fullchain.pem",
	PrivKey:       "privkey.pem",
	HAPort:        8123,
	ExposeMetrics: false,
	OutputFile:    "/data/envoy.json",
}

type options struct {
	Domain        string `json:"domain"`
	FullChain     string `json:"ssl_fullchain"`
	PrivKey       string `json:"ssl_privkey"`
	HAPort        int    `json:"ha_port"`
	ExposeMetrics bool   `json:"expose_metrics"`
	OutputFile    string `json:"output_file"`
}

func init() {
	defer func() {
		out, err := json.Marshal(opt)
		if err != nil {
			panic(err)
		}
		log.Printf("using config options: %s", out)
	}()

	configPath := "/data/options.json"

	// allow override for testing
	if os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Printf("could not open config file, continue with defaults: %v", err)
		return
	}
	defer configFile.Close()

	var fileOpts options
	if err := json.NewDecoder(configFile).Decode(&fileOpts); err != nil {
		log.Printf("could not parse config file, continue with defaults: %v", err)
		return
	}

	patch.Struct(&opt, fileOpts)
}

func main() {
	routerConfig, _ := ptypes.MarshalAny(&router.Router{})
	routes := []*route.Route{
		{
			Match: &route.RouteMatch{
				PathSpecifier: &route.RouteMatch_Prefix{
					Prefix: "/",
				},
			},
			Action: &route.Route_Route{
				Route: &route.RouteAction{
					ClusterSpecifier: &route.RouteAction_Cluster{
						Cluster: "service_homeassistant",
					},
				},
			},
		},
	}

	if opt.ExposeMetrics {
		metricsRoute := &route.Route{
			Match: &route.RouteMatch{
				PathSpecifier: &route.RouteMatch_Path{
					Path: "/internal/metrics",
				},
			},
			Action: &route.Route_Route{
				Route: &route.RouteAction{
					ClusterSpecifier: &route.RouteAction_Cluster{
						Cluster: "admin_interface",
					},
					PrefixRewrite: "/stats/prometheus",
				},
			},
		}
		routes = append([]*route.Route{metricsRoute}, routes...)
	}

	// HttpConnectionManager config
	httpManager := &hcm.HttpConnectionManager{
		StatPrefix: "ingress_http",
		UpgradeConfigs: []*hcm.HttpConnectionManager_UpgradeConfig{
			{
				UpgradeType: "websocket",
			},
		},
		RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{
			RouteConfig: &route.RouteConfiguration{
				Name: "local_route",
				VirtualHosts: []*route.VirtualHost{
					{
						Name:    "local_service",
						Domains: []string{opt.Domain},
						Routes:  routes,
					},
				},
			},
		},
		HttpFilters: []*hcm.HttpFilter{
			{
				Name: wellknown.Router,
				ConfigType: &hcm.HttpFilter_TypedConfig{
					TypedConfig: routerConfig,
				},
			},
		},
	}

	managerConfig, _ := ptypes.MarshalAny(httpManager)

	// Tls context config
	tlsContext := &tls.DownstreamTlsContext{
		CommonTlsContext: &tls.CommonTlsContext{
			TlsCertificates: []*tls.TlsCertificate{
				{
					CertificateChain: &core.DataSource{
						Specifier: &core.DataSource_Filename{
							Filename: "/ssl/" + opt.FullChain,
						},
					},
					PrivateKey: &core.DataSource{
						Specifier: &core.DataSource_Filename{
							Filename: "/ssl/" + opt.PrivKey,
						},
					},
				},
			},
		},
	}

	tlsConfig, _ := ptypes.MarshalAny(tlsContext)

	// Listeners
	listeners := []*listener.Listener{
		{
			Name: "listener_0",
			Address: &core.Address{
				Address: &core.Address_SocketAddress{
					SocketAddress: &core.SocketAddress{
						Address: "0.0.0.0",
						PortSpecifier: &core.SocketAddress_PortValue{
							PortValue: 443,
						},
					},
				},
			},
			FilterChains: []*listener.FilterChain{
				{
					Filters: []*listener.Filter{
						{
							Name: wellknown.HTTPConnectionManager,
							ConfigType: &listener.Filter_TypedConfig{
								TypedConfig: managerConfig,
							},
						},
					},
					TransportSocket: &core.TransportSocket{
						Name: wellknown.TransportSocketTls,
						ConfigType: &core.TransportSocket_TypedConfig{
							TypedConfig: tlsConfig,
						},
					},
				},
			},
		},
	}

	admin := &bootstrap.Admin{
		Address: &core.Address{
			Address: &core.Address_Pipe{
				Pipe: &core.Pipe{
					Path: "/tmp/admin.sock",
				},
			},
		},
	}

	// Clusters
	clusters := []*cluster.Cluster{
		{
			Name: "service_homeassistant",
			ClusterDiscoveryType: &cluster.Cluster_Type{
				Type: cluster.Cluster_LOGICAL_DNS,
			},
			ConnectTimeout: ptypes.DurationProto(250 * time.Millisecond),
			LoadAssignment: &endpoint.ClusterLoadAssignment{
				ClusterName: "service_homeassistant",
				Endpoints: []*endpoint.LocalityLbEndpoints{
					{
						LbEndpoints: []*endpoint.LbEndpoint{
							{
								HostIdentifier: &endpoint.LbEndpoint_Endpoint{
									Endpoint: &endpoint.Endpoint{
										Address: &core.Address{
											Address: &core.Address_SocketAddress{
												SocketAddress: &core.SocketAddress{
													Address: "homeassistant.local.hass.io",
													PortSpecifier: &core.SocketAddress_PortValue{
														PortValue: uint32(opt.HAPort),
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			LbPolicy:        cluster.Cluster_ROUND_ROBIN,
			DnsLookupFamily: cluster.Cluster_V4_ONLY,
		},
		{
			Name: "admin_interface",
			ClusterDiscoveryType: &cluster.Cluster_Type{
				Type: cluster.Cluster_STATIC,
			},
			ConnectTimeout: ptypes.DurationProto(250 * time.Millisecond),
			LoadAssignment: &endpoint.ClusterLoadAssignment{
				ClusterName: "admin_interface",
				Endpoints: []*endpoint.LocalityLbEndpoints{
					{
						LbEndpoints: []*endpoint.LbEndpoint{
							{
								HostIdentifier: &endpoint.LbEndpoint_Endpoint{
									Endpoint: &endpoint.Endpoint{
										Address: &core.Address{
											Address: &core.Address_Pipe{
												Pipe: &core.Pipe{
													Path: "/tmp/admin.sock",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			LbPolicy: cluster.Cluster_ROUND_ROBIN,
		},
	}

	// Construct the overall config
	config := &bootstrap.Bootstrap{
		StaticResources: &bootstrap.Bootstrap_StaticResources{
			Listeners: listeners,
			Clusters:  clusters,
		},
		Admin: admin,
	}

	marshaler := &jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(config)
	if err != nil {
		log.Fatalf("Failed to marshal config to JSON: %v", err)
	}

	err = ioutil.WriteFile(opt.OutputFile, []byte(str), os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to write config to file: %v", err)
	}
}
