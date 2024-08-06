package store

import (
	"testing"
	"todo/models"
	"todo/test"
)

func getSimpleTask() models.Task {
	return models.Task{
		Description: "do the dishes",
		Complete:    false,
	}
}

func TestMemStore_Add(t *testing.T) {
	type fields struct {
		list map[string]models.Task
	}

	tests := []struct {
		name   string
		fields fields
		arg    models.Task
		want   int
	}{
		{
			name: "Add to empty map",
			fields: fields{
				map[string]models.Task{},
			},
			arg:  getSimpleTask(),
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}
			err := m.Add(tt.arg)

			test.AssertNil(t, err)
			test.AssertEquals(t, len(tt.fields.list), tt.want)
		})
	}
}

func TestMemStore_Get(t *testing.T) {
	type fields struct {
		list map[string]models.Task
	}

	tests := []struct {
		name    string
		fields  fields
		arg     string
		want    models.Task
		wantErr error
	}{
		{
			name: "Found task 1",
			fields: fields{
				map[string]models.Task{
					"1": getSimpleTask(),
				},
			},
			arg:  "1",
			want: getSimpleTask(),
		},
		{
			name: "Not found task 2",
			fields: fields{
				map[string]models.Task{},
			},
			arg:     "2",
			want:    models.Task{},
			wantErr: NotFoundErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}

			got, err := m.Get(tt.arg)

			test.AssertEquals(t, err, tt.wantErr)
			test.AssertEquals(t, got, tt.want)
		})
	}
}

func TestMemStore_List(t *testing.T) {
	type fields struct {
		list map[string]models.Task
	}

	tests := []struct {
		name   string
		fields fields
		arg    string
		want   map[string]models.Task
	}{
		{
			name: "List tasks",
			fields: fields{
				map[string]models.Task{
					"1": getSimpleTask(),
				},
			},
			arg: "1",
			want: map[string]models.Task{
				"1": getSimpleTask(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}

			got, err := m.List()

			test.AssertNil(t, err)
			test.AssertDeepEquals(t, got, tt.want)
		})
	}
}

func TestMemStore_Remove(t *testing.T) {
	type fields struct {
		list map[string]models.Task
	}

	tests := []struct {
		name    string
		fields  fields
		arg     string
		want    int
		wantErr error
	}{
		{
			name: "Remove task",
			fields: fields{
				map[string]models.Task{
					"1": getSimpleTask(),
				},
			},
			arg:  "1",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}

			err := m.Remove(tt.arg)

			test.AssertNil(t, err)
			test.AssertEquals(t, len(m.list), tt.want)
		})
	}
}

func TestMemStore_Update(t *testing.T) {
	type fields struct {
		list map[string]models.Task
	}

	tests := []struct {
		name    string
		fields  fields
		arg     models.Task
		want    int
		wantErr error
	}{
		{
			name: "Update task",
			fields: fields{
				map[string]models.Task{
					"1": getSimpleTask(),
				},
			},
			arg: models.Task{
				Description: "do the dishes",
				Complete:    true,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Task not found",
			fields: fields{
				map[string]models.Task{
					"2": getSimpleTask(),
				},
			},
			arg: models.Task{
				Description: "do the dishes",
				Complete:    true,
			},
			want:    1,
			wantErr: NotFoundErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}

			err := m.Update("1", tt.arg)

			test.AssertEquals(t, err, tt.wantErr)
			test.AssertEquals(t, len(m.list), tt.want)
		})
	}
}
