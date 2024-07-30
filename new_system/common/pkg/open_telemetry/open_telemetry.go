package open_telemetry

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/dollarkillerx/common/pkg/config"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var Tracer = otel.Tracer("github.com")

// InitLog 初始化日志 必须异步调用
func InitLog(conf config.OpenTelemetryLogsConfig) {
	time.Sleep(time.Second)
	open, err := os.Open(conf.File)
	if err != nil {
		log.Error().Msgf("[OpenTelemetry] Failed to open log file: %s %s", err, conf.File)
		panic(err)
	}
	defer open.Close()

	client := resty.New().SetDisableWarn(true)

	reader := bufio.NewReader(open)
	open.Seek(0, io.SeekEnd)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// 如果到达文件末尾，则等待新内容
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			return
		}

		resp, err := client.R().
			SetHeader("Accept", "application/json").
			SetBasicAuth(conf.User, conf.Password).
			SetBody([]byte(strings.TrimSpace(line))).
			Post(conf.HTTPEndpoint)
		if err != nil {
			fmt.Printf("[Internal Error OpenTelemetry]: %s \n", err)
			continue
		}

		if resp.StatusCode() != 200 {
			fmt.Printf("[Internal Error OpenTelemetry]: %s \n", resp.String())
			continue
		}
	}
}

func InitTracerHTTP(conf config.OpenTelemetryTracesConfig) *sdktrace.TracerProvider {
	Tracer = otel.Tracer(conf.ServerName)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	otlptracehttp.NewClient()

	otlpHTTPExporter, err := otlptracehttp.New(context.TODO(),
		otlptracehttp.WithInsecure(), // use http & not https
		otlptracehttp.WithEndpoint(conf.HTTPEndpoint),
		otlptracehttp.WithURLPath(conf.Path),
		otlptracehttp.WithHeaders(map[string]string{
			"Authorization": conf.Authorization,
		}),
	)

	if err != nil {
		fmt.Println("Error creating HTTP OTLP exporter: ", err)
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(conf.ServerName),
		semconv.ServiceVersionKey.String("0.0.1"),
		//attribute.String("environment", "test"),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(otlpHTTPExporter),
	)
	otel.SetTracerProvider(tp)

	return tp
}
