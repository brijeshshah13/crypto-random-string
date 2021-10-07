package cryptorandomstring

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Generator
	}{
		{
			name: "Generator",
			want: &Generator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCharacters(t *testing.T) {
	defaultGenerator = New() // this is to garbage values set to defaultGenerator by previous tests
	type args struct {
		characters string
	}
	tests := []struct {
		name string
		args args
		want *Generator
	}{
		{
			name: "GeneratorWithCharacters",
			args: args{characters: "abc"},
			want: &Generator{characters: "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithCharacters(tt.args.characters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithKind(t *testing.T) {
	type args struct {
		kind string
	}
	tests := []struct {
		name string
		args args
		want *Generator
	}{
		{
			name: "GeneratorWithKind",
			args: args{kind: "numeric"},
			want: &Generator{kind: "numeric", characters: "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithKind(tt.args.kind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLength(t *testing.T) {
	type args struct {
		length uint64
	}
	tests := []struct {
		name string
		args args
		want *Generator
	}{
		{
			name: "GeneratorWithLength",
			args: args{length: 10},
			want: &Generator{length: 10, kind: "numeric", characters: "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLength(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
