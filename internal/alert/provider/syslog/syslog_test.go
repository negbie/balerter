package syslog

import (
	"github.com/balerter/balerter/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log/syslog"
	"testing"
)

func TestNew(t *testing.T) {
	s, err := New(config.ChannelSyslog{
		Name: "foo",
	}, nil)

	require.NoError(t, err)
	assert.IsType(t, &Syslog{}, s)
	assert.Equal(t, "foo", s.name)
}

func TestName(t *testing.T) {
	s := &Syslog{name: "foo"}
	assert.Equal(t, "foo", s.Name())
}

func Test_parsePriority(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want syslog.Priority
	}{
		{
			name: "empty",
			args: args{s: ""},
			want: syslog.LOG_EMERG,
		},
		{name: "severity ALERT", args: args{s: "ALERT"}, want: syslog.LOG_ALERT},
		{name: "severity CRIT", args: args{s: "CRIT"}, want: syslog.LOG_CRIT},
		{name: "severity ERR", args: args{s: "ERR"}, want: syslog.LOG_ERR},
		{name: "severity WARNING", args: args{s: "WARNING"}, want: syslog.LOG_WARNING},
		{name: "severity NOTICE", args: args{s: "NOTICE"}, want: syslog.LOG_NOTICE},
		{name: "severity INFO", args: args{s: "INFO"}, want: syslog.LOG_INFO},
		{name: "severity DEBUG", args: args{s: "DEBUG"}, want: syslog.LOG_DEBUG},
		{name: "severity and facility", args: args{s: "ALERT|KERN"}, want: syslog.LOG_ALERT | syslog.LOG_KERN},
		{name: "severity and facility", args: args{s: "ALERT|USER"}, want: syslog.LOG_ALERT | syslog.LOG_USER},
		{name: "severity and facility", args: args{s: "ALERT|MAIL"}, want: syslog.LOG_ALERT | syslog.LOG_MAIL},
		{name: "severity and facility", args: args{s: "ALERT|DAEMON"}, want: syslog.LOG_ALERT | syslog.LOG_DAEMON},
		{name: "severity and facility", args: args{s: "ALERT|AUTH"}, want: syslog.LOG_ALERT | syslog.LOG_AUTH},
		{name: "severity and facility", args: args{s: "ALERT|SYSLOG"}, want: syslog.LOG_ALERT | syslog.LOG_SYSLOG},
		{name: "severity and facility", args: args{s: "ALERT|LPR"}, want: syslog.LOG_ALERT | syslog.LOG_LPR},
		{name: "severity and facility", args: args{s: "ALERT|NEWS"}, want: syslog.LOG_ALERT | syslog.LOG_NEWS},
		{name: "severity and facility", args: args{s: "ALERT|UUCP"}, want: syslog.LOG_ALERT | syslog.LOG_UUCP},
		{name: "severity and facility", args: args{s: "ALERT|CRON"}, want: syslog.LOG_ALERT | syslog.LOG_CRON},
		{name: "severity and facility", args: args{s: "ALERT|AUTHPRIV"}, want: syslog.LOG_ALERT | syslog.LOG_AUTHPRIV},
		{name: "severity and facility", args: args{s: "ALERT|FTP"}, want: syslog.LOG_ALERT | syslog.LOG_FTP},
		{name: "severity and facility", args: args{s: "ALERT|LOCAL0"}, want: syslog.LOG_ALERT | syslog.LOG_LOCAL0},
		{name: "severity and facility", args: args{s: "ALERT|LOCAL1"}, want: syslog.LOG_ALERT | syslog.LOG_LOCAL1},
		{name: "severity and facility", args: args{s: "ALERT|LOCAL2"}, want: syslog.LOG_ALERT | syslog.LOG_LOCAL2},
		{name: "severity and facility", args: args{s: "ALERT|LOCAL3"}, want: syslog.LOG_ALERT | syslog.LOG_LOCAL3},
		{name: "severity and facility", args: args{s: "ALERT|LOCAL4"}, want: syslog.LOG_ALERT | syslog.LOG_LOCAL4},
		{name: "severity and facility", args: args{s: "ALERT|LOCAL5"}, want: syslog.LOG_ALERT | syslog.LOG_LOCAL5},
		{name: "severity and facility", args: args{s: "ALERT|LOCAL6"}, want: syslog.LOG_ALERT | syslog.LOG_LOCAL6},
		{name: "severity and facility", args: args{s: "ALERT|LOCAL7"}, want: syslog.LOG_ALERT | syslog.LOG_LOCAL7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePriority(tt.args.s); got != tt.want {
				t.Errorf("parsePriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
