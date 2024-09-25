package log

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	type args struct {
		format string
		level  string
	}

	type expectation struct {
		level  logrus.Level
		format logrus.Formatter
	}

	tests := []struct {
		name     string
		args     args
		expected expectation
	}{
		{
			name: "success debug level",
			args: args{
				format: "json",
				level:  "debug",
			},
			expected: expectation{
				level: logrus.DebugLevel,
			},
		},
		{
			name: "success warning level",
			args: args{
				format: "json",
				level:  "warning",
			},
			expected: expectation{
				level: logrus.WarnLevel,
			},
		},
		{
			name: "success panic level",
			args: args{
				format: "json",
				level:  "panic",
			},
			expected: expectation{
				level: logrus.PanicLevel,
			},
		},
		{
			name: "success fatal level",
			args: args{
				format: "json",
				level:  "fatal",
			},
			expected: expectation{
				level: logrus.FatalLevel,
			},
		},
		{
			name: "success error level",
			args: args{
				format: "json",
				level:  "error",
			},
			expected: expectation{
				level: logrus.ErrorLevel,
			},
		},
		{
			name: "success info level",
			args: args{
				format: "json",
				level:  "info",
			},
			expected: expectation{
				level: logrus.InfoLevel,
			},
		},
		{
			name: "success use default",
			args: args{
				format: "json",
				level:  "invalid",
			},
			expected: expectation{
				level: logrus.InfoLevel,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Configure(tt.args.format, tt.args.level)

			assert.Equal(t, logrus.GetLevel(), tt.expected.level, "expected %s got %s", logrus.GetLevel(), tt.expected.level)
		})
	}
}

func TestGetLevel(t *testing.T) {
	test := []struct {
		name  string
		level logrus.Level
		want  string
	}{
		{
			name:  "success",
			level: logrus.DebugLevel,
			want:  "DEBUG",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			logrus.SetLevel(tt.level)
			if got := GetLevel(); got != tt.want {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetLevel(t *testing.T) {
	type args struct {
		level string
	}
	test := []struct {
		name      string
		args      args
		wantLevel string
	}{
		{
			name: "success",
			args: args{
				level: "debug",
			},
			wantLevel: "DEBUG",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.args.level)
			assert.Equal(t, GetLevel(), tt.wantLevel, "got %s expected %s", logrus.GetLevel(), tt.wantLevel)
		})
	}
}

func TestGetLogger(t *testing.T) {
	type args struct {
		ctx    context.Context
		pkg    string
		fnName string
	}
	test := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				pkg:    "package",
				fnName: "fName",
			},
			wantErr: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			l := GetLogger(tt.args.ctx, tt.args.pkg, tt.args.fnName)
			if !tt.wantErr {
				assert.NotNil(t, l)
			} else {
				assert.Nil(t, l)
			}
		})
	}
}

func TestWithContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	test := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			l := WithContext(tt.args.ctx)
			if !tt.wantErr {
				assert.NotNil(t, l)
			} else {
				assert.Nil(t, l)
			}
		})
	}
}

func TestUsage(t *testing.T) {
	ctx := context.Background()
	Configure("json", "debug")
	logrus.SetReportCaller(true)
	WithContext(ctx).Debug("test")
}

func TestLogLevelField(t *testing.T) {
	assertTest := assert.New(t)
	ctx := context.Background()
	Configure("json", "debug")
	logTest := GetLogger(ctx, "service", "PatchAWB")
	assertTest.Equal("DEBUG", logTest.Data["level"])
}
