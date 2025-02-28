// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileexporter

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestCreateMetricsExporterError(t *testing.T) {
	cfg := &Config{
		FormatType: formatTypeJSON,
	}
	_, err := createMetricsExporter(
		context.Background(),
		exportertest.NewNopCreateSettings(),
		cfg)
	assert.Error(t, err)
}

func TestCreateMetricsExporter(t *testing.T) {
	cfg := &Config{
		FormatType: formatTypeJSON,
		Path:       tempFileName(t),
	}
	exp, err := createMetricsExporter(
		context.Background(),
		exportertest.NewNopCreateSettings(),
		cfg)
	assert.NoError(t, err)
	require.NotNil(t, exp)
}

func TestCreateTracesExporter(t *testing.T) {
	cfg := &Config{
		FormatType: formatTypeJSON,
		Path:       tempFileName(t),
	}
	exp, err := createTracesExporter(
		context.Background(),
		exportertest.NewNopCreateSettings(),
		cfg)
	assert.NoError(t, err)
	require.NotNil(t, exp)
}

func TestCreateTracesExporterError(t *testing.T) {
	cfg := &Config{
		FormatType: formatTypeJSON,
	}
	_, err := createTracesExporter(
		context.Background(),
		exportertest.NewNopCreateSettings(),
		cfg)
	assert.Error(t, err)
}

func TestCreateLogsExporter(t *testing.T) {
	cfg := &Config{
		FormatType: formatTypeJSON,
		Path:       tempFileName(t),
	}
	exp, err := createLogsExporter(
		context.Background(),
		exportertest.NewNopCreateSettings(),
		cfg)
	assert.NoError(t, err)
	require.NotNil(t, exp)
}

func TestCreateLogsExporterError(t *testing.T) {
	cfg := &Config{
		FormatType: formatTypeJSON,
	}
	_, err := createLogsExporter(
		context.Background(),
		exportertest.NewNopCreateSettings(),
		cfg)
	assert.Error(t, err)
}

func TestBuildFileWriter(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name     string
		args     args
		want     io.WriteCloser
		validate func(*testing.T, io.WriteCloser)
	}{
		{
			name: "single file",
			args: args{
				cfg: &Config{
					Path: tempFileName(t),
				},
			},
			validate: func(t *testing.T, closer io.WriteCloser) {
				_, ok := closer.(*bufferedWriteCloser)
				assert.Equal(t, true, ok)
			},
		},
		{
			name: "rotation file",
			args: args{
				cfg: &Config{
					Path: tempFileName(t),
					Rotation: &Rotation{
						MaxBackups: defaultMaxBackups,
					},
				},
			},
			validate: func(t *testing.T, closer io.WriteCloser) {
				bwc, ok := closer.(*bufferedWriteCloser)
				assert.Equal(t, true, ok)
				writer, ok := bwc.wrapped.(*lumberjack.Logger)
				assert.Equal(t, true, ok)
				assert.Equal(t, defaultMaxBackups, writer.MaxBackups)
			},
		},
		{
			name: "rotation file with user's configuration",
			args: args{
				cfg: &Config{
					Path: tempFileName(t),
					Rotation: &Rotation{
						MaxMegabytes: 30,
						MaxDays:      100,
						MaxBackups:   3,
						LocalTime:    true,
					},
				},
			},
			validate: func(t *testing.T, closer io.WriteCloser) {
				bwc, ok := closer.(*bufferedWriteCloser)
				assert.Equal(t, true, ok)
				writer, ok := bwc.wrapped.(*lumberjack.Logger)
				assert.Equal(t, true, ok)
				assert.Equal(t, 3, writer.MaxBackups)
				assert.Equal(t, 30, writer.MaxSize)
				assert.Equal(t, 100, writer.MaxAge)
				assert.Equal(t, true, writer.LocalTime)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildFileWriter(tt.args.cfg)
			assert.NoError(t, err)
			tt.validate(t, got)
		})
	}
}
