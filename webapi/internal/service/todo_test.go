package service

import (
	"context"
	"github.com/tingwei628/pgo/webapi/internal/entity"
	"reflect"
	"strings"
	"testing"
)

type MockDB struct {
	items []entity.Item
}

func (m *MockDB) InsertItem(ctx context.Context, item entity.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDB) GetAllItems(ctx context.Context) ([]entity.Item, error) {
	return m.items, nil
}
func (m *MockDB) GetItemsByKeyword(ctx context.Context, keyword string) ([]entity.Item, error) {
	var result []entity.Item
	for _, item := range m.items {
		if strings.Contains(strings.ToLower(item.Task), strings.ToLower(keyword)) {
			result = append(result, item)
		}
	}
	return result, nil
}

func TestTodoService_Search(t *testing.T) {
	tests := []struct {
		name           string
		toDosToAdd     []entity.Item // inputs
		query          string        // keyword
		expectedResult []entity.Item
	}{
		{
			name:           "t1",
			toDosToAdd:     []entity.Item{{Task: "c", Status: Completed}},
			query:          "c",
			expectedResult: []entity.Item{{Task: "c", Status: Completed}},
		},
		{
			name:           "t2",
			toDosToAdd:     []entity.Item{{Task: "d", Status: Completed}},
			query:          "d",
			expectedResult: []entity.Item{{Task: "d", Status: Completed}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &MockDB{}
			service := NewTodoService(mockDB)

			// setup test input
			for _, toAdd := range tt.toDosToAdd {
				err := service.Add(toAdd)
				if err != nil {
					t.Error(err)
				}
			}
			if got, _ := service.Search(tt.query); !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
