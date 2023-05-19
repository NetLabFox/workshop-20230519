package main

import "testing"

func Test_readEnv(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test01",
			want: "測試",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readEnv(); got != tt.want {
				t.Errorf("readEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkMain(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		readEnv()
	}
}
