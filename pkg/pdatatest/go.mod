module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest

go 1.19

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.76.3
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.76.3
	github.com/stretchr/testify v1.8.2
	go.opentelemetry.io/collector/pdata v1.0.0-rcv0011.0.20230508160606-833cbdaf75cc
	go.uber.org/multierr v1.11.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../pdatautil

retract (
	v0.76.2
	v0.76.1
)
