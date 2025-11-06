package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestIsTimeInRangeStr(t *testing.T) {
	type args struct {
		timeStr  string
		startStr string
		endStr   string
		layout   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				timeStr:  "2020-01-01 00:00:00",
				startStr: "2020-01-01 00:00:00",
				endStr:   "2020-01-01 23:59:59",
				layout:   "2006-01-02 15:04:05",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				timeStr:  "10:59:59",
				startStr: "00:00:00",
				endStr:   "23:59:59",
				layout:   "15:04:05",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsTimeInRangeStr(tt.args.timeStr, tt.args.startStr, tt.args.endStr, tt.args.layout)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsTimeInRangeStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsTimeInRangeStr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomTimeInRange(t *testing.T) {
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name        string
		args        args
		expectError bool
		validateFn  func(got time.Time, start time.Time, end time.Time) bool // 自定义校验逻辑
	}{
		{
			name: "valid range",
			args: args{
				start: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				end:   time.Date(2020, 1, 1, 23, 59, 59, 0, time.Local),
			},
			expectError: false,
			validateFn: func(got time.Time, start time.Time, end time.Time) bool {
				return !got.Before(start) && !got.After(end)
			},
		},
		{
			name: "different hours valid range",
			args: args{
				start: time.Date(2020, 6, 15, 8, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 6, 15, 8, 30, 0, 0, time.UTC),
			},
			expectError: false,
			validateFn: func(got time.Time, start time.Time, end time.Time) bool {
				return !got.Before(start) && !got.After(end)
			},
		},
		{
			name: "same start and end",
			args: args{
				start: time.Date(2020, 6, 15, 13, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 6, 15, 13, 0, 0, 0, time.UTC),
			},
			expectError: false,
			validateFn: func(got time.Time, start time.Time, end time.Time) bool {
				return got.Equal(start)
			},
		},
		{
			name: "invalid range start after end",
			args: args{
				start: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
				end:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			expectError: true,
			validateFn:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomTimeInRange(tt.args.start, tt.args.end)
			if (err != nil) != tt.expectError {
				t.Errorf("RandomTimeInRange() error = %v, expectError %v", err, tt.expectError)
				return
			}

			// 只有在没有预期错误时才进行后续操作
			if !tt.expectError {
				t.Logf("Generated time: %v", got)
				if got != nil {
					fmt.Println(got.Format("15:04:05"))
				}

				// 执行自定义验证函数
				if tt.validateFn != nil && got != nil {
					if !tt.validateFn(*got, tt.args.start, tt.args.end) {
						t.Errorf("RandomTimeInRange() got = %v, which does not satisfy validation criteria for [%v, %v]", got, tt.args.start, tt.args.end)
					}
				}
			}
		})
	}
}
