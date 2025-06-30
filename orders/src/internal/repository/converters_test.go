package repository

import (
	"github.com/BernsteinMondy/goorder/orders/src/internal/domain"
	"reflect"
	"testing"
)

func TestOrderStatusScanner_Scan(t *testing.T) {
	type args struct {
		src interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    domain.OrderStatus
		wantErr bool
	}{
		{
			name: "\"created\" string argument results in domain.OrderStatusCreated",
			want: domain.OrderStatusCreated,
			args: args{
				src: "created",
			},
		},
		{
			name: "\"cancelled\" string argument results in domain.OrderStatusCancelled",
			want: domain.OrderStatusCancelled,
			args: args{
				src: "cancelled",
			},
		},
		{
			name: "\"success\" string argument results in domain.OrderStatusSuccess",
			want: domain.OrderStatusSuccess,
			args: args{
				src: "success",
			},
		},
		{
			name: "\"payed\" string argument results in domain.OrderStatusPayed",
			want: domain.OrderStatusPayed,
			args: args{
				src: "payed",
			},
		},
		{
			name: "\"created\" []byte argument results in domain.OrderStatusCreated",
			want: domain.OrderStatusCreated,
			args: args{
				src: []byte("created"),
			},
		},
		{
			name: "\"cancelled\" []byte argument results in domain.OrderStatusCancelled",
			want: domain.OrderStatusCancelled,
			args: args{
				src: []byte("cancelled"),
			},
		},
		{
			name: "\"success\" []byte argument results in domain.OrderStatusSuccess",
			want: domain.OrderStatusSuccess,
			args: args{
				src: []byte("success"),
			},
		},
		{
			name: "\"payed\" []byte argument results in domain.OrderStatusPayed",
			want: domain.OrderStatusPayed,
			args: args{
				src: []byte("payed")},
		},
		{
			name: "incorrect string argument results in a non-nil error",
			args: args{
				src: "incorrect",
			},
			wantErr: true,
		},
		{
			name: "incorrect []byte argument results in a non-nil error",
			args: args{
				src: []byte("incorrect"),
			},
			wantErr: true,
		},
		{
			name: "empty string argument results in a non-nil error",
			args: args{
				src: "",
			},
			wantErr: true,
		},
		{
			name: "empty []byte argument results in a non-nil error",
			args: args{
				src: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "nil argument results in a non-nil error",
			args: args{
				src: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &OrderStatusScanner{}
			err := c.Scan(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v", err)
			}
			if got := c.Status; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderStatusToSQLEnum(t *testing.T) {
	type args struct {
		status domain.OrderStatus
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantPanic bool
	}{
		{
			name: "\"domain.OrderStatusCreated\" results in \"created\" string",
			args: args{
				status: domain.OrderStatusCreated,
			},
			want: "created",
		},
		{
			name: "\"domain.OrderStatusCancelled\" results in \"cancelled\" string",
			args: args{
				status: domain.OrderStatusCancelled,
			},
			want: "cancelled",
		},
		{
			name: "\"domain.OrderStatusPayed\" results in \"payed\" string",
			args: args{
				status: domain.OrderStatusPayed,
			},
			want: "payed",
		},
		{
			name: "\"domain.OrderStatusSuccess\" results in \"success\" string",
			args: args{
				status: domain.OrderStatusSuccess,
			},
			want: "success",
		},
		{
			name: "Unknown domain.OrderStatus results in panic",
			args: args{
				status: domain.OrderStatus(0),
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if tt.wantPanic {
					if r := recover(); r == nil {
						t.Errorf("orderStatusToSQLEnum() expected panic, but did not panic")
					}
				}
			}()

			if got := orderStatusToSQLEnum(tt.args.status); got != tt.want {
				t.Errorf("orderStatusToSQLEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductSizeScanner_Scan(t *testing.T) {
	type args struct {
		src interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    domain.ProductSize
		wantErr bool
	}{
		{
			name: "\"small\" string argument results in domain.ProductSizeSmall",
			want: domain.ProductSizeSmall,
			args: args{
				src: "small",
			},
		},
		{
			name: "\"medium\" string argument results in domain.ProductSizeMedium",
			want: domain.ProductSizeMedium,
			args: args{
				src: "medium",
			},
		},
		{
			name: "\"large\" string argument results in domain.ProductSizeLarge",
			want: domain.ProductSizeLarge,
			args: args{
				src: "large",
			},
		},
		{
			name: "\"extra_large\" string argument results in domain.ProductSizeExtraLarge",
			want: domain.ProductSizeExtraLarge,
			args: args{
				src: "extra_large",
			},
		},
		{
			name: "\"small\" []byte argument results in domain.ProductSizeSmall",
			want: domain.ProductSizeSmall,
			args: args{
				src: []byte("small"),
			},
		},
		{
			name: "\"medium\" []byte argument results in domain.ProductSizeMedium",
			want: domain.ProductSizeMedium,
			args: args{
				src: []byte("medium"),
			},
		},
		{
			name: "\"large\" []byte argument results in domain.ProductSizeLarge",
			want: domain.ProductSizeLarge,
			args: args{
				src: []byte("large"),
			},
		},
		{
			name: "\"extra_large\" []byte argument results in domain.ProductSizeExtraLarge",
			want: domain.ProductSizeExtraLarge,
			args: args{
				src: []byte("extra_large"),
			},
		},
		{
			name: "incorrect string argument results in a non-nil error",
			args: args{
				src: "incorrect",
			},
			wantErr: true,
		},
		{
			name: "incorrect []byte argument results in a non-nil error",
			args: args{
				src: []byte("incorrect"),
			},
			wantErr: true,
		},
		{
			name: "empty string argument results in a non-nil error",
			args: args{
				src: "",
			},
			wantErr: true,
		},
		{
			name: "empty []byte argument results in a non-nil error",
			args: args{
				src: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "nil argument results in a non-nil error",
			args: args{
				src: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ProductSizeScanner{}
			err := c.Scan(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v", err)
			}
			if got := c.Size; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productSizeToSQLEnum(t *testing.T) {
	type args struct {
		size domain.ProductSize
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantPanic bool
	}{
		{
			name: "\"domain.ProductSizeLarge\" results in \"large\" string",
			args: args{
				size: domain.ProductSizeLarge,
			},
			want: "large",
		},
		{
			name: "\"domain.ProductSizeMedium\" results in \"medium\" string",
			args: args{
				size: domain.ProductSizeMedium,
			},
			want: "medium",
		},
		{
			name: "\"domain.ProductSizeSmall\" results in \"small\" string",
			args: args{
				size: domain.ProductSizeSmall,
			},
			want: "small",
		},
		{
			name: "\"domain.ProductSizeExtraLarge\" results in \"extra_large\" string",
			args: args{
				size: domain.ProductSizeExtraLarge,
			},
			want: "extra_large",
		},
		{
			name: "Unknown domain.ProductSize results in panic",
			args: args{
				size: domain.ProductSize(0),
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if tt.wantPanic {
					if r := recover(); r == nil {
						t.Errorf("productSizeToSQLEnum() expected panic, but did not panic")
					}
				}
			}()

			if got := productSizeToSQLEnum(tt.args.size); got != tt.want {
				t.Errorf("productSizeToSQLEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}
